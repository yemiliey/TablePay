package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Basic interface for anything that needs to handle signals and shutdown gracefully.
type runnable interface {
	Run() error
	Shutdown(context.Context) error
}

func signals() <-chan os.Signal {
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	return signals
}

func graceful(r runnable, timeout time.Duration) error {
	errC := make(chan error)
	go func() {
		defer close(errC)
		if err := r.Run(); err != nil {
			errC <- err
		}
	}()

	select {
	case err := <-errC:
		return errors.Wrap(err, "error during run")
	case <-signals():
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		if err := r.Shutdown(ctx); err != nil {
			return errors.Wrap(err, "error during shutdown")
		}
		return nil
	}
}

// run allows us to use defer and os.Exit at the same time.
func run(args []string) int {
	// Get config.
	cfg := settings.Load()

	// Configure logger.
	logger, err := platformlogger.NewFromSettings(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create logger: %s\n", err)
		return 1
	}
	defer logger.Close()
	logger.Log(hatchet.L{
		"message": "logger configured",
	})

	listenAddr := cfg.StringDflt("listen", ":80")
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		logger.Log(hatchet.L{
			"message": "failed to listen",
			"addr":    listenAddr,
			"error":   err,
		})
		return 1
	}

	// create pos data api client
	posDataAPIAddress, err := cfg.String("pos-data-api.endpoint")
	if err != nil {
		logger.Log(hatchet.L{
			"message": "error getting pos data api endpoint",
			"error":   err,
		})
	}
	contxt, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	dataAPIConn, err := grpc.DialContext(
		contxt,
		posDataAPIAddress,
		// block until the dial succeeds or fails
		grpc.WithBlock(),
		// non-tls
		grpc.WithInsecure(),
		grpc.WithBalancerName("round_robin"),
		// retry failed calls due to Unavailable with exponential backoff
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithBackoff(grpc_retry.BackoffExponentialWithJitter(100*time.Millisecond, 0.1)),
			grpc_retry.WithMax(10),
			grpc_retry.WithCodes(codes.Unavailable),
		)),
	)
	if err != nil {
		logger.Log(hatchet.L{
			"message": "error dialing pos data api",
			"error":   err,
		})
		return 1
	}
	defer dataAPIConn.Close()
	posDataAPIClient := servicepb.NewDataAPIClient(dataAPIConn)

	tablePayServer := server.New(listener, posDataAPIClient, logger)

	logger.Log(hatchet.L{"message": "service started"})
	logger.Log(hatchet.L{
		"addr": tablePayServer.Addr(),
	})

	if err := graceful(tablePayServer, time.Second*30); err != nil {
		msg := "table pay api server failed"
		fmt.Fprintf(os.Stderr, "%s: %s\n", msg, err)
		logger.Log(hatchet.L{
			"message": msg,
			"error":   err,
		})
		return 1
	}

	logger.Log(hatchet.L{"message": "sevice stopped"})

	return 0
}

func main() {
	os.Exit(run(os.Args))
}

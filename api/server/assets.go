package server

import (
	"github.com/goadesign/goa"
)

// AssetsController implements the assets resource.
type AssetsController struct {
	*goa.Controller
}

// NewAssetsController creates a assets controller.
func NewAssetsController(service *goa.Service) *AssetsController {
	return &AssetsController{Controller: service.NewController("AssetsController")}
}

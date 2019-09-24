import React, { Component } from 'react';
import {BrowserRouter, Route, Switch} from 'react-router-dom'
import Navbar from './components/Navbar'
import OrderViewPage from './components/OrderViewPage'
import HomePage from './components/HomePage'
// import Cart from './components/Cart'

class App extends Component {
  render() {
    return (
       <BrowserRouter>
            <div className="App">

              <Navbar/>
                <Switch>
                    <Route exact path="/orderview" component={OrderViewPage}/>
                    <Route exact path="/" component={HomePage}/>
                </Switch>
             </div>
       </BrowserRouter>

    );
  }
}

export default App;

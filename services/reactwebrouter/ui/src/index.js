import {Router, Reducer, View} from 'redux-director';

function Initialize(appname, ins, mod, settings){
  console.log("Initializing router")
  Application.Register("Reducers",'router', Reducer);
  Window.redirect = Router.redirect;
}

function connect(store) {
  Router.connect(store)
  Router.setRoutes(Application.Registry.Routes, 'home');
}

export {
  Initialize,
  View,
  connect
}

import {Router, Reducer} from 'redux-director';

function Initialize(appname, settings){
  console.log("Initializing router")
  Application.Register("Reducers",'router', Reducer)
}

function connect(store) {
  Router.connect(store)
  Router.setRoutes(Application.Registry.Routes, 'home');
}

export {
  Initialize,
  connect
}
import {} from 'redux-director';

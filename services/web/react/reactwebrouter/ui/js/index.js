import {Router, Reducer, View} from 'redux-director';

function Initialize(appname, ins, mod, settings){
  console.log("Initializing router")
  Application.Register("Reducers",'router', Reducer);
  Window.redirect = function(url, newpage) {
    if(newpage) {
      window.location.href = url;
    } else {
      Router.redirect(url);
    }
  }
}

function connect(store) {
  Router.connect(store)
  Router.setRoutes(Application.AllRegItems("Routes"), 'home');
}

export {
  Initialize,
  View,
  connect
}

import {Router, Reducer, View} from 'redux-director';
import {formatUrl} from 'uicommon'

function Initialize(appname, ins, mod, settings){
  console.log("Initializing router")
  Application.Register("Reducers",'router', Reducer);
  Window.redirect = function(url, newpage) {
    console.log("rediecting to url", url)
    if(newpage) {
      window.location.href = url;
    } else {
      Router.redirect(url);
    }
  }
  Window.redirectPage = function(pageName, params) {
    let page = _reg('Pages', pageName)
    console.log("rediecting to page", pageName)
    if(page) {
      let formattedUrl = formatUrl(page.route, params);
      Window.redirect(formattedUrl);
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

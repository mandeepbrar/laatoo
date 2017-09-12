import { createAction} from 'uicommon'
import {Actions} from './actions'
import { Provider } from 'react-redux';
import configureStore from './stores';

this.appname = 'application';
this.settings = {};

function Initialize(app, s, def, req) {
  //anonymous permissions
  if(!Storage.permissions) {
    Storage.permissions= this.settings.defaultPermissions;
    /*[
      "View Mehfil",
      "View Post",
      "View Comment",
      "View Media",
      "View Article"
    ]*/
  }
  if(s.application === app) {
    this.appname = app;
    this.settings = s;
  }
  this.req = req
}

function StartApplication() {
  console.log("Starting application ", this.appname, this.settings)
  let {router, uikit, theme} = this.settings;
  console.log("router", router, "uikit", uikit)
  const store = configureStore();
  console.log("store", store);

  createMessageDialogs(store)
  let Router = this.req(router)
  if(Router.default) {
    Router = Router.default
  }
  Router.connect(store);
  let Uikit = this.req(uikit)
  if(Uikit.default) {
    Uikit = Uikit.default
  }
  let ThemeMod = this.req(theme)
  if(ThemeMod.default) {
    ThemeMod = ThemeMod.default
  }
  let Theme = ThemeMod.Theme
  //let theme = React.createElement(settings.theme)
//  <theme/>
  Uikit.render(
    <Provider store={store}>
      <App uikit={Uikit}/>
    </Provider>, document.getElementById('app')
  );
}

function createMessageDialogs(store) {
  Window.showMessage = function(messageObj) {
    store.dispatch(createAction(Actions.SHOW_MESSAGE, {message: messageObj.Default}, null))
  }
  Window.showError = function(errObj, resp) {
    try {
      console.log("error response", resp, errObj)
      if(errObj) {
        store.dispatch(createAction(Actions.DISPLAY_ERROR, {message: errObj.Default}, null))
      } else {
        console.log("Error not found", errObj)
      }
    }catch(Ex) {
      console.log(Ex)
    }
  }
  Window.showDialog = function(title, component, actions, contentStyle) {
    store.dispatch(createAction(Actions.SHOW_DIALOG, {Title: title, Component: component, Actions: actions, ContentStyle: contentStyle}, null))
  }
  Window.closeDialog = function() {
    store.dispatch(createAction(Actions.CLOSE_DIALOG, {}, null))
  }
}


export {
  Initialize,
  StartApplication
}

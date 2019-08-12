import { createAction} from 'uicommon'
import Actions from './actions'
import { Provider } from 'react-redux';
import configureStore from './stores';
import React from 'react';
import {App} from './App';
import {ProcessPages} from 'reactpages';

var module;

function Initialize(app, ins, mod, s, def, req) {
  module = this;
  module.appname = 'application';
  module.settings = {};
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
  console.log("react application initialize", app, s, req)
  if(s.uiapplication === app) {
    this.appname = app;
    this.settings = s;
    if(s.Backend) {
      Application.Backend = s.Backend
    } else {
      Application.Backend = window.location.origin
    }
  }
  this.req = req
}

function StartApplication() {
  console.log("Starting application ", module.appname, module);
  let {router, uikit, theme} = module.settings;
  let Uikit = this.req(uikit)
  if(Uikit.default) {
    Uikit = Uikit.default
  }
  Application.setUikit(Uikit);
  console.log("theme for application", theme);
  let ThemeMod = this.req(theme)
  console.log("Theme mod", ThemeMod);
  if(ThemeMod.default) {
    ThemeMod = ThemeMod.default
  }
  if(ThemeMod.Start) {
    ThemeMod.Start(module.appname, Uikit);
  }
  ProcessPages(ThemeMod, Uikit);
  let Theme = ThemeMod.Theme
  //let theme = React.createElement(settings.theme)
//  <theme/>
  let Router = this.req(router)
  if(Router.default) {
    Router = Router.default
  }
  Application.setRouter(Router);
  const store = configureStore();
  createMessageDialogs(store)
  Router.connect(store);
  Uikit.render(
    <Provider store={store}>
      <App router={Router} theme={Theme}/>
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
  Window.handleError = function(errObj, resp) {
    Window.showError(errObj, resp)
  }
  Window.showDialog = function(title, component, onClose, actions, contentStyle, titleStyle) {
    store.dispatch(createAction(Actions.SHOW_DIALOG, {Title: title, Component: component, OnClose: onClose, Actions: actions, ContentStyle: contentStyle, TitleStyle: titleStyle}, null))
  }
  Window.closeDialog = function() {
    store.dispatch(createAction(Actions.CLOSE_DIALOG, {}, null))
  }
}


export {
  Initialize,
  StartApplication
}

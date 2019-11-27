import { createAction} from 'uicommon'
import Actions from './actions'
import Interactions from './components/Interactions'
import './reducers/Interactions'
import './sagas/LaatooActions'
import { Provider } from 'react-redux';
import configureStore from './stores';
import React from 'react';
import {App} from './App';
import {ProcessPages} from 'reactpages';
import "styles/app.scss";

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
  console.log("Starting application1 ", module.appname, module, this, this.req);
  let {router, uikit, theme} = module.settings;
  let Uikit = this.req(uikit)
  if(Uikit.default) {
    Uikit = Uikit.default
  }
  Application.setUikit(Uikit);
  console.log("theme for application", theme, this.req);
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
  Application.store = configureStore();
  createMessageDialogs(Application.store)
  Router.connect(Application.store);
  if(Application.Registry.Bootmethods) {
    for (let [methodName, method] of Object.entries(Application.AllRegItems("Bootmethods"))) {
      method(Application.store, Application, Uikit, Theme, Router)
    }
  }
  Uikit.render(
    <Provider store={Application.store}>
      <App router={Router} theme={Theme}/>
      <Interactions/>
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
  Window.showInteraction = function(interactiontype, title, component, onClose, actions, contentStyle, titleStyle) {
    store.dispatch(createAction(Actions.SHOW_INTERACTION_COMP, {title, component, onClose, actions, contentStyle, titleStyle, interactiontype}, null))
  }

  Window.closeInteraction = function(interactiontype) {
    store.dispatch(createAction(Actions.CLOSE_INTERACTION_COMP, {interactiontype: interactiontype}, null))
  }

  Window.executeAction = function(action, params) {
    store.dispatch(createAction(Actions.LAATOO_ACTION, params, action))    
  }

  Window.showDialog = function(title, component, onClose, actions, contentStyle, titleStyle) {
    Window.showInteraction("Dialog", title, component, onClose, actions, contentStyle, titleStyle)
  }
  Window.closeDialog = function() {
    Window.closeInteraction("Dialog")
  }
}


export {
  Initialize,
  StartApplication
}

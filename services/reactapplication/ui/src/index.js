import React from 'react';
import { Application, Window, Sagas, createAction, Storage } from 'reactuibase'
import {Actions} from './actions'
import { render } from 'react-dom';
import { Provider } from 'react-redux';
import configureStore from './stores';

function Initialize(appname, settings) {
  //anonymous permissions
  if(!Storage.permissions) {
    Storage.permissions= settings.defaultPermissions;
    /*[
      "View Mehfil",
      "View Post",
      "View Comment",
      "View Media",
      "View Article"
    ]*/
  }
  if(settings.application === appname) {
    Window.StartApplication = function() {
      const store = configureStore();
      createMessageDialogs()
      let router = require(settings.router)
      router.connect(store);
      let uikit = require(settings.uikit)
      let theme = React.createElement(settings.theme)
      uikit.render(
        <Provider store={store}>
          <uikit.Dialogs/>
          <theme/>
        </Provider>
      );
    }
  }
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


export default {
  Initialize
}

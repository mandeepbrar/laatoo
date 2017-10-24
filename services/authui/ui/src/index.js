import React from 'react';
import {LoginComponent} from './LoginComponent';
import {LoginValidator} from './LoginValidator';
import {renderWebLogin} from './WebLoginForm';
const PropTypes = require('prop-types');
import './reducers/Security';
import './sagas/Security';

var module = this;

function Initialize(appName, ins, mod, settings, def, req) {
  module.properties = Application.Properties[ins]
  module.settings = settings;
  if(Object.keys(settings).length != 0 ) {
    Application.Security = {
      googleAuthUrl: settings.googleAuthUrl,
      loginService: settings.loginService,
      validateService: settings.validateService,
      loginServiceURL: settings.loginServiceURL,
      realm: settings.realm
    }
  } else {
    Application.Security = {
      loginService: "login",
      validateService: "validate",
      realm: ""
    }
    let loginSvc = _reg("Services")
    if(!loginSvc) {
      Application.Register('Services', 'login', {url:"/login", method:'POST'})
      Application.Register('Services', 'validate', {url:"/validate", method:'POST'})
    }
  }
  if(settings.AuthToken) {
    Application.Security.AuthToken = settings.AuthToken
  } else {
    Application.Security.AuthToken = "x-auth-token"
  }
}



const WebLoginForm = (props, context) => {
  console.log("render logiform", LoginComponent)
  return (
    <LoginComponent className={props.className} renderLogin={renderWebLogin(context.uikit, module.settings, module.properties)} realm={props.realm} loginService={props.loginService}
      loginServiceURL={props.loginServiceURL} googleAuthUrl={props.googleAuthUrl}/>
  )
}

WebLoginForm.contextTypes = {
  uikit: PropTypes.object
};

export {
  Initialize,
  LoginComponent,
  renderWebLogin,
  WebLoginForm,
  LoginValidator
}

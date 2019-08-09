import React from 'react';
import {LoginComponent} from './LoginComponent';
import {LoginValidator} from './LoginValidator';
import {renderWebLogin} from './WebLoginForm';
import {renderSignup} from './SignupForm';
import {SignupComponent} from './SignupComponent';
const PropTypes = require('prop-types');
import './reducers/Security';
import './sagas/Security';
import {UserBlock} from './UserBlock';

var module;
function Initialize(appName, ins, mod, settings, def, req) {
  module=this;
  module.properties = Application.Properties[ins]
  console.log("authui initialization", Application, ins)
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
      signupService: "signup",
      validateService: "validate",
      realm: ""
    }
    let loginSvc = _reg("Services")
    if(!loginSvc) {
      Application.Register('Services', 'login', {url:"/login", method:'POST'})
      Application.Register('Services', 'signup', {url:"/register", method:'POST'})
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
    <LoginComponent className={props.className} renderLogin={renderWebLogin(module.settings, module.properties)} loginService={props.loginService}
      googleAuthUrl={props.googleAuthUrl}/>
  )
}


const SignupForm = (props, context) => {
  console.log("render signup form", SignupComponent)
  return (
    <SignupComponent className={props.className} renderSignup={renderSignup(module.settings, module.properties)} module={module}/>
  )
}


function userBlockDisplay(ctx, desc, className) {
  return (
    <UserBlock className={ctx.className} module={module}/>
  )
}

Application.Register('Blocks', "userBlock", userBlockDisplay)

export {
  Initialize,
  LoginComponent,
  renderWebLogin,
  renderSignup,
  SignupComponent,
  SignupForm,
  WebLoginForm,
  LoginValidator
}

import React from 'react';
import {LoginComponent} from './LoginComponent';
import {LoginValidator, setModule} from './LoginValidator';
import {renderWebLogin} from './WebLoginForm';
import {renderSignup} from './SignupForm';
import {SignupComponent} from './SignupComponent';
const PropTypes = require('prop-types');
import './reducers/Security';
import './sagas/Security';
import {UserBlock, LogoutButton} from './UserBlock';

var module;
function Initialize(appName, ins, mod, settings, def, req) {
  module=this;
  module.properties = Application.Properties[ins]
  console.log("authui initialization", Application, ins)
  module.settings = settings;
  setModule(module);
  Application.Security = Object.assign({
    loginService: "login",
    logoutService: "logout",
    signupService: "signup",
    validateService: "validate",
    AuthToken: "X-Auth-Token",
    realm: ""
  }, settings)
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
  LogoutButton,
  renderWebLogin,
  renderSignup,
  SignupComponent,
  SignupForm,
  WebLoginForm,
  LoginValidator
}

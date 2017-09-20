import React from 'react'
import {renderLogin} from './Login'
import {LoginComponent} from 'authui'
import './styles/app.scss'

const PropTypes = require('prop-types');

var module = this;

function Initialize(appName, ins, mod, settings, def, req) {
  module.properties = Application.Properties[ins]
  module.settings = settings;
  //Application.Register('Actions', 'loginAction', {actiontype: "method"})
  //Application.Register('Actions', 'googleAuth', {actiontype: "method"})
}

const LoginForm = (props, context) => {
  console.log("render logiform", LoginComponent)
  return (
    <LoginComponent className={props.className} renderLogin={renderLogin(context.uikit, module.settings, module.properties)} realm={props.realm} loginService={props.loginService}
      loginServiceURL={props.loginServiceURL} googleAuthUrl={props.googleAuthUrl}/>
  )
}

LoginForm.contextTypes = {
  uikit: PropTypes.object
};

export {
  Initialize,
  LoginForm
}

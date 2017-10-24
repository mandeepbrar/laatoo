import React from 'react'
import {Form} from './Form'
import './FormReducer'
import './sagas/Entity'
//import './styles/app.scss'

const PropTypes = require('prop-types');

var module = this;

function Initialize(appName, ins, mod, settings, def, req) {
  module.properties = Application.Properties[ins]
  module.settings = settings;
  //Application.Register('Actions', 'loginAction', {actiontype: "method"})
  //Application.Register('Actions', 'googleAuth', {actiontype: "method"})
}


export {
  Initialize,
  Form
}

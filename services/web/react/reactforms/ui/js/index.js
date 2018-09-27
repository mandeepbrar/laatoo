import React from 'react'
import {Form} from './Form'
import './FormReducer'
import './sagas/Entity'
import {Field, Initialize as InitializeFieldMod} from './Field';

import './styles/app.scss'

const PropTypes = require('prop-types');

var reactforms;

function Initialize(appName, ins, mod, settings, def, req) {
  reactforms = this;
  reactforms.properties = Application.Properties[ins];
  reactforms.settings = settings;
  InitializeFieldMod(appName, ins, mod, settings, def, req)
  //Application.Register('Actions', 'loginAction', {actiontype: "method"})
  //Application.Register('Actions', 'googleAuth', {actiontype: "method"})
}


export {
  Initialize,
  Form,
  Field
}
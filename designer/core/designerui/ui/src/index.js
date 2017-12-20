import React from 'react';
import {UsersView} from './pages/UsersView'
import './styles/app'

function Form_Instance_Transform_Modules(props) {
  console.log("transformer called", props)
  return props
}

function Initialize(appName, ins, mod, settings, def, req) {
  console.log("Initializing ui");
  _r("Methods", "Form_Instance_Transform_Modules",Form_Instance_Transform_Modules);
  console.log("registering method",Application);
}


export {
  Initialize,
  UsersView
}

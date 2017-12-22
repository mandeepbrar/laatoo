import React from 'react';
import {UsersView} from './pages/UsersView'
import './styles/app'

function Form_Instance_Transform_Modules(fieldProps, formValue, field, allfields, props, state,  form) {
  console.log("transformer called", props, state, fieldProps, field, allfields)
  console.log("parent value", props.parentFormValue)
  let modules = props.parentFormValue.Modules
  let items = []
  if(modules) {
    modules.forEach((module)=>{
      console.log("module.....", module)
      items.push({text: module.Name, value: module.Name})
    })
  }
  console.log("setting drop down", props, state, items)
  state.additionalProperties["items"] = items
  return fieldProps
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

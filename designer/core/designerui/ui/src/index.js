import React from 'react';
import {UsersView} from './pages/UsersView'
import './styles/app'
import Actions from './actions';
import './sagas'
import {  createAction } from 'uicommon';

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
  state.additionalProperties["items"] = items
  return fieldProps
}

function Form_SyncModules(form, submit, setData, dispatch) {
  let type=  form.config.entity.toLowerCase()
  console.log("form sync", type)
  return (data) => {
    console.log("data", data)
    dispatch(createAction(Actions.SYNC_OBJECTS, data, {type, setData}));
  }
}

function AbstractServer_Actions(form, submit, reset, uikit, setData, dispatch) {
  console.log("my action buttons", dispatch)
  return (
    <uikit.Block>
      <uikit.ActionButton onClick={submit()} className="submitBtn">Save</uikit.ActionButton>
      <uikit.ActionButton onClick={submit(Form_SyncModules(form, submit, setData, dispatch))} className="">Sync Modules</uikit.ActionButton>
    </uikit.Block>
  )
}

function Initialize(appName, ins, mod, settings, def, req) {
  console.log("Initializing ui");
  _r("Methods", "Form_Instance_Transform_Modules",Form_Instance_Transform_Modules);
  _r("Methods", "AbstractServer_Actions", AbstractServer_Actions);
  console.log("registering method",Application);
}


export {
  Initialize,
  UsersView
}

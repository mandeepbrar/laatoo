import React from 'react';
import {UsersView} from './pages/UsersView'
import './styles/app'
import Actions from './actions';
import './sagas'
import {  createAction } from 'uicommon';
import {RequestBuilder, DataSource, EntityData} from 'uicommon'


function Form_Instance_Modules(props, context, callback) {
  console.log("Form_Instance_Modules  called", props, context, callback)
  let modules = context.parentFormValue && context.parentFormValue.Modules? context.parentFormValue.Modules: []
  /*let items = []
  if(modules) {
    modules.forEach((module)=>{
      console.log("module.....", module)
      items.push({Name: module.Name, Id: module.Id})
    })
  }*/
  //state.additionalProperties["items"] = modules
  //return fieldProps
  callback(modules)
}

function Form_SyncModules(form, submit, setData, dispatch) {
  let type=  form.config.entity.toLowerCase()
  console.log("form sync", type)
  return (data) => {
    console.log("data", data)
    dispatch(createAction(Actions.SYNC_OBJECTS, data, {type, setData}));
  }
}

function AbstractServer_Available_Modules(callback) {
  return (resp) => {
    console.log("received entity data", resp);
    if(resp.data && resp.data.Modules) {
      callback(resp.data.Modules)
    }
  }
}

function AbstractServer_Solution_Modules(props, context, callback) {
  console.log("abstract server solution modules--------", props, context)
  let solution = context.formValue.Solution
  if(solution) {
    EntityData.GetEntity("Solution", solution).then(AbstractServer_Available_Modules(callback), (err)=>{console.log("Error in fetching solution modules", err)});
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

function ModulesRepo_ViewModule(params) {
  console.log("Params******** modulesrepo", params)
  params.ctx.panel.overlayComponent(<h2>my module</h2>)
}

function Initialize(appName, ins, mod, settings, def, req) {
  console.log("Initializing ui");
  _r("Methods", "Form_Instance_Modules",Form_Instance_Modules);
  _r("Methods", "AbstractServer_Actions", AbstractServer_Actions);
  _r("Methods", "AbstractServer_Solution_Modules", AbstractServer_Solution_Modules);
  _r("Actions", "ModulesRepo_viewModule", {actiontype: "method", method: ModulesRepo_ViewModule})
  console.log("registering method",Application);
}


export {
  Initialize,
  UsersView
}

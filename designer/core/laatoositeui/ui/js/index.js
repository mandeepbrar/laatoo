import React from 'react';
import './styles/app'
import {  createAction } from 'uicommon';
import {RequestBuilder, DataSource, EntityData} from 'uicommon';


function Initialize(appName, ins, mod, settings, def, req) {
  console.log("Initializing ui");
  _r("Methods", "demoActionButtons", demoActionButtons)
}

function demoActionButtons(form, submitFunc, reset, setData, dispatch) {
  return (
    <_uikit.Block>
      <_uikit.Action widget="button" name="googleAuth" className="googleAuthAction">Sign in with google</_uikit.Action>
      <_uikit.ActionButton onClick={submitFunc()} className="submitBtn">Sign up</_uikit.ActionButton>
    </_uikit.Block>
  )
}


export {
  Initialize
}

import React from 'react';
import './styles/app'
import {  createAction } from 'uicommon';
import {RequestBuilder, DataSource, EntityData} from 'uicommon';
import {OauthButton} from 'oauthui';

console.log("oauth ui ", OauthButton)

function Initialize(appName, ins, mod, settings, def, req) {
  console.log("Initializing ui");
  _r("Methods", "demoActionButtons", demoActionButtons)
}

function demoActionButtons(form, submitFunc, reset, setData, dispatch) {
  console.log("oauth button in demo actions", OauthButton)
  return (
    <_uikit.Block>
      <OauthButton className="googleAuthAction s10 btn-google"><i className="s10 fa fa-google"/>Google</OauthButton>
      <_uikit.ActionButton onClick={submitFunc()} className="submitBtn s10">Sign up</_uikit.ActionButton>
    </_uikit.Block>
  )
}


export {
  Initialize
}

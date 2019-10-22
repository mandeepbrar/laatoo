import React from 'react';
import './styles/app'
import {  createAction } from 'uicommon';
import {RequestBuilder, DataSource, EntityData} from 'uicommon';
import {OauthButton} from 'oauthui';

console.log("oauth ui ", OauthButton)

function Initialize(appName, ins, mod, settings, def, req) {
  console.log("Initializing ui");
  _r("Methods", "demoActionButtons", demoActionButtons)
  _r("Methods", "demoPreSubmit", demoPreSubmit);
}

function demoActionButtons(form, submitFunc, reset, setData, dispatch) {
  console.log("oauth button in demo actions", OauthButton)
  return (
    <_uikit.Block>
      <OauthButton className="googleAuthAction s10 bg-white blue"><_uikit.Image src="images/google-icon.png" className="s10"/>Google</OauthButton>
      <_uikit.ActionButton onClick={submitFunc()} className="submitBtn s10">Sign up</_uikit.ActionButton>
    </_uikit.Block>
  )
}

function demoPreSubmit(data) {
  console.log(data)
  data["Username"] = data["Email"]
  return data
}


export {
  Initialize
}

import React from 'react';
import './styles/app'
import {  createAction } from 'uicommon';
import {RequestBuilder, DataSource, EntityData} from 'uicommon';


function Initialize(appName, ins, mod, settings, def, req) {
  console.log("Initializing ui");
  _r("Methods", "selectSolution", selectSolution)
}

function selectSolution(params) {
  console.log(params)
}

export {
  Initialize
}

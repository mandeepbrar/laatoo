import React from 'react';
import './styles/app'
import Designer from './designer'

function Initialize(appName, ins, mod, settings, def, req) {
  console.log("Initializing module designer", settings);
}


export {
  Initialize,
  Designer
}

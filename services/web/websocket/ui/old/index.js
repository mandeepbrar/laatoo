import React from 'react'
import {InitializeSocket} from './sagas'
const PropTypes = require('prop-types');

var module;
function Initialize(appName, ins, mod, settings, def, req) {
  module=this;
  module.properties = Application.Properties[ins]
  module.settings = settings;
  module.req = req;
  InitializeSocket(settings.url)
}

export {
    Initialize
}
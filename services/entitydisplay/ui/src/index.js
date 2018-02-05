import React from 'react';
const PropTypes = require('prop-types');


var module = this;

function Initialize(appName, ins, mod, settings, def, req) {
  module.properties = Application.Properties[ins]
  module.settings = settings;
  console.log(settings.objects)
  if(settings.objects) {
    settings.objects.forEach(function(obj) {
      console.log("object name", obj)
    })
  }
}

export {
  Initialize
}

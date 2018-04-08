import React from 'react';
import {  createAction } from 'uicommon';
const PropTypes = require('prop-types');
import {Action} from 'reactwebcommon';
import {Panel} from 'reactpages';
import {ItemDetailView} from 'itemdetailview';
import 'styles/app.scss'

function Initialize(appName, ins, mod, settings, def, req) {
}

class ModuleSelect extends React.Component {
  getItem = (view, x, i) => {
    console.log("xxxxxxx", x)
    return <div>acd</div>
  }
  render() {
    return (
      <ItemDetailView id="repositoryview" entityName="ModuleDefinition"></ItemDetailView>
    )
  }
}

/**<ItemDetailView module="itemdetailview" id="modulesrepo" entityName="ModuleDefinition"></ItemDetailView>*/


ModuleSelect.contextTypes = {
  uikit: PropTypes.object
};

export {
  Initialize,
  ModuleSelect
}

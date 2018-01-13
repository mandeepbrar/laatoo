import React from 'react';
import {  createAction } from 'uicommon';
const PropTypes = require('prop-types');
import {Action} from 'reactwebcommon';
import {Panel} from 'reactpages';

function Initialize(appName, ins, mod, settings, def, req) {
}

class ModuleSelect extends React.Component {
  getItem = (view, x, i) => {
    console.log("xxxxxxx", x)
    return <div>acd</div>
  }
  render() {
    return (
      <Panel getItem={this.getItem} description={{type:"view", id: "repositoryview"}} />
    )
  }
}

ModuleSelect.contextTypes = {
  uikit: PropTypes.object
};

export {
  Initialize,
  ModuleSelect
}

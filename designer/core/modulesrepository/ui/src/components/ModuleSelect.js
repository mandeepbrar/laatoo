import React from 'react';
import {  createAction } from 'uicommon';
const PropTypes = require('prop-types');
import {Action} from 'reactwebcommon';
import {Panel} from 'reactpages';
import {ItemDetailView} from 'itemdetailview';


class ModuleSelect extends React.Component {
  constructor(props) {
    super(props)
    this.view = React.createRef();
  }
  getItem = (view, x, i) => {
    var methods = view.methods;
    let select = ()=> {
      console.log("selected ", i)
      methods.itemSelectionChange(i, true)
    }
    let uikit = this.context.uikit;
    return <uikit.Block>
      <uikit.Block className="row center valigncenter ma10">
      {x.Name}
      </uikit.Block>
      <uikit.Block className="row m10">
        <Action className="p10" action={{actiontype:"method", method: select, params:{}}}>Select</Action>
        <Action className="p10" action={{actiontype:"method", method: methods.openDetail, params:{data: x, index: i}}}>Details</Action>
      </uikit.Block>
    </uikit.Block>
    return action
  }
  submit = () => {
    let items = this.view.current.selectedItems()
    this.props.description.add(items, null, null, true)
  }
  render() {
    let uikit = this.context.uikit;
    return (
      <uikit.Block>
        <uikit.Block className=" w100 right ">
          <Action widget="button" className="p10" action={{actiontype:"method", method: this.submit, params:{}}}>Submit</Action>
        </uikit.Block>
        <ItemDetailView id="repositoryview" viewRef={this.view} getItem={this.getItem} editable={true} entityName="ModuleDefinition"></ItemDetailView>
      </uikit.Block>
    )
  }
}

/**<ItemDetailView module="itemdetailview" id="modulesrepo" entityName="ModuleDefinition"></ItemDetailView>*/


ModuleSelect.contextTypes = {
  uikit: PropTypes.object
};

export {
    ModuleSelect
}

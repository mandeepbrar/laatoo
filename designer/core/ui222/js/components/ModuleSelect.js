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
    return <_uikit.Block>
      <_uikit.Block className="row center valigncenter ma10">
      {x.Name}
      </_uikit.Block>
      <_uikit.Block className="row m10">
        <Action className="p10" action={{actiontype:"method", method: select, params:{}}}>Select</Action>
        <Action className="p10" action={{actiontype:"method", method: methods.openDetail, params:{data: x, index: i}}}>Details</Action>
      </_uikit.Block>
    </_uikit.Block>
    return action
  }
  submit = () => {
    let items = this.view.current.selectedItems()
    this.props.description.add(items, null, null, true)
  }
  render() {
    return (
      <_uikit.Block>
        <_uikit.Block className=" w100 right ">
          <Action widget="button" className="p10" action={{actiontype:"method", method: this.submit, params:{}}}>Submit</Action>
        </_uikit.Block>
        <ItemDetailView id="repositoryview" viewRef={this.view} getItem={this.getItem} editable={true} entityName="ModuleDefinition"></ItemDetailView>
      </_uikit.Block>
    )
  }
}

export {
    ModuleSelect
}

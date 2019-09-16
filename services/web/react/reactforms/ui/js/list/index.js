import React from 'react'
import {LoadableComponent} from 'reactwebcommon'
import ItemList from './itemlist';
import { ListSummary } from './summary';
const PropTypes = require('prop-types');

class List extends LoadableComponent {
  constructor(props, ctx) {
    super(props)
    let value = props.value? props.value: [] 
    this.state = {value}

    let entityName = _tn(props.name, props.entity)
    let label = _tn(props.label, entityName)
    let itemLabel = _tn(props.itemLabel, label)

    this.widgetProps={
      name: props.name, onChange:this.change, overlayComponent: ctx.overlayComponent, mode:props.mode, label:label,
      listItem:props.listComponent, editComponent:props.editComponent, itemForm: props.itemForm, skipLabel:props.skipLabel,
      itemLabel:itemLabel, inline:props.inline, entity:props.entity, field:props.field, titleField:props.titleField
    }


    console.log("show list", props, this.widgetProps)
  }

  componentWillReceiveProps(nextProps) {
    console.log("componentWillReceiveProps  for list", nextProps)
    let value = nextProps.value? nextProps.value: [] 
    this.setState(Object.assign({}, this.state, {value}))
  }

  dataLoaded = (data) => {
    this.setState(Object.assign({}, this.state, {itemData: data}))
  }

  getLoadContext = () => {
    console.log("get load context called", this.context)
    let context = {formValue: this.context.getFormValue()}
    return context
  }

  change = (value, name, evt) => {
    console.log("changing list", value, this.props, name, evt)
    this.props.onChange(value, this.props.name, evt)
    this.setState(Object.assign({}, this.state, {value}))
  }

  render() {
    console.log("list render ", this.state)

    if(this.props.mode == "summary") {
      return <ListSummary  itemData={this.state.itemData} value={this.state.value} {...this.widgetProps}/>
    } else {
      //field should not be used in any of the forms items to keep it generic
      return <ItemList  itemData={this.state.itemData} value={this.state.value} {...this.widgetProps}/>
    }
  }
}

List.contextTypes = {
  getFormValue: PropTypes.func,
  overlayComponent: PropTypes.func
};

export {
  List as List
}

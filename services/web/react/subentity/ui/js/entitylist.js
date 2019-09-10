import React from 'react'
import {Action, LoadableComponent} from 'reactwebcommon'
import Item from './item';
import {Panel} from 'reactpages'
import SelectEntity from './selectentity'
const PropTypes = require('prop-types');

class EntityListField extends React.Component {
  constructor(props, ctx) {
    super(props)
    console.log("entity list field ", props)
    this.state = {value: props.value, selectOptions: props.selectOptions, formOpen:false, inlineComp: null}
  }

  componentWillReceiveProps(nextProps) {
    console.log("entity list field : componentWillReceiveProps", nextProps)
    let st={}
    if(this.state.value != nextProps.value) {
      st.value = nextProps.value
    }
    if(this.state.selectOptions != nextProps.selectOptions) {
      st.selectOptions = nextProps.selectOptions
    }
    if(Object.keys(st).length >0) {
      this.setState(Object.assign({}, this.state, st))
    }
  }

  closeForm = () => {
    console.log("close form", this.props.field.mode)
    switch(this.props.field.mode) {
      /*case "select":
      case "inline":
        this.setInlineComp(null)
        break;*/
      case "dialog":
        Window.closeDialog()
        break;
      case "overlay":
      default:
        if(this.props.overlayComponent) {
          this.props.overlayComponent(null)
        }
    }
    this.setState(Object.assign({}, this.state, {formOpen: false, inlineComp: null}))
  }

  submitAddData = (data, success, failure, multipleItems) => {
      console.log("adding subentity ", data)
      let items = this.state.value.slice();
      if(multipleItems && data && Array.isArray(data)) {
        data.forEach(function(k) {
          items.push(k)
        })
      } else {
        items.push(data)
      }
      console.log(" items in add", items, data, this.state)
      this.props.onChange(items)
      this.closeForm()
  }

  removeItem = (index) => {
    let items = this.state.value.slice();
    if (index > -1) {
      items.splice(index, 1);
    }
    this.props.onChange(items)
  }

  edit = (data, index) => { 
    let items = this.state.value;
    items[index] = data
    this.props.onChange(items)
    this.closeForm()
  }

  actions = (f, submit, reset)=> {
    console.log("actios returned", f, submit, reset)
    return (
      <_uikit.Block className="right p20">
        {!this.props.field.inline?
        <_uikit.ActionButton onClick={reset}>
        Reset
        </_uikit.ActionButton>
        :null}
        <_uikit.ActionButton onClick={submit}>
        Save
        </_uikit.ActionButton>
      </_uikit.Block>
    )
  }

  openForm = (formData, submit, title, comp) => {
    let fld = this.props.field
    if(!comp) {
      comp = <Panel actions={this.actions} inline={true} formData={formData} title={title} closePanel={this.closeForm} onSubmit={submit} description={this.props.formDesc} /> //, actions, contentStyle)
    }
    switch(fld.mode) {
      case "inline":
        this.setInlineComp(comp)
        break;
      case "select":
        this.setInlineComp(<SelectEntity fld={fld} submit={this.edit} items={this.state.selectOptions} entity={formData} close={this.closeForm}/>)
        break;
      case "dialog":
        console.log("show subentity dialog", comp)        
        Window.showDialog(title, comp, this.closeForm)
        break;
      case "overlay":
      default:
        if(this.props.overlayComponent) {
          this.props.overlayComponent(comp)
        }
    }
    this.setState(Object.assign({}, this.state, {formOpen: true}))  
  }

  addItem = () => {
    let fld = this.props.field
    let title = "Add "+this.props.label
    let comp = fld.addwidget?
      <Panel title={title} description={{type:"component", componentName: fld.addwidget, module:fld.addwidgetmodule, add: this.props.submitAddData}} closePanel={this.closeForm} />
      :null
    this.openForm(null, this.submitAddData, title, comp)
  }

  setInlineComp = (comp) => {
    this.setState(Object.assign({}, this.state, {inlineComp: comp}))
  }

  getFormValue = () => {
    let parentFormValue = this.props.getFormValue();
    console.log("parent form value", this.props, parentFormValue);
    return parentFormValue;
  }

  render() {
    let items = []
    console.log("rendering items in entity list", this.props, this.state)
    let comp = this
    let props = this.props
    this.state.value.forEach(function(k, index) {
      if(!k) { return; }
      items.push(
        <Item value={k} index={index} removeItem={comp.removeItem} field={props.field} edit={comp.edit} openForm={comp.openForm}/>
      )
    })
    if(this.state.inlineComp) {
      items.push(this.state.inlineComp)
    }
    if(items.length == 0) {
      items.push("No data")
    }
    console.log("subentity items ", items)
    let actions = [  <Action action={{actiontype: "method"}} className="p10" method={this.addItem}> <_uikit.Icons.NewIcon /> </Action>]
    return (
      <_uikit.Block className={" entitylistfield "} title={this.props.title} titleBarActions={actions}>
        {items}
      </_uikit.Block>
    )
  }
}

export default EntityListField
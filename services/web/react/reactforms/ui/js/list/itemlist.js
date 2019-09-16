import React from 'react'
import {Action, LoadableComponent} from 'reactwebcommon'
import {Form} from '../Form'
import {Field} from '../Field'
import Item from './listitem';
import {Panel} from 'reactpages'
const PropTypes = require('prop-types');

class ItemList extends React.Component {
  constructor(props, ctx) {
    super(props)
    console.log("list field ", props)
    let listItem = props.listItem
    if(listItem) {
      this.listItemComponent = _res(listItem.module, listItem.component)
    }

    let editComp = props.editComponent
    if(editComp) {
      this.editComponent = _res(editComp.module, editComp.component)
    } 

    this.state = {value: props.value, itemData: props.itemData, formOpen:false, editComp: null}
  }

  componentWillReceiveProps(nextProps) {
    console.log("item list field : componentWillReceiveProps", nextProps)
    this.setState({value: nextProps.value, itemData: nextProps.itemData})
  }

  closeForm = () => {
    console.log("close form in forms", this.props.mode)
    switch(this.props.mode) {
      case "dialog":
        Window.closeDialog()
        break;
      case "overlay":
      default:
        if(this.props.overlayComponent) {
          console.log("overlay component set to null")
          this.props.overlayComponent(null)
        }
      break;
    }
    console.log("closing form ")
    this.setState(Object.assign({}, this.state, {formOpen: false, editComp: null}))
  }

  submitAddData = (data, success, failure, multipleItems) => {
      console.log("adding entry to list ", data)
      let items = this.state.value.slice();
      if(multipleItems && data && Array.isArray(data)) {
        data.forEach(function(k) {
          items.push(k)
        })
      } else {
        items.push(data)
      }
      console.log(" items in add", items, data, this.state)
      this.props.onChange(items, this.props.name)
      this.closeForm()
  }

  removeItem = (index) => {
    let items = this.state.value.slice();
    if (index > -1) {
      items.splice(index, 1);
    }
    this.props.onChange(items)
  }

  submitEditData = (data, index) => { 
    console.log("editing item", data, index)
    let items = this.state.value;
    items[index] = data
    console.log("new items", items)
    this.props.onChange(items, this.props.name)
    this.closeForm()
  }

  actions = (customSubmit) => {
    let comp = this
    return (f, submit, reset)=> {
      console.log("actions returned", f, submit, reset)
      return (
        <_uikit.Block className=" right p20 ">
          {!comp.props.inline?
          <_uikit.ActionButton onClick={reset}>
          Reset
          </_uikit.ActionButton>
          :null}
          <_uikit.ActionButton onClick={submit(customSubmit)}>
          Save
          </_uikit.ActionButton>
        </_uikit.Block>
      )
    }
  }

  addItem = () => {
    let title = "Add " + this.props.itemLabel
    /*let comp = fld.addwidget?
      <Panel title={title} description={{type:"component", componentName: fld.addwidget, module:fld.addwidgetmodule, add: this.props.submitAddData}} closePanel={this.closeForm} />
      :null*/
    this.openForm(null, this.submitAddData, title)
  }

  getRenderedItems = () => {
    let items = []
    let comp = this
    let props = this.props
    let itemComp = this.listItemComponent? this.listItemComponent: Item

    this.state.value.forEach(function(k, index) {
      if(!k) { return; }
      let itemProps = {value:k, index:index, removeItem: comp.removeItem, field: props.field, submitEditData: comp.submitEditData, 
        titleField: props.titleField, submitAddData: comp.submitAddData , openForm: comp.openForm, mode: props.mode}
      items.push(_ce(itemComp, itemProps))
    })
    return items        
  }

  getEditComp = (formData, submit, title, index, props) => {
    console.log("get edit comp", formData, props, title)
    let compProps = {actions:this.actions(submit), formData:formData, index: index, removeItem: this.removeItem, field: props.field, entity: props.entity,
      titleField: props.titleField, onFormSubmit: submit, mode: props.mode, closePanel: this.closeForm, title: title, itemLabel: props.itemLabel}
    if(this.editComponent) {
      return _ce(this.editComponent, compProps)
    } else if(props.itemForm || props.entity) {
      let formName = props.itemForm? props.itemForm: "new_form_" + props.entity.toLowerCase()
      let formDesc = {type: "form", id: formName}
      return <Panel description={formDesc} {...compProps}/> //, actions, contentStyle)    
//      return <Entity {...compProps}/> //, actions, contentStyle)    
    } 
    //let formsubmit = (data)=> {console.log("submit text", data); submit(data, index)}
    let beforeValueSet = (data) => {
      return { value: data}
    }
    let preSubmit = (data) => {
      console.log("adding item to list", "data")
      return _tn(data["value"], null)
    }
    return (
      <Form form="list_item_add" onFormSubmit={submit} beforeValueSet={beforeValueSet} preSubmit={preSubmit} formData={formData}>
        <Field name="value" className={props.className + " w100"}/>    
      </Form>
    )
  }

  openForm = (formData, submit, title, index) => {
    let props = this.props
    let comp = this.getEditComp(formData, submit, title, index, props)
    console.log("opening form with mode", props.mode, comp)
    let editComp = null;
    switch(this.props.mode) {
      case "inline":
      case "panes":
        console.log("opening form", comp)
        editComp = comp
        break;
      case "dialog":
        console.log("show dialog for entry", comp)        
        Window.showDialog(title, comp, this.closeForm)
        break;
      case "overlay":
      default:
        if(this.props.overlayComponent) {
          this.props.overlayComponent(comp)
        }
    }
    this.setState(Object.assign({}, this.state, {formOpen: true, editComp: editComp}))  
  }

  render() {
    let props = this.props
    console.log("list render ", this.state, props)

    let formPane = null
    let items = this.getRenderedItems()
    let itemsBlockClass = " w100 "
    if(this.state.editComp) {
      if(props.mode == "panes") {
        formPane = this.state.editComp
        itemsBlockClass = " left w30 "
      } else {
        items.push(this.state.editComp)
      }
    }
    if(items.length == 0) {
      items.push("No data")
    }

    let actions = [  <Action action={{actiontype: "method"}} className="p10" method={this.addItem}> <_uikit.Icons.NewIcon /> </Action>]
    return (
      <_uikit.Block className={" list " + props.label} contentClass="row" title={props.label} titleBarActions={actions}>
        <_uikit.Block className={" items " + itemsBlockClass} >
          {items}
        </_uikit.Block>
        { formPane?
          <_uikit.Block className={" editpane right fdgrow "} >          
          {formPane}
          </_uikit.Block>
          :null
        }
      </_uikit.Block>
    )
  }
}

export default ItemList
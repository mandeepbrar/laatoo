import React from 'react'
import {Action} from 'reactwebcommon'
import {Panel} from 'reactpages'
const PropTypes = require('prop-types');

class EntityListField extends React.Component {
  constructor(props, ctx) {
    super(props)
    this.state = {value: props.value, formOpen:false}
    this.uikit = props.uikit
  }

  componentWillReceiveProps(nextProps) {
    this.setState(Object.assign({}, this.state, {value: nextProps.value}))
  }

  closeForm = () => {
    switch(this.props.field.mode) {
      case "inline":
        this.inlineRow = null
        break;
      case "dialog":
        Window.closeDialog()
        break;
      case "overlay":
      default:
        if(this.props.overlayComponent) {
          this.props.overlayComponent(null)
        }
    }
    this.setState(Object.assign({}, this.state, {formOpen: false}))
  }

  actions = (f, submit, reset)=> {
    console.log("actios returned", f, submit, reset)
    return (
      <this.uikit.Block className="right p20">
        {!this.props.field.inline?
        <this.uikit.ActionButton onClick={reset}>
        Reset
        </this.uikit.ActionButton>
        :null}
        <this.uikit.ActionButton onClick={submit}>
        Save
        </this.uikit.ActionButton>
      </this.uikit.Block>
    )
  }

  edit = (data, index, success, failure) => {
    let items = this.state.value
    if(items && items.length > index) {
      items[index] = data.data
      console.log(" items", items)
      this.props.onChange(items)
      this.closeForm()
    }
  }

  add = (data, success, failure) => {
      let items = this.state.value.slice();
      items.push(data.data)
      console.log(" items", items)
      this.props.onChange(items)
      this.closeForm()
  }

  removeItem = (item, index) => {
    let items = this.state.value.slice();
    if (index > -1) {
      items.splice(index, 1);
    }
    this.props.onChange(items)
  }

  getFormValue = () => {
    let parentFormValue = this.props.getFormValue();
    console.log("parent form value", parentFormValue);
    return parentFormValue;
  }

  openForm = (formData, index) => {
    console.log("opened form", this.props, this.context)
    let cl = this;
    let submit = formData? (data, success, failure)=>{return cl.edit(data, index, success, failure)}: this.add
    let comp = <Panel actions={this.actions} inline={true} formData={formData} title={"Add "+this.props.label} parent={this} subform={true} closePanel={this.closeForm} onSubmit={submit} description={this.props.formDesc} /> //, actions, contentStyle)
    switch(this.props.field.mode) {
      case "inline":
        this.inlineRow = comp
        break;
      case "dialog":
        Window.showDialog(null, comp, this.closeForm)
        break;
      case "overlay":
      default:
        if(this.props.overlayComponent) {
          this.props.overlayComponent(comp)
        }
    }
    this.setState(Object.assign({}, this.state, {formOpen: true}))
  }

  render() {
    let items = []
    console.log("rendering items in entity list", this.props, this.state)
    let comp = this
    this.state.value.forEach(function(k, index) {
      var removeItem = () => {
        comp.removeItem(k, index)
      }
      var editItem = () => {
        comp.openForm(k, index)
      }
      let textField = comp.props.entityText? comp.props.entityText: "Name"
      let text = k[textField];
      text = text? text: k["Title"]
      items.push(
        <comp.uikit.Block  className="row between-xs">
          <comp.uikit.Block className="left" >
          {text}
          </comp.uikit.Block>
          <comp.uikit.Block className="right">
            <Action action={{actiontype:"method"}} className="edit p10" method={editItem}>
              <comp.uikit.Icons.EditIcon />
            </Action>
            <Action action={{actiontype:"method"}} className="remove p10" method={removeItem}>
              <comp.uikit.Icons.DeleteIcon />
            </Action>
          </comp.uikit.Block>
        </comp.uikit.Block>
      )
    })
    let inlinerow = null
    if(this.state.formOpen && this.inlineRow) {
      items.push(this.inlineRow)
    }
    if(items.length == 0) {
      items.push("No data")
    }
    let actions = [  <Action action={{actiontype: "method"}} className="p10" method={this.openForm}> <this.uikit.Icons.NewIcon /> </Action>]
    return (
      <this.uikit.Block className={"entitylistfield "} titleBarActions={actions}>
        {items}
      </this.uikit.Block>
    )
  }
}

class SubEntity extends React.Component {
  constructor(props, ctx) {
    super(props)
    this.list = props.field.list? true: false
    this.label = props.field.label? props.field.label : props.field.entity
    let formName = props.field.form? props.field.form : "new_form_"+props.field.entity.toLowerCase()
    this.formDesc = {type: "form", id: formName}
    this.uikit = ctx.uikit;
    let value = props.input.value? props.input.value: (this.list? [] : {})
    this.state = {value}
  }

  componentWillReceiveProps(nextProps) {
    let value = nextProps.input.value? nextProps.input.value: (this.list? [] : {})
    this.setState(Object.assign({}, this.state, {value}))
  }

  change = (value) => {
    console.log("charnging subentity", value)
    this.props.input.onChange(value)
  }

  render() {
    let title = this.props.field.skipLabel? null: this.label
    return (
      <this.uikit.Block className={"subentity "+this.label} title={title}>
        {this.list?
          <EntityListField uikit={this.uikit} getFormValue={this.context.getFormValue} field={this.props.field} onChange={this.change} label={this.label}
          overlayComponent={this.context.overlayComponent} formDesc={this.formDesc} value={this.state.value}/>
        : <Panel actions={()=>{}} formData={this.state.value} onChange={this.change} trackChanges={true} description={this.formDesc} />
        }
      </this.uikit.Block>
    )
  }
}

SubEntity.contextTypes = {
  uikit: PropTypes.object,
  getFormValue: PropTypes.func,
  overlayComponent: PropTypes.func
};

export {
  SubEntity as SubEntity
}

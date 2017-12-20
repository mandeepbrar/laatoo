import React from 'react'
import {Action} from 'reactwebcommon'
import {Panel} from 'reactpages'
const PropTypes = require('prop-types');

class EntityListField extends React.Component {
  constructor(props, ctx) {
    super(props)
    let items = props.input.value? props.input.value: []
    this.label = props.field.label? props.field.label : props.field.entity
    let formName = props.field.form? props.field.form : "new_form_"+props.field.entity.toLowerCase()
    this.formDesc = {type: "form", id: formName}
    this.uikit = ctx.uikit;
    let formOpen = false
    this.state = {items, formOpen}
  }

  componentWillReceiveProps(nextProps) {
    let items = nextProps.input.value? nextProps.input.value: []
    this.setState(Object.assign({}, this.state, {items}))
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
        if(this.context.overlayComponent) {
          this.context.overlayComponent(null)
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
    let items = this.state.items
    if(items && items.length > index) {
      items[index] = data.data
      console.log(" items", items)
      this.props.input.onChange(items)
      this.closeForm()
    }
  }

  editItem = (item, index) => {
    this.openForm(item, index)
  }

  add = (data, success, failure) => {
      let items = this.addItem(data.data)
      console.log(" items", items)
      this.props.input.onChange(items)
      this.closeForm()
  }

  addItem = (item) => {
    let items = this.state.items.slice();
    items.push(item)
    return items
    //console.log("items set to ", items, item)
    //this.setState(Object.assign({}, {items: items}))
  }

  openForm = (formData, index) => {
    console.log("opened form", this.props, this.context)
    let cl = this;
    let submit = formData? (data, success, failure)=>{return cl.edit(data, index, success, failure)}: this.add
    let comp = <Panel actions={this.actions} inline={true} formData={formData} title={"Add "+this.label} closePanel={this.closeForm} onSubmit={submit} description={this.formDesc} /> //, actions, contentStyle)
    switch(this.props.field.mode) {
      case "inline":
        this.inlineRow = comp
        break;
      case "dialog":
        Window.showDialog(null, comp, this.closeForm)
        break;
      case "overlay":
      default:
        if(this.context.overlayComponent) {
          this.context.overlayComponent(comp)
        }
    }
    this.setState(Object.assign({}, this.state, {formOpen: true}))
  }


  removeItem = (item, index) => {
    let items = this.state.items.slice();
    if (index > -1) {
      items.splice(index, 1);
    }
    this.props.input.onChange(items)
  }

  render() {
    let items = []
    console.log("rendering items in entity list", this.props, this.state)
    let comp = this
    this.state.items.forEach(function(k, index) {
      var removeItem = () => {
        comp.removeItem(k, index)
      }
      var editItem = () => {
        comp.editItem(k, index)
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
    let title = this.props.field.skipLabel? null: this.label
    let actions = [  <Action name="listfield_new_entity" className="p10" method={this.openForm}> <this.uikit.Icons.NewIcon /> </Action>]
    return (
      <this.uikit.Block className={"entitylistfield "+this.label} title={title} titleBarActions={actions}>
        {items}
      </this.uikit.Block>
    )
  }
}

EntityListField.contextTypes = {
  uikit: PropTypes.object,
  overlayComponent: PropTypes.func
};

export {
  EntityListField as EntityList
}

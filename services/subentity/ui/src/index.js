import React from 'react'
import {Action, LoadableComponent} from 'reactwebcommon'
import {Panel} from 'reactpages'
const PropTypes = require('prop-types');

class SelectEntity extends React.Component {
  constructor(props) {
    super(props)
    this.state={value: props.value, items: props.items}
  }
  componentWillReceiveProps(nextProps) {
    console.log("on change of select entity--", this.state, nextProps)
    this.setState(Object.assign({}, this.state, {value: nextProps.value, items: nextProps.items}))
  }
  onChange=(value)=> {
    console.log("on change of select entity--", value)
    this.setState(Object.assign({}, this.state, {value}))
  }
  saveValue=()=> {
    console.log("svaing value", this.state, this.props)
    this.props.submit(this.state.value)
  }
  render() {
    console.log("rendering select entity", this.state)
    let {fld, uikit} = this.props
    let fldDesc = {label: fld.label, name: fld.name, widget: "Select", selectItem: true, type: "entity", items: this.state.items}
    return <uikit.Block  className="row between-xs">
      <uikit.Block className="left col-xs-10" >
        <uikit.Forms.FieldWidget className="w100" field={fldDesc} fieldChange={this.onChange} value={this.state.value}/>
      </uikit.Block>
      <uikit.Block className="right">
        <Action action={{actiontype:"method"}} className="edit p10" method={this.saveValue}>
          <uikit.Icons.EditIcon />
        </Action>
        <Action action={{actiontype:"method"}} className="remove p10" method={this.props.close}>
          <uikit.Icons.DeleteIcon />
        </Action>
      </uikit.Block>
    </uikit.Block>
  }
}

class EntityListField extends React.Component {
  constructor(props, ctx) {
    super(props)
    console.log("entity list field ", props)
    this.state = {value: props.value, formOpen:false, selectOptions: props.selectOptions}
    this.uikit = props.uikit
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
      case "select":
      case "inline":
        console.log("inline row close form")
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
      items[index] = data
      console.log(" items in edit", items[index], index, data, this.state)
      this.props.onChange(items)
      this.closeForm()
    }
  }

  add = (data, success, failure, multipleItems) => {
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

  removeItem = (item, index) => {
    let items = this.state.value.slice();
    if (index > -1) {
      items.splice(index, 1);
    }
    this.props.onChange(items)
  }

  getFormValue = () => {
    let parentFormValue = this.props.getFormValue();
    console.log("parent form value", this.props, parentFormValue);
    return parentFormValue;
  }

  openForm = (formData, index) => {
    console.log("opened form", this.props, this.context)
    let cl = this;
    let fld = this.props.field
    let submit = formData? (data, success, failure)=>{return cl.edit(data, index, success, failure)}: this.add
    let comp = fld.addwidget?
      <Panel title={"Add "+this.props.label} description={{type:"component", componentName: fld.addwidget, module:fld.addwidgetmodule, add: this.add}} parentFormRef={this} subform={true} closePanel={this.closeForm} autoSubmitOnChange={this.props.autoSubmitOnChange}/>
    : <Panel actions={this.actions} inline={true} formData={formData} title={"Add "+this.props.label} parentFormRef={this}  subform={true} closePanel={this.closeForm} onSubmit={submit} description={this.props.formDesc} autoSubmitOnChange={this.props.autoSubmitOnChange}/> //, actions, contentStyle)
    switch(this.props.field.mode) {
      case "inline":
        this.inlineRow = comp
        break;
      case "select":
        this.inlineRow = <SelectEntity fld={fld} uikit={this.uikit} submit={submit} items={this.state.selectOptions} entity={formData} index={index} close={this.closeForm}/>
        break;
      case "dialog":
        console.log("show subentity dialog", comp)
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
    let fld = this.props.field
    let comp = this
    this.state.value.forEach(function(k, index) {
      if(!k) { return; }
      var removeItem = () => {
        comp.removeItem(k, index)
      }
      var editItem = () => {
        comp.openForm(k, index)
      }
      console.log("entity list ", k, fld)
      let textField = fld.textField? fld.textField: "Name"
      let text = k[textField];
      text = text? text: k["Title"]
      console.log("entity text ", text, textField)
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
    console.log("subentity items ", items)
    let actions = [  <Action action={{actiontype: "method"}} className="p10" method={this.openForm}> <this.uikit.Icons.NewIcon /> </Action>]
    return (
      <this.uikit.Block className={"entitylistfield "} title={this.props.title} titleBarActions={actions}>
        {items}
      </this.uikit.Block>
    )
  }
}

class SubEntity extends LoadableComponent {
  constructor(props, ctx) {
    super(props)
    this.list = props.field.list? true: false
    this.label = props.field.label? props.field.label : props.field.entity
    let formName = props.field.form? props.field.form : "new_form_"+props.field.entity.toLowerCase()
    this.formDesc = {type: "form", id: formName}
    this.uikit = ctx.uikit;
    let value = props.input.value? props.input.value: (this.list? [] : {})
    this.state = {value}
    console.log("show subentity", this.formDesc, props, ctx)
  }

  componentWillReceiveProps(nextProps) {
    console.log("componentWillReceiveProps  for SubEntity", nextProps)
    let value = nextProps.input.value? nextProps.input.value: (this.list? [] : {})
    if(this.state.value != value) {
      this.setState(Object.assign({}, this.state, {value}))
    }
  }

  dataLoaded = (data) => {
    if(this.props.field.mode == "select") {
      console.log("data loaded for SubEntity", data)
      this.setState(Object.assign({}, this.state, {selectOptions: data}))
    }
  }

  getLoadContext = () => {
    console.log("get load context called", this.context)
    let context = {formValue: this.context.getFormValue()}
    if(this.context.getParentFormValue) {
      context.parentFormValue = this.context.getParentFormValue()
    }
    return context
  }

  selectSubEntity = () => {
    let fld = this.props.field
    let fldDesc = {label: fld.label, name: fld.name, widget: "Select", selectItem: true, type: "entity"}
    return <this.uikit.Forms.FieldWidget className="w100" field={fldDesc} fieldChange={this.change} items={this.state.selectOptions} value={this.state.value}/>
  }

  change = (value) => {
    console.log("charnging subentity", value, this.props)
    this.props.fieldChange(value, this.props.name)
    this.setState(Object.assign({}, this.state, {value}))
  }

  render() {
    console.log("subentity ", this.state, this.props)
    let field = this.props.field
    let title = field.skipLabel? null: this.label
    let autoSubmit = null
    return (
      <this.uikit.Block className={"subentity "+this.label}>
        {this.list?
        <EntityListField uikit={this.uikit} getFormValue={this.context.getFormValue} field={this.props.field} onChange={this.change} label={this.label} form={this.props.form} formRef={this.props.formRef} autoSubmitOnChange={this.props.autoSubmitOnChange}
          selectOptions= {this.state.selectOptions} overlayComponent={this.context.overlayComponent}  parentFormRef={this.props.parentFormRef} formDesc={this.formDesc} title={title} value={this.state.value}/>
        : ((field.mode=="select")?
        this.selectSubEntity()
        :
        <Panel actions={()=>{}} formData={this.state.value} title={title}  autoSubmitOnChange={true} onChange={this.change} trackChanges={true} subform={this.props.subform} formRef={this.props.formRef} parentFormRef={this.props.parentFormRef} description={this.formDesc} />)
        }
      </this.uikit.Block>
    )
  }
}

SubEntity.contextTypes = {
  uikit: PropTypes.object,
  getFormValue: PropTypes.func,
  getParentFormValue: PropTypes.func,
  overlayComponent: PropTypes.func
};

export {
  SubEntity as SubEntity
}

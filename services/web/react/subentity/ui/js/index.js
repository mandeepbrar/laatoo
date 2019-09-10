import React from 'react'
import {Action, LoadableComponent} from 'reactwebcommon'
import EntityList from './entitylist';
import {Panel} from 'reactpages'
const PropTypes = require('prop-types');

class SubEntity extends LoadableComponent {
  constructor(props, ctx) {
    super(props)
    let field = props.field
    this.list = field.list? true: false
    this.label = field.label? field.label : field.entity
    let formName = field.form? field.form : "new_form_"+field.entity.toLowerCase()
    this.formDesc = {type: "form", id: formName}
    let value = props.value? props.value: (this.list? [] : {})
    this.widgetMode = field.widget && field.widget.props? field.widget.props.mode: null;
    this.state = {value}
    console.log("show subentity", this.formDesc, props, ctx, this.state)
  }

  componentWillReceiveProps(nextProps) {
    console.log("componentWillReceiveProps  for SubEntity", nextProps)
    let value = nextProps.value? nextProps.value: (this.list? [] : {})
    if(this.state.value != value) {
      this.setState(Object.assign({}, this.state, {value}))
    }
  }

  dataLoaded = (data) => {
    if(this.widgetMode == "select") {
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
    let fldDesc = {label: fld.label, name: fld.name, widget: {"name": "Select"}, type: "entity"}
    return <_uikit.Field className="w100" field={fldDesc} onChange={this.change}  selectItem={true} items={this.state.selectOptions} value={this.state.value}/>
  }

  change = (value, name, evt) => {
    console.log("charnging subentity", value, this.props, name, evt)
    this.props.onChange(value, this.props.name, evt)
    this.setState(Object.assign({}, this.state, {value}))
  }

  render() {
    console.log("subentity ", this.state, this.props)
    let field = this.props.field
    let title = field.skipLabel? null: this.label
    let autoSubmit = null
    return (
      <_uikit.Block className={"subentity "+this.label}>
        {this.list?
        <EntityList field={field} onChange={this.change} label={this.label} formRef={this.props.formRef} selectOptions= {this.state.selectOptions} 
          overlayComponent={this.context.overlayComponent}  formDesc={this.formDesc} title={title} value={this.state.value}/>
        : ((this.widgetMode=="select")?
        this.selectSubEntity()
        :
        <Panel actions={()=>{}} formData={this.state.value} title={title} onChange={this.change} subform={true} formRef={this.props.formRef} description={this.formDesc} />)
        }
      </_uikit.Block>
    )
  }
}

SubEntity.contextTypes = {
  getFormValue: PropTypes.func,
  getParentFormValue: PropTypes.func,
  overlayComponent: PropTypes.func
};

export {
  SubEntity as SubEntity
}

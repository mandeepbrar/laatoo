/*'use strict';

import React from 'react';
import { Field } from 'redux-form'
const PropTypes = require('prop-types');
import {FldList} from './FldList';


class FieldWrapper extends React.Component {
  constructor(props, context) {
    super(props)
    let field = props.fields[props.name]
    this.field = field
    console.log("fields created", props, this.field, context)
    if(!this.field.label && !this.field.skipLabel) {
      this.field.label = props.name
    }
    let additionalProperties = {entity: field.entity, dataService: field.dataService, dataServiceParams: field.dataServiceParams, loader: field.loader, skipDataLoad: field.skipDataLoad}

    this.state = {time: props.time? props.time: Date.now(), additionalProperties}

    if(this.field.transformer) {
      let method = _reg("Methods", this.field.transformer)
      this.transformer = method;
    }
    if(this.field.type == "storableref") [
      this.isRef = true
    ]
    if(this.field.module) {
      this.fldWidget = _res(this.field.module, this.field.widget);
    }
  }

  componentWillReceiveProps(nextProps, nextState) {
    console.log("receive props field wrapper", nextProps, nextState)
    if(nextProps.time > this.state.time) {
      this.setState(Object.assign({}, this.state, {time: nextProps.time}))
    }
  }

  fieldChange = (fldProps) => {
    let comp=this
    return (data, name, evt)=> {
      console.log("fld change", data, name, evt, this.props, this.context, fldProps, fldProps.input.onChange, comp.isRef)
      if(fldProps.input.onChange) {
        if(comp.isRef) {
          let myRefObj = {}
          myRefObj[comp.field.name] = {"Id": data, "Type": comp.field.entity}
          data = myRefObj
          console.log("set ref value data", myRefObj, data)
        }
        console.log("setting fld value", comp, data, name)
        fldProps.input.onChange(data, name)
      }
    }
  }

  component = (fieldProps) => {
    console.log("component", this.state, fieldProps, this.props)
    let newProps = fieldProps
    if(this.transformer) {
      newProps = this.transformer(fieldProps, this.props.formValue, this.field, this.context.fields, this.props, this.state,  this)
    }
    if(this.isRef) {
      let ref = newProps.input.value[this.field.name]
      if(ref) {
        newProps.input.value = ref.Id
      }
      console.log("ref props", newProps)
    }
    let comp = null
    let baseComp = null
    if(this.fldWidget) {
      return <this.fldWidget name={this.props.name} className={this.props.className} {...this.state.additionalProperties} time={this.state.time}
          formValue={this.props.formValue} field={this.field} fieldChange={this.fieldChange(fieldProps)} subFormChange={this.props.subFormChange} autoSubmitOnChange={this.props.autoSubmitOnChange}
          subform={this.props.subform} formRef={this.props.formRef} parentFormRef={this.props.parentFormRef} {...newProps}/>
    } else {
      if(this.field.list) {
        return <FldList name={this.props.name} baseComponent={this.context.uikit.Forms.FieldWidget} className={this.props.className} ap={this.state.additionalProperties}
         time={this.state.time}  formValue={this.props.formValue} field={this.field}  fieldChange={this.fieldChange(fieldProps)} autoSubmitOnChange={this.props.autoSubmitOnChange}
         subFormChange={this.props.subFormChange} subform={this.props.subform} formRef={this.props.formRef} parentFormRef={this.props.parentFormRef} baseProps={newProps}/>
      } else {
        return <this.context.uikit.Forms.FieldWidget  name={this.props.name} className={this.props.className} {...this.state.additionalProperties}
          time={this.state.time}  formValue={this.props.formValue} field={this.field}  fieldChange={this.fieldChange(fieldProps)} subFormChange={this.props.subFormChange} autoSubmitOnChange={this.props.autoSubmitOnChange}
          subform={this.props.subform} formRef={this.props.formRef} parentFormRef={this.props.parentFormRef} {...newProps}/>
      }
    }
  }

  render() {
    console.log("changing state", this.state)
    return (
      <Field key={this.props.name} name={this.props.name} time={this.state.time} component={this.component}/>
    )
  }

}

FieldWrapper.contextTypes = {
  fields: PropTypes.object,
  uikit:  PropTypes.object
};

export { FieldWrapper as Field}
*/
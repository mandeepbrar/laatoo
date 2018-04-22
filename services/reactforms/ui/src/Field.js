'use strict';

import React from 'react';
import { Field } from 'redux-form'
const PropTypes = require('prop-types');
import {FldList} from './FldList';

var modrequire = null;

function Initialize(appName, ins, mod, settings, def, req) {
  modrequire = req;
}

class FieldWrapper extends React.Component {
  constructor(props, context) {
    super(props)
    let field = context.fields[props.name]
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
    if(this.field.module) {
      let mod = modrequire(this.field.module);
      this.fldWidget = mod[this.field.widget];
    }
  }

  componentWillReceiveProps(nextProps, nextState) {
    if(nextProps.time > this.state.time) {
      this.setState(Object.assign({}, this.state, {time: nextProps.time}))
    }
  }

  component = (fieldProps) => {
    console.log("component", this.state, fieldProps)
    let newProps = fieldProps
    if(this.transformer) {
      newProps = this.transformer(fieldProps, this.props.formValue, this.field, this.context.fields, this.props, this.state,  this)
    }
    let comp = null
    let baseComp = null
    if(this.fldWidget) {
      return <this.fldWidget name={this.props.name} className={this.props.className} {...this.state.additionalProperties} time={this.state.time} formValue={this.props.formValue} field={this.field} {...newProps}/>
    } else {
      if(this.field.list) {
        return <FldList name={this.props.name} baseComponent={this.context.uikit.Forms.FieldWidget} className={this.props.className} ap={this.state.additionalProperties} time={this.state.time}  formValue={this.props.formValue} field={this.field} baseProps={newProps}/>
      } else {
        return <this.context.uikit.Forms.FieldWidget  name={this.props.name} className={this.props.className} {...this.state.additionalProperties} time={this.state.time}  formValue={this.props.formValue} field={this.field} {...newProps}/>
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

export { FieldWrapper as Field, Initialize}

'use strict';

import React from 'react';
import { Field } from 'redux-form'
const PropTypes = require('prop-types');

class FieldWrapper extends React.Component {
  constructor(props, context) {
    super(props)
    console.log("fields created", props, this.field)
    this.field = context.fields[props.name]
    this.additionalProperties={}
    if (this.field.widget == "SelectField") {
      if(this.field.items){
        this.additionalProperties.items = this.field.items
      }
    }
  }

  render() {
    return (
      <Field name={this.props.name} className={this.props.className} {...this.additionalProperties} field={this.field} component={this.context.uikit.Forms.FieldWidget}/>
    )
  }

}

FieldWrapper.contextTypes = {
  fields: PropTypes.object,
  uikit:  PropTypes.object
};

export { FieldWrapper as Field}

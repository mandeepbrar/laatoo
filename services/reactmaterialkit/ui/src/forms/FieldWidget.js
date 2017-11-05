import React from 'react';
import { Field as RFField } from 'redux-form'
import {
  Checkbox,
  RadioButtonGroup,
  SelectField,
  TextField,
  Toggle,
  DatePicker,
  TimePicker
} from 'redux-form-material-ui'
/*function GetFieldWidget(field) {
  switch(field.widget) {
    case "TextField":
      return renderTextField(field)
  }
  return null
}*/

class FieldWidget extends React.Component {
  constructor(props) {
    super(props)
    let field = props.field
    if (field) {
      switch(field.widget) {
        case "TextField":
          this.renderer = this.renderTextField
          break
        case "NumberField":
          this.renderer = this.renderTextField
          break
        case "Radio":
          this.renderer = this.renderTextField
          break
        case "Checkbox":
          this.renderer = this.renderTextField
          break
        case "Toggle":
          this.renderer = this.renderTextField
          break
        case "ListField":
          this.renderer = this.renderTextField
          break
        case "DatePicker":
          this.renderer = this.renderTextField
          break
        case "TimePicker":
          this.renderer = this.renderTextField
          break
        case "ImagePicker":
          this.renderer = this.renderTextField
          break
      }
    }
    /*if(field && field.widget) {
    }*/
  }
  renderTextField = (fld, props) => {
    let {input, meta} = props
    /*return (
      <TextField name={fld.name}  hintText={fld.label} errorText={meta.touched && meta.error} onChange={input.onChange} onBlur={input.onBlur}
        onFocus={input.onFocus} floatingLabelText={fld.label} className={fld.name + " textfield " + (fld.className?fld.className:"")}/>
    )*/
    return (
      <TextField name={fld.name} name={fld.name} errorText={meta.touched && meta.error} onChange={input.onChange} onBlur={input.onBlur}
        onFocus={input.onFocus} floatingLabelText={fld.label}
         hintText={fld.label} {...props} className={fld.name + " textfield " + (fld.className?fld.className:"")}/>
    )
  }
  render() {
    if(this.renderer) {
      return this.renderer(this.props.field, this.props)
    } else {
      return(
        <div>
        </div>
      )
    }
  }
}

export {
  FieldWidget
}

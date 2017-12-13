import React from 'react';
import { Field as RFField } from 'redux-form'
import {
  Checkbox,
  RadioButtonGroup,
  TextField,
  Toggle,
  DatePicker,
  Switch,
  Select,
  TimePicker
} from 'material-ui';
import { MenuItem } from 'material-ui/Menu';
import { FormControl, FormControlLabel, FormGroup } from 'material-ui/Form';
import { InputLabel } from 'material-ui/Input';
import PropTypes from 'prop-types';

//import {MenuItem, SelectField} from 'material-ui';
//import injectTapEventPlugin from "react-tap-event-plugin";

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
    //injectTapEventPlugin()
    console.log("react material kit fieldwidget", props)
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
          this.renderer = this.renderCheckbox
          break
        case "Switch":
          this.renderer = this.renderSwitch
          break
        case "Select":
          this.renderer = this.renderSelect
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
        default:
          if(field.widgetMod) {

          } else {
            this.renderer = this.renderTextField
          }
          break
      }
    }
    /*if(field && field.widget) {
    }*/
  }

  renderCheckbox = (fld, props) =>  {
    let {input, meta} = props
    return (
      <FormControlLabel className={fld.name + " checkbox "+ (fld.className?fld.className:"")}  label={fld.label}
          control={ <Checkbox name={fld.name} checked={props.value ? true : false} onCheck={props.onChange}
              className={fld.name + " " + (fld.controlClassName?fld.controlClassName:"")}/> }/>
    )
  }

  renderSelect = (fld, props) =>  {
    let {input, meta} = props
    let items=[]
    if(props.items) {
      props.items.forEach(function(item) {
        console.log("menu item", MenuItem, item)
        items.push(
          <MenuItem className={fld.itemClass} value={item.value}>{item.text}</MenuItem>
        )
      })
    }

    return (
      <FormControl className={fld.name + " formcontrol "+ (fld.className?fld.className:"")}>
        <InputLabel htmlFor={fld.name}>{fld.label}</InputLabel>
        <Select name={fld.name} floatingLabelText={fld.label} label={fld.label} errorText={meta.touched && meta.error}
          onChange={(event, index, value) => {input.onChange(event.target.value)}} value={input.value}
          className={fld.name + " select " + (fld.controlClassName?fld.controlClassName:"")}>
        {items}
        </Select>
      </FormControl>
    )
  }

  renderSwitch = (fld, props) =>  {
    let {input, meta} = props
    console.log("switch props", props)
    return (
      <FormControlLabel control={
            <Switch name={fld.name} onChange={input.onChange} checked={input.value}
              className={fld.name + " toggle " + (fld.controlClassName?fld.controlClassName:"")}/>
          } label={fld.label} className={fld.name + " " + (fld.className?fld.className:"")}  />
    )
  }


  renderTextField = (fld, props) => {
    let {input, meta} = props
    return (
      <TextField name={fld.name} errorText={meta.touched && meta.error} onChange={input.onChange} onBlur={input.onBlur}
        onFocus={input.onFocus} floatingLabelText={fld.label} label={fld.label} value={input.value}
        hintText={fld.label} {...props} className={fld.name + " textfield " + (fld.className?fld.className:"")}/>
    )
  }
  render() {
    if(this.renderer) {
      console.log("rendering material ui field", this.props)
      return this.renderer(this.props.field, this.props)
    } else {
      return(
        <div>
        </div>
      )
    }
  }
}

FieldWidget.propTypes = {
  classes: PropTypes.object.isRequired,
};

export {
  FieldWidget
}

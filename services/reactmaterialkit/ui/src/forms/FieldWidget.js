import React from 'react';
import { Field as RFField } from 'redux-form'
import {
  Checkbox,
  RadioButtonGroup,
  TextField,
  Toggle,
  DatePicker,
  Select,
  TimePicker
} from 'material-ui';
import { MenuItem } from 'material-ui/Menu';

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
        case "Toggle":
          this.renderer = this.renderToggle
          break
        case "SelectField":
          this.renderer = this.renderSelectField
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
      <Checkbox name={fld.name} label={fld.label} checked={props.value ? true : false} onCheck={props.onChange}
        className={fld.name + " checkbox " + (fld.className?fld.className:"")}/>
    )
  }

  renderSelectField = (fld, props) =>  {
    let {input, meta} = props
    let items=[]
    if(props.items) {
      props.items.forEach(function(item) {
        console.log("menu item", MenuItem, item)
        items.push(
          <MenuItem className={fld.itemClass} value={item.value} primaryText={item.text} />
        )
      })
    }

    console.log("select field items", items, props)
    return (
      <Select name={fld.name} floatingLabelText={fld.label} label={fld.label} errorText={meta.touched && meta.error}
        onChange={(event, index, value) => input.onChange(value)} value={input.value}
        className={fld.name + " select " + (fld.className?fld.className:"")}>
      {items}
      </Select>
    )
  }

  renderToggle = (fld, props) =>  {
    let {input, meta} = props
    return (
      <Toggle name={fld.name} label={fld.label} errorText={meta.touched && meta.error} {...props} onToggle={input.onChange}
        className={fld.name + " toggle " + (fld.className?fld.className:"")}/>
    )
  }


  renderTextField = (fld, props) => {
    let {input, meta} = props
    return (
      <TextField name={fld.name} errorText={meta.touched && meta.error} onChange={input.onChange} onBlur={input.onBlur}
        onFocus={input.onFocus} floatingLabelText={fld.label}
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

export {
  FieldWidget
}

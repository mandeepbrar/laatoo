import React from 'react';
import { Field as RFField } from 'redux-form'
import {
  Checkbox,
  RadioButtonGroup,
  TextField,
  Toggle,
  DatePicker,
  Switch,
  TimePicker
} from 'material-ui';
import { MenuItem } from 'material-ui/Menu';
import { FormControl, FormControlLabel, FormGroup } from 'material-ui/Form';
import { InputLabel } from 'material-ui/Input';
import {Select} from '../components/Select';
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
    let {input, meta, className, fieldChange} = props
    return (
      <FormControlLabel className={(className?className + " ":"") + fld.name + " checkbox "+ (fld.className?fld.className:"")}  label={fld.label}
          control={ <Checkbox name={fld.name} checked={props.value ? true : false} onCheck={(evt, checked)=>fieldChange(checked, evt.target.name, evt)}
              className={fld.name + " " + (fld.controlClassName?fld.controlClassName:"")}/> }/>
    )
  }

  renderSelect = (fld, props) =>  {
    console.log("render select ", fld, props)
    let {input, meta, className, fieldChange} = props
    let value = props.value? props.value: (input? input.value: null)
    let et = meta? meta.touched && meta.error: null
    className= (className?className + " ":"") + " "+ (fld.className?fld.className:"")
    let isEntity = fld.type=="entity"
    let entity= isEntity ? fld.entity: null
    let items = props.items? props.items: fld.items
    let textField = fld.textField? fld.textField:isEntity? "Name": "text"
    let valueField = fld.valueField? fld.valueField: isEntity? "Id" : "value"
    console.log("field change for select==", fieldChange)

    return <Select items={items} itemClass={fld.itemClass} className={className} onChange={fieldChange} value={value} dataServiceParams={fld.dataServiceParams}
        label={fld.label} name={fld.name} errorText={et} loader={fld.loader} skipDataLoad={fld.skipDataLoad} dataService={fld.dataService} selectItem={fld.selectItem}
        isEntity={isEntity} textField={textField} valueField={valueField} entity={entity} controlClassName={fld.name + " select " + (fld.controlClassName?fld.controlClassName:"")} />
  }

  renderSwitch = (fld, props) =>  {
    let {input, meta, className, fieldChange} = props
    return (
      <FormControlLabel control={
            <Switch name={fld.name} onChange={fieldChange} checked={input.value}
              className={fld.name + " toggle " + (fld.controlClassName?fld.controlClassName:"")}/>
          } label={fld.label} className={(className?className + " ":"") + fld.name + " " + (fld.className?fld.className:"")}  />
    )
  }


  renderTextField = (fld, props) => {
    let {input, meta, className, fieldChange} = props
    return (
      <TextField name={fld.name} errorText={meta.touched && meta.error} onChange={(evt)=>fieldChange(evt.target.value, evt.target.name, evt)} onBlur={input.onBlur}
        onFocus={input.onFocus} floatingLabelText={fld.label} label={fld.label} value={input.value}
        hintText={fld.label} {...props} className={(className?className + " ":"") + fld.name + " textfield " + (fld.className?fld.className:"")}/>
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

FieldWidget.propTypes = {
  classes: PropTypes.object.isRequired,
};

export {
  FieldWidget
}

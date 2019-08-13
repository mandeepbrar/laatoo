import React from 'react';
import { Field as RFField } from 'redux-form'
import {
  Checkbox,
  RadioButtonGroup,
  TextField,
  Toggle,
  DatePicker,
  MenuItem,
  FormControl, FormControlLabel, FormGroup,
  InputLabel, 
  Switch,
  TimePicker
} from '@material-ui/core';
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
    this.renderer = this.renderTextField
    if (field && field.widget) {
      switch(field.widget.name) {
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
      }
    }
    this.state = {value: props.value}
    this.widgetProps = field.widget && field.widget.props? field.widget.props:{}
    this.className = (props.className?props.className :"") + " " + field.name +" " + (this.widgetProps.className?this.widgetProps.className:" ") 
    this.controlClassName= field.name + " " + (this.widgetProps.controlClassName?this.widgetProps.controlClassName:"")
    console.log("constructing material kit field", field, props)
    /*if(field && field.widget) {
    }*/
  }

  componentWillReceiveProps(nextProps, nextState) {
    this.setState(Object.assign({}, this.state, {value: nextProps.value}))
  }    

  renderCheckbox = (fld, props) =>  {
    return (
      <FormControlLabel className={this.className + " checkbox "}  label={fld.label}
          control={ <Checkbox name={fld.name} checked={props.value ? true : false} onCheck={(evt, checked)=>props.onChange(checked, evt.target.name, evt)}
              className={this.controlClassName}/> }/>
    )
  }

  renderSelect = (fld, props) =>  {
    console.log("render select ", fld, props)
    let isEntity = fld.type=="entity"
    let entity= isEntity ? fld.entity: null
    let items = props.items? props.items: this.widgetProps.items
    let textField = this.widgetProps.textField? this.widgetProps.textField:isEntity? "Name": "text"
    let valueField = this.widgetProps.valueField? this.widgetProps.valueField: isEntity? "Id" : "value"
    let selectItem = props.selectItem? props.selectItem: this.widgetProps.selectItem

    return <Select items={items} itemClass={this.widgetProps.itemClass} className={this.className + "select"} onChange={props.onChange} value={this.state.value} dataServiceParams={this.widgetProps.dataServiceParams}
        errorText={props.errorText} loader={this.widgetProps.loader} loadData={this.widgetProps.loadData} dataService={this.widgetProps.dataService} selectItem={selectItem}
        label={fld.label} name={fld.name} isEntity={isEntity} textField={textField} valueField={valueField} entity={entity} controlClassName={this.controlClassName + " select "} />
  }

  renderSwitch = (fld, props) =>  {
    return (
      <FormControlLabel control={
            <Switch name={fld.name} onChange={props.onChange} checked={this.state.value} className={this.controlClassName + " toggle "}/>
          } label={fld.label} className={this.className + " switch"}  />
    )
  }

  textChange = (evt) => {
    if(this.props.onChange) {
      console.log("text change ", evt.target.value)
      this.props.onChange(evt.target.value, evt.target.name, evt)
    }
  }

  renderTextField = (fld, props) => {
    console.log("rendertext field", props, fld)
    return (
      <TextField name={fld.name} errorText={props.errorText} onChange={this.textChange} onBlur={props.onBlur} onFocus={props.onFocus} 
        floatingLabelText={fld.label} label={fld.label} value={this.state.value} hintText={fld.label} className={this.className + " textfield " }/>
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

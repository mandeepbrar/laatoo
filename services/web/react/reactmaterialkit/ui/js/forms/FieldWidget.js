import React from 'react';
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
import Slider from '@material-ui/lab/Slider';
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
    let widgetName = _tn(props.widgetName, "TextField")
    this.controlProps = {}
    switch(widgetName) {
      case "TextField":
        this.renderer = this.renderTextField
        break
      case "NumberField":
        this.controlProps = {type: "number"}
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
      case "Slider":
        this.renderer = this.renderSlider
        break
      default:
        this.renderer = this.renderTextField
    }
    this.state = {value: props.value}
    this.label = _tn(props.label, props.name)
    this.className = _tn(props.className, "") + widgetName
    this.controlClassName= _tn(props.controlClassName, "")
    console.log("material kit field widget", props)
    /*if(field && field.widget) {
    }*/
  }

  componentWillReceiveProps(nextProps, nextState) {
    this.setState(Object.assign({}, this.state, {value: nextProps.value}))
  }    

  renderCheckbox = (props) =>  {
    return (
      <FormControlLabel className={this.className + " checkbox "}  label={this.label}
          control={ <Checkbox name={this.name} checked={this.state.value ? true : false} onCheck={this.change}
              className={this.controlClassName}/> }/>
    )
  }

  renderSelect = (props) =>  {
    console.log("render select ", props)
    let textField = _tn(props.textField, "text")
    let valueField = _tn(props.valueField, "value") 

    return <Select items={props.items} itemClass={props.itemClass} className={this.className + " select "} onChange={this.change} value={this.state.value} dataServiceParams={props.dataServiceParams}
        errorText={props.errorText} loader={props.loader} loadData={props.loadData} dataService={props.dataService} selectItem={props.selectItem}
        label={this.label} name={props.name} textField={textField} valueField={valueField} controlClassName={this.controlClassName + " select "} />
  }

  renderSwitch = (props) =>  {
    return (
      <FormControlLabel control={
            <Switch name={props.name} onChange={this.change} checked={this.state.value} className={props.controlClassName + " toggle "}/>
          } label={this.label} className={this.className + " switch "}  />
    )
  }

  renderSlider = (props) =>  {
    console.log("rendering slider", Slider, this.state.value)
    return (
      <FormControlLabel control={
        <Slider name={props.name} onChange={(evt, val)=> {console.log(val);this.props.onChange(val, this.props.name, evt)}} value={this.state.value} 
          {...props.widgetProps} className={props.controlClassName + " slider "}/>
      } label={this.label} className={this.className + " slider "}  />
    )
  }

  change = (evt) => {
    if(this.props.onChange) {
      console.log("change ", evt, evt.target.value)
      this.props.onChange(evt.target.value, this.props.name, evt)
    }
  }

  renderTextField = (props) => {
    console.log("rendertext field", props)
    return (
      <TextField name={this.name} errorText={props.errorText} onChange={this.change} onBlur={props.onBlur} onFocus={props.onFocus} {...this.controlProps}
        floatingLabelText={this.label} label={this.label} value={this.state.value} hintText={this.label} className={this.className + " textfield " }/>
    )
  }
  render() {
    if(this.renderer) {
      return this.renderer(this.props)
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

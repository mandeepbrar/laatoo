import React from 'react';
import {Select} from '../components/Select';
import {TextField} from '../components/TextField';
import PropTypes from 'prop-types';
import Form from 'react-bootstrap/Form'


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
    switch(widgetName) {
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
        this.renderer = this.renderTextField
    }
    this.state = {value: props.value}
    this.className = _tn(props.className, "") + widgetName
    this.controlClassName= _tn(props.controlClassName, "")
    /*if(field && field.widget) {
    }*/
  }

  componentWillReceiveProps(nextProps, nextState) {
    this.setState(Object.assign({}, this.state, {value: nextProps.value}))
  }    

  renderCheckbox = (props) =>  {
    return (

      <Form.Group controlId={props.name} className={this.className + " checkbox "} >
        <Form.Check type="checkbox" label={props.label}  value={this.state.value} onCheck={this.change} className={this.controlClassName} />
      </Form.Group>
    )
  }

  renderSelect = (props) =>  {
    console.log("render select ", props)
    let textField = _tn(props.textField, "text")
    let valueField = _tn(props.valueField, "value") 

    return <Select items={props.items} itemClass={props.itemClass} className={this.className} onChange={this.change} value={this.state.value} dataServiceParams={props.dataServiceParams}
        errorText={props.errorText} loader={props.loader} loadData={props.loadData} dataService={props.dataService} selectItem={props.selectItem}
        label={props.label} name={props.name} textField={textField} valueField={valueField} controlClassName={this.controlClassName + " select "} />
  }

  renderSwitch = (props) =>  {
    return this.renderCheckbox(props)
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
      <TextField label={props.label} name={props.name} placeholder={props.placeholder} onChange={this.change} className={this.className}
        onBlur={props.onBlur} onFocus={props.onFocus} value={this.state.value} errorText={props.errorText}/>
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

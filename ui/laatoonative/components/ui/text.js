import React from 'react';
import {Input} from 'native-base';
import {Window} from 'laatoocommon';

class Text extends React.Component {
  constructor(props) {
    super(props)

    this.value = props.value? props.value: "";
    this.state = {value: this.value};
    this.setValue = this.setValue.bind(this)
    this.keyPress = this.keyPress.bind(this)
    this.change = this.change.bind(this);
  }
  componentWillReceiveProps(nextprops) {
    if(nextprops.value) {
      this.setValue(nextprops.value)
    }
  }
  setLanguage(lang) {
    this.language = lang
  }
  setValue(value, successFunc) {
    this.setState({value: value}, successFunc)
    this.value = value
    if(this.props.onChange) {
      let evt={target:{name: this.props.name, value: value}}
      this.props.onChange(evt)
    }
  }
  keyPress(event) {
    console.log("key press", event)
/*    console.log("key press event ", event, this.props)
    let target = event.target
    if(this.props.onEnterKey && event.which == 13) {
      this.props.onEnterKey()
      return
    }*/
  }
  change(val) {
    this.setValue(val)
  }
  render() {
    let config={}
    if(this.props.config) {
      config = this.props.config
    }
    console.log("this is rendering my textfield")
    return (
      <Input name={this.props.name} className={this.props.className}  onKeyPress={this.keyPress} value={this.state.value}
        defaultValue={this.props.defaultValue} rows={this.props.rows}  rowsMax={this.props.rows} multiLine={this.props.multiline}
        onKeyDown={this.props.onKeyDown} placeholder={this.props.placeholder} type={this.props.type}
        style={config.style} textareaStyle={{height: 'initial'}} inputStyle={config.inputStyle} onChangeText={this.change}/>
    )
  }
}


export {Text as TextField};

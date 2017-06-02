import React from 'react';
import {Item, Input, Label} from 'native-base';
import {Window} from 'laatoocommon';

class Text extends React.Component {
  constructor(props) {
    super(props)

    this.value = props.value? props.value: "";
    this.state = {value: this.value};
    this.setValue = this.setValue.bind(this)
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
/*    console.log("key press event ", event, this.props)
    let target = event.target
    if(this.props.onEnterKey && event.which == 13) {
      this.props.onEnterKey()
      return
    }
  }*/
  change(val) {
    this.setValue(val)
  }
  render() {
    let config={}
    if(this.props.config) {
      config = this.props.config
    }
    return (
      <Item fixedLabel={this.props.fixedLabel} stackedLabel={this.props.stackedLabel} floatingLabel={this.props.floatingLabel} inlineLabel={this.props.inlineLabel}>
        <Label>{this.props.label}</Label>
        {this.props.lefticon}
        <Input name={this.props.name} className={this.props.className}  onKeyPress={this.keyPress} value={this.state.value}
          defaultValue={this.props.defaultValue} rows={this.props.rows}  rowsMax={this.props.rows} multiLine={this.props.multiline}
          placeholder={this.props.placeholder} secureTextEntry={this.props.secureTextEntry} rounded={this.props.rouded}
          textColor={this.props.textColor} inputColorPlaceholder={this.props.inputColorPlaceholder} inputHeightBase={this.props.inputHeightBase}
          style={config.style} inputStyle={config.inputStyle} onChangeText={this.change}/>
        {this.props.icon}
      </Item>
    )
  }
}


export {Text as TextField};

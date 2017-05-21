import React from 'react';
import TextField from 'material-ui/TextField';
import {Window} from 'laatoo';
import t from 'tcomb-form';

class Text extends React.Component {
  constructor(props) {
    super(props)

    this.value = props.value? props.value: "";
    this.state = {value: this.value};
    this.setValue = this.setValue.bind(this)
    this.keyPress = this.keyPress.bind(this)
    this.keyMap = props.keyMap
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
    let target = event.target
    if(this.props.onEnterKey && event.which == 13) {
      this.props.onEnterKey()
      return
    }
    if(Window.language!="d") {
      if(event.which>32 && event.which<127) {
        let ch = event.which;
        event.preventDefault()
        let pos = target.selectionStart
        let v = this.state.value
        let newval = [v.slice(0, pos), this.keyMap[ch-33], v.slice(pos)].join('');
        this.setValue(newval, () => {target.selectionStart = target.selectionEnd = pos+1})
        return
      }
    }
  }
  change(event) {
    this.setValue(event.target.value)
  }
  render() {
    let config={}
    if(this.props.config) {
      config = this.props.config
    }
    return (
      <TextField name={this.props.name} className={this.props.className} onKeyPress={this.keyPress} value={this.state.value}
        defaultValue={this.props.defaultValue} rows={this.props.rows}  rowsMax={this.props.rows} multiLine={this.props.multiline}
        onKeyDown={this.props.onKeyDown} placeholder={this.props.placeholder} type={this.props.type}
        style={config.style} textareaStyle={{height: 'initial'}} inputStyle={config.inputStyle} onChange={this.change}/>
    )
  }
}


class FormTextField extends t.form.Component { // extend the base class
  getTemplate() {
    return (locals) => {
      let type="text"
      if(locals.config.type) {
        type = locals.config.type
      }
      let onChange= null;
      if(locals.onChange) {
        onChange= function(evt) {
          locals.onChange(evt.target.value);
        }
      }
      let val = locals.value
      if(!val) {
        val = ""
      }
      return (
        <TextField value={val} multiline={locals.config.multiline} className={locals.config.className} rows={locals.config.rows}
          onChange={onChange} type={type} config={locals.config} name={locals.config.name}>
        </TextField>
      );
    };
  }
}



export {Text as TextField, FormTextField as FormTextField};

import t from 'tcomb-form';
import React from 'react';

var configs = {
  theme: 'snow'
};
var toolbarOptions = [
  [{ size: ['small', false, 'large', 'huge'] }],
  ['bold', 'italic', 'underline'],
  [{ color: [] }, { background: [] }],    // Snow theme fills in values
  [{list: 'ordered'}, {list:'bullet'}, {align:[]}],
  [{ script: 'sub' }, { script: 'super' }],
  ['link','image']
];

class RichEdit extends React.Component {
  constructor(props) {
    super(props)
    if(props.value) {
      this.state = {value: props.value};
      this.value = props.value;
    }
    this.setValue = this.setValue.bind(this)
    this.keyPress = this.keyPress.bind(this)
    this.change = this.change.bind(this);
    this.keymaps = props.keymaps
  }
  componentDidMount() {
    this.editor = new Quill(this.refs.editor,
      {
      modules: {
        toolbar: toolbarOptions
      },
      theme: 'snow'
      });
    //quill.addModule('toolbar', this.refs.toolbar);
    this.editor.on('text-change', this.change);
  }

  componentWillReceiveProps(nextprops) {
    if(!this.value && nextprops.value ) {
      this.editor.pasteHTML(nextprops.value)
      //this.setValue(nextprops.value)
    }
  }
  setValue(value, successFunc) {
    this.setState({value}, successFunc)
    this.value = value
    if(this.props.onChange) {
      this.props.onChange(value)
    }
  }
  keyPress(event, some) {
    if(window.language && window.language!="d") {
      let target = event.target
      let keymap = this.keymaps[window.language]
      if(keymap && event.which>32 && event.which<127) {
        event.preventDefault()
        let ch = event.which;
        let selection = this.editor.getSelection();
        let pos = selection.index
        this.editor.insertText(selection.index, keymap[ch-33])
        this.editor.setSelection(pos+1, 0)
        return
      }
    }
  }
  change(value, target, prop1, prop3) {
    this.setValue(this.refs.editor.firstChild.innerHTML)
  }
  render() {
    let props = this.props
    return (
      <div style={this.props.style} className = {this.props.className}>
        <div id="toolbar" ref="toolbar">
        </div>
        <div ref="editor" onKeyPress={this.keyPress} style={this.props.textAreaStyle} className={this.props.textAreaClassName}>
        </div>
      </div>
    )
  }
}

class TextEdit extends t.form.Component { // extend the base class
  getTemplate() {
    return (locals) => {
      let config = locals.config
      return (
        <div>
          {(config && config.hideLabel)? null :<label>{locals.label}</label>}
          <RichEdit  style={config.style} className={config.className} keymaps={config.keymaps} value={locals.value}
            textAreaStyle={config.textAreaStyle}  textAreaClassName={config.textAreaStyle} onChange={locals.onChange} />
        </div>
      );
    };
  }
}


export {
  TextEdit as TextEdit,
  RichEdit as RichEdit
};

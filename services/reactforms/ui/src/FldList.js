import React from 'react';
import {Action} from 'reactwebcommon';
const PropTypes = require('prop-types');

class ListEditor extends React.Component {
  constructor(props) {
    super(props)
    this.state = {values: props.values}
  }
  componentWillReceiveProps(nextProps) {
    let values = nextProps.values
    this.setState(Object.assign({}, this.state, {values}))
  }
  addItem=()=> {
    let items = this.state.values.slice();
    items.push("")
    this.onChange(items)
  }
  getValues=()=> {
    return this.state.values
  }
  changeValue = (value, index) => {
    if(this.state.values.length > index) {
      let items = this.state.values.slice();
      items[index] = value
      this.onChange(items)
    }
  }
  onChange = (values) => {
    if(this.props.onChange) {
      this.props.onChange(values)
    } else {
      this.setState({values: values})
    }
  }
  render() {
    let editor = this
    let props = this.props
    let comps = []
    this.state.values.forEach(function(str, index) {
      let editText = (evt)=> {console.log("edit text", evt); editor.changeValue(evt.target.value, index)}
      let newProps = Object.assign({}, props.baseProps, {input:{value: str}})
      if(props.field.widget == "Select") {
        comps.push(<props.uikit.Forms.FieldWidget className={props.className + " w100"} value={str} onChange={editText} field={props.field} {...props.baseProps} {...props.ap} time={editor.state.time}/>)
      }else {
        comps.push(<props.uikit.TextField className={props.className + " w100"} value={str} onChange={editText} time={editor.state.time} />)
      }
    })
    return <props.uikit.Block className="w100" titleBarActions={[<Action action={{actiontype:"method"}} className="left" method={this.addItem}><props.uikit.Icons.NewIcon/></Action>]}>{comps}</props.uikit.Block>
  }
}

class FldList extends React.Component {
  constructor(props) {
    super(props)
    let vals = props.baseProps.input.value? props.baseProps.input.value: []
    this.state = {values: vals, time: props.time}
    this.field = props.field
  }
  componentWillReceiveProps(nextProps) {
    let values = nextProps.baseProps.input.value? nextProps.baseProps.input.value: []
    this.setState(Object.assign({}, this.state, {values}))
  }
  editingComplete = () => {
    //this.setState(Object.assign({}, this.state, {values}))
    Window.closeDialog()
    let values = this.editor.getValues();
    this.onChange(values)
  }
  onChange = (values) => {
    this.props.baseProps.input.onChange(values)
  }
  editList = () => {
    console.log("edit list")
    let actions = [<Action action={{actiontype:"method"}} widget="button" className="right" method={this.editingComplete}>Save</Action>]
    Window.showDialog("Items", <ListEditor ref={(editor) => {this.editor = editor;}} uikit={this.context.uikit} values={this.state.values} {...this.props}/>, null, actions )
  }
  render() {
    let uikit = this.context.uikit
    console.log("reder stringlist", this.state, this.props, this.context, this.field)
    let cl = this.props.className? this.props.className:""
    if(this.field.inplace) {
      return (
        <uikit.Block time={this.state.time} className={cl + " w100 row stringlist_inplace"}>
          <ListEditor ref={(editor) => {this.editor = editor;}} uikit={this.context.uikit} onChange={this.onChange} values={this.state.values} {...this.props}/>
        </uikit.Block>
      )
    } else {
      let val = this.state.values.join()
      return (
        <uikit.Block time={this.state.time} className={cl + " row stringlist"}>
          <uikit.Block className="col-xs-12 label">{this.props.name}</uikit.Block>
          <uikit.Block className="value col-xs-10">
          {val? val: "<No Data>"}
          </uikit.Block>
          <Action action={{actiontype:"method"}} className=" col-xs-2" method={this.editList}>
            <uikit.Icons.EditIcon/>
          </Action>
        </uikit.Block>
      )
    }
  }
}

FldList.contextTypes = {
  uikit: PropTypes.object
};

export {FldList}

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
    console.log("componentWillReceiveProps list editor fldlist", values)
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
    console.log("value change triggered ", value, index)
    if(this.state.values.length > index) {
      let items = this.state.values.slice();
      items[index] = value
      this.onChange(items)
    }
  }
  onChange = (values) => {
    console.log("value onchange triggered ", values, this.props)
    if(this.props.fieldChange) {
      this.props.fieldChange(values, this.props.name, null)
    }
    this.setState({values: values})
  }
  render() {
    let editor = this
    let props = this.props
    let comps = []
    console.log("list editor ", props, this.state)
    this.state.values.forEach(function(str, index) {
      let editText = (evt, val)=> {console.log("edit text", evt, val); editor.changeValue(event.target.value, index)}
      let newProps = Object.assign({}, props, {input:{value: str}})
      if(props.field.widget == "Select") {
        console.log("creating select val", str)
        comps.push(<_uikit.Field className={props.className + " w100"} value={str} onChange={editText} field={props.field} {...props}/>)
      }else {
        comps.push(<_uikit.TextField className={props.className + " w100"} value={str} onChange={editText} />)
      }
    })
    return <_uikit.Block className="w100" titleBarActions={[<Action action={{actiontype:"method"}} className="left" method={this.addItem}><_uikit.Icons.NewIcon/></Action>]}>{comps}</_uikit.Block>
  }
}

class FldList extends React.Component {
  constructor(props) {
    super(props)
    console.log("fld list", props)
    this.state = {values: this.createState(props.value)}
    this.field = props.field
  }
  componentWillReceiveProps(nextProps) {
    console.log("componentWillReceiveProps fldlist", nextProps)
    this.setState(Object.assign({}, this.state, {values: this.createState(nextProps.value)}))
  }
  createState(val) {
    if(val) {
      if(Array.isArray(val)) {
        return val
      }
    }
    return []
  }
  editingComplete = () => {
    //this.setState(Object.assign({}, this.state, {values}))
    Window.closeDialog()
    let values = this.editor.getValues();
    this.onChange(values)
  }
  onChange = (values) => {
    console.log("onchange fldlist", values)
    this.props.onChange(values)
  }
  editList = () => {
    console.log("edit list")
    let actions = [<Action action={{actiontype:"method"}} widget="button" className="right" method={this.editingComplete}>Save</Action>]
    Window.showDialog("Items", <ListEditor ref={(editor) => {this.editor = editor;}} values={this.state.values} {...this.props}/>, null, actions )
  }
  render() {
    console.log("reder stringlist", this.state, this.props, this.context, this.field)
    let cl = this.props.className? this.props.className:""
    if(this.field.inplace) {
      return (
        <_uikit.Block className={cl + " w100 row stringlist_inplace"}>
          <ListEditor ref={(editor) => {this.editor = editor;}} onChange={this.onChange} values={this.state.values} {...this.props}/>
        </_uikit.Block>
      )
    } else {
      console.log("non list editor", this.state, "props", this.props)
      let val = this.state.values.join()
      console.log("non list editor1", val, _uikit)
      return (
        <_uikit.Block className={cl + " row stringlist"}>
          <_uikit.Block className="col-xs-12 label">{this.props.name}</_uikit.Block>
          <_uikit.Block className="value col-xs-10">
          {val? val: "<No Data>"}
          </_uikit.Block>
          <Action action={{actiontype:"method"}} className=" col-xs-2" method={this.editList}>
            <_uikit.Icons.EditIcon/>
          </Action>
        </_uikit.Block>
      )
    }
  }
}

export {FldList}

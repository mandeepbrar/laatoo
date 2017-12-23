import React from 'react';
import {Action} from 'reactwebcommon';
const PropTypes = require('prop-types');

class ListEditor extends React.Component {
  constructor(props) {
    super(props)
    this.state = {values: props.values}
  }
  addItem=()=> {
    let items = this.state.values.slice();
    items.push("")
    this.setState({values: items})
  }
  getValues=()=> {
    return this.state.values
  }
  changeValue = (value, index) => {
    if(this.state.values.length > index) {
      let items = this.state.values.slice();
      items[index] = value
      this.setState({values: items})
    }
  }
  render() {
    let editor = this
    let props = this.props
    let comps = []
    this.state.values.forEach(function(str, index) {
      let editText = (evt)=> {editor.changeValue(evt.target.value, index)}
      let newProps = Object.assign({}, props.baseProps, {input:{value: str}})
      if(props.field.widget == "Select") {

      }else {
        comps.push(<props.uikit.TextField className={props.className + " w100"} value={str} onChange={editText} time={editor.state.time} />)
      }
    })
    return <props.uikit.Block titleBarActions={[<Action action={{actiontype:"method"}} className="right" method={this.addItem}><props.uikit.Icons.NewIcon/></Action>]}>{comps}</props.uikit.Block>
  }
}

class Stringlist extends React.Component {
  constructor(props) {
    super(props)
    let vals = props.baseProps.input.value? props.baseProps.input.value: []
    this.state = {values: vals, time: props.time}
  }
  componentWillReceiveProps(nextProps) {
    let values = nextProps.baseProps.input.value? nextProps.baseProps.input.value: []
    this.setState(Object.assign({}, this.state, {values}))
  }
  editingComplete = () => {
    let values =this.editor.getValues();
    //this.setState(Object.assign({}, this.state, {values}))
    Window.closeDialog()
    console.log("editing values", values)
    this.props.baseProps.input.onChange(values)
  }
  editList = () => {
    console.log("edit list")
    let actions = [<Action action={{actiontype:"method"}} widget="button" className="right" method={this.editingComplete}>Save</Action>]

    Window.showDialog("Items", <ListEditor ref={(editor) => {this.editor = editor;}} uikit={this.context.uikit} values={this.state.values} {...this.props}/>, null, actions )
  }
  render() {
    let uikit = this.context.uikit
    console.log("reder stringlist", this.state, this.props, this.context, Action)
    let val = this.state.values.join()
    let cl = this.props.className? this.props.className:""
    return (<uikit.Block time={this.state.time} className={cl + " row stringlist"}>
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

Stringlist.contextTypes = {
  uikit: PropTypes.object
};

export {Stringlist}

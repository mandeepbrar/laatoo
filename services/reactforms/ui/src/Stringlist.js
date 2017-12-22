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
  render() {
    let state = this.state
    let props = this.props
    let comps = []
    console.log("list eidot", this.props)
    state.values.forEach(function(str) {
      let newProps = Object.assign({}, props.baseProps, {input:{value: str}})
      comps.push(<props.baseComponent className={props.className} {...props.ap} time={state.time} field={props.field} {...newProps}/>)
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
  editList = () => {
    console.log("edit list")
    Window.showDialog("Items", <ListEditor uikit={this.context.uikit} values={this.state.values} {...this.props}/>)
  }
  render() {
    console.log("reder stringlist", this.state, this.props, this.context, Action)
    let val = this.state.values.join()
    console.log('val', val);
    return <this.context.uikit.Block time={this.state.time} className="w100">
      <b>{this.props.name}</b>{val? val: "<No Data>"}
      <Action action={{actiontype:"method"}} className="right" method={this.editList}><this.context.uikit.Icons.EditIcon/></Action>
      </this.context.uikit.Block>
  }
}

Stringlist.contextTypes = {
  uikit: PropTypes.object
};

export {Stringlist}

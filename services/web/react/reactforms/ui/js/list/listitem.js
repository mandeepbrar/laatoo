import React from 'react'
import {Action} from 'reactwebcommon'
const PropTypes = require('prop-types');

class Item extends React.Component {
  constructor(props, ctx) {
    super(props)
    console.log(" list item ", props)
    let title = this.getTitle(props)
    this.state = {value: props.value, index: props.index, title: title}
    this.titleClick = (props.mode == "panes")? this.edit: null;
  }

  getTitle = (props) => {
    let val = props.value? props.value: {}
    let title = val[props.titleField];
    title = title? title: val
    console.log("edit item title", title, props.titleField, val)
    return title.toString()
  }

  componentWillReceiveProps(nextProps) {
    console.log("entity list : componentWillReceiveProps", nextProps)
    let title = this.getTitle(nextProps)
    this.setState(Object.assign({}, this.state, {value: nextProps.value, index: nextProps.index, title: title}))
  }

  /*submitEditData = (data, success, failure) => {
    this.props.edit(data, this.state.index)
  }*/

  edit = () => {
    let title = "Edit "+this.state.title
    this.props.openForm(this.state.value, this.submit, title, this.state.index)
  }
  submit=(data)=> {
    console.log("submitting data for item", data, this.state.index)
    this.props.submitEditData(data, this.state.index)
  }
  removeItem = () => {
    this.props.removeItem(this.state.index)
  }

  actions = () => {
    return (
      <_uikit.Block className="right">
        <Action action={{actiontype:"method"}} className="edit p10" method={this.edit}>
            <_uikit.Icons.EditIcon />
        </Action>
        <Action action={{actiontype:"method"}} className="remove p10" method={this.removeItem}>
            <_uikit.Icons.DeleteIcon />
        </Action>
      </_uikit.Block>
    )
  }

  title = () => {
    return (
      <_uikit.Block className="left" onClick={this.titleClick}>
      {this.state.title}
      </_uikit.Block>  
    )
  }

  render() {
    console.log("rendering items in entity list", this.props, this.state)
    return (
        <_uikit.Block  className="row between-xs">
          {this.title()}
          {
            (this.props.mode != "panes")?
            this.actions():
            null
          }
        </_uikit.Block>
    )
  }
}

export default Item
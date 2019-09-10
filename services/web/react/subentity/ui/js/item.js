import React from 'react'
import {Action, LoadableComponent} from 'reactwebcommon'
import {Panel} from 'reactpages'
import SelectEntity from './selectentity'
const PropTypes = require('prop-types');

class Item extends React.Component {
  constructor(props, ctx) {
    super(props)
    console.log("entity list field ", props)
    let title = this.getTitle(props)
    this.state = {value: props.value, index: props.index, title: title}
  }

  getTitle = (props) => {
    let fld = props.field
    let val = props.value? props.value: {}
    let titleField = fld.textField? fld.textField: "Name" //textfield? or titleField
    let title = val[titleField];
    title = title? title: val["Title"]
    return title
  }

  componentWillReceiveProps(nextProps) {
    console.log("entity list field : componentWillReceiveProps", nextProps)
    let title = this.getTitle(nextProps)
    this.setState(Object.assign({}, this.state, {value: nextProps.value, index: nextProps.index, title: title}))
  }

  submitEditData = (data, success, failure) => {
    this.props.edit(data, this.state.index)
  }

  edit = () => {
    let fld = this.props.field
    let title = "Edit "+this.state.title
    this.props.openForm(this.state.value, this.submitEditData, title)
  }

  removeItem = () => {
    this.props.removeItem(this.state.index)
  }


  render() {
    console.log("rendering items in entity list", this.props, this.state)
    return (
        <_uikit.Block  className="row between-xs">
            <_uikit.Block className="left" >
            {this.state.title}
            </_uikit.Block>
            <_uikit.Block className="right">
                <Action action={{actiontype:"method"}} className="edit p10" method={this.edit}>
                    <_uikit.Icons.EditIcon />
                </Action>
                <Action action={{actiontype:"method"}} className="remove p10" method={this.removeItem}>
                    <_uikit.Icons.DeleteIcon />
                </Action>
            </_uikit.Block>
        </_uikit.Block>
    )
  }
}

export default Item
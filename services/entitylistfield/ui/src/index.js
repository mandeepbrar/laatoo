import React from 'react'
import {Action} from 'reactwebcommon'
import {Panel} from 'reactpages'
const PropTypes = require('prop-types');

class EntityListField extends React.Component {
  constructor(props, ctx) {
    super(props)
    let items = props.input.value? props.input.value: []
    console.log("props...i entity list", props)
    let formName = props.field.form? props.field.form : "new_form_"+props.field.entity.toLowerCase()
    this.formDesc = {type: "form", id: formName}
    this.uikit = ctx.uikit;
    this.state = {items}
  }

  componentWillReceiveProps(nextProps) {
    console.log("componentWillReceiveProps(nextProps).", nextProps)
    let items = nextProps.input.value? nextProps.input.value: []
    this.setState({items})
  }

  actions = (f, submit, reset)=> {
    console.log("actios returned", f, submit, reset)
    return (
      <this.uikit.Block className="right p20">
        <this.uikit.ActionButton onClick={reset}>
        Reset
        </this.uikit.ActionButton>
        <this.uikit.ActionButton onClick={submit}>
        Add
        </this.uikit.ActionButton>
      </this.uikit.Block>
    )
  }

  submit = (data, success, failure) => {
      let items = this.addItem(data.data)
      console.log(" items", items)
      this.props.input.onChange(items)
  }

  openForm = () => {
    console.log("opened form", this.props)
    Window.showDialog(<h1>Add entity</h1>, <Panel actions={this.actions} onSubmit={this.submit} description={this.formDesc} />) //, actions, contentStyle)
  }

  addItem = (item) => {
    let items = this.state.items.slice();
    items.push(item)
    return items
    //console.log("items set to ", items, item)
    //this.setState(Object.assign({}, {items: items}))
  }

  removeItem = (item, index) => {
    let items = this.state.items.slice();
    if (index > -1) {
      items.splice(index, 1);
    }
    this.props.input.onChange(items)
  }

  render() {
    let items = []
    console.log("rendering items in entity list", this.props, this.state)
    let comp = this
    this.state.items.forEach(function(k, index) {
      var removeItem = () => {
        comp.removeItem(k, index)
      }
      let textField = comp.props.entityText? comp.props.entityText: "Name"
      let text = k[textField];
      text = text? text: k["Title"]
      items.push(
        <comp.uikit.Block>
          <comp.uikit.Block>
          {text}
          </comp.uikit.Block>
          <comp.uikit.ActionButton className="removeButton" onClick={removeItem}>
          X
          </comp.uikit.ActionButton>
        </comp.uikit.Block>
      )
    })
    return (
      <this.uikit.Block>
        <this.uikit.Block className="right">
          <Action name="listfield_new_entity" method={this.openForm}>+</Action>
        </this.uikit.Block>
        {items}
      </this.uikit.Block>
    )
  }
}

EntityListField.contextTypes = {
  uikit: PropTypes.object
};

export {
  EntityListField as EntityList
}

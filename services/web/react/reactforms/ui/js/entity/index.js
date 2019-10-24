import React from 'react'
import {Action, LoadableComponent} from 'reactwebcommon'
import {Panel} from 'reactpages'
const PropTypes = require('prop-types');

class Entity extends LoadableComponent {
  constructor(props, ctx) {
    super(props)
    let formName = _tn(props.form, "new_form_" + props.entity.toLowerCase())
    this.formDesc = {type: "form", id: formName}
    this.state = {value: props.value}
  }

  componentWillReceiveProps(nextProps) {
    this.setState({value: nextProps.value})
  }

  change = (value, name, evt) => {
    console.log("changing entity", value, this.props, name, evt)
    this.props.onChange(value, this.props.name, evt)
    this.setState(Object.assign({}, this.state, {value}))
  }

  render() {
    let props = this.props
    console.log("subentity ", this.state, props)
    let value=_tn(this.state.value, {})
    let title = props.skipLabel? null: props.label
    return (
      <Panel actions={props.actions} formData={value} subform={props.subform} closePanel={props.closePanel} title={title} onChange={this.change} onFormSubmit={props.onFormSubmit} description={this.formDesc} />
    )
  }
}

Entity.contextTypes = {
  getFormValue: PropTypes.func,
  overlayComponent: PropTypes.func
};

export {
  Entity as Entity
}

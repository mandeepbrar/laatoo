'use strict';

import React from 'react';
import { connect } from 'react-redux';
import {ActionNames} from './Actions';
import {createAction} from 'uicommon';
import { Field, reduxForm } from 'redux-form'
const PropTypes = require('prop-types');

class WebFormUI extends React.Component {
  constructor(props, context) {
    super(props);
    this.uikit = context.uikit
    console.log("web form ui constructor", this.props, context)
    let desc = this.props.description
    ///const { handleSubmit } = props
  /*  this.submitForm = this.submitForm.bind(this);
    this.setValue = this.setValue.bind(this);
    this.submit = this.submit.bind(this);
    this.getValue = this.getValue.bind(this);
    this.onChange = this.onChange.bind(this);
    if(this.props.actionButtons) {
      this.actionButtons = this.props.actionButtons
    } else {

    }
    this.lookupSchemaOptions = desc.lookupSchemaOptions
    let so = desc.lookupSchemaOptions? desc.lookupSchemaOptions(this) : desc.schemaOptions

    this.state = {formValue: props.entityData, so : so, key: "entityform" + (new Date())}
    if(props.refCallback) {
      props.refCallback(this)
    }*/
  }

  setValue(val) {
    this.setState(Object.assign(this.state,{formValue: val}))
  }

  getValue() {
    return this.refs.form.getValue()
  }

  componentWillReceiveProps(nextprops) {
    let ed = nextprops.data? nextprops.data: this.state.formValue
    let so = this.lookupSchemaOptions? this.lookupSchemaOptions(this, ed, {}, "", this.state.so): this.state.so
    if(so) {
      this.setState( {formValue: nextprops.data, so: so, key: "entityform" + (new Date())})
    }
    if(this.props.refCallback) {
      this.props.refCallback(this)
    }
  }

  submitForm(evt) {
    evt.preventDefault();
    let validationRes = this.refs.form.validate()
    let data = this.refs.form.getValue()
    if (!data) {
      if(this.props.failureCallback) {
        this.props.failureCallback(validationRes)
      }
      console.log(validationRes);
      return;
    }
    data = Object.assign({}, data);
    if(this.props.preSave) {
      data = this.props.preSave(data);
    }
    console.log("data to submit", data)
    if (!data) {
      return;
    }
    if(!this.props.onSubmit) {
      this.props.onSubmit(data);
    }
  }

  onChange (val, path) {
    if(this.props.onChange) {
      this.props.onChange(val, path)
    }

    console.log("value of the form ", val, this.state)
    let fv = Object.assign({}, this.state.formValue, val)
    let st = Object.assign({}, this.state, {formValue:fv})
    if(this.lookupSchemaOptions) {
      let so = this.lookupSchemaOptions(this, fv, val, path, this.state.so)
      if(so) {
        st.so = so
        st.key = "entityform" + (new Date())
      }
    }
    this.setState(st)
  }
  fields = () => {
    let fieldsArr = new Array()
    let f = this.props.description.fields
    let uikit = this.uikit
    if(f) {
      Object.keys(f).forEach(function(k) {
        let fd = f[k]
        console.log("field found", fd)
        fieldsArr.push(
          <uikit.Block>
             <label htmlFor="firstName">{fd.label}</label>
             <Field name={fd.name} component="input" type="text" />
         </uikit.Block>
        )
      })
    }

    return fieldsArr
  }
  render() {
    console.log("rendering web form")
    return (
      <form onSubmit={this.submitForm} className="webform">
      {this.fields()}
      </form>
    )
  }
}

WebFormUI.contextTypes = {
  uikit: PropTypes.object
};
/*
  <t.form.Form ref="form" key={state.key} type={state.schema} value={state.formValue} options={state.so} onChange={this.onChange}/>
  {this.actionButtons}*/

const ReduxForm = props => {
  console.log("redux form ", props, reduxForm, WebFormUI)
  let val = reduxForm({
    // a unique name for the form
    form: props.description.id
  })(WebFormUI)
  return React.createElement(val, {description: props.description});
}

const mapStateToProps = (state, ownProps) => {
  return {
    id: ownProps.id,
    description: ownProps.description,
    refCallback: ownProps.refCallback,
    actionButtons: ownProps.actionButtons,
    onChange: ownProps.onChange,
    onSubmit: ownProps.onSubmit,
    children: ownProps.children
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  let onSubmit = null
  if(ownProps.onSubmit) {
      onSubmit = ownProps.onSubmit
  }
  else {
    let formDesc = ownProps.description
    if(formDesc.entityName) {
      switch(formDesc.formType) {
        case "Create":
          onSubmit = (data) => {
            dispatch(createAction(ActionNames.ENTITY_SAVE, {data:data, entityName: ownProps.name}, {reducer: ownProps.reducer, successCallback: ownProps.postSave, failureCallback: ownProps.failureCallback}));
          }
          break;
        case "Put":
          onSubmit = (data) => {
            dispatch(createAction(ActionNames.ENTITY_PUT, {data:data, entityId: ownProps.id, entityName: ownProps.name}, {reducer: ownProps.reducer, reload: ownProps.reloadOnUpdate, successCallback: ownProps.postSave, failureCallback: ownProps.failureCallback}));
          }
          break;
        case "Update":
          onSubmit = (data) => {
            dispatch(createAction(ActionNames.ENTITY_UPDATE, {data:data, entityId: ownProps.id, entityName: ownProps.name}, {reducer: ownProps.reducer, reload: ownProps.reloadOnUpdate, successCallback: ownProps.postSave, failureCallback: ownProps.failureCallback}));
          }
          break;
      }
    } else {
      let svc = formDesc.serviceName
      onSubmit = (data) => {
        dispatch(createAction(ActionNames.SUBMIT_FORM, data, {successCallback: ownProps.postSave, failureCallback: ownProps.failureCallback}));
      }
    }
  }
  return { onSubmit }
}

const Form = connect(
  mapStateToProps,
  mapDispatchToProps
)(ReduxForm);



export {Form} ;

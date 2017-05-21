'use strict';

import React from 'react';
import Entity from './Entity';
import t from 'tcomb-form';
import { connect } from 'react-redux';
import {ActionNames, createAction} from 'laatoocommon';

class TCombWebForm extends React.Component {
  constructor(props) {
    super(props);
    this.submitForm = this.submitForm.bind(this);
    this.setValue = this.setValue.bind(this);
    this.submit = this.submit.bind(this);
    this.getValue = this.getValue.bind(this);
    this.onChange = this.onChange.bind(this);
    let so = props.lookupSchemaOptions? props.lookupSchemaOptions(props.entityData) : props.schemaOptions
    this.state = {formValue: props.entityData, so : so, key: "entityform" + (new Date())}
    if(props.refCallback) {
      props.refCallback(this)
    }
  }
  setValue(val) {
    this.setState(Object.assign(this.state,{formValue: val}))
  }
  getValue() {
    return this.refs.form.getValue()
  }
  componentWillReceiveProps(nextprops) {
    let ed = nextprops.entityData? nextprops.entityData: this.state.formValue
    let so = this.props.lookupSchemaOptions? this.props.lookupSchemaOptions(this, ed, {}, "", this.state.so): this.props.schemaOptions
    if(so) {
      this.setState( {formValue: nextprops.entityData, so: so, key: "entityform" + (new Date())})
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
    if(!this.props.id || this.props.id==="") {
      this.props.save(data);
    } else {
      if(this.props.usePut) {
        this.props.put(data);
      } else {
        this.props.update(data);
      }
    }
  }
  onChange (val, path) {
    if(this.props.onChange) {
      this.props.onChange(val, path)
    }
    console.log("value of the entity form ", val, this.state)
    let fv = Object.assign({}, this.state.formValue, val)
    let st = Object.assign({}, this.state, {formValue:fv})
    if(this.props.lookupSchemaOptions) {
      let so = this.props.lookupSchemaOptions(this, fv, val, path, this.state.so)
      if(so) {
        st.so = so
        st.key = "entityform" + (new Date())
      }
    }
    this.setState(st)
  }
  submit() {
    if(this.props.actionButtons) {
      return this.props.actionButtons(this)
    }
    if(this.props.hideSubmit) {
      return null
    } else {
      return (
        <div className="entityformsubmit">
          <button type="submit">Save</button>
        </div>
      )
    }
  }
  render() {
    let state = this.state
    return (
      <form onSubmit={this.submitForm} className="entityform">
        <t.form.Form ref="form" key={state.key} type={this.props.schema} value={state.formValue} options={state.so} onChange={this.onChange}/>
        {this.submit()}
      </form>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    id: ownProps.id,
    name: ownProps.name,
    entityData: ownProps.entityData,
    refCallback: ownProps.refCallback,
    schema: ownProps.schema,
    preSave: ownProps.preSave,
    failureCallback: ownProps.failureCallback,
    usePut: ownProps.usePut,
    actionButtons: ownProps.actionButtons,
    schemaOptions: ownProps.schemaOptions
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    save: (data) => {
      dispatch(createAction(ActionNames.ENTITY_SAVE, {data:data, entityName: ownProps.name}, {reducer: ownProps.reducer, successCallback: ownProps.postSave, failureCallback: ownProps.failureCallback}));
    },
    put: (data) => {
      dispatch(createAction(ActionNames.ENTITY_PUT, {data:data, entityId: ownProps.id, entityName: ownProps.name}, {reducer: ownProps.reducer, reload: ownProps.reloadOnUpdate, successCallback: ownProps.postSave, failureCallback: ownProps.failureCallback}));
    },
    update: (data) => {
      dispatch(createAction(ActionNames.ENTITY_UPDATE, {data:data, entityId: ownProps.id, entityName: ownProps.name}, {reducer: ownProps.reducer, reload: ownProps.reloadOnUpdate, successCallback: ownProps.postSave, failureCallback: ownProps.failureCallback}));
    }
  }
}

const EntityForm = connect(
  mapStateToProps,
  mapDispatchToProps
)(TCombWebForm);

export {EntityForm as EntityForm} ;

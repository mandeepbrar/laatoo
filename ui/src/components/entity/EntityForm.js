'use strict';

import React from 'react';
import Entity from './Entity';
import t from 'tcomb-form';
import { connect } from 'react-redux';
import {ActionNames} from '../../actions/ActionNames';
import {createAction} from '../../utils';

class TCombWebForm extends React.Component {
  constructor(props) {
    super(props);
    this.submitForm = this.submitForm.bind(this);
    this.setValue = this.setValue.bind(this);
    this.getValue = this.getValue.bind(this);
    this.state = {formValue: props.entityData}
    if(this.props.refCallback) {
      this.props.refCallback(this)
    }
  }
  setValue(val) {
    this.setState(Object.assign(this.state,{formValue: val}))
  }
  getValue() {
    return this.refs.form.getValue()
  }
  componentWillReceiveProps(nextprops) {
    this.setState(Object.assign(this.state, {formValue: nextprops.entityData}))
    if(this.props.refCallback) {
      this.props.refCallback(this)
    }
  }
  submitForm(evt) {
    evt.preventDefault();
    let data = this.refs.form.getValue()
    if (!data) {
      return;
    }
    data = Object.assign({}, data);
    if(this.props.preSave) {
      data = this.props.preSave(data);
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

  render() {
    let state = this.state
    return (
      <form onSubmit={this.submitForm} className="entityform">
        <t.form.Form ref="form" type={this.props.schema} value={state.formValue} options={this.props.schemaOptions}/>
        <div className="entityformsubmit">
          <button type="submit">Save</button>
        </div>
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
    usePut: ownProps.usePut,
    schemaOptions: ownProps.schemaOptions
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    save: (data) => {
      dispatch(createAction(ActionNames.ENTITY_SAVE, {data:data, entityName: ownProps.name}, {reducer: ownProps.reducer}));
    },
    put: (data) => {
      dispatch(createAction(ActionNames.ENTITY_PUT, {data:data, entityId: ownProps.id, entityName: ownProps.name}, {reducer: ownProps.reducer}));
    },
    update: (data) => {
      dispatch(createAction(ActionNames.ENTITY_UPDATE, {data:data, entityId: ownProps.id, entityName: ownProps.name}, {reducer: ownProps.reducer}));
    }
  }
}

const EntityForm = connect(
  mapStateToProps,
  mapDispatchToProps
)(TCombWebForm);

export {EntityForm as EntityForm} ;

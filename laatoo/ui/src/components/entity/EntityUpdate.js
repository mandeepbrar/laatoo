'use strict';

import React from 'react';
import {EntityForm} from './EntityForm'
import t from 'tcomb-form';
import {  Response,  DataSource,  RequestBuilder } from '../../sources/DataSource';
import { browserHistory } from 'react-router';
import { connect } from 'react-redux';
import { createAction } from '../../utils';
import  {ActionNames} from '../../actions/ActionNames';

class UpdateForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {schema: props.schema};
  }
  componentDidMount() {
    this.props.loadEntity();
    if(this.props.mountForm) {
      this.props.mountForm(this);
    }
  }
  componentWillReceiveProps(nextprops) {
    if(nextprops.status == "Updated") {
      if(this.props.postSave) {
        this.props.postSave(this, nextprops);
      }
    }
  }

  render() {
    return (
      <div>
        <h1>Update Entity</h1>
        <EntityForm name={this.props.name} entityData={this.props.data} id={this.props.id} schema={this.state.schema} reducer={this.props.reducer} preSave={this.props.preSave} schemaOptions={this.props.schemaOptions}>
        </EntityForm>
      </div>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  let props = {
    name: ownProps.name,
    id: ownProps.id,
    schema: ownProps.schema,
    schemaOptions: ownProps.schemaOptions,
    reducer: ownProps.reducer,
    mountForm: ownProps.mountForm,
    preSave: ownProps.preSave,
    postSave: ownProps.postSave
  };
  if(state.router && state.router.routeStore) {
    let form = state.router.routeStore[ownProps.reducer];
    if(form) {
      props.status = form.status
      props.data = form.data
    }
  }
  return props;
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    loadEntity: () => {
      let payload = {entityName: ownProps.name, entityId: ownProps.id};
      let meta = {reducer: ownProps.reducer};
      dispatch(createAction(ActionNames.ENTITY_GET, payload, meta));
    }
  }
}

const UpdateEntity = connect(
  mapStateToProps,
  mapDispatchToProps
)(UpdateForm);

export {UpdateEntity as UpdateEntity}

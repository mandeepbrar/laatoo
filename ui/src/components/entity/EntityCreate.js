'use strict';

import React from 'react';
import {EntityForm} from './EntityForm'
import t from 'tcomb-form';
import {  Response,  DataSource,  RequestBuilder } from '../../sources/DataSource';
import { connect } from 'react-redux';
import { createAction } from '../../utils';
import  {ActionNames} from '../../actions/ActionNames';

class CreateForm extends React.Component {
  constructor(props) {
    super(props);
    this.title = this.title.bind(this);
    this.state = {schema: props.schema};
  }
  componentDidMount() {
    if(this.props.idToDuplicate) {
      this.props.loadEntity();
    }
    if(this.props.mountForm) {
      this.props.mountForm(this);
    }
  }
  componentWillReceiveProps(nextprops) {
    if(nextprops.status == "Saved") {
      if(this.props.postSave) {
        this.props.postSave(this, nextprops);
      }
    }
  }
  title() {
    if(this.props.schemaOptions && this.props.schemaOptions.template) {
      return null
    } else {
      return (
        <h1>Create {this.props.name}</h1>
      )
    }
  }
  render() {
    return (
      <div>
        {this.title()}
        <EntityForm name={this.props.name} actionButtons={this.props.actionButtons} refCallback={this.props.refCallback} schema={this.state.schema}
            entityData={this.props.data} reducer={this.props.reducer} preSave={this.props.preSave} schemaOptions={this.props.schemaOptions}>
        </EntityForm>
      </div>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  console.log("own props in create", ownProps)
  let props = {
    idToDuplicate: ownProps.idToDuplicate,
    name: ownProps.name,
    schema: ownProps.schema,
    schemaOptions: ownProps.schemaOptions,
    refCallback: ownProps.refCallback,
    reducer: ownProps.reducer,
    mountForm: ownProps.mountForm,
    actionButtons: ownProps.actionButtons,
    preSave: ownProps.preSave,
    postSave: ownProps.postSave
  };
  if(state.router && state.router.routeStore) {
    let form = state.router.routeStore[ownProps.reducer];
    if(form) {
      props.status = form.status
      props.data = form.data
      props.data.Id = ""
    }
  }
  return props;
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    loadEntity: () => {
      let payload = {entityName: ownProps.name, entityId: ownProps.idToDuplicate};
      let meta = {reducer: ownProps.reducer};
      dispatch(createAction(ActionNames.ENTITY_GET, payload, meta));
    }
  }
}

const CreateEntity = connect(
  mapStateToProps,
  mapDispatchToProps
)(CreateForm);

export {CreateEntity as CreateEntity}

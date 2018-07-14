'use strict';

import React from 'react';
import {EntityForm} from './EntityForm'
import t from 'tcomb-form';
import { createAction, ActionNames, Response,  DataSource,  RequestBuilder } from 'reactuibase';
import { connect } from 'react-redux';

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
    let schemaOptions = this.props.schemaOptions? this.props.schemaOptions: {}
    if(this.props.params) {
      if(schemaOptions.config) {
        schemaOptions.config.routeParams = this.props.params
      } else {
        schemaOptions.config = {routeParams: this.props.params}
      }
    }
    return (
      <div>
        {this.title()}
        <EntityForm name={this.props.name} actionButtons={this.props.actionButtons} refCallback={this.props.refCallback} schema={this.state.schema}
            entityData={this.props.data} reducer={this.props.reducer} preSave={this.props.preSave} schemaOptions={schemaOptions}>
        </EntityForm>
      </div>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  let props = {
    idToDuplicate: ownProps.idToDuplicate,
    name: ownProps.name,
    schema: ownProps.schema,
    schemaOptions: ownProps.schemaOptions,
    params: ownProps.params,
    refCallback: ownProps.refCallback,
    reducer: ownProps.reducer,
    mountForm: ownProps.mountForm,
    actionButtons: ownProps.actionButtons,
    preSave: ownProps.preSave,
    postSave: ownProps.postSave
  };
  let data = ownProps.data
  if(state.router && state.router.routeStore) {
    let form = state.router.routeStore[ownProps.reducer];
    if(form) {
      props.status = form.status
      if(form.data) {
        data = Object.assign({}, data, form.data)
      }
      data.Id = ""
    }
  }
  props.data = data
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

'use strict';

import React from 'react';
import {EntityForm} from './EntityForm'
import t from 'tcomb-form';
import {  Response,  DataSource,  RequestBuilder } from '../../sources/DataSource';
import { connect } from 'react-redux';

class CreateForm extends React.Component {
  constructor(props) {
    super(props);
    this.title = this.title.bind(this);
    this.state = {schema: props.schema};
  }
  componentDidMount() {
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
    if(this.props.schemaOptions.template) {
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
        <EntityForm name={this.props.name} schema={this.state.schema} reducer={this.props.reducer} preSave={this.props.preSave} schemaOptions={this.props.schemaOptions}>
        </EntityForm>
      </div>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  let props = {
    name: ownProps.name,
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
    }
  }
  return props;
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {}
}

const CreateEntity = connect(
  mapStateToProps,
  mapDispatchToProps
)(CreateForm);

export {CreateEntity as CreateEntity}

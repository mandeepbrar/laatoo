'use strict';

import React from 'react';
import { connect } from 'react-redux';
import {ActionNames} from '../../actions/ActionNames';
import {Action} from '../action/Action';
import {createAction} from '../../utils';
import {  Response,  EntityData } from '../../sources/DataSource';

class EntitiesViewTable extends React.Component {
  constructor(props) {
    super(props);
    this.deleteEntity = this.deleteEntity.bind(this);
    this.getRow = this.getRow.bind(this);
  }
  componentDidMount() {
    if(this.props.load) {
      this.props.loadView();
    }
  }
  componentWillReceiveProps(nextprops) {
    if(nextprops.load) {
      nextprops.loadView();
    }
  }
  deleteEntity(params) {
    let table = this;
    let successMethod = function(response) {
      table.props.loadView();
    };
    let failureMethod = function(errorResponse) {
    };
    EntityData.DeleteEntity(this.props.name, params.id).then(successMethod, failureMethod);
  }
  getRow(x, i) {
    let id = x[this.props.idField];
    let title = x[this.props.titleField];
    let encodedid = encodeURIComponent(id)
    return (
      <tr key={i + 1}>
        <td style={{width:"40%"}}>
          <Action name={"Edit "+this.props.name} params={{ id: encodedid }}>{title}</Action>
        </td>
        <td>
          <Action name={"Edit "+this.props.name} params={{ id: encodedid }}>{id}</Action>
        </td>
        <td>
          <Action name={"Delete "+this.props.name} method={this.deleteEntity} params={{ id: encodedid }}>delete</Action>
        </td>
      </tr>
    )

  }
  render() {
    return (
      <div className="container">
        <div className="row">
          <Action className="pull-right  m20" widget="button" key={"Create "+this.props.name} name={"Create "+this.props.name}>{"Create "+this.props.name}</Action>
        </div>
        <table className="table table-striped ">
          <thead>
            <tr>
              <th>
                Entity
              </th>
              <th>
                Id
              </th>
              <th>
              </th>
            </tr>
          </thead>
          <tbody>
          {[...this.props.items].map((x, i) =>
            this.getRow(x,i)
          )}
          </tbody>
        </table>
      </div>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  let props = {
    name: ownProps.name,
    reducer: ownProps.reducer,
    idField: ownProps.idField,
    titleField: ownProps.titleField,
    load: false,
    items: []
  };
  if(state.router && state.router.routeStore) {
    let entityView = state.router.routeStore[ownProps.reducer];
    if(entityView) {
      if(entityView.status == "Loaded") {
          props.items = entityView.data
          return props
      }
      if(entityView.status == "NotLoaded") {
          props.load = true
          return props
      }
    }
  }
  return props;
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    loadView: () => {
      let payload = {};
      let meta = {serviceName: ownProps.viewService, reducer: ownProps.reducer};
      dispatch(createAction(ActionNames.VIEW_FETCH, payload, meta));
    }
  }
}

const EntityView = connect(
  mapStateToProps,
  mapDispatchToProps
)(EntitiesViewTable);

export {EntityView as EntityView}

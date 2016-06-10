'use strict';

import React from 'react';
import { connect } from 'react-redux';
import {ActionNames} from '../../actions/ActionNames';
import {Action} from '../action/Action';
import {createAction} from '../../utils';
import {  Response,  EntityData } from '../../sources/DataSource';
import Paginator from 'react-pagify';
import pagifyBootstrapPreset from 'react-pagify-preset-bootstrap';
import segmentize from 'segmentize';

class EntitiesViewTable extends React.Component {
  constructor(props) {
    super(props);
    this.deleteEntity = this.deleteEntity.bind(this);
    this.getRow = this.getRow.bind(this);
    this.getHeader = this.getHeader.bind(this);
    this.pagination = this.pagination.bind(this);
  }
  componentDidMount() {
    if(this.props.load) {
      this.props.loadView(this.props.currentPage);
    }
  }
  componentWillReceiveProps(nextprops) {
    if(nextprops.load) {
      nextprops.loadView(nextprops.currentPage);
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
  getRow(view, x, i) {
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
  getHeader() {
    return (
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
    )
  }
  pagination() {
    if(this.props.paginate) {
      let pages = this.props.totalPages
      let page = this.props.currentPage
      return (
        <Paginator.Context
          {...pagifyBootstrapPreset}
          segments={segmentize({
              pages,
              page,
              beginPages: 1,
              endPages: 1,
              sidePages: 3
          })}
          onSelect={(newPage, event) => {
              event.preventDefault();
              this.props.loadView(newPage);
          }}
        >
          <Paginator.Button page={page - 1}>Previous</Paginator.Button>
          <Paginator.Segment field="beginPages" />
          <Paginator.Ellipsis previousField="beginPages" nextField="previousPages" />
          <Paginator.Segment field="previousPages" />
          <Paginator.Segment field="centerPage" className="active" />
          <Paginator.Segment field="nextPages" />
          <Paginator.Ellipsis previousField="nextPages" nextField="endPages" />
          <Paginator.Segment field="endPages" />
          <Paginator.Button page={page + 1}>Next</Paginator.Button>
        </Paginator.Context>
      )
    } else {
      return null
    }
  }
  render() {
    return (
      <div className="container">
        <div className="row">
          <Action className="pull-right  m20" widget="button" key={"Create "+this.props.name} name={"Create "+this.props.name}>{"Create "+this.props.name}</Action>
        </div>
        <table className="table table-striped ">
          <thead>
            {this.props.header? this.props.header(): this.getHeader()}
          </thead>
          <tbody>
          {[...this.props.items].map((x, i) =>
            this.props.row? this.props.row(this, x, i): this.getRow(this, x,i)
          )}
          </tbody>
        </table>
        {this.pagination()}
      </div>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  let props = {
    name: ownProps.name,
    reducer: ownProps.reducer,
    idField: ownProps.idField,
    header: ownProps.header,
    paginate: ownProps.paginate,
    pageSize: ownProps.pageSize,
    currentPage: 1,
    totalPages: 1,
    row: ownProps.row,
    titleField: ownProps.titleField,
    load: false,
    items: []
  };
  if(state.router && state.router.routeStore) {
    let entityView = state.router.routeStore[ownProps.reducer];
    if(entityView) {
      if(entityView.status == "Loaded") {
          props.items = entityView.data
          props.currentPage = entityView.currentPage
          props.totalPages = entityView.totalPages
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
    loadView: (pagenum) => {
      let queryParams={}
      if (ownProps.paginate) {
            queryParams.pagesize = ownProps.pageSize;
            queryParams.pagenum = pagenum;
        }
      console.log("query params", queryParams)
      let payload = {queryParams};
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

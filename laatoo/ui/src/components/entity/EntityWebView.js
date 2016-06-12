'use strict';

import React from 'react';
import {ActionNames} from '../../actions/ActionNames';
import {Action} from '../action/Action';
import {createAction} from '../../utils';
import {WebTableView} from '../view/WebTableView';
import {  Response,  EntityData } from '../../sources/DataSource';

class EntityView extends React.Component {
  constructor(props) {
    super(props);
    this.name = this.props.name;
    this.deleteEntity = this.deleteEntity.bind(this);
  }
  componentDidMount() {
      this.refs.view.addMethod("deleteItem", this.deleteEntity)
  }
  deleteEntity(params) {
    let methods = this.refs.view.methods();
    let successMethod = function(response) {
      methods.reload();
    };
    let failureMethod = function(errorResponse) {
    };
    EntityData.DeleteEntity(this.name, params.id).then(successMethod, failureMethod);
  }
  render() {
    return (
      <div className="container">
        <div className="row ma20">
          <Action className="pull-right  m20" widget="button" key={"Create "+this.name} name={"Create "+this.name}>{"Create "+this.name}</Action>
        </div>
        <WebTableView className="ma20" ref="view" name={this.name}
          filterTitle="Search" filterGo="Go" filterForm={this.props.filterForm} defaultFilter={this.props.defaultFilter}
          idField={this.props.idField} paginate={this.props.paginate} pageSize={this.props.pageSize} getHeader={this.props.getHeader}
          getItem={this.props.getItem} reducer={this.props.reducer} titleField={this.props.titleField} viewService={this.props.viewService}
          urlParams={this.props.urlParams} postArgs={this.props.postArgs} currentPage={this.props.currentPage}>
        </WebTableView>
      </div>
    )
  }
}


export {EntityView as EntityView}

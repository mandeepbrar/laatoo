'use strict';

import React from 'react';
import {ActionNames, createAction,  Response,  EntityData } from 'reactuibase';
import {Action} from '../action/Action';
import {WebTableView} from '../view/WebTableView';

class EntityView extends React.Component {
  constructor(props) {
    super(props);
    this.name = this.props.name;
    this.deleteEntity = this.deleteEntity.bind(this);
    if(props.actions) {
      this.actions = props.actions
    } else {
      this.actions = [
        <Action className="rightalign  m20" widget="button" key={"Create "+this.name} name={"Create "+this.name}>{"Create "+this.name}</Action>
      ]
    }

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
      <div className="entitywebview">
        <div className="ma20 actionbar" style={{display:"block", clear:"both"}}>
          {this.actions}
        </div>
        <WebTableView className="ma20" ref="view" name={this.name}
          filterTitle="Search" filterGo="Go" filterForm={this.props.filterForm} defaultFilter={this.props.defaultFilter}
          idField={this.props.idField} loader={this.props.loader} paginate={this.props.paginate} pageSize={this.props.pageSize} getHeader={this.props.getHeader}
          getItem={this.props.getItem} reducer={this.props.reducer} titleField={this.props.titleField} viewService={this.props.viewService}
          urlParams={this.props.urlParams} postArgs={this.props.postArgs} currentPage={this.props.currentPage}>
        </WebTableView>
      </div>
    )
  }
}


export {EntityView as EntityView}

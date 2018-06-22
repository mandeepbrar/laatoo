'use strict';

import React from 'react';
import {ActionNames} from 'reactuibase';
import {Action} from '../action/Action';
import {WebView} from './WebView'

class WebTableView extends React.Component {
  constructor(props) {
    super(props);
    this.getItemGroup = this.getItemGroup.bind(this);
    this.getView = this.getView.bind(this);
    this.getItem = this.getItem.bind(this);
    this.getHeader = this.getHeader.bind(this);
    this.methods = this.methods.bind(this);
    this.addMethod = this.addMethod.bind(this);
  }
  methods() {
    return this.refs.view.methods()
  }
  addMethod(name, method) {
    this.refs.view.addMethod(name, method)
  }
  getView(view, header, groups) {
    return (
      <table  className="webtableview" style={{width: '100%'}}>
        <thead>
          {header}
        </thead>
        <tbody>
          {groups}
        </tbody>
      </table>
    )
  }
  getItemGroup(view, x) {
    return x
  }
  getHeader(view) {
    if(this.props.getHeader) {
      return this.props.getHeader(view)
    }
    return(
      <tr>
        <th>
          Title
        </th>
        <th>
        </th>
        <th>
        </th>
      </tr>
    )
  }
  getItem(view, x, i) {
    if(this.props.getItem) {
      return this.props.getItem(view, x, i)
    }
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
          <Action name={"Delete "+this.props.name} method={view.methods.deleteItem} params={{ id: encodedid }}>delete</Action>
        </td>
      </tr>
    )
  }

  render() {
    return (
      <WebView
        ref="view"
        reducer = {this.props.reducer}
        paginate = {this.props.paginate}
        pageSize = {this.props.pageSize}
        viewService = {this.props.viewService}
        urlParams = {this.props.urlParams}
        postArgs = {this.props.postArgs}
        defaultFilter = {this.props.defaultFilter}
        loader = {this.props.loader}
        currentPage = {this.props.currentPage}
        filterTitle= {this.props.filterTitle}
        filterForm={this.props.filterForm}
        filterGo={ this.props.filterGo}
        getFilter={this.props.getFilter}
        globalReducer={this.props.globalReducer}
        getView =  {this.getView}
        getItem = {this.getItem}
        getItemGroup = {this.getItemGroup}
        getHeader = {this.getHeader}
        className={this.props.className}
      >
      </WebView>
    )
  }
}

export {
  WebTableView
}

'use strict';

import React from 'react';
import {ActionNames} from '../../actions/ActionNames';
import {Action} from '../action/Action';
import {WebView} from './WebView'
import {ScrollListener} from '../main/ScrollListener';

class WebListView extends React.Component {
  constructor(props) {
    super(props);
    this.getItemGroup = this.getItemGroup.bind(this);
    this.getView = this.getView.bind(this);
    this.getItem = this.getItem.bind(this);
    this.getHeader = this.getHeader.bind(this);
    this.methods = this.methods.bind(this);
    this.addMethod = this.addMethod.bind(this);
    this.onScrollEnd = this.onScrollEnd.bind(this);
  }
  methods() {
    return this.refs.view.methods()
  }
  addMethod(name, method) {
    this.refs.view.addMethod(name, method)
  }
  onScrollEnd() {
    let methods = this.methods();
    methods.loadMore();
  }
  getView(view, header, groups) {
    if(this.props.getView) {
        return this.props.getView(view, header, groups)
    }
    if(this.props.incrementalLoad) {
      return (
        <ScrollListener key={this.props.key} className={this.props.className} style={this.props.style} onScrollEnd={this.onScrollEnd}>
          <div className={this.props.headerClass}>
            {header}
          </div>
          <div className={this.props.contentClass}>
            {groups}
          </div>
        </ScrollListener>
      )
    } else {
      return (
        <div key={this.props.key} >
          <div className={this.props.headerClass}>
            {header}
          </div>
          <div className={this.props.contentClass}>
            {groups}
          </div>
        </div>
      )
    }
  }
  getItemGroup(view, x) {
    return x
  }
  getHeader(view) {
    if(this.props.getHeader) {
      return this.props.getHeader(view)
    }
    return null
  }
  getItem(view, x, i) {
    if(this.props.getItem) {
      return this.props.getItem(view, x, i)
    }
    return React.Children.map(this.props.children, (child) => React.cloneElement(child, { item: x, index: i }) );
  }

  render() {
    return (
      <WebView
        ref="view"
        key={this.props.key}
        reducer = {this.props.reducer}
        paginate = {this.props.paginate}
        pageSize = {this.props.pageSize}
        hidePaginationControl = {this.props.hidePaginationControl}
        viewService = {this.props.viewService}
        urlParams = {this.props.urlParams}
        postArgs = {this.props.postArgs}
        defaultFilter = {this.props.defaultFilter}
        currentPage = {this.props.currentPage}
        filterTitle= {this.props.filterTitle}
        loader = {this.props.loader}
        filterForm={this.props.filterForm}
        style={this.props.style}
        globalReducer={this.props.globalReducer}
        filterGo={ this.props.filterGo}
        getFilter={this.props.getFilter}
        getView =  {this.getView}
        getItem = {this.getItem}
        incrementalLoad={this.props.incrementalLoad}
        getItemGroup = {this.getItemGroup}
        getHeader = {this.getHeader}
        className={this.props.className}
      >
      </WebView>
    )
  }
}

export {
  WebListView
}

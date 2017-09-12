'use strict';

import React from 'react';
import { connect } from 'react-redux';
import {ActionNames} from '../actions/ActionNames';
import {createAction} from 'uicommon';
import {ViewData} from './ViewData'

class View extends React.Component {
  constructor(props) {
    super(props);
    this.getItemGroup = this.getItemGroup.bind(this);
    this.getView = this.getView.bind(this);
    this.renderView = this.renderView.bind(this);
    this.getItem = this.getItem.bind(this);
    this.getHeader = this.getHeader.bind(this);
    this.getPagination = this.getPagination.bind(this);
    this.getFilter = this.getFilter.bind(this);
    this.addMethod = this.addMethod.bind(this);
    this.numItems = 0
  }

  addMethod(name, method) {
    return this.viewdata.addMethod(name, method)
  }

  getView(header, groups, pagination, filter) {
    if(this.props.getView) {
      return this.props.getView(this, header, groups, pagination, filter)
    }
    return null
  }
  getFilter(filterTitle, filterForm, filterGo) {
    if(this.props.getFilter) {
      return this.props.getFilter(this, filterTitle, filterForm, filterGo)
    }
    return null
  }
  getItemGroup(x) {
    if(this.props.getItemGroup) {
      return this.props.getItemGroup(this, x)
    }
    return null
  }
  getItem(x, i) {
    if(this.props.getItem) {
      return this.props.getItem(this, x, i)
    }
    return null
  }
  getHeader() {
    if(this.props.getHeader) {
      return this.props.getHeader(this)
    }
    return null
  }
  getPagination() {
    if(this.props.paginate && this.props.getPagination) {
        let pages = this.props.totalPages
        let page = this.props.currentPage
        return this.props.getPagination(this, pages, page)
    }
    return null
  }
  renderView(viewdata, items, currentPage, totalPages) {
    this.viewdata = viewdata
    let groups=[]
    let groupsize = 1
    let group=[]

    if(items) {
      let keys = Object.keys(items);
      this.numItems = keys.length
      for (var i in keys) {
        let x = items[keys[i]]
        if (x) {
          let item = this.getItem(x, keys[i])
          group.push(item)
          if((i % groupsize) == 0) {
            let itemGrp = this.getItemGroup(group)
            groups.push(itemGrp)
            group = []
          }
        }
      }
    } else {
      if(this.props.loader) {
        groups.push(this.props.loader)
      }
    }
    let header = this.getHeader()
    let filterCtrl = this.getFilter(this.props.filterTitle, this.props.filterForm, this.props.filterGo, this.filter)
    let pagination = this.getPagination()
    return this.getView(header, groups, pagination, filterCtrl)
  }

  render() {
    return (
      <ViewData
        getView={this.renderView}
        key={this.props.key}
        reducer={this.props.reducer}
        paginate={this.props.paginate}
        pageSize={this.props.pageSize}
        viewService={this.props.viewService}
        urlParams = {this.props.urlParams}
        postArgs = {this.props.postArgs}
        defaultFilter = {this.props.defaultFilter}
        currentPage={this.props.currentPage}
        style={this.props.style}
        className={this.props.className}
        incrementalLoad={this.props.incrementalLoad}
        globalReducer={this.props.globalReducer}>
      </ViewData>
    )
  }
}

export {View as View}

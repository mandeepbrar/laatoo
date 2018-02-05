'use strict';

import React from 'react';
import {ViewData} from './ViewData'
const PropTypes = require('prop-types');

class ViewUI extends ViewData {
  constructor(props, context) {
    super(props);
    this.getItemGroup = this.getItemGroup.bind(this);
    this.getView = this.getView.bind(this);
    this.renderView = this.renderView.bind(this);
    this.getItem = this.getItem.bind(this);
    this.getHeader = this.getHeader.bind(this);
    this.getPagination = this.getPagination.bind(this);
    this.getFilter = this.getFilter.bind(this);
    this.uikit = context.uikit;
    this.div = this.uikit.Block;
    this.scroll = this.uikit.Scroll;
    this.onScrollEnd = this.onScrollEnd.bind(this);
    this.numItems = 0
  }

  onItemCheckboxChange(view){
    return (evt)=> {
      let cb = evt.target
      let item = view.refs[cb.value]
      if (cb.checked) {
        item.selected = true
      } else {
        item.selected = false
      }
    }
  }

/*
getFilter(view, filterTitle, filterForm, filterGo, filter) {
  if(this.props.getFilter) {
    return this.props.getFilter(filterTitle, filterForm, filterGo, filter)
  }
  if(!filterForm) {
    return null
  }
  if(!filterTitle) {
    filterTitle="Search"
  }
  if(!filterGo) {
    filterGo = "Go"
  }
  return (
    <ViewFilter title={filterTitle} schema={filterForm} defaultFilter={filter} setFilter={view.methods.setFilter} goBtn={filterGo} >
      <div className="row m20">
        <i className="fa fa-search pull-right"></i>
      </div>
    </ViewFilter>
  )
}
*/
  onScrollEnd() {
    let methods = this.methods();
    methods.loadMore();
  }
  getView(header, groups, pagination, filter) {
    if(this.props.editable) {
      this.addMethod('onItemCheckboxChange', this.onItemCheckboxChange(view))
    }
    if(this.props.getView) {
        return this.props.getView(this, header, groups, pagination, filter, this.props)
    }
    if(this.props.incrementalLoad) {
      return (
        <this.uikit.scroll key={this.props.key} className={this.props.className} onScrollEnd={this.onScrollEnd}>
          {filter}
          {header}
          {groups}
          {pagination}
        </this.uikit.scroll>
      )
    } else {
      return (
        <this.uikit.Block key={this.props.key} className={this.props.className} style={this.props.style} >
        {filter}
        {header}
        {groups}
        {pagination}
        </this.uikit.Block>
      )
    }
  }

  getFilter() {
    if(this.props.getFilter) {
      return this.props.getFilter(this, this.props.defaultFilter)
    }
    return null
  }
  getItemGroup(x) {
    if(this.props.getItemGroup) {
      return this.props.getItemGroup(this, x)
    }
    return <this.uikit.Block className="group">x</this.uikit.Block>
  }
  getRenderedItem = (x, i) => {
    console.log("get rendered item", this.props.children, this.props)
    return React.Children.map(this.props.children, (child) => React.cloneElement(child, { data: x, index: i }) )
  }
  getItem(x, i) {
    if(this.props.getItem) {
      return this.props.getItem(this, x, i)
    }
    return this.getRenderedItem(x, i);
  }
  getHeader() {
    if(this.props.getHeader) {
      return this.props.getHeader(this)
    } else if (this.props.header) {
      return this.props.header
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
  renderView(items, currentPage, totalPages) {
    //this.viewdata = viewdata
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
          if(groupsize == 1) {
            groups.push(item)
          } else {
            group.push(item)
            if((i % groupsize) == 0) {
              let itemGrp = this.getItemGroup(group)
              groups.push(itemGrp)
              group = []
            }
          }
        }
      }
    } else {
      if(this.props.loader) {
        groups.push(this.props.loader)
      }
    }
    let header = this.getHeader()
    let filterCtrl = this.getFilter()
    let pagination = this.getPagination()
    return this.getView(header, groups, pagination, filterCtrl)
  }
/*
  render() {
    return (
      <ViewData
        getView={this.renderView}
        key={this.props.key}
        reducer={this.props.reducer}
        paginate={this.props.paginate}
        pageSize={this.props.pageSize}
        viewService={this.props.viewService}
        loader = {this.props.loader}
        urlParams = {this.props.urlParams}
        postArgs = {this.props.postArgs}
        defaultFilter = {this.props.defaultFilter}
        currentPage={this.props.currentPage}
        style={this.props.style}
        className={this.props.className}
        incrementalLoad={this.props.incrementalLoad}
        getPagination={this.props.incrementalLoad || this.props.hidePaginationControl ? null : this.props.getPagination} >
      </ViewData>
    )
  }*/
}

ViewUI.contextTypes = {
  uikit: PropTypes.object
};
export {ViewUI as ViewUI}

import React from 'react';
import {ViewData, ViewItem} from './ViewData'
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
    this.onScrollEnd = this.onScrollEnd.bind(this);
  }

  onScrollEnd() {
    let methods = this.methods();
    methods.loadMore();
  }
  getView(header, groups, pagination, filter) {
    console.log("view ui getView", this.props)
    if(this.props.contentOnly) {
      return groups
    }
    if(this.props.getView) {
        return this.props.getView(this, header, groups, pagination, filter, this.props)
    }
    if(this.props.incrementalLoad) {
      return (
        <_uikit.scroll key={this.props.key} className={this.props.className} onScrollEnd={this.onScrollEnd}>
          {filter}
          {header}
          {groups}
          {pagination}
        </_uikit.scroll>
      )
    } else {
      return (
        <_uikit.Block key={this.props.key} className={this.props.className} style={this.props.style} >
        {filter}
        {header}
        {groups}
        {pagination}
        </_uikit.Block>
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
    return <_uikit.Block className="group">x</_uikit.Block>
  }
  getRenderedItem = (x, i) => {
    console.log("get rendered item", this.props.children, this.props)
    return React.Children.map(this.props.children, (child) => React.cloneElement(child, { data: x, index: i }) )
  }
  getItem(x, i) {
    let renderedComp = null;
    if(this.props.getItem) {
      renderedComp = this.props.getItem(this, x, i);
    } else {
      renderedComp = this.getRenderedItem(x, i);
    }
    let viewItem = new ViewItem()
    viewItem.index = i;
    viewItem.data = x;
    viewItem.renderedItem = renderedComp;
    console.log("pushing item", viewItem)
    super.pushItem(viewItem)
    return renderedComp
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
      //super.setNumItems(keys.length);
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

export {ViewUI as ViewUI}

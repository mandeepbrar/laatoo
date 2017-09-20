'use strict';

import React from 'react';

class ViewData extends React.Component {
  constructor(props) {
    super(props);
    this.setPage = this.setPage.bind(this);
    this.selectedItems = this.selectedItems.bind(this);
    this.itemStatus = this.itemStatus.bind(this);
    this.viewrefs = this.viewrefs.bind(this);
    this.itemCount = this.itemCount.bind(this);
    this.reload = this.reload.bind(this);
    this.setFilter = this.setFilter.bind(this);
    this.loadMore = this.loadMore.bind(this);
    this.canLoadMore = this.canLoadMore.bind(this);
    //this.getView = this.getView.bind(this);
    this.methods = {reload: this.reload, canLoadMore: this.canLoadMore, loadMore: this.loadMore, setFilter:this.setFilter,
        itemCount: this.itemCount, viewrefs: this.viewrefs, itemStatus: this.itemStatus,
        selectedItems: this.selectedItems, setPage: this.setPage}
    this.addMethod = this.addMethod.bind(this);
    this.state = {lastLoadTime: -1}
    this.numItems = 0
  }
  componentWillMount() {
    this.filter = this.props.defaultFilter
  }
  componentDidMount() {
    if(this.props.load && !this.props.externalLoad) {
      this.props.loadView(this.props.currentPage, this.filter);
    }
  }
  componentWillReceiveProps(nextprops) {
    if(nextprops.load) {
      nextprops.loadView(nextprops.currentPage, this.filter);
    }
  }
  shouldComponentUpdate(nextProps, nextState) {
    if(!nextProps.forceUpdate && this.lastRenderTime) {
      if(nextProps.lastUpdateTime) {
        if(this.lastRenderTime >= nextProps.lastUpdateTime) {
          return false
        }
      } else {
        return false
      }
    }
    return true;
  }
  addMethod(name, method) {
    this.methods[name] = method
  }
  reload() {
    this.props.loadView(this.props.currentPage, this.filter);
  }
  canLoadMore() {
    return this.props.currentPage < this.props.totalPages
  }
  viewrefs() {
    return this.refs
  }
  itemCount() {
    return this.numItems
  }
  selectedItems() {
    let selectedItems = []
    for(var i=0; i<this.numItems;i++) {
      let refName = "item"+i
      let item = this.refs[refName]
      if(item.selected) {
        selectedItems.push(item.id)
      }
    }
    return selectedItems
  }
  itemStatus() {
    let items = {}
    for(var i=0; i<this.numItems;i++) {
      let refName = "item"+i
      let item = this.refs[refName]
      items[item.id] = item.selected
    }
    return items
  }
  setPage(newPage) {
    this.props.loadView(newPage, this.filter)
  }
  setFilter(filter) {
    this.filter = filter
    this.props.loadView(1, this.filter)
  }
/*  getView(items, currentPage, totalPages) {
    if(this.props.getView) {
      return this.props.getView(this, items, currentPage, totalPages)
    }
    return null
  }*/
  loadMore() {
    if(this.props.currentPage>=this.props.totalPages) {
      return false
    } else {
      if(this.props.currentPage) {
        this.props.loadIncrementally(this.props.currentPage + 1, this.filter)
        return true
      }
    }
  }
  render() {
    this.lastRenderTime = this.props.lastUpdateTime
    let view = this.renderView(this.props.items, this.props.currentPage, this.props.totalPages)
    this.items=this.props.items
    return view
  }
}


export {ViewData }

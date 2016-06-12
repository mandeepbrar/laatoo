'use strict';

import React from 'react';
import { connect } from 'react-redux';
import {ActionNames} from '../../actions/ActionNames';
import {Action} from '../action/Action';
import {createAction} from '../../utils';
import {  Response } from '../../sources/DataSource';

class ViewDisplay extends React.Component {
  constructor(props) {
    super(props);
    this.getItemGroup = this.getItemGroup.bind(this);
    this.getView = this.getView.bind(this);
    this.getItem = this.getItem.bind(this);
    this.getHeader = this.getHeader.bind(this);
    this.getPagination = this.getPagination.bind(this);
    this.setPage = this.setPage.bind(this);
    this.selectedItems = this.selectedItems.bind(this);
    this.itemStatus = this.itemStatus.bind(this);
    this.viewrefs = this.viewrefs.bind(this);
    this.itemCount = this.itemCount.bind(this);
    this.getFilter = this.getFilter.bind(this);
    this.reload = this.reload.bind(this);
    this.setFilter = this.setFilter.bind(this);
    this.methods = {reload: this.reload, setFilter:this.setFilter, itemCount: this.itemCount, viewrefs: this.viewrefs, itemStatus: this.itemStatus, selectedItems: this.selectedItems, setPage: this.setPage}
    this.addMethod = this.addMethod.bind(this);
    this.numItems = 0
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
  addMethod(name, method) {
    this.methods[name] = method
  }
  reload() {
    this.props.loadView(this.props.currentPage);
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
    this.props.loadView(newPage)
  }
  setFilter(filter) {
    console.log("filter", filter)
    this.props.loadView(1, filter)
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
  render() {
    let groups=[]
    let groupsize = 1
    let group=[]
    this.numItems = this.props.items.length
    for (var i in this.props.items) {
      let x = this.props.items[i]
      let item = this.getItem(x, i)
      group.push(item)
      if((i % groupsize) == 0) {
        let itemGrp = this.getItemGroup(group)
        groups.push(itemGrp)
        group = []
      }
    }
    let header = this.getHeader()
    let filter = this.getFilter(this.props.filterTitle, this.props.filterForm, this.props.filterGo)
    let pagination = this.getPagination()
    let view = this.getView(header, groups, pagination, filter)
    return view
  }
}

const mapStateToProps = (state, ownProps) => {
  let props = {
    reducer: ownProps.reducer,
    paginate: ownProps.paginate,
    pageSize: ownProps.pageSize,
    filterTitle: ownProps.filterTitle,
    filterForm: ownProps.filterForm,
    filterGo: ownProps.filterGo,
    getView: ownProps.getView,
    getItem: ownProps.getItem,
    getItemGroup: ownProps.getItemGroup,
    getHeader: ownProps.getHeader,
    getFilter: ownProps.getFilter,
    getPagination: ownProps.getPagination,
    currentPage: ownProps.currentPage,
    className: ownProps.className,
    totalPages: 1,
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
    loadView: (pagenum, filter) => {
      if(!pagenum) {
        pagenum = 1
      }
      let queryParams={}
      if (ownProps.paginate) {
            queryParams.pagesize = ownProps.pageSize;
            queryParams.pagenum = pagenum;
        }
      let postArgs = Object.assign({}, ownProps.viewParams, filter);
      let payload = {queryParams, postArgs};
      let meta = {serviceName: ownProps.viewService, reducer: ownProps.reducer};
      dispatch(createAction(ActionNames.VIEW_FETCH, payload, meta));
    }
  }
}

const View = connect(
  mapStateToProps,
  mapDispatchToProps,
  null,
  {withRef: true}
)(ViewDisplay);

export {View as View}

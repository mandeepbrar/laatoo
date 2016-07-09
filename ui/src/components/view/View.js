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
    this.loadMore = this.loadMore.bind(this);
    this.methods = {reload: this.reload, loadMore: this.loadMore, setFilter:this.setFilter, itemCount: this.itemCount, viewrefs: this.viewrefs, itemStatus: this.itemStatus, selectedItems: this.selectedItems, setPage: this.setPage}
    this.addMethod = this.addMethod.bind(this);
    this.state = {lastLoadTime: -1}
    this.numItems = 0
  }
  componentWillMount() {
    this.filter = this.props.defaultFilter
  }
  componentDidMount() {
    if(this.props.load) {
      this.props.loadView(this.props.currentPage, this.filter);
    }
  }
  componentWillReceiveProps(nextprops) {
    if(nextprops.load) {
      nextprops.loadView(nextprops.currentPage, this.filter);
    }
  }
  shouldComponentUpdate(nextProps, nextState) {
    if(this.lastRenderTime) {
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
    let groups=[]
    let groupsize = 1
    let group=[]

    if(this.props.items) {
      let keys = Object.keys(this.props.items);
      this.numItems = keys.length
      for (var i in keys) {
        let x = this.props.items[keys[i]]
        let item = this.getItem(x, i)
        group.push(item)
        if((i % groupsize) == 0) {
          let itemGrp = this.getItemGroup(group)
          groups.push(itemGrp)
          group = []
        }
      }
    } else {
      if(this.props.loader) {
        groups.push(this.props.loader)
      }
    }
    this.items=this.props.items
    let header = this.getHeader()
    let filterCtrl = this.getFilter(this.props.filterTitle, this.props.filterForm, this.props.filterGo, this.filter)
    let pagination = this.getPagination()
    let view = this.getView(header, groups, pagination, filterCtrl)
    return view
  }
}

const mapStateToProps = (state, ownProps) => {
  let props = {
    reducer: ownProps.reducer,
    paginate: ownProps.paginate,
    pageSize: ownProps.pageSize,
    defaultFilter: ownProps.defaultFilter,
    urlParams: ownProps.urlParams,
    postArgs: ownProps.postArgs,
    loader : ownProps.loader,
    filterTitle: ownProps.filterTitle,
    filterForm: ownProps.filterForm,
    filterGo: ownProps.filterGo,
    getView: ownProps.getView,
    getItem: ownProps.getItem,
    incrementalLoad: ownProps.incrementalLoad,
    getItemGroup: ownProps.getItemGroup,
    getHeader: ownProps.getHeader,
    getFilter: ownProps.getFilter,
    getPagination: ownProps.getPagination,
    currentPage: ownProps.currentPage,
    className: ownProps.className,
    totalPages: 1,
    load: false,
    items: null
  };
  let view = null;
  if(!ownProps.globalReducer) {
    if(state.router && state.router.routeStore) {
      view = state.router.routeStore[ownProps.reducer];
    }
  } else {
    view = state[ownProps.reducer];
  }
  if(view) {
    if(view.status == "Loaded") {
        props.items = view.data
        props.currentPage = view.currentPage
        props.totalPages = view.totalPages
        props.lastUpdateTime = view.lastUpdateTime
        props.latestPageData = view.latestPageData
        return props
    }
    if(view.status == "NotLoaded") {
        props.load = true
        return props
    }
  }
  return props;
}

function loadData(dispatch, ownProps, pagenum, filter, incrementalLoad) {
  if(!pagenum) {
    pagenum = 1
  }
  let queryParams={}
  if (ownProps.paginate) {
        queryParams.pagesize = ownProps.pageSize;
        queryParams.pagenum = pagenum;
    }
  let postArgs = Object.assign({}, ownProps.postArgs, filter);
  let payload = {queryParams, postArgs};
  let meta = {serviceName: ownProps.viewService, reducer: ownProps.reducer, incrementalLoad: incrementalLoad};
  dispatch(createAction(ActionNames.VIEW_FETCH, payload, meta));
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    loadView: (pagenum, filter) => {
      loadData(dispatch, ownProps, pagenum, filter, false)
    },
    loadIncrementally: (pagenum, filter) => {
      loadData(dispatch, ownProps, pagenum, filter, true)
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

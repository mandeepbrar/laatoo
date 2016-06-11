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
  setPage(newPage) {
    this.props.loadView(newPage)
  }
  getView(header, groups, pagination) {
    if(this.props.getView) {
      return this.props.getView(this, header, groups, pagination)
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
    let pagination = this.getPagination()
    let view = this.getView(header, groups, pagination)
    return view
  }
}

const mapStateToProps = (state, ownProps) => {
  let props = {
    reducer: ownProps.reducer,
    paginate: ownProps.paginate,
    pageSize: ownProps.pageSize,
    getView: ownProps.getView,
    getItem: ownProps.getItem,
    getItemGroup: ownProps.getItemGroup,
    getHeader: ownProps.getHeader,
    getPagination: ownProps.getPagination,
    currentPage: ownProps.currentPage,
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
    loadView: (pagenum) => {
      if(!pagenum) {
        pagenum = 1
      }
      let queryParams={}
      if (ownProps.paginate) {
            queryParams.pagesize = ownProps.pageSize;
            queryParams.pagenum = pagenum;
        }
      console.log("query params", queryParams)
      let payload = {queryParams};
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

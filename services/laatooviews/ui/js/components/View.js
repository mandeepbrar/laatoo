import { connect } from 'react-redux';
import {ActionNames} from '../Actions';
import {createAction} from 'uicommon';
import {ViewUI} from './ViewUI';
import React from 'react';

function getSvc(ownProps) {
  return ownProps.serviceName? ownProps.serviceName: ownProps.name
}

const mapStateToProps = (state, ownProps) => {
  let svc = getSvc(ownProps);
  let viewname = ownProps.name;
  let red = ownProps.reducer? ownProps.reducer: svc;
  let props = {
    name: viewname,
    global: ownProps.global,
    paginate: ownProps.paginate,
    pageSize: ownProps.pageSize,
    header: ownProps.header,
    getView: ownProps.getView,
    getItem: ownProps.getItem,
    getHeader: ownProps.getHeader,
    defaultFilter: ownProps.defaultFilter,
    externalLoad: ownProps.externalLoad,
    urlParams: ownProps.urlParams,
    postArgs: ownProps.postArgs,
    getView: ownProps.getView,
    incrementalLoad: ownProps.incrementalLoad,
    currentPage: ownProps.currentPage,
    className: ownProps.className,
    ref: ownProps.viewRef,
    loader:ownProps.loader,
    getPagination: ownProps.incrementalLoad || ownProps.hidePaginationControl ? null : ownProps.getPagination,
    style: ownProps.style,
    totalPages: 1,
    load: false,
    items: null
  };
  let viewReducer = state["views"];
  if(viewReducer && viewname) {
    let view = viewReducer.views[viewname]
    if(view) {
      if(view.status == "Loaded") {
          props.items = view.data
          props.currentPage = view.currentPage
          props.totalPages = view.totalPages
          props.lastUpdateTime = view.lastUpdateTime
          props.latestPageData = view.latestPageData
      }
      if(view.status == "NotLoaded") {
          props.load = true
      }
    } else {
      props.load = true
    }
  }
  return props;
}

function loadData(dispatch, ownProps, pagenum, filter, incrementalLoad) {
  let svc = getSvc(ownProps);
  console.log("service.....", svc, ownProps)
  if(!pagenum) {
    pagenum = 1
  }
  let queryParams={}
  if (ownProps.paginate) {
        queryParams.pagesize = ownProps.pageSize;
        queryParams.pagenum = pagenum;
    }
  let serviceObject = ownProps.dataurl?{url: ownProps.dataurl, method: "POST"}:null
  let postArgs = Object.assign({}, ownProps.postArgs, filter);
  let payload = {queryParams, postArgs};
  let meta = {serviceName: svc, global: ownProps.global, viewname: ownProps.name, serviceObject: serviceObject, incrementalLoad: incrementalLoad};
  dispatch(createAction(ActionNames.VIEW_FETCH, payload, meta));
}

const mapDispatchToProps = (dispatch, ownProps) => {
  console.log("view comp properties", ownProps);
  return {
    loadView: (pagenum, filter) => {
      loadData(dispatch, ownProps, pagenum, filter, false)
    },
    loadIncrementally: (pagenum, filter) => {
      loadData(dispatch, ownProps, pagenum, filter, true)
    }
  }
}
console.log("react-redux connect in laatooviews", connect, require('react-redux'));
const ViewComponent = connect(
  mapStateToProps,
  mapDispatchToProps
)(ViewUI);

const View = (props) => {
  console.log("props of the view...", props)
  let view = props.description;
  if(!view && props.id) {
    view = _reg("Views", props.id)
  }
  if(view) {
    console.log("view.....", view)
    let args = props.postArgs? props.postArgs: view.postArgs;
    let params = props.urlparams? props.urlparams: view.urlparams;
    let viewname = view.name? view.name : props.id;
    let className = props.className+ " view_"+viewname + " ";
    if(view.className) {
      className = className + view.className;
    }
    let item = props.children;
    return <ViewComponent serviceObject={view.service} serviceName={view.serviceName} name={viewname} global={view.global}
      className={className} incrementalLoad={view.incrementalLoad} paginate={view.paginate} header={props.header} getHeader={props.getHeader}
       editable={props.editable} getView={props.getView} getItem={props.getItem} urlparams={params} postArgs={args} viewRef={props.viewRef}>
       {item}
       </ViewComponent>
  }
  return null
}

export {
  View,
  ViewComponent
}

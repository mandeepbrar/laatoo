import { connect } from 'react-redux';
import {ActionNames} from '../Actions';
import {createAction} from 'uicommon';

function getSvc(ownProps) {
  return ownProps.dataservice? ownProps.dataservice: ownProps.name
}

const mapStateToProps = (state, ownProps) => {
  let svc = getSvc(ownProps);
  let red = ownProps.reducer? ownProps.reducer: svc;
  let props = {
    reducer: red,
    global: ownProps.global,
    paginate: ownProps.paginate,
    pageSize: ownProps.pageSize,
    defaultFilter: ownProps.defaultFilter,
    externalLoad: ownProps.externalLoad,
    urlParams: ownProps.urlParams,
    postArgs: ownProps.postArgs,
    getView: ownProps.getView,
    incrementalLoad: ownProps.incrementalLoad,
    currentPage: ownProps.currentPage,
    className: ownProps.className,
    loader:ownProps.loader,
    getPagination:{ownProps.incrementalLoad || ownProps.hidePaginationControl ? null : ownProps.getPagination},
    style: ownProps.style,
    totalPages: 1,
    load: false,
    items: null
  };
  let view = null;
  if(!ownProps.global) {
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
  let svc = getSvc(ownProps);
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
  let meta = {serviceName: svc, reducer: svc, serviceObject: serviceObject, incrementalLoad: incrementalLoad};
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
  mapDispatchToProps
)(ViewUI);

export {
  View
}

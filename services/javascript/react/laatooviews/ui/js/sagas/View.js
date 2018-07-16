import { takeEvery, takeLatest } from 'redux-saga'
import { call, put } from 'redux-saga/effects'
import  {ActionNames} from '../Actions';
import { createAction, Response,  DataSource,  RequestBuilder } from 'uicommon';

function* fetchViewData(action) {
  try {
    yield put(createAction(ActionNames.VIEW_FETCHING, action.payload,{global: action.meta.global, viewname: action.meta.viewname}));
    let req = RequestBuilder.DefaultRequest(action.payload.queryParams, action.payload.postArgs, action.payload.headers);
    let resp = null
    if(action.meta.serviceObject) {
      resp = yield call(DataSource.ExecuteServiceObject, action.meta.serviceObject, req);
    } else {
      resp = yield call(DataSource.ExecuteService, action.meta.serviceName, req);
    }
    yield put(createAction(ActionNames.VIEW_FETCH_SUCCESS, resp.data, {info: resp.info, incrementalLoad: action.meta.incrementalLoad, global: action.meta.global, viewname: action.meta.viewname}));
  } catch (e) {
    yield put(createAction(ActionNames.VIEW_FETCH_FAILED, e, action.meta));
    if(action.meta.failureCallback) {
      action.meta.failureCallback(e)
    } else {
      if(Window.handleError) {
        Window.handleError(e)
      }
    }
  }
}

function* viewSaga() {
  yield* takeEvery(ActionNames.VIEW_FETCH, fetchViewData);
}

//export {viewSaga as viewSaga};
Application.Register('Sagas', "views", viewSaga)

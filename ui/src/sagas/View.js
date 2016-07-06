import { takeEvery, takeLatest } from 'redux-saga'
import { call, put } from 'redux-saga/effects'
import  {ActionNames} from '../actions/ActionNames';
import {  Response,  DataSource,  RequestBuilder } from '../sources/DataSource';
import { createAction } from '../utils';
import 'babel-polyfill';

function* fetchViewData(action) {
  try {
    yield put(createAction(ActionNames.VIEW_FETCHING, action.payload,{reducer: action.meta.reducer}));
    let req = RequestBuilder.DefaultRequest(action.payload.queryParams, action.payload.postArgs);
    const resp = yield call(DataSource.ExecuteService, action.meta.serviceName, req);
    yield put(createAction(ActionNames.VIEW_FETCH_SUCCESS, resp.data, {info: resp.info, incrementalLoad: action.meta.incrementalLoad, reducer: action.meta.reducer}));
  } catch (e) {
    yield put(createAction(ActionNames.VIEW_FETCH_FAILED, e, {reducer: action.meta.reducer, incrementalLoad: action.meta.incrementalLoad}));
  }
}

function* viewSaga() {
  yield* takeEvery(ActionNames.VIEW_FETCH, fetchViewData);
}

export {viewSaga as viewSaga};

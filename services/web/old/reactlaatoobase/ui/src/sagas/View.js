import { takeEvery, takeLatest } from 'redux-saga'
import { call, put } from 'redux-saga/effects'
import  {ActionNames} from '../actions/ActionNames';
import { createAction, Response,  DataSource,  RequestBuilder } from 'uicommon';

function* fetchViewData(action) {
  try {
    yield put(createAction(ActionNames.VIEW_FETCHING, action.payload,{reducer: action.meta.reducer}));
    let req = RequestBuilder.DefaultRequest(action.payload.queryParams, action.payload.postArgs, action.payload.headers);
    const resp = yield call(DataSource.ExecuteService, action.meta.serviceName, req);
    yield put(createAction(ActionNames.VIEW_FETCH_SUCCESS, resp.data, {info: resp.info, incrementalLoad: action.meta.incrementalLoad, reducer: action.meta.reducer}));
  } catch (e) {
    yield put(createAction(ActionNames.VIEW_FETCH_FAILED, e, action.meta));
    if(action.meta.failureCallback) {
      action.meta.failureCallback(e)
    } else {
      if(window.handleError) {
        window.handleError(e)
      }
    }
  }
}

function* viewSaga() {
  yield takeEvery(ActionNames.VIEW_FETCH, fetchViewData);
}

//export {viewSaga as viewSaga};
Application.Register('Sagas', "viewSaga", viewSaga)

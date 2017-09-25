import { takeEvery, takeLatest } from 'redux-saga'
import { call, put } from 'redux-saga/effects'
import  {ActionNames} from '../Actions';
import {  createAction, Response,  EntityData } from 'uicommon';

function* getEntityViewData(action) {
  try {
    yield put(createAction(ActionNames.ENTITY_VIEW_FETCHING, action.payload, {global: action.meta.global, entityId: payload.entityId}));
    const resp = yield call(EntityData.GetEntity, action.payload.entityName, action.payload.entityId, action.payload.headers, action.payload.svc);
    resp.data.isOwner = (resp.data.CreatedBy === Storage.user);
    yield put(createAction(ActionNames.ENTITY_VIEW_FETCH_SUCCESS, resp, action.meta));
    if(action.meta.successCallback) {
      action.meta.successCallback({resp: resp, payload: action.payload})
    }
  } catch (e) {
    yield put(createAction(ActionNames.ENTITY_VIEW_FETCH_FAILED, e, action.meta));
    if(action.meta.failureCallback) {
      action.meta.failureCallback(e)
    } else {
      if(window.handleError) {
        window.handleError(e)
      }
    }
  }
}


function* entityViewSaga() {
  yield [
    takeEvery(ActionNames.ENTITY_VIEW_FETCH, getEntityViewData)
  ]
}

//export {entitySaga as entitySaga};
Application.Register('Sagas', "entityViewSaga", entityViewSaga)

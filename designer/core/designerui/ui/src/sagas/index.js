import { takeEvery, takeLatest } from 'redux-saga'
import { call, put } from 'redux-saga/effects'
import Actions from '../actions';
import {  createAction, Response, RequestBuilder, DataSource, EntityData } from 'uicommon';

function* syncObjects(action) {
  console.log("syncing objects", action)
/*  try {
    yield put(createAction(ActionNames.ENTITY_GETTING, action.payload, {}));
    const resp = yield call(EntityData.GetEntity, action.payload.entityName, action.payload.entityId, action.payload.headers, action.payload.svc);
    resp.data.isOwner = (resp.data.CreatedBy === Storage.user);
    yield put(createAction(ActionNames.ENTITY_GET_SUCCESS, resp, action.meta));
    if(action.meta.successCallback) {
      action.meta.successCallback({resp: resp, payload: action.payload})
    }
  } catch (e) {
    yield put(createAction(ActionNames.ENTITY_GET_FAILED, e, action.meta));
    if(action.meta.failureCallback) {
      action.meta.failureCallback(e)
    } else {
      if(Window.handleError) {
        Window.handleError(e)
      }
    }
  }*/
}

//console.log("Action names ", ActionNames)
function* designerSaga() {
  yield [
    takeEvery(Actions.SYNC_OBJECTS, syncObjects)
  ]
}
//takeEvery(ActionNames.ENTITY_DELETE, deleteEntityData)

//export {entitySaga as entitySaga};
Application.Register('Sagas', "designerSaga", designerSaga)

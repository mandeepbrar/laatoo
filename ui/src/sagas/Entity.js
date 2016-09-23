import { takeEvery, takeLatest } from 'redux-saga'
import { call, put } from 'redux-saga/effects'
import  {ActionNames} from '../actions/ActionNames';
import {  Response,  EntityData } from '../sources/DataSource';
import { createAction } from '../utils';
import 'babel-polyfill';

function* getEntityData(action) {
  try {
    yield put(createAction(ActionNames.ENTITY_GETTING, action.payload, {reducer: action.meta.reducer}));
    const resp = yield call(EntityData.GetEntity, action.payload.entityName, action.payload.entityId);
    resp.data.isOwner = (resp.data.CreatedBy === localStorage.user);
    yield put(createAction(ActionNames.ENTITY_GET_SUCCESS, resp, action.meta));
    if(action.meta.successCallback) {
      action.meta.successCallback(action.payload)
    }
  } catch (e) {
    yield put(createAction(ActionNames.ENTITY_GET_FAILED, e, action.meta));
    if(action.meta.failureCallback) {
      action.meta.failureCallback(e)
    } else {
      if(window.handleError) {
        window.handleError(e)
      }
    }
  }
}

function* deleteEntityData(action) {
  try {
    yield put(createAction(ActionNames.ENTITY_DELETING, action.payload, {reducer: action.meta.reducer}));
    const resp = yield call(EntityData.DeleteEntity, action.payload.entityName, action.payload.entityId);
    yield put(createAction(ActionNames.ENTITY_DELETE_SUCCESS, resp, action.meta));
    if(action.meta.successCallback) {
      action.meta.successCallback(action.payload)
    }
  } catch (e) {
    yield put(createAction(ActionNames.ENTITY_DELETE_FAILURE, e, action.meta));
    if(action.meta.failureCallback) {
      action.meta.failureCallback(e)
    } else {
      if(window.handleError) {
        window.handleError(e)
      }
    }
  }
}

function* saveEntityData(action) {
  try {
    yield put(createAction(ActionNames.ENTITY_SAVING, action.payload, {reducer: action.meta.reducer}));
    const resp = yield call(EntityData.SaveEntity, action.payload.entityName, action.payload.data);
    yield put(createAction(ActionNames.ENTITY_SAVE_SUCCESS, {}, {reducer: action.meta.reducer}));
    if(action.meta.successCallback) {
      action.meta.successCallback(action.payload)
    }
  } catch (e) {
    yield put(createAction(ActionNames.ENTITY_SAVE_FAILURE, e, action.meta));
    if(action.meta.failureCallback) {
      action.meta.failureCallback(e)
    } else {
      if(window.handleError) {
        window.handleError(e)
      }
    }
  }
}

function* putEntityData(action) {
  try {
    yield put(createAction(ActionNames.ENTITY_PUTTING, action.payload, {reducer: action.meta.reducer}));
    const resp = yield call(EntityData.PutEntity, action.payload.entityName, action.payload.entityId, action.payload.data);
    yield put(createAction(ActionNames.ENTITY_PUT_SUCCESS, {}, action.meta));
    if(action.meta.reload) {
      yield put(createAction(ActionNames.ENTITY_GET, action.payload, action.meta));
    }
    if(action.meta.successCallback) {
      action.meta.successCallback(action.payload)
    }
  } catch (e) {
    yield put(createAction(ActionNames.ENTITY_PUT_FAILURE, e, action.meta));
    if(action.meta.failureCallback) {
      action.meta.failureCallback(e)
    } else {
      if(window.handleError) {
        window.handleError(e)
      }
    }
  }
}

function* updateEntityData(action) {
  try {
    yield put(createAction(ActionNames.ENTITY_UPDATING, action.payload, {reducer: action.meta.reducer}));
    const resp = yield call(EntityData.UpdateEntity, action.payload.entityName, action.payload.entityId, action.payload.data);
    yield put(createAction(ActionNames.ENTITY_UPDATE_SUCCESS, {}, action.meta));
    if(action.meta.reload) {
      yield put(createAction(ActionNames.ENTITY_GET, action.payload, action.meta));
    }
    if(action.meta.successCallback) {
      action.meta.successCallback(action.payload)
    }
  } catch (e) {
    yield put(createAction(ActionNames.ENTITY_UPDATE_FAILURE, e, action.meta));
    if(action.meta.failureCallback) {
      action.meta.failureCallback(e)
    } else {
      if(window.handleError) {
        window.handleError(e)
      }
    }
  }
}


function* entitySaga() {
  yield [
    takeEvery(ActionNames.ENTITY_GET, getEntityData),
    takeEvery(ActionNames.ENTITY_SAVE, saveEntityData),
    takeEvery(ActionNames.ENTITY_UPDATE, updateEntityData),
    takeEvery(ActionNames.ENTITY_PUT, putEntityData),
    takeEvery(ActionNames.ENTITY_DELETE, deleteEntityData)
  ]
}

export {entitySaga as entitySaga};
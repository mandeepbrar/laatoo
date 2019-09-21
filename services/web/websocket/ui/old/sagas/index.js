import { takeEvery, takeLatest } from 'redux-saga'
import { call, put } from 'redux-saga/effects'
import  Actions from './Actions';
import {  createAction, Response,  EntityData } from 'uicommon';
const Sockette = require('sockette');

var websocket

function InitializeSocket(url, options) {
    websocket = new Sockette(url, {
        timeout: 500,
        maxAttempts: 10,
        onopen: e => console.log('Connected!', e),
        onmessage: e => console.log('Received:', e),
        onreconnect: e => console.log('Reconnecting...', e),
        onmaximum: e => console.log('Stop Attempting!', e),
        onclose: e => console.log('Closed!', e),
        onerror: e => console.log('Error:', e)
      });
}

function* sendMessage(action) {
  try {
    if(websocket) {
        websocket.json(action.payload);
    }
    /*yield put(createAction(ActionNames.ENTITY_VIEW_FETCHING, action.payload, meta));
    const resp = yield call(EntityData.GetEntity, action.payload.entityName, action.payload.entityId, action.payload.headers, action.payload.svc);
    resp.data.isOwner = (resp.data.CreatedBy === Storage.user);
    yield put(createAction(ActionNames.ENTITY_VIEW_FETCH_SUCCESS, resp, meta));*/
    if(action.meta.successCallback) {
      action.meta.successCallback(action.payload)
    }
  } catch (e) {
    if(action.meta.failureCallback) {
      action.meta.failureCallback(e)
    } else {
      if(Window.handleError) {
        Window.handleError(e)
      }
    }
  }
}


function* entityViewSaga() {
  yield takeEvery(Actions.SEND_SOCKET_MESSAGE, sendMessage)
}

//export {entitySaga as entitySaga};
Application.Register('Sagas', "sendMessageSaga", sendMessage)

export {
    InitializeSocket
}
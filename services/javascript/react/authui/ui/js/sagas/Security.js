import { takeEvery, takeLatest } from 'redux-saga'
import { call, put } from 'redux-saga/effects'
import  {ActionNames} from '../actions';
import {  createAction, Response,  DataSource,  RequestBuilder } from 'uicommon';

function* login(action) {
  try {
    yield put(createAction(ActionNames.LOGGING_IN));
    let req = RequestBuilder.DefaultRequest(null, action.payload);
    const resp = yield call(DataSource.ExecuteService, action.meta.serviceName, req);
    let authToken = Application.Security.AuthToken.toLowerCase();
    let token = resp.info[authToken];
    let user = resp.data;
    let userId = resp.data.Id;
    let permissions = resp.data.Permissions;
    let loginaction = createAction(ActionNames.LOGIN_SUCCESS, {userId, token, permissions, user});
    yield put(loginaction);
    console.log("dispatched login action &&&&")
  } catch (e) {
    yield put(createAction(ActionNames.LOGIN_FAILURE, e));
  }
}

function* logout(action) {
  yield put(createAction(ActionNames.LOGOUT_SUCCESS, {}));
}

function* loginSaga() {
  yield [
    takeLatest(ActionNames.LOGIN, login),
    takeLatest(ActionNames.LOGOUT, logout)
  ];
}

//export {loginSaga as loginSaga};
Application.Register('Sagas', "loginSaga", loginSaga)

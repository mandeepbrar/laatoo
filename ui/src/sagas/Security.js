import { takeEvery, takeLatest } from 'redux-saga'
import { call, put } from 'redux-saga/effects'
import  {ActionNames} from '../actions/ActionNames';
import {  Response,  DataSource,  RequestBuilder } from '../sources/DataSource';
import {createAction} from '../utils';
import {Application, Storage} from '../Globals'

function* login(action) {
  try {
    yield put(createAction(ActionNames.LOGGING_IN));
    let req = RequestBuilder.DefaultRequest(null, action.payload);
    const resp = yield call(DataSource.ExecuteService, action.meta.serviceName, req);
    let authToken = Application.Security.AuthToken.toLowerCase();
    let token = resp.info[authToken];
    let userId = resp.data.Id;
    let permissions = resp.data.Permissions;
    let loginaction = createAction(ActionNames.LOGIN_SUCCESS, {userId, token, permissions});
    yield put(loginaction);
  } catch (e) {
    yield put(createAction(ActionNames.LOGIN_FAILURE, e));
  }
}

function* loginSaga() {
  yield* takeLatest(ActionNames.LOGIN, login);
}

export {loginSaga as loginSaga};

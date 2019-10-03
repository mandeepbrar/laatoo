import { takeEvery, takeLatest, fork, call, put } from 'redux-saga/effects'
import  {ActionNames} from '../actions';
import {  createAction, Response,  DataSource,  RequestBuilder } from 'uicommon';


function* signup(action) {
  try {
    yield put(createAction(ActionNames.SIGNING_UP));
    let req = RequestBuilder.DefaultRequest(null, action.payload);
    const resp = yield call(DataSource.ExecuteService, action.meta.serviceName, req);
    let signupaction = createAction(ActionNames.SIGNUP_SUCCESS, {});
    yield put(signupaction);
    console.log("dispatched signup action success");
  } catch (e) {
    yield put(createAction(ActionNames.SIGNUP_FAILURE, e));
    Window.handleError(e);
  }
}


function* login(action) {
  try {
    console.log("received login action", action)
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
    Window.handleError(e);
  }
}

function* logout(action) {
  let req = RequestBuilder.DefaultRequest(null, null);
  const resp = yield call(DataSource.ExecuteService, action.meta.serviceName, req);
  console.log("Logout response", resp)
  yield put(createAction(ActionNames.LOGOUT_SUCCESS, {}));
}

function* loginSaga() {
  yield takeLatest(ActionNames.LOGIN, login)
}

function* signupSaga() {
  yield takeLatest(ActionNames.SIGN_UP, signup)
}

function* logoutSaga() {
  yield takeLatest(ActionNames.LOGOUT, logout)  
}

function* authSaga() {
  console.log("take latest in auth saga", fork)
  yield fork(loginSaga)
  yield fork(signupSaga)
  yield fork(logoutSaga)
}

//export {loginSaga as loginSaga};
Application.Register('Sagas', "authSaga", authSaga)

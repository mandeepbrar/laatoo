import { takeEvery, call, put } from 'redux-saga/effects'
import Actions from '../actions';
import {  createAction, Response, RequestBuilder, DataSource, EntityData } from 'uicommon';

function* syncObjects(action) {
  console.log("syncing objects", action)
  try {
    let req = RequestBuilder.DefaultRequest(null, action.payload);
    const resp = yield call(DataSource.ExecuteService, action.meta.type+"resolver", req);
    console.log("resolver", resp)
  } catch (e) {
    console.log("sync objects", e)
    if(Window.handleError) {
      Window.handleError(e)
    }
  }
}

//console.log("Action names ", ActionNames)
function* designerSaga() {
  yield takeEvery(Actions.SYNC_OBJECTS, syncObjects)
}
//takeEvery(ActionNames.ENTITY_DELETE, deleteEntityData)

//export {entitySaga as entitySaga};
Application.Register('Sagas', "designerSaga", designerSaga)

import { takeEvery, takeLatest } from 'redux-saga'
import { call, put } from 'redux-saga/effects'
import  {ActionNames} from '../Actions';
import { createAction, Response,  EntityData, DataSource, RequestBuilder } from 'uicommon';

function* loadGroup(action) {
  try {
    let request = {}
    console.log("load group saga", action.payload)
    Object.keys(action.payload).forEach(function (key) {
      let service = action.payload[key]
      console.log("load group ", service, "key", key)
      if(service.type == "entity") {
        request[key] = {Params:{id: service.entityId}, Body: {}}
      }
      if(service.type == "view") {
        request[key] = {Params: service.queryParams, Body: service.postArgs}
      }
    });
    /*Object.keys(action.payload).forEach(function* (service, key) {
      console.log("load group saga", service, "key", key)
      if(service.type == "entity") {
        yield put(createAction(ActionNames.ENTITY_GETTING, service.payload, {reducer: service.meta.reducer}));
      }
      if(service.type == "view") {
        yield put(createAction(ActionNames.VIEW_FETCHING, service.payload,{reducer: service.meta.reducer}));
      }
    });*/
    console.log("created request", request)
    let req = RequestBuilder.DefaultRequest(null, request);
    const resp = yield call(DataSource.ExecuteService, action.meta.serviceName, req);
    console.log("resp", resp)
    let actions = new Array()
    Object.keys(action.payload).forEach(function (key) {
      let service = action.payload[key]
      let servResponse = resp.data[key]
      console.log("service ", service, servResponse)
      let itemResp = {data: servResponse.Data, statuscode: servResponse.Status, info: servResponse.Info}
      if(service.type == "entity") {
        if(servResponse.Status == 200) {
          itemResp.data.isOwner = (itemResp.data.CreatedBy === Storage.user);
          actions.push(createAction(ActionNames.ENTITY_GET_SUCCESS, itemResp, {reducer: service.meta.reducer}));
        } else {
          actions.push(createAction(ActionNames.ENTITY_GET_FAILED, itemResp,  {reducer: service.meta.reducer}));
        }
      }
      if(service.type == "view") {
        if(servResponse.Status == 200) {
          actions.push(createAction(ActionNames.VIEW_FETCH_SUCCESS, itemResp.data, {info: itemResp.info, incrementalLoad: action.meta.incrementalLoad, reducer: service.meta.reducer}));
        } else {
          actions.push(createAction(ActionNames.VIEW_FETCH_FAILED, itemResp.data, {reducer: service.reducer}));
        }
      }
    });
    if(action.meta.successCallback) {
      action.meta.successCallback({resp: resp, payload: action.payload})
    }
    console.log("actions i", actions)
    for (let i=0; i<actions.length; i++) {
      console.log("actions i",i,  actions[i])
      yield put(actions[i])
    }
  } catch (e) {
    if (action.meta.services) {
      let actions = new Array()
      Object.keys(action.meta.services).forEach(function (key) {
        let service = action.meta.services[key]
        if(service.type == "entity") {
          actions.push(createAction(ActionNames.ENTITY_GET_FAILED, e,  {reducer: service.reducer}));
        }
        if(service.type == "view") {
          actions.push(createAction(ActionNames.VIEW_FETCH_FAILED, e,{reducer: service.reducer}));
        }
      });
      for (let i=0; i<actions.length; i++) {
        console.log("actions i",i,  actions[i])
        yield put(actions[i])
      }
    }
    if(action.meta.failureCallback) {
      action.meta.failureCallback(e)
    } else {
      if(Window.handleError) {
        Window.handleError(e)
      }
    }
  }
}

function* groupLoadSaga() {
  yield [
    takeEvery(ActionNames.GROUP_LOAD, loadGroup)
  ]
}

Application.Register('Sagas', "groupLoadSaga", groupLoadSaga)

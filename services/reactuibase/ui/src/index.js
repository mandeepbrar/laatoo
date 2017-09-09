const redux = require('redux');
import {Application, Storage, Window} from './Globals'
import { Response, DataSource, RequestBuilder, EntityData } from './sources/DataSource';
import {Reducers} from './reducers';
import {ActionNames} from './actions/ActionNames';
import {createAction, formatUrl, LaatooError, hasPermission} from './utils';
import {DisplayEntity} from './entity/EntityDisplay';
import {ViewReducer} from './reducers/View';
import {EntityReducer} from './reducers/Entity';
import createSagaMiddleware from 'redux-saga';
import {Sagas} from './sagas';
import {GroupLoad} from './components/GroupLoad';
import {View} from './components/View';
import {ViewData} from './components/ViewData';
import {LoginComponent} from './components/LoginComponent';
import 'babel-polyfill'
import Color from './colors'
import {LoginValidator} from './components/LoginValidator';
/*

*/
/*
function createStore(reducers, initialState, middleware, sagas, enhancers) {
  const sagaMiddleware = createSagaMiddleware();
  enhancers = redux.compose(redux.applyMiddleware(sagaMiddleware, ...middleware), ...enhancers);
  if(!reducers) {
    reducers = {};
  }
  // mount it on the Store
  const store = redux.createStore( redux.combineReducers(reducers), initialState, enhancers);

  // then run the saga
  runSagas(sagaMiddleware, sagas);
  return store;
}*/

function RegisterRoute(routeId, routeData) {
  Application.Routes[routeId] = routeData
}

function RegisterReducer(reducerId, reducer) {
  Application.Reducers[reducerId] = reducer
}

function RegisterSaga(saga) {
  Application.Sagas.push(saga)
}

console.log("color from laatoo ", Color);

export {
  Storage,
  Application,
  Window,
  LoginValidator,
  RequestBuilder,
  DisplayEntity,
  DataSource,
  Response,
  EntityData,
  Reducers,
  ViewReducer,
  Colors,
  View,
  ViewData,
  EntityReducer,
  LoginComponent,
  ActionNames,
  formatUrl,
  createAction,
  LaatooError,
  hasPermission,
  RegisterRoute,
  RegisterReducer,
  RegisterSaga,
  Sagas
}

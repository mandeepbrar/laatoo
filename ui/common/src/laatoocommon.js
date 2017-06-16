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
import {Sagas, runSagas} from './sagas';
import {GroupLoad} from './components/GroupLoad';
import {View} from './components/View';
import {ViewData} from './components/ViewData';
import {LoginComponent} from './components/LoginComponent';
import 'babel-polyfill'
import GurmukhiKeymap from './utils/gurmukhikeymap'
import Color from './colors'
import {LoginValidator} from './components/LoginValidator';
/*

*/

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
}

console.log("color from laatoo ", Color)
let moduleExports = {
  Storage: Storage,
  Application: Application,
  Window: Window,
  GurmukhiKeymap: GurmukhiKeymap,
  LoginValidator: LoginValidator,
  RequestBuilder: RequestBuilder,
  DisplayEntity: DisplayEntity,
  DataSource:DataSource,
  Response: Response,
  EntityData: EntityData,
  Reducers: Reducers,
  ViewReducer: ViewReducer,
  Colors: Color,
  View: View,
  ViewData: ViewData,
  EntityReducer: EntityReducer,
  LoginComponent: LoginComponent,
  ActionNames: ActionNames,
  formatUrl: formatUrl,
  createStore: createStore,
  createAction: createAction,
  LaatooError: LaatooError,
  hasPermission: hasPermission,
  GroupLoad: GroupLoad,
  Sagas: Sagas
}

module.exports = moduleExports;

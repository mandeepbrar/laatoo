const redux = require('redux');
import {Application, Storage, Window} from './Globals'
import { Response, DataSource, RequestBuilder, EntityData } from './sources/DataSource';
import {Reducers} from './reducers';
import {ActionNames} from './actions/ActionNames';
import {createAction, formatUrl, LaatooError, hasPermission} from './utils';
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


let moduleExports = {
  Storage: Storage,
  Application: Application,
  Window: Window,
  GurmukhiKeymap: GurmukhiKeymap,
  RequestBuilder: RequestBuilder,
  DataSource:DataSource,
  Response: Response,
  EntityData: EntityData,
  Reducers: Reducers,
  ViewReducer: ViewReducer,
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

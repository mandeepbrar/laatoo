const redux = require('redux');
import {Application, Storage, Window} from './Globals'
import { Response, DataSource, RequestBuilder, EntityData } from './sources/DataSource';
import {Reducers} from './reducers';
import {Action} from './components/action/Action';
import {ActionNames} from './actions/ActionNames';
import {createAction, formatUrl} from './utils';
import {ViewReducer} from './reducers/View';
import {EntityReducer} from './reducers/Entity';
import createSagaMiddleware from 'redux-saga';
import {Sagas, runSagas} from './sagas';
import {LoginComponent} from './components/login/LoginComponent';
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
  RequestBuilder: RequestBuilder,
  Window: Window,
  DataSource:DataSource,
  Response: Response,
  EntityData: EntityData,
  Reducers: Reducers,
  ViewReducer: ViewReducer,
  EntityReducer: EntityReducer,
  LoginComponent: LoginComponent,
  Action: Action,
  ActionNames: ActionNames,
  formatUrl: formatUrl,
  createStore: createStore,
  createAction: createAction,
  Sagas: Sagas
}

//export {LoginComponent as LoginComponent};//
//export {DataSource as DataSource};
module.exports = moduleExports;

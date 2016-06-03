require('styles/App.css')

const redux = require('redux');
import LoginComponent from './components/login/LoginComponent';
import { Response, DataSource, RequestBuilder, EntityData } from './sources/DataSource';
import {Reducers} from './reducers';
import {Action} from './components/action/Action';
import {Entity} from './components/entity/Entity';
import {EntityForm} from './components/entity/EntityForm';
import {ActionNames} from './actions/ActionNames';
import createSagaMiddleware from 'redux-saga';
import {createAction} from './utils';
import {formatUrl} from './utils';
import {Sagas, runSagas} from './sagas';

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

//export {LoginComponent as LoginComponent};//
//export {DataSource as DataSource};
module.exports = {
    LoginComponent: LoginComponent,
    DataSource:DataSource,
    Response: Response,
    Reducers: Reducers,
    RequestBuilder: RequestBuilder,
    Action: Action,
    ActionNames: ActionNames,
    Entity: Entity,
    EntityData: EntityData,
    EntityForm:EntityForm,
    createStore: createStore,
    createAction: createAction,
    formatUrl: formatUrl,
    Sagas: Sagas
};

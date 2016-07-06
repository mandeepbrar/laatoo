require('styles/App.css')

const redux = require('redux');
import LoginComponent from './components/login/LoginComponent';
import { Response, DataSource, RequestBuilder, EntityData } from './sources/DataSource';
import {Reducers} from './reducers';
import {Action} from './components/action/Action';
import {Entity} from './components/entity/Entity';
import {DisplayEntity} from './components/entity/EntityDisplay';
import {EntityForm} from './components/entity/EntityForm';
import {ActionNames} from './actions/ActionNames';
import createSagaMiddleware from 'redux-saga';
import {createAction} from './utils';
import {formatUrl} from './utils';
import {Sagas, runSagas} from './sagas';
import {VideoEdit} from './components/form/videoedit';
import {TextEdit} from './components/form/textedit';
import {ImageEdit} from './components/form/imageedit';
import {WebTableView} from './components/view/WebTableView';
import {ViewReducer} from './reducers/View';
import {EntityReducer} from './reducers/Entity';
import {ViewFilter} from './components/view/Filter';
import {Image} from './components/main/Image';
import {View} from './components/view/View';
import {WebView} from './components/view/WebView';
import {WebListView} from './components/view/WebListView';
import {Html} from './components/main/Html';
import {ScrollListener} from './components/main/ScrollListener';

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
    ViewReducer: ViewReducer,
    EntityReducer: EntityReducer,
    Action: Action,
    ActionNames: ActionNames,
    Entity: Entity,
    DisplayEntity: DisplayEntity,
    EntityData: EntityData,
    EntityForm:EntityForm,
    createStore: createStore,
    createAction: createAction,
    WebTableView: WebTableView,
    VideoEdit: VideoEdit,
    TextEdit: TextEdit,
    ScrollListener: ScrollListener,
    WebView: WebView,
    WebListView: WebListView,
    Html : Html,
    View: View,
    Image: Image,
    ImageEdit: ImageEdit,
    ViewFilter: ViewFilter,
    formatUrl: formatUrl,
    Sagas: Sagas
};

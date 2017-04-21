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
  Window: Window,
  RequestBuilder: RequestBuilder,
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

if(!Application.native) {
  let videoedit = require('./components/form/videoedit');
  let textedit = require('./components/form/textedit');
  let imageedit = require('./components/form/imageedit');
  let webtableview = require('./components/view/WebTableView');
  let entity = require('./components/entity/Entity');
  let entitydisplay = require('./components/entity/EntityDisplay');
  let entityform = require('./components/entity/EntityForm');
  let entityupdate = require('./components/entity/EntityUpdate');
  require('./styles/App.css');
  require('babel-polyfill');
  let viewfilter = require('./components/view/Filter');
  let image = require('./components/main/Image');
  let view = require('./components/view/View');
  let webview = require('./components/view/WebView');
  let weblistview = require('./components/view/WebListView');
  let html = require('./components/main/Html');
  let scrolllistener = require('./components/main/ScrollListener');
  let groupload = require( './components/main/GroupLoad');

  moduleExports = Object.assign(moduleExports, {
    GroupLoad: groupload.GroupLoad,
    Entity: entity.Entity,
    DisplayEntity: entitydisplay.DisplayEntity,
    EntityForm: entityform.EntityForm,
    UpdateEntity: entityupdate.UpdateEntity,
    WebTableView: webtableview.WebTableView,
    VideoEdit: videoedit.VideoEdit,
    TextEdit: textedit.TextEdit,
    RichEdit: textedit.RichEdit,
    ScrollListener: scrolllistener.ScrollListener,
    WebView: webview.WebView,
    WebListView: weblistview.WebListView,
    Html : html.Html,
    View: view.View,
    Image: image.Image,
    ImageChooser: imageedit.ImageChooser,
    ImageEdit: imageedit.ImageEdit,
    ViewFilter: viewfilter.ViewFilter
  })
}

//export {LoginComponent as LoginComponent};//
//export {DataSource as DataSource};
module.exports = moduleExports;

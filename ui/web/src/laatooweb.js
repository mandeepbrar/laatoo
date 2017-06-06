const redux = require('redux');
import {Entity} from './components/entity/Entity';
import {VideoEdit} from './components/form/videoedit';
import {TextEdit, RichEdit} from './components/form/textedit';
import {ImageEdit, ImageChooser} from './components/form/imageedit';
import {DisplayEntity} from './components/entity/EntityDisplay';
import {EntityForm} from './components/entity/EntityForm';
import {UpdateEntity} from './components/entity/EntityUpdate';
import {ViewFilter} from './components/view/Filter';
import {WebTableView} from './components/view/WebTableView';
import {WebView} from './components/view/WebView';
import {WebListView} from './components/view/WebListView';
import {Image} from './components/ui/Image';
import {Html} from './components/ui/Html';
import {ScrollListener} from './components/ui/ScrollListener';
import {Action} from './components/action/Action';
import {Router} from 'redux-director'
import './styles/App.css';

import {
  Storage,
  Application,
  Window,
  RequestBuilder,
  DataSource,
  Response,
  EntityData,
  Reducers,
  ViewReducer,
  EntityReducer,
  LoginComponent,
  LoginValidator,
  ActionNames,
  formatUrl,
  Colors,
  createStore,
  createAction,
  GroupLoad,
  GurmukhiKeymap,
  Sagas
} from 'laatoocommon';

module.exports = {
    Entity: Entity,
    DisplayEntity: DisplayEntity,
    EntityForm: EntityForm,
    Colors: Colors,
    UpdateEntity: UpdateEntity,
    WebTableView: WebTableView,
    VideoEdit: VideoEdit,
    Action: Action,
    TextEdit: TextEdit,
    RichEdit: RichEdit,
    ScrollListener: ScrollListener,
    WebView: WebView,
    WebListView: WebListView,
    ListView: WebListView,
    Html : Html,
    Image: Image,
    ImageChooser: ImageChooser,
    ImageEdit: ImageEdit,
    redirect: Router.redirect,
    ViewFilter: ViewFilter,
    Storage: Storage,
    LoginValidator: LoginValidator,
    Application: Application,
    GurmukhiKeymap: GurmukhiKeymap,
    Window: Window,
    RequestBuilder: RequestBuilder,
    DataSource:DataSource,
    Response: Response,
    EntityData: EntityData,
    Reducers: Reducers,
    ViewReducer: ViewReducer,
    EntityReducer: EntityReducer,
    LoginComponent: LoginComponent,
    ActionNames: ActionNames,
    formatUrl: formatUrl,
    createStore: createStore,
    createAction: createAction,
    GroupLoad: GroupLoad,
    Sagas: Sagas
}

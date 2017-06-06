
//import { NativeModules } from 'react-native';
import {TextField} from './components/ui/text';
import {Image} from './components/ui/image';
import {ListView} from './components/ui/listview';
import {
  Storage,
  Application,
  LoginValidator,
  Window,
  RequestBuilder,
  DataSource,
  Response,
  EntityData,
  Reducers,
  ViewReducer,
  EntityReducer,
  LoginComponent,
  ActionNames,
  formatUrl,
  createStore,
  createAction,
  Colors,
  GroupLoad,
  GurmukhiKeymap,
  Sagas
} from 'laatoocommon';

import {Dimensions} from 'react-native'

Application.ScreenDimensions = Dimensions.get('window');
Application.MaxHeight = Math.ceil(9*Application.ScreenDimensions.width/16)
console.log("application dimensions ", Application)
//const { RNLaatoonative } = NativeModules;

module.exports = {
  //RNLaatoonative as RNLaatoonative,
  Image: Image,
  TextField : TextField,
  Colors: Colors,
  ListView : ListView,
  Storage : Storage,
  LoginValidator: LoginValidator,
  Application : Application,
  GurmukhiKeymap : GurmukhiKeymap,
  Window : Window,
  RequestBuilder : RequestBuilder,
  DataSource : DataSource,
  Response : Response,
  EntityData : EntityData,
  Reducers : Reducers,
  ViewReducer : ViewReducer,
  EntityReducer : EntityReducer,
  LoginComponent : LoginComponent,
  ActionNames : ActionNames,
  formatUrl : formatUrl,
  createStore : createStore,
  createAction : createAction,
  GroupLoad : GroupLoad,
  Sagas : Sagas
};

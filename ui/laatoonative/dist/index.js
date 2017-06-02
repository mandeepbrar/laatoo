'use strict';

var _text=require('./components/ui/text');
var _image=require('./components/ui/image');
var _listview=require('./components/ui/listview');
var _laatoocommon=require('laatoocommon');





















var _reactNative=require('react-native');

_laatoocommon.Application.ScreenDimensions=_reactNative.Dimensions.get('window');
_laatoocommon.Application.MaxHeight=Math.ceil(9*_laatoocommon.Application.ScreenDimensions.width/16);
console.log("application dimensions ",_laatoocommon.Application);


module.exports={

Image:_image.Image,
TextField:_text.TextField,
Colors:_laatoocommon.Colors,
ListView:_listview.ListView,
Storage:_laatoocommon.Storage,
Application:_laatoocommon.Application,
GurmukhiKeymap:_laatoocommon.GurmukhiKeymap,
Window:_laatoocommon.Window,
RequestBuilder:_laatoocommon.RequestBuilder,
DataSource:_laatoocommon.DataSource,
Response:_laatoocommon.Response,
EntityData:_laatoocommon.EntityData,
Reducers:_laatoocommon.Reducers,
ViewReducer:_laatoocommon.ViewReducer,
EntityReducer:_laatoocommon.EntityReducer,
LoginComponent:_laatoocommon.LoginComponent,
ActionNames:_laatoocommon.ActionNames,
formatUrl:_laatoocommon.formatUrl,
createStore:_laatoocommon.createStore,
createAction:_laatoocommon.createAction,
GroupLoad:_laatoocommon.GroupLoad,
Sagas:_laatoocommon.Sagas};
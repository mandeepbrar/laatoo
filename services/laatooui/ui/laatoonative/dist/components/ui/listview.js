'use strict';Object.defineProperty(exports,"__esModule",{value:true});exports.ListView=undefined;var _extends=Object.assign||function(target){for(var i=1;i<arguments.length;i++){var source=arguments[i];for(var key in source){if(Object.prototype.hasOwnProperty.call(source,key)){target[key]=source[key];}}}return target;};var _createClass=function(){function defineProperties(target,props){for(var i=0;i<props.length;i++){var descriptor=props[i];descriptor.enumerable=descriptor.enumerable||false;descriptor.configurable=true;if("value"in descriptor)descriptor.writable=true;Object.defineProperty(target,descriptor.key,descriptor);}}return function(Constructor,protoProps,staticProps){if(protoProps)defineProperties(Constructor.prototype,protoProps);if(staticProps)defineProperties(Constructor,staticProps);return Constructor;};}();

var _react=require('react');var _react2=_interopRequireDefault(_react);
var _reactNative=require('react-native');
var _laatoocommon=require('laatoocommon');function _interopRequireDefault(obj){return obj&&obj.__esModule?obj:{default:obj};}function _classCallCheck(instance,Constructor){if(!(instance instanceof Constructor)){throw new TypeError("Cannot call a class as a function");}}function _possibleConstructorReturn(self,call){if(!self){throw new ReferenceError("this hasn't been initialised - super() hasn't been called");}return call&&(typeof call==="object"||typeof call==="function")?call:self;}function _inherits(subClass,superClass){if(typeof superClass!=="function"&&superClass!==null){throw new TypeError("Super expression must either be null or a function, not "+typeof superClass);}subClass.prototype=Object.create(superClass&&superClass.prototype,{constructor:{value:subClass,enumerable:false,writable:true,configurable:true}});if(superClass)Object.setPrototypeOf?Object.setPrototypeOf(subClass,superClass):subClass.__proto__=superClass;}var

ListView=function(_React$Component){_inherits(ListView,_React$Component);
function ListView(props){_classCallCheck(this,ListView);var _this=_possibleConstructorReturn(this,(ListView.__proto__||Object.getPrototypeOf(ListView)).call(this,
props));
_this.getView=_this.getView.bind(_this);
_this.renderView=_this.renderView.bind(_this);
_this.renderItem=_this.renderItem.bind(_this);
_this.getItem=_this.getItem.bind(_this);
_this.getHeader=_this.getHeader.bind(_this);
_this.getPagination=_this.getPagination.bind(_this);
_this.getFilter=_this.getFilter.bind(_this);
_this.onScrollEnd=_this.onScrollEnd.bind(_this);
_this.addMethod=_this.addMethod.bind(_this);
_this.style=_extends({},{flex:1},_this.props.style);
_this.contentStyle=_extends({},{flex:1},_this.props.contentStyle);
_this.numItems=0;return _this;
}_createClass(ListView,[{key:'addMethod',value:function addMethod(

name,method){
return this.viewdata.addMethod(name,method);
}},{key:'methods',value:function methods()
{
return this.viewdata.methods;
}},{key:'onScrollEnd',value:function onScrollEnd()

{
var methods=this.methods();
methods.loadMore();
}},{key:'getView',value:function getView(

header,groups,pagination,filter){
console.log("getting the view ",this.style,this.contentStyle);
if(this.props.getView){
return this.props.getView(this,header,groups,pagination,filter);
}
return(
_react2.default.createElement(_reactNative.View,{style:this.style},
_react2.default.createElement(_reactNative.View,{style:this.props.headerStyle},
header),

filter?filter:null,
_react2.default.createElement(_reactNative.View,{style:this.contentStyle},
groups),

pagination?pagination:null));


}},{key:'getFilter',value:function getFilter(
filterTitle,filterForm,filterGo){
if(this.props.getFilter){
return this.props.getFilter(this,filterTitle,filterForm,filterGo);
}
return null;
}},{key:'renderItem',value:function renderItem(
itemInfo){
console.log("rendering item*****",itemInfo);
return this.getItem(itemInfo.item,itemInfo.index);
}},{key:'getItem',value:function getItem(
x,i){
if(this.props.getItem){
return this.props.getItem(this,x,i);
}
var child=_react2.default.Children.only(this.props.children);
console.log("child",child);
return _react2.default.cloneElement(child,{item:x,index:i});
}},{key:'getHeader',value:function getHeader()
{
if(this.props.getHeader){
return this.props.getHeader(this);
}
return null;
}},{key:'getPagination',value:function getPagination()
{
if(this.props.paginate&&this.props.getPagination){
var pages=this.props.totalPages;
var page=this.props.currentPage;
return this.props.getPagination(this,pages,page);
}
return null;
}},{key:'renderView',value:function renderView(
viewdata,items,currentPage,totalPages){
this.viewdata=viewdata;
var body=[];
if(items){
if(this.props.incrementalLoad){
body.push(_react2.default.createElement(_reactNative.FlatList,{data:items,onEndReached:this.onScrollEnd,renderItem:this.renderItem,horizontal:this.props.horizontal,numColumns:this.props.numColumns,columnWrapperStyle:this.props.columnWrapperStyle,onPress:this.props.onPress}));
}else{
body.push(_react2.default.createElement(_reactNative.FlatList,{data:items,horizontal:this.props.horizontal,numColumns:this.props.numColumns,renderItem:this.renderItem}));
}
}else{
if(this.props.loader){
body.push(this.props.loader);
}
}
var header=this.getHeader();
var filterCtrl=this.getFilter(this.props.filterTitle,this.props.filterForm,this.props.filterGo,this.filter);
var pagination=this.getPagination();
return this.getView(header,body,pagination,filterCtrl);
}},{key:'render',value:function render()

{
return(
_react2.default.createElement(_laatoocommon.ViewData,{
getView:this.renderView,
key:this.props.key,
reducer:this.props.reducer,
paginate:this.props.paginate,
pageSize:this.props.pageSize,
viewService:this.props.viewService,
urlParams:this.props.urlParams,
postArgs:this.props.postArgs,
defaultFilter:this.props.defaultFilter,
currentPage:this.props.currentPage,
style:this.props.style,
className:this.props.className,
incrementalLoad:this.props.incrementalLoad,
globalReducer:this.props.globalReducer}));

}}]);return ListView;}(_react2.default.Component);exports.



ListView=ListView;
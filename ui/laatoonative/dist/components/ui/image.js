'use strict';Object.defineProperty(exports,"__esModule",{value:true});exports.Image=undefined;var _extends=Object.assign||function(target){for(var i=1;i<arguments.length;i++){var source=arguments[i];for(var key in source){if(Object.prototype.hasOwnProperty.call(source,key)){target[key]=source[key];}}}return target;};var _createClass=function(){function defineProperties(target,props){for(var i=0;i<props.length;i++){var descriptor=props[i];descriptor.enumerable=descriptor.enumerable||false;descriptor.configurable=true;if("value"in descriptor)descriptor.writable=true;Object.defineProperty(target,descriptor.key,descriptor);}}return function(Constructor,protoProps,staticProps){if(protoProps)defineProperties(Constructor.prototype,protoProps);if(staticProps)defineProperties(Constructor,staticProps);return Constructor;};}();var _react=require('react');var _react2=_interopRequireDefault(_react);
var _laatoocommon=require('laatoocommon');
var _reactNative=require('react-native');function _interopRequireDefault(obj){return obj&&obj.__esModule?obj:{default:obj};}function _classCallCheck(instance,Constructor){if(!(instance instanceof Constructor)){throw new TypeError("Cannot call a class as a function");}}function _possibleConstructorReturn(self,call){if(!self){throw new ReferenceError("this hasn't been initialised - super() hasn't been called");}return call&&(typeof call==="object"||typeof call==="function")?call:self;}function _inherits(subClass,superClass){if(typeof superClass!=="function"&&superClass!==null){throw new TypeError("Super expression must either be null or a function, not "+typeof superClass);}subClass.prototype=Object.create(superClass&&superClass.prototype,{constructor:{value:subClass,enumerable:false,writable:true,configurable:true}});if(superClass)Object.setPrototypeOf?Object.setPrototypeOf(subClass,superClass):subClass.__proto__=superClass;}var


PrefixImage=function(_React$Component){_inherits(PrefixImage,_React$Component);
function PrefixImage(props){_classCallCheck(this,PrefixImage);return _possibleConstructorReturn(this,(PrefixImage.__proto__||Object.getPrototypeOf(PrefixImage)).call(this,
props));
}_createClass(PrefixImage,[{key:'render',value:function render()
{
var skipPrefix=this.props.skipPrefix?this.props.skipPrefix:false;
var source=this.props.src;
if(!source||source.length==0){
if(this.props.children){
return this.props.children;
}
return null;
}
if(!skipPrefix&&!this.props.src.startsWith("http")){
source=this.props.prefix+source;
}
console.log("source file ",source,this.props);
var srcobj=null;
if(this.props.local){
srcobj=require(source);
}else{
srcobj={uri:source};
}
console.log("src obj--",srcobj);

var style=_extends({},{flex:1,resizeMode:'cover',height:_laatoocommon.Application.MaxHeight,width:null},this.props.style);
console.log("style of the image*************",style);
var i=_react2.default.createElement(_reactNative.Image,_extends({style:style},this.props.modifier,{source:srcobj}));
if(this.props.link){
return i;



}else{
return i;
}
}}]);return PrefixImage;}(_react2.default.Component);exports.



Image=PrefixImage;
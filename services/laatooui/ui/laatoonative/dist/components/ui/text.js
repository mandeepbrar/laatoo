'use strict';Object.defineProperty(exports,"__esModule",{value:true});exports.TextField=undefined;var _createClass=function(){function defineProperties(target,props){for(var i=0;i<props.length;i++){var descriptor=props[i];descriptor.enumerable=descriptor.enumerable||false;descriptor.configurable=true;if("value"in descriptor)descriptor.writable=true;Object.defineProperty(target,descriptor.key,descriptor);}}return function(Constructor,protoProps,staticProps){if(protoProps)defineProperties(Constructor.prototype,protoProps);if(staticProps)defineProperties(Constructor,staticProps);return Constructor;};}();var _react=require('react');var _react2=_interopRequireDefault(_react);
var _nativeBase=require('native-base');
var _laatoocommon=require('laatoocommon');function _interopRequireDefault(obj){return obj&&obj.__esModule?obj:{default:obj};}function _classCallCheck(instance,Constructor){if(!(instance instanceof Constructor)){throw new TypeError("Cannot call a class as a function");}}function _possibleConstructorReturn(self,call){if(!self){throw new ReferenceError("this hasn't been initialised - super() hasn't been called");}return call&&(typeof call==="object"||typeof call==="function")?call:self;}function _inherits(subClass,superClass){if(typeof superClass!=="function"&&superClass!==null){throw new TypeError("Super expression must either be null or a function, not "+typeof superClass);}subClass.prototype=Object.create(superClass&&superClass.prototype,{constructor:{value:subClass,enumerable:false,writable:true,configurable:true}});if(superClass)Object.setPrototypeOf?Object.setPrototypeOf(subClass,superClass):subClass.__proto__=superClass;}var

Text=function(_React$Component){_inherits(Text,_React$Component);
function Text(props){_classCallCheck(this,Text);var _this=_possibleConstructorReturn(this,(Text.__proto__||Object.getPrototypeOf(Text)).call(this,
props));

_this.value=props.value?props.value:"";
_this.state={value:_this.value};
_this.setValue=_this.setValue.bind(_this);
_this.change=_this.change.bind(_this);return _this;
}_createClass(Text,[{key:'componentWillReceiveProps',value:function componentWillReceiveProps(
nextprops){
if(nextprops.value){
this.setValue(nextprops.value);
}
}},{key:'setLanguage',value:function setLanguage(
lang){
this.language=lang;
}},{key:'setValue',value:function setValue(
value,successFunc){
this.setState({value:value},successFunc);
this.value=value;
if(this.props.onChange){
var evt={target:{name:this.props.name,value:value}};
this.props.onChange(evt);
}
}},{key:'change',value:function change(







val){
this.setValue(val);
}},{key:'render',value:function render()
{
var config={};
if(this.props.config){
config=this.props.config;
}
return(
_react2.default.createElement(_nativeBase.Item,{fixedLabel:this.props.fixedLabel,stackedLabel:this.props.stackedLabel,floatingLabel:this.props.floatingLabel,inlineLabel:this.props.inlineLabel},
_react2.default.createElement(_nativeBase.Label,null,this.props.label),
this.props.lefticon,
_react2.default.createElement(_nativeBase.Input,{name:this.props.name,className:this.props.className,onKeyPress:this.keyPress,value:this.state.value,
defaultValue:this.props.defaultValue,rows:this.props.rows,rowsMax:this.props.rows,multiLine:this.props.multiline,
placeholder:this.props.placeholder,secureTextEntry:this.props.secureTextEntry,rounded:this.props.rouded,
textColor:this.props.textColor,inputColorPlaceholder:this.props.inputColorPlaceholder,inputHeightBase:this.props.inputHeightBase,
style:config.style,inputStyle:config.inputStyle,onChangeText:this.change}),
this.props.icon));


}}]);return Text;}(_react2.default.Component);exports.



TextField=Text;
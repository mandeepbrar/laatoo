'use strict';Object.defineProperty(exports,"__esModule",{value:true});exports.Action=undefined;var _createClass=function(){function defineProperties(target,props){for(var i=0;i<props.length;i++){var descriptor=props[i];descriptor.enumerable=descriptor.enumerable||false;descriptor.configurable=true;if("value"in descriptor)descriptor.writable=true;Object.defineProperty(target,descriptor.key,descriptor);}}return function(Constructor,protoProps,staticProps){if(protoProps)defineProperties(Constructor.prototype,protoProps);if(staticProps)defineProperties(Constructor,staticProps);return Constructor;};}();

var _react=require('react');var _react2=_interopRequireDefault(_react);
var _actionbutton=require('./actionbutton');var _actionbutton2=_interopRequireDefault(_actionbutton);
var _reactRedux=require('react-redux');
var _actionlink=require('./actionlink');var _actionlink2=_interopRequireDefault(_actionlink);
var _laatoocommon=require('laatoocommon');
var _reactNavigation=require('react-navigation');function _interopRequireDefault(obj){return obj&&obj.__esModule?obj:{default:obj};}function _classCallCheck(instance,Constructor){if(!(instance instanceof Constructor)){throw new TypeError("Cannot call a class as a function");}}function _possibleConstructorReturn(self,call){if(!self){throw new ReferenceError("this hasn't been initialised - super() hasn't been called");}return call&&(typeof call==="object"||typeof call==="function")?call:self;}function _inherits(subClass,superClass){if(typeof superClass!=="function"&&superClass!==null){throw new TypeError("Super expression must either be null or a function, not "+typeof superClass);}subClass.prototype=Object.create(superClass&&superClass.prototype,{constructor:{value:subClass,enumerable:false,writable:true,configurable:true}});if(superClass)Object.setPrototypeOf?Object.setPrototypeOf(subClass,superClass):subClass.__proto__=superClass;}var

ActionComp=function(_React$Component){_inherits(ActionComp,_React$Component);
function ActionComp(props){_classCallCheck(this,ActionComp);var _this=_possibleConstructorReturn(this,(ActionComp.__proto__||Object.getPrototypeOf(ActionComp)).call(this,
props));
_this.renderView=_this.renderView.bind(_this);
_this.dispatchAction=_this.dispatchAction.bind(_this);
_this.actionFunc=_this.actionFunc.bind(_this);
_this.hasPermission=false;
console.log("Action",props);
var action=_laatoocommon.Application.Actions[props.name];
if(action){
_this.action=action;
_this.hasPermission=(0,_laatoocommon.hasPermission)(action.permission);
}return _this;
}_createClass(ActionComp,[{key:'dispatchAction',value:function dispatchAction()

{
var payload={};
if(this.props.params){
payload=this.props.params;
}
this.props.dispatch((0,_laatoocommon.createAction)(this.action.action,payload,{successCallback:this.props.successCallback,failureCallback:this.props.failureCallback}));
}},{key:'actionFunc',value:function actionFunc()
{
if(this.props.confirm){
if(!this.props.confirm(this.props)){
return false;
}
}
console.log("exceuting action",this.action);
switch(this.action.actiontype){
case"dispatchaction":
this.dispatchAction();
return false;
case"method":
var params=this.props.params;
var method=this.props.method;
method(params);
return false;








default:
if(this.action.target){

console.log(this.action.target);

this.props.dispatch(_reactNavigation.NavigationActions.navigate({routeName:this.action.target,params:this.props.params}));


}
return false;}

}},{key:'renderView',value:function renderView()

{
if(!this.hasPermission){
return null;
}
var actionF=this.actionFunc;
console.log("action props",this.props);
switch(this.props.widget){
case'button':{
return(
_react2.default.createElement(_actionbutton2.default,{style:this.props.style,actionFunc:actionF,key:this.props.name+"_comp",transparent:this.props.transparent,
iconRight:this.props.iconRight,iconLeft:this.props.iconLeft,actionchildren:this.props.children}));


}
default:{
return(
_react2.default.createElement(_actionlink2.default,{style:this.props.style,actionFunc:actionF,key:this.props.name+"_comp",actionchildren:this.props.children}));


}}

}},{key:'render',value:function render()
{
return this.renderView();
}}]);return ActionComp;}(_react2.default.Component);


ActionComp.propTypes={
name:_react2.default.PropTypes.string.isRequired};



var Action=(0,_reactRedux.connect)()(ActionComp);exports.
Action=Action;
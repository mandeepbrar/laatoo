'use strict';Object.defineProperty(exports,"__esModule",{value:true});

var _react=require('react');var _react2=_interopRequireDefault(_react);
var _nativeBase=require('native-base');function _interopRequireDefault(obj){return obj&&obj.__esModule?obj:{default:obj};}

var ActionLink=function ActionLink(props){return(
_react2.default.createElement(_nativeBase.Button,{style:props.style,transparent:true,iconLeft:true,onPress:props.actionFunc},
props.actionchildren));};exports.default=



ActionLink;
define("oauthui",["reactwebcommon","react","react-redux","prop-types"],(function(t,n,e,o){return function(t){var n={};function e(o){if(n[o])return n[o].exports;var r=n[o]={i:o,l:!1,exports:{}};return t[o].call(r.exports,r,r.exports,e),r.l=!0,r.exports}return e.m=t,e.c=n,e.d=function(t,n,o){e.o(t,n)||Object.defineProperty(t,n,{enumerable:!0,get:o})},e.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},e.t=function(t,n){if(1&n&&(t=e(t)),8&n)return t;if(4&n&&"object"==typeof t&&t&&t.__esModule)return t;var o=Object.create(null);if(e.r(o),Object.defineProperty(o,"default",{enumerable:!0,value:t}),2&n&&"string"!=typeof t)for(var r in t)e.d(o,r,function(n){return t[n]}.bind(null,r));return o},e.n=function(t){var n=t&&t.__esModule?function(){return t.default}:function(){return t};return e.d(n,"a",n),n},e.o=function(t,n){return Object.prototype.hasOwnProperty.call(t,n)},e.p="/",e(e.s=4)}([function(n,e){n.exports=t},function(t,e){t.exports=n},function(t,n){t.exports=e},function(t,n){t.exports=o},function(t,n,e){"use strict";e.r(n);var o=e(1),r=e.n(o),i=e(0),u=e(2);e.d(n,"Initialize",(function(){return c})),e.d(n,"OauthButton",(function(){return l})),e.d(n,"oauthLogin",(function(){return a}));e(3);function c(t,n,e,o,r,i){}function a(t){window.open(t.oauthurl+t.realm,"_blank","height=500,width=400,toolbar=no,resizable=yes,menubar=no,location=0");var n,e,o=(n=t.oauthurl,(e=document.createElement("a")).href=n,e);loginSite=o.protocol+"//"+o.hostname+(o.port?":"+o.port:"")}var l=Object(u.connect)((function(t,n){return{realm:Application.Security.realm}}),(function(t,n){return console.log("map dispatch of oauth login compoent"),{handleOauthLogin:function(n){t(createAction((void 0).LOGIN_SUCCESS,{userId:n.id,token:n.token,permissions:n.permissions}))}}}))((function(t){return window.addEventListener("message",(function(n){n.origin===loginSite&&"LoginSuccess"==n.data.message&&t.handleOauthLogin(n.data)})),console.log("returning Action from oauth btn",i.Action),r.a.createElement(i.Action,{widget:"button",method:function(){a(t)},action:{actiontype:"method"},className:"oauthlogin "+_tn(t.className,"")},t.children)}))}])}));
//# sourceMappingURL=index.js.map
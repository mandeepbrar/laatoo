define("laatoositeui",["react","oauthui","uicommon"],(function(t,e,n){return function(t){var e={};function n(o){if(e[o])return e[o].exports;var u=e[o]={i:o,l:!1,exports:{}};return t[o].call(u.exports,u,u.exports,n),u.l=!0,u.exports}return n.m=t,n.c=e,n.d=function(t,e,o){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:o})},n.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var o=Object.create(null);if(n.r(o),Object.defineProperty(o,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var u in t)n.d(o,u,function(e){return t[e]}.bind(null,u));return o},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=3)}([function(e,n){e.exports=t},function(t,n){t.exports=e},function(t,e){t.exports=n},function(t,e,n){"use strict";n.r(e),n.d(e,"Initialize",(function(){return i}));var o=n(0),u=n.n(o),r=(n(4),n(2),n(1));function i(t,e,n,o,u,r){console.log("Initializing ui"),_r("Methods","demoActionButtons",c)}function c(t,e,n,o,i){return console.log("oauth button in demo actions",r.OauthButton),u.a.createElement(_uikit.Block,null,u.a.createElement(r.OauthButton,{className:"googleAuthAction s10 btn-google"},u.a.createElement("i",{className:"s10 fa fa-google"}),"Google"),u.a.createElement(_uikit.ActionButton,{onClick:e(),className:"submitBtn s10"},"Sign up"))}console.log("oauth ui ",r.OauthButton)},function(t,e){}])}));
//# sourceMappingURL=index.js.map
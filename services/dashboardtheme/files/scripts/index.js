define("dashboardtheme",["authui","reactwebcommon","reactpages","react-redux","react"],function(t,e,n,r,o){return function(t){function e(r){if(n[r])return n[r].exports;var o=n[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,e),o.l=!0,o.exports}var n={};return e.m=t,e.c=n,e.d=function(t,n,r){e.o(t,n)||Object.defineProperty(t,n,{configurable:!1,enumerable:!0,get:r})},e.n=function(t){var n=t&&t.__esModule?function(){return t.default}:function(){return t};return e.d(n,"a",n),n},e.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},e.p="/",e(e.s=25)}([function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,e,n){var r=n(20)("wks"),o=n(21),i=n(4).Symbol;t.exports=function(t){return r[t]||(r[t]=i&&i[t]||(i||o)("Symbol."+t))}},function(t,e,n){var r=n(4),o=n(1),i=n(14),u=function(t,e,n){var c,a,s,f=t&u.F,l=t&u.G,p=t&u.S,d=t&u.P,m=t&u.B,y=t&u.W,h=l?o:o[e]||(o[e]={}),v=l?r:p?r[e]:(r[e]||{}).prototype;l&&(n=e);for(c in n)(a=!f&&v&&c in v)&&c in h||(s=a?v[c]:n[c],h[c]=l&&"function"!=typeof v[c]?n[c]:m&&a?i(s,r):y&&v[c]==s?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(s):d&&"function"==typeof s?i(Function.call,s):s,d&&((h.prototype||(h.prototype={}))[c]=s))};u.F=1,u.G=2,u.S=4,u.P=8,u.B=16,u.W=32,t.exports=u},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e,n){var r=n(48),o=n(7);t.exports=function(t){return r(o(t))}},function(t,e){t.exports=o},function(t,e){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e,n){var r=n(0),o=n(10);t.exports=n(19)?function(t,e,n){return r.setDesc(t,e,o(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports={}},function(t,e,n){var r=n(0).setDesc,o=n(11),i=n(2)("toStringTag");t.exports=function(t,e,n){t&&!o(t=n?t:t.prototype,i)&&r(t,i,{configurable:!0,value:e})}},function(t,e,n){var r=n(32);t.exports=function(t,e,n){if(r(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,r){return t.call(e,n,r)};case 3:return function(n,r,o){return t.call(e,n,r,o)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){"use strict";function r(t){return t&&t.__esModule?t:{default:t}}e.__esModule=!0;var o=n(38),i=r(o),u=n(49),c=r(u),a="function"==typeof c.default&&"symbol"==typeof i.default?function(t){return typeof t}:function(t){return t&&"function"==typeof c.default&&t.constructor===c.default&&t!==c.default.prototype?"symbol":typeof t};e.default="function"==typeof c.default&&"symbol"===a(i.default)?function(t){return void 0===t?"undefined":a(t)}:function(t){return t&&"function"==typeof c.default&&t.constructor===c.default&&t!==c.default.prototype?"symbol":void 0===t?"undefined":a(t)}},function(t,e,n){"use strict";var r=n(17),o=n(3),i=n(18),u=n(9),c=n(11),a=n(12),s=n(43),f=n(13),l=n(0).getProto,p=n(2)("iterator"),d=!([].keys&&"next"in[].keys()),m=function(){return this};t.exports=function(t,e,n,y,h,v,g){s(n,e,y);var _,b,x=function(t){if(!d&&t in S)return S[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},k=e+" Iterator",w="values"==h,O=!1,S=t.prototype,P=S[p]||S["@@iterator"]||h&&S[h],E=P||x(h);if(P){var N=l(E.call(new t));f(N,k,!0),!r&&c(S,"@@iterator")&&u(N,p,m),w&&"values"!==P.name&&(O=!0,E=function(){return P.call(this)})}if(r&&!g||!d&&!O&&S[p]||u(S,p,E),a[e]=E,a[k]=m,h)if(_={values:w?E:x("values"),keys:v?E:x("keys"),entries:w?x("entries"):E},g)for(b in _)b in S||i(S,b,_[b]);else o(o.P+o.F*(d||O),e,_);return _}},function(t,e){t.exports=!0},function(t,e,n){t.exports=n(9)},function(t,e,n){t.exports=!n(8)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){var r=n(4),o=r["__core-js_shared__"]||(r["__core-js_shared__"]={});t.exports=function(t){return o[t]||(o[t]={})}},function(t,e){var n=0,r=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+r).toString(36))}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){var r=n(24);t.exports=function(t){if(!r(t))throw TypeError(t+" is not an object!");return t}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){"use strict";function r(e,n,r,i,u,a){if(console.log("appname",e,"ins ",n,"mod",r,"settings",i),t.properties=Application.Properties[n],t.settings=i,t.skipAuth=i.skipAuth,!_reg("Pages","home")){var f={id:"home",route:"/",components:{main:{type:"component",component:c.a.createElement(s.a,{modProps:t.properties})}}};Application.Register("Pages","home",f),Application.Register("Actions","Page_home",{url:"/"})}var l="authui",p="WebLoginForm";i&&!i.skipAuth&&i.loginModule&&(l=i.loginModule,p=i.loginComp);var d=a(l);d&&(t.logInComp=d[p]),o()}function o(){var e=[],n=[];n=t.settings&&t.settings.menu?t.settings.menu:Application.Properties.menu,console.log("dashboard menu",n),n?n.forEach(function(t){t.page?e.push({title:t.title,action:"Page_"+t.page}):e.push({title:t.title,action:t.action})}):e.push({title:"Home",action:"Page_home"}),t.menu=e}function i(e,n){return e.components.menu=function(e){return c.a.createElement(n.Navbar,{items:t.menu,vertical:t.settings.vertical})},e}Object.defineProperty(e,"__esModule",{value:!0}),n.d(e,"Initialize",function(){return r}),n.d(e,"ProcessRoute",function(){return i}),n.d(e,"Theme",function(){return d});var u=n(6),c=n.n(u),a=n(26),s=n(67),f=n(68),l=(n.n(f),n(69)),t=(n.n(l),this),p=function(e){var n=!!t.settings.vertical,r=(t.logInComp,c.a.createElement(e.uikit.Block,{className:"body"},c.a.createElement(e.uikit.Block,{className:n?"vertmenu":"horizmenu"},c.a.createElement(e.router.View,{name:"menu"})),c.a.createElement(e.uikit.Block,{className:n?"vertbody":"horizbody"},c.a.createElement(e.router.View,{name:"main"}))));return e.loggedIn||t.skipAuth?r:c.a.createElement(e.uikit.Block,{className:"dashlogin"},c.a.createElement(t.logInComp,null))},d=function(e){return c.a.createElement(e.uikit.UIWrapper,null,c.a.createElement(e.uikit.Block,{className:t.settings.className?t.settings.className+" dashboard":"dashboard"},c.a.createElement(a.a,{uikit:e.uikit,module:t}),c.a.createElement(l.LoginValidator,{validateService:"validate"},c.a.createElement(p,{uikit:e.uikit,router:e.router})),c.a.createElement(e.uikit.Block,{className:"footer"},e.children)))}},function(t,e,n){"use strict";var r=n(27),o=n.n(r),i=n(33),u=n.n(i),c=n(34),a=n.n(c),s=n(37),f=n.n(s),l=n(57),p=n.n(l),d=n(6),m=n.n(d),y=n(64),h=(n.n(y),n(65)),v=(n.n(h),n(66)),g=(n.n(v),function(t){function e(t){return u()(this,e),f()(this,(e.__proto__||o()(e)).call(this,t))}return p()(e,t),a()(e,[{key:"render",value:function(){var t=this.props,e=t.module,n=e.properties.header,r=e.settings,o=r.infoBlock?r.infoBlock:"userBlock";return console.log("header properties",n,t),m.a.createElement(t.uikit.Block,{className:n.className?n.className:"header"},m.a.createElement(t.uikit.Block,{className:"logo"},n.image?m.a.createElement(t.uikit.Block,{className:"image"},m.a.createElement(y.Image,{src:n.image})):null,n.title?m.a.createElement(t.uikit.Block,{className:"title"},n.title):null),t.loggedIn?m.a.createElement(h.Panel,{id:o,className:"infoBlock"}):null)}}]),e}(m.a.Component)),_=function(t,e){return{loggedIn:"LoggedIn"==t.Security.status}},b=Object(v.connect)(_)(g);e.a=b},function(t,e,n){t.exports={default:n(28),__esModule:!0}},function(t,e,n){n(29),t.exports=n(1).Object.getPrototypeOf},function(t,e,n){var r=n(30);n(31)("getPrototypeOf",function(t){return function(e){return t(r(e))}})},function(t,e,n){var r=n(7);t.exports=function(t){return Object(r(t))}},function(t,e,n){var r=n(3),o=n(1),i=n(8);t.exports=function(t,e){var n=(o.Object||{})[t]||Object[t],u={};u[t]=e(n),r(r.S+r.F*i(function(){n(1)}),"Object",u)}},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){"use strict";e.__esModule=!0;var r=n(35),o=function(t){return t&&t.__esModule?t:{default:t}}(r);e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var r=e[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),(0,o.default)(t,r.key,r)}}return function(e,n,r){return n&&t(e.prototype,n),r&&t(e,r),e}}()},function(t,e,n){t.exports={default:n(36),__esModule:!0}},function(t,e,n){var r=n(0);t.exports=function(t,e,n){return r.setDesc(t,e,n)}},function(t,e,n){"use strict";e.__esModule=!0;var r=n(15),o=function(t){return t&&t.__esModule?t:{default:t}}(r);e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,o.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){t.exports={default:n(39),__esModule:!0}},function(t,e,n){n(40),n(44),t.exports=n(2)("iterator")},function(t,e,n){"use strict";var r=n(41)(!0);n(16)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=r(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){var r=n(42),o=n(7);t.exports=function(t){return function(e,n){var i,u,c=String(o(e)),a=r(n),s=c.length;return a<0||a>=s?t?"":void 0:(i=c.charCodeAt(a),i<55296||i>56319||a+1===s||(u=c.charCodeAt(a+1))<56320||u>57343?t?c.charAt(a):i:t?c.slice(a,a+2):u-56320+(i-55296<<10)+65536)}}},function(t,e){var n=Math.ceil,r=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?r:n)(t)}},function(t,e,n){"use strict";var r=n(0),o=n(10),i=n(13),u={};n(9)(u,n(2)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=r.create(u,{next:o(1,n)}),i(t,e+" Iterator")}},function(t,e,n){n(45);var r=n(12);r.NodeList=r.HTMLCollection=r.Array},function(t,e,n){"use strict";var r=n(46),o=n(47),i=n(12),u=n(5);t.exports=n(16)(Array,"Array",function(t,e){this._t=u(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,o(1)):"keys"==e?o(0,n):"values"==e?o(0,t[n]):o(0,[n,t[n]])},"values"),i.Arguments=i.Array,r("keys"),r("values"),r("entries")},function(t,e){t.exports=function(){}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e,n){var r=n(22);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==r(t)?t.split(""):Object(t)}},function(t,e,n){t.exports={default:n(50),__esModule:!0}},function(t,e,n){n(51),n(56),t.exports=n(1).Symbol},function(t,e,n){"use strict";var r=n(0),o=n(4),i=n(11),u=n(19),c=n(3),a=n(18),s=n(8),f=n(20),l=n(13),p=n(21),d=n(2),m=n(52),y=n(53),h=n(54),v=n(55),g=n(23),_=n(5),b=n(10),x=r.getDesc,k=r.setDesc,w=r.create,O=y.get,S=o.Symbol,P=o.JSON,E=P&&P.stringify,N=!1,j=d("_hidden"),M=r.isEnum,A=f("symbol-registry"),B=f("symbols"),I="function"==typeof S,D=Object.prototype,C=u&&s(function(){return 7!=w(k({},"a",{get:function(){return k(this,"a",{value:7}).a}})).a})?function(t,e,n){var r=x(D,e);r&&delete D[e],k(t,e,n),r&&t!==D&&k(D,e,r)}:k,F=function(t){var e=B[t]=w(S.prototype);return e._k=t,u&&N&&C(D,t,{configurable:!0,set:function(e){i(this,j)&&i(this[j],t)&&(this[j][t]=!1),C(this,t,b(1,e))}}),e},T=function(t){return"symbol"==typeof t},W=function(t,e,n){return n&&i(B,e)?(n.enumerable?(i(t,j)&&t[j][e]&&(t[j][e]=!1),n=w(n,{enumerable:b(0,!1)})):(i(t,j)||k(t,j,b(1,{})),t[j][e]=!0),C(t,e,n)):k(t,e,n)},L=function(t,e){g(t);for(var n,r=h(e=_(e)),o=0,i=r.length;i>o;)W(t,n=r[o++],e[n]);return t},z=function(t,e){return void 0===e?w(t):L(w(t),e)},J=function(t){var e=M.call(this,t);return!(e||!i(this,t)||!i(B,t)||i(this,j)&&this[j][t])||e},R=function(t,e){var n=x(t=_(t),e);return!n||!i(B,e)||i(t,j)&&t[j][e]||(n.enumerable=!0),n},G=function(t){for(var e,n=O(_(t)),r=[],o=0;n.length>o;)i(B,e=n[o++])||e==j||r.push(e);return r},K=function(t){for(var e,n=O(_(t)),r=[],o=0;n.length>o;)i(B,e=n[o++])&&r.push(B[e]);return r},V=function(t){if(void 0!==t&&!T(t)){for(var e,n,r=[t],o=1,i=arguments;i.length>o;)r.push(i[o++]);return e=r[1],"function"==typeof e&&(n=e),!n&&v(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!T(e))return e}),r[1]=e,E.apply(P,r)}},H=s(function(){var t=S();return"[null]"!=E([t])||"{}"!=E({a:t})||"{}"!=E(Object(t))});I||(S=function(){if(T(this))throw TypeError("Symbol is not a constructor");return F(p(arguments.length>0?arguments[0]:void 0))},a(S.prototype,"toString",function(){return this._k}),T=function(t){return t instanceof S},r.create=z,r.isEnum=J,r.getDesc=R,r.setDesc=W,r.setDescs=L,r.getNames=y.get=G,r.getSymbols=K,u&&!n(17)&&a(D,"propertyIsEnumerable",J,!0));var U={for:function(t){return i(A,t+="")?A[t]:A[t]=S(t)},keyFor:function(t){return m(A,t)},useSetter:function(){N=!0},useSimple:function(){N=!1}};r.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=d(t);U[t]=I?e:F(e)}),N=!0,c(c.G+c.W,{Symbol:S}),c(c.S,"Symbol",U),c(c.S+c.F*!I,"Object",{create:z,defineProperty:W,defineProperties:L,getOwnPropertyDescriptor:R,getOwnPropertyNames:G,getOwnPropertySymbols:K}),P&&c(c.S+c.F*(!I||H),"JSON",{stringify:V}),l(S,"Symbol"),l(Math,"Math",!0),l(o.JSON,"JSON",!0)},function(t,e,n){var r=n(0),o=n(5);t.exports=function(t,e){for(var n,i=o(t),u=r.getKeys(i),c=u.length,a=0;c>a;)if(i[n=u[a++]]===e)return n}},function(t,e,n){var r=n(5),o=n(0).getNames,i={}.toString,u="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[],c=function(t){try{return o(t)}catch(t){return u.slice()}};t.exports.get=function(t){return u&&"[object Window]"==i.call(t)?c(t):o(r(t))}},function(t,e,n){var r=n(0);t.exports=function(t){var e=r.getKeys(t),n=r.getSymbols;if(n)for(var o,i=n(t),u=r.isEnum,c=0;i.length>c;)u.call(t,o=i[c++])&&e.push(o);return e}},function(t,e,n){var r=n(22);t.exports=Array.isArray||function(t){return"Array"==r(t)}},function(t,e){},function(t,e,n){"use strict";function r(t){return t&&t.__esModule?t:{default:t}}e.__esModule=!0;var o=n(58),i=r(o),u=n(62),c=r(u),a=n(15),s=r(a);e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,s.default)(e)));t.prototype=(0,c.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(i.default?(0,i.default)(t,e):t.__proto__=e)}},function(t,e,n){t.exports={default:n(59),__esModule:!0}},function(t,e,n){n(60),t.exports=n(1).Object.setPrototypeOf},function(t,e,n){var r=n(3);r(r.S,"Object",{setPrototypeOf:n(61).set})},function(t,e,n){var r=n(0).getDesc,o=n(24),i=n(23),u=function(t,e){if(i(t),!o(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,o){try{o=n(14)(Function.call,r(Object.prototype,"__proto__").set,2),o(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return u(t,n),e?t.__proto__=n:o(t,n),t}}({},!1):void 0),check:u}},function(t,e,n){t.exports={default:n(63),__esModule:!0}},function(t,e,n){var r=n(0);t.exports=function(t,e){return r.create(t,e)}},function(t,n){t.exports=e},function(t,e){t.exports=n},function(t,e){t.exports=r},function(t,e,n){"use strict";var r=n(6),o=n.n(r),i=function(t){return o.a.createElement("div",{className:"welcomepage"},t.modProps.welcome.text)};e.a=i},function(t,e){},function(e,n){e.exports=t}])});
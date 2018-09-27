define("dashboardtheme",["react","authui","react-redux","reactpages","reactwebcommon"],function(t,e,n,r,o){return function(t){var e={};function n(r){if(e[r])return e[r].exports;var o=e[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=t,n.c=e,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},n.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var o in t)n.d(r,o,function(e){return t[e]}.bind(null,o));return r},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=34)}([function(e,n){e.exports=t},function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(t,e,n){var r=n(18)("wks"),o=n(17),i=n(5).Symbol;t.exports=function(t){return r[t]||(r[t]=i&&i[t]||(i||o)("Symbol."+t))}},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,e,n){var r=n(50),o=n(13);t.exports=function(t){return r(o(t))}},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e,n){var r=n(5),o=n(3),i=n(24),u=function(t,e,n){var c,a,s,f=t&u.F,l=t&u.G,p=t&u.S,d=t&u.P,m=t&u.B,y=t&u.W,h=l?o:o[e]||(o[e]={}),v=l?r:p?r[e]:(r[e]||{}).prototype;for(c in l&&(n=e),n)(a=!f&&v&&c in v)&&c in h||(s=a?v[c]:n[c],h[c]=l&&"function"!=typeof v[c]?n[c]:m&&a?i(s,r):y&&v[c]==s?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(s):d&&"function"==typeof s?i(Function.call,s):s,d&&((h.prototype||(h.prototype={}))[c]=s))};u.F=1,u.G=2,u.S=4,u.P=8,u.B=16,u.W=32,t.exports=u},function(t,e,n){var r=n(1).setDesc,o=n(9),i=n(2)("toStringTag");t.exports=function(t,e,n){t&&!o(t=n?t:t.prototype,i)&&r(t,i,{configurable:!0,value:e})}},function(t,e){t.exports={}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e,n){var r=n(1),o=n(10);t.exports=n(19)?function(t,e,n){return r.setDesc(t,e,o(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){var r=n(14);t.exports=function(t){if(!r(t))throw TypeError(t+" is not an object!");return t}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e){var n=0,r=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+r).toString(36))}},function(t,e,n){var r=n(5),o=r["__core-js_shared__"]||(r["__core-js_shared__"]={});t.exports=function(t){return o[t]||(o[t]={})}},function(t,e,n){t.exports=!n(12)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){t.exports=n(11)},function(t,e){t.exports=!0},function(t,e,n){"use strict";var r=n(21),o=n(6),i=n(20),u=n(11),c=n(9),a=n(8),s=n(55),f=n(7),l=n(1).getProto,p=n(2)("iterator"),d=!([].keys&&"next"in[].keys()),m=function(){return this};t.exports=function(t,e,n,y,h,v,g){s(n,e,y);var b,_,x=function(t){if(!d&&t in O)return O[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},k=e+" Iterator",w="values"==h,S=!1,O=t.prototype,P=O[p]||O["@@iterator"]||h&&O[h],E=P||x(h);if(P){var j=l(E.call(new t));f(j,k,!0),!r&&c(O,"@@iterator")&&u(j,p,m),w&&"values"!==P.name&&(S=!0,E=function(){return P.call(this)})}if(r&&!g||!d&&!S&&O[p]||u(O,p,E),a[e]=E,a[k]=m,h)if(b={values:w?E:x("values"),keys:v?E:x("keys"),entries:w?x("entries"):E},g)for(_ in b)_ in O||i(O,_,b[_]);else o(o.P+o.F*(d||S),e,b);return b}},function(t,e,n){"use strict";e.__esModule=!0;var r=u(n(60)),o=u(n(49)),i="function"==typeof o.default&&"symbol"==typeof r.default?function(t){return typeof t}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":typeof t};function u(t){return t&&t.__esModule?t:{default:t}}e.default="function"==typeof o.default&&"symbol"===i(r.default)?function(t){return void 0===t?"undefined":i(t)}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":void 0===t?"undefined":i(t)}},function(t,e,n){var r=n(63);t.exports=function(t,e,n){if(r(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,r){return t.call(e,n,r)};case 3:return function(n,r,o){return t.call(e,n,r,o)}}return function(){return t.apply(e,arguments)}}},function(t,n){t.exports=e},function(t,e){t.exports=n},function(t,e){t.exports=r},function(t,e){t.exports=o},function(t,e,n){"use strict";e.__esModule=!0;var r=u(n(41)),o=u(n(37)),i=u(n(23));function u(t){return t&&t.__esModule?t:{default:t}}e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,i.default)(e)));t.prototype=(0,o.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(r.default?(0,r.default)(t,e):t.__proto__=e)}},function(t,e,n){"use strict";e.__esModule=!0;var r=function(t){return t&&t.__esModule?t:{default:t}}(n(23));e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,r.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){"use strict";e.__esModule=!0;var r=function(t){return t&&t.__esModule?t:{default:t}}(n(62));e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var o=e[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),(0,r.default)(t,o.key,o)}}return function(e,n,r){return n&&t(e.prototype,n),r&&t(e,r),e}}()},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){t.exports={default:n(67),__esModule:!0}},function(t,e,n){"use strict";n.r(e);var r,o=n(0),i=n.n(o),u=n(33),c=n.n(u),a=n(32),s=n.n(a),f=n(31),l=n.n(f),p=n(30),d=n.n(p),m=n(29),y=n.n(m),h=n(28),v=n(27),g=n(26),b=function(t){function e(t){return s()(this,e),d()(this,(e.__proto__||c()(e)).call(this,t))}return y()(e,t),l()(e,[{key:"render",value:function(){var t=this.props,e=t.module,n=e.properties.header,r=e.settings,o=r.infoBlock?r.infoBlock:"userBlock";return console.log("header properties",n,t),i.a.createElement(t.uikit.Block,{className:n.className?n.className:"header"},i.a.createElement(t.uikit.Block,{className:"logo"},n.image?i.a.createElement(t.uikit.Block,{className:"image"},i.a.createElement(h.Image,{src:n.image})):null,n.title?i.a.createElement(t.uikit.Block,{className:"title"},n.title):null),t.loggedIn?i.a.createElement(v.Panel,{id:o,className:"infoBlock"}):null)}}]),e}(i.a.Component),_=Object(g.connect)(function(t,e){return{loggedIn:"LoggedIn"==t.Security.status}})(b),x=function(t){return i.a.createElement("div",{className:"welcomepage"},t.modProps.welcome.text)},k=(n(35),n(25));function w(t,e,n,o,u,c){if(console.log("appname = ",t,"ins ",e,"mod",n,"settings",o),(r=this).properties=Application.Properties[e],r.settings=o,r.skipAuth=o.skipAuth,!_reg("Pages","home")){var a={id:"home",route:"/",components:{main:{type:"component",component:i.a.createElement(x,{modProps:r.properties})}}};Application.Register("Pages","home",a),Application.Register("Actions","Page_home",{url:"/"})}var s="authui",f="WebLoginForm";o&&!o.skipAuth&&o.loginModule&&(s=o.loginModule,f=o.loginComp);var l=c(s);l&&(r.logInComp=l[f]),function(){var t=[],e=[];e=r.settings&&r.settings.menu?r.settings.menu:Application.Properties.menu;console.log("dashboard menu",e),e?e.forEach(function(e){e.page?t.push({title:e.title,action:"Page_"+e.page}):t.push({title:e.title,action:e.action})}):t.push({title:"Home",action:"Page_home"});r.menu=t}()}function S(t,e){return t.components.menu=function(t){return i.a.createElement(e.Navbar,{items:r.menu,vertical:r.settings.vertical})},t}n.d(e,"Initialize",function(){return w}),n.d(e,"ProcessRoute",function(){return S}),n.d(e,"Theme",function(){return P});var O=function(t){var e=!!r.settings.vertical,n=(r.logInComp,i.a.createElement(t.uikit.Block,{className:"body"},i.a.createElement(t.uikit.Block,{className:e?"vertmenu":"horizmenu"},i.a.createElement(t.router.View,{name:"menu"})),i.a.createElement(t.uikit.Block,{className:e?"vertbody":"horizbody"},i.a.createElement(t.router.View,{name:"main"}))));return t.loggedIn||r.skipAuth?n:i.a.createElement(t.uikit.Block,{className:"dashlogin"},i.a.createElement(r.logInComp,null))},P=function(t){return i.a.createElement(t.uikit.UIWrapper,null,i.a.createElement(t.uikit.Block,{className:r.settings.className?r.settings.className+" dashboard":"dashboard"},i.a.createElement(_,{uikit:t.uikit,module:r}),i.a.createElement(k.LoginValidator,{validateService:"validate"},i.a.createElement(O,{uikit:t.uikit,router:t.router})),i.a.createElement(t.uikit.Block,{className:"footer"},t.children)))}},function(t,e){},function(t,e,n){var r=n(1);t.exports=function(t,e){return r.create(t,e)}},function(t,e,n){t.exports={default:n(36),__esModule:!0}},function(t,e,n){var r=n(1).getDesc,o=n(14),i=n(15),u=function(t,e){if(i(t),!o(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,o){try{(o=n(24)(Function.call,r(Object.prototype,"__proto__").set,2))(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return u(t,n),e?t.__proto__=n:o(t,n),t}}({},!1):void 0),check:u}},function(t,e,n){var r=n(6);r(r.S,"Object",{setPrototypeOf:n(38).set})},function(t,e,n){n(39),t.exports=n(3).Object.setPrototypeOf},function(t,e,n){t.exports={default:n(40),__esModule:!0}},function(t,e){},function(t,e,n){var r=n(16);t.exports=Array.isArray||function(t){return"Array"==r(t)}},function(t,e,n){var r=n(1);t.exports=function(t){var e=r.getKeys(t),n=r.getSymbols;if(n)for(var o,i=n(t),u=r.isEnum,c=0;i.length>c;)u.call(t,o=i[c++])&&e.push(o);return e}},function(t,e,n){var r=n(4),o=n(1).getNames,i={}.toString,u="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[];t.exports.get=function(t){return u&&"[object Window]"==i.call(t)?function(t){try{return o(t)}catch(t){return u.slice()}}(t):o(r(t))}},function(t,e,n){var r=n(1),o=n(4);t.exports=function(t,e){for(var n,i=o(t),u=r.getKeys(i),c=u.length,a=0;c>a;)if(i[n=u[a++]]===e)return n}},function(t,e,n){"use strict";var r=n(1),o=n(5),i=n(9),u=n(19),c=n(6),a=n(20),s=n(12),f=n(18),l=n(7),p=n(17),d=n(2),m=n(46),y=n(45),h=n(44),v=n(43),g=n(15),b=n(4),_=n(10),x=r.getDesc,k=r.setDesc,w=r.create,S=y.get,O=o.Symbol,P=o.JSON,E=P&&P.stringify,j=!1,M=d("_hidden"),N=r.isEnum,A=f("symbol-registry"),B=f("symbols"),I="function"==typeof O,D=Object.prototype,T=u&&s(function(){return 7!=w(k({},"a",{get:function(){return k(this,"a",{value:7}).a}})).a})?function(t,e,n){var r=x(D,e);r&&delete D[e],k(t,e,n),r&&t!==D&&k(D,e,r)}:k,C=function(t){var e=B[t]=w(O.prototype);return e._k=t,u&&j&&T(D,t,{configurable:!0,set:function(e){i(this,M)&&i(this[M],t)&&(this[M][t]=!1),T(this,t,_(1,e))}}),e},F=function(t){return"symbol"==typeof t},W=function(t,e,n){return n&&i(B,e)?(n.enumerable?(i(t,M)&&t[M][e]&&(t[M][e]=!1),n=w(n,{enumerable:_(0,!1)})):(i(t,M)||k(t,M,_(1,{})),t[M][e]=!0),T(t,e,n)):k(t,e,n)},L=function(t,e){g(t);for(var n,r=h(e=b(e)),o=0,i=r.length;i>o;)W(t,n=r[o++],e[n]);return t},z=function(t,e){return void 0===e?w(t):L(w(t),e)},J=function(t){var e=N.call(this,t);return!(e||!i(this,t)||!i(B,t)||i(this,M)&&this[M][t])||e},R=function(t,e){var n=x(t=b(t),e);return!n||!i(B,e)||i(t,M)&&t[M][e]||(n.enumerable=!0),n},G=function(t){for(var e,n=S(b(t)),r=[],o=0;n.length>o;)i(B,e=n[o++])||e==M||r.push(e);return r},K=function(t){for(var e,n=S(b(t)),r=[],o=0;n.length>o;)i(B,e=n[o++])&&r.push(B[e]);return r},V=s(function(){var t=O();return"[null]"!=E([t])||"{}"!=E({a:t})||"{}"!=E(Object(t))});I||(a((O=function(){if(F(this))throw TypeError("Symbol is not a constructor");return C(p(arguments.length>0?arguments[0]:void 0))}).prototype,"toString",function(){return this._k}),F=function(t){return t instanceof O},r.create=z,r.isEnum=J,r.getDesc=R,r.setDesc=W,r.setDescs=L,r.getNames=y.get=G,r.getSymbols=K,u&&!n(21)&&a(D,"propertyIsEnumerable",J,!0));var H={for:function(t){return i(A,t+="")?A[t]:A[t]=O(t)},keyFor:function(t){return m(A,t)},useSetter:function(){j=!0},useSimple:function(){j=!1}};r.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=d(t);H[t]=I?e:C(e)}),j=!0,c(c.G+c.W,{Symbol:O}),c(c.S,"Symbol",H),c(c.S+c.F*!I,"Object",{create:z,defineProperty:W,defineProperties:L,getOwnPropertyDescriptor:R,getOwnPropertyNames:G,getOwnPropertySymbols:K}),P&&c(c.S+c.F*(!I||V),"JSON",{stringify:function(t){if(void 0!==t&&!F(t)){for(var e,n,r=[t],o=1,i=arguments;i.length>o;)r.push(i[o++]);return"function"==typeof(e=r[1])&&(n=e),!n&&v(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!F(e))return e}),r[1]=e,E.apply(P,r)}}}),l(O,"Symbol"),l(Math,"Math",!0),l(o.JSON,"JSON",!0)},function(t,e,n){n(47),n(42),t.exports=n(3).Symbol},function(t,e,n){t.exports={default:n(48),__esModule:!0}},function(t,e,n){var r=n(16);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==r(t)?t.split(""):Object(t)}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e){t.exports=function(){}},function(t,e,n){"use strict";var r=n(52),o=n(51),i=n(8),u=n(4);t.exports=n(22)(Array,"Array",function(t,e){this._t=u(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,o(1)):o(0,"keys"==e?n:"values"==e?t[n]:[n,t[n]])},"values"),i.Arguments=i.Array,r("keys"),r("values"),r("entries")},function(t,e,n){n(53);var r=n(8);r.NodeList=r.HTMLCollection=r.Array},function(t,e,n){"use strict";var r=n(1),o=n(10),i=n(7),u={};n(11)(u,n(2)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=r.create(u,{next:o(1,n)}),i(t,e+" Iterator")}},function(t,e){var n=Math.ceil,r=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?r:n)(t)}},function(t,e,n){var r=n(56),o=n(13);t.exports=function(t){return function(e,n){var i,u,c=String(o(e)),a=r(n),s=c.length;return a<0||a>=s?t?"":void 0:(i=c.charCodeAt(a))<55296||i>56319||a+1===s||(u=c.charCodeAt(a+1))<56320||u>57343?t?c.charAt(a):i:t?c.slice(a,a+2):u-56320+(i-55296<<10)+65536}}},function(t,e,n){"use strict";var r=n(57)(!0);n(22)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=r(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){n(58),n(54),t.exports=n(2)("iterator")},function(t,e,n){t.exports={default:n(59),__esModule:!0}},function(t,e,n){var r=n(1);t.exports=function(t,e,n){return r.setDesc(t,e,n)}},function(t,e,n){t.exports={default:n(61),__esModule:!0}},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){var r=n(6),o=n(3),i=n(12);t.exports=function(t,e){var n=(o.Object||{})[t]||Object[t],u={};u[t]=e(n),r(r.S+r.F*i(function(){n(1)}),"Object",u)}},function(t,e,n){var r=n(13);t.exports=function(t){return Object(r(t))}},function(t,e,n){var r=n(65);n(64)("getPrototypeOf",function(t){return function(e){return t(r(e))}})},function(t,e,n){n(66),t.exports=n(3).Object.getPrototypeOf}])});
//# sourceMappingURL=index.js.map
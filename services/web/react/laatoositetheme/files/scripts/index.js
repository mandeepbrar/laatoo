define("laatoositetheme",["react","reactpages","react-redux","authui","reactwebcommon","prop-types"],function(t,e,n,r,o,i){return function(t){var e={};function n(r){if(e[r])return e[r].exports;var o=e[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=t,n.c=e,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},n.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var o in t)n.d(r,o,function(e){return t[e]}.bind(null,o));return r},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=68)}([function(e,n){e.exports=t},function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(t,n){t.exports=e},function(t,e,n){t.exports={default:n(34),__esModule:!0}},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){"use strict";e.__esModule=!0;var r=function(t){return t&&t.__esModule?t:{default:t}}(n(39));e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var o=e[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),(0,r.default)(t,o.key,o)}}return function(e,n,r){return n&&t(e.prototype,n),r&&t(e,r),e}}()},function(t,e,n){"use strict";e.__esModule=!0;var r=function(t){return t&&t.__esModule?t:{default:t}}(n(22));e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,r.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){"use strict";e.__esModule=!0;var r=u(n(60)),o=u(n(64)),i=u(n(22));function u(t){return t&&t.__esModule?t:{default:t}}e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,i.default)(e)));t.prototype=(0,o.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(r.default?(0,r.default)(t,e):t.__proto__=e)}},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,e,n){var r=n(27)("wks"),o=n(28),i=n(11).Symbol;t.exports=function(t){return r[t]||(r[t]=i&&i[t]||(i||o)("Symbol."+t))}},function(t,e,n){var r=n(11),o=n(8),i=n(21),u=function(t,e,n){var c,a,s,l=t&u.F,f=t&u.G,p=t&u.S,d=t&u.P,m=t&u.B,g=t&u.W,y=f?o:o[e]||(o[e]={}),h=f?r:p?r[e]:(r[e]||{}).prototype;for(c in f&&(n=e),n)(a=!l&&h&&c in h)&&c in y||(s=a?h[c]:n[c],y[c]=f&&"function"!=typeof h[c]?n[c]:m&&a?i(s,r):g&&h[c]==s?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(s):d&&"function"==typeof s?i(Function.call,s):s,d&&((y.prototype||(y.prototype={}))[c]=s))};u.F=1,u.G=2,u.S=4,u.P=8,u.B=16,u.W=32,t.exports=u},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e,n){var r=n(51),o=n(13);t.exports=function(t){return r(o(t))}},function(t,e){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e,n){var r=n(1),o=n(16);t.exports=n(26)?function(t,e,n){return r.setDesc(t,e,o(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports={}},function(t,e,n){var r=n(1).setDesc,o=n(17),i=n(9)("toStringTag");t.exports=function(t,e,n){t&&!o(t=n?t:t.prototype,i)&&r(t,i,{configurable:!0,value:e})}},function(t,e){t.exports=n},function(t,e,n){var r=n(38);t.exports=function(t,e,n){if(r(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,r){return t.call(e,n,r)};case 3:return function(n,r,o){return t.call(e,n,r,o)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){"use strict";e.__esModule=!0;var r=u(n(41)),o=u(n(52)),i="function"==typeof o.default&&"symbol"==typeof r.default?function(t){return typeof t}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":typeof t};function u(t){return t&&t.__esModule?t:{default:t}}e.default="function"==typeof o.default&&"symbol"===i(r.default)?function(t){return void 0===t?"undefined":i(t)}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":void 0===t?"undefined":i(t)}},function(t,e,n){"use strict";var r=n(24),o=n(10),i=n(25),u=n(15),c=n(17),a=n(18),s=n(46),l=n(19),f=n(1).getProto,p=n(9)("iterator"),d=!([].keys&&"next"in[].keys()),m=function(){return this};t.exports=function(t,e,n,g,y,h,v){s(n,e,g);var b,_,x=function(t){if(!d&&t in P)return P[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},S=e+" Iterator",k="values"==y,w=!1,P=t.prototype,O=P[p]||P["@@iterator"]||y&&P[y],E=O||x(y);if(O){var j=f(E.call(new t));l(j,S,!0),!r&&c(P,"@@iterator")&&u(j,p,m),k&&"values"!==O.name&&(w=!0,E=function(){return O.call(this)})}if(r&&!v||!d&&!w&&P[p]||u(P,p,E),a[e]=E,a[S]=m,y)if(b={values:k?E:x("values"),keys:h?E:x("keys"),entries:k?x("entries"):E},v)for(_ in b)_ in P||i(P,_,b[_]);else o(o.P+o.F*(d||w),e,b);return b}},function(t,e){t.exports=!0},function(t,e,n){t.exports=n(15)},function(t,e,n){t.exports=!n(14)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){var r=n(11),o=r["__core-js_shared__"]||(r["__core-js_shared__"]={});t.exports=function(t){return o[t]||(o[t]={})}},function(t,e){var n=0,r=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+r).toString(36))}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){var r=n(31);t.exports=function(t){if(!r(t))throw TypeError(t+" is not an object!");return t}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e){t.exports=r},function(t,e){t.exports=o},function(t,e,n){n(35),t.exports=n(8).Object.getPrototypeOf},function(t,e,n){var r=n(36);n(37)("getPrototypeOf",function(t){return function(e){return t(r(e))}})},function(t,e,n){var r=n(13);t.exports=function(t){return Object(r(t))}},function(t,e,n){var r=n(10),o=n(8),i=n(14);t.exports=function(t,e){var n=(o.Object||{})[t]||Object[t],u={};u[t]=e(n),r(r.S+r.F*i(function(){n(1)}),"Object",u)}},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){t.exports={default:n(40),__esModule:!0}},function(t,e,n){var r=n(1);t.exports=function(t,e,n){return r.setDesc(t,e,n)}},function(t,e,n){t.exports={default:n(42),__esModule:!0}},function(t,e,n){n(43),n(47),t.exports=n(9)("iterator")},function(t,e,n){"use strict";var r=n(44)(!0);n(23)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=r(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){var r=n(45),o=n(13);t.exports=function(t){return function(e,n){var i,u,c=String(o(e)),a=r(n),s=c.length;return a<0||a>=s?t?"":void 0:(i=c.charCodeAt(a))<55296||i>56319||a+1===s||(u=c.charCodeAt(a+1))<56320||u>57343?t?c.charAt(a):i:t?c.slice(a,a+2):u-56320+(i-55296<<10)+65536}}},function(t,e){var n=Math.ceil,r=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?r:n)(t)}},function(t,e,n){"use strict";var r=n(1),o=n(16),i=n(19),u={};n(15)(u,n(9)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=r.create(u,{next:o(1,n)}),i(t,e+" Iterator")}},function(t,e,n){n(48);var r=n(18);r.NodeList=r.HTMLCollection=r.Array},function(t,e,n){"use strict";var r=n(49),o=n(50),i=n(18),u=n(12);t.exports=n(23)(Array,"Array",function(t,e){this._t=u(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,o(1)):o(0,"keys"==e?n:"values"==e?t[n]:[n,t[n]])},"values"),i.Arguments=i.Array,r("keys"),r("values"),r("entries")},function(t,e){t.exports=function(){}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e,n){var r=n(29);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==r(t)?t.split(""):Object(t)}},function(t,e,n){t.exports={default:n(53),__esModule:!0}},function(t,e,n){n(54),n(59),t.exports=n(8).Symbol},function(t,e,n){"use strict";var r=n(1),o=n(11),i=n(17),u=n(26),c=n(10),a=n(25),s=n(14),l=n(27),f=n(19),p=n(28),d=n(9),m=n(55),g=n(56),y=n(57),h=n(58),v=n(30),b=n(12),_=n(16),x=r.getDesc,S=r.setDesc,k=r.create,w=g.get,P=o.Symbol,O=o.JSON,E=O&&O.stringify,j=!1,M=d("_hidden"),N=r.isEnum,I=l("symbol-registry"),A=l("symbols"),C="function"==typeof P,B=Object.prototype,T=u&&s(function(){return 7!=k(S({},"a",{get:function(){return S(this,"a",{value:7}).a}})).a})?function(t,e,n){var r=x(B,e);r&&delete B[e],S(t,e,n),r&&t!==B&&S(B,e,r)}:S,D=function(t){var e=A[t]=k(P.prototype);return e._k=t,u&&j&&T(B,t,{configurable:!0,set:function(e){i(this,M)&&i(this[M],t)&&(this[M][t]=!1),T(this,t,_(1,e))}}),e},F=function(t){return"symbol"==typeof t},K=function(t,e,n){return n&&i(A,e)?(n.enumerable?(i(t,M)&&t[M][e]&&(t[M][e]=!1),n=k(n,{enumerable:_(0,!1)})):(i(t,M)||S(t,M,_(1,{})),t[M][e]=!0),T(t,e,n)):S(t,e,n)},W=function(t,e){v(t);for(var n,r=y(e=b(e)),o=0,i=r.length;i>o;)K(t,n=r[o++],e[n]);return t},L=function(t,e){return void 0===e?k(t):W(k(t),e)},R=function(t){var e=N.call(this,t);return!(e||!i(this,t)||!i(A,t)||i(this,M)&&this[M][t])||e},z=function(t,e){var n=x(t=b(t),e);return!n||!i(A,e)||i(t,M)&&t[M][e]||(n.enumerable=!0),n},J=function(t){for(var e,n=w(b(t)),r=[],o=0;n.length>o;)i(A,e=n[o++])||e==M||r.push(e);return r},G=function(t){for(var e,n=w(b(t)),r=[],o=0;n.length>o;)i(A,e=n[o++])&&r.push(A[e]);return r},V=s(function(){var t=P();return"[null]"!=E([t])||"{}"!=E({a:t})||"{}"!=E(Object(t))});C||(a((P=function(){if(F(this))throw TypeError("Symbol is not a constructor");return D(p(arguments.length>0?arguments[0]:void 0))}).prototype,"toString",function(){return this._k}),F=function(t){return t instanceof P},r.create=L,r.isEnum=R,r.getDesc=z,r.setDesc=K,r.setDescs=W,r.getNames=g.get=J,r.getSymbols=G,u&&!n(24)&&a(B,"propertyIsEnumerable",R,!0));var H={for:function(t){return i(I,t+="")?I[t]:I[t]=P(t)},keyFor:function(t){return m(I,t)},useSetter:function(){j=!0},useSimple:function(){j=!1}};r.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=d(t);H[t]=C?e:D(e)}),j=!0,c(c.G+c.W,{Symbol:P}),c(c.S,"Symbol",H),c(c.S+c.F*!C,"Object",{create:L,defineProperty:K,defineProperties:W,getOwnPropertyDescriptor:z,getOwnPropertyNames:J,getOwnPropertySymbols:G}),O&&c(c.S+c.F*(!C||V),"JSON",{stringify:function(t){if(void 0!==t&&!F(t)){for(var e,n,r=[t],o=1,i=arguments;i.length>o;)r.push(i[o++]);return"function"==typeof(e=r[1])&&(n=e),!n&&h(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!F(e))return e}),r[1]=e,E.apply(O,r)}}}),f(P,"Symbol"),f(Math,"Math",!0),f(o.JSON,"JSON",!0)},function(t,e,n){var r=n(1),o=n(12);t.exports=function(t,e){for(var n,i=o(t),u=r.getKeys(i),c=u.length,a=0;c>a;)if(i[n=u[a++]]===e)return n}},function(t,e,n){var r=n(12),o=n(1).getNames,i={}.toString,u="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[];t.exports.get=function(t){return u&&"[object Window]"==i.call(t)?function(t){try{return o(t)}catch(t){return u.slice()}}(t):o(r(t))}},function(t,e,n){var r=n(1);t.exports=function(t){var e=r.getKeys(t),n=r.getSymbols;if(n)for(var o,i=n(t),u=r.isEnum,c=0;i.length>c;)u.call(t,o=i[c++])&&e.push(o);return e}},function(t,e,n){var r=n(29);t.exports=Array.isArray||function(t){return"Array"==r(t)}},function(t,e){},function(t,e,n){t.exports={default:n(61),__esModule:!0}},function(t,e,n){n(62),t.exports=n(8).Object.setPrototypeOf},function(t,e,n){var r=n(10);r(r.S,"Object",{setPrototypeOf:n(63).set})},function(t,e,n){var r=n(1).getDesc,o=n(31),i=n(30),u=function(t,e){if(i(t),!o(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,o){try{(o=n(21)(Function.call,r(Object.prototype,"__proto__").set,2))(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return u(t,n),e?t.__proto__=n:o(t,n),t}}({},!1):void 0),check:u}},function(t,e,n){t.exports={default:n(65),__esModule:!0}},function(t,e,n){var r=n(1);t.exports=function(t,e){return r.create(t,e)}},function(t,e){t.exports=i},function(t,e){},function(t,e,n){"use strict";n.r(e);var r,o=n(0),i=n.n(o),u=n(3),c=n.n(u),a=n(4),s=n.n(a),l=n(5),f=n.n(l),p=n(6),d=n.n(p),m=n(7),g=n.n(m),y=n(2),h=n(20),v=n(32),b=n(33),_=function(t){function e(t){return s()(this,e),d()(this,(e.__proto__||c()(e)).call(this,t))}return g()(e,t),f()(e,[{key:"render",value:function(){var t=this.props,e=t.module,n=e.properties.header,r=e.settings,o=r.infoBlock?r.infoBlock:"userBlock";return console.log("header properties",n,t),i.a.createElement(t.uikit.Block,{className:n.className?n.className:"header"},i.a.createElement(t.uikit.Block,{className:"logo"},n.image?i.a.createElement(t.uikit.Block,{className:"image"},i.a.createElement(b.Image,{src:n.image})):null,n.title?i.a.createElement(t.uikit.Block,{className:"title"},n.title):null),t.loggedIn?i.a.createElement(y.Panel,{id:o,className:"infoBlock"}):null)}}]),e}(i.a.Component),x=Object(h.connect)(function(t,e){return{loggedIn:"LoggedIn"==t.Security.status}})(_);var S=function(t){function e(t){return s()(this,e),d()(this,(e.__proto__||c()(e)).call(this,t))}return g()(e,t),f()(e,[{key:"render",value:function(){console.log("rendering theme ui",this.props);var t=this.props,e=!!r.settings.vertical;return i.a.createElement(t.uikit.Block,{className:r.settings.className?r.settings.className+" dashboard":"dashboard"},i.a.createElement(x,{uikit:t.uikit,module:r}),i.a.createElement(t.uikit.Block,{className:"body"},r.settings.showMenu?i.a.createElement(t.uikit.Block,{className:e?"vertmenu":"horizmenu"},i.a.createElement(t.router.View,{name:"menu",loggedIn:t.loggedIn})):null,i.a.createElement(t.uikit.Block,{className:e?"vertbody":"horizbody"},i.a.createElement(t.router.View,{name:"main",loggedIn:t.loggedIn}))),i.a.createElement(t.uikit.Block,{className:"footer"},t.children))}}]),e}(i.a.Component),k=function(t,e,n){r.settings.vertical;return i.a.createElement(t.uikit.UIWrapper,null,i.a.createElement(v.LoginValidator,{validateService:"validate"},i.a.createElement(S,{uikit:t.uikit,router:t.router})))};function w(t,e,n,o,u){return r.settings.showMenu&&(console.log("pre processing.......",t),t.menu=function(t){return console.log("menu items ",t,r.menu),i.a.createElement(u.Navbar,{items:r.menu,vertical:r.settings.vertical})}),t}var P,O=n(66);function E(t,e,n,r,o,u){return console.log("RenderPageComponent",r,e,t,o),"main"==e?i.a.createElement(j,{pageComp:t,pageKey:e,pageId:n,routerState:r,page:o,uikit:u}):"menu"==e?i.a.createElement(M,{menu:t,pageKey:e,pageId:n,routerState:r,page:o,uikit:u}):void 0}var j=function(t,e){return console.log("rendering site page",t,e),e.loggedIn?i.a.createElement(N,{pageId:t.pageId,placeholder:t.pageKey,routerState:t.routerState,description:t.pageComp}):i.a.createElement(t.uikit.Block,{className:"dashlogin"},i.a.createElement(P.logInComp,null))};j.contextTypes={routeParams:O.object,routerState:O.object,loggedIn:O.bool,user:O.object};var M=function(t,e){return console.log("rendering site menu",t,e),e.loggedIn?t.menu:null};M.contextTypes={routeParams:O.object,routerState:O.object,loggedIn:O.bool,user:O.object};var N=function(t){function e(){return s()(this,e),d()(this,(e.__proto__||c()(e)).apply(this,arguments))}return g()(e,t),f()(e,[{key:"getChildContext",value:function(){return{routeParams:this.props.routerState.params,routerState:this.props.routerState}}},{key:"render",value:function(){console.log("render page component**************",this.props);var t=this.props.pageId+this.props.placeholder;return i.a.createElement(y.Panel,{key:t,description:this.props.description})}}]),e}(i.a.Component);N.childContextTypes={routeParams:O.object,routerState:O.object};var I,A=function(t){return i.a.createElement("div",{className:"welcomepage"},t.modProps.welcome.text)};n(67);function C(t,e,n,o,u,c){if(console.log("appname = ",t,"ins ",e,"mod",n,"settings",o),function(t){P=t}(I=this),function(t){r=t}(I),I.properties=Application.Properties[e],I.settings=o,I.skipAuth=o.skipAuth,!_reg("Pages","home")){var a={id:"home",route:"/",components:{main:{type:"component",component:i.a.createElement(A,{modProps:I.properties})}}};Application.Register("Pages","home",a),Application.Register("Actions","Page_home",{url:"/"})}var s="authui",l="WebLoginForm";o&&!o.skipAuth&&o.loginModule&&(s=o.loginModule,l=o.loginComp);var f=c(s);f&&(I.logInComp=f[l]),I.settings.showMenu&&function(){var t=[],e=[];e=I.settings&&I.settings.menu?I.settings.menu:Application.Properties.menu;console.log("dashboard menu",e),e?e.forEach(function(e){e.page?t.push({title:e.title,action:"Page_"+e.page}):t.push({title:e.title,action:e.action})}):t.push({title:"Home",action:"Page_home"});I.menu=t}()}n.d(e,"Initialize",function(){return C}),n.d(e,"PreprocessPageComponents",function(){return w}),n.d(e,"RenderPageComponent",function(){return E}),n.d(e,"Theme",function(){return k})}])});
//# sourceMappingURL=index.js.map
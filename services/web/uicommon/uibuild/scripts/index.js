define("uicommon",["axios"],function(e){return function(e){var t={};function n(r){if(t[r])return t[r].exports;var o=t[r]={i:r,l:!1,exports:{}};return e[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=e,n.c=t,n.d=function(e,t,r){n.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:r})},n.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},n.t=function(e,t){if(1&t&&(e=n(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var o in e)n.d(r,o,function(t){return e[t]}.bind(null,o));return r},n.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return n.d(t,"a",t),t},n.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},n.p="/",n(n.s=89)}([function(e,t){var n=Object;e.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(e,t,n){var r=n(32)("wks"),o=n(33),i=n(4).Symbol;e.exports=function(e){return r[e]||(r[e]=i&&i[e]||(i||o)("Symbol."+e))}},function(e,t){var n=e.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(e,t,n){"use strict";t.__esModule=!0,t.default=function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}},function(e,t){var n=e.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(e,t,n){"use strict";t.__esModule=!0;var r,o=n(43),i=(r=o)&&r.__esModule?r:{default:r};t.default=function(){function e(e,t){for(var n=0;n<t.length;n++){var r=t[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),(0,i.default)(e,r.key,r)}}return function(t,n,r){return n&&e(t.prototype,n),r&&e(t,r),t}}()},function(e,t,n){e.exports={default:n(45),__esModule:!0}},function(e,t,n){var r=n(4),o=n(2),i=n(9),u=function(e,t,n){var a,f,c,s=e&u.F,l=e&u.G,d=e&u.S,p=e&u.P,v=e&u.B,h=e&u.W,b=l?o:o[t]||(o[t]={}),y=l?r:d?r[t]:(r[t]||{}).prototype;for(a in l&&(n=t),n)(f=!s&&y&&a in y)&&a in b||(c=f?y[a]:n[a],b[a]=l&&"function"!=typeof y[a]?n[a]:v&&f?i(c,r):h&&y[a]==c?function(e){var t=function(t){return this instanceof e?new e(t):e(t)};return t.prototype=e.prototype,t}(c):p&&"function"==typeof c?i(Function.call,c):c,p&&((b.prototype||(b.prototype={}))[a]=c))};u.F=1,u.G=2,u.S=4,u.P=8,u.B=16,u.W=32,e.exports=u},function(e,t,n){var r=n(16);e.exports=function(e){if(!r(e))throw TypeError(e+" is not an object!");return e}},function(e,t,n){var r=n(17);e.exports=function(e,t,n){if(r(e),void 0===t)return e;switch(n){case 1:return function(n){return e.call(t,n)};case 2:return function(n,r){return e.call(t,n,r)};case 3:return function(n,r,o){return e.call(t,n,r,o)}}return function(){return e.apply(t,arguments)}}},function(e,t){var n={}.toString;e.exports=function(e){return n.call(e).slice(8,-1)}},function(e,t){e.exports={}},function(e,t){e.exports=function(e){try{return!!e()}catch(e){return!0}}},function(e,t,n){e.exports=!n(12)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(e,t,n){var r=n(0).setDesc,o=n(24),i=n(1)("toStringTag");e.exports=function(e,t,n){e&&!o(e=n?e:e.prototype,i)&&r(e,i,{configurable:!0,value:t})}},function(e,t,n){var r=n(26),o=n(19);e.exports=function(e){return r(o(e))}},function(e,t){e.exports=function(e){return"object"==typeof e?null!==e:"function"==typeof e}},function(e,t){e.exports=function(e){if("function"!=typeof e)throw TypeError(e+" is not a function!");return e}},function(e,t,n){var r=n(19);e.exports=function(e){return Object(r(e))}},function(e,t){e.exports=function(e){if(null==e)throw TypeError("Can't call method on  "+e);return e}},function(e,t){e.exports=!0},function(e,t,n){e.exports=n(22)},function(e,t,n){var r=n(0),o=n(23);e.exports=n(13)?function(e,t,n){return r.setDesc(e,t,o(1,n))}:function(e,t,n){return e[t]=n,e}},function(e,t){e.exports=function(e,t){return{enumerable:!(1&e),configurable:!(2&e),writable:!(4&e),value:t}}},function(e,t){var n={}.hasOwnProperty;e.exports=function(e,t){return n.call(e,t)}},function(e,t,n){e.exports={default:n(69),__esModule:!0}},function(e,t,n){var r=n(10);e.exports=Object("z").propertyIsEnumerable(0)?Object:function(e){return"String"==r(e)?e.split(""):Object(e)}},function(e,t,n){var r=n(7),o=n(2),i=n(12);e.exports=function(e,t){var n=(o.Object||{})[e]||Object[e],u={};u[e]=t(n),r(r.S+r.F*i(function(){n(1)}),"Object",u)}},function(e,t,n){"use strict";t.__esModule=!0;var r=u(n(50)),o=u(n(57)),i="function"==typeof o.default&&"symbol"==typeof r.default?function(e){return typeof e}:function(e){return e&&"function"==typeof o.default&&e.constructor===o.default&&e!==o.default.prototype?"symbol":typeof e};function u(e){return e&&e.__esModule?e:{default:e}}t.default="function"==typeof o.default&&"symbol"===i(r.default)?function(e){return void 0===e?"undefined":i(e)}:function(e){return e&&"function"==typeof o.default&&e.constructor===o.default&&e!==o.default.prototype?"symbol":void 0===e?"undefined":i(e)}},function(e,t,n){"use strict";var r=n(52)(!0);n(31)(String,"String",function(e){this._t=String(e),this._i=0},function(){var e,t=this._t,n=this._i;return n>=t.length?{value:void 0,done:!0}:(e=r(t,n),this._i+=e.length,{value:e,done:!1})})},function(e,t){var n=Math.ceil,r=Math.floor;e.exports=function(e){return isNaN(e=+e)?0:(e>0?r:n)(e)}},function(e,t,n){"use strict";var r=n(20),o=n(7),i=n(21),u=n(22),a=n(24),f=n(11),c=n(53),s=n(14),l=n(0).getProto,d=n(1)("iterator"),p=!([].keys&&"next"in[].keys()),v=function(){return this};e.exports=function(e,t,n,h,b,y,m){c(n,t,h);var E,A,g=function(e){if(!p&&e in _)return _[e];switch(e){case"keys":case"values":return function(){return new n(this,e)}}return function(){return new n(this,e)}},x=t+" Iterator",S="values"==b,R=!1,_=e.prototype,P=_[d]||_["@@iterator"]||b&&_[b],O=P||g(b);if(P){var D=l(O.call(new e));s(D,x,!0),!r&&a(_,"@@iterator")&&u(D,d,v),S&&"values"!==P.name&&(R=!0,O=function(){return P.call(this)})}if(r&&!m||!p&&!R&&_[d]||u(_,d,O),f[t]=O,f[x]=v,b)if(E={values:S?O:g("values"),keys:y?O:g("keys"),entries:S?g("entries"):O},m)for(A in E)A in _||i(_,A,E[A]);else o(o.P+o.F*(p||R),t,E);return E}},function(e,t,n){var r=n(4),o=r["__core-js_shared__"]||(r["__core-js_shared__"]={});e.exports=function(e){return o[e]||(o[e]={})}},function(e,t){var n=0,r=Math.random();e.exports=function(e){return"Symbol(".concat(void 0===e?"":e,")_",(++n+r).toString(36))}},function(e,t,n){n(54);var r=n(11);r.NodeList=r.HTMLCollection=r.Array},function(e,t){},function(e,t,n){var r=n(0).getDesc,o=n(16),i=n(8),u=function(e,t){if(i(e),!o(t)&&null!==t)throw TypeError(t+": can't set as prototype!")};e.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(e,t,o){try{(o=n(9)(Function.call,r(Object.prototype,"__proto__").set,2))(e,[]),t=!(e instanceof Array)}catch(e){t=!0}return function(e,n){return u(e,n),t?e.__proto__=n:o(e,n),e}}({},!1):void 0),check:u}},function(e,t,n){var r=n(10),o=n(1)("toStringTag"),i="Arguments"==r(function(){return arguments}());e.exports=function(e){var t,n,u;return void 0===e?"Undefined":null===e?"Null":"string"==typeof(n=(t=Object(e))[o])?n:i?r(t):"Object"==(u=r(t))&&"function"==typeof t.callee?"Arguments":u}},function(e,t,n){e.exports={default:n(48),__esModule:!0}},function(e,t,n){"use strict";t.__esModule=!0;var r,o=n(28),i=(r=o)&&r.__esModule?r:{default:r};t.default=function(e,t){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!t||"object"!==(void 0===t?"undefined":(0,i.default)(t))&&"function"!=typeof t?e:t}},function(e,t,n){"use strict";t.__esModule=!0;var r=u(n(64)),o=u(n(67)),i=u(n(28));function u(e){return e&&e.__esModule?e:{default:e}}t.default=function(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Super expression must either be null or a function, not "+(void 0===t?"undefined":(0,i.default)(t)));e.prototype=(0,o.default)(t&&t.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),t&&(r.default?(0,r.default)(e,t):e.__proto__=t)}},function(e,t,n){e.exports={default:n(71),__esModule:!0}},function(t,n){t.exports=e},function(e,t,n){e.exports={default:n(44),__esModule:!0}},function(e,t,n){var r=n(0);e.exports=function(e,t,n){return r.setDesc(e,t,n)}},function(e,t,n){n(46),e.exports=n(2).Object.assign},function(e,t,n){var r=n(7);r(r.S+r.F,"Object",{assign:n(47)})},function(e,t,n){var r=n(0),o=n(18),i=n(26);e.exports=n(12)(function(){var e=Object.assign,t={},n={},r=Symbol(),o="abcdefghijklmnopqrst";return t[r]=7,o.split("").forEach(function(e){n[e]=e}),7!=e({},t)[r]||Object.keys(e({},n)).join("")!=o})?function(e,t){for(var n=o(e),u=arguments,a=u.length,f=1,c=r.getKeys,s=r.getSymbols,l=r.isEnum;a>f;)for(var d,p=i(u[f++]),v=s?c(p).concat(s(p)):c(p),h=v.length,b=0;h>b;)l.call(p,d=v[b++])&&(n[d]=p[d]);return n}:Object.assign},function(e,t,n){n(49),e.exports=n(2).Object.getPrototypeOf},function(e,t,n){var r=n(18);n(27)("getPrototypeOf",function(e){return function(t){return e(r(t))}})},function(e,t,n){e.exports={default:n(51),__esModule:!0}},function(e,t,n){n(29),n(34),e.exports=n(1)("iterator")},function(e,t,n){var r=n(30),o=n(19);e.exports=function(e){return function(t,n){var i,u,a=String(o(t)),f=r(n),c=a.length;return f<0||f>=c?e?"":void 0:(i=a.charCodeAt(f))<55296||i>56319||f+1===c||(u=a.charCodeAt(f+1))<56320||u>57343?e?a.charAt(f):i:e?a.slice(f,f+2):u-56320+(i-55296<<10)+65536}}},function(e,t,n){"use strict";var r=n(0),o=n(23),i=n(14),u={};n(22)(u,n(1)("iterator"),function(){return this}),e.exports=function(e,t,n){e.prototype=r.create(u,{next:o(1,n)}),i(e,t+" Iterator")}},function(e,t,n){"use strict";var r=n(55),o=n(56),i=n(11),u=n(15);e.exports=n(31)(Array,"Array",function(e,t){this._t=u(e),this._i=0,this._k=t},function(){var e=this._t,t=this._k,n=this._i++;return!e||n>=e.length?(this._t=void 0,o(1)):o(0,"keys"==t?n:"values"==t?e[n]:[n,e[n]])},"values"),i.Arguments=i.Array,r("keys"),r("values"),r("entries")},function(e,t){e.exports=function(){}},function(e,t){e.exports=function(e,t){return{value:t,done:!!e}}},function(e,t,n){e.exports={default:n(58),__esModule:!0}},function(e,t,n){n(59),n(35),e.exports=n(2).Symbol},function(e,t,n){"use strict";var r=n(0),o=n(4),i=n(24),u=n(13),a=n(7),f=n(21),c=n(12),s=n(32),l=n(14),d=n(33),p=n(1),v=n(60),h=n(61),b=n(62),y=n(63),m=n(8),E=n(15),A=n(23),g=r.getDesc,x=r.setDesc,S=r.create,R=h.get,_=o.Symbol,P=o.JSON,O=P&&P.stringify,D=!1,w=p("_hidden"),j=r.isEnum,L=s("symbol-registry"),T=s("symbols"),F="function"==typeof _,k=Object.prototype,B=u&&c(function(){return 7!=S(x({},"a",{get:function(){return x(this,"a",{value:7}).a}})).a})?function(e,t,n){var r=g(k,t);r&&delete k[t],x(e,t,n),r&&e!==k&&x(k,t,r)}:x,M=function(e){var t=T[e]=S(_.prototype);return t._k=e,u&&D&&B(k,e,{configurable:!0,set:function(t){i(this,w)&&i(this[w],e)&&(this[w][e]=!1),B(this,e,A(1,t))}}),t},G=function(e){return"symbol"==typeof e},q=function(e,t,n){return n&&i(T,t)?(n.enumerable?(i(e,w)&&e[w][t]&&(e[w][t]=!1),n=S(n,{enumerable:A(0,!1)})):(i(e,w)||x(e,w,A(1,{})),e[w][t]=!0),B(e,t,n)):x(e,t,n)},N=function(e,t){m(e);for(var n,r=b(t=E(t)),o=0,i=r.length;i>o;)q(e,n=r[o++],t[n]);return e},U=function(e,t){return void 0===t?S(e):N(S(e),t)},C=function(e){var t=j.call(this,e);return!(t||!i(this,e)||!i(T,e)||i(this,w)&&this[w][e])||t},I=function(e,t){var n=g(e=E(e),t);return!n||!i(T,t)||i(e,w)&&e[w][t]||(n.enumerable=!0),n},H=function(e){for(var t,n=R(E(e)),r=[],o=0;n.length>o;)i(T,t=n[o++])||t==w||r.push(t);return r},W=function(e){for(var t,n=R(E(e)),r=[],o=0;n.length>o;)i(T,t=n[o++])&&r.push(T[t]);return r},Y=c(function(){var e=_();return"[null]"!=O([e])||"{}"!=O({a:e})||"{}"!=O(Object(e))});F||(f((_=function(){if(G(this))throw TypeError("Symbol is not a constructor");return M(d(arguments.length>0?arguments[0]:void 0))}).prototype,"toString",function(){return this._k}),G=function(e){return e instanceof _},r.create=U,r.isEnum=C,r.getDesc=I,r.setDesc=q,r.setDescs=N,r.getNames=h.get=H,r.getSymbols=W,u&&!n(20)&&f(k,"propertyIsEnumerable",C,!0));var K={for:function(e){return i(L,e+="")?L[e]:L[e]=_(e)},keyFor:function(e){return v(L,e)},useSetter:function(){D=!0},useSimple:function(){D=!1}};r.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(e){var t=p(e);K[e]=F?t:M(t)}),D=!0,a(a.G+a.W,{Symbol:_}),a(a.S,"Symbol",K),a(a.S+a.F*!F,"Object",{create:U,defineProperty:q,defineProperties:N,getOwnPropertyDescriptor:I,getOwnPropertyNames:H,getOwnPropertySymbols:W}),P&&a(a.S+a.F*(!F||Y),"JSON",{stringify:function(e){if(void 0!==e&&!G(e)){for(var t,n,r=[e],o=1,i=arguments;i.length>o;)r.push(i[o++]);return"function"==typeof(t=r[1])&&(n=t),!n&&y(t)||(t=function(e,t){if(n&&(t=n.call(this,e,t)),!G(t))return t}),r[1]=t,O.apply(P,r)}}}),l(_,"Symbol"),l(Math,"Math",!0),l(o.JSON,"JSON",!0)},function(e,t,n){var r=n(0),o=n(15);e.exports=function(e,t){for(var n,i=o(e),u=r.getKeys(i),a=u.length,f=0;a>f;)if(i[n=u[f++]]===t)return n}},function(e,t,n){var r=n(15),o=n(0).getNames,i={}.toString,u="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[];e.exports.get=function(e){return u&&"[object Window]"==i.call(e)?function(e){try{return o(e)}catch(e){return u.slice()}}(e):o(r(e))}},function(e,t,n){var r=n(0);e.exports=function(e){var t=r.getKeys(e),n=r.getSymbols;if(n)for(var o,i=n(e),u=r.isEnum,a=0;i.length>a;)u.call(e,o=i[a++])&&t.push(o);return t}},function(e,t,n){var r=n(10);e.exports=Array.isArray||function(e){return"Array"==r(e)}},function(e,t,n){e.exports={default:n(65),__esModule:!0}},function(e,t,n){n(66),e.exports=n(2).Object.setPrototypeOf},function(e,t,n){var r=n(7);r(r.S,"Object",{setPrototypeOf:n(36).set})},function(e,t,n){e.exports={default:n(68),__esModule:!0}},function(e,t,n){var r=n(0);e.exports=function(e,t){return r.create(e,t)}},function(e,t,n){n(70),e.exports=n(2).Object.keys},function(e,t,n){var r=n(18);n(27)("keys",function(e){return function(t){return e(r(t))}})},function(e,t,n){n(35),n(29),n(34),n(72),e.exports=n(2).Promise},function(e,t,n){"use strict";var r,o=n(0),i=n(20),u=n(4),a=n(9),f=n(37),c=n(7),s=n(16),l=n(8),d=n(17),p=n(73),v=n(74),h=n(36).set,b=n(79),y=n(1)("species"),m=n(80),E=n(81),A=u.process,g="process"==f(A),x=u.Promise,S=function(){},R=function(e){var t,n=new x(S);return e&&(n.constructor=function(e){e(S,S)}),(t=x.resolve(n)).catch(S),t===n},_=function(){var e=!1;function t(e){var n=new x(e);return h(n,t.prototype),n}try{if(e=x&&x.resolve&&R(),h(t,x),t.prototype=o.create(x.prototype,{constructor:{value:t}}),t.resolve(5).then(function(){})instanceof t||(e=!1),e&&n(13)){var r=!1;x.resolve(o.setDesc({},"then",{get:function(){r=!0}})),e=r}}catch(t){e=!1}return e}(),P=function(e){var t=l(e)[y];return null!=t?t:e},O=function(e){var t;return!(!s(e)||"function"!=typeof(t=e.then))&&t},D=function(e){var t,n;this.promise=new e(function(e,r){if(void 0!==t||void 0!==n)throw TypeError("Bad Promise constructor");t=e,n=r}),this.resolve=d(t),this.reject=d(n)},w=function(e){try{e()}catch(e){return{error:e}}},j=function(e,t){if(!e.n){e.n=!0;var n=e.c;E(function(){for(var r=e.v,o=1==e.s,i=0,a=function(t){var n,i,u=o?t.ok:t.fail,a=t.resolve,f=t.reject;try{u?(o||(e.h=!0),(n=!0===u?r:u(r))===t.promise?f(TypeError("Promise-chain cycle")):(i=O(n))?i.call(n,a,f):a(n)):f(r)}catch(e){f(e)}};n.length>i;)a(n[i++]);n.length=0,e.n=!1,t&&setTimeout(function(){var t,n,o=e.p;L(o)&&(g?A.emit("unhandledRejection",r,o):(t=u.onunhandledrejection)?t({promise:o,reason:r}):(n=u.console)&&n.error&&n.error("Unhandled promise rejection",r)),e.a=void 0},1)})}},L=function(e){var t,n=e._d,r=n.a||n.c,o=0;if(n.h)return!1;for(;r.length>o;)if((t=r[o++]).fail||!L(t.promise))return!1;return!0},T=function(e){var t=this;t.d||(t.d=!0,(t=t.r||t).v=e,t.s=2,t.a=t.c.slice(),j(t,!0))},F=function(e){var t,n=this;if(!n.d){n.d=!0,n=n.r||n;try{if(n.p===e)throw TypeError("Promise can't be resolved itself");(t=O(e))?E(function(){var r={r:n,d:!1};try{t.call(e,a(F,r,1),a(T,r,1))}catch(e){T.call(r,e)}}):(n.v=e,n.s=1,j(n,!1))}catch(e){T.call({r:n,d:!1},e)}}};_||(x=function(e){d(e);var t=this._d={p:p(this,x,"Promise"),c:[],a:void 0,s:0,d:!1,v:void 0,h:!1,n:!1};try{e(a(F,t,1),a(T,t,1))}catch(e){T.call(t,e)}},n(86)(x.prototype,{then:function(e,t){var n=new D(m(this,x)),r=n.promise,o=this._d;return n.ok="function"!=typeof e||e,n.fail="function"==typeof t&&t,o.c.push(n),o.a&&o.a.push(n),o.s&&j(o,!1),r},catch:function(e){return this.then(void 0,e)}})),c(c.G+c.W+c.F*!_,{Promise:x}),n(14)(x,"Promise"),n(87)("Promise"),r=n(2).Promise,c(c.S+c.F*!_,"Promise",{reject:function(e){var t=new D(this);return(0,t.reject)(e),t.promise}}),c(c.S+c.F*(!_||R(!0)),"Promise",{resolve:function(e){if(e instanceof x&&(t=e.constructor,n=this,i&&t===x&&n===r||b(t,n)))return e;var t,n,o=new D(this);return(0,o.resolve)(e),o.promise}}),c(c.S+c.F*!(_&&n(88)(function(e){x.all(e).catch(function(){})})),"Promise",{all:function(e){var t=P(this),n=new D(t),r=n.resolve,i=n.reject,u=[],a=w(function(){v(e,!1,u.push,u);var n=u.length,a=Array(n);n?o.each.call(u,function(e,o){var u=!1;t.resolve(e).then(function(e){u||(u=!0,a[o]=e,--n||r(a))},i)}):r(a)});return a&&i(a.error),n.promise},race:function(e){var t=P(this),n=new D(t),r=n.reject,o=w(function(){v(e,!1,function(e){t.resolve(e).then(n.resolve,r)})});return o&&r(o.error),n.promise}})},function(e,t){e.exports=function(e,t,n){if(!(e instanceof t))throw TypeError(n+": use the 'new' operator!");return e}},function(e,t,n){var r=n(9),o=n(75),i=n(76),u=n(8),a=n(77),f=n(78);e.exports=function(e,t,n,c){var s,l,d,p=f(e),v=r(n,c,t?2:1),h=0;if("function"!=typeof p)throw TypeError(e+" is not iterable!");if(i(p))for(s=a(e.length);s>h;h++)t?v(u(l=e[h])[0],l[1]):v(e[h]);else for(d=p.call(e);!(l=d.next()).done;)o(d,v,l.value,t)}},function(e,t,n){var r=n(8);e.exports=function(e,t,n,o){try{return o?t(r(n)[0],n[1]):t(n)}catch(t){var i=e.return;throw void 0!==i&&r(i.call(e)),t}}},function(e,t,n){var r=n(11),o=n(1)("iterator"),i=Array.prototype;e.exports=function(e){return void 0!==e&&(r.Array===e||i[o]===e)}},function(e,t,n){var r=n(30),o=Math.min;e.exports=function(e){return e>0?o(r(e),9007199254740991):0}},function(e,t,n){var r=n(37),o=n(1)("iterator"),i=n(11);e.exports=n(2).getIteratorMethod=function(e){if(null!=e)return e[o]||e["@@iterator"]||i[r(e)]}},function(e,t){e.exports=Object.is||function(e,t){return e===t?0!==e||1/e==1/t:e!=e&&t!=t}},function(e,t,n){var r=n(8),o=n(17),i=n(1)("species");e.exports=function(e,t){var n,u=r(e).constructor;return void 0===u||null==(n=r(u)[i])?t:o(n)}},function(e,t,n){var r,o,i,u=n(4),a=n(82).set,f=u.MutationObserver||u.WebKitMutationObserver,c=u.process,s=u.Promise,l="process"==n(10)(c),d=function(){var e,t,n;for(l&&(e=c.domain)&&(c.domain=null,e.exit());r;)t=r.domain,n=r.fn,t&&t.enter(),n(),t&&t.exit(),r=r.next;o=void 0,e&&e.enter()};if(l)i=function(){c.nextTick(d)};else if(f){var p=1,v=document.createTextNode("");new f(d).observe(v,{characterData:!0}),i=function(){v.data=p=-p}}else i=s&&s.resolve?function(){s.resolve().then(d)}:function(){a.call(u,d)};e.exports=function(e){var t={fn:e,next:void 0,domain:l&&c.domain};o&&(o.next=t),r||(r=t,i()),o=t}},function(e,t,n){var r,o,i,u=n(9),a=n(83),f=n(84),c=n(85),s=n(4),l=s.process,d=s.setImmediate,p=s.clearImmediate,v=s.MessageChannel,h=0,b={},y=function(){var e=+this;if(b.hasOwnProperty(e)){var t=b[e];delete b[e],t()}},m=function(e){y.call(e.data)};d&&p||(d=function(e){for(var t=[],n=1;arguments.length>n;)t.push(arguments[n++]);return b[++h]=function(){a("function"==typeof e?e:Function(e),t)},r(h),h},p=function(e){delete b[e]},"process"==n(10)(l)?r=function(e){l.nextTick(u(y,e,1))}:v?(i=(o=new v).port2,o.port1.onmessage=m,r=u(i.postMessage,i,1)):s.addEventListener&&"function"==typeof postMessage&&!s.importScripts?(r=function(e){s.postMessage(e+"","*")},s.addEventListener("message",m,!1)):r="onreadystatechange"in c("script")?function(e){f.appendChild(c("script")).onreadystatechange=function(){f.removeChild(this),y.call(e)}}:function(e){setTimeout(u(y,e,1),0)}),e.exports={set:d,clear:p}},function(e,t){e.exports=function(e,t,n){var r=void 0===n;switch(t.length){case 0:return r?e():e.call(n);case 1:return r?e(t[0]):e.call(n,t[0]);case 2:return r?e(t[0],t[1]):e.call(n,t[0],t[1]);case 3:return r?e(t[0],t[1],t[2]):e.call(n,t[0],t[1],t[2]);case 4:return r?e(t[0],t[1],t[2],t[3]):e.call(n,t[0],t[1],t[2],t[3])}return e.apply(n,t)}},function(e,t,n){e.exports=n(4).document&&document.documentElement},function(e,t,n){var r=n(16),o=n(4).document,i=r(o)&&r(o.createElement);e.exports=function(e){return i?o.createElement(e):{}}},function(e,t,n){var r=n(21);e.exports=function(e,t){for(var n in t)r(e,n,t[n]);return e}},function(e,t,n){"use strict";var r=n(2),o=n(0),i=n(13),u=n(1)("species");e.exports=function(e){var t=r[e];i&&t&&!t[u]&&o.setDesc(t,u,{configurable:!0,get:function(){return this}})}},function(e,t,n){var r=n(1)("iterator"),o=!1;try{var i=[7][r]();i.return=function(){o=!0},Array.from(i,function(){throw 2})}catch(e){}e.exports=function(e,t){if(!t&&!o)return!1;var n=!1;try{var i=[7],u=i[r]();u.next=function(){return{done:n=!0}},i[r]=function(){return u},e(i)}catch(e){}return n}},function(e,t,n){"use strict";n.r(t);var r=n(3),o=n.n(r),i=n(5),u=n.n(i),a=n(6),f=n.n(a),c=function(){function e(){o()(this,e),this.ParameterSeparatorRequest=this.ParameterSeparatorRequest.bind(this),this.DefaultRequest=this.DefaultRequest.bind(this),this.URLParamsRequest=this.URLParamsRequest.bind(this)}return u()(e,[{key:"ParameterSeparatorRequest",value:function(e,t,n,r){var o={};return null==t&&(t={}),o.params=e,o.data=t,o.urlparams=n,o.headers=r,o.GetRequest=function(t){if("http"==t){var n={};if(n.data=o.data,n.headers=o.headers,null==o.params)return n.params=null,n.urlparams=null,n;var i={},u={},a=0;if(null!=o.urlparams)for(var c in o.urlparams)c in o.params&&(i[c]=o.params[c],a+=1);if(a>0){var s=0;for(var c in o.params)c in i||(u[c]=o.params[c],s+=1);return s>0?(n.urlparams=i,n.params=u):(n.urlparams=i,n.params=null),n}return n.urlparams=null,n.params=e,n}var l={};return l.data=o.data,l.params=e,l.params=f()({},e,r),l},o}},{key:"DefaultRequest",value:function(e,t,n){var r={};return null==t&&(t={}),r.params=e,r.data=t,r.headers=n,r.GetRequest=function(e){var t={};return t.data=r.data,t.params=r.params,t.urlparams=null,t.headers=r.headers,t},r}},{key:"URLParamsRequest",value:function(e,t,n){var r={};return null==t&&(t={}),r.data=t,r.urlparams=e,r.headers=n,r.GetRequest=function(e){if("http"==e){var t={};return t.data=r.data,t.params=null,t.urlparams=r.urlparams,t.headers=r.headers,t}var n={};return n.data=r.data,n.params=f()({},r.urlparams,r.headers),n},r}}]),e}(),s=function(){function e(t,n){var r=this;o()(this,e),this.SetPrefix=function(e){r.EntityPrefix=e},this.DataSource=t,this.RequestBuilder=n,this.GetEntity=this.GetEntity.bind(this),this.SaveEntity=this.SaveEntity.bind(this),this.DeleteEntity=this.DeleteEntity.bind(this),this.PutEntity=this.PutEntity.bind(this),this.UpdateEntity=this.UpdateEntity.bind(this),this.EntityPrefix="/"}return u()(e,[{key:"GetEntity",value:function(e,t,n,r){if(r){var o=this.RequestBuilder.URLParamsRequest({":id":t},null,n);return this.DataSource.ExecuteService(r,o)}var i={method:"GET"};i.url=this.EntityPrefix+e.toLowerCase()+"/"+t;o=this.RequestBuilder.DefaultRequest(null,null,n);return this.DataSource.ExecuteServiceObject(i,o)}},{key:"SaveEntity",value:function(e,t,n,r){var o=this.RequestBuilder.DefaultRequest(null,t,n);if(r)return this.DataSource.ExecuteService(r,o);var i={method:"POST"};return i.url=this.EntityPrefix+e.toLowerCase(),this.DataSource.ExecuteServiceObject(i,o)}},{key:"DeleteEntity",value:function(e,t,n,r){if(r){var o=this.RequestBuilder.URLParamsRequest({":id":t},null,n);return this.DataSource.ExecuteService(r,o)}var i={method:"DELETE"};i.url=this.EntityPrefix+e.toLowerCase()+"/"+t;o=this.RequestBuilder.DefaultRequest(null,null,n);return this.DataSource.ExecuteServiceObject(i,o)}},{key:"PutEntity",value:function(e,t,n,r,o){if(o){var i=this.RequestBuilder.URLParamsRequest({":id":t},null,r);return this.DataSource.ExecuteService(o,i)}var u={method:"PUT"};u.url=this.EntityPrefix+e.toLowerCase()+"/"+t;i=this.RequestBuilder.DefaultRequest(null,n,r);return this.DataSource.ExecuteServiceObject(u,i)}},{key:"ListEntities",value:function(e,t,n,r){if(r){var o=this.RequestBuilder.URLParamsRequest(t,null,n);return this.DataSource.ExecuteService(r,o)}var i={method:"POST"};i.url=this.EntityPrefix+e.toLowerCase()+"/view";o=this.RequestBuilder.DefaultRequest(t,null,n);return this.DataSource.ExecuteServiceObject(i,o)}},{key:"UpdateEntity",value:function(e,t,n,r,o){if(o){var i=this.RequestBuilder.URLParamsRequest({":id":t},null,r);return this.DataSource.ExecuteService(o,i)}var u={method:"PUT"};u.url=this.EntityPrefix+e.toLowerCase()+"/"+t;i=this.RequestBuilder.DefaultRequest(null,n,r);return this.DataSource.ExecuteServiceObject(u,i)}}]),e}();console.log("uicommon - application services",Application);var l=new(function(){function e(){o()(this,e),this.ExecuteService=this.ExecuteService.bind(this),this.ExecuteServiceObject=this.ExecuteServiceObject.bind(this)}return u()(e,[{key:"ExecuteService",value:function(e,t){var n=arguments.length>2&&void 0!==arguments[2]?arguments[2]:null,r=_reg("Services",e);if(null!=r&&null!=t)return this.ExecuteServiceObject(r,t,n);throw new Error("Service not found "+e)}},{key:"ExecuteServiceObject",value:function(e,t){var n=arguments.length>2&&void 0!==arguments[2]?arguments[2]:null;if(e.protocol||(e.protocol="http"),null!=e&&null!=t){var r=_reg("DataSourceHandlers",e.protocol);if(null==r)throw console.log("Requested service for handler",e),new Error("Invalid protocol handler");return r.ExecuteServiceObject(e,t,n)}}}]),e}()),d=new c,p=new s(l,d),v={Success:"Success",Unauthorized:"Unauthorized",InternalError:"InternalError",BadRequest:"BadRequest",Failure:"Failure"},h=n(38),b=n.n(h),y=n(39),m=n.n(y),E=n(40),A=n.n(E),g=function(e){function t(e,n,r){o()(this,t);var i=m()(this,(t.__proto__||b()(t)).call(this,e));return i.name=i.constructor.name,i.message=e,"function"==typeof Error.captureStackTrace?Error.captureStackTrace(i,i.constructor):i.stack=new Error(e).stack,i.type=e,i.rootError=n,i.args=r,i}return A()(t,e),t}(Error);function x(e,t,n){var r=t instanceof Error;return console.log("created action",e,t,n,r),{type:e,payload:t,meta:n,error:r}}function S(e,t){var n=e;if(t)for(var r in t){var o=t[r];n=n.replace(new RegExp(":"+r,"g"),o)}return n}function R(e){var t=!0;if(e&&""!=e){var n=localStorage.permissions;n&&n.indexOf(e)<0&&(t=!1)}return t}var _={WHITE:"#ffffff",BLACK:"#000000",TRANSPARENT:"transparent",DELTAGREY:{50:"#e4e4e4",100:"#bdbdbd",200:"#a1a1a1",300:"#7e7e7e",400:"#6e6e6e",500:"#5f5f5f",600:"#505050",700:"#404040",800:"#313131",900:"#222222",A100:"#e4e4e4",A200:"#bdbdbd",A400:"#6e6e6e",A700:"#404040"},DELTABLUE:{50:"#e6f5ff",100:"#9ad8ff",200:"#62c2ff",300:"#1aa7ff",400:"#009afb",500:"#0087dc",600:"#0074bd",700:"#00619f",800:"#004f80",900:"#003c62",A100:"#e6f5ff",A200:"#9ad8ff",A400:"#009afb",A700:"#00619f"},DELTAORANGE:{50:"#FFFDFA",100:"#FFDAAE",200:"#FFC076",300:"#FF9F2E",400:"#FF9110",500:"#F08200",600:"#D17100",700:"#B36100",800:"#945000",900:"#764000",A100:"#FFFDFA",A200:"#FFDAAE",A400:"#FF9110",A700:"#B36100"},DELTAGREEN:{50:"#C5F8CD",100:"#81EF91",200:"#50E965",300:"#1BD636",400:"#17BB2F",500:"#14A028",600:"#118521",700:"#0D6A1A",800:"#0A4E14",900:"#06330D",A100:"#C5F8CD",A200:"#81EF91",A400:"#17BB2F",A700:"#0D6A1A"},RED:{50:"ffebee",100:"#ffcdd2",200:"#ef9a9a",300:"#e57373",400:"#ef5350",500:"#f44336",600:"#e53935",700:"#d32f2f",800:"#c62828",900:"#b71c1c",A100:"#ff8a80",A200:"#ff5252",A400:"#ff1744",A700:"#d50000"},PINK:{50:"#fce4ec",100:"#f8bbd0",200:"#f48fb1",300:"#f06292",400:"#ec407a",500:"#e91e63",600:"#d81b60",700:"#c2185b",800:"#ad1457",900:"#880e4f",A100:"#ff80ab",A200:"#ff4081",A400:"#f50057",A700:"#c51162"},PURPLE:{50:"#f3e5f5",100:"#e1bee7",200:"#ce93d8",300:"#ba68c8",400:"#ab47bc",500:"#9c27b0",600:"#8e24aa",700:"#7b1fa2",800:"#6a1b9a",900:"#4a148c",A100:"#ea80fc",A200:"#e040fb",A400:"#d500f9",A700:"#aa00ff"},DEEPPRUPLE:{50:"#ede7f6",100:"#d1c4e9",200:"#b39ddb",300:"#9575cd",400:"#7e57c2",500:"#673ab7",600:"#5e35b1",700:"#512DA8",800:"#4527A0",900:"#311B92",A100:"#b388ff",A200:"#7c4dff",A400:"#651fff",A700:"#6200ea"},INDIGO:{50:"#e8eaf6",100:"#c5cae9",200:"#9fa8da",300:"#7986cb",400:"#5c6bc0",500:"#3f51b5",600:"#3949ab",700:"#303F9F",800:"#283593",900:"#1A237E",A100:"#8c9eff",A200:"#536dfe",A400:"#3d5afe",A700:"#304ffe"},BLUE:{50:"#e3f2fd",100:"#bbdefb",200:"#90caf9",300:"#64b5f6",400:"#42a5f5",500:"#2196f3",600:"#1e88e5",700:"#1976d2",800:"#1565c0",900:"#0d47a1",A100:"#82b1ff",A200:"#448aff",A400:"#2979ff",A700:"#2962ff"},LIGHTBLUE:{50:"#e1f5fe",100:"#b3e5fc",200:"#81d4fa",300:"#4fc3f7",400:"#29b6f6",500:"#03a9f4",600:"#039be5",700:"#0288d1",800:"#0277bd",900:"#01579b",A100:"#80d8ff",A200:"#40c4ff",A400:"#00b0ff",A700:"#0091ea"},CYAN:{50:"#e0f7fa",100:"#b2ebf2",200:"#80deea",300:"#4dd0e1",400:"#26c6da",500:"#00bcd4",600:"#00acc1",700:"#0097a7",800:"#00838f",900:"#006064",A100:"#84ffff",A200:"#18ffff",A400:"#00e5ff",A700:"#00b8d4"},TEAL:{50:"#e0f2f1",100:"#b2dfdb",200:"#80cbc4",300:"#4db6ac",400:"#26a69a",500:"#009688",600:"#00897b",700:"#00796b",800:"#00695c",900:"#004d40",A100:"#a7ffeb",A200:"#64ffda",A400:"#1de9b6",A700:"#00bfa5"},GREEN:{50:"#e8f5e9",100:"#c8e6c9",200:"#a5d6a7",300:"#81c784",400:"#66bb6a",500:"#4caf50",600:"#43a047",700:"#388e3c",800:"#2e7d32",900:"#1b5e20",A100:"#b9f6ca",A200:"#69f0ae",A400:"#00e676",A700:"#00c853"},LIGHTGREEN:{50:"#f1f8e9",100:"#dcedc8",200:"#c5e1a5",300:"#aed581",400:"#9ccc65",500:"#8bc34a",600:"#7cb342",700:"#689f38",800:"#558b2f",900:"#33691e",A100:"#ccff90",A200:"#b2ff59",A400:"#76ff03",A700:"#64dd17"},LIME:{50:"#f9fbe7",100:"#f0f4c3",200:"#e6ee9c",300:"#dce775",400:"#d4e157",500:"#cddc39",600:"#c0ca33",700:"#afb42b",800:"#9e9d24",900:"#827717",A100:"#f4ff81",A200:"#eeff41",A400:"#c6ff00",A700:"#aeea00"},YELLOW:{50:"#fffde7",100:"#fff9c4",200:"#fff59d",300:"#fff176",400:"#ffee58",500:"#ffeb3b",600:"#fdd835",700:"#fbc02d",800:"#f9a825",900:"#f57f17",A100:"#ffff8d",A200:"#ffff00",A400:"#ffea00",A700:"#ffd600"},AMBER:{50:"#fff8e1",100:"#ffecb3",200:"#ffe082",300:"#ffd54f",400:"#ffca28",500:"#ffc107",600:"#ffb300",700:"#ffa000",800:"#ff8f00",900:"#ff6f00",A100:"#ffe57f",A200:"#ffd740",A400:"#ffc400",A700:"#ffab00"},ORANGE:{50:"#fff3e0",100:"#ffe0b2",200:"#ffcc80",300:"#ffb74d",400:"#ffa726",500:"#ff9800",600:"#fb8c00",700:"#f57c00",800:"#ef6c00",900:"#e65100",A100:"#ffd180",A200:"#ffab40",A400:"#ff9100",A700:"#ff6d00"},DEEPORANGE:{50:"#fbe9e7",100:"#ffccbc",200:"#ffab91",300:"#ff8a65",400:"#ff7043",500:"#ff5722",600:"#f4511e",700:"#e64a19",800:"#d84315",900:"#bf360c",A100:"#ff9e80",A200:"#ff6e40",A400:"#ff3d00",A700:"#dd2c00"},BROWN:{50:"#efebe9",100:"#d7ccc8",200:"#bcaaa4",300:"#a1887f",400:"#8d6e63",500:"#795548",600:"#6d4c41",700:"#5d4037",800:"#4e342e",900:"#3e2723",A100:"#ece2df",A200:"#cfb7af",A400:"#8c6253",A700:"#533a31"},BLUEGREY:{50:"#eceff1",100:"#cfd8dc",200:"#b0bec5",300:"#90a4ae",400:"#78909c",500:"#607d8b",600:"#546e7a",700:"#455a64",800:"#37474f",900:"#263238",A100:"#f9fafb",A200:"#ccd7dc",A400:"#6e8d9b",A700:"#475c67"},GREY:{50:"#fafafa",100:"#f5f5f5",200:"#eeeeee",300:"#e0e0e0",400:"#bdbdbd",500:"#9e9e9e",600:"#757575",700:"#616161",800:"#424242",900:"#212121",A100:"#ffffff",A200:"#fcfcfc",A400:"#adadad",A700:"#7f7f7f"}},P=_.DELTABLUE,O=_.DELTAORANGE,D=_.DELTAGREEN,w=_.RED,j=_.PINK,L=_.PURPLE,T=_.DEEPPRUPLE,F=_.INDIGO,k=_.BLUE,B=_.LIGHTBLUE,M=_.CYAN,G=_.TEAL,q=_.GREEN,N=_.LIGHTGREEN,U=_.LIME,C=_.YELLOW,I=_.AMBER,H=_.ORANGE,W=_.DEEPORANGE,Y=_.BROWN,K=_.BLUEGREY,z=_.GREY,J={White:_.WHITE,Black:_.BLACK,Transparent:_.TRANSPARENT,DeltaGrey:_.DELTAGREY[500],DeltaBlue:P[500],DeltaOrange:O[500],DeltaGreen:D[500],Red:w[500],Pink:j[500],Purple:L[500],DeepPruple:T[500],Indigo:F[500],Blue:k[500],LightBlue:B[500],Cyan:M[500],Teal:G[500],Green:q[500],LightGreen:N[500],Lime:U[500],Yellow:C[500],Amber:I[500],Orange:H[500],DeepOrange:W[500],Brown:Y[500],BlueGrey:K[500],Grey:z[500]},Q=f()({},_,J),V=n(25),X=n.n(V),Z=n(41),$=n.n(Z),ee=n(42),te=n.n(ee),ne=function(){function e(){o()(this,e),this.ExecuteServiceObject=this.ExecuteServiceObject.bind(this),this.HttpCall=this.HttpCall.bind(this),this.buildHttpSvcResponse=this.buildHttpSvcResponse.bind(this)}return u()(e,[{key:"ExecuteServiceObject",value:function(e,t,n){var r=this.getMethod(e),o=t.GetRequest("http"),i=this.getURL(e,o);return this.HttpCall(i,r,o.params,o.data,o.headers,n)}},{key:"HttpCall",value:function(e,t,n,r,o){var i=arguments.length>5&&void 0!==arguments[5]?arguments[5]:null,u=this,a=new $.a(function(a,c){if(""!==t&&""!==e){"DELETE"!=t&&"GET"!=t||(r=null),o||(o={}),o[Application.Security.AuthToken]=Storage.auth;var s={method:t,url:e,data:r,headers:o,params:n,responseType:"json"};i&&(s=f()({},s,i)),console.log("Request.. ",s),te()(s).then(function(e){if(e.status<300){var t=u.buildHttpSvcResponse(Response.Success,"",e);a(t)}else c(u.buildHttpSvcResponse(Response.Failure,"",e))},function(e){c(u.buildHttpSvcResponse(Response.Failure,"",e))})}else c(u.buildHttpSvcResponse(Response.InternalError,"Could not build request",e))});return a}},{key:"createFullUrl",value:function(e,t){return null!=t&&0!=X()(t).length?e+"?"+X()(data).map(function(e){return[e,data[e]].map(encodeURIComponent).join("=")}).join("&"):e}},{key:"buildHttpSvcResponse",value:function(e,t,n){return n instanceof Error?this.buildSvcResponse(e,t,n,{}):this.buildSvcResponse(e,t,n.data,n.headers,n.status)}},{key:"buildSvcResponse",value:function(e,t,n,r,o){var i={};return i.code=e,i.message=t,i.data=n,i.info=r,i.statuscode=o,console.log(i),i}},{key:"getURL",value:function(e,t){var n=e.url;if(null!=t.urlparams)for(var r in t.urlparams)n=n.replace(":"+r,t.urlparams[r]);return n.startsWith("http")?n:Application.Backend+n}},{key:"getMethod",value:function(e){return e.method?e.method:"GET"}}]),e}();function re(e,t,n,r,o,i){console.log("****************Init uicommon*****"),_r("DataSourceHandlers","http",new ne),r&&r.EntityPrefix&&p.SetPrefix(r.EntityPrefix)}n.d(t,"Initialize",function(){return re}),n.d(t,"RequestBuilder",function(){return d}),n.d(t,"DataSource",function(){return l}),n.d(t,"Response",function(){return v}),n.d(t,"EntityData",function(){return p}),n.d(t,"Colors",function(){return Q}),n.d(t,"formatUrl",function(){return S}),n.d(t,"createAction",function(){return x}),n.d(t,"LaatooError",function(){return g}),n.d(t,"hasPermission",function(){return R})}])});
//# sourceMappingURL=index.js.map
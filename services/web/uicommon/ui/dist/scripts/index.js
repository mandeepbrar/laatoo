define("uicommon",["axios"],function(e){return function(e){var t={};function r(n){if(t[n])return t[n].exports;var o=t[n]={i:n,l:!1,exports:{}};return e[n].call(o.exports,o,o.exports,r),o.l=!0,o.exports}return r.m=e,r.c=t,r.d=function(e,t,n){r.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},r.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},r.t=function(e,t){if(1&t&&(e=r(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(r.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var o in e)r.d(n,o,function(t){return e[t]}.bind(null,o));return n},r.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return r.d(t,"a",t),t},r.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},r.p="/",r(r.s=89)}([function(e,t){var r=Object;e.exports={create:r.create,getProto:r.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:r.getOwnPropertyDescriptor,setDesc:r.defineProperty,setDescs:r.defineProperties,getKeys:r.keys,getNames:r.getOwnPropertyNames,getSymbols:r.getOwnPropertySymbols,each:[].forEach}},function(e,t,r){var n=r(32)("wks"),o=r(33),i=r(4).Symbol;e.exports=function(e){return n[e]||(n[e]=i&&i[e]||(i||o)("Symbol."+e))}},function(e,t){var r=e.exports={version:"1.2.6"};"number"==typeof __e&&(__e=r)},function(e,t,r){"use strict";t.__esModule=!0,t.default=function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}},function(e,t){var r=e.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=r)},function(e,t,r){"use strict";t.__esModule=!0;var n=function(e){return e&&e.__esModule?e:{default:e}}(r(43));t.default=function(){function e(e,t){for(var r=0;r<t.length;r++){var o=t[r];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),(0,n.default)(e,o.key,o)}}return function(t,r,n){return r&&e(t.prototype,r),n&&e(t,n),t}}()},function(e,t,r){e.exports={default:r(45),__esModule:!0}},function(e,t,r){var n=r(4),o=r(2),i=r(9),u=function(e,t,r){var a,f,c,s=e&u.F,l=e&u.G,d=e&u.S,p=e&u.P,v=e&u.B,h=e&u.W,b=l?o:o[t]||(o[t]={}),y=l?n:d?n[t]:(n[t]||{}).prototype;for(a in l&&(r=t),r)(f=!s&&y&&a in y)&&a in b||(c=f?y[a]:r[a],b[a]=l&&"function"!=typeof y[a]?r[a]:v&&f?i(c,n):h&&y[a]==c?function(e){var t=function(t){return this instanceof e?new e(t):e(t)};return t.prototype=e.prototype,t}(c):p&&"function"==typeof c?i(Function.call,c):c,p&&((b.prototype||(b.prototype={}))[a]=c))};u.F=1,u.G=2,u.S=4,u.P=8,u.B=16,u.W=32,e.exports=u},function(e,t,r){var n=r(16);e.exports=function(e){if(!n(e))throw TypeError(e+" is not an object!");return e}},function(e,t,r){var n=r(17);e.exports=function(e,t,r){if(n(e),void 0===t)return e;switch(r){case 1:return function(r){return e.call(t,r)};case 2:return function(r,n){return e.call(t,r,n)};case 3:return function(r,n,o){return e.call(t,r,n,o)}}return function(){return e.apply(t,arguments)}}},function(e,t){var r={}.toString;e.exports=function(e){return r.call(e).slice(8,-1)}},function(e,t){e.exports={}},function(e,t){e.exports=function(e){try{return!!e()}catch(e){return!0}}},function(e,t,r){e.exports=!r(12)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(e,t,r){var n=r(0).setDesc,o=r(24),i=r(1)("toStringTag");e.exports=function(e,t,r){e&&!o(e=r?e:e.prototype,i)&&n(e,i,{configurable:!0,value:t})}},function(e,t,r){var n=r(26),o=r(19);e.exports=function(e){return n(o(e))}},function(e,t){e.exports=function(e){return"object"==typeof e?null!==e:"function"==typeof e}},function(e,t){e.exports=function(e){if("function"!=typeof e)throw TypeError(e+" is not a function!");return e}},function(e,t,r){var n=r(19);e.exports=function(e){return Object(n(e))}},function(e,t){e.exports=function(e){if(void 0==e)throw TypeError("Can't call method on  "+e);return e}},function(e,t){e.exports=!0},function(e,t,r){e.exports=r(22)},function(e,t,r){var n=r(0),o=r(23);e.exports=r(13)?function(e,t,r){return n.setDesc(e,t,o(1,r))}:function(e,t,r){return e[t]=r,e}},function(e,t){e.exports=function(e,t){return{enumerable:!(1&e),configurable:!(2&e),writable:!(4&e),value:t}}},function(e,t){var r={}.hasOwnProperty;e.exports=function(e,t){return r.call(e,t)}},function(e,t,r){e.exports={default:r(69),__esModule:!0}},function(e,t,r){var n=r(10);e.exports=Object("z").propertyIsEnumerable(0)?Object:function(e){return"String"==n(e)?e.split(""):Object(e)}},function(e,t,r){var n=r(7),o=r(2),i=r(12);e.exports=function(e,t){var r=(o.Object||{})[e]||Object[e],u={};u[e]=t(r),n(n.S+n.F*i(function(){r(1)}),"Object",u)}},function(e,t,r){"use strict";t.__esModule=!0;var n=u(r(50)),o=u(r(57)),i="function"==typeof o.default&&"symbol"==typeof n.default?function(e){return typeof e}:function(e){return e&&"function"==typeof o.default&&e.constructor===o.default&&e!==o.default.prototype?"symbol":typeof e};function u(e){return e&&e.__esModule?e:{default:e}}t.default="function"==typeof o.default&&"symbol"===i(n.default)?function(e){return void 0===e?"undefined":i(e)}:function(e){return e&&"function"==typeof o.default&&e.constructor===o.default&&e!==o.default.prototype?"symbol":void 0===e?"undefined":i(e)}},function(e,t,r){"use strict";var n=r(52)(!0);r(31)(String,"String",function(e){this._t=String(e),this._i=0},function(){var e,t=this._t,r=this._i;return r>=t.length?{value:void 0,done:!0}:(e=n(t,r),this._i+=e.length,{value:e,done:!1})})},function(e,t){var r=Math.ceil,n=Math.floor;e.exports=function(e){return isNaN(e=+e)?0:(e>0?n:r)(e)}},function(e,t,r){"use strict";var n=r(20),o=r(7),i=r(21),u=r(22),a=r(24),f=r(11),c=r(53),s=r(14),l=r(0).getProto,d=r(1)("iterator"),p=!([].keys&&"next"in[].keys()),v=function(){return this};e.exports=function(e,t,r,h,b,y,m){c(r,t,h);var E,A,g=function(e){if(!p&&e in R)return R[e];switch(e){case"keys":case"values":return function(){return new r(this,e)}}return function(){return new r(this,e)}},x=t+" Iterator",S="values"==b,_=!1,R=e.prototype,P=R[d]||R["@@iterator"]||b&&R[b],O=P||g(b);if(P){var w=l(O.call(new e));s(w,x,!0),!n&&a(R,"@@iterator")&&u(w,d,v),S&&"values"!==P.name&&(_=!0,O=function(){return P.call(this)})}if(n&&!m||!p&&!_&&R[d]||u(R,d,O),f[t]=O,f[x]=v,b)if(E={values:S?O:g("values"),keys:y?O:g("keys"),entries:S?g("entries"):O},m)for(A in E)A in R||i(R,A,E[A]);else o(o.P+o.F*(p||_),t,E);return E}},function(e,t,r){var n=r(4),o=n["__core-js_shared__"]||(n["__core-js_shared__"]={});e.exports=function(e){return o[e]||(o[e]={})}},function(e,t){var r=0,n=Math.random();e.exports=function(e){return"Symbol(".concat(void 0===e?"":e,")_",(++r+n).toString(36))}},function(e,t,r){r(54);var n=r(11);n.NodeList=n.HTMLCollection=n.Array},function(e,t){},function(e,t,r){var n=r(0).getDesc,o=r(16),i=r(8),u=function(e,t){if(i(e),!o(t)&&null!==t)throw TypeError(t+": can't set as prototype!")};e.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(e,t,o){try{(o=r(9)(Function.call,n(Object.prototype,"__proto__").set,2))(e,[]),t=!(e instanceof Array)}catch(e){t=!0}return function(e,r){return u(e,r),t?e.__proto__=r:o(e,r),e}}({},!1):void 0),check:u}},function(e,t,r){var n=r(10),o=r(1)("toStringTag"),i="Arguments"==n(function(){return arguments}());e.exports=function(e){var t,r,u;return void 0===e?"Undefined":null===e?"Null":"string"==typeof(r=(t=Object(e))[o])?r:i?n(t):"Object"==(u=n(t))&&"function"==typeof t.callee?"Arguments":u}},function(e,t,r){e.exports={default:r(48),__esModule:!0}},function(e,t,r){"use strict";t.__esModule=!0;var n=function(e){return e&&e.__esModule?e:{default:e}}(r(28));t.default=function(e,t){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!t||"object"!==(void 0===t?"undefined":(0,n.default)(t))&&"function"!=typeof t?e:t}},function(e,t,r){"use strict";t.__esModule=!0;var n=u(r(64)),o=u(r(67)),i=u(r(28));function u(e){return e&&e.__esModule?e:{default:e}}t.default=function(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Super expression must either be null or a function, not "+(void 0===t?"undefined":(0,i.default)(t)));e.prototype=(0,o.default)(t&&t.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),t&&(n.default?(0,n.default)(e,t):e.__proto__=t)}},function(e,t,r){e.exports={default:r(71),__esModule:!0}},function(t,r){t.exports=e},function(e,t,r){e.exports={default:r(44),__esModule:!0}},function(e,t,r){var n=r(0);e.exports=function(e,t,r){return n.setDesc(e,t,r)}},function(e,t,r){r(46),e.exports=r(2).Object.assign},function(e,t,r){var n=r(7);n(n.S+n.F,"Object",{assign:r(47)})},function(e,t,r){var n=r(0),o=r(18),i=r(26);e.exports=r(12)(function(){var e=Object.assign,t={},r={},n=Symbol(),o="abcdefghijklmnopqrst";return t[n]=7,o.split("").forEach(function(e){r[e]=e}),7!=e({},t)[n]||Object.keys(e({},r)).join("")!=o})?function(e,t){for(var r=o(e),u=arguments,a=u.length,f=1,c=n.getKeys,s=n.getSymbols,l=n.isEnum;a>f;)for(var d,p=i(u[f++]),v=s?c(p).concat(s(p)):c(p),h=v.length,b=0;h>b;)l.call(p,d=v[b++])&&(r[d]=p[d]);return r}:Object.assign},function(e,t,r){r(49),e.exports=r(2).Object.getPrototypeOf},function(e,t,r){var n=r(18);r(27)("getPrototypeOf",function(e){return function(t){return e(n(t))}})},function(e,t,r){e.exports={default:r(51),__esModule:!0}},function(e,t,r){r(29),r(34),e.exports=r(1)("iterator")},function(e,t,r){var n=r(30),o=r(19);e.exports=function(e){return function(t,r){var i,u,a=String(o(t)),f=n(r),c=a.length;return f<0||f>=c?e?"":void 0:(i=a.charCodeAt(f))<55296||i>56319||f+1===c||(u=a.charCodeAt(f+1))<56320||u>57343?e?a.charAt(f):i:e?a.slice(f,f+2):u-56320+(i-55296<<10)+65536}}},function(e,t,r){"use strict";var n=r(0),o=r(23),i=r(14),u={};r(22)(u,r(1)("iterator"),function(){return this}),e.exports=function(e,t,r){e.prototype=n.create(u,{next:o(1,r)}),i(e,t+" Iterator")}},function(e,t,r){"use strict";var n=r(55),o=r(56),i=r(11),u=r(15);e.exports=r(31)(Array,"Array",function(e,t){this._t=u(e),this._i=0,this._k=t},function(){var e=this._t,t=this._k,r=this._i++;return!e||r>=e.length?(this._t=void 0,o(1)):o(0,"keys"==t?r:"values"==t?e[r]:[r,e[r]])},"values"),i.Arguments=i.Array,n("keys"),n("values"),n("entries")},function(e,t){e.exports=function(){}},function(e,t){e.exports=function(e,t){return{value:t,done:!!e}}},function(e,t,r){e.exports={default:r(58),__esModule:!0}},function(e,t,r){r(59),r(35),e.exports=r(2).Symbol},function(e,t,r){"use strict";var n=r(0),o=r(4),i=r(24),u=r(13),a=r(7),f=r(21),c=r(12),s=r(32),l=r(14),d=r(33),p=r(1),v=r(60),h=r(61),b=r(62),y=r(63),m=r(8),E=r(15),A=r(23),g=n.getDesc,x=n.setDesc,S=n.create,_=h.get,R=o.Symbol,P=o.JSON,O=P&&P.stringify,w=!1,D=p("_hidden"),j=n.isEnum,L=s("symbol-registry"),T=s("symbols"),F="function"==typeof R,k=Object.prototype,B=u&&c(function(){return 7!=S(x({},"a",{get:function(){return x(this,"a",{value:7}).a}})).a})?function(e,t,r){var n=g(k,t);n&&delete k[t],x(e,t,r),n&&e!==k&&x(k,t,n)}:x,M=function(e){var t=T[e]=S(R.prototype);return t._k=e,u&&w&&B(k,e,{configurable:!0,set:function(t){i(this,D)&&i(this[D],e)&&(this[D][e]=!1),B(this,e,A(1,t))}}),t},G=function(e){return"symbol"==typeof e},q=function(e,t,r){return r&&i(T,t)?(r.enumerable?(i(e,D)&&e[D][t]&&(e[D][t]=!1),r=S(r,{enumerable:A(0,!1)})):(i(e,D)||x(e,D,A(1,{})),e[D][t]=!0),B(e,t,r)):x(e,t,r)},N=function(e,t){m(e);for(var r,n=b(t=E(t)),o=0,i=n.length;i>o;)q(e,r=n[o++],t[r]);return e},U=function(e,t){return void 0===t?S(e):N(S(e),t)},C=function(e){var t=j.call(this,e);return!(t||!i(this,e)||!i(T,e)||i(this,D)&&this[D][e])||t},I=function(e,t){var r=g(e=E(e),t);return!r||!i(T,t)||i(e,D)&&e[D][t]||(r.enumerable=!0),r},H=function(e){for(var t,r=_(E(e)),n=[],o=0;r.length>o;)i(T,t=r[o++])||t==D||n.push(t);return n},W=function(e){for(var t,r=_(E(e)),n=[],o=0;r.length>o;)i(T,t=r[o++])&&n.push(T[t]);return n},Y=c(function(){var e=R();return"[null]"!=O([e])||"{}"!=O({a:e})||"{}"!=O(Object(e))});F||(f((R=function(){if(G(this))throw TypeError("Symbol is not a constructor");return M(d(arguments.length>0?arguments[0]:void 0))}).prototype,"toString",function(){return this._k}),G=function(e){return e instanceof R},n.create=U,n.isEnum=C,n.getDesc=I,n.setDesc=q,n.setDescs=N,n.getNames=h.get=H,n.getSymbols=W,u&&!r(20)&&f(k,"propertyIsEnumerable",C,!0));var K={for:function(e){return i(L,e+="")?L[e]:L[e]=R(e)},keyFor:function(e){return v(L,e)},useSetter:function(){w=!0},useSimple:function(){w=!1}};n.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(e){var t=p(e);K[e]=F?t:M(t)}),w=!0,a(a.G+a.W,{Symbol:R}),a(a.S,"Symbol",K),a(a.S+a.F*!F,"Object",{create:U,defineProperty:q,defineProperties:N,getOwnPropertyDescriptor:I,getOwnPropertyNames:H,getOwnPropertySymbols:W}),P&&a(a.S+a.F*(!F||Y),"JSON",{stringify:function(e){if(void 0!==e&&!G(e)){for(var t,r,n=[e],o=1,i=arguments;i.length>o;)n.push(i[o++]);return"function"==typeof(t=n[1])&&(r=t),!r&&y(t)||(t=function(e,t){if(r&&(t=r.call(this,e,t)),!G(t))return t}),n[1]=t,O.apply(P,n)}}}),l(R,"Symbol"),l(Math,"Math",!0),l(o.JSON,"JSON",!0)},function(e,t,r){var n=r(0),o=r(15);e.exports=function(e,t){for(var r,i=o(e),u=n.getKeys(i),a=u.length,f=0;a>f;)if(i[r=u[f++]]===t)return r}},function(e,t,r){var n=r(15),o=r(0).getNames,i={}.toString,u="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[];e.exports.get=function(e){return u&&"[object Window]"==i.call(e)?function(e){try{return o(e)}catch(e){return u.slice()}}(e):o(n(e))}},function(e,t,r){var n=r(0);e.exports=function(e){var t=n.getKeys(e),r=n.getSymbols;if(r)for(var o,i=r(e),u=n.isEnum,a=0;i.length>a;)u.call(e,o=i[a++])&&t.push(o);return t}},function(e,t,r){var n=r(10);e.exports=Array.isArray||function(e){return"Array"==n(e)}},function(e,t,r){e.exports={default:r(65),__esModule:!0}},function(e,t,r){r(66),e.exports=r(2).Object.setPrototypeOf},function(e,t,r){var n=r(7);n(n.S,"Object",{setPrototypeOf:r(36).set})},function(e,t,r){e.exports={default:r(68),__esModule:!0}},function(e,t,r){var n=r(0);e.exports=function(e,t){return n.create(e,t)}},function(e,t,r){r(70),e.exports=r(2).Object.keys},function(e,t,r){var n=r(18);r(27)("keys",function(e){return function(t){return e(n(t))}})},function(e,t,r){r(35),r(29),r(34),r(72),e.exports=r(2).Promise},function(e,t,r){"use strict";var n,o=r(0),i=r(20),u=r(4),a=r(9),f=r(37),c=r(7),s=r(16),l=r(8),d=r(17),p=r(73),v=r(74),h=r(36).set,b=r(79),y=r(1)("species"),m=r(80),E=r(81),A=u.process,g="process"==f(A),x=u.Promise,S=function(){},_=function(e){var t,r=new x(S);return e&&(r.constructor=function(e){e(S,S)}),(t=x.resolve(r)).catch(S),t===r},R=function(){var e=!1;function t(e){var r=new x(e);return h(r,t.prototype),r}try{if(e=x&&x.resolve&&_(),h(t,x),t.prototype=o.create(x.prototype,{constructor:{value:t}}),t.resolve(5).then(function(){})instanceof t||(e=!1),e&&r(13)){var n=!1;x.resolve(o.setDesc({},"then",{get:function(){n=!0}})),e=n}}catch(t){e=!1}return e}(),P=function(e){var t=l(e)[y];return void 0!=t?t:e},O=function(e){var t;return!(!s(e)||"function"!=typeof(t=e.then))&&t},w=function(e){var t,r;this.promise=new e(function(e,n){if(void 0!==t||void 0!==r)throw TypeError("Bad Promise constructor");t=e,r=n}),this.resolve=d(t),this.reject=d(r)},D=function(e){try{e()}catch(e){return{error:e}}},j=function(e,t){if(!e.n){e.n=!0;var r=e.c;E(function(){for(var n=e.v,o=1==e.s,i=0,a=function(t){var r,i,u=o?t.ok:t.fail,a=t.resolve,f=t.reject;try{u?(o||(e.h=!0),(r=!0===u?n:u(n))===t.promise?f(TypeError("Promise-chain cycle")):(i=O(r))?i.call(r,a,f):a(r)):f(n)}catch(e){f(e)}};r.length>i;)a(r[i++]);r.length=0,e.n=!1,t&&setTimeout(function(){var t,r,o=e.p;L(o)&&(g?A.emit("unhandledRejection",n,o):(t=u.onunhandledrejection)?t({promise:o,reason:n}):(r=u.console)&&r.error&&r.error("Unhandled promise rejection",n)),e.a=void 0},1)})}},L=function(e){var t,r=e._d,n=r.a||r.c,o=0;if(r.h)return!1;for(;n.length>o;)if((t=n[o++]).fail||!L(t.promise))return!1;return!0},T=function(e){var t=this;t.d||(t.d=!0,(t=t.r||t).v=e,t.s=2,t.a=t.c.slice(),j(t,!0))},F=function(e){var t,r=this;if(!r.d){r.d=!0,r=r.r||r;try{if(r.p===e)throw TypeError("Promise can't be resolved itself");(t=O(e))?E(function(){var n={r:r,d:!1};try{t.call(e,a(F,n,1),a(T,n,1))}catch(e){T.call(n,e)}}):(r.v=e,r.s=1,j(r,!1))}catch(e){T.call({r:r,d:!1},e)}}};R||(x=function(e){d(e);var t=this._d={p:p(this,x,"Promise"),c:[],a:void 0,s:0,d:!1,v:void 0,h:!1,n:!1};try{e(a(F,t,1),a(T,t,1))}catch(e){T.call(t,e)}},r(86)(x.prototype,{then:function(e,t){var r=new w(m(this,x)),n=r.promise,o=this._d;return r.ok="function"!=typeof e||e,r.fail="function"==typeof t&&t,o.c.push(r),o.a&&o.a.push(r),o.s&&j(o,!1),n},catch:function(e){return this.then(void 0,e)}})),c(c.G+c.W+c.F*!R,{Promise:x}),r(14)(x,"Promise"),r(87)("Promise"),n=r(2).Promise,c(c.S+c.F*!R,"Promise",{reject:function(e){var t=new w(this);return(0,t.reject)(e),t.promise}}),c(c.S+c.F*(!R||_(!0)),"Promise",{resolve:function(e){if(e instanceof x&&function(e,t){return!(!i||e!==x||t!==n)||b(e,t)}(e.constructor,this))return e;var t=new w(this);return(0,t.resolve)(e),t.promise}}),c(c.S+c.F*!(R&&r(88)(function(e){x.all(e).catch(function(){})})),"Promise",{all:function(e){var t=P(this),r=new w(t),n=r.resolve,i=r.reject,u=[],a=D(function(){v(e,!1,u.push,u);var r=u.length,a=Array(r);r?o.each.call(u,function(e,o){var u=!1;t.resolve(e).then(function(e){u||(u=!0,a[o]=e,--r||n(a))},i)}):n(a)});return a&&i(a.error),r.promise},race:function(e){var t=P(this),r=new w(t),n=r.reject,o=D(function(){v(e,!1,function(e){t.resolve(e).then(r.resolve,n)})});return o&&n(o.error),r.promise}})},function(e,t){e.exports=function(e,t,r){if(!(e instanceof t))throw TypeError(r+": use the 'new' operator!");return e}},function(e,t,r){var n=r(9),o=r(75),i=r(76),u=r(8),a=r(77),f=r(78);e.exports=function(e,t,r,c){var s,l,d,p=f(e),v=n(r,c,t?2:1),h=0;if("function"!=typeof p)throw TypeError(e+" is not iterable!");if(i(p))for(s=a(e.length);s>h;h++)t?v(u(l=e[h])[0],l[1]):v(e[h]);else for(d=p.call(e);!(l=d.next()).done;)o(d,v,l.value,t)}},function(e,t,r){var n=r(8);e.exports=function(e,t,r,o){try{return o?t(n(r)[0],r[1]):t(r)}catch(t){var i=e.return;throw void 0!==i&&n(i.call(e)),t}}},function(e,t,r){var n=r(11),o=r(1)("iterator"),i=Array.prototype;e.exports=function(e){return void 0!==e&&(n.Array===e||i[o]===e)}},function(e,t,r){var n=r(30),o=Math.min;e.exports=function(e){return e>0?o(n(e),9007199254740991):0}},function(e,t,r){var n=r(37),o=r(1)("iterator"),i=r(11);e.exports=r(2).getIteratorMethod=function(e){if(void 0!=e)return e[o]||e["@@iterator"]||i[n(e)]}},function(e,t){e.exports=Object.is||function(e,t){return e===t?0!==e||1/e==1/t:e!=e&&t!=t}},function(e,t,r){var n=r(8),o=r(17),i=r(1)("species");e.exports=function(e,t){var r,u=n(e).constructor;return void 0===u||void 0==(r=n(u)[i])?t:o(r)}},function(e,t,r){var n,o,i,u=r(4),a=r(82).set,f=u.MutationObserver||u.WebKitMutationObserver,c=u.process,s=u.Promise,l="process"==r(10)(c),d=function(){var e,t,r;for(l&&(e=c.domain)&&(c.domain=null,e.exit());n;)t=n.domain,r=n.fn,t&&t.enter(),r(),t&&t.exit(),n=n.next;o=void 0,e&&e.enter()};if(l)i=function(){c.nextTick(d)};else if(f){var p=1,v=document.createTextNode("");new f(d).observe(v,{characterData:!0}),i=function(){v.data=p=-p}}else i=s&&s.resolve?function(){s.resolve().then(d)}:function(){a.call(u,d)};e.exports=function(e){var t={fn:e,next:void 0,domain:l&&c.domain};o&&(o.next=t),n||(n=t,i()),o=t}},function(e,t,r){var n,o,i,u=r(9),a=r(83),f=r(84),c=r(85),s=r(4),l=s.process,d=s.setImmediate,p=s.clearImmediate,v=s.MessageChannel,h=0,b={},y=function(){var e=+this;if(b.hasOwnProperty(e)){var t=b[e];delete b[e],t()}},m=function(e){y.call(e.data)};d&&p||(d=function(e){for(var t=[],r=1;arguments.length>r;)t.push(arguments[r++]);return b[++h]=function(){a("function"==typeof e?e:Function(e),t)},n(h),h},p=function(e){delete b[e]},"process"==r(10)(l)?n=function(e){l.nextTick(u(y,e,1))}:v?(i=(o=new v).port2,o.port1.onmessage=m,n=u(i.postMessage,i,1)):s.addEventListener&&"function"==typeof postMessage&&!s.importScripts?(n=function(e){s.postMessage(e+"","*")},s.addEventListener("message",m,!1)):n="onreadystatechange"in c("script")?function(e){f.appendChild(c("script")).onreadystatechange=function(){f.removeChild(this),y.call(e)}}:function(e){setTimeout(u(y,e,1),0)}),e.exports={set:d,clear:p}},function(e,t){e.exports=function(e,t,r){var n=void 0===r;switch(t.length){case 0:return n?e():e.call(r);case 1:return n?e(t[0]):e.call(r,t[0]);case 2:return n?e(t[0],t[1]):e.call(r,t[0],t[1]);case 3:return n?e(t[0],t[1],t[2]):e.call(r,t[0],t[1],t[2]);case 4:return n?e(t[0],t[1],t[2],t[3]):e.call(r,t[0],t[1],t[2],t[3])}return e.apply(r,t)}},function(e,t,r){e.exports=r(4).document&&document.documentElement},function(e,t,r){var n=r(16),o=r(4).document,i=n(o)&&n(o.createElement);e.exports=function(e){return i?o.createElement(e):{}}},function(e,t,r){var n=r(21);e.exports=function(e,t){for(var r in t)n(e,r,t[r]);return e}},function(e,t,r){"use strict";var n=r(2),o=r(0),i=r(13),u=r(1)("species");e.exports=function(e){var t=n[e];i&&t&&!t[u]&&o.setDesc(t,u,{configurable:!0,get:function(){return this}})}},function(e,t,r){var n=r(1)("iterator"),o=!1;try{var i=[7][n]();i.return=function(){o=!0},Array.from(i,function(){throw 2})}catch(e){}e.exports=function(e,t){if(!t&&!o)return!1;var r=!1;try{var i=[7],u=i[n]();u.next=function(){return{done:r=!0}},i[n]=function(){return u},e(i)}catch(e){}return r}},function(e,t,r){"use strict";r.r(t);var n=r(3),o=r.n(n),i=r(5),u=r.n(i),a=r(6),f=r.n(a),c=function(){function e(){o()(this,e),this.ParameterSeparatorRequest=this.ParameterSeparatorRequest.bind(this),this.DefaultRequest=this.DefaultRequest.bind(this),this.URLParamsRequest=this.URLParamsRequest.bind(this)}return u()(e,[{key:"ParameterSeparatorRequest",value:function(e,t,r,n){var o={};return null==t&&(t={}),o.params=e,o.data=t,o.urlparams=r,o.headers=n,o.GetRequest=function(t){if("http"==t){var r={};if(r.data=o.data,r.headers=o.headers,null==o.params)return r.params=null,r.urlparams=null,r;var i={},u={},a=0;if(null!=o.urlparams)for(var c in o.urlparams)c in o.params&&(i[c]=o.params[c],a+=1);if(a>0){var s=0;for(var c in o.params)c in i||(u[c]=o.params[c],s+=1);return s>0?(r.urlparams=i,r.params=u):(r.urlparams=i,r.params=null),r}return r.urlparams=null,r.params=e,r}var l={};return l.data=o.data,l.params=e,l.params=f()({},e,n),l},o}},{key:"DefaultRequest",value:function(e,t,r){var n={};return null==t&&(t={}),n.params=e,n.data=t,n.headers=r,n.GetRequest=function(e){var t={};return t.data=n.data,t.params=n.params,t.urlparams=null,t.headers=n.headers,t},n}},{key:"URLParamsRequest",value:function(e,t,r){var n={};return null==t&&(t={}),n.data=t,n.urlparams=e,n.headers=r,n.GetRequest=function(e){if("http"==e){var t={};return t.data=n.data,t.params=null,t.urlparams=n.urlparams,t.headers=n.headers,t}var r={};return r.data=n.data,r.params=f()({},n.urlparams,n.headers),r},n}}]),e}(),s=function(){function e(t,r){var n=this;o()(this,e),this.SetPrefix=function(e){n.EntityPrefix=e},this.DataSource=t,this.RequestBuilder=r,this.GetEntity=this.GetEntity.bind(this),this.SaveEntity=this.SaveEntity.bind(this),this.DeleteEntity=this.DeleteEntity.bind(this),this.PutEntity=this.PutEntity.bind(this),this.UpdateEntity=this.UpdateEntity.bind(this),this.EntityPrefix="/"}return u()(e,[{key:"GetEntity",value:function(e,t,r,n){if(n){var o=this.RequestBuilder.URLParamsRequest({":id":t},null,r);return this.DataSource.ExecuteService(n,o)}var i={method:"GET"};i.url=this.EntityPrefix+e.toLowerCase()+"/"+t;o=this.RequestBuilder.DefaultRequest(null,null,r);return this.DataSource.ExecuteServiceObject(i,o)}},{key:"SaveEntity",value:function(e,t,r,n){var o=this.RequestBuilder.DefaultRequest(null,t,r);if(n)return this.DataSource.ExecuteService(n,o);var i={method:"POST"};return i.url=this.EntityPrefix+e.toLowerCase(),this.DataSource.ExecuteServiceObject(i,o)}},{key:"DeleteEntity",value:function(e,t,r,n){if(n){var o=this.RequestBuilder.URLParamsRequest({":id":t},null,r);return this.DataSource.ExecuteService(n,o)}var i={method:"DELETE"};i.url=this.EntityPrefix+e.toLowerCase()+"/"+t;o=this.RequestBuilder.DefaultRequest(null,null,r);return this.DataSource.ExecuteServiceObject(i,o)}},{key:"PutEntity",value:function(e,t,r,n,o){if(o){var i=this.RequestBuilder.URLParamsRequest({":id":t},null,n);return this.DataSource.ExecuteService(o,i)}var u={method:"PUT"};u.url=this.EntityPrefix+e.toLowerCase()+"/"+t;i=this.RequestBuilder.DefaultRequest(null,r,n);return this.DataSource.ExecuteServiceObject(u,i)}},{key:"ListEntities",value:function(e,t,r,n){if(n){var o=this.RequestBuilder.URLParamsRequest(t,null,r);return this.DataSource.ExecuteService(n,o)}var i={method:"POST"};i.url=this.EntityPrefix+e.toLowerCase()+"/view";o=this.RequestBuilder.DefaultRequest(t,null,r);return this.DataSource.ExecuteServiceObject(i,o)}},{key:"UpdateEntity",value:function(e,t,r,n,o){if(o){var i=this.RequestBuilder.URLParamsRequest({":id":t},null,n);return this.DataSource.ExecuteService(o,i)}var u={method:"PUT"};u.url=this.EntityPrefix+e.toLowerCase()+"/"+t;i=this.RequestBuilder.DefaultRequest(null,r,n);return this.DataSource.ExecuteServiceObject(u,i)}}]),e}();console.log("uicommon - application services",Application);var l=new(function(){function e(){o()(this,e),this.ExecuteService=this.ExecuteService.bind(this),this.ExecuteServiceObject=this.ExecuteServiceObject.bind(this)}return u()(e,[{key:"ExecuteService",value:function(e,t){var r=arguments.length>2&&void 0!==arguments[2]?arguments[2]:null,n=_reg("Services",e);if(null!=n&&null!=t)return this.ExecuteServiceObject(n,t,r);throw new Error("Service not found "+e)}},{key:"ExecuteServiceObject",value:function(e,t){var r=arguments.length>2&&void 0!==arguments[2]?arguments[2]:null;if(e.protocol||(e.protocol="http"),null!=e&&null!=t){var n=_reg("DataSourceHandlers",e.protocol);if(null==n)throw console.log("Requested service for handler",e),new Error("Invalid protocol handler");return n.ExecuteServiceObject(e,t,r)}}}]),e}()),d=new c,p=new s(l,d),v={Success:"Success",Unauthorized:"Unauthorized",InternalError:"InternalError",BadRequest:"BadRequest",Failure:"Failure"},h=r(38),b=r.n(h),y=r(39),m=r.n(y),E=r(40),A=r.n(E),g=function(e){function t(e,r,n){o()(this,t);var i=m()(this,(t.__proto__||b()(t)).call(this,e));return i.name=i.constructor.name,i.message=e,"function"==typeof Error.captureStackTrace?Error.captureStackTrace(i,i.constructor):i.stack=new Error(e).stack,i.type=e,i.rootError=r,i.args=n,i}return A()(t,e),t}(Error);function x(e,t,r){var n=t instanceof Error;return console.log("created action",e,t,r,n),{type:e,payload:t,meta:r,error:n}}function S(e,t){var r=e;if(t)for(var n in t){var o=t[n];r=r.replace(new RegExp(":"+n,"g"),o)}return r}function _(e){var t=!0;if(e&&""!=e){var r=localStorage.permissions;r&&r.indexOf(e)<0&&(t=!1)}return t}var R={WHITE:"#ffffff",BLACK:"#000000",TRANSPARENT:"transparent",DELTAGREY:{50:"#e4e4e4",100:"#bdbdbd",200:"#a1a1a1",300:"#7e7e7e",400:"#6e6e6e",500:"#5f5f5f",600:"#505050",700:"#404040",800:"#313131",900:"#222222",A100:"#e4e4e4",A200:"#bdbdbd",A400:"#6e6e6e",A700:"#404040"},DELTABLUE:{50:"#e6f5ff",100:"#9ad8ff",200:"#62c2ff",300:"#1aa7ff",400:"#009afb",500:"#0087dc",600:"#0074bd",700:"#00619f",800:"#004f80",900:"#003c62",A100:"#e6f5ff",A200:"#9ad8ff",A400:"#009afb",A700:"#00619f"},DELTAORANGE:{50:"#FFFDFA",100:"#FFDAAE",200:"#FFC076",300:"#FF9F2E",400:"#FF9110",500:"#F08200",600:"#D17100",700:"#B36100",800:"#945000",900:"#764000",A100:"#FFFDFA",A200:"#FFDAAE",A400:"#FF9110",A700:"#B36100"},DELTAGREEN:{50:"#C5F8CD",100:"#81EF91",200:"#50E965",300:"#1BD636",400:"#17BB2F",500:"#14A028",600:"#118521",700:"#0D6A1A",800:"#0A4E14",900:"#06330D",A100:"#C5F8CD",A200:"#81EF91",A400:"#17BB2F",A700:"#0D6A1A"},RED:{50:"ffebee",100:"#ffcdd2",200:"#ef9a9a",300:"#e57373",400:"#ef5350",500:"#f44336",600:"#e53935",700:"#d32f2f",800:"#c62828",900:"#b71c1c",A100:"#ff8a80",A200:"#ff5252",A400:"#ff1744",A700:"#d50000"},PINK:{50:"#fce4ec",100:"#f8bbd0",200:"#f48fb1",300:"#f06292",400:"#ec407a",500:"#e91e63",600:"#d81b60",700:"#c2185b",800:"#ad1457",900:"#880e4f",A100:"#ff80ab",A200:"#ff4081",A400:"#f50057",A700:"#c51162"},PURPLE:{50:"#f3e5f5",100:"#e1bee7",200:"#ce93d8",300:"#ba68c8",400:"#ab47bc",500:"#9c27b0",600:"#8e24aa",700:"#7b1fa2",800:"#6a1b9a",900:"#4a148c",A100:"#ea80fc",A200:"#e040fb",A400:"#d500f9",A700:"#aa00ff"},DEEPPRUPLE:{50:"#ede7f6",100:"#d1c4e9",200:"#b39ddb",300:"#9575cd",400:"#7e57c2",500:"#673ab7",600:"#5e35b1",700:"#512DA8",800:"#4527A0",900:"#311B92",A100:"#b388ff",A200:"#7c4dff",A400:"#651fff",A700:"#6200ea"},INDIGO:{50:"#e8eaf6",100:"#c5cae9",200:"#9fa8da",300:"#7986cb",400:"#5c6bc0",500:"#3f51b5",600:"#3949ab",700:"#303F9F",800:"#283593",900:"#1A237E",A100:"#8c9eff",A200:"#536dfe",A400:"#3d5afe",A700:"#304ffe"},BLUE:{50:"#e3f2fd",100:"#bbdefb",200:"#90caf9",300:"#64b5f6",400:"#42a5f5",500:"#2196f3",600:"#1e88e5",700:"#1976d2",800:"#1565c0",900:"#0d47a1",A100:"#82b1ff",A200:"#448aff",A400:"#2979ff",A700:"#2962ff"},LIGHTBLUE:{50:"#e1f5fe",100:"#b3e5fc",200:"#81d4fa",300:"#4fc3f7",400:"#29b6f6",500:"#03a9f4",600:"#039be5",700:"#0288d1",800:"#0277bd",900:"#01579b",A100:"#80d8ff",A200:"#40c4ff",A400:"#00b0ff",A700:"#0091ea"},CYAN:{50:"#e0f7fa",100:"#b2ebf2",200:"#80deea",300:"#4dd0e1",400:"#26c6da",500:"#00bcd4",600:"#00acc1",700:"#0097a7",800:"#00838f",900:"#006064",A100:"#84ffff",A200:"#18ffff",A400:"#00e5ff",A700:"#00b8d4"},TEAL:{50:"#e0f2f1",100:"#b2dfdb",200:"#80cbc4",300:"#4db6ac",400:"#26a69a",500:"#009688",600:"#00897b",700:"#00796b",800:"#00695c",900:"#004d40",A100:"#a7ffeb",A200:"#64ffda",A400:"#1de9b6",A700:"#00bfa5"},GREEN:{50:"#e8f5e9",100:"#c8e6c9",200:"#a5d6a7",300:"#81c784",400:"#66bb6a",500:"#4caf50",600:"#43a047",700:"#388e3c",800:"#2e7d32",900:"#1b5e20",A100:"#b9f6ca",A200:"#69f0ae",A400:"#00e676",A700:"#00c853"},LIGHTGREEN:{50:"#f1f8e9",100:"#dcedc8",200:"#c5e1a5",300:"#aed581",400:"#9ccc65",500:"#8bc34a",600:"#7cb342",700:"#689f38",800:"#558b2f",900:"#33691e",A100:"#ccff90",A200:"#b2ff59",A400:"#76ff03",A700:"#64dd17"},LIME:{50:"#f9fbe7",100:"#f0f4c3",200:"#e6ee9c",300:"#dce775",400:"#d4e157",500:"#cddc39",600:"#c0ca33",700:"#afb42b",800:"#9e9d24",900:"#827717",A100:"#f4ff81",A200:"#eeff41",A400:"#c6ff00",A700:"#aeea00"},YELLOW:{50:"#fffde7",100:"#fff9c4",200:"#fff59d",300:"#fff176",400:"#ffee58",500:"#ffeb3b",600:"#fdd835",700:"#fbc02d",800:"#f9a825",900:"#f57f17",A100:"#ffff8d",A200:"#ffff00",A400:"#ffea00",A700:"#ffd600"},AMBER:{50:"#fff8e1",100:"#ffecb3",200:"#ffe082",300:"#ffd54f",400:"#ffca28",500:"#ffc107",600:"#ffb300",700:"#ffa000",800:"#ff8f00",900:"#ff6f00",A100:"#ffe57f",A200:"#ffd740",A400:"#ffc400",A700:"#ffab00"},ORANGE:{50:"#fff3e0",100:"#ffe0b2",200:"#ffcc80",300:"#ffb74d",400:"#ffa726",500:"#ff9800",600:"#fb8c00",700:"#f57c00",800:"#ef6c00",900:"#e65100",A100:"#ffd180",A200:"#ffab40",A400:"#ff9100",A700:"#ff6d00"},DEEPORANGE:{50:"#fbe9e7",100:"#ffccbc",200:"#ffab91",300:"#ff8a65",400:"#ff7043",500:"#ff5722",600:"#f4511e",700:"#e64a19",800:"#d84315",900:"#bf360c",A100:"#ff9e80",A200:"#ff6e40",A400:"#ff3d00",A700:"#dd2c00"},BROWN:{50:"#efebe9",100:"#d7ccc8",200:"#bcaaa4",300:"#a1887f",400:"#8d6e63",500:"#795548",600:"#6d4c41",700:"#5d4037",800:"#4e342e",900:"#3e2723",A100:"#ece2df",A200:"#cfb7af",A400:"#8c6253",A700:"#533a31"},BLUEGREY:{50:"#eceff1",100:"#cfd8dc",200:"#b0bec5",300:"#90a4ae",400:"#78909c",500:"#607d8b",600:"#546e7a",700:"#455a64",800:"#37474f",900:"#263238",A100:"#f9fafb",A200:"#ccd7dc",A400:"#6e8d9b",A700:"#475c67"},GREY:{50:"#fafafa",100:"#f5f5f5",200:"#eeeeee",300:"#e0e0e0",400:"#bdbdbd",500:"#9e9e9e",600:"#757575",700:"#616161",800:"#424242",900:"#212121",A100:"#ffffff",A200:"#fcfcfc",A400:"#adadad",A700:"#7f7f7f"}},P=R.DELTABLUE,O=R.DELTAORANGE,w=R.DELTAGREEN,D=R.RED,j=R.PINK,L=R.PURPLE,T=R.DEEPPRUPLE,F=R.INDIGO,k=R.BLUE,B=R.LIGHTBLUE,M=R.CYAN,G=R.TEAL,q=R.GREEN,N=R.LIGHTGREEN,U=R.LIME,C=R.YELLOW,I=R.AMBER,H=R.ORANGE,W=R.DEEPORANGE,Y=R.BROWN,K=R.BLUEGREY,z=R.GREY,J={White:R.WHITE,Black:R.BLACK,Transparent:R.TRANSPARENT,DeltaGrey:R.DELTAGREY[500],DeltaBlue:P[500],DeltaOrange:O[500],DeltaGreen:w[500],Red:D[500],Pink:j[500],Purple:L[500],DeepPruple:T[500],Indigo:F[500],Blue:k[500],LightBlue:B[500],Cyan:M[500],Teal:G[500],Green:q[500],LightGreen:N[500],Lime:U[500],Yellow:C[500],Amber:I[500],Orange:H[500],DeepOrange:W[500],Brown:Y[500],BlueGrey:K[500],Grey:z[500]},$=(f()({},R,J),r(25)),Q=r.n($),V=r(41),X=r.n(V),Z=r(42),ee=r.n(Z),te=function(){function e(){o()(this,e),this.ExecuteServiceObject=this.ExecuteServiceObject.bind(this),this.HttpCall=this.HttpCall.bind(this),this.buildHttpSvcResponse=this.buildHttpSvcResponse.bind(this);var t=_$.laatoobrowser_wasm;null!=t&&(this.Application=t.initialize())}return u()(e,[{key:"ExecuteServiceObject",value:function(e,t,r){console.log("********Laatoo browser",this.Application,this.Application.js_get_registered_item("sdfdf","sdf"));var n=this.getMethod(e),o=t.GetRequest("http"),i=this.getURL(e,o);return this.HttpCall(i,n,o.params,o.data,o.headers,r)}},{key:"HttpCall",value:function(e,t,r,n,o){var i=arguments.length>5&&void 0!==arguments[5]?arguments[5]:null;console.log("http call...",this.Browser,this.Browser.execute_service),this.Browser.execute_service(e,t,r);var u=this;return new X.a(function(a,c){if(""!==t&&""!==e){"DELETE"!=t&&"GET"!=t||(n=null),o||(o={}),o[Application.Security.AuthToken]=Storage.auth;var s={method:t,url:e,data:n,headers:o,params:r,responseType:"json"};i&&(s=f()({},s,i)),console.log("Request.. ",s),ee()(s).then(function(e){if(e.status<300){var t=u.buildHttpSvcResponse(Response.Success,"",e);a(t)}else c(u.buildHttpSvcResponse(Response.Failure,"",e))},function(e){c(u.buildHttpSvcResponse(Response.Failure,"",e))})}else c(u.buildHttpSvcResponse(Response.InternalError,"Could not build request",e))})}},{key:"createFullUrl",value:function(e,t){return null!=t&&0!=Q()(t).length?e+"?"+Q()(data).map(function(e){return[e,data[e]].map(encodeURIComponent).join("=")}).join("&"):e}},{key:"buildHttpSvcResponse",value:function(e,t,r){return r instanceof Error?this.buildSvcResponse(e,t,r,{}):this.buildSvcResponse(e,t,r.data,r.headers,r.status)}},{key:"buildSvcResponse",value:function(e,t,r,n,o){var i={};return i.code=e,i.message=t,i.data=r,i.info=n,i.statuscode=o,console.log(i),i}},{key:"getURL",value:function(e,t){var r=e.url;if(null!=t.urlparams)for(var n in t.urlparams)r=r.replace(":"+n,t.urlparams[n]);return r.startsWith("http")?r:Application.Backend+r}},{key:"getMethod",value:function(e){return e.method?e.method:"GET"}}]),e}();function re(e,t,r,n,o,i){console.log("****************Init uicommon*****"),_r("DataSourceHandlers","http",new te),n&&n.EntityPrefix&&p.SetPrefix(n.EntityPrefix)}r.d(t,"Initialize",function(){return re}),r.d(t,"Colors",function(){return Colors}),r.d(t,"RequestBuilder",function(){return d}),r.d(t,"DataSource",function(){return l}),r.d(t,"Response",function(){return v}),r.d(t,"EntityData",function(){return p}),r.d(t,"formatUrl",function(){return S}),r.d(t,"createAction",function(){return x}),r.d(t,"LaatooError",function(){return g}),r.d(t,"hasPermission",function(){return _})}])});
//# sourceMappingURL=index.js.map
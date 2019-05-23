define("entity",["react","prop-types"],function(t,e){return function(t){var e={};function n(r){if(e[r])return e[r].exports;var o=e[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=t,n.c=e,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},n.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var o in t)n.d(r,o,function(e){return t[e]}.bind(null,o));return r},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=33)}([function(e,n){e.exports=t},function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,e,n){var r=n(23)("wks"),o=n(24),u=n(5).Symbol;t.exports=function(t){return r[t]||(r[t]=u&&u[t]||(u||o)("Symbol."+t))}},function(t,e,n){var r=n(5),o=n(2),u=n(18),i=function(t,e,n){var c,f,a,s=t&i.F,l=t&i.G,p=t&i.S,d=t&i.P,y=t&i.B,v=t&i.W,b=l?o:o[e]||(o[e]={}),h=l?r:p?r[e]:(r[e]||{}).prototype;for(c in l&&(n=e),n)(f=!s&&h&&c in h)&&c in b||(a=f?h[c]:n[c],b[c]=l&&"function"!=typeof h[c]?n[c]:y&&f?u(a,r):v&&h[c]==a?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(a):d&&"function"==typeof a?u(Function.call,a):a,d&&((b.prototype||(b.prototype={}))[c]=a))};i.F=1,i.G=2,i.S=4,i.P=8,i.B=16,i.W=32,t.exports=i},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e,n){var r=n(48),o=n(8);t.exports=function(t){return r(o(t))}},function(t,e,n){"use strict";e.__esModule=!0;var r=i(n(38)),o=i(n(49)),u="function"==typeof o.default&&"symbol"==typeof r.default?function(t){return typeof t}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":typeof t};function i(t){return t&&t.__esModule?t:{default:t}}e.default="function"==typeof o.default&&"symbol"===u(r.default)?function(t){return void 0===t?"undefined":u(t)}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":void 0===t?"undefined":u(t)}},function(t,e){t.exports=function(t){if(null==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e,n){var r=n(1),o=n(11);t.exports=n(22)?function(t,e,n){return r.setDesc(t,e,o(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports={}},function(t,e,n){var r=n(1).setDesc,o=n(12),u=n(3)("toStringTag");t.exports=function(t,e,n){t&&!o(t=n?t:t.prototype,u)&&r(t,u,{configurable:!0,value:e})}},function(t,e,n){"use strict";e.__esModule=!0;var r,o=n(7),u=(r=o)&&r.__esModule?r:{default:r};e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,u.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){var r=n(8);t.exports=function(t){return Object(r(t))}},function(t,e,n){var r=n(4),o=n(2),u=n(9);t.exports=function(t,e){var n=(o.Object||{})[t]||Object[t],i={};i[t]=e(n),r(r.S+r.F*u(function(){n(1)}),"Object",i)}},function(t,e,n){var r=n(37);t.exports=function(t,e,n){if(r(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,r){return t.call(e,n,r)};case 3:return function(n,r,o){return t.call(e,n,r,o)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){"use strict";var r=n(20),o=n(4),u=n(21),i=n(10),c=n(12),f=n(13),a=n(43),s=n(14),l=n(1).getProto,p=n(3)("iterator"),d=!([].keys&&"next"in[].keys()),y=function(){return this};t.exports=function(t,e,n,v,b,h,_){a(n,e,v);var g,m,x=function(t){if(!d&&t in S)return S[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},O=e+" Iterator",w="values"==b,j=!1,S=t.prototype,P=S[p]||S["@@iterator"]||b&&S[b],M=P||x(b);if(P){var E=l(M.call(new t));s(E,O,!0),!r&&c(S,"@@iterator")&&i(E,p,y),w&&"values"!==P.name&&(j=!0,M=function(){return P.call(this)})}if(r&&!_||!d&&!j&&S[p]||i(S,p,M),f[e]=M,f[O]=y,b)if(g={values:w?M:x("values"),keys:h?M:x("keys"),entries:w?x("entries"):M},_)for(m in g)m in S||u(S,m,g[m]);else o(o.P+o.F*(d||j),e,g);return g}},function(t,e){t.exports=!0},function(t,e,n){t.exports=n(10)},function(t,e,n){t.exports=!n(9)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){var r=n(5),o=r["__core-js_shared__"]||(r["__core-js_shared__"]={});t.exports=function(t){return o[t]||(o[t]={})}},function(t,e){var n=0,r=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+r).toString(36))}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){var r=n(27);t.exports=function(t){if(!r(t))throw TypeError(t+" is not an object!");return t}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){t.exports={default:n(35),__esModule:!0}},function(t,e,n){t.exports={default:n(57),__esModule:!0}},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){"use strict";e.__esModule=!0;var r,o=n(59),u=(r=o)&&r.__esModule?r:{default:r};e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var r=e[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),(0,u.default)(t,r.key,r)}}return function(e,n,r){return n&&t(e.prototype,n),r&&t(e,r),e}}()},function(t,e,n){"use strict";e.__esModule=!0;var r=i(n(61)),o=i(n(65)),u=i(n(7));function i(t){return t&&t.__esModule?t:{default:t}}e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,u.default)(e)));t.prototype=(0,o.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(r.default?(0,r.default)(t,e):t.__proto__=e)}},function(t,e,n){"use strict";n.r(e),function(t){n.d(e,"Initialize",function(){return g});var r=n(28),o=n.n(r),u=n(7),i=n.n(u),c=n(29),f=n.n(c),a=n(30),s=n.n(a),l=n(31),p=n.n(l),d=n(15),y=n.n(d),v=n(32),b=n.n(v),h=(n(67),n(0)),_=n.n(h);n(68);function g(e,n,r,o,u,i){if(t.properties=Application.Properties[n],t.settings=o,console.log("entity initialize",o,o.object),o.object){var c=o.object+"_default";_reg("Blocks",c)||_r("Blocks",c,function(t,e,n){return console.log("rendering default display",c,t,e,n),_.a.createElement("h1",null,"default display")})}}Window.displayDefaultEntity=function(t,e,n){return _.a.createElement(m,{desc:e,uikit:n,ctx:t})};var m=function(t){function e(){var t,n,r,u;s()(this,e);for(var c=arguments.length,a=Array(c),l=0;l<c;l++)a[l]=arguments[l];return n=r=y()(this,(t=e.__proto__||f()(e)).call.apply(t,[this].concat(a))),r.createField=function(t,e,n,o,u,i){var c=r.createObjFields(t,n+1,o,u,i);return _.a.createElement("div",{className:"field "+e},_.a.createElement("div",{className:"name"},e),_.a.createElement("div",{className:"value"},c))},r.createObjFields=function(t,e,n,u,c){if(null==t)return null;if(t instanceof Array){for(var f=new Array,a=0;a<t.length;a++)f.push(_.a.createElement("div",{className:"entityarrayitem"},r.createObjFields(t[a],e+1,n,u,c)));return f}if("object"==(void 0===t?"undefined":i()(t))){var s=new Array,l=new Array,p=r;return o()(t).forEach(function(r){var o=t[r],i=p.createField(o,r,e,n,u,c);console.log("field",r,"fieldVal",o," level ",e),o instanceof Array&&0==e?l.push(_.a.createElement(c.Tab,{label:r},i)):s.push(i)}),0!=e?s:_.a.createElement(c.Tabset,null,_.a.createElement(c.Tab,{label:"General"},s),l)}return t},u=n,y()(r,u)}return b()(e,t),p()(e,[{key:"render",value:function(){var t=this.props,e=t.ctx,n=t.desc,r=t.uikit;return console.log(e,n,r),_.a.createElement("div",{className:"entity "},this.createObjFields(e.data,0,e,n,r))}}]),e}(_.a.Component)}.call(this,n(34)(t))},function(t,e){t.exports=function(t){if(!t.webpackPolyfill){var e=Object.create(t);e.children||(e.children=[]),Object.defineProperty(e,"loaded",{enumerable:!0,get:function(){return e.l}}),Object.defineProperty(e,"id",{enumerable:!0,get:function(){return e.i}}),Object.defineProperty(e,"exports",{enumerable:!0}),e.webpackPolyfill=1}return e}},function(t,e,n){n(36),t.exports=n(2).Object.keys},function(t,e,n){var r=n(16);n(17)("keys",function(t){return function(e){return t(r(e))}})},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){t.exports={default:n(39),__esModule:!0}},function(t,e,n){n(40),n(44),t.exports=n(3)("iterator")},function(t,e,n){"use strict";var r=n(41)(!0);n(19)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=r(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){var r=n(42),o=n(8);t.exports=function(t){return function(e,n){var u,i,c=String(o(e)),f=r(n),a=c.length;return f<0||f>=a?t?"":void 0:(u=c.charCodeAt(f))<55296||u>56319||f+1===a||(i=c.charCodeAt(f+1))<56320||i>57343?t?c.charAt(f):u:t?c.slice(f,f+2):i-56320+(u-55296<<10)+65536}}},function(t,e){var n=Math.ceil,r=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?r:n)(t)}},function(t,e,n){"use strict";var r=n(1),o=n(11),u=n(14),i={};n(10)(i,n(3)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=r.create(i,{next:o(1,n)}),u(t,e+" Iterator")}},function(t,e,n){n(45);var r=n(13);r.NodeList=r.HTMLCollection=r.Array},function(t,e,n){"use strict";var r=n(46),o=n(47),u=n(13),i=n(6);t.exports=n(19)(Array,"Array",function(t,e){this._t=i(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,o(1)):o(0,"keys"==e?n:"values"==e?t[n]:[n,t[n]])},"values"),u.Arguments=u.Array,r("keys"),r("values"),r("entries")},function(t,e){t.exports=function(){}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e,n){var r=n(25);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==r(t)?t.split(""):Object(t)}},function(t,e,n){t.exports={default:n(50),__esModule:!0}},function(t,e,n){n(51),n(56),t.exports=n(2).Symbol},function(t,e,n){"use strict";var r=n(1),o=n(5),u=n(12),i=n(22),c=n(4),f=n(21),a=n(9),s=n(23),l=n(14),p=n(24),d=n(3),y=n(52),v=n(53),b=n(54),h=n(55),_=n(26),g=n(6),m=n(11),x=r.getDesc,O=r.setDesc,w=r.create,j=v.get,S=o.Symbol,P=o.JSON,M=P&&P.stringify,E=!1,k=d("_hidden"),A=r.isEnum,N=s("symbol-registry"),F=s("symbols"),D="function"==typeof S,T=Object.prototype,C=i&&a(function(){return 7!=w(O({},"a",{get:function(){return O(this,"a",{value:7}).a}})).a})?function(t,e,n){var r=x(T,e);r&&delete T[e],O(t,e,n),r&&t!==T&&O(T,e,r)}:O,I=function(t){var e=F[t]=w(S.prototype);return e._k=t,i&&E&&C(T,t,{configurable:!0,set:function(e){u(this,k)&&u(this[k],t)&&(this[k][t]=!1),C(this,t,m(1,e))}}),e},W=function(t){return"symbol"==typeof t},B=function(t,e,n){return n&&u(F,e)?(n.enumerable?(u(t,k)&&t[k][e]&&(t[k][e]=!1),n=w(n,{enumerable:m(0,!1)})):(u(t,k)||O(t,k,m(1,{})),t[k][e]=!0),C(t,e,n)):O(t,e,n)},G=function(t,e){_(t);for(var n,r=b(e=g(e)),o=0,u=r.length;u>o;)B(t,n=r[o++],e[n]);return t},J=function(t,e){return void 0===e?w(t):G(w(t),e)},z=function(t){var e=A.call(this,t);return!(e||!u(this,t)||!u(F,t)||u(this,k)&&this[k][t])||e},K=function(t,e){var n=x(t=g(t),e);return!n||!u(F,e)||u(t,k)&&t[k][e]||(n.enumerable=!0),n},L=function(t){for(var e,n=j(g(t)),r=[],o=0;n.length>o;)u(F,e=n[o++])||e==k||r.push(e);return r},H=function(t){for(var e,n=j(g(t)),r=[],o=0;n.length>o;)u(F,e=n[o++])&&r.push(F[e]);return r},R=a(function(){var t=S();return"[null]"!=M([t])||"{}"!=M({a:t})||"{}"!=M(Object(t))});D||(f((S=function(){if(W(this))throw TypeError("Symbol is not a constructor");return I(p(arguments.length>0?arguments[0]:void 0))}).prototype,"toString",function(){return this._k}),W=function(t){return t instanceof S},r.create=J,r.isEnum=z,r.getDesc=K,r.setDesc=B,r.setDescs=G,r.getNames=v.get=L,r.getSymbols=H,i&&!n(20)&&f(T,"propertyIsEnumerable",z,!0));var V={for:function(t){return u(N,t+="")?N[t]:N[t]=S(t)},keyFor:function(t){return y(N,t)},useSetter:function(){E=!0},useSimple:function(){E=!1}};r.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=d(t);V[t]=D?e:I(e)}),E=!0,c(c.G+c.W,{Symbol:S}),c(c.S,"Symbol",V),c(c.S+c.F*!D,"Object",{create:J,defineProperty:B,defineProperties:G,getOwnPropertyDescriptor:K,getOwnPropertyNames:L,getOwnPropertySymbols:H}),P&&c(c.S+c.F*(!D||R),"JSON",{stringify:function(t){if(void 0!==t&&!W(t)){for(var e,n,r=[t],o=1,u=arguments;u.length>o;)r.push(u[o++]);return"function"==typeof(e=r[1])&&(n=e),!n&&h(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!W(e))return e}),r[1]=e,M.apply(P,r)}}}),l(S,"Symbol"),l(Math,"Math",!0),l(o.JSON,"JSON",!0)},function(t,e,n){var r=n(1),o=n(6);t.exports=function(t,e){for(var n,u=o(t),i=r.getKeys(u),c=i.length,f=0;c>f;)if(u[n=i[f++]]===e)return n}},function(t,e,n){var r=n(6),o=n(1).getNames,u={}.toString,i="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[];t.exports.get=function(t){return i&&"[object Window]"==u.call(t)?function(t){try{return o(t)}catch(t){return i.slice()}}(t):o(r(t))}},function(t,e,n){var r=n(1);t.exports=function(t){var e=r.getKeys(t),n=r.getSymbols;if(n)for(var o,u=n(t),i=r.isEnum,c=0;u.length>c;)i.call(t,o=u[c++])&&e.push(o);return e}},function(t,e,n){var r=n(25);t.exports=Array.isArray||function(t){return"Array"==r(t)}},function(t,e){},function(t,e,n){n(58),t.exports=n(2).Object.getPrototypeOf},function(t,e,n){var r=n(16);n(17)("getPrototypeOf",function(t){return function(e){return t(r(e))}})},function(t,e,n){t.exports={default:n(60),__esModule:!0}},function(t,e,n){var r=n(1);t.exports=function(t,e,n){return r.setDesc(t,e,n)}},function(t,e,n){t.exports={default:n(62),__esModule:!0}},function(t,e,n){n(63),t.exports=n(2).Object.setPrototypeOf},function(t,e,n){var r=n(4);r(r.S,"Object",{setPrototypeOf:n(64).set})},function(t,e,n){var r=n(1).getDesc,o=n(27),u=n(26),i=function(t,e){if(u(t),!o(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,o){try{(o=n(18)(Function.call,r(Object.prototype,"__proto__").set,2))(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return i(t,n),e?t.__proto__=n:o(t,n),t}}({},!1):void 0),check:i}},function(t,e,n){t.exports={default:n(66),__esModule:!0}},function(t,e,n){var r=n(1);t.exports=function(t,e){return r.create(t,e)}},function(t,e){},function(t,n){t.exports=e}])});
//# sourceMappingURL=index.js.map
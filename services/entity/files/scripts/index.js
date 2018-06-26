define("entity",["react","prop-types"],function(t,e){return function(t){var e={};function n(r){if(e[r])return e[r].exports;var o=e[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=t,n.c=e,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},n.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var o in t)n.d(r,o,function(e){return t[e]}.bind(null,o));return r},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=92)}([function(e,n){e.exports=t},function(t,e){var n=t.exports={version:"2.5.7"};"number"==typeof __e&&(__e=n)},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e,n){t.exports=!n(11)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){var r=n(12),o=n(39),i=n(24),u=Object.defineProperty;e.f=n(3)?Object.defineProperty:function(t,e,n){if(r(t),e=i(e,!0),r(n),o)try{return u(t,e,n)}catch(t){}if("get"in n||"set"in n)throw TypeError("Accessors not supported!");return"value"in n&&(t[e]=n.value),t}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e,n){var r=n(26)("wks"),o=n(14),i=n(2).Symbol,u="function"==typeof i;(t.exports=function(t){return r[t]||(r[t]=u&&i[t]||(u?i:o)("Symbol."+t))}).store=r},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){var r=n(4),o=n(13);t.exports=n(3)?function(t,e,n){return r.f(t,e,o(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e,n){var r=n(2),o=n(1),i=n(40),u=n(8),c=n(5),f=function(t,e,n){var a,s,l,p=t&f.F,y=t&f.G,d=t&f.S,v=t&f.P,b=t&f.B,h=t&f.W,m=y?o:o[e]||(o[e]={}),_=m.prototype,g=y?r:d?r[e]:(r[e]||{}).prototype;for(a in y&&(n=e),n)(s=!p&&g&&void 0!==g[a])&&c(m,a)||(l=s?g[a]:n[a],m[a]=y&&"function"!=typeof g[a]?n[a]:b&&s?i(l,r):h&&g[a]==l?function(t){var e=function(e,n,r){if(this instanceof t){switch(arguments.length){case 0:return new t;case 1:return new t(e);case 2:return new t(e,n)}return new t(e,n,r)}return t.apply(this,arguments)};return e.prototype=t.prototype,e}(l):v&&"function"==typeof l?i(Function.call,l):l,v&&((m.virtual||(m.virtual={}))[a]=l,t&f.R&&_&&!_[a]&&u(_,a,l)))};f.F=1,f.G=2,f.S=4,f.P=8,f.B=16,f.W=32,f.U=64,f.R=128,t.exports=f},function(t,e,n){var r=n(88),o=n(29);t.exports=function(t){return r(o(t))}},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e,n){var r=n(7);t.exports=function(t){if(!r(t))throw TypeError(t+" is not an object!");return t}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e){var n=0,r=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+r).toString(36))}},function(t,e){t.exports=!0},function(t,e,n){var r=n(43),o=n(25);t.exports=Object.keys||function(t){return r(t,o)}},function(t,e,n){"use strict";e.__esModule=!0;var r=u(n(83)),o=u(n(72)),i="function"==typeof o.default&&"symbol"==typeof r.default?function(t){return typeof t}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":typeof t};function u(t){return t&&t.__esModule?t:{default:t}}e.default="function"==typeof o.default&&"symbol"===i(r.default)?function(t){return void 0===t?"undefined":i(t)}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":void 0===t?"undefined":i(t)}},function(t,e){e.f={}.propertyIsEnumerable},function(t,e,n){var r=n(2),o=n(1),i=n(15),u=n(20),c=n(4).f;t.exports=function(t){var e=o.Symbol||(o.Symbol=i?{}:r.Symbol||{});"_"==t.charAt(0)||t in e||c(e,t,{value:u.f(t)})}},function(t,e,n){e.f=n(6)},function(t,e,n){var r=n(4).f,o=n(5),i=n(6)("toStringTag");t.exports=function(t,e,n){t&&!o(t=n?t:t.prototype,i)&&r(t,i,{configurable:!0,value:e})}},function(t,e,n){var r=n(12),o=n(78),i=n(25),u=n(27)("IE_PROTO"),c=function(){},f=function(){var t,e=n(38)("iframe"),r=i.length;for(e.style.display="none",n(77).appendChild(e),e.src="javascript:",(t=e.contentWindow.document).open(),t.write("<script>document.F=Object<\/script>"),t.close(),f=t.F;r--;)delete f.prototype[i[r]];return f()};t.exports=Object.create||function(t,e){var n;return null!==t?(c.prototype=r(t),n=new c,c.prototype=null,n[u]=t):n=f(),void 0===e?n:o(n,e)}},function(t,e){t.exports={}},function(t,e,n){var r=n(7);t.exports=function(t,e){if(!r(t))return t;var n,o;if(e&&"function"==typeof(n=t.toString)&&!r(o=n.call(t)))return o;if("function"==typeof(n=t.valueOf)&&!r(o=n.call(t)))return o;if(!e&&"function"==typeof(n=t.toString)&&!r(o=n.call(t)))return o;throw TypeError("Can't convert object to primitive value")}},function(t,e){t.exports="constructor,hasOwnProperty,isPrototypeOf,propertyIsEnumerable,toLocaleString,toString,valueOf".split(",")},function(t,e,n){var r=n(1),o=n(2),i=o["__core-js_shared__"]||(o["__core-js_shared__"]={});(t.exports=function(t,e){return i[t]||(i[t]=void 0!==e?e:{})})("versions",[]).push({version:r.version,mode:n(15)?"pure":"global",copyright:"© 2018 Denis Pushkarev (zloirock.ru)"})},function(t,e,n){var r=n(26)("keys"),o=n(14);t.exports=function(t){return r[t]||(r[t]=o(t))}},function(t,e){var n=Math.ceil,r=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?r:n)(t)}},function(t,e){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e,n){var r=n(29);t.exports=function(t){return Object(r(t))}},function(t,e,n){"use strict";e.__esModule=!0;var r=function(t){return t&&t.__esModule?t:{default:t}}(n(17));e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,r.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){var r=n(18),o=n(13),i=n(10),u=n(24),c=n(5),f=n(39),a=Object.getOwnPropertyDescriptor;e.f=n(3)?a:function(t,e){if(t=i(t),e=u(e,!0),f)try{return a(t,e)}catch(t){}if(c(t,e))return o(!r.f.call(t,e),t[e])}},function(t,e,n){var r=n(43),o=n(25).concat("length","prototype");e.f=Object.getOwnPropertyNames||function(t){return r(t,o)}},function(t,e){e.f=Object.getOwnPropertySymbols},function(t,e,n){var r=n(5),o=n(30),i=n(27)("IE_PROTO"),u=Object.prototype;t.exports=Object.getPrototypeOf||function(t){return t=o(t),r(t,i)?t[i]:"function"==typeof t.constructor&&t instanceof t.constructor?t.constructor.prototype:t instanceof Object?u:null}},function(t,e,n){t.exports=n(8)},function(t,e,n){"use strict";var r=n(15),o=n(9),i=n(36),u=n(8),c=n(23),f=n(79),a=n(21),s=n(35),l=n(6)("iterator"),p=!([].keys&&"next"in[].keys()),y=function(){return this};t.exports=function(t,e,n,d,v,b,h){f(n,e,d);var m,_,g,O=function(t){if(!p&&t in w)return w[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},x=e+" Iterator",S="values"==v,j=!1,w=t.prototype,P=w[l]||w["@@iterator"]||v&&w[v],E=P||O(v),M=v?S?O("entries"):E:void 0,T="Array"==e&&w.entries||P;if(T&&(g=s(T.call(new t)))!==Object.prototype&&g.next&&(a(g,x,!0),r||"function"==typeof g[l]||u(g,l,y)),S&&P&&"values"!==P.name&&(j=!0,E=function(){return P.call(this)}),r&&!h||!p&&!j&&w[l]||u(w,l,E),c[e]=E,c[x]=y,v)if(m={values:S?E:O("values"),keys:b?E:O("keys"),entries:M},h)for(_ in m)_ in w||i(w,_,m[_]);else o(o.P+o.F*(p||j),e,m);return m}},function(t,e,n){var r=n(7),o=n(2).document,i=r(o)&&r(o.createElement);t.exports=function(t){return i?o.createElement(t):{}}},function(t,e,n){t.exports=!n(3)&&!n(11)(function(){return 7!=Object.defineProperty(n(38)("div"),"a",{get:function(){return 7}}).a})},function(t,e,n){var r=n(84);t.exports=function(t,e,n){if(r(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,r){return t.call(e,n,r)};case 3:return function(n,r,o){return t.call(e,n,r,o)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){var r=n(9),o=n(1),i=n(11);t.exports=function(t,e){var n=(o.Object||{})[t]||Object[t],u={};u[t]=e(n),r(r.S+r.F*i(function(){n(1)}),"Object",u)}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){var r=n(5),o=n(10),i=n(87)(!1),u=n(27)("IE_PROTO");t.exports=function(t,e){var n,c=o(t),f=0,a=[];for(n in c)n!=u&&r(c,n)&&a.push(n);for(;e.length>f;)r(c,n=e[f++])&&(~i(a,n)||a.push(n));return a}},function(t,e,n){"use strict";e.__esModule=!0;var r=u(n(57)),o=u(n(53)),i=u(n(17));function u(t){return t&&t.__esModule?t:{default:t}}e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,i.default)(e)));t.prototype=(0,o.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(r.default?(0,r.default)(t,e):t.__proto__=e)}},function(t,e,n){"use strict";e.__esModule=!0;var r=function(t){return t&&t.__esModule?t:{default:t}}(n(60));e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var o=e[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),(0,r.default)(t,o.key,o)}}return function(e,n,r){return n&&t(e.prototype,n),r&&t(e,r),e}}()},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){t.exports={default:n(62),__esModule:!0}},function(t,e,n){t.exports={default:n(90),__esModule:!0}},function(t,n){t.exports=e},function(t,e){},function(t,e,n){var r=n(9);r(r.S,"Object",{create:n(22)})},function(t,e,n){n(51);var r=n(1).Object;t.exports=function(t,e){return r.create(t,e)}},function(t,e,n){t.exports={default:n(52),__esModule:!0}},function(t,e,n){var r=n(7),o=n(12),i=function(t,e){if(o(t),!r(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,r){try{(r=n(40)(Function.call,n(32).f(Object.prototype,"__proto__").set,2))(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return i(t,n),e?t.__proto__=n:r(t,n),t}}({},!1):void 0),check:i}},function(t,e,n){var r=n(9);r(r.S,"Object",{setPrototypeOf:n(54).set})},function(t,e,n){n(55),t.exports=n(1).Object.setPrototypeOf},function(t,e,n){t.exports={default:n(56),__esModule:!0}},function(t,e,n){var r=n(9);r(r.S+r.F*!n(3),"Object",{defineProperty:n(4).f})},function(t,e,n){n(58);var r=n(1).Object;t.exports=function(t,e,n){return r.defineProperty(t,e,n)}},function(t,e,n){t.exports={default:n(59),__esModule:!0}},function(t,e,n){var r=n(30),o=n(35);n(41)("getPrototypeOf",function(){return function(t){return o(r(t))}})},function(t,e,n){n(61),t.exports=n(1).Object.getPrototypeOf},function(t,e,n){n(19)("observable")},function(t,e,n){n(19)("asyncIterator")},function(t,e){},function(t,e,n){var r=n(10),o=n(33).f,i={}.toString,u="object"==typeof window&&window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[];t.exports.f=function(t){return u&&"[object Window]"==i.call(t)?function(t){try{return o(t)}catch(t){return u.slice()}}(t):o(r(t))}},function(t,e,n){var r=n(42);t.exports=Array.isArray||function(t){return"Array"==r(t)}},function(t,e,n){var r=n(16),o=n(34),i=n(18);t.exports=function(t){var e=r(t),n=o.f;if(n)for(var u,c=n(t),f=i.f,a=0;c.length>a;)f.call(t,u=c[a++])&&e.push(u);return e}},function(t,e,n){var r=n(14)("meta"),o=n(7),i=n(5),u=n(4).f,c=0,f=Object.isExtensible||function(){return!0},a=!n(11)(function(){return f(Object.preventExtensions({}))}),s=function(t){u(t,r,{value:{i:"O"+ ++c,w:{}}})},l=t.exports={KEY:r,NEED:!1,fastKey:function(t,e){if(!o(t))return"symbol"==typeof t?t:("string"==typeof t?"S":"P")+t;if(!i(t,r)){if(!f(t))return"F";if(!e)return"E";s(t)}return t[r].i},getWeak:function(t,e){if(!i(t,r)){if(!f(t))return!0;if(!e)return!1;s(t)}return t[r].w},onFreeze:function(t){return a&&l.NEED&&f(t)&&!i(t,r)&&s(t),t}}},function(t,e,n){"use strict";var r=n(2),o=n(5),i=n(3),u=n(9),c=n(36),f=n(69).KEY,a=n(11),s=n(26),l=n(21),p=n(14),y=n(6),d=n(20),v=n(19),b=n(68),h=n(67),m=n(12),_=n(7),g=n(10),O=n(24),x=n(13),S=n(22),j=n(66),w=n(32),P=n(4),E=n(16),M=w.f,T=P.f,k=j.f,L=r.Symbol,A=r.JSON,F=A&&A.stringify,N=y("_hidden"),C=y("toPrimitive"),I={}.propertyIsEnumerable,D=s("symbol-registry"),G=s("symbols"),R=s("op-symbols"),V=Object.prototype,W="function"==typeof L,z=r.QObject,B=!z||!z.prototype||!z.prototype.findChild,H=i&&a(function(){return 7!=S(T({},"a",{get:function(){return T(this,"a",{value:7}).a}})).a})?function(t,e,n){var r=M(V,e);r&&delete V[e],T(t,e,n),r&&t!==V&&T(V,e,r)}:T,J=function(t){var e=G[t]=S(L.prototype);return e._k=t,e},K=W&&"symbol"==typeof L.iterator?function(t){return"symbol"==typeof t}:function(t){return t instanceof L},Y=function(t,e,n){return t===V&&Y(R,e,n),m(t),e=O(e,!0),m(n),o(G,e)?(n.enumerable?(o(t,N)&&t[N][e]&&(t[N][e]=!1),n=S(n,{enumerable:x(0,!1)})):(o(t,N)||T(t,N,x(1,{})),t[N][e]=!0),H(t,e,n)):T(t,e,n)},q=function(t,e){m(t);for(var n,r=b(e=g(e)),o=0,i=r.length;i>o;)Y(t,n=r[o++],e[n]);return t},Q=function(t){var e=I.call(this,t=O(t,!0));return!(this===V&&o(G,t)&&!o(R,t))&&(!(e||!o(this,t)||!o(G,t)||o(this,N)&&this[N][t])||e)},U=function(t,e){if(t=g(t),e=O(e,!0),t!==V||!o(G,e)||o(R,e)){var n=M(t,e);return!n||!o(G,e)||o(t,N)&&t[N][e]||(n.enumerable=!0),n}},X=function(t){for(var e,n=k(g(t)),r=[],i=0;n.length>i;)o(G,e=n[i++])||e==N||e==f||r.push(e);return r},Z=function(t){for(var e,n=t===V,r=k(n?R:g(t)),i=[],u=0;r.length>u;)!o(G,e=r[u++])||n&&!o(V,e)||i.push(G[e]);return i};W||(c((L=function(){if(this instanceof L)throw TypeError("Symbol is not a constructor!");var t=p(arguments.length>0?arguments[0]:void 0),e=function(n){this===V&&e.call(R,n),o(this,N)&&o(this[N],t)&&(this[N][t]=!1),H(this,t,x(1,n))};return i&&B&&H(V,t,{configurable:!0,set:e}),J(t)}).prototype,"toString",function(){return this._k}),w.f=U,P.f=Y,n(33).f=j.f=X,n(18).f=Q,n(34).f=Z,i&&!n(15)&&c(V,"propertyIsEnumerable",Q,!0),d.f=function(t){return J(y(t))}),u(u.G+u.W+u.F*!W,{Symbol:L});for(var $="hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),tt=0;$.length>tt;)y($[tt++]);for(var et=E(y.store),nt=0;et.length>nt;)v(et[nt++]);u(u.S+u.F*!W,"Symbol",{for:function(t){return o(D,t+="")?D[t]:D[t]=L(t)},keyFor:function(t){if(!K(t))throw TypeError(t+" is not a symbol!");for(var e in D)if(D[e]===t)return e},useSetter:function(){B=!0},useSimple:function(){B=!1}}),u(u.S+u.F*!W,"Object",{create:function(t,e){return void 0===e?S(t):q(S(t),e)},defineProperty:Y,defineProperties:q,getOwnPropertyDescriptor:U,getOwnPropertyNames:X,getOwnPropertySymbols:Z}),A&&u(u.S+u.F*(!W||a(function(){var t=L();return"[null]"!=F([t])||"{}"!=F({a:t})||"{}"!=F(Object(t))})),"JSON",{stringify:function(t){for(var e,n,r=[t],o=1;arguments.length>o;)r.push(arguments[o++]);if(n=e=r[1],(_(e)||void 0!==t)&&!K(t))return h(e)||(e=function(t,e){if("function"==typeof n&&(e=n.call(this,t,e)),!K(e))return e}),r[1]=e,F.apply(A,r)}}),L.prototype[C]||n(8)(L.prototype,C,L.prototype.valueOf),l(L,"Symbol"),l(Math,"Math",!0),l(r.JSON,"JSON",!0)},function(t,e,n){n(70),n(65),n(64),n(63),t.exports=n(1).Symbol},function(t,e,n){t.exports={default:n(71),__esModule:!0}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e){t.exports=function(){}},function(t,e,n){"use strict";var r=n(74),o=n(73),i=n(23),u=n(10);t.exports=n(37)(Array,"Array",function(t,e){this._t=u(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,o(1)):o(0,"keys"==e?n:"values"==e?t[n]:[n,t[n]])},"values"),i.Arguments=i.Array,r("keys"),r("values"),r("entries")},function(t,e,n){n(75);for(var r=n(2),o=n(8),i=n(23),u=n(6)("toStringTag"),c="CSSRuleList,CSSStyleDeclaration,CSSValueList,ClientRectList,DOMRectList,DOMStringList,DOMTokenList,DataTransferItemList,FileList,HTMLAllCollection,HTMLCollection,HTMLFormElement,HTMLSelectElement,MediaList,MimeTypeArray,NamedNodeMap,NodeList,PaintRequestList,Plugin,PluginArray,SVGLengthList,SVGNumberList,SVGPathSegList,SVGPointList,SVGStringList,SVGTransformList,SourceBufferList,StyleSheetList,TextTrackCueList,TextTrackList,TouchList".split(","),f=0;f<c.length;f++){var a=c[f],s=r[a],l=s&&s.prototype;l&&!l[u]&&o(l,u,a),i[a]=i.Array}},function(t,e,n){var r=n(2).document;t.exports=r&&r.documentElement},function(t,e,n){var r=n(4),o=n(12),i=n(16);t.exports=n(3)?Object.defineProperties:function(t,e){o(t);for(var n,u=i(e),c=u.length,f=0;c>f;)r.f(t,n=u[f++],e[n]);return t}},function(t,e,n){"use strict";var r=n(22),o=n(13),i=n(21),u={};n(8)(u,n(6)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=r(u,{next:o(1,n)}),i(t,e+" Iterator")}},function(t,e,n){var r=n(28),o=n(29);t.exports=function(t){return function(e,n){var i,u,c=String(o(e)),f=r(n),a=c.length;return f<0||f>=a?t?"":void 0:(i=c.charCodeAt(f))<55296||i>56319||f+1===a||(u=c.charCodeAt(f+1))<56320||u>57343?t?c.charAt(f):i:t?c.slice(f,f+2):u-56320+(i-55296<<10)+65536}}},function(t,e,n){"use strict";var r=n(80)(!0);n(37)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=r(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){n(81),n(76),t.exports=n(20).f("iterator")},function(t,e,n){t.exports={default:n(82),__esModule:!0}},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){var r=n(28),o=Math.max,i=Math.min;t.exports=function(t,e){return(t=r(t))<0?o(t+e,0):i(t,e)}},function(t,e,n){var r=n(28),o=Math.min;t.exports=function(t){return t>0?o(r(t),9007199254740991):0}},function(t,e,n){var r=n(10),o=n(86),i=n(85);t.exports=function(t){return function(e,n,u){var c,f=r(e),a=o(f.length),s=i(u,a);if(t&&n!=n){for(;a>s;)if((c=f[s++])!=c)return!0}else for(;a>s;s++)if((t||s in f)&&f[s]===n)return t||s||0;return!t&&-1}}},function(t,e,n){var r=n(42);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==r(t)?t.split(""):Object(t)}},function(t,e,n){var r=n(30),o=n(16);n(41)("keys",function(){return function(t){return o(r(t))}})},function(t,e,n){n(89),t.exports=n(1).Object.keys},function(t,e){t.exports=function(t){if(!t.webpackPolyfill){var e=Object.create(t);e.children||(e.children=[]),Object.defineProperty(e,"loaded",{enumerable:!0,get:function(){return e.l}}),Object.defineProperty(e,"id",{enumerable:!0,get:function(){return e.i}}),Object.defineProperty(e,"exports",{enumerable:!0}),e.webpackPolyfill=1}return e}},function(t,e,n){"use strict";n.r(e),function(t){n.d(e,"Initialize",function(){return _});var r=n(48),o=n.n(r),i=n(17),u=n.n(i),c=n(47),f=n.n(c),a=n(46),s=n.n(a),l=n(45),p=n.n(l),y=n(31),d=n.n(y),v=n(44),b=n.n(v),h=(n(50),n(0)),m=n.n(h);n(49);function _(e,n,r,o,i,u){if(t.properties=Application.Properties[n],t.settings=o,console.log("entity initialize",o,o.object),o.object){var c=o.object+"_default";_reg("Blocks",c)||_r("Blocks",c,function(t,e,n){return console.log("rendering default display",c,t,e,n),m.a.createElement("h1",null,"default display")})}}Window.displayDefaultEntity=function(t,e,n){return m.a.createElement(g,{desc:e,uikit:n,ctx:t})};var g=function(t){function e(){var t,n,r,i;s()(this,e);for(var c=arguments.length,a=Array(c),l=0;l<c;l++)a[l]=arguments[l];return n=r=d()(this,(t=e.__proto__||f()(e)).call.apply(t,[this].concat(a))),r.createField=function(t,e,n,o,i,u){var c=r.createObjFields(t,n+1,o,i,u);return m.a.createElement("div",{className:"field "+e},m.a.createElement("div",{className:"name"},e),m.a.createElement("div",{className:"value"},c))},r.createObjFields=function(t,e,n,i,c){if(null==t)return null;if(t instanceof Array){for(var f=new Array,a=0;a<t.length;a++)f.push(m.a.createElement("div",{className:"entityarrayitem"},r.createObjFields(t[a],e+1,n,i,c)));return f}if("object"==(void 0===t?"undefined":u()(t))){var s=new Array,l=new Array,p=r;return o()(t).forEach(function(r){var o=t[r],u=p.createField(o,r,e,n,i,c);console.log("field",r,"fieldVal",o," level ",e),o instanceof Array&&0==e?l.push(m.a.createElement(c.Tab,{label:r},u)):s.push(u)}),0!=e?s:m.a.createElement(c.Tabset,null,m.a.createElement(c.Tab,{label:"General"},s),l)}return t},i=n,d()(r,i)}return b()(e,t),p()(e,[{key:"render",value:function(){var t=this.props,e=t.ctx,n=t.desc,r=t.uikit;return console.log(e,n,r),m.a.createElement("div",{className:"entity "},this.createObjFields(e.data,0,e,n,r))}}]),e}(m.a.Component)}.call(this,n(91)(t))}])});
//# sourceMappingURL=index.js.map
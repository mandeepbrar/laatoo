define("authui",["react","uicommon","prop-types","react-redux","reactwebcommon","redux-saga","md5","redux"],function(t,e,n,r,o,i,a,u){return function(t){var e={};function n(r){if(e[r])return e[r].exports;var o=e[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=t,n.c=e,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},n.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var o in t)n.d(r,o,function(e){return t[e]}.bind(null,o));return r},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=104)}([function(e,n){e.exports=t},function(t,n){t.exports=e},function(t,e){var n=t.exports={version:"2.5.7"};"number"==typeof __e&&(__e=n)},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e){t.exports=n},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e,n){var r=n(3),o=n(2),i=n(47),a=n(11),u=n(5),c=function(t,e,n){var s,l,f,p=t&c.F,d=t&c.G,h=t&c.S,g=t&c.P,v=t&c.B,y=t&c.W,m=d?o:o[e]||(o[e]={}),b=m.prototype,S=d?r:h?r[e]:(r[e]||{}).prototype;for(s in d&&(n=e),n)(l=!p&&S&&void 0!==S[s])&&u(m,s)||(f=l?S[s]:n[s],m[s]=d&&"function"!=typeof S[s]?n[s]:v&&l?i(f,r):y&&S[s]==f?function(t){var e=function(e,n,r){if(this instanceof t){switch(arguments.length){case 0:return new t;case 1:return new t(e);case 2:return new t(e,n)}return new t(e,n,r)}return t.apply(this,arguments)};return e.prototype=t.prototype,e}(f):g&&"function"==typeof f?i(Function.call,f):f,g&&((m.virtual||(m.virtual={}))[s]=f,t&c.R&&b&&!b[s]&&a(b,s,f)))};c.F=1,c.G=2,c.S=4,c.P=8,c.B=16,c.W=32,c.U=64,c.R=128,t.exports=c},function(t,e,n){var r=n(22),o=n(48),i=n(34),a=Object.defineProperty;e.f=n(8)?Object.defineProperty:function(t,e,n){if(r(t),e=i(e,!0),r(n),o)try{return a(t,e,n)}catch(t){}if("get"in n||"set"in n)throw TypeError("Accessors not supported!");return"value"in n&&(t[e]=n.value),t}},function(t,e,n){t.exports=!n(13)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){t.exports=n(102)},function(t,e,n){var r=n(44),o=n(29);t.exports=function(t){return r(o(t))}},function(t,e,n){var r=n(7),o=n(27);t.exports=n(8)?function(t,e,n){return r.f(t,e,o(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e,n){var r=n(32)("wks"),o=n(26),i=n(3).Symbol,a="function"==typeof i;(t.exports=function(t){return r[t]||(r[t]=a&&i[t]||(a?i:o)("Symbol."+t))}).store=r},function(t,e,n){t.exports={default:n(64),__esModule:!0}},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){"use strict";e.__esModule=!0;var r=function(t){return t&&t.__esModule?t:{default:t}}(n(66));e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var o=e[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),(0,r.default)(t,o.key,o)}}return function(e,n,r){return n&&t(e.prototype,n),r&&t(e,r),e}}()},function(t,e,n){"use strict";e.__esModule=!0;var r=function(t){return t&&t.__esModule?t:{default:t}}(n(51));e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,r.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){"use strict";e.__esModule=!0;var r=a(n(90)),o=a(n(94)),i=a(n(51));function a(t){return t&&t.__esModule?t:{default:t}}e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,i.default)(e)));t.prototype=(0,o.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(r.default?(0,r.default)(t,e):t.__proto__=e)}},function(t,e){t.exports=r},function(t,e,n){var r=n(43),o=n(33);t.exports=Object.keys||function(t){return r(t,o)}},function(t,e,n){var r=n(12);t.exports=function(t){if(!r(t))throw TypeError(t+" is not an object!");return t}},function(t,e){t.exports=o},function(t,e,n){var r=n(29);t.exports=function(t){return Object(r(t))}},function(t,e){t.exports=!0},function(t,e){var n=0,r=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+r).toString(36))}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e){e.f={}.propertyIsEnumerable},function(t,e){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e){var n=Math.ceil,r=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?r:n)(t)}},function(t,e,n){var r=n(32)("keys"),o=n(26);t.exports=function(t){return r[t]||(r[t]=o(t))}},function(t,e,n){var r=n(2),o=n(3),i=o["__core-js_shared__"]||(o["__core-js_shared__"]={});(t.exports=function(t,e){return i[t]||(i[t]=void 0!==e?e:{})})("versions",[]).push({version:r.version,mode:n(25)?"pure":"global",copyright:"© 2018 Denis Pushkarev (zloirock.ru)"})},function(t,e){t.exports="constructor,hasOwnProperty,isPrototypeOf,propertyIsEnumerable,toLocaleString,toString,valueOf".split(",")},function(t,e,n){var r=n(12);t.exports=function(t,e){if(!r(t))return t;var n,o;if(e&&"function"==typeof(n=t.toString)&&!r(o=n.call(t)))return o;if("function"==typeof(n=t.valueOf)&&!r(o=n.call(t)))return o;if(!e&&"function"==typeof(n=t.toString)&&!r(o=n.call(t)))return o;throw TypeError("Can't convert object to primitive value")}},function(t,e){t.exports={}},function(t,e,n){var r=n(22),o=n(74),i=n(33),a=n(31)("IE_PROTO"),u=function(){},c=function(){var t,e=n(49)("iframe"),r=i.length;for(e.style.display="none",n(75).appendChild(e),e.src="javascript:",(t=e.contentWindow.document).open(),t.write("<script>document.F=Object<\/script>"),t.close(),c=t.F;r--;)delete c.prototype[i[r]];return c()};t.exports=Object.create||function(t,e){var n;return null!==t?(u.prototype=r(t),n=new u,u.prototype=null,n[a]=t):n=c(),void 0===e?n:o(n,e)}},function(t,e,n){var r=n(7).f,o=n(5),i=n(14)("toStringTag");t.exports=function(t,e,n){t&&!o(t=n?t:t.prototype,i)&&r(t,i,{configurable:!0,value:e})}},function(t,e,n){e.f=n(14)},function(t,e,n){var r=n(3),o=n(2),i=n(25),a=n(38),u=n(7).f;t.exports=function(t){var e=o.Symbol||(o.Symbol=i?{}:r.Symbol||{});"_"==t.charAt(0)||t in e||u(e,t,{value:a.f(t)})}},function(t,e){e.f=Object.getOwnPropertySymbols},function(t,e,n){t.exports={default:n(99),__esModule:!0}},function(t,e){t.exports=i},function(t,e,n){var r=n(5),o=n(10),i=n(60)(!1),a=n(31)("IE_PROTO");t.exports=function(t,e){var n,u=o(t),c=0,s=[];for(n in u)n!=a&&r(u,n)&&s.push(n);for(;e.length>c;)r(u,n=e[c++])&&(~i(s,n)||s.push(n));return s}},function(t,e,n){var r=n(45);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==r(t)?t.split(""):Object(t)}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){var r=n(6),o=n(2),i=n(13);t.exports=function(t,e){var n=(o.Object||{})[t]||Object[t],a={};a[t]=e(n),r(r.S+r.F*i(function(){n(1)}),"Object",a)}},function(t,e,n){var r=n(63);t.exports=function(t,e,n){if(r(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,r){return t.call(e,n,r)};case 3:return function(n,r,o){return t.call(e,n,r,o)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){t.exports=!n(8)&&!n(13)(function(){return 7!=Object.defineProperty(n(49)("div"),"a",{get:function(){return 7}}).a})},function(t,e,n){var r=n(12),o=n(3).document,i=r(o)&&r(o.createElement);t.exports=function(t){return i?o.createElement(t):{}}},function(t,e,n){var r=n(5),o=n(24),i=n(31)("IE_PROTO"),a=Object.prototype;t.exports=Object.getPrototypeOf||function(t){return t=o(t),r(t,i)?t[i]:"function"==typeof t.constructor&&t instanceof t.constructor?t.constructor.prototype:t instanceof Object?a:null}},function(t,e,n){"use strict";e.__esModule=!0;var r=a(n(69)),o=a(n(80)),i="function"==typeof o.default&&"symbol"==typeof r.default?function(t){return typeof t}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":typeof t};function a(t){return t&&t.__esModule?t:{default:t}}e.default="function"==typeof o.default&&"symbol"===i(r.default)?function(t){return void 0===t?"undefined":i(t)}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":void 0===t?"undefined":i(t)}},function(t,e,n){"use strict";var r=n(25),o=n(6),i=n(53),a=n(11),u=n(35),c=n(73),s=n(37),l=n(50),f=n(14)("iterator"),p=!([].keys&&"next"in[].keys()),d=function(){return this};t.exports=function(t,e,n,h,g,v,y){c(n,e,h);var m,b,S,O=function(t){if(!p&&t in L)return L[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},x=e+" Iterator",_="values"==g,w=!1,L=t.prototype,E=L[f]||L["@@iterator"]||g&&L[g],j=E||O(g),k=g?_?O("entries"):j:void 0,N="Array"==e&&L.entries||E;if(N&&(S=l(N.call(new t)))!==Object.prototype&&S.next&&(s(S,x,!0),r||"function"==typeof S[f]||a(S,f,d)),_&&E&&"values"!==E.name&&(w=!0,j=function(){return E.call(this)}),r&&!y||!p&&!w&&L[f]||a(L,f,j),u[e]=j,u[x]=d,g)if(m={values:_?j:O("values"),keys:v?j:O("keys"),entries:k},y)for(b in m)b in L||i(L,b,m[b]);else o(o.P+o.F*(p||w),e,m);return m}},function(t,e,n){t.exports=n(11)},function(t,e,n){var r=n(43),o=n(33).concat("length","prototype");e.f=Object.getOwnPropertyNames||function(t){return r(t,o)}},function(t,e,n){var r=n(28),o=n(27),i=n(10),a=n(34),u=n(5),c=n(48),s=Object.getOwnPropertyDescriptor;e.f=n(8)?s:function(t,e){if(t=i(t),e=a(e,!0),c)try{return s(t,e)}catch(t){}if(u(t,e))return o(!r.f.call(t,e),t[e])}},function(t,e,n){t.exports={default:n(58),__esModule:!0}},function(t,e){t.exports=a},function(t,e,n){n(59),t.exports=n(2).Object.keys},function(t,e,n){var r=n(24),o=n(21);n(46)("keys",function(){return function(t){return o(r(t))}})},function(t,e,n){var r=n(10),o=n(61),i=n(62);t.exports=function(t){return function(e,n,a){var u,c=r(e),s=o(c.length),l=i(a,s);if(t&&n!=n){for(;s>l;)if((u=c[l++])!=u)return!0}else for(;s>l;l++)if((t||l in c)&&c[l]===n)return t||l||0;return!t&&-1}}},function(t,e,n){var r=n(30),o=Math.min;t.exports=function(t){return t>0?o(r(t),9007199254740991):0}},function(t,e,n){var r=n(30),o=Math.max,i=Math.min;t.exports=function(t,e){return(t=r(t))<0?o(t+e,0):i(t,e)}},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){n(65),t.exports=n(2).Object.getPrototypeOf},function(t,e,n){var r=n(24),o=n(50);n(46)("getPrototypeOf",function(){return function(t){return o(r(t))}})},function(t,e,n){t.exports={default:n(67),__esModule:!0}},function(t,e,n){n(68);var r=n(2).Object;t.exports=function(t,e,n){return r.defineProperty(t,e,n)}},function(t,e,n){var r=n(6);r(r.S+r.F*!n(8),"Object",{defineProperty:n(7).f})},function(t,e,n){t.exports={default:n(70),__esModule:!0}},function(t,e,n){n(71),n(76),t.exports=n(38).f("iterator")},function(t,e,n){"use strict";var r=n(72)(!0);n(52)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=r(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){var r=n(30),o=n(29);t.exports=function(t){return function(e,n){var i,a,u=String(o(e)),c=r(n),s=u.length;return c<0||c>=s?t?"":void 0:(i=u.charCodeAt(c))<55296||i>56319||c+1===s||(a=u.charCodeAt(c+1))<56320||a>57343?t?u.charAt(c):i:t?u.slice(c,c+2):a-56320+(i-55296<<10)+65536}}},function(t,e,n){"use strict";var r=n(36),o=n(27),i=n(37),a={};n(11)(a,n(14)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=r(a,{next:o(1,n)}),i(t,e+" Iterator")}},function(t,e,n){var r=n(7),o=n(22),i=n(21);t.exports=n(8)?Object.defineProperties:function(t,e){o(t);for(var n,a=i(e),u=a.length,c=0;u>c;)r.f(t,n=a[c++],e[n]);return t}},function(t,e,n){var r=n(3).document;t.exports=r&&r.documentElement},function(t,e,n){n(77);for(var r=n(3),o=n(11),i=n(35),a=n(14)("toStringTag"),u="CSSRuleList,CSSStyleDeclaration,CSSValueList,ClientRectList,DOMRectList,DOMStringList,DOMTokenList,DataTransferItemList,FileList,HTMLAllCollection,HTMLCollection,HTMLFormElement,HTMLSelectElement,MediaList,MimeTypeArray,NamedNodeMap,NodeList,PaintRequestList,Plugin,PluginArray,SVGLengthList,SVGNumberList,SVGPathSegList,SVGPointList,SVGStringList,SVGTransformList,SourceBufferList,StyleSheetList,TextTrackCueList,TextTrackList,TouchList".split(","),c=0;c<u.length;c++){var s=u[c],l=r[s],f=l&&l.prototype;f&&!f[a]&&o(f,a,s),i[s]=i.Array}},function(t,e,n){"use strict";var r=n(78),o=n(79),i=n(35),a=n(10);t.exports=n(52)(Array,"Array",function(t,e){this._t=a(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,o(1)):o(0,"keys"==e?n:"values"==e?t[n]:[n,t[n]])},"values"),i.Arguments=i.Array,r("keys"),r("values"),r("entries")},function(t,e){t.exports=function(){}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e,n){t.exports={default:n(81),__esModule:!0}},function(t,e,n){n(82),n(87),n(88),n(89),t.exports=n(2).Symbol},function(t,e,n){"use strict";var r=n(3),o=n(5),i=n(8),a=n(6),u=n(53),c=n(83).KEY,s=n(13),l=n(32),f=n(37),p=n(26),d=n(14),h=n(38),g=n(39),v=n(84),y=n(85),m=n(22),b=n(12),S=n(10),O=n(34),x=n(27),_=n(36),w=n(86),L=n(55),E=n(7),j=n(21),k=L.f,N=E.f,I=w.f,A=r.Symbol,T=r.JSON,P=T&&T.stringify,C=d("_hidden"),G=d("toPrimitive"),F={}.propertyIsEnumerable,R=l("symbol-registry"),M=l("symbols"),U=l("op-symbols"),D=Object.prototype,B="function"==typeof A,q=r.QObject,W=!q||!q.prototype||!q.prototype.findChild,V=i&&s(function(){return 7!=_(N({},"a",{get:function(){return N(this,"a",{value:7}).a}})).a})?function(t,e,n){var r=k(D,e);r&&delete D[e],N(t,e,n),r&&t!==D&&N(D,e,r)}:N,H=function(t){var e=M[t]=_(A.prototype);return e._k=t,e},K=B&&"symbol"==typeof A.iterator?function(t){return"symbol"==typeof t}:function(t){return t instanceof A},z=function(t,e,n){return t===D&&z(U,e,n),m(t),e=O(e,!0),m(n),o(M,e)?(n.enumerable?(o(t,C)&&t[C][e]&&(t[C][e]=!1),n=_(n,{enumerable:x(0,!1)})):(o(t,C)||N(t,C,x(1,{})),t[C][e]=!0),V(t,e,n)):N(t,e,n)},J=function(t,e){m(t);for(var n,r=v(e=S(e)),o=0,i=r.length;i>o;)z(t,n=r[o++],e[n]);return t},Y=function(t){var e=F.call(this,t=O(t,!0));return!(this===D&&o(M,t)&&!o(U,t))&&(!(e||!o(this,t)||!o(M,t)||o(this,C)&&this[C][t])||e)},Q=function(t,e){if(t=S(t),e=O(e,!0),t!==D||!o(M,e)||o(U,e)){var n=k(t,e);return!n||!o(M,e)||o(t,C)&&t[C][e]||(n.enumerable=!0),n}},X=function(t){for(var e,n=I(S(t)),r=[],i=0;n.length>i;)o(M,e=n[i++])||e==C||e==c||r.push(e);return r},Z=function(t){for(var e,n=t===D,r=I(n?U:S(t)),i=[],a=0;r.length>a;)!o(M,e=r[a++])||n&&!o(D,e)||i.push(M[e]);return i};B||(u((A=function(){if(this instanceof A)throw TypeError("Symbol is not a constructor!");var t=p(arguments.length>0?arguments[0]:void 0),e=function(n){this===D&&e.call(U,n),o(this,C)&&o(this[C],t)&&(this[C][t]=!1),V(this,t,x(1,n))};return i&&W&&V(D,t,{configurable:!0,set:e}),H(t)}).prototype,"toString",function(){return this._k}),L.f=Q,E.f=z,n(54).f=w.f=X,n(28).f=Y,n(40).f=Z,i&&!n(25)&&u(D,"propertyIsEnumerable",Y,!0),h.f=function(t){return H(d(t))}),a(a.G+a.W+a.F*!B,{Symbol:A});for(var $="hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),tt=0;$.length>tt;)d($[tt++]);for(var et=j(d.store),nt=0;et.length>nt;)g(et[nt++]);a(a.S+a.F*!B,"Symbol",{for:function(t){return o(R,t+="")?R[t]:R[t]=A(t)},keyFor:function(t){if(!K(t))throw TypeError(t+" is not a symbol!");for(var e in R)if(R[e]===t)return e},useSetter:function(){W=!0},useSimple:function(){W=!1}}),a(a.S+a.F*!B,"Object",{create:function(t,e){return void 0===e?_(t):J(_(t),e)},defineProperty:z,defineProperties:J,getOwnPropertyDescriptor:Q,getOwnPropertyNames:X,getOwnPropertySymbols:Z}),T&&a(a.S+a.F*(!B||s(function(){var t=A();return"[null]"!=P([t])||"{}"!=P({a:t})||"{}"!=P(Object(t))})),"JSON",{stringify:function(t){for(var e,n,r=[t],o=1;arguments.length>o;)r.push(arguments[o++]);if(n=e=r[1],(b(e)||void 0!==t)&&!K(t))return y(e)||(e=function(t,e){if("function"==typeof n&&(e=n.call(this,t,e)),!K(e))return e}),r[1]=e,P.apply(T,r)}}),A.prototype[G]||n(11)(A.prototype,G,A.prototype.valueOf),f(A,"Symbol"),f(Math,"Math",!0),f(r.JSON,"JSON",!0)},function(t,e,n){var r=n(26)("meta"),o=n(12),i=n(5),a=n(7).f,u=0,c=Object.isExtensible||function(){return!0},s=!n(13)(function(){return c(Object.preventExtensions({}))}),l=function(t){a(t,r,{value:{i:"O"+ ++u,w:{}}})},f=t.exports={KEY:r,NEED:!1,fastKey:function(t,e){if(!o(t))return"symbol"==typeof t?t:("string"==typeof t?"S":"P")+t;if(!i(t,r)){if(!c(t))return"F";if(!e)return"E";l(t)}return t[r].i},getWeak:function(t,e){if(!i(t,r)){if(!c(t))return!0;if(!e)return!1;l(t)}return t[r].w},onFreeze:function(t){return s&&f.NEED&&c(t)&&!i(t,r)&&l(t),t}}},function(t,e,n){var r=n(21),o=n(40),i=n(28);t.exports=function(t){var e=r(t),n=o.f;if(n)for(var a,u=n(t),c=i.f,s=0;u.length>s;)c.call(t,a=u[s++])&&e.push(a);return e}},function(t,e,n){var r=n(45);t.exports=Array.isArray||function(t){return"Array"==r(t)}},function(t,e,n){var r=n(10),o=n(54).f,i={}.toString,a="object"==typeof window&&window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[];t.exports.f=function(t){return a&&"[object Window]"==i.call(t)?function(t){try{return o(t)}catch(t){return a.slice()}}(t):o(r(t))}},function(t,e){},function(t,e,n){n(39)("asyncIterator")},function(t,e,n){n(39)("observable")},function(t,e,n){t.exports={default:n(91),__esModule:!0}},function(t,e,n){n(92),t.exports=n(2).Object.setPrototypeOf},function(t,e,n){var r=n(6);r(r.S,"Object",{setPrototypeOf:n(93).set})},function(t,e,n){var r=n(12),o=n(22),i=function(t,e){if(o(t),!r(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,r){try{(r=n(47)(Function.call,n(55).f(Object.prototype,"__proto__").set,2))(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return i(t,n),e?t.__proto__=n:r(t,n),t}}({},!1):void 0),check:i}},function(t,e,n){t.exports={default:n(95),__esModule:!0}},function(t,e,n){n(96);var r=n(2).Object;t.exports=function(t,e){return r.create(t,e)}},function(t,e,n){var r=n(6);r(r.S,"Object",{create:n(36)})},function(t,e){t.exports=u},function(t,e){},function(t,e,n){n(100),t.exports=n(2).Object.assign},function(t,e,n){var r=n(6);r(r.S+r.F,"Object",{assign:n(101)})},function(t,e,n){"use strict";var r=n(21),o=n(40),i=n(28),a=n(24),u=n(44),c=Object.assign;t.exports=!c||n(13)(function(){var t={},e={},n=Symbol(),r="abcdefghijklmnopqrst";return t[n]=7,r.split("").forEach(function(t){e[t]=t}),7!=c({},t)[n]||Object.keys(c({},e)).join("")!=r})?function(t,e){for(var n=a(t),c=arguments.length,s=1,l=o.f,f=i.f;c>s;)for(var p,d=u(arguments[s++]),h=l?r(d).concat(l(d)):r(d),g=h.length,v=0;g>v;)f.call(d,p=h[v++])&&(n[p]=d[p]);return n}:c},function(t,e,n){var r=function(){return this}()||Function("return this")(),o=r.regeneratorRuntime&&Object.getOwnPropertyNames(r).indexOf("regeneratorRuntime")>=0,i=o&&r.regeneratorRuntime;if(r.regeneratorRuntime=void 0,t.exports=n(103),o)r.regeneratorRuntime=i;else try{delete r.regeneratorRuntime}catch(t){r.regeneratorRuntime=void 0}},function(t,e){!function(e){"use strict";var n,r=Object.prototype,o=r.hasOwnProperty,i="function"==typeof Symbol?Symbol:{},a=i.iterator||"@@iterator",u=i.asyncIterator||"@@asyncIterator",c=i.toStringTag||"@@toStringTag",s="object"==typeof t,l=e.regeneratorRuntime;if(l)s&&(t.exports=l);else{(l=e.regeneratorRuntime=s?t.exports:{}).wrap=S;var f="suspendedStart",p="suspendedYield",d="executing",h="completed",g={},v={};v[a]=function(){return this};var y=Object.getPrototypeOf,m=y&&y(y(A([])));m&&m!==r&&o.call(m,a)&&(v=m);var b=w.prototype=x.prototype=Object.create(v);_.prototype=b.constructor=w,w.constructor=_,w[c]=_.displayName="GeneratorFunction",l.isGeneratorFunction=function(t){var e="function"==typeof t&&t.constructor;return!!e&&(e===_||"GeneratorFunction"===(e.displayName||e.name))},l.mark=function(t){return Object.setPrototypeOf?Object.setPrototypeOf(t,w):(t.__proto__=w,c in t||(t[c]="GeneratorFunction")),t.prototype=Object.create(b),t},l.awrap=function(t){return{__await:t}},L(E.prototype),E.prototype[u]=function(){return this},l.AsyncIterator=E,l.async=function(t,e,n,r){var o=new E(S(t,e,n,r));return l.isGeneratorFunction(e)?o:o.next().then(function(t){return t.done?t.value:o.next()})},L(b),b[c]="Generator",b[a]=function(){return this},b.toString=function(){return"[object Generator]"},l.keys=function(t){var e=[];for(var n in t)e.push(n);return e.reverse(),function n(){for(;e.length;){var r=e.pop();if(r in t)return n.value=r,n.done=!1,n}return n.done=!0,n}},l.values=A,I.prototype={constructor:I,reset:function(t){if(this.prev=0,this.next=0,this.sent=this._sent=n,this.done=!1,this.delegate=null,this.method="next",this.arg=n,this.tryEntries.forEach(N),!t)for(var e in this)"t"===e.charAt(0)&&o.call(this,e)&&!isNaN(+e.slice(1))&&(this[e]=n)},stop:function(){this.done=!0;var t=this.tryEntries[0].completion;if("throw"===t.type)throw t.arg;return this.rval},dispatchException:function(t){if(this.done)throw t;var e=this;function r(r,o){return u.type="throw",u.arg=t,e.next=r,o&&(e.method="next",e.arg=n),!!o}for(var i=this.tryEntries.length-1;i>=0;--i){var a=this.tryEntries[i],u=a.completion;if("root"===a.tryLoc)return r("end");if(a.tryLoc<=this.prev){var c=o.call(a,"catchLoc"),s=o.call(a,"finallyLoc");if(c&&s){if(this.prev<a.catchLoc)return r(a.catchLoc,!0);if(this.prev<a.finallyLoc)return r(a.finallyLoc)}else if(c){if(this.prev<a.catchLoc)return r(a.catchLoc,!0)}else{if(!s)throw new Error("try statement without catch or finally");if(this.prev<a.finallyLoc)return r(a.finallyLoc)}}}},abrupt:function(t,e){for(var n=this.tryEntries.length-1;n>=0;--n){var r=this.tryEntries[n];if(r.tryLoc<=this.prev&&o.call(r,"finallyLoc")&&this.prev<r.finallyLoc){var i=r;break}}i&&("break"===t||"continue"===t)&&i.tryLoc<=e&&e<=i.finallyLoc&&(i=null);var a=i?i.completion:{};return a.type=t,a.arg=e,i?(this.method="next",this.next=i.finallyLoc,g):this.complete(a)},complete:function(t,e){if("throw"===t.type)throw t.arg;return"break"===t.type||"continue"===t.type?this.next=t.arg:"return"===t.type?(this.rval=this.arg=t.arg,this.method="return",this.next="end"):"normal"===t.type&&e&&(this.next=e),g},finish:function(t){for(var e=this.tryEntries.length-1;e>=0;--e){var n=this.tryEntries[e];if(n.finallyLoc===t)return this.complete(n.completion,n.afterLoc),N(n),g}},catch:function(t){for(var e=this.tryEntries.length-1;e>=0;--e){var n=this.tryEntries[e];if(n.tryLoc===t){var r=n.completion;if("throw"===r.type){var o=r.arg;N(n)}return o}}throw new Error("illegal catch attempt")},delegateYield:function(t,e,r){return this.delegate={iterator:A(t),resultName:e,nextLoc:r},"next"===this.method&&(this.arg=n),g}}}function S(t,e,n,r){var o=e&&e.prototype instanceof x?e:x,i=Object.create(o.prototype),a=new I(r||[]);return i._invoke=function(t,e,n){var r=f;return function(o,i){if(r===d)throw new Error("Generator is already running");if(r===h){if("throw"===o)throw i;return T()}for(n.method=o,n.arg=i;;){var a=n.delegate;if(a){var u=j(a,n);if(u){if(u===g)continue;return u}}if("next"===n.method)n.sent=n._sent=n.arg;else if("throw"===n.method){if(r===f)throw r=h,n.arg;n.dispatchException(n.arg)}else"return"===n.method&&n.abrupt("return",n.arg);r=d;var c=O(t,e,n);if("normal"===c.type){if(r=n.done?h:p,c.arg===g)continue;return{value:c.arg,done:n.done}}"throw"===c.type&&(r=h,n.method="throw",n.arg=c.arg)}}}(t,n,a),i}function O(t,e,n){try{return{type:"normal",arg:t.call(e,n)}}catch(t){return{type:"throw",arg:t}}}function x(){}function _(){}function w(){}function L(t){["next","throw","return"].forEach(function(e){t[e]=function(t){return this._invoke(e,t)}})}function E(t){var e;this._invoke=function(n,r){function i(){return new Promise(function(e,i){!function e(n,r,i,a){var u=O(t[n],t,r);if("throw"!==u.type){var c=u.arg,s=c.value;return s&&"object"==typeof s&&o.call(s,"__await")?Promise.resolve(s.__await).then(function(t){e("next",t,i,a)},function(t){e("throw",t,i,a)}):Promise.resolve(s).then(function(t){c.value=t,i(c)},a)}a(u.arg)}(n,r,e,i)})}return e=e?e.then(i,i):i()}}function j(t,e){var r=t.iterator[e.method];if(r===n){if(e.delegate=null,"throw"===e.method){if(t.iterator.return&&(e.method="return",e.arg=n,j(t,e),"throw"===e.method))return g;e.method="throw",e.arg=new TypeError("The iterator does not provide a 'throw' method")}return g}var o=O(r,t.iterator,e.arg);if("throw"===o.type)return e.method="throw",e.arg=o.arg,e.delegate=null,g;var i=o.arg;return i?i.done?(e[t.resultName]=i.value,e.next=t.nextLoc,"return"!==e.method&&(e.method="next",e.arg=n),e.delegate=null,g):i:(e.method="throw",e.arg=new TypeError("iterator result is not an object"),e.delegate=null,g)}function k(t){var e={tryLoc:t[0]};1 in t&&(e.catchLoc=t[1]),2 in t&&(e.finallyLoc=t[2],e.afterLoc=t[3]),this.tryEntries.push(e)}function N(t){var e=t.completion||{};e.type="normal",delete e.arg,t.completion=e}function I(t){this.tryEntries=[{tryLoc:"root"}],t.forEach(k,this),this.reset(!0)}function A(t){if(t){var e=t[a];if(e)return e.call(t);if("function"==typeof t.next)return t;if(!isNaN(t.length)){var r=-1,i=function e(){for(;++r<t.length;)if(o.call(t,r))return e.value=t[r],e.done=!1,e;return e.value=n,e.done=!0,e};return i.next=i}}return{next:T}}function T(){return{value:n,done:!0}}}(function(){return this}()||Function("return this")())},function(t,e,n){"use strict";n.r(e);var r=n(56),o=n.n(r),i=n(0),a=n.n(i),u=n(15),c=n.n(u),s=n(16),l=n.n(s),f=n(17),p=n.n(f),d=n(18),h=n.n(d),g=n(19),v=n.n(g),y=n(4),m=function(t){function e(t){l()(this,e);var n=h()(this,(e.__proto__||c()(e)).call(this,t));console.log("costructor of login web"),n.state={email:"",password:""},n.handleLogin=n.handleLogin.bind(n),n.handleChange=n.handleChange.bind(n);return t.realm&&"?Realm="+t.realm,n}return v()(e,t),p()(e,[{key:"handleChange",value:function(t){var e={};e[t.target.name]=t.target.value,this.setState(e)}},{key:"handleLogin",value:function(){this.props.handleLogin(this.state.email,this.state.password)}},{key:"render",value:function(){return console.log("login ui",this.props),this.props.renderLogin(this.state,this.handleChange,this.handleLogin,this.oauthLogin,this.props)}}]),e}(a.a.Component);m.propTypes={handleOauthLogin:y.func.isRequired,handleLogin:y.func.isRequired},m.contextTypes={uikit:y.object};var b=n(57),S=n.n(b),O=n(20),x={LOGIN:"LOGIN",LOGGING_IN:"LOGGING_IN",LOGIN_SUCCESS:"LOGIN_SUCCESS",LOGIN_FAILURE:"LOGIN_FAILURE",LOGOUT:"LOGOUT",LOGOUT_SUCCESS:"LOGOUT_SUCCESS"},_=n(1),w=n(4),L=n.n(w),E=Object(O.connect)(function(t,e){return{realm:Application.Security.realm,renderLogin:e.renderLogin,signup:e.signup}},function(t,e){console.log("map dispatch of login compoent");var n="";return Application.Security.realm&&(n=Application.Security.realm),{handleLogin:function(e,r){var o={Username:e,Password:S()(r),Realm:n},i={serviceName:Application.Security.loginService};t(Object(_.createAction)(x.LOGIN,o,i))},handleOauthLogin:function(e){t(Object(_.createAction)(x.LOGIN_SUCCESS,{userId:e.id,token:e.token,permissions:e.permissions}))}}})(m);E.propTypes={loginService:L.a.string.isRequired,successpage:L.a.string,realm:L.a.string,signup:L.a.string};n(97);var j=n(4),k=function(t){function e(t){l()(this,e);var n=h()(this,(e.__proto__||c()(e)).call(this,t));return n.validatetoken=n.validatetoken.bind(n),n.state={loggedIn:t.loggedIn,validation:t.validation},t.validation&&n.validatetoken(),n}return v()(e,t),p()(e,[{key:"componentWillReceiveProps",value:function(t){t.loggedIn==this.state.loggedIn&&t.validation==this.state.validation||this.setState({loggedIn:t.loggedIn,validation:t.validation})}},{key:"validatetoken",value:function(){var t=this,e=this.props.logout,n=this.props.login,r=_.RequestBuilder.DefaultRequest({},{});_.DataSource.ExecuteService(this.props.validateService,r).then(function(t){n(t.data.Id,t.data.Permissions)},function(n){e(),t.setState({loggedIn:!1,validation:!1})})}},{key:"getChildContext",value:function(){return{loggedIn:this.state.loggedIn}}},{key:"render",value:function(){return this.state.validation?null:this.props.children?a.a.cloneElement(this.props.children,{loggedIn:this.state.loggedIn,validation:this.state.validation}):null}}]),e}(a.a.Component);k.childContextTypes={loggedIn:j.bool,user:j.object};var N=Object(O.connect)(function(t,e){if(null!=Storage.auth&&""!=Storage.auth)return"LoggedIn"!=t.Security.status?{validation:!0,loggedIn:!1,validateService:e.validateService}:{validation:!1,loggedIn:!0,validateService:e.validateService}},function(t,e){return{login:function(e,n){t(Object(_.createAction)(x.LOGIN_SUCCESS,{userId:e,token:Storage.auth,user:Storage.user,permissions:n}))},logout:function(){t(Object(_.createAction)(x.LOGOUT,null,null))}}})(k),I=n(23);n(98);function A(t,e,n){return function(r,o,i,u,c){return console.log("renderLogin",n,"uikit",t,"settigs",e,"props",c),a.a.createElement("div",{className:c.className?c.className:" loginbox "},a.a.createElement("div",{className:"logintext"},n.loginForm.formtext),a.a.createElement("div",{className:"sociallogin"},a.a.createElement(I.Action,{widget:"button",method:function(){u(Application.Security.googleAuthUrl)},name:"googleAuth",className:"googleAuthAction"},n.loginForm.google)),a.a.createElement("div",{className:"separator"},n.loginForm.separator),a.a.createElement("div",{className:"main"},a.a.createElement(t.Form,{role:"form"},a.a.createElement("div",{className:"userfield"},a.a.createElement("label",{htmlFor:"email"},n.loginForm.userlabel),a.a.createElement(t.TextField,{className:"text",name:"email",value:r.email,placeholder:n.loginForm.userplaceholder,onChange:o})),a.a.createElement("div",{className:"passwordfield"},a.a.createElement("label",{htmlFor:"inputPassword"},n.loginForm.passwordlabel),a.a.createElement(t.TextField,{type:"password",className:"text",name:"password",value:r.password,placeholder:n.loginForm.passwordplaceholder,onChange:o})),a.a.createElement("a",{className:"pull-right",href:"#"},"Forgot password?"),a.a.createElement("div",{className:"checkbox"},a.a.createElement("label",null,a.a.createElement("input",{type:"checkbox"}),"Remember me")),a.a.createElement("div",{className:"actionbuttons"},a.a.createElement(I.Action,{widget:"button",className:"loginBtn",name:"loginAction",method:i},n.loginForm.loginBtnText)))))}}var T=n(41),P=n.n(T),C={status:"NotLogged",token:"",userId:"",permissions:[]};Application.Register("Reducers","Security",function(t,e){if(e.type)switch(e.type){case x.LOGGING_IN:return P()({},t,{status:"LoggingIn"});case x.LOGIN_SUCCESS:return t.authToken===e.payload.token?t:(Storage.auth=e.payload.token,Storage.permissions=e.payload.permissions,Storage.userId=e.payload.userId,Storage.userFullName=e.payload.user.Name,Storage.userName=e.payload.user.Username,Storage.email=e.payload.user.Email,Storage.user=e.payload.user,P()({},t,{status:"LoggedIn",authToken:e.payload.token,userId:e.payload.userId,permissions:e.payload.permissions}));case x.LOGIN_FAILURE:case x.LOGOUT_SUCCESS:return Storage.auth="",Storage.permissions=[],Storage.userId="",Storage.userName="",Storage.userFullName="",Storage.email="",Storage.user=null,C;default:return t||C}});var G=n(9),F=n.n(G),R=n(42),M=(Object.assign,"function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t}),U=function(t){return"@@redux-saga/"+t},D=U("TASK"),B=U("HELPER");function q(t,e,n){if(!e(t))throw function(t,e){var n=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"";"undefined"==typeof window?console.log("redux-saga "+t+": "+e+"\n"+(n&&n.stack||n)):console[t](e,n)}("error","uncaught at check",n),new Error(n)}var W=Object.prototype.hasOwnProperty;function V(t,e){return H.notUndef(t)&&W.call(t,e)}var H={undef:function(t){return null===t||void 0===t},notUndef:function(t){return null!==t&&void 0!==t},func:function(t){return"function"==typeof t},number:function(t){return"number"==typeof t},string:function(t){return"string"==typeof t},array:Array.isArray,object:function(t){return t&&!H.array(t)&&"object"===(void 0===t?"undefined":M(t))},promise:function(t){return t&&H.func(t.then)},iterator:function(t){return t&&H.func(t.next)&&H.func(t.throw)},iterable:function(t){return t&&H.func(Symbol)?H.func(t[Symbol.iterator]):H.array(t)},task:function(t){return t&&t[D]},observable:function(t){return t&&H.func(t.subscribe)},buffer:function(t){return t&&H.func(t.isEmpty)&&H.func(t.take)&&H.func(t.put)},pattern:function(t){return t&&(H.string(t)||"symbol"===(void 0===t?"undefined":M(t))||H.func(t)||H.array(t))},channel:function(t){return t&&H.func(t.take)&&H.func(t.close)},helper:function(t){return t&&t[B]},stringableFunc:function(t){return H.func(t)&&V(t,"toString")}};function K(t,e){return function(){return t.apply(void 0,arguments)}}Object.assign;var z=U("IO"),J="TAKE",Y="PUT",Q="CALL",X=function(t,e){var n;return(n={})[z]=!0,n[t]=e,n};function Z(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"*";if(arguments.length&&q(arguments[0],H.notUndef,"take(patternOrChannel): patternOrChannel is undefined"),H.pattern(t))return X(J,{pattern:t});if(H.channel(t))return X(J,{channel:t});throw new Error("take(patternOrChannel): argument "+String(t)+" is not valid channel or a valid pattern")}Z.maybe=function(){var t=Z.apply(void 0,arguments);return t[J].maybe=!0,t};Z.maybe;function $(t,e){return arguments.length>1?(q(t,H.notUndef,"put(channel, action): argument channel is undefined"),q(t,H.channel,"put(channel, action): argument "+t+" is not a valid channel"),q(e,H.notUndef,"put(channel, action): argument action is undefined")):(q(t,H.notUndef,"put(action): argument action is undefined"),e=t,t=null),X(Y,{channel:t,action:e})}function tt(t,e,n){q(e,H.notUndef,t+": argument fn is undefined");var r=null;if(H.array(e)){var o=e;r=o[0],e=o[1]}else if(e.fn){var i=e;r=i.context,e=i.fn}return r&&H.string(e)&&H.func(r[e])&&(e=r[e]),q(e,H.func,t+": argument "+e+" is not a function"),{context:r,fn:e,args:n}}function et(t){for(var e=arguments.length,n=Array(e>1?e-1:0),r=1;r<e;r++)n[r-1]=arguments[r];return X(Q,tt("call",t,n))}$.resolve=function(){var t=$.apply(void 0,arguments);return t[Y].resolve=!0,t},$.sync=K($.resolve);var nt=F.a.mark(it),rt=F.a.mark(at),ot=F.a.mark(ut);function it(t){var e,n,r,o,i,a,u,c;return F.a.wrap(function(s){for(;;)switch(s.prev=s.next){case 0:return s.prev=0,s.next=3,$(Object(_.createAction)(x.LOGGING_IN));case 3:return e=_.RequestBuilder.DefaultRequest(null,t.payload),s.next=6,et(_.DataSource.ExecuteService,t.meta.serviceName,e);case 6:return n=s.sent,r=Application.Security.AuthToken.toLowerCase(),o=n.info[r],i=n.data,a=n.data.Id,u=n.data.Permissions,c=Object(_.createAction)(x.LOGIN_SUCCESS,{userId:a,token:o,permissions:u,user:i}),s.next=15,$(c);case 15:console.log("dispatched login action &&&&"),s.next=22;break;case 18:return s.prev=18,s.t0=s.catch(0),s.next=22,$(Object(_.createAction)(x.LOGIN_FAILURE,s.t0));case 22:case"end":return s.stop()}},nt,this,[[0,18]])}function at(t){return F.a.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,$(Object(_.createAction)(x.LOGOUT_SUCCESS,{}));case 2:case"end":return t.stop()}},rt,this)}function ut(){return F.a.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,[Object(R.takeLatest)(x.LOGIN,it),Object(R.takeLatest)(x.LOGOUT,at)];case 2:case"end":return t.stop()}},ot,this)}Application.Register("Sagas","loginSaga",ut);var ct=function(t){function e(){return l()(this,e),h()(this,(e.__proto__||c()(e)).apply(this,arguments))}return v()(e,t),p()(e,[{key:"render",value:function(){var t=this.props,e=t.module.properties?t.module.properties:{},n=e.logoutText?e.logoutText:"Logout";return t.loggedIn?a.a.createElement(t.uikit.Block,{className:"userblock "+t.className},a.a.createElement(t.uikit.Block,{className:"username"},Storage.userFullName?Storage.userFullName:Storage.userName),a.a.createElement(I.Action,{name:"logout",method:t.logout,className:"logout"},n)):null}}]),e}(a.a.Component),st=Object(O.connect)(function(t,e){return{loggedIn:"LoggedIn"==t.Security.status}},function(t,e){return{logout:function(){t(Object(_.createAction)(x.LOGOUT,null,null))}}})(ct);n.d(e,"Initialize",function(){return pt}),n.d(e,"WebLoginForm",function(){return dt}),n.d(e,"LoginComponent",function(){return E}),n.d(e,"renderWebLogin",function(){return A}),n.d(e,"LoginValidator",function(){return N});var lt,ft=n(4);function pt(t,e,n,r,i,a){((lt=this).properties=Application.Properties[e],lt.settings=r,0!=o()(r).length)?Application.Security={googleAuthUrl:r.googleAuthUrl,loginService:r.loginService,validateService:r.validateService,loginServiceURL:r.loginServiceURL,realm:r.realm}:(Application.Security={loginService:"login",validateService:"validate",realm:""},_reg("Services")||(Application.Register("Services","login",{url:"/login",method:"POST"}),Application.Register("Services","validate",{url:"/validate",method:"POST"})));r.AuthToken?Application.Security.AuthToken=r.AuthToken:Application.Security.AuthToken="x-auth-token"}var dt=function(t,e){return console.log("render logiform",E),a.a.createElement(E,{className:t.className,renderLogin:A(e.uikit,lt.settings,lt.properties),realm:t.realm,loginService:t.loginService,loginServiceURL:t.loginServiceURL,googleAuthUrl:t.googleAuthUrl})};dt.contextTypes={uikit:ft.object},Application.Register("Blocks","userBlock",function(t,e,n,r){return a.a.createElement(st,{className:t.className,uikit:n,module:lt})})}])});
//# sourceMappingURL=index.js.map
define("authui",["react","uicommon","prop-types","react-redux","reactwebcommon","redux-saga","md5","redux"],function(t,e,n,r,o,i,a,u){return function(t){var e={};function n(r){if(e[r])return e[r].exports;var o=e[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=t,n.c=e,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},n.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var o in t)n.d(r,o,function(e){return t[e]}.bind(null,o));return r},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=80)}([function(e,n){e.exports=t},function(t,n){t.exports=e},function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(t,e){t.exports=n},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,e,n){t.exports=n(78)},function(t,e,n){t.exports={default:n(45),__esModule:!0}},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){"use strict";e.__esModule=!0;var r=function(t){return t&&t.__esModule?t:{default:t}}(n(47));e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var o=e[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),(0,r.default)(t,o.key,o)}}return function(e,n,r){return n&&t(e.prototype,n),r&&t(e,r),e}}()},function(t,e,n){"use strict";e.__esModule=!0;var r=function(t){return t&&t.__esModule?t:{default:t}}(n(29));e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,r.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){"use strict";e.__esModule=!0;var r=a(n(67)),o=a(n(71)),i=a(n(29));function a(t){return t&&t.__esModule?t:{default:t}}e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,i.default)(e)));t.prototype=(0,o.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(r.default?(0,r.default)(t,e):t.__proto__=e)}},function(t,e){t.exports=r},function(t,e,n){var r=n(15),o=n(4),i=n(28),a=function(t,e,n){var u,c,s,l=t&a.F,f=t&a.G,p=t&a.S,d=t&a.P,h=t&a.B,g=t&a.W,v=f?o:o[e]||(o[e]={}),y=f?r:p?r[e]:(r[e]||{}).prototype;for(u in f&&(n=e),n)(c=!l&&y&&u in y)&&u in v||(s=c?y[u]:n[u],v[u]=f&&"function"!=typeof y[u]?n[u]:h&&c?i(s,r):g&&y[u]==s?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(s):d&&"function"==typeof s?i(Function.call,s):s,d&&((v.prototype||(v.prototype={}))[u]=s))};a.F=1,a.G=2,a.S=4,a.P=8,a.B=16,a.W=32,t.exports=a},function(t,e,n){var r=n(34)("wks"),o=n(35),i=n(15).Symbol;t.exports=function(t){return r[t]||(r[t]=i&&i[t]||(i||o)("Symbol."+t))}},function(t,e){t.exports=o},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e,n){var r=n(36),o=n(19);t.exports=function(t){return r(o(t))}},function(t,e,n){var r=n(19);t.exports=function(t){return Object(r(t))}},function(t,e){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e,n){var r=n(2),o=n(21);t.exports=n(33)?function(t,e,n){return r.setDesc(t,e,o(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports={}},function(t,e,n){var r=n(2).setDesc,o=n(22),i=n(13)("toStringTag");t.exports=function(t,e,n){t&&!o(t=n?t:t.prototype,i)&&r(t,i,{configurable:!0,value:e})}},function(t,e,n){t.exports={default:n(75),__esModule:!0}},function(t,e){t.exports=i},function(t,e,n){var r=n(12),o=n(4),i=n(16);t.exports=function(t,e){var n=(o.Object||{})[t]||Object[t],a={};a[t]=e(n),r(r.S+r.F*i(function(){n(1)}),"Object",a)}},function(t,e,n){var r=n(44);t.exports=function(t,e,n){if(r(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,r){return t.call(e,n,r)};case 3:return function(n,r,o){return t.call(e,n,r,o)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){"use strict";e.__esModule=!0;var r=a(n(49)),o=a(n(59)),i="function"==typeof o.default&&"symbol"==typeof r.default?function(t){return typeof t}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":typeof t};function a(t){return t&&t.__esModule?t:{default:t}}e.default="function"==typeof o.default&&"symbol"===i(r.default)?function(t){return void 0===t?"undefined":i(t)}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":void 0===t?"undefined":i(t)}},function(t,e,n){"use strict";var r=n(31),o=n(12),i=n(32),a=n(20),u=n(22),c=n(23),s=n(54),l=n(24),f=n(2).getProto,p=n(13)("iterator"),d=!([].keys&&"next"in[].keys()),h=function(){return this};t.exports=function(t,e,n,g,v,y,m){s(n,e,g);var b,S,x=function(t){if(!d&&t in L)return L[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},_=e+" Iterator",O="values"==v,w=!1,L=t.prototype,E=L[p]||L["@@iterator"]||v&&L[v],k=E||x(v);if(E){var N=f(k.call(new t));l(N,_,!0),!r&&u(L,"@@iterator")&&a(N,p,h),O&&"values"!==E.name&&(w=!0,k=function(){return E.call(this)})}if(r&&!m||!d&&!w&&L[p]||a(L,p,k),c[e]=k,c[_]=h,v)if(b={values:O?k:x("values"),keys:y?k:x("keys"),entries:O?x("entries"):k},m)for(S in b)S in L||i(L,S,b[S]);else o(o.P+o.F*(d||w),e,b);return b}},function(t,e){t.exports=!0},function(t,e,n){t.exports=n(20)},function(t,e,n){t.exports=!n(16)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){var r=n(15),o=r["__core-js_shared__"]||(r["__core-js_shared__"]={});t.exports=function(t){return o[t]||(o[t]={})}},function(t,e){var n=0,r=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+r).toString(36))}},function(t,e,n){var r=n(37);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==r(t)?t.split(""):Object(t)}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){var r=n(39);t.exports=function(t){if(!r(t))throw TypeError(t+" is not an object!");return t}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){t.exports={default:n(42),__esModule:!0}},function(t,e){t.exports=a},function(t,e,n){n(43),t.exports=n(4).Object.keys},function(t,e,n){var r=n(18);n(27)("keys",function(t){return function(e){return t(r(e))}})},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){n(46),t.exports=n(4).Object.getPrototypeOf},function(t,e,n){var r=n(18);n(27)("getPrototypeOf",function(t){return function(e){return t(r(e))}})},function(t,e,n){t.exports={default:n(48),__esModule:!0}},function(t,e,n){var r=n(2);t.exports=function(t,e,n){return r.setDesc(t,e,n)}},function(t,e,n){t.exports={default:n(50),__esModule:!0}},function(t,e,n){n(51),n(55),t.exports=n(13)("iterator")},function(t,e,n){"use strict";var r=n(52)(!0);n(30)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=r(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){var r=n(53),o=n(19);t.exports=function(t){return function(e,n){var i,a,u=String(o(e)),c=r(n),s=u.length;return c<0||c>=s?t?"":void 0:(i=u.charCodeAt(c))<55296||i>56319||c+1===s||(a=u.charCodeAt(c+1))<56320||a>57343?t?u.charAt(c):i:t?u.slice(c,c+2):a-56320+(i-55296<<10)+65536}}},function(t,e){var n=Math.ceil,r=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?r:n)(t)}},function(t,e,n){"use strict";var r=n(2),o=n(21),i=n(24),a={};n(20)(a,n(13)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=r.create(a,{next:o(1,n)}),i(t,e+" Iterator")}},function(t,e,n){n(56);var r=n(23);r.NodeList=r.HTMLCollection=r.Array},function(t,e,n){"use strict";var r=n(57),o=n(58),i=n(23),a=n(17);t.exports=n(30)(Array,"Array",function(t,e){this._t=a(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,o(1)):o(0,"keys"==e?n:"values"==e?t[n]:[n,t[n]])},"values"),i.Arguments=i.Array,r("keys"),r("values"),r("entries")},function(t,e){t.exports=function(){}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e,n){t.exports={default:n(60),__esModule:!0}},function(t,e,n){n(61),n(66),t.exports=n(4).Symbol},function(t,e,n){"use strict";var r=n(2),o=n(15),i=n(22),a=n(33),u=n(12),c=n(32),s=n(16),l=n(34),f=n(24),p=n(35),d=n(13),h=n(62),g=n(63),v=n(64),y=n(65),m=n(38),b=n(17),S=n(21),x=r.getDesc,_=r.setDesc,O=r.create,w=g.get,L=o.Symbol,E=o.JSON,k=E&&E.stringify,N=!1,j=d("_hidden"),I=r.isEnum,A=l("symbol-registry"),P=l("symbols"),T="function"==typeof L,G=Object.prototype,C=a&&s(function(){return 7!=O(_({},"a",{get:function(){return _(this,"a",{value:7}).a}})).a})?function(t,e,n){var r=x(G,e);r&&delete G[e],_(t,e,n),r&&t!==G&&_(G,e,r)}:_,F=function(t){var e=P[t]=O(L.prototype);return e._k=t,a&&N&&C(G,t,{configurable:!0,set:function(e){i(this,j)&&i(this[j],t)&&(this[j][t]=!1),C(this,t,S(1,e))}}),e},U=function(t){return"symbol"==typeof t},R=function(t,e,n){return n&&i(P,e)?(n.enumerable?(i(t,j)&&t[j][e]&&(t[j][e]=!1),n=O(n,{enumerable:S(0,!1)})):(i(t,j)||_(t,j,S(1,{})),t[j][e]=!0),C(t,e,n)):_(t,e,n)},M=function(t,e){m(t);for(var n,r=v(e=b(e)),o=0,i=r.length;i>o;)R(t,n=r[o++],e[n]);return t},D=function(t,e){return void 0===e?O(t):M(O(t),e)},B=function(t){var e=I.call(this,t);return!(e||!i(this,t)||!i(P,t)||i(this,j)&&this[j][t])||e},q=function(t,e){var n=x(t=b(t),e);return!n||!i(P,e)||i(t,j)&&t[j][e]||(n.enumerable=!0),n},W=function(t){for(var e,n=w(b(t)),r=[],o=0;n.length>o;)i(P,e=n[o++])||e==j||r.push(e);return r},K=function(t){for(var e,n=w(b(t)),r=[],o=0;n.length>o;)i(P,e=n[o++])&&r.push(P[e]);return r},J=s(function(){var t=L();return"[null]"!=k([t])||"{}"!=k({a:t})||"{}"!=k(Object(t))});T||(c((L=function(){if(U(this))throw TypeError("Symbol is not a constructor");return F(p(arguments.length>0?arguments[0]:void 0))}).prototype,"toString",function(){return this._k}),U=function(t){return t instanceof L},r.create=D,r.isEnum=B,r.getDesc=q,r.setDesc=R,r.setDescs=M,r.getNames=g.get=W,r.getSymbols=K,a&&!n(31)&&c(G,"propertyIsEnumerable",B,!0));var z={for:function(t){return i(A,t+="")?A[t]:A[t]=L(t)},keyFor:function(t){return h(A,t)},useSetter:function(){N=!0},useSimple:function(){N=!1}};r.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=d(t);z[t]=T?e:F(e)}),N=!0,u(u.G+u.W,{Symbol:L}),u(u.S,"Symbol",z),u(u.S+u.F*!T,"Object",{create:D,defineProperty:R,defineProperties:M,getOwnPropertyDescriptor:q,getOwnPropertyNames:W,getOwnPropertySymbols:K}),E&&u(u.S+u.F*(!T||J),"JSON",{stringify:function(t){if(void 0!==t&&!U(t)){for(var e,n,r=[t],o=1,i=arguments;i.length>o;)r.push(i[o++]);return"function"==typeof(e=r[1])&&(n=e),!n&&y(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!U(e))return e}),r[1]=e,k.apply(E,r)}}}),f(L,"Symbol"),f(Math,"Math",!0),f(o.JSON,"JSON",!0)},function(t,e,n){var r=n(2),o=n(17);t.exports=function(t,e){for(var n,i=o(t),a=r.getKeys(i),u=a.length,c=0;u>c;)if(i[n=a[c++]]===e)return n}},function(t,e,n){var r=n(17),o=n(2).getNames,i={}.toString,a="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[];t.exports.get=function(t){return a&&"[object Window]"==i.call(t)?function(t){try{return o(t)}catch(t){return a.slice()}}(t):o(r(t))}},function(t,e,n){var r=n(2);t.exports=function(t){var e=r.getKeys(t),n=r.getSymbols;if(n)for(var o,i=n(t),a=r.isEnum,u=0;i.length>u;)a.call(t,o=i[u++])&&e.push(o);return e}},function(t,e,n){var r=n(37);t.exports=Array.isArray||function(t){return"Array"==r(t)}},function(t,e){},function(t,e,n){t.exports={default:n(68),__esModule:!0}},function(t,e,n){n(69),t.exports=n(4).Object.setPrototypeOf},function(t,e,n){var r=n(12);r(r.S,"Object",{setPrototypeOf:n(70).set})},function(t,e,n){var r=n(2).getDesc,o=n(39),i=n(38),a=function(t,e){if(i(t),!o(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,o){try{(o=n(28)(Function.call,r(Object.prototype,"__proto__").set,2))(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return a(t,n),e?t.__proto__=n:o(t,n),t}}({},!1):void 0),check:a}},function(t,e,n){t.exports={default:n(72),__esModule:!0}},function(t,e,n){var r=n(2);t.exports=function(t,e){return r.create(t,e)}},function(t,e){t.exports=u},function(t,e){},function(t,e,n){n(76),t.exports=n(4).Object.assign},function(t,e,n){var r=n(12);r(r.S+r.F,"Object",{assign:n(77)})},function(t,e,n){var r=n(2),o=n(18),i=n(36);t.exports=n(16)(function(){var t=Object.assign,e={},n={},r=Symbol(),o="abcdefghijklmnopqrst";return e[r]=7,o.split("").forEach(function(t){n[t]=t}),7!=t({},e)[r]||Object.keys(t({},n)).join("")!=o})?function(t,e){for(var n=o(t),a=arguments,u=a.length,c=1,s=r.getKeys,l=r.getSymbols,f=r.isEnum;u>c;)for(var p,d=i(a[c++]),h=l?s(d).concat(l(d)):s(d),g=h.length,v=0;g>v;)f.call(d,p=h[v++])&&(n[p]=d[p]);return n}:Object.assign},function(t,e,n){var r=function(){return this}()||Function("return this")(),o=r.regeneratorRuntime&&Object.getOwnPropertyNames(r).indexOf("regeneratorRuntime")>=0,i=o&&r.regeneratorRuntime;if(r.regeneratorRuntime=void 0,t.exports=n(79),o)r.regeneratorRuntime=i;else try{delete r.regeneratorRuntime}catch(t){r.regeneratorRuntime=void 0}},function(t,e){!function(e){"use strict";var n,r=Object.prototype,o=r.hasOwnProperty,i="function"==typeof Symbol?Symbol:{},a=i.iterator||"@@iterator",u=i.asyncIterator||"@@asyncIterator",c=i.toStringTag||"@@toStringTag",s="object"==typeof t,l=e.regeneratorRuntime;if(l)s&&(t.exports=l);else{(l=e.regeneratorRuntime=s?t.exports:{}).wrap=S;var f="suspendedStart",p="suspendedYield",d="executing",h="completed",g={},v={};v[a]=function(){return this};var y=Object.getPrototypeOf,m=y&&y(y(A([])));m&&m!==r&&o.call(m,a)&&(v=m);var b=w.prototype=_.prototype=Object.create(v);O.prototype=b.constructor=w,w.constructor=O,w[c]=O.displayName="GeneratorFunction",l.isGeneratorFunction=function(t){var e="function"==typeof t&&t.constructor;return!!e&&(e===O||"GeneratorFunction"===(e.displayName||e.name))},l.mark=function(t){return Object.setPrototypeOf?Object.setPrototypeOf(t,w):(t.__proto__=w,c in t||(t[c]="GeneratorFunction")),t.prototype=Object.create(b),t},l.awrap=function(t){return{__await:t}},L(E.prototype),E.prototype[u]=function(){return this},l.AsyncIterator=E,l.async=function(t,e,n,r){var o=new E(S(t,e,n,r));return l.isGeneratorFunction(e)?o:o.next().then(function(t){return t.done?t.value:o.next()})},L(b),b[c]="Generator",b[a]=function(){return this},b.toString=function(){return"[object Generator]"},l.keys=function(t){var e=[];for(var n in t)e.push(n);return e.reverse(),function n(){for(;e.length;){var r=e.pop();if(r in t)return n.value=r,n.done=!1,n}return n.done=!0,n}},l.values=A,I.prototype={constructor:I,reset:function(t){if(this.prev=0,this.next=0,this.sent=this._sent=n,this.done=!1,this.delegate=null,this.method="next",this.arg=n,this.tryEntries.forEach(j),!t)for(var e in this)"t"===e.charAt(0)&&o.call(this,e)&&!isNaN(+e.slice(1))&&(this[e]=n)},stop:function(){this.done=!0;var t=this.tryEntries[0].completion;if("throw"===t.type)throw t.arg;return this.rval},dispatchException:function(t){if(this.done)throw t;var e=this;function r(r,o){return u.type="throw",u.arg=t,e.next=r,o&&(e.method="next",e.arg=n),!!o}for(var i=this.tryEntries.length-1;i>=0;--i){var a=this.tryEntries[i],u=a.completion;if("root"===a.tryLoc)return r("end");if(a.tryLoc<=this.prev){var c=o.call(a,"catchLoc"),s=o.call(a,"finallyLoc");if(c&&s){if(this.prev<a.catchLoc)return r(a.catchLoc,!0);if(this.prev<a.finallyLoc)return r(a.finallyLoc)}else if(c){if(this.prev<a.catchLoc)return r(a.catchLoc,!0)}else{if(!s)throw new Error("try statement without catch or finally");if(this.prev<a.finallyLoc)return r(a.finallyLoc)}}}},abrupt:function(t,e){for(var n=this.tryEntries.length-1;n>=0;--n){var r=this.tryEntries[n];if(r.tryLoc<=this.prev&&o.call(r,"finallyLoc")&&this.prev<r.finallyLoc){var i=r;break}}i&&("break"===t||"continue"===t)&&i.tryLoc<=e&&e<=i.finallyLoc&&(i=null);var a=i?i.completion:{};return a.type=t,a.arg=e,i?(this.method="next",this.next=i.finallyLoc,g):this.complete(a)},complete:function(t,e){if("throw"===t.type)throw t.arg;return"break"===t.type||"continue"===t.type?this.next=t.arg:"return"===t.type?(this.rval=this.arg=t.arg,this.method="return",this.next="end"):"normal"===t.type&&e&&(this.next=e),g},finish:function(t){for(var e=this.tryEntries.length-1;e>=0;--e){var n=this.tryEntries[e];if(n.finallyLoc===t)return this.complete(n.completion,n.afterLoc),j(n),g}},catch:function(t){for(var e=this.tryEntries.length-1;e>=0;--e){var n=this.tryEntries[e];if(n.tryLoc===t){var r=n.completion;if("throw"===r.type){var o=r.arg;j(n)}return o}}throw new Error("illegal catch attempt")},delegateYield:function(t,e,r){return this.delegate={iterator:A(t),resultName:e,nextLoc:r},"next"===this.method&&(this.arg=n),g}}}function S(t,e,n,r){var o=e&&e.prototype instanceof _?e:_,i=Object.create(o.prototype),a=new I(r||[]);return i._invoke=function(t,e,n){var r=f;return function(o,i){if(r===d)throw new Error("Generator is already running");if(r===h){if("throw"===o)throw i;return P()}for(n.method=o,n.arg=i;;){var a=n.delegate;if(a){var u=k(a,n);if(u){if(u===g)continue;return u}}if("next"===n.method)n.sent=n._sent=n.arg;else if("throw"===n.method){if(r===f)throw r=h,n.arg;n.dispatchException(n.arg)}else"return"===n.method&&n.abrupt("return",n.arg);r=d;var c=x(t,e,n);if("normal"===c.type){if(r=n.done?h:p,c.arg===g)continue;return{value:c.arg,done:n.done}}"throw"===c.type&&(r=h,n.method="throw",n.arg=c.arg)}}}(t,n,a),i}function x(t,e,n){try{return{type:"normal",arg:t.call(e,n)}}catch(t){return{type:"throw",arg:t}}}function _(){}function O(){}function w(){}function L(t){["next","throw","return"].forEach(function(e){t[e]=function(t){return this._invoke(e,t)}})}function E(t){var e;this._invoke=function(n,r){function i(){return new Promise(function(e,i){!function e(n,r,i,a){var u=x(t[n],t,r);if("throw"!==u.type){var c=u.arg,s=c.value;return s&&"object"==typeof s&&o.call(s,"__await")?Promise.resolve(s.__await).then(function(t){e("next",t,i,a)},function(t){e("throw",t,i,a)}):Promise.resolve(s).then(function(t){c.value=t,i(c)},a)}a(u.arg)}(n,r,e,i)})}return e=e?e.then(i,i):i()}}function k(t,e){var r=t.iterator[e.method];if(r===n){if(e.delegate=null,"throw"===e.method){if(t.iterator.return&&(e.method="return",e.arg=n,k(t,e),"throw"===e.method))return g;e.method="throw",e.arg=new TypeError("The iterator does not provide a 'throw' method")}return g}var o=x(r,t.iterator,e.arg);if("throw"===o.type)return e.method="throw",e.arg=o.arg,e.delegate=null,g;var i=o.arg;return i?i.done?(e[t.resultName]=i.value,e.next=t.nextLoc,"return"!==e.method&&(e.method="next",e.arg=n),e.delegate=null,g):i:(e.method="throw",e.arg=new TypeError("iterator result is not an object"),e.delegate=null,g)}function N(t){var e={tryLoc:t[0]};1 in t&&(e.catchLoc=t[1]),2 in t&&(e.finallyLoc=t[2],e.afterLoc=t[3]),this.tryEntries.push(e)}function j(t){var e=t.completion||{};e.type="normal",delete e.arg,t.completion=e}function I(t){this.tryEntries=[{tryLoc:"root"}],t.forEach(N,this),this.reset(!0)}function A(t){if(t){var e=t[a];if(e)return e.call(t);if("function"==typeof t.next)return t;if(!isNaN(t.length)){var r=-1,i=function e(){for(;++r<t.length;)if(o.call(t,r))return e.value=t[r],e.done=!1,e;return e.value=n,e.done=!0,e};return i.next=i}}return{next:P}}function P(){return{value:n,done:!0}}}(function(){return this}()||Function("return this")())},function(t,e,n){"use strict";n.r(e);var r=n(40),o=n.n(r),i=n(0),a=n.n(i),u=n(6),c=n.n(u),s=n(7),l=n.n(s),f=n(8),p=n.n(f),d=n(9),h=n.n(d),g=n(10),v=n.n(g),y=n(3),m=function(t){function e(t){l()(this,e);var n=h()(this,(e.__proto__||c()(e)).call(this,t));console.log("costructor of login web"),n.state={email:"",password:""},n.handleLogin=n.handleLogin.bind(n),n.handleChange=n.handleChange.bind(n);return t.realm&&"?Realm="+t.realm,n}return v()(e,t),p()(e,[{key:"handleChange",value:function(t){var e={};e[t.target.name]=t.target.value,this.setState(e)}},{key:"handleLogin",value:function(){this.props.handleLogin(this.state.email,this.state.password)}},{key:"render",value:function(){return console.log("login ui",this.props),this.props.renderLogin(this.state,this.handleChange,this.handleLogin,this.oauthLogin,this.props)}}]),e}(a.a.Component);m.propTypes={handleOauthLogin:y.func.isRequired,handleLogin:y.func.isRequired},m.contextTypes={uikit:y.object};var b=n(41),S=n.n(b),x=n(11),_={LOGIN:"LOGIN",LOGGING_IN:"LOGGING_IN",LOGIN_SUCCESS:"LOGIN_SUCCESS",LOGIN_FAILURE:"LOGIN_FAILURE",LOGOUT:"LOGOUT",LOGOUT_SUCCESS:"LOGOUT_SUCCESS"},O=n(1),w=n(3),L=n.n(w),E=Object(x.connect)(function(t,e){return{realm:Application.Security.realm,renderLogin:e.renderLogin,signup:e.signup}},function(t,e){console.log("map dispatch of login compoent");var n="";return Application.Security.realm&&(n=Application.Security.realm),{handleLogin:function(e,r){var o={Username:e,Password:S()(r),Realm:n},i={serviceName:Application.Security.loginService};t(Object(O.createAction)(_.LOGIN,o,i))},handleOauthLogin:function(e){t(Object(O.createAction)(_.LOGIN_SUCCESS,{userId:e.id,token:e.token,permissions:e.permissions}))}}})(m);E.propTypes={loginService:L.a.string.isRequired,successpage:L.a.string,realm:L.a.string,signup:L.a.string};n(73);var k=n(3),N=function(t){function e(t){l()(this,e);var n=h()(this,(e.__proto__||c()(e)).call(this,t));return n.validatetoken=n.validatetoken.bind(n),n.state={loggedIn:t.loggedIn,validation:t.validation},t.validation&&n.validatetoken(),n}return v()(e,t),p()(e,[{key:"componentWillReceiveProps",value:function(t){t.loggedIn==this.state.loggedIn&&t.validation==this.state.validation||this.setState({loggedIn:t.loggedIn,validation:t.validation})}},{key:"validatetoken",value:function(){var t=this,e=this.props.logout,n=this.props.login,r=O.RequestBuilder.DefaultRequest({},{});O.DataSource.ExecuteService(this.props.validateService,r).then(function(t){n(t.data.Id,t.data.Permissions)},function(n){e(),t.setState({loggedIn:!1,validation:!1})})}},{key:"getChildContext",value:function(){return{loggedIn:this.state.loggedIn}}},{key:"render",value:function(){return this.state.validation?null:this.props.children?a.a.cloneElement(this.props.children,{loggedIn:this.state.loggedIn,validation:this.state.validation}):null}}]),e}(a.a.Component);N.childContextTypes={loggedIn:k.bool,user:k.object};var j=Object(x.connect)(function(t,e){return null==Storage.auth?{validation:!1,loggedIn:!1,validateService:e.validateService}:""!=Storage.auth?"LoggedIn"!=t.Security.status?{validation:!0,loggedIn:!1,validateService:e.validateService}:{validation:!1,loggedIn:!0,validateService:e.validateService}:void 0},function(t,e){return{login:function(e,n){t(Object(O.createAction)(_.LOGIN_SUCCESS,{userId:e,token:Storage.auth,user:Storage.user,permissions:n}))},logout:function(){t(Object(O.createAction)(_.LOGOUT,null,null))}}})(N),I=n(14);n(74);function A(t,e,n){return function(r,o,i,u,c){return console.log("renderLogin",n,"uikit",t,"settigs",e,"props",c),a.a.createElement("div",{className:c.className?c.className:" loginbox "},a.a.createElement("div",{className:"logintext"},n.loginForm.formtext),a.a.createElement("div",{className:"sociallogin"},a.a.createElement(I.Action,{widget:"button",method:function(){u(Application.Security.googleAuthUrl)},name:"googleAuth",className:"googleAuthAction"},n.loginForm.google)),a.a.createElement("div",{className:"separator"},n.loginForm.separator),a.a.createElement("div",{className:"main"},a.a.createElement(t.Form,{role:"form"},a.a.createElement("div",{className:"userfield"},a.a.createElement("label",{htmlFor:"email"},n.loginForm.userlabel),a.a.createElement(t.TextField,{className:"text",name:"email",value:r.email,placeholder:n.loginForm.userplaceholder,onChange:o})),a.a.createElement("div",{className:"passwordfield"},a.a.createElement("label",{htmlFor:"inputPassword"},n.loginForm.passwordlabel),a.a.createElement(t.TextField,{type:"password",className:"text",name:"password",value:r.password,placeholder:n.loginForm.passwordplaceholder,onChange:o})),a.a.createElement("a",{className:"pull-right",href:"#"},"Forgot password?"),a.a.createElement("div",{className:"checkbox"},a.a.createElement("label",null,a.a.createElement("input",{type:"checkbox"}),"Remember me")),a.a.createElement("div",{className:"actionbuttons"},a.a.createElement(I.Action,{widget:"button",className:"loginBtn",name:"loginAction",method:i},n.loginForm.loginBtnText)))))}}var P=n(25),T=n.n(P),G={status:"NotLogged",token:"",userId:"",permissions:[]};Application.Register("Reducers","Security",function(t,e){if(e.type)switch(e.type){case _.LOGGING_IN:return T()({},t,{status:"LoggingIn"});case _.LOGIN_SUCCESS:return t.authToken===e.payload.token?t:(Storage.auth=e.payload.token,Storage.permissions=e.payload.permissions,Storage.userId=e.payload.userId,Storage.userFullName=e.payload.user.Name,Storage.userName=e.payload.user.Username,Storage.email=e.payload.user.Email,Storage.user=e.payload.user,T()({},t,{status:"LoggedIn",authToken:e.payload.token,userId:e.payload.userId,permissions:e.payload.permissions}));case _.LOGIN_FAILURE:case _.LOGOUT_SUCCESS:return Storage.auth="",Storage.permissions=[],Storage.userId="",Storage.userName="",Storage.userFullName="",Storage.email="",Storage.user=null,G;default:return t||G}});var C=n(5),F=n.n(C),U=n(26),R=(Object.assign,"function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t}),M=function(t){return"@@redux-saga/"+t},D=M("TASK"),B=M("HELPER");function q(t,e,n){if(!e(t))throw function(t,e){var n=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"";"undefined"==typeof window?console.log("redux-saga "+t+": "+e+"\n"+(n&&n.stack||n)):console[t](e,n)}("error","uncaught at check",n),new Error(n)}var W=Object.prototype.hasOwnProperty;function K(t,e){return J.notUndef(t)&&W.call(t,e)}var J={undef:function(t){return null===t||void 0===t},notUndef:function(t){return null!==t&&void 0!==t},func:function(t){return"function"==typeof t},number:function(t){return"number"==typeof t},string:function(t){return"string"==typeof t},array:Array.isArray,object:function(t){return t&&!J.array(t)&&"object"===(void 0===t?"undefined":R(t))},promise:function(t){return t&&J.func(t.then)},iterator:function(t){return t&&J.func(t.next)&&J.func(t.throw)},iterable:function(t){return t&&J.func(Symbol)?J.func(t[Symbol.iterator]):J.array(t)},task:function(t){return t&&t[D]},observable:function(t){return t&&J.func(t.subscribe)},buffer:function(t){return t&&J.func(t.isEmpty)&&J.func(t.take)&&J.func(t.put)},pattern:function(t){return t&&(J.string(t)||"symbol"===(void 0===t?"undefined":R(t))||J.func(t)||J.array(t))},channel:function(t){return t&&J.func(t.take)&&J.func(t.close)},helper:function(t){return t&&t[B]},stringableFunc:function(t){return J.func(t)&&K(t,"toString")}};function z(t,e){return function(){return t.apply(void 0,arguments)}}var H=M("IO"),Y="TAKE",V="PUT",Q="CALL",X=function(t,e){var n;return(n={})[H]=!0,n[t]=e,n};function Z(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"*";if(arguments.length&&q(arguments[0],J.notUndef,"take(patternOrChannel): patternOrChannel is undefined"),J.pattern(t))return X(Y,{pattern:t});if(J.channel(t))return X(Y,{channel:t});throw new Error("take(patternOrChannel): argument "+String(t)+" is not valid channel or a valid pattern")}Z.maybe=function(){var t=Z.apply(void 0,arguments);return t[Y].maybe=!0,t};Z.maybe;function $(t,e){return arguments.length>1?(q(t,J.notUndef,"put(channel, action): argument channel is undefined"),q(t,J.channel,"put(channel, action): argument "+t+" is not a valid channel"),q(e,J.notUndef,"put(channel, action): argument action is undefined")):(q(t,J.notUndef,"put(action): argument action is undefined"),e=t,t=null),X(V,{channel:t,action:e})}function tt(t,e,n){q(e,J.notUndef,t+": argument fn is undefined");var r=null;if(J.array(e)){var o=e;r=o[0],e=o[1]}else if(e.fn){var i=e;r=i.context,e=i.fn}return r&&J.string(e)&&J.func(r[e])&&(e=r[e]),q(e,J.func,t+": argument "+e+" is not a function"),{context:r,fn:e,args:n}}function et(t){for(var e=arguments.length,n=Array(e>1?e-1:0),r=1;r<e;r++)n[r-1]=arguments[r];return X(Q,tt("call",t,n))}$.resolve=function(){var t=$.apply(void 0,arguments);return t[V].resolve=!0,t},$.sync=z($.resolve);Object.assign;var nt=F.a.mark(it),rt=F.a.mark(at),ot=F.a.mark(ut);function it(t){var e,n,r,o,i,a,u,c;return F.a.wrap(function(s){for(;;)switch(s.prev=s.next){case 0:return s.prev=0,s.next=3,$(Object(O.createAction)(_.LOGGING_IN));case 3:return e=O.RequestBuilder.DefaultRequest(null,t.payload),s.next=6,et(O.DataSource.ExecuteService,t.meta.serviceName,e);case 6:return n=s.sent,r=Application.Security.AuthToken.toLowerCase(),o=n.info[r],i=n.data,a=n.data.Id,u=n.data.Permissions,c=Object(O.createAction)(_.LOGIN_SUCCESS,{userId:a,token:o,permissions:u,user:i}),s.next=15,$(c);case 15:console.log("dispatched login action &&&&"),s.next=22;break;case 18:return s.prev=18,s.t0=s.catch(0),s.next=22,$(Object(O.createAction)(_.LOGIN_FAILURE,s.t0));case 22:case"end":return s.stop()}},nt,this,[[0,18]])}function at(t){return F.a.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,$(Object(O.createAction)(_.LOGOUT_SUCCESS,{}));case 2:case"end":return t.stop()}},rt,this)}function ut(){return F.a.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,[Object(U.takeLatest)(_.LOGIN,it),Object(U.takeLatest)(_.LOGOUT,at)];case 2:case"end":return t.stop()}},ot,this)}Application.Register("Sagas","loginSaga",ut);var ct=function(t){function e(){return l()(this,e),h()(this,(e.__proto__||c()(e)).apply(this,arguments))}return v()(e,t),p()(e,[{key:"render",value:function(){var t=this.props,e=t.module.properties?t.module.properties:{},n=e.logoutText?e.logoutText:"Logout";return t.loggedIn?a.a.createElement(t.uikit.Block,{className:"userblock "+t.className},a.a.createElement(t.uikit.Block,{className:"username"},Storage.userFullName?Storage.userFullName:Storage.userName),a.a.createElement(I.Action,{name:"logout",method:t.logout,className:"logout"},n)):null}}]),e}(a.a.Component),st=Object(x.connect)(function(t,e){return{loggedIn:"LoggedIn"==t.Security.status}},function(t,e){return{logout:function(){t(Object(O.createAction)(_.LOGOUT,null,null))}}})(ct);n.d(e,"Initialize",function(){return pt}),n.d(e,"WebLoginForm",function(){return dt}),n.d(e,"LoginComponent",function(){return E}),n.d(e,"renderWebLogin",function(){return A}),n.d(e,"LoginValidator",function(){return j});var lt,ft=n(3);function pt(t,e,n,r,i,a){((lt=this).properties=Application.Properties[e],lt.settings=r,0!=o()(r).length)?Application.Security={googleAuthUrl:r.googleAuthUrl,loginService:r.loginService,validateService:r.validateService,loginServiceURL:r.loginServiceURL,realm:r.realm}:(Application.Security={loginService:"login",validateService:"validate",realm:""},_reg("Services")||(Application.Register("Services","login",{url:"/login",method:"POST"}),Application.Register("Services","validate",{url:"/validate",method:"POST"})));r.AuthToken?Application.Security.AuthToken=r.AuthToken:Application.Security.AuthToken="x-auth-token"}var dt=function(t,e){return console.log("render logiform",E),a.a.createElement(E,{className:t.className,renderLogin:A(e.uikit,lt.settings,lt.properties),realm:t.realm,loginService:t.loginService,loginServiceURL:t.loginServiceURL,googleAuthUrl:t.googleAuthUrl})};dt.contextTypes={uikit:ft.object},Application.Register("Blocks","userBlock",function(t,e,n,r){return a.a.createElement(st,{className:t.className,uikit:n,module:lt})})}])});
//# sourceMappingURL=index.js.map
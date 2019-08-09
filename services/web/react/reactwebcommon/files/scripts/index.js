define("reactwebcommon",["react","uicommon","sanitize-html","react-redux"],function(t,e,n,o){return function(t){var e={};function n(o){if(e[o])return e[o].exports;var r=e[o]={i:o,l:!1,exports:{}};return t[o].call(r.exports,r,r.exports,n),r.l=!0,r.exports}return n.m=t,n.c=e,n.d=function(t,e,o){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:o})},n.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var o=Object.create(null);if(n.r(o),Object.defineProperty(o,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var r in t)n.d(o,r,function(e){return t[e]}.bind(null,r));return o},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=78)}([function(e,n){e.exports=t},function(t,e,n){t.exports=n(71)()},function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(t,e,n){t.exports={default:n(42),__esModule:!0}},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){"use strict";e.__esModule=!0;var o,r=n(45),i=(o=r)&&o.__esModule?o:{default:o};e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var o=e[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),(0,i.default)(t,o.key,o)}}return function(e,n,o){return n&&t(e.prototype,n),o&&t(e,o),e}}()},function(t,e,n){"use strict";e.__esModule=!0;var o,r=n(25),i=(o=r)&&o.__esModule?o:{default:o};e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,i.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){"use strict";e.__esModule=!0;var o=s(n(65)),r=s(n(69)),i=s(n(25));function s(t){return t&&t.__esModule?t:{default:t}}e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,i.default)(e)));t.prototype=(0,r.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(o.default?(0,o.default)(t,e):t.__proto__=e)}},function(t,n){t.exports=e},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,e,n){var o=n(12),r=n(9),i=n(21),s=function(t,e,n){var c,a,u,l=t&s.F,f=t&s.G,p=t&s.S,h=t&s.P,d=t&s.B,y=t&s.W,m=f?r:r[e]||(r[e]={}),v=f?o:p?o[e]:(o[e]||{}).prototype;for(c in f&&(n=e),n)(a=!l&&v&&c in v)&&c in m||(u=a?v[c]:n[c],m[c]=f&&"function"!=typeof v[c]?n[c]:d&&a?i(u,o):y&&v[c]==u?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(u):h&&"function"==typeof u?i(Function.call,u):u,h&&((m.prototype||(m.prototype={}))[c]=u))};s.F=1,s.G=2,s.S=4,s.P=8,s.B=16,s.W=32,t.exports=s},function(t,e,n){var o=n(30)("wks"),r=n(31),i=n(12).Symbol;t.exports=function(t){return o[t]||(o[t]=i&&i[t]||(i||r)("Symbol."+t))}},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e,n){var o=n(23),r=n(15);t.exports=function(t){return o(r(t))}},function(t,e){t.exports=function(t){if(null==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e,n){var o=n(2),r=n(17);t.exports=n(29)?function(t,e,n){return o.setDesc(t,e,r(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports={}},function(t,e,n){var o=n(2).setDesc,r=n(18),i=n(11)("toStringTag");t.exports=function(t,e,n){t&&!r(t=n?t:t.prototype,i)&&o(t,i,{configurable:!0,value:e})}},function(t,e,n){var o=n(40);t.exports=function(t,e,n){if(o(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,o){return t.call(e,n,o)};case 3:return function(n,o,r){return t.call(e,n,o,r)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){var o=n(15);t.exports=function(t){return Object(o(t))}},function(t,e,n){var o=n(24);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==o(t)?t.split(""):Object(t)}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){"use strict";e.__esModule=!0;var o=s(n(47)),r=s(n(57)),i="function"==typeof r.default&&"symbol"==typeof o.default?function(t){return typeof t}:function(t){return t&&"function"==typeof r.default&&t.constructor===r.default&&t!==r.default.prototype?"symbol":typeof t};function s(t){return t&&t.__esModule?t:{default:t}}e.default="function"==typeof r.default&&"symbol"===i(o.default)?function(t){return void 0===t?"undefined":i(t)}:function(t){return t&&"function"==typeof r.default&&t.constructor===r.default&&t!==r.default.prototype?"symbol":void 0===t?"undefined":i(t)}},function(t,e,n){"use strict";var o=n(27),r=n(10),i=n(28),s=n(16),c=n(18),a=n(19),u=n(52),l=n(20),f=n(2).getProto,p=n(11)("iterator"),h=!([].keys&&"next"in[].keys()),d=function(){return this};t.exports=function(t,e,n,y,m,v,b){u(n,e,y);var _,g,w=function(t){if(!h&&t in k)return k[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},x=e+" Iterator",S="values"==m,O=!1,k=t.prototype,P=k[p]||k["@@iterator"]||m&&k[m],j=P||w(m);if(P){var M=f(j.call(new t));l(M,x,!0),!o&&c(k,"@@iterator")&&s(M,p,d),S&&"values"!==P.name&&(O=!0,j=function(){return P.call(this)})}if(o&&!b||!h&&!O&&k[p]||s(k,p,j),a[e]=j,a[x]=d,m)if(_={values:S?j:w("values"),keys:v?j:w("keys"),entries:S?w("entries"):j},b)for(g in _)g in k||i(k,g,_[g]);else r(r.P+r.F*(h||O),e,_);return _}},function(t,e){t.exports=!0},function(t,e,n){t.exports=n(16)},function(t,e,n){t.exports=!n(13)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){var o=n(12),r=o["__core-js_shared__"]||(o["__core-js_shared__"]={});t.exports=function(t){return r[t]||(r[t]={})}},function(t,e){var n=0,o=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+o).toString(36))}},function(t,e,n){var o=n(33);t.exports=function(t){if(!o(t))throw TypeError(t+" is not an object!");return t}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){"use strict";e.__esModule=!0;var o,r=n(37),i=(o=r)&&o.__esModule?o:{default:o};e.default=i.default||function(t){for(var e=1;e<arguments.length;e++){var n=arguments[e];for(var o in n)Object.prototype.hasOwnProperty.call(n,o)&&(t[o]=n[o])}return t}},function(t,e){t.exports=n},function(t,e){t.exports=o},function(t,e,n){t.exports={default:n(38),__esModule:!0}},function(t,e,n){n(39),t.exports=n(9).Object.assign},function(t,e,n){var o=n(10);o(o.S+o.F,"Object",{assign:n(41)})},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){var o=n(2),r=n(22),i=n(23);t.exports=n(13)(function(){var t=Object.assign,e={},n={},o=Symbol(),r="abcdefghijklmnopqrst";return e[o]=7,r.split("").forEach(function(t){n[t]=t}),7!=t({},e)[o]||Object.keys(t({},n)).join("")!=r})?function(t,e){for(var n=r(t),s=arguments,c=s.length,a=1,u=o.getKeys,l=o.getSymbols,f=o.isEnum;c>a;)for(var p,h=i(s[a++]),d=l?u(h).concat(l(h)):u(h),y=d.length,m=0;y>m;)f.call(h,p=d[m++])&&(n[p]=h[p]);return n}:Object.assign},function(t,e,n){n(43),t.exports=n(9).Object.getPrototypeOf},function(t,e,n){var o=n(22);n(44)("getPrototypeOf",function(t){return function(e){return t(o(e))}})},function(t,e,n){var o=n(10),r=n(9),i=n(13);t.exports=function(t,e){var n=(r.Object||{})[t]||Object[t],s={};s[t]=e(n),o(o.S+o.F*i(function(){n(1)}),"Object",s)}},function(t,e,n){t.exports={default:n(46),__esModule:!0}},function(t,e,n){var o=n(2);t.exports=function(t,e,n){return o.setDesc(t,e,n)}},function(t,e,n){t.exports={default:n(48),__esModule:!0}},function(t,e,n){n(49),n(53),t.exports=n(11)("iterator")},function(t,e,n){"use strict";var o=n(50)(!0);n(26)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=o(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){var o=n(51),r=n(15);t.exports=function(t){return function(e,n){var i,s,c=String(r(e)),a=o(n),u=c.length;return a<0||a>=u?t?"":void 0:(i=c.charCodeAt(a))<55296||i>56319||a+1===u||(s=c.charCodeAt(a+1))<56320||s>57343?t?c.charAt(a):i:t?c.slice(a,a+2):s-56320+(i-55296<<10)+65536}}},function(t,e){var n=Math.ceil,o=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?o:n)(t)}},function(t,e,n){"use strict";var o=n(2),r=n(17),i=n(20),s={};n(16)(s,n(11)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=o.create(s,{next:r(1,n)}),i(t,e+" Iterator")}},function(t,e,n){n(54);var o=n(19);o.NodeList=o.HTMLCollection=o.Array},function(t,e,n){"use strict";var o=n(55),r=n(56),i=n(19),s=n(14);t.exports=n(26)(Array,"Array",function(t,e){this._t=s(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,r(1)):r(0,"keys"==e?n:"values"==e?t[n]:[n,t[n]])},"values"),i.Arguments=i.Array,o("keys"),o("values"),o("entries")},function(t,e){t.exports=function(){}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e,n){t.exports={default:n(58),__esModule:!0}},function(t,e,n){n(59),n(64),t.exports=n(9).Symbol},function(t,e,n){"use strict";var o=n(2),r=n(12),i=n(18),s=n(29),c=n(10),a=n(28),u=n(13),l=n(30),f=n(20),p=n(31),h=n(11),d=n(60),y=n(61),m=n(62),v=n(63),b=n(32),_=n(14),g=n(17),w=o.getDesc,x=o.setDesc,S=o.create,O=y.get,k=r.Symbol,P=r.JSON,j=P&&P.stringify,M=!1,E=h("_hidden"),N=o.isEnum,T=l("symbol-registry"),C=l("symbols"),D="function"==typeof k,F=Object.prototype,A=s&&u(function(){return 7!=S(x({},"a",{get:function(){return x(this,"a",{value:7}).a}})).a})?function(t,e,n){var o=w(F,e);o&&delete F[e],x(t,e,n),o&&t!==F&&x(F,e,o)}:x,I=function(t){var e=C[t]=S(k.prototype);return e._k=t,s&&M&&A(F,t,{configurable:!0,set:function(e){i(this,E)&&i(this[E],t)&&(this[E][t]=!1),A(this,t,g(1,e))}}),e},R=function(t){return"symbol"==typeof t},L=function(t,e,n){return n&&i(C,e)?(n.enumerable?(i(t,E)&&t[E][e]&&(t[E][e]=!1),n=S(n,{enumerable:g(0,!1)})):(i(t,E)||x(t,E,g(1,{})),t[E][e]=!0),A(t,e,n)):x(t,e,n)},W=function(t,e){b(t);for(var n,o=m(e=_(e)),r=0,i=o.length;i>r;)L(t,n=o[r++],e[n]);return t},B=function(t,e){return void 0===e?S(t):W(S(t),e)},q=function(t){var e=N.call(this,t);return!(e||!i(this,t)||!i(C,t)||i(this,E)&&this[E][t])||e},H=function(t,e){var n=w(t=_(t),e);return!n||!i(C,e)||i(t,E)&&t[E][e]||(n.enumerable=!0),n},U=function(t){for(var e,n=O(_(t)),o=[],r=0;n.length>r;)i(C,e=n[r++])||e==E||o.push(e);return o},V=function(t){for(var e,n=O(_(t)),o=[],r=0;n.length>r;)i(C,e=n[r++])&&o.push(C[e]);return o},J=u(function(){var t=k();return"[null]"!=j([t])||"{}"!=j({a:t})||"{}"!=j(Object(t))});D||(a((k=function(){if(R(this))throw TypeError("Symbol is not a constructor");return I(p(arguments.length>0?arguments[0]:void 0))}).prototype,"toString",function(){return this._k}),R=function(t){return t instanceof k},o.create=B,o.isEnum=q,o.getDesc=H,o.setDesc=L,o.setDescs=W,o.getNames=y.get=U,o.getSymbols=V,s&&!n(27)&&a(F,"propertyIsEnumerable",q,!0));var K={for:function(t){return i(T,t+="")?T[t]:T[t]=k(t)},keyFor:function(t){return d(T,t)},useSetter:function(){M=!0},useSimple:function(){M=!1}};o.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=h(t);K[t]=D?e:I(e)}),M=!0,c(c.G+c.W,{Symbol:k}),c(c.S,"Symbol",K),c(c.S+c.F*!D,"Object",{create:B,defineProperty:L,defineProperties:W,getOwnPropertyDescriptor:H,getOwnPropertyNames:U,getOwnPropertySymbols:V}),P&&c(c.S+c.F*(!D||J),"JSON",{stringify:function(t){if(void 0!==t&&!R(t)){for(var e,n,o=[t],r=1,i=arguments;i.length>r;)o.push(i[r++]);return"function"==typeof(e=o[1])&&(n=e),!n&&v(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!R(e))return e}),o[1]=e,j.apply(P,o)}}}),f(k,"Symbol"),f(Math,"Math",!0),f(r.JSON,"JSON",!0)},function(t,e,n){var o=n(2),r=n(14);t.exports=function(t,e){for(var n,i=r(t),s=o.getKeys(i),c=s.length,a=0;c>a;)if(i[n=s[a++]]===e)return n}},function(t,e,n){var o=n(14),r=n(2).getNames,i={}.toString,s="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[];t.exports.get=function(t){return s&&"[object Window]"==i.call(t)?function(t){try{return r(t)}catch(t){return s.slice()}}(t):r(o(t))}},function(t,e,n){var o=n(2);t.exports=function(t){var e=o.getKeys(t),n=o.getSymbols;if(n)for(var r,i=n(t),s=o.isEnum,c=0;i.length>c;)s.call(t,r=i[c++])&&e.push(r);return e}},function(t,e,n){var o=n(24);t.exports=Array.isArray||function(t){return"Array"==o(t)}},function(t,e){},function(t,e,n){t.exports={default:n(66),__esModule:!0}},function(t,e,n){n(67),t.exports=n(9).Object.setPrototypeOf},function(t,e,n){var o=n(10);o(o.S,"Object",{setPrototypeOf:n(68).set})},function(t,e,n){var o=n(2).getDesc,r=n(33),i=n(32),s=function(t,e){if(i(t),!r(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,r){try{(r=n(21)(Function.call,o(Object.prototype,"__proto__").set,2))(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return s(t,n),e?t.__proto__=n:r(t,n),t}}({},!1):void 0),check:s}},function(t,e,n){t.exports={default:n(70),__esModule:!0}},function(t,e,n){var o=n(2);t.exports=function(t,e){return o.create(t,e)}},function(t,e,n){"use strict";var o=n(72),r=n(73),i=n(74);t.exports=function(){function t(t,e,n,o,s,c){c!==i&&r(!1,"Calling PropTypes validators directly is not supported by the `prop-types` package. Use PropTypes.checkPropTypes() to call them. Read more at http://fb.me/use-check-prop-types")}function e(){return t}t.isRequired=t;var n={array:t,bool:t,func:t,number:t,object:t,string:t,symbol:t,any:t,arrayOf:e,element:t,instanceOf:e,node:t,objectOf:e,oneOf:e,oneOfType:e,shape:e,exact:e};return n.checkPropTypes=o,n.PropTypes=n,n}},function(t,e,n){"use strict";function o(t){return function(){return t}}var r=function(){};r.thatReturns=o,r.thatReturnsFalse=o(!1),r.thatReturnsTrue=o(!0),r.thatReturnsNull=o(null),r.thatReturnsThis=function(){return this},r.thatReturnsArgument=function(t){return t},t.exports=r},function(t,e,n){"use strict";var o=function(t){};t.exports=function(t,e,n,r,i,s,c,a){if(o(e),!t){var u;if(void 0===e)u=new Error("Minified exception occurred; use the non-minified dev environment for the full error message and additional helpful warnings.");else{var l=[n,r,i,s,c,a],f=0;(u=new Error(e.replace(/%s/g,function(){return l[f++]}))).name="Invariant Violation"}throw u.framesToPop=1,u}}},function(t,e,n){"use strict";t.exports="SECRET_DO_NOT_PASS_THIS_OR_YOU_WILL_BE_FIRED"},function(t,e){},function(t,e){},function(t,e){},function(t,e,n){"use strict";n.r(e);var o=n(34),r=n.n(o),i=n(3),s=n.n(i),c=n(4),a=n.n(c),u=n(5),l=n.n(u),f=n(6),p=n.n(f),h=n(7),d=n.n(h),y=n(0),m=n.n(y),v=function(t){function e(t){a()(this,e);var n=p()(this,(e.__proto__||s()(e)).call(this,t));return t.skipPrefix||(t.skipPrefix=!1),n}return d()(e,t),l()(e,[{key:"render",value:function(){var t=this.props.src;if(!t||0==t.length)return this.props.children?this.props.children:(console.log("No src tag provieded for image",this.props),null);!this.props.prefix||this.props.skipPrefix||this.props.src.startsWith("http")||(t=this.props.prefix+t);var e=m.a.createElement("img",r()({src:t},this.props.modifier,{className:this.props.className,style:this.props.style}));return this.props.link?m.a.createElement("a",{target:this.props.target,href:this.props.link},e):e}}]),e}(m.a.Component),b=n(35),_=n.n(b);var g=function(t){return m.a.createElement("div",{className:t.className,style:t.style,dangerouslySetInnerHTML:(e=t.children,n=t.sanitize,o=e,n&&(o=_()(e)),{__html:o})});var e,n,o},w=function(t){function e(t){a()(this,e);var n=p()(this,(e.__proto__||s()(e)).call(this,t));if(n.handleScroll=n.handleScroll.bind(n),t.windowScroll){var o=Math.floor(window.scrollY/window.innerHeight);n.state={windowNumber:o}}return n.state={scrolledOut:!1,scrolledIn:!1},n}return d()(e,t),l()(e,[{key:"handleScroll",value:function(t){if(this.props.windowScroll){var e=Math.floor(window.scrollY/window.innerHeight);e!=this.state.windowNumber&&(this.props.windowScroll(e),this.setState({windowNumber:e}))}if(this.props.onScrollEnd||this.props.onScrollIn){var n=this.refs.scrollListener,o=window.innerHeight;if(this.props.scrollEndPos&&(o=this.props.scrollEndPos),null!=n){var r=n.getBoundingClientRect();r.bottom<=o&&!this.state.scrolledOut&&(!this.state.scrolledOut&&this.props.onScrollEnd&&this.props.onScrollEnd(r.bottom),this.setState({scrolledOut:!0,scrolledIn:!1})),this.state.scrolledOut&&r.bottom>o&&(!this.state.scrolledIn&&this.props.onScrollIn&&this.props.onScrollIn(r.bottom),this.setState({scrolledOut:!1,scrolledIn:!0}))}}}},{key:"componentDidMount",value:function(){window.addEventListener("scroll",this.handleScroll)}},{key:"componentWillUnmount",value:function(){window.removeEventListener("scroll",this.handleScroll)}},{key:"render",value:function(){return m.a.createElement("div",{ref:"scrollListener",key:this.props.key,style:this.props.style,className:this.props.className},this.props.children)}}]),e}(m.a.Component),x=n(1),S=n.n(x),O=function(t,e){return _uikit.ActionButton?m.a.createElement(_uikit.ActionButton,{className:t.className+" actionbutton",onClick:t.actionFunc,btnProps:t},t.actionchildren):m.a.createElement("a",{className:t.className+" actionbutton",onClick:t.actionFunc,role:"button"},t.actionchildren)};O.propTypes={actionFunc:S.a.func.isRequired,actionchildren:S.a.oneOfType([S.a.array,S.a.string])};var k=O,P=n(36),j=function(t){return m.a.createElement("a",{className:t.className+" actionlink",href:"javascript:void(0)",onClick:t.actionFunc},t.actionchildren)};j.propTypes={actionFunc:S.a.func.isRequired,actionchildren:S.a.oneOfType([S.a.array,S.a.string])};var M=j,E=n(8),N=function(t){function e(t){a()(this,e);var n=p()(this,(e.__proto__||s()(e)).call(this,t));return console.log("action comp creation",t),n.renderView=n.renderView.bind(n),n.dispatchAction=n.dispatchAction.bind(n),n.actionFunc=n.actionFunc.bind(n),n.hasPermission=!1,null!=t.action?n.action=t.action:n.action=_reg("Actions",t.name),console.log("action",n.action),"method"==n.action.actiontype&&(n.props.method?n.actionMethod=n.props.method?n.props.method:_reg("Methods",n.action.method):n.action.method&&"function"==typeof n.action.method?n.actionMethod=n.action.method:n.actionMethod=_reg("Methods",n.action.method)),n.action&&(n.hasPermission=Object(E.hasPermission)(n.action.permission)),n}return d()(e,t),l()(e,[{key:"dispatchAction",value:function(){var t={};this.props.params&&(t=this.props.params),this.props.dispatch(Object(E.createAction)(this.action.action,t,{successCallback:this.props.successCallback,failureCallback:this.props.failureCallback}))}},{key:"actionFunc",value:function(t){if(console.log("action executed",this.props.name,this.props,this.action),t.preventDefault(),this.props.confirm&&!this.props.confirm(this.props))return!1;switch(this.action.actiontype){case"dispatchaction":return this.dispatchAction(),!1;case"method":var e=this.props.params?this.props.params:this.action.params;return this.actionMethod(e),!1;case"showdialog":var n=Window.resolvePanel("block",this.action.id);console.log("show dialog",this.action,n);var o=this.props.onClose?this.props.onClose:_reg("Methods",this.action.onClose);return Window.showDialog(this.action.title,n,o,this.action.actions,this.action.contentStyle,this.action.titleStyle),!1;case"newwindow":if(this.action.url){var r=Object(E.formatUrl)(this.action.url,this.props.params);return console.log(r),window.open(r),!1}default:if(this.action.url){var i=Object(E.formatUrl)(this.action.url,this.props.params);console.log(i),Window.redirect(i,this.action.newpage)}return!1}}},{key:"renderView",value:function(){if(!this.hasPermission)return null;var t=this.props.children?this.props.children:this.props.label,e=this.actionFunc;switch(this.props.widget){case"button":return m.a.createElement(k,{className:this.props.className,actionFunc:e,key:this.props.name+"_comp",actionchildren:t});case"component":return m.a.createElement(this.props.component,{actionFunc:e,key:this.props.name+"_comp",actionchildren:t});default:return m.a.createElement(M,{className:this.props.className,actionFunc:e,key:this.props.name+"_comp",actionchildren:t})}}},{key:"render",value:function(){return this.renderView()}}]),e}(m.a.Component);N.propTypes={name:S.a.string.isRequired};var T=Object(P.connect)()(N),C=function(t){function e(t){a()(this,e);var n=p()(this,(e.__proto__||s()(e)).call(this,t));n.actions=[];var o=t.description;console.log("action bar",t),n.description=o,n.className=t.className?t.className:o.className;var r=n;return o&&o.actions&&o.actions.forEach(function(t){r.actions.push(m.a.createElement(T,{name:t.name,label:t.label,widget:t.widget,className:" action "}))}),n}return d()(e,t),l()(e,[{key:"render",value:function(){return m.a.createElement(_uikit.Block,{className:" actionbar "+this.className},this.actions)}}]),e}(m.a.Component),D=function(t){function e(t){a()(this,e);var n=p()(this,(e.__proto__||s()(e)).call(this,t));return n.errorMethod=function(t){console.log("could not load data",t)},n.response=function(t){console.log("loadable component:---------response",t),t&&t.data&&n.dataLoaded(t.data)},t.loader&&(n.method=_reg("Methods",t.loader)),t.loadData&&(n.loadData=t.loadData),n}return d()(e,t),l()(e,[{key:"componentWillMount",value:function(){console.log("loadable component:---------",this.method,this.props);var t=this.props;if(this.loadData)if(this.method)this.method(t,this.getLoadContext?this.getLoadContext():{},this.dataLoaded);else if(t.dataService){var e=E.RequestBuilder.DefaultRequest(null,t.dataServiceParams);E.DataSource.ExecuteService(t.dataService,e).then(this.response,this.errorMethod)}else t.entity&&E.EntityData.ListEntities(t.entity).then(this.response,this.errorMethod)}}]),e}(m.a.Component);n(75),n(76),n(77);n.d(e,"ScrollListener",function(){return w}),n.d(e,"Action",function(){return T}),n.d(e,"Html",function(){return g}),n.d(e,"ActionBar",function(){return C}),n.d(e,"LoadableComponent",function(){return D}),n.d(e,"Image",function(){return v})}])});
//# sourceMappingURL=index.js.map
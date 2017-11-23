define("reactwebcommon",["react","react-redux","uicommon","sanitize-html"],function(t,n,e,r){return function(t){function n(r){if(e[r])return e[r].exports;var o=e[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}var e={};return n.m=t,n.c=e,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{configurable:!1,enumerable:!0,get:r})},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,n){return Object.prototype.hasOwnProperty.call(t,n)},n.p="/",n(n.s=35)}([function(t,n){var e=Object;t.exports={create:e.create,getProto:e.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:e.getOwnPropertyDescriptor,setDesc:e.defineProperty,setDescs:e.defineProperties,getKeys:e.keys,getNames:e.getOwnPropertyNames,getSymbols:e.getOwnPropertySymbols,each:[].forEach}},function(n,e){n.exports=t},function(t,n){var e=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=e)},function(t,n,e){var r=e(5),o=e(2),i=e(20),c=function(t,n,e){var s,u,a,l=t&c.F,f=t&c.G,p=t&c.S,h=t&c.P,d=t&c.B,y=t&c.W,v=f?o:o[n]||(o[n]={}),m=f?r:p?r[n]:(r[n]||{}).prototype;f&&(e=n);for(s in e)(u=!l&&m&&s in m)&&s in v||(a=u?m[s]:e[s],v[s]=f&&"function"!=typeof m[s]?e[s]:d&&u?i(a,r):y&&m[s]==a?function(t){var n=function(n){return this instanceof t?new t(n):t(n)};return n.prototype=t.prototype,n}(a):h&&"function"==typeof a?i(Function.call,a):a,h&&((v.prototype||(v.prototype={}))[s]=a))};c.F=1,c.G=2,c.S=4,c.P=8,c.B=16,c.W=32,t.exports=c},function(t,n,e){var r=e(29)("wks"),o=e(30),i=e(5).Symbol;t.exports=function(t){return r[t]||(r[t]=i&&i[t]||(i||o)("Symbol."+t))}},function(t,n){var e=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=e)},function(t,n){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,n,e){t.exports={default:e(43),__esModule:!0}},function(t,n,e){"use strict";n.__esModule=!0,n.default=function(t,n){if(!(t instanceof n))throw new TypeError("Cannot call a class as a function")}},function(t,n,e){"use strict";n.__esModule=!0;var r=e(46),o=function(t){return t&&t.__esModule?t:{default:t}}(r);n.default=function(){function t(t,n){for(var e=0;e<n.length;e++){var r=n[e];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),(0,o.default)(t,r.key,r)}}return function(n,e,r){return e&&t(n.prototype,e),r&&t(n,r),n}}()},function(t,n,e){"use strict";n.__esModule=!0;var r=e(24),o=function(t){return t&&t.__esModule?t:{default:t}}(r);n.default=function(t,n){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!n||"object"!==(void 0===n?"undefined":(0,o.default)(n))&&"function"!=typeof n?t:n}},function(t,n,e){var r=e(22),o=e(14);t.exports=function(t){return r(o(t))}},function(t,n,e){"use strict";function r(t){return t&&t.__esModule?t:{default:t}}n.__esModule=!0;var o=e(66),i=r(o),c=e(70),s=r(c),u=e(24),a=r(u);n.default=function(t,n){if("function"!=typeof n&&null!==n)throw new TypeError("Super expression must either be null or a function, not "+(void 0===n?"undefined":(0,a.default)(n)));t.prototype=(0,s.default)(n&&n.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),n&&(i.default?(0,i.default)(t,n):t.__proto__=n)}},function(t,n,e){t.exports=e(76)()},function(t,n){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,n,e){var r=e(0),o=e(16);t.exports=e(28)?function(t,n,e){return r.setDesc(t,n,o(1,e))}:function(t,n,e){return t[n]=e,t}},function(t,n){t.exports=function(t,n){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:n}}},function(t,n){var e={}.hasOwnProperty;t.exports=function(t,n){return e.call(t,n)}},function(t,n){t.exports={}},function(t,n,e){var r=e(0).setDesc,o=e(17),i=e(4)("toStringTag");t.exports=function(t,n,e){t&&!o(t=e?t:t.prototype,i)&&r(t,i,{configurable:!0,value:n})}},function(t,n,e){var r=e(41);t.exports=function(t,n,e){if(r(t),void 0===n)return t;switch(e){case 1:return function(e){return t.call(n,e)};case 2:return function(e,r){return t.call(n,e,r)};case 3:return function(e,r,o){return t.call(n,e,r,o)}}return function(){return t.apply(n,arguments)}}},function(t,n,e){var r=e(14);t.exports=function(t){return Object(r(t))}},function(t,n,e){var r=e(23);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==r(t)?t.split(""):Object(t)}},function(t,n){var e={}.toString;t.exports=function(t){return e.call(t).slice(8,-1)}},function(t,n,e){"use strict";function r(t){return t&&t.__esModule?t:{default:t}}n.__esModule=!0;var o=e(48),i=r(o),c=e(58),s=r(c),u="function"==typeof s.default&&"symbol"==typeof i.default?function(t){return typeof t}:function(t){return t&&"function"==typeof s.default&&t.constructor===s.default&&t!==s.default.prototype?"symbol":typeof t};n.default="function"==typeof s.default&&"symbol"===u(i.default)?function(t){return void 0===t?"undefined":u(t)}:function(t){return t&&"function"==typeof s.default&&t.constructor===s.default&&t!==s.default.prototype?"symbol":void 0===t?"undefined":u(t)}},function(t,n,e){"use strict";var r=e(26),o=e(3),i=e(27),c=e(15),s=e(17),u=e(18),a=e(53),l=e(19),f=e(0).getProto,p=e(4)("iterator"),h=!([].keys&&"next"in[].keys()),d=function(){return this};t.exports=function(t,n,e,y,v,m,b){a(e,n,y);var _,g,w=function(t){if(!h&&t in k)return k[t];switch(t){case"keys":case"values":return function(){return new e(this,t)}}return function(){return new e(this,t)}},x=n+" Iterator",O="values"==v,S=!1,k=t.prototype,j=k[p]||k["@@iterator"]||v&&k[v],P=j||w(v);if(j){var E=f(P.call(new t));l(E,x,!0),!r&&s(k,"@@iterator")&&c(E,p,d),O&&"values"!==j.name&&(S=!0,P=function(){return j.call(this)})}if(r&&!b||!h&&!S&&k[p]||c(k,p,P),u[n]=P,u[x]=d,v)if(_={values:O?P:w("values"),keys:m?P:w("keys"),entries:O?w("entries"):P},b)for(g in _)g in k||i(k,g,_[g]);else o(o.P+o.F*(h||S),n,_);return _}},function(t,n){t.exports=!0},function(t,n,e){t.exports=e(15)},function(t,n,e){t.exports=!e(6)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,n,e){var r=e(5),o=r["__core-js_shared__"]||(r["__core-js_shared__"]={});t.exports=function(t){return o[t]||(o[t]={})}},function(t,n){var e=0,r=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++e+r).toString(36))}},function(t,n,e){var r=e(32);t.exports=function(t){if(!r(t))throw TypeError(t+" is not an object!");return t}},function(t,n){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,n,e){"use strict";e.d(n,"a",function(){return x});var r=e(7),o=e.n(r),i=e(8),c=e.n(i),s=e(9),u=e.n(s),a=e(10),l=e.n(a),f=e(12),p=e.n(f),h=e(1),d=e.n(h),y=e(75),v=e(34),m=(e.n(v),e(80)),b=e(81),_=(e.n(b),e(13)),g=e.n(_),w=function(t){function n(t){c()(this,n);var e=l()(this,(n.__proto__||o()(n)).call(this,t));return console.log("action comp creation",t),e.renderView=e.renderView.bind(e),e.dispatchAction=e.dispatchAction.bind(e),e.actionFunc=e.actionFunc.bind(e),e.hasPermission=!1,null!=t.action?e.action=t.action:e.action=_reg("Actions",t.name),e.action&&(e.hasPermission=Object(b.hasPermission)(e.action.permission)),e}return p()(n,t),u()(n,[{key:"dispatchAction",value:function(){var t={};this.props.params&&(t=this.props.params),this.props.dispatch(Object(b.createAction)(this.action.action,t,{successCallback:this.props.successCallback,failureCallback:this.props.failureCallback}))}},{key:"actionFunc",value:function(t){if(console.log("action executed",this.props.name,this.props),t.preventDefault(),this.props.confirm&&!this.props.confirm(this.props))return!1;switch(this.action.actiontype){case"dispatchaction":return this.dispatchAction(),!1;case"method":var n=this.props.params;return(0,this.props.method)(n),!1;case"newwindow":if(this.action.url){var e=Object(b.formatUrl)(this.action.url,this.props.params);return console.log(e),window.open(e),!1}default:if(this.action.url){var r=Object(b.formatUrl)(this.action.url,this.props.params);console.log(r),Window.redirect(r)}return!1}}},{key:"renderView",value:function(){if(!this.hasPermission)return null;var t=this.props.children?this.props.children:this.props.label,n=this.actionFunc;switch(this.props.widget){case"button":return d.a.createElement(y.a,{className:this.props.className,actionFunc:n,key:this.props.name+"_comp",actionchildren:t});case"component":return d.a.createElement(this.props.component,{actionFunc:n,key:this.props.name+"_comp",actionchildren:t});default:return d.a.createElement(m.a,{className:this.props.className,actionFunc:n,key:this.props.name+"_comp",actionchildren:t})}}},{key:"render",value:function(){return this.renderView()}}]),n}(d.a.Component);w.propTypes={name:g.a.string.isRequired},console.log("react-redux connect in reactwebcommon",e(34));var x=Object(v.connect)()(w)},function(t,e){t.exports=n},function(t,n,e){"use strict";Object.defineProperty(n,"__esModule",{value:!0});var r=e(36),o=e(72),i=e(74),c=e(33),s=e(82),u=e(83),a=(e.n(u),e(84));e.n(a);e.d(n,"ScrollListener",function(){return i.a}),e.d(n,"Action",function(){return c.a}),e.d(n,"Html",function(){return o.a}),e.d(n,"ActionBar",function(){return s.a}),e.d(n,"Image",function(){return r.a})},function(t,n,e){"use strict";e.d(n,"a",function(){return m});var r=e(37),o=e.n(r),i=e(7),c=e.n(i),s=e(8),u=e.n(s),a=e(9),l=e.n(a),f=e(10),p=e.n(f),h=e(12),d=e.n(h),y=e(1),v=e.n(y),m=function(t){function n(t){u()(this,n);var e=p()(this,(n.__proto__||c()(n)).call(this,t));return t.skipPrefix||(t.skipPrefix=!1),e}return d()(n,t),l()(n,[{key:"render",value:function(){var t=this.props.src;if(!t||0==t.length)return this.props.children?this.props.children:null;!this.props.prefix||this.props.skipPrefix||this.props.src.startsWith("http")||(t=this.props.prefix+t);var n=v.a.createElement("img",o()({src:t},this.props.modifier,{className:this.props.className,style:this.props.style}));return this.props.link?v.a.createElement("a",{target:this.props.target,href:this.props.link},n):n}}]),n}(v.a.Component)},function(t,n,e){"use strict";n.__esModule=!0;var r=e(38),o=function(t){return t&&t.__esModule?t:{default:t}}(r);n.default=o.default||function(t){for(var n=1;n<arguments.length;n++){var e=arguments[n];for(var r in e)Object.prototype.hasOwnProperty.call(e,r)&&(t[r]=e[r])}return t}},function(t,n,e){t.exports={default:e(39),__esModule:!0}},function(t,n,e){e(40),t.exports=e(2).Object.assign},function(t,n,e){var r=e(3);r(r.S+r.F,"Object",{assign:e(42)})},function(t,n){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,n,e){var r=e(0),o=e(21),i=e(22);t.exports=e(6)(function(){var t=Object.assign,n={},e={},r=Symbol(),o="abcdefghijklmnopqrst";return n[r]=7,o.split("").forEach(function(t){e[t]=t}),7!=t({},n)[r]||Object.keys(t({},e)).join("")!=o})?function(t,n){for(var e=o(t),c=arguments,s=c.length,u=1,a=r.getKeys,l=r.getSymbols,f=r.isEnum;s>u;)for(var p,h=i(c[u++]),d=l?a(h).concat(l(h)):a(h),y=d.length,v=0;y>v;)f.call(h,p=d[v++])&&(e[p]=h[p]);return e}:Object.assign},function(t,n,e){e(44),t.exports=e(2).Object.getPrototypeOf},function(t,n,e){var r=e(21);e(45)("getPrototypeOf",function(t){return function(n){return t(r(n))}})},function(t,n,e){var r=e(3),o=e(2),i=e(6);t.exports=function(t,n){var e=(o.Object||{})[t]||Object[t],c={};c[t]=n(e),r(r.S+r.F*i(function(){e(1)}),"Object",c)}},function(t,n,e){t.exports={default:e(47),__esModule:!0}},function(t,n,e){var r=e(0);t.exports=function(t,n,e){return r.setDesc(t,n,e)}},function(t,n,e){t.exports={default:e(49),__esModule:!0}},function(t,n,e){e(50),e(54),t.exports=e(4)("iterator")},function(t,n,e){"use strict";var r=e(51)(!0);e(25)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,n=this._t,e=this._i;return e>=n.length?{value:void 0,done:!0}:(t=r(n,e),this._i+=t.length,{value:t,done:!1})})},function(t,n,e){var r=e(52),o=e(14);t.exports=function(t){return function(n,e){var i,c,s=String(o(n)),u=r(e),a=s.length;return u<0||u>=a?t?"":void 0:(i=s.charCodeAt(u),i<55296||i>56319||u+1===a||(c=s.charCodeAt(u+1))<56320||c>57343?t?s.charAt(u):i:t?s.slice(u,u+2):c-56320+(i-55296<<10)+65536)}}},function(t,n){var e=Math.ceil,r=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?r:e)(t)}},function(t,n,e){"use strict";var r=e(0),o=e(16),i=e(19),c={};e(15)(c,e(4)("iterator"),function(){return this}),t.exports=function(t,n,e){t.prototype=r.create(c,{next:o(1,e)}),i(t,n+" Iterator")}},function(t,n,e){e(55);var r=e(18);r.NodeList=r.HTMLCollection=r.Array},function(t,n,e){"use strict";var r=e(56),o=e(57),i=e(18),c=e(11);t.exports=e(25)(Array,"Array",function(t,n){this._t=c(t),this._i=0,this._k=n},function(){var t=this._t,n=this._k,e=this._i++;return!t||e>=t.length?(this._t=void 0,o(1)):"keys"==n?o(0,e):"values"==n?o(0,t[e]):o(0,[e,t[e]])},"values"),i.Arguments=i.Array,r("keys"),r("values"),r("entries")},function(t,n){t.exports=function(){}},function(t,n){t.exports=function(t,n){return{value:n,done:!!t}}},function(t,n,e){t.exports={default:e(59),__esModule:!0}},function(t,n,e){e(60),e(65),t.exports=e(2).Symbol},function(t,n,e){"use strict";var r=e(0),o=e(5),i=e(17),c=e(28),s=e(3),u=e(27),a=e(6),l=e(29),f=e(19),p=e(30),h=e(4),d=e(61),y=e(62),v=e(63),m=e(64),b=e(31),_=e(11),g=e(16),w=r.getDesc,x=r.setDesc,O=r.create,S=y.get,k=o.Symbol,j=o.JSON,P=j&&j.stringify,E=!1,N=h("_hidden"),M=r.isEnum,T=l("symbol-registry"),F=l("symbols"),A="function"==typeof k,C=Object.prototype,I=c&&a(function(){return 7!=O(x({},"a",{get:function(){return x(this,"a",{value:7}).a}})).a})?function(t,n,e){var r=w(C,n);r&&delete C[n],x(t,n,e),r&&t!==C&&x(C,n,r)}:x,D=function(t){var n=F[t]=O(k.prototype);return n._k=t,c&&E&&I(C,t,{configurable:!0,set:function(n){i(this,N)&&i(this[N],t)&&(this[N][t]=!1),I(this,t,g(1,n))}}),n},R=function(t){return"symbol"==typeof t},L=function(t,n,e){return e&&i(F,n)?(e.enumerable?(i(t,N)&&t[N][n]&&(t[N][n]=!1),e=O(e,{enumerable:g(0,!1)})):(i(t,N)||x(t,N,g(1,{})),t[N][n]=!0),I(t,n,e)):x(t,n,e)},B=function(t,n){b(t);for(var e,r=v(n=_(n)),o=0,i=r.length;i>o;)L(t,e=r[o++],n[e]);return t},W=function(t,n){return void 0===n?O(t):B(O(t),n)},H=function(t){var n=M.call(this,t);return!(n||!i(this,t)||!i(F,t)||i(this,N)&&this[N][t])||n},q=function(t,n){var e=w(t=_(t),n);return!e||!i(F,n)||i(t,N)&&t[N][n]||(e.enumerable=!0),e},U=function(t){for(var n,e=S(_(t)),r=[],o=0;e.length>o;)i(F,n=e[o++])||n==N||r.push(n);return r},V=function(t){for(var n,e=S(_(t)),r=[],o=0;e.length>o;)i(F,n=e[o++])&&r.push(F[n]);return r},J=function(t){if(void 0!==t&&!R(t)){for(var n,e,r=[t],o=1,i=arguments;i.length>o;)r.push(i[o++]);return n=r[1],"function"==typeof n&&(e=n),!e&&m(n)||(n=function(t,n){if(e&&(n=e.call(this,t,n)),!R(n))return n}),r[1]=n,P.apply(j,r)}},K=a(function(){var t=k();return"[null]"!=P([t])||"{}"!=P({a:t})||"{}"!=P(Object(t))});A||(k=function(){if(R(this))throw TypeError("Symbol is not a constructor");return D(p(arguments.length>0?arguments[0]:void 0))},u(k.prototype,"toString",function(){return this._k}),R=function(t){return t instanceof k},r.create=W,r.isEnum=H,r.getDesc=q,r.setDesc=L,r.setDescs=B,r.getNames=y.get=U,r.getSymbols=V,c&&!e(26)&&u(C,"propertyIsEnumerable",H,!0));var z={for:function(t){return i(T,t+="")?T[t]:T[t]=k(t)},keyFor:function(t){return d(T,t)},useSetter:function(){E=!0},useSimple:function(){E=!1}};r.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var n=h(t);z[t]=A?n:D(n)}),E=!0,s(s.G+s.W,{Symbol:k}),s(s.S,"Symbol",z),s(s.S+s.F*!A,"Object",{create:W,defineProperty:L,defineProperties:B,getOwnPropertyDescriptor:q,getOwnPropertyNames:U,getOwnPropertySymbols:V}),j&&s(s.S+s.F*(!A||K),"JSON",{stringify:J}),f(k,"Symbol"),f(Math,"Math",!0),f(o.JSON,"JSON",!0)},function(t,n,e){var r=e(0),o=e(11);t.exports=function(t,n){for(var e,i=o(t),c=r.getKeys(i),s=c.length,u=0;s>u;)if(i[e=c[u++]]===n)return e}},function(t,n,e){var r=e(11),o=e(0).getNames,i={}.toString,c="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[],s=function(t){try{return o(t)}catch(t){return c.slice()}};t.exports.get=function(t){return c&&"[object Window]"==i.call(t)?s(t):o(r(t))}},function(t,n,e){var r=e(0);t.exports=function(t){var n=r.getKeys(t),e=r.getSymbols;if(e)for(var o,i=e(t),c=r.isEnum,s=0;i.length>s;)c.call(t,o=i[s++])&&n.push(o);return n}},function(t,n,e){var r=e(23);t.exports=Array.isArray||function(t){return"Array"==r(t)}},function(t,n){},function(t,n,e){t.exports={default:e(67),__esModule:!0}},function(t,n,e){e(68),t.exports=e(2).Object.setPrototypeOf},function(t,n,e){var r=e(3);r(r.S,"Object",{setPrototypeOf:e(69).set})},function(t,n,e){var r=e(0).getDesc,o=e(32),i=e(31),c=function(t,n){if(i(t),!o(n)&&null!==n)throw TypeError(n+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,n,o){try{o=e(20)(Function.call,r(Object.prototype,"__proto__").set,2),o(t,[]),n=!(t instanceof Array)}catch(t){n=!0}return function(t,e){return c(t,e),n?t.__proto__=e:o(t,e),t}}({},!1):void 0),check:c}},function(t,n,e){t.exports={default:e(71),__esModule:!0}},function(t,n,e){var r=e(0);t.exports=function(t,n){return r.create(t,n)}},function(t,n,e){"use strict";function r(t,n){var e=t;return n&&(e=s()(t)),{__html:e}}e.d(n,"a",function(){return u});var o=e(1),i=e.n(o),c=e(73),s=e.n(c),u=function(t){return i.a.createElement("div",{className:t.className,style:t.style,dangerouslySetInnerHTML:r(t.children,t.sanitize)})}},function(t,n){t.exports=r},function(t,n,e){"use strict";e.d(n,"a",function(){return y});var r=e(7),o=e.n(r),i=e(8),c=e.n(i),s=e(9),u=e.n(s),a=e(10),l=e.n(a),f=e(12),p=e.n(f),h=e(1),d=e.n(h),y=function(t){function n(t){c()(this,n);var e=l()(this,(n.__proto__||o()(n)).call(this,t));if(e.handleScroll=e.handleScroll.bind(e),t.windowScroll){var r=Math.floor(window.scrollY/window.innerHeight);e.state={windowNumber:r}}return e.state={scrolledOut:!1,scrolledIn:!1},e}return p()(n,t),u()(n,[{key:"handleScroll",value:function(t){if(this.props.windowScroll){var n=Math.floor(window.scrollY/window.innerHeight);n!=this.state.windowNumber&&(this.props.windowScroll(n),this.setState({windowNumber:n}))}if(this.props.onScrollEnd||this.props.onScrollIn){var e=this.refs.scrollListener,r=window.innerHeight;if(this.props.scrollEndPos&&(r=this.props.scrollEndPos),null!=e){var o=e.getBoundingClientRect();o.bottom<=r&&!this.state.scrolledOut&&(!this.state.scrolledOut&&this.props.onScrollEnd&&this.props.onScrollEnd(o.bottom),this.setState({scrolledOut:!0,scrolledIn:!1})),this.state.scrolledOut&&o.bottom>r&&(!this.state.scrolledIn&&this.props.onScrollIn&&this.props.onScrollIn(o.bottom),this.setState({scrolledOut:!1,scrolledIn:!0}))}}}},{key:"componentDidMount",value:function(){window.addEventListener("scroll",this.handleScroll)}},{key:"componentWillUnmount",value:function(){window.removeEventListener("scroll",this.handleScroll)}},{key:"render",value:function(){return d.a.createElement("div",{ref:"scrollListener",key:this.props.key,style:this.props.style,className:this.props.className},this.props.children)}}]),n}(d.a.Component)},function(t,n,e){"use strict";var r=e(1),o=e.n(r),i=e(13),c=e.n(i),s=function(t,n){return n.uikit&&n.uikit.ActionButton?o.a.createElement(n.uikit.ActionButton,{className:t.className+" actionbutton",onClick:t.actionFunc,btnProps:t},t.actionchildren):o.a.createElement("a",{className:t.className+" actionbutton",onClick:t.actionFunc,role:"button"},t.actionchildren)};s.contextTypes={uikit:c.a.object},s.propTypes={actionFunc:c.a.func.isRequired,actionchildren:c.a.oneOfType([c.a.array,c.a.string])},n.a=s},function(t,n,e){"use strict";var r=e(77),o=e(78),i=e(79);t.exports=function(){function t(t,n,e,r,c,s){s!==i&&o(!1,"Calling PropTypes validators directly is not supported by the `prop-types` package. Use PropTypes.checkPropTypes() to call them. Read more at http://fb.me/use-check-prop-types")}function n(){return t}t.isRequired=t;var e={array:t,bool:t,func:t,number:t,object:t,string:t,symbol:t,any:t,arrayOf:n,element:t,instanceOf:n,node:t,objectOf:n,oneOf:n,oneOfType:n,shape:n,exact:n};return e.checkPropTypes=r,e.PropTypes=e,e}},function(t,n,e){"use strict";function r(t){return function(){return t}}var o=function(){};o.thatReturns=r,o.thatReturnsFalse=r(!1),o.thatReturnsTrue=r(!0),o.thatReturnsNull=r(null),o.thatReturnsThis=function(){return this},o.thatReturnsArgument=function(t){return t},t.exports=o},function(t,n,e){"use strict";function r(t,n,e,r,i,c,s,u){if(o(n),!t){var a;if(void 0===n)a=new Error("Minified exception occurred; use the non-minified dev environment for the full error message and additional helpful warnings.");else{var l=[e,r,i,c,s,u],f=0;a=new Error(n.replace(/%s/g,function(){return l[f++]})),a.name="Invariant Violation"}throw a.framesToPop=1,a}}var o=function(t){};t.exports=r},function(t,n,e){"use strict";t.exports="SECRET_DO_NOT_PASS_THIS_OR_YOU_WILL_BE_FIRED"},function(t,n,e){"use strict";var r=e(1),o=e.n(r),i=e(13),c=e.n(i),s=function(t){return o.a.createElement("a",{className:t.className+" actionlink",href:"javascript:void(0)",onClick:t.actionFunc},t.actionchildren)};s.propTypes={actionFunc:c.a.func.isRequired,actionchildren:c.a.oneOfType([c.a.array,c.a.string])},n.a=s},function(t,n){t.exports=e},function(t,n,e){"use strict";e.d(n,"a",function(){return b});var r=e(7),o=e.n(r),i=e(8),c=e.n(i),s=e(9),u=e.n(s),a=e(10),l=e.n(a),f=e(12),p=e.n(f),h=e(1),d=e.n(h),y=e(33),v=e(13),m=e.n(v),b=function(t){function n(t){c()(this,n);var e=l()(this,(n.__proto__||o()(n)).call(this,t));e.actions=[];var r=t.description;console.log("action bar",t),e.description=r,e.className=t.className?t.className:r.className;var i=e;return r&&r.actions&&r.actions.forEach(function(t){i.actions.push(d.a.createElement(y.a,{name:t.name,label:t.label,actiontype:r.actiontype,widget:r.widget,className:" action "}))}),e}return p()(n,t),u()(n,[{key:"render",value:function(){return d.a.createElement(this.context.uikit.Block,{className:" actionbar "+this.className},this.actions)}}]),n}(d.a.Component);b.contextTypes={uikit:m.a.object}},function(t,n){},function(t,n){}])});
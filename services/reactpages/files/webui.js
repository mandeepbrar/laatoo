define("reactpages",["react","redux"],function(t,e){return function(t){function e(r){if(n[r])return n[r].exports;var o=n[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,e),o.l=!0,o.exports}var n={};return e.m=t,e.c=n,e.d=function(t,n,r){e.o(t,n)||Object.defineProperty(t,n,{configurable:!1,enumerable:!0,get:r})},e.n=function(t){var n=t&&t.__esModule?function(){return t.default}:function(){return t};return e.d(n,"a",n),n},e.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},e.p="/",e(e.s=34)}([function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,e,n){var r=n(4),o=n(1),i=n(16),u=function(t,e,n){var c,a,s,f=t&u.F,l=t&u.G,p=t&u.S,d=t&u.P,y=t&u.B,m=t&u.W,h=l?o:o[e]||(o[e]={}),v=l?r:p?r[e]:(r[e]||{}).prototype;l&&(n=e);for(c in n)(a=!f&&v&&c in v)&&c in h||(s=a?v[c]:n[c],h[c]=l&&"function"!=typeof v[c]?n[c]:y&&a?i(s,r):m&&v[c]==s?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(s):d&&"function"==typeof s?i(Function.call,s):s,d&&((h.prototype||(h.prototype={}))[c]=s))};u.F=1,u.G=2,u.S=4,u.P=8,u.B=16,u.W=32,t.exports=u},function(t,e,n){var r=n(25)("wks"),o=n(26),i=n(4).Symbol;t.exports=function(t){return r[t]||(r[t]=i&&i[t]||(i||o)("Symbol."+t))}},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e,n){var r=n(27),o=n(8);t.exports=function(t){return r(o(t))}},function(t,e,n){var r=n(8);t.exports=function(t){return Object(r(t))}},function(t,e){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e,n){var r=n(0),o=n(10);t.exports=n(24)?function(t,e,n){return r.setDesc(t,e,o(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports={}},function(t,e,n){var r=n(0).setDesc,o=n(11),i=n(3)("toStringTag");t.exports=function(t,e,n){t&&!o(t=n?t:t.prototype,i)&&r(t,i,{configurable:!0,value:e})}},function(t,e,n){t.exports={default:n(35),__esModule:!0}},function(t,e,n){var r=n(2),o=n(1),i=n(5);t.exports=function(t,e){var n=(o.Object||{})[t]||Object[t],u={};u[t]=e(n),r(r.S+r.F*i(function(){n(1)}),"Object",u)}},function(t,e,n){var r=n(37);t.exports=function(t,e,n){if(r(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,r){return t.call(e,n,r)};case 3:return function(n,r,o){return t.call(e,n,r,o)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){"use strict";e.__esModule=!0;var r=n(38),o=function(t){return t&&t.__esModule?t:{default:t}}(r);e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var r=e[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),(0,o.default)(t,r.key,r)}}return function(e,n,r){return n&&t(e.prototype,n),r&&t(e,r),e}}()},function(t,e,n){"use strict";e.__esModule=!0;var r=n(20),o=function(t){return t&&t.__esModule?t:{default:t}}(r);e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,o.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){"use strict";function r(t){return t&&t.__esModule?t:{default:t}}e.__esModule=!0;var o=n(40),i=r(o),u=n(50),c=r(u),a="function"==typeof c.default&&"symbol"==typeof i.default?function(t){return typeof t}:function(t){return t&&"function"==typeof c.default&&t.constructor===c.default&&t!==c.default.prototype?"symbol":typeof t};e.default="function"==typeof c.default&&"symbol"===a(i.default)?function(t){return void 0===t?"undefined":a(t)}:function(t){return t&&"function"==typeof c.default&&t.constructor===c.default&&t!==c.default.prototype?"symbol":void 0===t?"undefined":a(t)}},function(t,e,n){"use strict";var r=n(22),o=n(2),i=n(23),u=n(9),c=n(11),a=n(12),s=n(45),f=n(13),l=n(0).getProto,p=n(3)("iterator"),d=!([].keys&&"next"in[].keys()),y=function(){return this};t.exports=function(t,e,n,m,h,v,g){s(n,e,m);var _,b,k=function(t){if(!d&&t in P)return P[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},x=e+" Iterator",w="values"==h,O=!1,P=t.prototype,E=P[p]||P["@@iterator"]||h&&P[h],S=E||k(h);if(E){var j=l(S.call(new t));f(j,x,!0),!r&&c(P,"@@iterator")&&u(j,p,y),w&&"values"!==E.name&&(O=!0,S=function(){return E.call(this)})}if(r&&!g||!d&&!O&&P[p]||u(P,p,S),a[e]=S,a[x]=y,h)if(_={values:w?S:k("values"),keys:v?S:k("keys"),entries:w?k("entries"):S},g)for(b in _)b in P||i(P,b,_[b]);else o(o.P+o.F*(d||O),e,_);return _}},function(t,e){t.exports=!0},function(t,e,n){t.exports=n(9)},function(t,e,n){t.exports=!n(5)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){var r=n(4),o=r["__core-js_shared__"]||(r["__core-js_shared__"]={});t.exports=function(t){return o[t]||(o[t]={})}},function(t,e){var n=0,r=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+r).toString(36))}},function(t,e,n){var r=n(28);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==r(t)?t.split(""):Object(t)}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){var r=n(30);t.exports=function(t){if(!r(t))throw TypeError(t+" is not an object!");return t}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){"use strict";function r(t){return t&&t.__esModule?t:{default:t}}e.__esModule=!0;var o=n(58),i=r(o),u=n(62),c=r(u),a=n(20),s=r(a);e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,s.default)(e)));t.prototype=(0,c.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(i.default?(0,i.default)(t,e):t.__proto__=e)}},function(e,n){e.exports=t},function(t,e,n){t.exports=n(73)()},function(t,e,n){"use strict";function r(e,n,r,o,i,u){t.properties=Application.Properties[n],t.settings=o,t.req=u,k.a.setModule(t)}function o(t,e){var n=Application.AllRegItems("Pages");if(n)for(var r in n)try{!function(){var o=n[r],u=i(o),c=o.components;o.component&&(c={main:o.component});var a={};v()(c).forEach(function(t){a[t]=function(e,n){return function(r){return _.a.createElement(O,{pageId:n,placeholder:t,routerState:r,description:e})}}(c[t],r)});var s={pattern:o.route,components:a,reducer:Object(x.combineReducers)(u)},f=s;t&&t.ProcessRoute&&(f=t.ProcessRoute(s,e)),Application.Register("Routes",r,f),Application.Register("Actions","Page_"+r,{url:f.pattern})}()}catch(t){console.log(t)}}function i(e){var n={};for(var r in e.datasources)try{var o=_reg("Datasources",r),i={};o.type;var u=o.module;if(u){var c=t.req(u);c&&(i=c[o.processor])}i&&(n[r]=i)}catch(t){}return n}Object.defineProperty(e,"__esModule",{value:!0}),n.d(e,"Initialize",function(){return r}),n.d(e,"ProcessPages",function(){return o});var u=n(14),c=n.n(u),a=n(17),s=n.n(a),f=n(18),l=n.n(f),p=n(19),d=n.n(p),y=n(31),m=n.n(y),h=n(64),v=n.n(h),g=n(32),_=n.n(g),b=n(67),k=(n.n(b),n(68)),x=n(77);n.n(x);n.d(e,"Panel",function(){return k.a});var w=n(33),t=this,O=function(t){function e(){return s()(this,e),d()(this,(e.__proto__||c()(e)).apply(this,arguments))}return m()(e,t),l()(e,[{key:"getChildContext",value:function(){return{routeParams:this.props.routerState.params}}},{key:"render",value:function(){var t=this.props.pageId+this.props.placeholder;return _.a.createElement(k.a,{key:t,description:this.props.description})}}]),e}(_.a.Component);O.childContextTypes={routeParams:w.object}},function(t,e,n){n(36),t.exports=n(1).Object.getPrototypeOf},function(t,e,n){var r=n(7);n(15)("getPrototypeOf",function(t){return function(e){return t(r(e))}})},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){t.exports={default:n(39),__esModule:!0}},function(t,e,n){var r=n(0);t.exports=function(t,e,n){return r.setDesc(t,e,n)}},function(t,e,n){t.exports={default:n(41),__esModule:!0}},function(t,e,n){n(42),n(46),t.exports=n(3)("iterator")},function(t,e,n){"use strict";var r=n(43)(!0);n(21)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=r(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){var r=n(44),o=n(8);t.exports=function(t){return function(e,n){var i,u,c=String(o(e)),a=r(n),s=c.length;return a<0||a>=s?t?"":void 0:(i=c.charCodeAt(a),i<55296||i>56319||a+1===s||(u=c.charCodeAt(a+1))<56320||u>57343?t?c.charAt(a):i:t?c.slice(a,a+2):u-56320+(i-55296<<10)+65536)}}},function(t,e){var n=Math.ceil,r=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?r:n)(t)}},function(t,e,n){"use strict";var r=n(0),o=n(10),i=n(13),u={};n(9)(u,n(3)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=r.create(u,{next:o(1,n)}),i(t,e+" Iterator")}},function(t,e,n){n(47);var r=n(12);r.NodeList=r.HTMLCollection=r.Array},function(t,e,n){"use strict";var r=n(48),o=n(49),i=n(12),u=n(6);t.exports=n(21)(Array,"Array",function(t,e){this._t=u(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,o(1)):"keys"==e?o(0,n):"values"==e?o(0,t[n]):o(0,[n,t[n]])},"values"),i.Arguments=i.Array,r("keys"),r("values"),r("entries")},function(t,e){t.exports=function(){}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e,n){t.exports={default:n(51),__esModule:!0}},function(t,e,n){n(52),n(57),t.exports=n(1).Symbol},function(t,e,n){"use strict";var r=n(0),o=n(4),i=n(11),u=n(24),c=n(2),a=n(23),s=n(5),f=n(25),l=n(13),p=n(26),d=n(3),y=n(53),m=n(54),h=n(55),v=n(56),g=n(29),_=n(6),b=n(10),k=r.getDesc,x=r.setDesc,w=r.create,O=m.get,P=o.Symbol,E=o.JSON,S=E&&E.stringify,j=!1,M=d("_hidden"),N=r.isEnum,T=f("symbol-registry"),D=f("symbols"),A="function"==typeof P,B=Object.prototype,I=u&&s(function(){return 7!=w(x({},"a",{get:function(){return x(this,"a",{value:7}).a}})).a})?function(t,e,n){var r=k(B,e);r&&delete B[e],x(t,e,n),r&&t!==B&&x(B,e,r)}:x,C=function(t){var e=D[t]=w(P.prototype);return e._k=t,u&&j&&I(B,t,{configurable:!0,set:function(e){i(this,M)&&i(this[M],t)&&(this[M][t]=!1),I(this,t,b(1,e))}}),e},F=function(t){return"symbol"==typeof t},R=function(t,e,n){return n&&i(D,e)?(n.enumerable?(i(t,M)&&t[M][e]&&(t[M][e]=!1),n=w(n,{enumerable:b(0,!1)})):(i(t,M)||x(t,M,b(1,{})),t[M][e]=!0),I(t,e,n)):x(t,e,n)},V=function(t,e){g(t);for(var n,r=h(e=_(e)),o=0,i=r.length;i>o;)R(t,n=r[o++],e[n]);return t},q=function(t,e){return void 0===e?w(t):V(w(t),e)},L=function(t){var e=N.call(this,t);return!(e||!i(this,t)||!i(D,t)||i(this,M)&&this[M][t])||e},H=function(t,e){var n=k(t=_(t),e);return!n||!i(D,e)||i(t,M)&&t[M][e]||(n.enumerable=!0),n},W=function(t){for(var e,n=O(_(t)),r=[],o=0;n.length>o;)i(D,e=n[o++])||e==M||r.push(e);return r},J=function(t){for(var e,n=O(_(t)),r=[],o=0;n.length>o;)i(D,e=n[o++])&&r.push(D[e]);return r},K=function(t){if(void 0!==t&&!F(t)){for(var e,n,r=[t],o=1,i=arguments;i.length>o;)r.push(i[o++]);return e=r[1],"function"==typeof e&&(n=e),!n&&v(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!F(e))return e}),r[1]=e,S.apply(E,r)}},G=s(function(){var t=P();return"[null]"!=S([t])||"{}"!=S({a:t})||"{}"!=S(Object(t))});A||(P=function(){if(F(this))throw TypeError("Symbol is not a constructor");return C(p(arguments.length>0?arguments[0]:void 0))},a(P.prototype,"toString",function(){return this._k}),F=function(t){return t instanceof P},r.create=q,r.isEnum=L,r.getDesc=H,r.setDesc=R,r.setDescs=V,r.getNames=m.get=W,r.getSymbols=J,u&&!n(22)&&a(B,"propertyIsEnumerable",L,!0));var z={for:function(t){return i(T,t+="")?T[t]:T[t]=P(t)},keyFor:function(t){return y(T,t)},useSetter:function(){j=!0},useSimple:function(){j=!1}};r.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=d(t);z[t]=A?e:C(e)}),j=!0,c(c.G+c.W,{Symbol:P}),c(c.S,"Symbol",z),c(c.S+c.F*!A,"Object",{create:q,defineProperty:R,defineProperties:V,getOwnPropertyDescriptor:H,getOwnPropertyNames:W,getOwnPropertySymbols:J}),E&&c(c.S+c.F*(!A||G),"JSON",{stringify:K}),l(P,"Symbol"),l(Math,"Math",!0),l(o.JSON,"JSON",!0)},function(t,e,n){var r=n(0),o=n(6);t.exports=function(t,e){for(var n,i=o(t),u=r.getKeys(i),c=u.length,a=0;c>a;)if(i[n=u[a++]]===e)return n}},function(t,e,n){var r=n(6),o=n(0).getNames,i={}.toString,u="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[],c=function(t){try{return o(t)}catch(t){return u.slice()}};t.exports.get=function(t){return u&&"[object Window]"==i.call(t)?c(t):o(r(t))}},function(t,e,n){var r=n(0);t.exports=function(t){var e=r.getKeys(t),n=r.getSymbols;if(n)for(var o,i=n(t),u=r.isEnum,c=0;i.length>c;)u.call(t,o=i[c++])&&e.push(o);return e}},function(t,e,n){var r=n(28);t.exports=Array.isArray||function(t){return"Array"==r(t)}},function(t,e){},function(t,e,n){t.exports={default:n(59),__esModule:!0}},function(t,e,n){n(60),t.exports=n(1).Object.setPrototypeOf},function(t,e,n){var r=n(2);r(r.S,"Object",{setPrototypeOf:n(61).set})},function(t,e,n){var r=n(0).getDesc,o=n(30),i=n(29),u=function(t,e){if(i(t),!o(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,o){try{o=n(16)(Function.call,r(Object.prototype,"__proto__").set,2),o(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return u(t,n),e?t.__proto__=n:o(t,n),t}}({},!1):void 0),check:u}},function(t,e,n){t.exports={default:n(63),__esModule:!0}},function(t,e,n){var r=n(0);t.exports=function(t,e){return r.create(t,e)}},function(t,e,n){t.exports={default:n(65),__esModule:!0}},function(t,e,n){n(66),t.exports=n(1).Object.keys},function(t,e,n){var r=n(7);n(15)("keys",function(t){return function(e){return t(r(e))}})},function(t,e){},function(t,e,n){"use strict";var t,r=n(69),o=n.n(r),i=n(14),u=n.n(i),c=n(17),a=n.n(c),s=n(18),f=n.n(s),l=n(19),p=n.n(l),d=n(31),y=n.n(d),m=n(32),h=n.n(m),v=n(33),g=n.n(v),_=function(e){function n(e,r){a()(this,n);var i=p()(this,(n.__proto__||u()(n)).call(this,e));b.call(i),i.uikit=r.uikit;var c=e.description;!c&&e.id&&(c=_reg("Panels",e.id)),console.log("creating panel",c,e,r,i.context);var s="panel ";if(c.id&&(s=s+" "+c.id),c.name&&(s=s+" "+c.name),c&&"string"==typeof c)i.processBlock(c,e,s);else if(c)switch(c.type){case"view":s+=" view ",i.processView(c,e,r,s);break;case"entity":s+=" entity ",i.processEntity(c,e,r,s);break;case"form":s+=" form ",i.processForm(c,e,r,s);break;case"html":s+=" html ",i.processHtml(c,e,r,s);break;case"block":s+=" block ",i.processBlock(c,e,r,s);break;case"layout":s+=" panel ",i.processLayout(c,e,r,s);break;case"component":if(c.component)i.getView=function(t,e,n){return c.component};else{var f=i.getComponent(c.module,c.componentName,t.req);console.log("rendering component",c,f);var l={className:s},d=c.props?o()({},c.props,l):l;i.getView=function(t,e){return function(n,r,o){return console.log("rendering comp",e,t),h.a.createElement(e,t)}}(d,f)}break;default:var y=_reg("PanelTypes",c.type);i.getView=function(t,e,n){return y.getComponent(c,t,e,n)}}return i}return y()(n,e),f()(n,[{key:"getDisplayFunc",value:function(t,e){if(!t)return null;if("string"==typeof t)return _reg("Blocks",t);var n=_reg("Blocks",t.block);return n||(n=_reg("Blocks",t.defaultBlock)),n}},{key:"render",value:function(){return this.getView?this.getView(this.props,this.context,this.state):h.a.createElement(this.context.uikit.Block,null)}}],[{key:"setModule",value:function(e){t=e}}]),n}(h.a.Component),b=function(){var e=this;this.getPanelItems=function(t){if(!t)return null;if(t instanceof Array){return t.map(function(t){return h.a.createElement(_,{description:t})})}return h.a.createElement(_,{description:t})},this.processLayout=function(t,n,r,o){var i=t;if(t.id&&(i=_reg("Panels",t.id)),i.layout){var u=e,c=null,a=function(t){return i[t]?h.a.createElement(u.uikit.Block,{className:t},u.getPanelItems(i[t])):null};switch(i.layout){case"2col":c=h.a.createElement(e.uikit.Block,{className:o+" twocol"},a("header"),h.a.createElement(e.uikit.Block,{className:"row"},a("left"),a("right")),a("footer"));break;case"3col":c=h.a.createElement(e.uikit.Block,{className:o+" threecol"},a("header"),h.a.createElement(e.uikit.Block,{className:"row"},a("left"),a("right")),a("footer"));break;default:c=h.a.createElement(e.uikit.Block,{className:o},a("items"))}e.getView=function(t,e,n){return c}}},this.processBlock=function(t,n,r,o){var i=e.getDisplayFunc(t,n);console.log("processing block",t,i,n),e.getView=i?function(e,n,r){return i(e.data,t,n.uikit)}:function(t,e,n){return h.a.createElement(e.uikit.Block,null)}},this.createMarkup=function(t){return{__html:t}},this.processHtml=function(t,n,r,o){t.html?e.getView=function(e,n,r){return console.log("rendering html",t.html),h.a.createElement("div",{className:o,dangerouslySetInnerHTML:this.createMarkup(t.html)})}:e.getView=function(t,e,n){return h.a.createElement(e.uikit.Block,null)}},this.processForm=function(n,r,i,u){console.log("processing form",n);var c=n;if(n.id&&(c=_reg("Forms",n.id)),c){console.log("processing form",n,c);var a=c.info;e.form||(e.form=e.getComponent("reactforms","Form",t.req)),e.form?e.getView=function(t,e,r){var i=o()({},a,e.routeParams);return console.log("form cfg",i,a),h.a.createElement(this.form,{form:n.id,config:i,description:c,id:n.id})}:e.getView=function(t,e,n){return h.a.createElement(e.uikit.Block,null)}}},this.processView=function(n,r,i,u){var c=n.viewid,a=n;c&&(a=_reg("Views",c));var s=o()({},a,n),f=s.header?h.a.createElement(_,{description:s.header}):null;e.view||(e.view=e.getComponent("laatooviews","View",t.req)),e.getView=function(t,e,n){return h.a.createElement(this.view,{params:t.params,header:f,id:c},h.a.createElement(_,{description:s.item}))}},this.processEntity=function(n,r,o,i){var u=n.entityDisplay?n.entityDisplay:"default";console.log("view entity description",n,u);var c="",a="";c=o.routeParams&&o.routeParams.entityId?o.routeParams.entityId:n.entityId,a=n.entityName,e.entity||(e.entity=e.getComponent("laatooviews","Entity",t.req)),console.log("processing entity",r," entity id ",c,"data",r.data);var s={type:"block",block:n.entityName+"_"+u,defaultBlock:n.entityName+"_default"};console.log("entity display",s),e.getView=function(t,e,r){return console.log("entity display in get view",s),h.a.createElement(this.entity,{id:c,name:a,entityDescription:n,data:t.data,index:t.index,uikit:e.uikit},h.a.createElement(_,{description:s}))}},this.getComponent=function(e,n,r){var o=e+n,i=t[o];if(!i){var u=r(e);u&&n&&(i=u[n],t[o]=i)}return i}};_.contextTypes={uikit:g.a.object,routeParams:g.a.object},e.a=_},function(t,e,n){t.exports={default:n(70),__esModule:!0}},function(t,e,n){n(71),t.exports=n(1).Object.assign},function(t,e,n){var r=n(2);r(r.S+r.F,"Object",{assign:n(72)})},function(t,e,n){var r=n(0),o=n(7),i=n(27);t.exports=n(5)(function(){var t=Object.assign,e={},n={},r=Symbol(),o="abcdefghijklmnopqrst";return e[r]=7,o.split("").forEach(function(t){n[t]=t}),7!=t({},e)[r]||Object.keys(t({},n)).join("")!=o})?function(t,e){for(var n=o(t),u=arguments,c=u.length,a=1,s=r.getKeys,f=r.getSymbols,l=r.isEnum;c>a;)for(var p,d=i(u[a++]),y=f?s(d).concat(f(d)):s(d),m=y.length,h=0;m>h;)l.call(d,p=y[h++])&&(n[p]=d[p]);return n}:Object.assign},function(t,e,n){"use strict";var r=n(74),o=n(75),i=n(76);t.exports=function(){function t(t,e,n,r,u,c){c!==i&&o(!1,"Calling PropTypes validators directly is not supported by the `prop-types` package. Use PropTypes.checkPropTypes() to call them. Read more at http://fb.me/use-check-prop-types")}function e(){return t}t.isRequired=t;var n={array:t,bool:t,func:t,number:t,object:t,string:t,symbol:t,any:t,arrayOf:e,element:t,instanceOf:e,node:t,objectOf:e,oneOf:e,oneOfType:e,shape:e,exact:e};return n.checkPropTypes=r,n.PropTypes=n,n}},function(t,e,n){"use strict";function r(t){return function(){return t}}var o=function(){};o.thatReturns=r,o.thatReturnsFalse=r(!1),o.thatReturnsTrue=r(!0),o.thatReturnsNull=r(null),o.thatReturnsThis=function(){return this},o.thatReturnsArgument=function(t){return t},t.exports=o},function(t,e,n){"use strict";function r(t,e,n,r,i,u,c,a){if(o(e),!t){var s;if(void 0===e)s=new Error("Minified exception occurred; use the non-minified dev environment for the full error message and additional helpful warnings.");else{var f=[n,r,i,u,c,a],l=0;s=new Error(e.replace(/%s/g,function(){return f[l++]})),s.name="Invariant Violation"}throw s.framesToPop=1,s}}var o=function(t){};t.exports=r},function(t,e,n){"use strict";t.exports="SECRET_DO_NOT_PASS_THIS_OR_YOU_WILL_BE_FIRED"},function(t,n){t.exports=e}])});
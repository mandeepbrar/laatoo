define("reactpages",["react","redux"],function(t,e){return function(t){function e(r){if(n[r])return n[r].exports;var o=n[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,e),o.l=!0,o.exports}var n={};return e.m=t,e.c=n,e.d=function(t,n,r){e.o(t,n)||Object.defineProperty(t,n,{configurable:!1,enumerable:!0,get:r})},e.n=function(t){var n=t&&t.__esModule?function(){return t.default}:function(){return t};return e.d(n,"a",n),n},e.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},e.p="/",e(e.s=34)}([function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,e,n){var r=n(4),o=n(1),i=n(16),u=function(t,e,n){var c,a,s,f=t&u.F,l=t&u.G,p=t&u.S,d=t&u.P,y=t&u.B,m=t&u.W,v=l?o:o[e]||(o[e]={}),g=l?r:p?r[e]:(r[e]||{}).prototype;l&&(n=e);for(c in n)(a=!f&&g&&c in g)&&c in v||(s=a?g[c]:n[c],v[c]=l&&"function"!=typeof g[c]?n[c]:y&&a?i(s,r):m&&g[c]==s?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(s):d&&"function"==typeof s?i(Function.call,s):s,d&&((v.prototype||(v.prototype={}))[c]=s))};u.F=1,u.G=2,u.S=4,u.P=8,u.B=16,u.W=32,t.exports=u},function(t,e,n){var r=n(25)("wks"),o=n(26),i=n(4).Symbol;t.exports=function(t){return r[t]||(r[t]=i&&i[t]||(i||o)("Symbol."+t))}},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e,n){var r=n(27),o=n(8);t.exports=function(t){return r(o(t))}},function(t,e,n){var r=n(8);t.exports=function(t){return Object(r(t))}},function(t,e){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e,n){var r=n(0),o=n(10);t.exports=n(24)?function(t,e,n){return r.setDesc(t,e,o(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports={}},function(t,e,n){var r=n(0).setDesc,o=n(11),i=n(3)("toStringTag");t.exports=function(t,e,n){t&&!o(t=n?t:t.prototype,i)&&r(t,i,{configurable:!0,value:e})}},function(t,e,n){t.exports={default:n(35),__esModule:!0}},function(t,e,n){var r=n(2),o=n(1),i=n(5);t.exports=function(t,e){var n=(o.Object||{})[t]||Object[t],u={};u[t]=e(n),r(r.S+r.F*i(function(){n(1)}),"Object",u)}},function(t,e,n){var r=n(37);t.exports=function(t,e,n){if(r(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,r){return t.call(e,n,r)};case 3:return function(n,r,o){return t.call(e,n,r,o)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){"use strict";e.__esModule=!0;var r=n(38),o=function(t){return t&&t.__esModule?t:{default:t}}(r);e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var r=e[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),(0,o.default)(t,r.key,r)}}return function(e,n,r){return n&&t(e.prototype,n),r&&t(e,r),e}}()},function(t,e,n){"use strict";e.__esModule=!0;var r=n(20),o=function(t){return t&&t.__esModule?t:{default:t}}(r);e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,o.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){"use strict";function r(t){return t&&t.__esModule?t:{default:t}}e.__esModule=!0;var o=n(40),i=r(o),u=n(50),c=r(u),a="function"==typeof c.default&&"symbol"==typeof i.default?function(t){return typeof t}:function(t){return t&&"function"==typeof c.default&&t.constructor===c.default&&t!==c.default.prototype?"symbol":typeof t};e.default="function"==typeof c.default&&"symbol"===a(i.default)?function(t){return void 0===t?"undefined":a(t)}:function(t){return t&&"function"==typeof c.default&&t.constructor===c.default&&t!==c.default.prototype?"symbol":void 0===t?"undefined":a(t)}},function(t,e,n){"use strict";var r=n(22),o=n(2),i=n(23),u=n(9),c=n(11),a=n(12),s=n(45),f=n(13),l=n(0).getProto,p=n(3)("iterator"),d=!([].keys&&"next"in[].keys()),y=function(){return this};t.exports=function(t,e,n,m,v,g,h){s(n,e,m);var _,b,x=function(t){if(!d&&t in O)return O[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},k=e+" Iterator",w="values"==v,P=!1,O=t.prototype,S=O[p]||O["@@iterator"]||v&&O[v],E=S||x(v);if(S){var j=l(E.call(new t));f(j,k,!0),!r&&c(O,"@@iterator")&&u(j,p,y),w&&"values"!==S.name&&(P=!0,E=function(){return S.call(this)})}if(r&&!h||!d&&!P&&O[p]||u(O,p,E),a[e]=E,a[k]=y,v)if(_={values:w?E:x("values"),keys:g?E:x("keys"),entries:w?x("entries"):E},h)for(b in _)b in O||i(O,b,_[b]);else o(o.P+o.F*(d||P),e,_);return _}},function(t,e){t.exports=!0},function(t,e,n){t.exports=n(9)},function(t,e,n){t.exports=!n(5)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){var r=n(4),o=r["__core-js_shared__"]||(r["__core-js_shared__"]={});t.exports=function(t){return o[t]||(o[t]={})}},function(t,e){var n=0,r=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+r).toString(36))}},function(t,e,n){var r=n(28);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==r(t)?t.split(""):Object(t)}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){var r=n(30);t.exports=function(t){if(!r(t))throw TypeError(t+" is not an object!");return t}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){"use strict";function r(t){return t&&t.__esModule?t:{default:t}}e.__esModule=!0;var o=n(58),i=r(o),u=n(62),c=r(u),a=n(20),s=r(a);e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,s.default)(e)));t.prototype=(0,c.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(i.default?(0,i.default)(t,e):t.__proto__=e)}},function(e,n){e.exports=t},function(t,e,n){t.exports=n(73)()},function(t,e,n){"use strict";function r(e,n,r,o,i,u){t.properties=Application.Properties[n],t.settings=o,t.req=u,Window.redirectPage||(Window.redirectPage=function(t,e){var n=_reg("Pages",t);if(console.log("redirect page",n),n){var r=formatUrl(n.url,e);Window.redirect(r)}}),x.a.setModule(t)}function o(t,e){var n=Application.AllRegItems("Pages");if(n)for(var r in n)try{!function(){var o=n[r],u=i(o),c=o.components;o.component&&(c={main:o.component});var a={};g()(c).forEach(function(t){a[t]=function(e,n){return function(r){return _.a.createElement(P,{pageId:n,placeholder:t,routerState:r,description:e})}}(c[t],r)});var s={pattern:o.route,components:a,reducer:Object(k.combineReducers)(u)},f=s;t&&t.ProcessRoute&&(f=t.ProcessRoute(s,e)),Application.Register("Routes",r,f),Application.Register("Actions","Page_"+r,{url:f.pattern})}()}catch(t){console.log(t)}}function i(e){var n={};for(var r in e.datasources)try{var o=_reg("Datasources",r),i={};o.type;var u=o.module;if(u){var c=t.req(u);c&&(i=c[o.processor])}i&&(n[r]=i)}catch(t){}return n}Object.defineProperty(e,"__esModule",{value:!0}),n.d(e,"Initialize",function(){return r}),n.d(e,"ProcessPages",function(){return o});var u=n(14),c=n.n(u),a=n(17),s=n.n(a),f=n(18),l=n.n(f),p=n(19),d=n.n(p),y=n(31),m=n.n(y),v=n(64),g=n.n(v),h=n(32),_=n.n(h),b=n(67),x=(n.n(b),n(68)),k=n(77);n.n(k);n.d(e,"Panel",function(){return x.a});var w=n(33),t=this,P=function(t){function e(){return s()(this,e),d()(this,(e.__proto__||c()(e)).apply(this,arguments))}return m()(e,t),l()(e,[{key:"getChildContext",value:function(){return{routeParams:this.props.routerState.params}}},{key:"render",value:function(){var t=this.props.pageId+this.props.placeholder;return _.a.createElement(x.a,{key:t,description:this.props.description})}}]),e}(_.a.Component);P.childContextTypes={routeParams:w.object}},function(t,e,n){n(36),t.exports=n(1).Object.getPrototypeOf},function(t,e,n){var r=n(7);n(15)("getPrototypeOf",function(t){return function(e){return t(r(e))}})},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){t.exports={default:n(39),__esModule:!0}},function(t,e,n){var r=n(0);t.exports=function(t,e,n){return r.setDesc(t,e,n)}},function(t,e,n){t.exports={default:n(41),__esModule:!0}},function(t,e,n){n(42),n(46),t.exports=n(3)("iterator")},function(t,e,n){"use strict";var r=n(43)(!0);n(21)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=r(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){var r=n(44),o=n(8);t.exports=function(t){return function(e,n){var i,u,c=String(o(e)),a=r(n),s=c.length;return a<0||a>=s?t?"":void 0:(i=c.charCodeAt(a),i<55296||i>56319||a+1===s||(u=c.charCodeAt(a+1))<56320||u>57343?t?c.charAt(a):i:t?c.slice(a,a+2):u-56320+(i-55296<<10)+65536)}}},function(t,e){var n=Math.ceil,r=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?r:n)(t)}},function(t,e,n){"use strict";var r=n(0),o=n(10),i=n(13),u={};n(9)(u,n(3)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=r.create(u,{next:o(1,n)}),i(t,e+" Iterator")}},function(t,e,n){n(47);var r=n(12);r.NodeList=r.HTMLCollection=r.Array},function(t,e,n){"use strict";var r=n(48),o=n(49),i=n(12),u=n(6);t.exports=n(21)(Array,"Array",function(t,e){this._t=u(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,o(1)):"keys"==e?o(0,n):"values"==e?o(0,t[n]):o(0,[n,t[n]])},"values"),i.Arguments=i.Array,r("keys"),r("values"),r("entries")},function(t,e){t.exports=function(){}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e,n){t.exports={default:n(51),__esModule:!0}},function(t,e,n){n(52),n(57),t.exports=n(1).Symbol},function(t,e,n){"use strict";var r=n(0),o=n(4),i=n(11),u=n(24),c=n(2),a=n(23),s=n(5),f=n(25),l=n(13),p=n(26),d=n(3),y=n(53),m=n(54),v=n(55),g=n(56),h=n(29),_=n(6),b=n(10),x=r.getDesc,k=r.setDesc,w=r.create,P=m.get,O=o.Symbol,S=o.JSON,E=S&&S.stringify,j=!1,M=d("_hidden"),N=r.isEnum,T=f("symbol-registry"),D=f("symbols"),A="function"==typeof O,B=Object.prototype,C=u&&s(function(){return 7!=w(k({},"a",{get:function(){return k(this,"a",{value:7}).a}})).a})?function(t,e,n){var r=x(B,e);r&&delete B[e],k(t,e,n),r&&t!==B&&k(B,e,r)}:k,I=function(t){var e=D[t]=w(O.prototype);return e._k=t,u&&j&&C(B,t,{configurable:!0,set:function(e){i(this,M)&&i(this[M],t)&&(this[M][t]=!1),C(this,t,b(1,e))}}),e},F=function(t){return"symbol"==typeof t},R=function(t,e,n){return n&&i(D,e)?(n.enumerable?(i(t,M)&&t[M][e]&&(t[M][e]=!1),n=w(n,{enumerable:b(0,!1)})):(i(t,M)||k(t,M,b(1,{})),t[M][e]=!0),C(t,e,n)):k(t,e,n)},V=function(t,e){h(t);for(var n,r=v(e=_(e)),o=0,i=r.length;i>o;)R(t,n=r[o++],e[n]);return t},q=function(t,e){return void 0===e?w(t):V(w(t),e)},W=function(t){var e=N.call(this,t);return!(e||!i(this,t)||!i(D,t)||i(this,M)&&this[M][t])||e},L=function(t,e){var n=x(t=_(t),e);return!n||!i(D,e)||i(t,M)&&t[M][e]||(n.enumerable=!0),n},H=function(t){for(var e,n=P(_(t)),r=[],o=0;n.length>o;)i(D,e=n[o++])||e==M||r.push(e);return r},J=function(t){for(var e,n=P(_(t)),r=[],o=0;n.length>o;)i(D,e=n[o++])&&r.push(D[e]);return r},K=function(t){if(void 0!==t&&!F(t)){for(var e,n,r=[t],o=1,i=arguments;i.length>o;)r.push(i[o++]);return e=r[1],"function"==typeof e&&(n=e),!n&&g(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!F(e))return e}),r[1]=e,E.apply(S,r)}},G=s(function(){var t=O();return"[null]"!=E([t])||"{}"!=E({a:t})||"{}"!=E(Object(t))});A||(O=function(){if(F(this))throw TypeError("Symbol is not a constructor");return I(p(arguments.length>0?arguments[0]:void 0))},a(O.prototype,"toString",function(){return this._k}),F=function(t){return t instanceof O},r.create=q,r.isEnum=W,r.getDesc=L,r.setDesc=R,r.setDescs=V,r.getNames=m.get=H,r.getSymbols=J,u&&!n(22)&&a(B,"propertyIsEnumerable",W,!0));var U={for:function(t){return i(T,t+="")?T[t]:T[t]=O(t)},keyFor:function(t){return y(T,t)},useSetter:function(){j=!0},useSimple:function(){j=!1}};r.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=d(t);U[t]=A?e:I(e)}),j=!0,c(c.G+c.W,{Symbol:O}),c(c.S,"Symbol",U),c(c.S+c.F*!A,"Object",{create:q,defineProperty:R,defineProperties:V,getOwnPropertyDescriptor:L,getOwnPropertyNames:H,getOwnPropertySymbols:J}),S&&c(c.S+c.F*(!A||G),"JSON",{stringify:K}),l(O,"Symbol"),l(Math,"Math",!0),l(o.JSON,"JSON",!0)},function(t,e,n){var r=n(0),o=n(6);t.exports=function(t,e){for(var n,i=o(t),u=r.getKeys(i),c=u.length,a=0;c>a;)if(i[n=u[a++]]===e)return n}},function(t,e,n){var r=n(6),o=n(0).getNames,i={}.toString,u="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[],c=function(t){try{return o(t)}catch(t){return u.slice()}};t.exports.get=function(t){return u&&"[object Window]"==i.call(t)?c(t):o(r(t))}},function(t,e,n){var r=n(0);t.exports=function(t){var e=r.getKeys(t),n=r.getSymbols;if(n)for(var o,i=n(t),u=r.isEnum,c=0;i.length>c;)u.call(t,o=i[c++])&&e.push(o);return e}},function(t,e,n){var r=n(28);t.exports=Array.isArray||function(t){return"Array"==r(t)}},function(t,e){},function(t,e,n){t.exports={default:n(59),__esModule:!0}},function(t,e,n){n(60),t.exports=n(1).Object.setPrototypeOf},function(t,e,n){var r=n(2);r(r.S,"Object",{setPrototypeOf:n(61).set})},function(t,e,n){var r=n(0).getDesc,o=n(30),i=n(29),u=function(t,e){if(i(t),!o(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,o){try{o=n(16)(Function.call,r(Object.prototype,"__proto__").set,2),o(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return u(t,n),e?t.__proto__=n:o(t,n),t}}({},!1):void 0),check:u}},function(t,e,n){t.exports={default:n(63),__esModule:!0}},function(t,e,n){var r=n(0);t.exports=function(t,e){return r.create(t,e)}},function(t,e,n){t.exports={default:n(65),__esModule:!0}},function(t,e,n){n(66),t.exports=n(1).Object.keys},function(t,e,n){var r=n(7);n(15)("keys",function(t){return function(e){return t(r(e))}})},function(t,e){},function(t,e,n){"use strict";var t,r=n(69),o=n.n(r),i=n(14),u=n.n(i),c=n(17),a=n.n(c),s=n(18),f=n.n(s),l=n(19),p=n.n(l),d=n(31),y=n.n(d),m=n(32),v=n.n(m),g=n(33),h=n.n(g),_=function(e){function n(e,r){a()(this,n);var i=p()(this,(n.__proto__||u()(n)).call(this,e));b.call(i),i.uikit=r.uikit;var c=e.description;if(!c&&e.id&&(c=_reg("Panels",e.id)),!c&&e.id){var s=e.type?e.type:"block";c=o()({type:s},e)}console.log("creating panel",c,e,r,i.context);var f=e.className?e.className+" panel ":" panel ";if(c.id&&(f=f+" "+c.id),c.name&&(f=f+" "+c.name),c.className&&(f=f+" "+c.className),c&&"string"==typeof c)i.processBlock(c,e,f);else if(c)switch(c.type){case"view":f+=" view ",i.processView(c,e,r,f);break;case"entity":f+=" entity ",i.processEntity(c,e,r,f);break;case"form":f+=" form ",i.processForm(c,e,r,f);break;case"html":f+=" html ",i.processHtml(c,e,r,f);break;case"block":f+=" block ",i.processBlock(c,e,r,f);break;case"layout":f+=" layout ",i.processLayout(c,e,r,f);break;case"component":if(c.component)i.getView=function(t,e,n){return c.component};else{var l=i.getComponent(c.module,c.componentName,t.req);console.log("rendering component",c,l);var d={description:c,className:f},y=c.props?o()({},c.props,d):d;i.getView=function(t,e){return function(n,r,o){return console.log("rendering comp",e,t),v.a.createElement(e,t)}}(y,l)}break;default:var m=_reg("PanelTypes",c.type);i.getView=function(t,e,n){return m.getComponent(c,t,e,n)}}return i}return y()(n,e),f()(n,[{key:"getDisplayFunc",value:function(t,e){if(console.log("getting block",t),!t)return null;if("string"==typeof t)return _reg("Blocks",t);var n=t.block?t.block:t.id,r=_reg("Blocks",n);return r||(r=_reg("Blocks",t.defaultBlock)),r}},{key:"render",value:function(){return console.log("rendering panel",this.props),this.getView?this.getView(this.props,this.context,this.state):v.a.createElement(this.context.uikit.Block,null)}}],[{key:"setModule",value:function(e){t=e}}]),n}(v.a.Component),b=function(){var e=this;this.getPanelItems=function(t){if(!t)return null;if(t instanceof Array){return t.map(function(t){return v.a.createElement(_,{description:t})})}return v.a.createElement(_,{description:t})},this.processLayout=function(t,n,r,o){var i=t;if(t.id&&(i=_reg("Panels",t.id)),i.layout){var u=e,c=null,a=function(t){return i[t]?v.a.createElement(u.uikit.Block,{className:t},u.getPanelItems(i[t])):null};switch(i.layout){case"2col":c=v.a.createElement(e.uikit.Block,{className:o+" twocol"},a("header"),v.a.createElement(e.uikit.Block,{className:"row"},a("left"),a("right")),a("footer"));break;case"3col":c=v.a.createElement(e.uikit.Block,{className:o+" threecol"},a("header"),v.a.createElement(e.uikit.Block,{className:"row"},a("left"),a("right")),a("footer"));break;default:c=v.a.createElement(e.uikit.Block,{className:o},a("items"))}e.getView=function(t,e,n){return c}}},this.processBlock=function(t,n,r,o){var i=e.getDisplayFunc(t,n);console.log("processing block",t,i,n),e.getView=i?function(e,n,r){return console.log("calling block func",e,n),i({data:e.data,className:o,routeParams:n.routeParams,storage:Storage},t,n.uikit)}:function(t,e,n){return v.a.createElement(e.uikit.Block,null)}},this.createMarkup=function(t){return{__html:t}},this.processHtml=function(t,n,r,o){t.html?e.getView=function(e,n,r){return console.log("rendering html",t.html),v.a.createElement("div",{className:o,dangerouslySetInnerHTML:this.createMarkup(t.html)})}:e.getView=function(t,e,n){return v.a.createElement(e.uikit.Block,null)}},this.processForm=function(n,r,i,u){console.log("processing form",n);var c=n;if(n.id&&(c=_reg("Forms",n.id)),c){console.log("processing form",n,c);var a=c.info;e.form||(e.form=e.getComponent("reactforms","Form",t.req)),e.form?e.getView=function(t,e,r){var i=o()({},a,e.routeParams);return console.log("form cfg",i,a),v.a.createElement(this.form,{form:n.id,formContext:{data:t.data,routeParams:e.routeParams,storage:Storage},config:i,onSubmit:t.onSubmit,actions:t.actions,description:c,id:n.id})}:e.getView=function(t,e,n){return v.a.createElement(e.uikit.Block,null)}}},this.processView=function(n,r,i,u){var c=n.viewid?n.viewid:n.id,a=n;c&&(a=_reg("Views",c));var s=o()({},a,n),f=s.header?v.a.createElement(_,{description:s.header}):null;e.view||(e.view=e.getComponent("laatooviews","View",t.req)),e.getView=function(t,e,n){return v.a.createElement(this.view,{params:t.params,description:a,className:u,header:f,id:c},v.a.createElement(_,{description:s.item}))}},this.processEntity=function(n,r,o,i){var u=n.entityDisplay?n.entityDisplay:"default";console.log("view entity description",n,u,r);var c="",a="";c=o.routeParams&&o.routeParams.entityId?o.routeParams.entityId:n.entityId,a=n.entityName,e.entity||(e.entity=e.getComponent("laatooviews","Entity",t.req));var s="";r.index&&(s=r.index%2?"oddindex":"evenindex");var f={type:"block",block:n.entityName+"_"+u,defaultBlock:n.entityName+"_default"};e.getView=function(t,e,r){return v.a.createElement(this.entity,{id:c,name:a,entityDescription:n,data:t.data,index:t.index,uikit:e.uikit},v.a.createElement(_,{description:f,className:s}))}},this.getComponent=function(e,n,r){var o=e+n,i=t[o];if(!i){var u=r(e);u&&n&&(i=u[n],t[o]=i)}return i}};_.contextTypes={uikit:h.a.object,routeParams:h.a.object},e.a=_},function(t,e,n){t.exports={default:n(70),__esModule:!0}},function(t,e,n){n(71),t.exports=n(1).Object.assign},function(t,e,n){var r=n(2);r(r.S+r.F,"Object",{assign:n(72)})},function(t,e,n){var r=n(0),o=n(7),i=n(27);t.exports=n(5)(function(){var t=Object.assign,e={},n={},r=Symbol(),o="abcdefghijklmnopqrst";return e[r]=7,o.split("").forEach(function(t){n[t]=t}),7!=t({},e)[r]||Object.keys(t({},n)).join("")!=o})?function(t,e){for(var n=o(t),u=arguments,c=u.length,a=1,s=r.getKeys,f=r.getSymbols,l=r.isEnum;c>a;)for(var p,d=i(u[a++]),y=f?s(d).concat(f(d)):s(d),m=y.length,v=0;m>v;)l.call(d,p=y[v++])&&(n[p]=d[p]);return n}:Object.assign},function(t,e,n){"use strict";var r=n(74),o=n(75),i=n(76);t.exports=function(){function t(t,e,n,r,u,c){c!==i&&o(!1,"Calling PropTypes validators directly is not supported by the `prop-types` package. Use PropTypes.checkPropTypes() to call them. Read more at http://fb.me/use-check-prop-types")}function e(){return t}t.isRequired=t;var n={array:t,bool:t,func:t,number:t,object:t,string:t,symbol:t,any:t,arrayOf:e,element:t,instanceOf:e,node:t,objectOf:e,oneOf:e,oneOfType:e,shape:e,exact:e};return n.checkPropTypes=r,n.PropTypes=n,n}},function(t,e,n){"use strict";function r(t){return function(){return t}}var o=function(){};o.thatReturns=r,o.thatReturnsFalse=r(!1),o.thatReturnsTrue=r(!0),o.thatReturnsNull=r(null),o.thatReturnsThis=function(){return this},o.thatReturnsArgument=function(t){return t},t.exports=o},function(t,e,n){"use strict";function r(t,e,n,r,i,u,c,a){if(o(e),!t){var s;if(void 0===e)s=new Error("Minified exception occurred; use the non-minified dev environment for the full error message and additional helpful warnings.");else{var f=[n,r,i,u,c,a],l=0;s=new Error(e.replace(/%s/g,function(){return f[l++]})),s.name="Invariant Violation"}throw s.framesToPop=1,s}}var o=function(t){};t.exports=r},function(t,e,n){"use strict";t.exports="SECRET_DO_NOT_PASS_THIS_OR_YOU_WILL_BE_FIRED"},function(t,n){t.exports=e}])});
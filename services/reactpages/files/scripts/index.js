define("reactpages",["react","redux","reactwebcommon"],function(t,e,n){return function(t){function e(r){if(n[r])return n[r].exports;var o=n[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,e),o.l=!0,o.exports}var n={};return e.m=t,e.c=n,e.d=function(t,n,r){e.o(t,n)||Object.defineProperty(t,n,{configurable:!1,enumerable:!0,get:r})},e.n=function(t){var n=t&&t.__esModule?function(){return t.default}:function(){return t};return e.d(n,"a",n),n},e.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},e.p="/",e(e.s=34)}([function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,e,n){var r=n(4),o=n(1),i=n(16),a=function(t,e,n){var c,u,s,l=t&a.F,f=t&a.G,p=t&a.S,y=t&a.P,d=t&a.B,m=t&a.W,h=f?o:o[e]||(o[e]={}),v=f?r:p?r[e]:(r[e]||{}).prototype;f&&(n=e);for(c in n)(u=!l&&v&&c in v)&&c in h||(s=u?v[c]:n[c],h[c]=f&&"function"!=typeof v[c]?n[c]:d&&u?i(s,r):m&&v[c]==s?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(s):y&&"function"==typeof s?i(Function.call,s):s,y&&((h.prototype||(h.prototype={}))[c]=s))};a.F=1,a.G=2,a.S=4,a.P=8,a.B=16,a.W=32,t.exports=a},function(t,e,n){var r=n(25)("wks"),o=n(26),i=n(4).Symbol;t.exports=function(t){return r[t]||(r[t]=i&&i[t]||(i||o)("Symbol."+t))}},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e,n){var r=n(27),o=n(8);t.exports=function(t){return r(o(t))}},function(t,e,n){var r=n(8);t.exports=function(t){return Object(r(t))}},function(t,e){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e,n){var r=n(0),o=n(10);t.exports=n(24)?function(t,e,n){return r.setDesc(t,e,o(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports={}},function(t,e,n){var r=n(0).setDesc,o=n(11),i=n(3)("toStringTag");t.exports=function(t,e,n){t&&!o(t=n?t:t.prototype,i)&&r(t,i,{configurable:!0,value:e})}},function(t,e,n){t.exports={default:n(35),__esModule:!0}},function(t,e,n){var r=n(2),o=n(1),i=n(5);t.exports=function(t,e){var n=(o.Object||{})[t]||Object[t],a={};a[t]=e(n),r(r.S+r.F*i(function(){n(1)}),"Object",a)}},function(t,e,n){var r=n(37);t.exports=function(t,e,n){if(r(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,r){return t.call(e,n,r)};case 3:return function(n,r,o){return t.call(e,n,r,o)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){"use strict";e.__esModule=!0;var r=n(38),o=function(t){return t&&t.__esModule?t:{default:t}}(r);e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var r=e[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),(0,o.default)(t,r.key,r)}}return function(e,n,r){return n&&t(e.prototype,n),r&&t(e,r),e}}()},function(t,e,n){"use strict";e.__esModule=!0;var r=n(20),o=function(t){return t&&t.__esModule?t:{default:t}}(r);e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,o.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){"use strict";function r(t){return t&&t.__esModule?t:{default:t}}e.__esModule=!0;var o=n(40),i=r(o),a=n(50),c=r(a),u="function"==typeof c.default&&"symbol"==typeof i.default?function(t){return typeof t}:function(t){return t&&"function"==typeof c.default&&t.constructor===c.default&&t!==c.default.prototype?"symbol":typeof t};e.default="function"==typeof c.default&&"symbol"===u(i.default)?function(t){return void 0===t?"undefined":u(t)}:function(t){return t&&"function"==typeof c.default&&t.constructor===c.default&&t!==c.default.prototype?"symbol":void 0===t?"undefined":u(t)}},function(t,e,n){"use strict";var r=n(22),o=n(2),i=n(23),a=n(9),c=n(11),u=n(12),s=n(45),l=n(13),f=n(0).getProto,p=n(3)("iterator"),y=!([].keys&&"next"in[].keys()),d=function(){return this};t.exports=function(t,e,n,m,h,v,g){s(n,e,m);var _,b,k=function(t){if(!y&&t in O)return O[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},x=e+" Iterator",w="values"==h,P=!1,O=t.prototype,E=O[p]||O["@@iterator"]||h&&O[h],S=E||k(h);if(E){var N=f(S.call(new t));l(N,x,!0),!r&&c(O,"@@iterator")&&a(N,p,d),w&&"values"!==E.name&&(P=!0,S=function(){return E.call(this)})}if(r&&!g||!y&&!P&&O[p]||a(O,p,S),u[e]=S,u[x]=d,h)if(_={values:w?S:k("values"),keys:v?S:k("keys"),entries:w?k("entries"):S},g)for(b in _)b in O||i(O,b,_[b]);else o(o.P+o.F*(y||P),e,_);return _}},function(t,e){t.exports=!0},function(t,e,n){t.exports=n(9)},function(t,e,n){t.exports=!n(5)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){var r=n(4),o=r["__core-js_shared__"]||(r["__core-js_shared__"]={});t.exports=function(t){return o[t]||(o[t]={})}},function(t,e){var n=0,r=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+r).toString(36))}},function(t,e,n){var r=n(28);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==r(t)?t.split(""):Object(t)}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){var r=n(30);t.exports=function(t){if(!r(t))throw TypeError(t+" is not an object!");return t}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){"use strict";function r(t){return t&&t.__esModule?t:{default:t}}e.__esModule=!0;var o=n(58),i=r(o),a=n(62),c=r(a),u=n(20),s=r(u);e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,s.default)(e)));t.prototype=(0,c.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(i.default?(0,i.default)(t,e):t.__proto__=e)}},function(e,n){e.exports=t},function(t,e,n){t.exports=n(73)()},function(t,e,n){"use strict";function r(e,n,r,o,i,a){t.properties=Application.Properties[n],t.settings=o,t.req=a,Window.redirectPage||(Window.redirectPage=function(t,e){var n=_reg("Pages",t);if(console.log("redirect page",n),n){var r=formatUrl(n.url,e);Window.redirect(r)}}),k.a.setModule(t)}function o(t,e){var n=Application.AllRegItems("Pages");if(n)for(var r in n)try{!function(){var o=n[r],a=i(o),c=o.components;o.component&&(c={main:o.component});var u={};v()(c).forEach(function(t){u[t]=function(e,n){return function(r){return _.a.createElement(P,{pageId:n,placeholder:t,routerState:r,description:e})}}(c[t],r)});var s={pattern:o.route,components:u,reducer:Object(x.combineReducers)(a)},l=s;t&&t.ProcessRoute&&(l=t.ProcessRoute(s,e)),Application.Register("Routes",r,l),Application.Register("Actions","Page_"+r,{url:l.pattern})}()}catch(t){console.log(t)}}function i(e){var n={};for(var r in e.datasources)try{var o=_reg("Datasources",r),i={};o.type;var a=o.module;if(a){var c=t.req(a);c&&(i=c[o.processor])}i&&(n[r]=i)}catch(t){}return n}Object.defineProperty(e,"__esModule",{value:!0}),n.d(e,"Initialize",function(){return r}),n.d(e,"ProcessPages",function(){return o});var a=n(14),c=n.n(a),u=n(17),s=n.n(u),l=n(18),f=n.n(l),p=n(19),y=n.n(p),d=n(31),m=n.n(d),h=n(64),v=n.n(h),g=n(32),_=n.n(g),b=n(67),k=(n.n(b),n(68)),x=n(78);n.n(x);n.d(e,"Panel",function(){return k.a});var w=n(33),t=this,P=function(t){function e(){return s()(this,e),y()(this,(e.__proto__||c()(e)).apply(this,arguments))}return m()(e,t),f()(e,[{key:"getChildContext",value:function(){return{routeParams:this.props.routerState.params}}},{key:"render",value:function(){var t=this.props.pageId+this.props.placeholder;return _.a.createElement(k.a,{key:t,description:this.props.description})}}]),e}(_.a.Component);P.childContextTypes={routeParams:w.object}},function(t,e,n){n(36),t.exports=n(1).Object.getPrototypeOf},function(t,e,n){var r=n(7);n(15)("getPrototypeOf",function(t){return function(e){return t(r(e))}})},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){t.exports={default:n(39),__esModule:!0}},function(t,e,n){var r=n(0);t.exports=function(t,e,n){return r.setDesc(t,e,n)}},function(t,e,n){t.exports={default:n(41),__esModule:!0}},function(t,e,n){n(42),n(46),t.exports=n(3)("iterator")},function(t,e,n){"use strict";var r=n(43)(!0);n(21)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=r(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){var r=n(44),o=n(8);t.exports=function(t){return function(e,n){var i,a,c=String(o(e)),u=r(n),s=c.length;return u<0||u>=s?t?"":void 0:(i=c.charCodeAt(u),i<55296||i>56319||u+1===s||(a=c.charCodeAt(u+1))<56320||a>57343?t?c.charAt(u):i:t?c.slice(u,u+2):a-56320+(i-55296<<10)+65536)}}},function(t,e){var n=Math.ceil,r=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?r:n)(t)}},function(t,e,n){"use strict";var r=n(0),o=n(10),i=n(13),a={};n(9)(a,n(3)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=r.create(a,{next:o(1,n)}),i(t,e+" Iterator")}},function(t,e,n){n(47);var r=n(12);r.NodeList=r.HTMLCollection=r.Array},function(t,e,n){"use strict";var r=n(48),o=n(49),i=n(12),a=n(6);t.exports=n(21)(Array,"Array",function(t,e){this._t=a(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,o(1)):"keys"==e?o(0,n):"values"==e?o(0,t[n]):o(0,[n,t[n]])},"values"),i.Arguments=i.Array,r("keys"),r("values"),r("entries")},function(t,e){t.exports=function(){}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e,n){t.exports={default:n(51),__esModule:!0}},function(t,e,n){n(52),n(57),t.exports=n(1).Symbol},function(t,e,n){"use strict";var r=n(0),o=n(4),i=n(11),a=n(24),c=n(2),u=n(23),s=n(5),l=n(25),f=n(13),p=n(26),y=n(3),d=n(53),m=n(54),h=n(55),v=n(56),g=n(29),_=n(6),b=n(10),k=r.getDesc,x=r.setDesc,w=r.create,P=m.get,O=o.Symbol,E=o.JSON,S=E&&E.stringify,N=!1,j=y("_hidden"),C=r.isEnum,M=l("symbol-registry"),T=l("symbols"),B="function"==typeof O,A=Object.prototype,D=a&&s(function(){return 7!=w(x({},"a",{get:function(){return x(this,"a",{value:7}).a}})).a})?function(t,e,n){var r=k(A,e);r&&delete A[e],x(t,e,n),r&&t!==A&&x(A,e,r)}:x,I=function(t){var e=T[t]=w(O.prototype);return e._k=t,a&&N&&D(A,t,{configurable:!0,set:function(e){i(this,j)&&i(this[j],t)&&(this[j][t]=!1),D(this,t,b(1,e))}}),e},F=function(t){return"symbol"==typeof t},R=function(t,e,n){return n&&i(T,e)?(n.enumerable?(i(t,j)&&t[j][e]&&(t[j][e]=!1),n=w(n,{enumerable:b(0,!1)})):(i(t,j)||x(t,j,b(1,{})),t[j][e]=!0),D(t,e,n)):x(t,e,n)},V=function(t,e){g(t);for(var n,r=h(e=_(e)),o=0,i=r.length;i>o;)R(t,n=r[o++],e[n]);return t},q=function(t,e){return void 0===e?w(t):V(w(t),e)},W=function(t){var e=C.call(this,t);return!(e||!i(this,t)||!i(T,t)||i(this,j)&&this[j][t])||e},L=function(t,e){var n=k(t=_(t),e);return!n||!i(T,e)||i(t,j)&&t[j][e]||(n.enumerable=!0),n},H=function(t){for(var e,n=P(_(t)),r=[],o=0;n.length>o;)i(T,e=n[o++])||e==j||r.push(e);return r},J=function(t){for(var e,n=P(_(t)),r=[],o=0;n.length>o;)i(T,e=n[o++])&&r.push(T[e]);return r},K=function(t){if(void 0!==t&&!F(t)){for(var e,n,r=[t],o=1,i=arguments;i.length>o;)r.push(i[o++]);return e=r[1],"function"==typeof e&&(n=e),!n&&v(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!F(e))return e}),r[1]=e,S.apply(E,r)}},G=s(function(){var t=O();return"[null]"!=S([t])||"{}"!=S({a:t})||"{}"!=S(Object(t))});B||(O=function(){if(F(this))throw TypeError("Symbol is not a constructor");return I(p(arguments.length>0?arguments[0]:void 0))},u(O.prototype,"toString",function(){return this._k}),F=function(t){return t instanceof O},r.create=q,r.isEnum=W,r.getDesc=L,r.setDesc=R,r.setDescs=V,r.getNames=m.get=H,r.getSymbols=J,a&&!n(22)&&u(A,"propertyIsEnumerable",W,!0));var U={for:function(t){return i(M,t+="")?M[t]:M[t]=O(t)},keyFor:function(t){return d(M,t)},useSetter:function(){N=!0},useSimple:function(){N=!1}};r.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=y(t);U[t]=B?e:I(e)}),N=!0,c(c.G+c.W,{Symbol:O}),c(c.S,"Symbol",U),c(c.S+c.F*!B,"Object",{create:q,defineProperty:R,defineProperties:V,getOwnPropertyDescriptor:L,getOwnPropertyNames:H,getOwnPropertySymbols:J}),E&&c(c.S+c.F*(!B||G),"JSON",{stringify:K}),f(O,"Symbol"),f(Math,"Math",!0),f(o.JSON,"JSON",!0)},function(t,e,n){var r=n(0),o=n(6);t.exports=function(t,e){for(var n,i=o(t),a=r.getKeys(i),c=a.length,u=0;c>u;)if(i[n=a[u++]]===e)return n}},function(t,e,n){var r=n(6),o=n(0).getNames,i={}.toString,a="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[],c=function(t){try{return o(t)}catch(t){return a.slice()}};t.exports.get=function(t){return a&&"[object Window]"==i.call(t)?c(t):o(r(t))}},function(t,e,n){var r=n(0);t.exports=function(t){var e=r.getKeys(t),n=r.getSymbols;if(n)for(var o,i=n(t),a=r.isEnum,c=0;i.length>c;)a.call(t,o=i[c++])&&e.push(o);return e}},function(t,e,n){var r=n(28);t.exports=Array.isArray||function(t){return"Array"==r(t)}},function(t,e){},function(t,e,n){t.exports={default:n(59),__esModule:!0}},function(t,e,n){n(60),t.exports=n(1).Object.setPrototypeOf},function(t,e,n){var r=n(2);r(r.S,"Object",{setPrototypeOf:n(61).set})},function(t,e,n){var r=n(0).getDesc,o=n(30),i=n(29),a=function(t,e){if(i(t),!o(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,o){try{o=n(16)(Function.call,r(Object.prototype,"__proto__").set,2),o(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return a(t,n),e?t.__proto__=n:o(t,n),t}}({},!1):void 0),check:a}},function(t,e,n){t.exports={default:n(63),__esModule:!0}},function(t,e,n){var r=n(0);t.exports=function(t,e){return r.create(t,e)}},function(t,e,n){t.exports={default:n(65),__esModule:!0}},function(t,e,n){n(66),t.exports=n(1).Object.keys},function(t,e,n){var r=n(7);n(15)("keys",function(t){return function(e){return t(r(e))}})},function(t,e){},function(t,e,n){"use strict";var t,r=n(69),o=n.n(r),i=n(14),a=n.n(i),c=n(17),u=n.n(c),s=n(18),l=n.n(s),f=n(19),p=n.n(f),y=n(31),d=n.n(y),m=n(32),h=n.n(m),v=n(33),g=n.n(v),_=n(77),b=(n.n(_),function(e){function n(e,r){u()(this,n);var i=p()(this,(n.__proto__||a()(n)).call(this,e));k.call(i),i.uikit=r.uikit;var c=e.description;if(!c&&e.id&&(c=_reg("Panels",e.id)),!c&&e.id){var s=e.type?e.type:"block";c=o()({type:s},e)}i.title=e.title?e.title:c.title?c.title:null,i.closePanel=e.closePanel?e.closePanel:null,console.log("creating panel",c,e,r,i.context);var l=e.className?e.className+" panel ":" panel ";if(c.id&&(l=l+" "+c.id),c.name&&(l=l+" "+c.name),c.className&&(l=l+" "+c.className),i.overlay=e.overlay?e.overlay:null,c&&"string"==typeof c)i.processBlock(c,e);else if(c)switch(c.type){case"view":l+=" view ",i.processView(c,e,r);break;case"entity":l+=" entity ",i.processEntity(c,e,r);break;case"form":l+=" form ",i.processForm(c,e,r);break;case"html":l+=" html ",i.processHtml(c,e,r);break;case"block":l+=" block ",i.processBlock(c,e,r);break;case"layout":l+=" layout ",i.processLayout(c,e,r);break;case"component":if(c.component)i.getView=function(t,e,n){return c.component};else{var f=i.getComponent(c.module,c.componentName,t.req);i.getView=function(t){return function(e,n,r,i){var a={description:c,className:i},u=c.props?o()({},c.props,a):a;return h.a.createElement(t,u)}}(f)}break;default:var y=_reg("PanelTypes",c.type);i.getView=function(t,e,n){return y.getComponent(c,t,e,n)}}return console.log("class name",i.className,l,c),i.className=l,i}return d()(n,e),l()(n,[{key:"getDisplayFunc",value:function(t,e){if(console.log("getting block",t),!t)return null;if("string"==typeof t)return _reg("Blocks",t);var n=t.block?t.block:t.id,r=_reg("Blocks",n);return r||(r=_reg("Blocks",t.defaultBlock)),r}},{key:"getChildContext",value:function(){return console.log("getting child contextoverlaying component",this.overlay,this.props,this.context),this.overlay?{overlayComponent:this.overlayComponent,closeOverlay:this.closeOverlay}:this.context&&this.context.overlayComponent?{overlayComponent:this.context.overlayComponent,closeOverlay:this.closeOverlay}:null}},{key:"render",value:function(){console.log("rendering panel",this.props);var t=null;return t=this.overlay&&this.state&&this.state.overlayComponent?this.state.overlayComponent:this.getView?this.getView(this.props,this.context,this.state,this.title?"":this.className):h.a.createElement(this.context.uikit.Block,null),this.title?h.a.createElement(this.uikit.Block,{className:this.className},h.a.createElement(this.uikit.Block,{className:"titlebar"},h.a.createElement(this.uikit.Block,{className:"title left"},this.title),this.closePanel?h.a.createElement(_.Action,{className:"right close action",action:{actiontype:"method"},method:this.closePanel},h.a.createElement(this.uikit.Icons.CloseIcon,null)):null),h.a.createElement(this.uikit.Block,null,t)):t}}],[{key:"setModule",value:function(e){t=e}}]),n}(h.a.Component)),k=function(){var e=this;this.getPanelItems=function(t){if(!t)return null;if(t instanceof Array){return t.map(function(t){return h.a.createElement(b,{description:t})})}return h.a.createElement(b,{description:t})},this.cfgPanel=function(t,n){!e.title&&t&&(e.title=t),!e.overlay&&n&&(e.overlay=n)},this.processLayout=function(t,n,r){var o=t;if(t.id&&(o=_reg("Panels",t.id)),o.layout){e.cfgPanel(o.title,o.overlay);var i=e,a=null,c=function(t){return o[t]?h.a.createElement(i.uikit.Block,{className:t},i.getPanelItems(o[t])):null};e.getView=function(t,e,n,r){switch(o.layout){case"2col":a=h.a.createElement(this.uikit.Block,{className:r+" twocol"},c("header"),h.a.createElement(this.uikit.Block,{className:"row"},c("left"),c("right")),c("footer"));break;case"3col":a=h.a.createElement(this.uikit.Block,{className:r+" threecol"},c("header"),h.a.createElement(this.uikit.Block,{className:"row"},c("left"),c("right")),c("footer"));break;default:a=h.a.createElement(this.uikit.Block,{className:r},c("items"))}return a}}},this.processBlock=function(t,n,r){var o=e.getDisplayFunc(t,n);e.cfgPanel(t.title,t.overlay),console.log("processing block",t,o,n),e.getView=o?function(e,n,r,i){return console.log("calling block func",e,n),o({data:e.data,className:i,routeParams:n.routeParams,storage:Storage},t,n.uikit)}:function(t,e,n,r){return h.a.createElement(e.uikit.Block,null)}},this.createMarkup=function(t){return{__html:t}},this.processHtml=function(t,n,r){e.cfgPanel(t.title,t.overlay),t.html?e.getView=function(e,n,r,o){return console.log("rendering html",t.html),h.a.createElement("div",{className:o,dangerouslySetInnerHTML:this.createMarkup(t.html)})}:e.getView=function(t,e,n,r){return h.a.createElement(e.uikit.Block,null)}},this.processForm=function(n,r,i){console.log("processing form",n);var a=n;if(n.id&&(a=_reg("Forms",n.id)),a){console.log("processing form",n),e.cfgPanel(a.info.title,a.info.overlay);var c=a.info;e.form||(e.form=e.getComponent("reactforms","Form",t.req)),e.form?e.getView=function(t,e,r,i){var u=o()({},c,e.routeParams);return console.log("form cfg",u,c),h.a.createElement(this.form,{form:n.id,formContext:{data:t.data,routeParams:e.routeParams,storage:Storage},config:u,inline:t.inline,onSubmit:t.onSubmit,actions:t.actions,description:a,className:i,id:n.id})}:e.getView=function(t,e,n,r){return h.a.createElement(e.uikit.Block,null)}}},this.processView=function(n,r,i){var a=n.viewid?n.viewid:n.id,c=n;a&&(c=_reg("Views",a)),e.cfgPanel(c.title,c.overlay);var u=o()({},c,n),s=u.header?h.a.createElement(b,{description:u.header}):null;e.view||(e.view=e.getComponent("laatooviews","View",t.req)),e.getView=function(t,e,n,r){return h.a.createElement(this.view,{params:t.params,description:c,className:r,header:s,id:a},h.a.createElement(b,{description:u.item}))}},this.processEntity=function(n,r,o){var i=n.entityDisplay?n.entityDisplay:"default";console.log("view entity description",n,i,r);var a="",c="";a=o.routeParams&&o.routeParams.entityId?o.routeParams.entityId:n.entityId,c=n.entityName,e.cfgPanel(n.title,n.overlay),e.entity||(e.entity=e.getComponent("laatooviews","Entity",t.req));var u="";r.index&&(u=r.index%2?"oddindex":"evenindex");var s={type:"block",block:n.entityName+"_"+i,defaultBlock:n.entityName+"_default"};e.getView=function(t,e,r,o){return h.a.createElement(this.entity,{id:a,name:c,entityDescription:n,data:t.data,index:t.index,uikit:e.uikit},h.a.createElement(b,{description:s,className:u}))}},this.overlayComponent=function(t){console.log("overlaying component"),e.setState(o()({},{overlayComponent:t}))},this.closeOverlay=function(){e.setState({})},this.getComponent=function(e,n,r){var o=e+n,i=t[o];if(!i){var a=r(e);a&&n&&(i=a[n],t[o]=i)}return i}};b.childContextTypes={overlayComponent:g.a.func,closeOverlay:g.a.func},b.contextTypes={overlayComponent:g.a.func,closeOverlay:g.a.func,uikit:g.a.object,routeParams:g.a.object},e.a=b},function(t,e,n){t.exports={default:n(70),__esModule:!0}},function(t,e,n){n(71),t.exports=n(1).Object.assign},function(t,e,n){var r=n(2);r(r.S+r.F,"Object",{assign:n(72)})},function(t,e,n){var r=n(0),o=n(7),i=n(27);t.exports=n(5)(function(){var t=Object.assign,e={},n={},r=Symbol(),o="abcdefghijklmnopqrst";return e[r]=7,o.split("").forEach(function(t){n[t]=t}),7!=t({},e)[r]||Object.keys(t({},n)).join("")!=o})?function(t,e){for(var n=o(t),a=arguments,c=a.length,u=1,s=r.getKeys,l=r.getSymbols,f=r.isEnum;c>u;)for(var p,y=i(a[u++]),d=l?s(y).concat(l(y)):s(y),m=d.length,h=0;m>h;)f.call(y,p=d[h++])&&(n[p]=y[p]);return n}:Object.assign},function(t,e,n){"use strict";var r=n(74),o=n(75),i=n(76);t.exports=function(){function t(t,e,n,r,a,c){c!==i&&o(!1,"Calling PropTypes validators directly is not supported by the `prop-types` package. Use PropTypes.checkPropTypes() to call them. Read more at http://fb.me/use-check-prop-types")}function e(){return t}t.isRequired=t;var n={array:t,bool:t,func:t,number:t,object:t,string:t,symbol:t,any:t,arrayOf:e,element:t,instanceOf:e,node:t,objectOf:e,oneOf:e,oneOfType:e,shape:e,exact:e};return n.checkPropTypes=r,n.PropTypes=n,n}},function(t,e,n){"use strict";function r(t){return function(){return t}}var o=function(){};o.thatReturns=r,o.thatReturnsFalse=r(!1),o.thatReturnsTrue=r(!0),o.thatReturnsNull=r(null),o.thatReturnsThis=function(){return this},o.thatReturnsArgument=function(t){return t},t.exports=o},function(t,e,n){"use strict";function r(t,e,n,r,i,a,c,u){if(o(e),!t){var s;if(void 0===e)s=new Error("Minified exception occurred; use the non-minified dev environment for the full error message and additional helpful warnings.");else{var l=[n,r,i,a,c,u],f=0;s=new Error(e.replace(/%s/g,function(){return l[f++]})),s.name="Invariant Violation"}throw s.framesToPop=1,s}}var o=function(t){};t.exports=r},function(t,e,n){"use strict";t.exports="SECRET_DO_NOT_PASS_THIS_OR_YOU_WILL_BE_FIRED"},function(t,e){t.exports=n},function(t,n){t.exports=e}])});
define("reactpages",["react","prop-types","redux","reactwebcommon"],function(e,t,n,o){return function(e){var t={};function n(o){if(t[o])return t[o].exports;var r=t[o]={i:o,l:!1,exports:{}};return e[o].call(r.exports,r,r.exports,n),r.l=!0,r.exports}return n.m=e,n.c=t,n.d=function(e,t,o){n.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:o})},n.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},n.t=function(e,t){if(1&t&&(e=n(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var o=Object.create(null);if(n.r(o),Object.defineProperty(o,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var r in e)n.d(o,r,function(t){return e[t]}.bind(null,r));return o},n.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return n.d(t,"a",t),t},n.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},n.p="/",n(n.s=73)}([function(t,n){t.exports=e},function(e,t){var n=Object;e.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(e,n){e.exports=t},function(e,t){var n=e.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(e,t,n){var o=n(7),r=n(3),i=n(18),a=function(e,t,n){var c,u,s,l=e&a.F,f=e&a.G,p=e&a.S,y=e&a.P,m=e&a.B,d=e&a.W,v=f?r:r[t]||(r[t]={}),g=f?o:p?o[t]:(o[t]||{}).prototype;for(c in f&&(n=t),n)(u=!l&&g&&c in g)&&c in v||(s=u?g[c]:n[c],v[c]=f&&"function"!=typeof g[c]?n[c]:m&&u?i(s,o):d&&g[c]==s?function(e){var t=function(t){return this instanceof e?new e(t):e(t)};return t.prototype=e.prototype,t}(s):y&&"function"==typeof s?i(Function.call,s):s,y&&((v.prototype||(v.prototype={}))[c]=s))};a.F=1,a.G=2,a.S=4,a.P=8,a.B=16,a.W=32,e.exports=a},function(e,t,n){var o=n(26)("wks"),r=n(27),i=n(7).Symbol;e.exports=function(e){return o[e]||(o[e]=i&&i[e]||(i||r)("Symbol."+e))}},function(e,t,n){e.exports={default:n(41),__esModule:!0}},function(e,t){var n=e.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(e,t){e.exports=function(e){try{return!!e()}catch(e){return!0}}},function(e,t,n){var o=n(19),r=n(11);e.exports=function(e){return o(r(e))}},function(e,t,n){var o=n(11);e.exports=function(e){return Object(o(e))}},function(e,t){e.exports=function(e){if(void 0==e)throw TypeError("Can't call method on  "+e);return e}},function(e,t,n){var o=n(1),r=n(13);e.exports=n(25)?function(e,t,n){return o.setDesc(e,t,r(1,n))}:function(e,t,n){return e[t]=n,e}},function(e,t){e.exports=function(e,t){return{enumerable:!(1&e),configurable:!(2&e),writable:!(4&e),value:t}}},function(e,t){var n={}.hasOwnProperty;e.exports=function(e,t){return n.call(e,t)}},function(e,t){e.exports={}},function(e,t,n){var o=n(1).setDesc,r=n(14),i=n(5)("toStringTag");e.exports=function(e,t,n){e&&!r(e=n?e:e.prototype,i)&&o(e,i,{configurable:!0,value:t})}},function(e,t,n){var o=n(4),r=n(3),i=n(8);e.exports=function(e,t){var n=(r.Object||{})[e]||Object[e],a={};a[e]=t(n),o(o.S+o.F*i(function(){n(1)}),"Object",a)}},function(e,t,n){var o=n(39);e.exports=function(e,t,n){if(o(e),void 0===t)return e;switch(n){case 1:return function(n){return e.call(t,n)};case 2:return function(n,o){return e.call(t,n,o)};case 3:return function(n,o,r){return e.call(t,n,o,r)}}return function(){return e.apply(t,arguments)}}},function(e,t,n){var o=n(20);e.exports=Object("z").propertyIsEnumerable(0)?Object:function(e){return"String"==o(e)?e.split(""):Object(e)}},function(e,t){var n={}.toString;e.exports=function(e){return n.call(e).slice(8,-1)}},function(e,t,n){"use strict";t.__esModule=!0;var o=a(n(48)),r=a(n(58)),i="function"==typeof r.default&&"symbol"==typeof o.default?function(e){return typeof e}:function(e){return e&&"function"==typeof r.default&&e.constructor===r.default&&e!==r.default.prototype?"symbol":typeof e};function a(e){return e&&e.__esModule?e:{default:e}}t.default="function"==typeof r.default&&"symbol"===i(o.default)?function(e){return void 0===e?"undefined":i(e)}:function(e){return e&&"function"==typeof r.default&&e.constructor===r.default&&e!==r.default.prototype?"symbol":void 0===e?"undefined":i(e)}},function(e,t,n){"use strict";var o=n(23),r=n(4),i=n(24),a=n(12),c=n(14),u=n(15),s=n(53),l=n(16),f=n(1).getProto,p=n(5)("iterator"),y=!([].keys&&"next"in[].keys()),m=function(){return this};e.exports=function(e,t,n,d,v,g,h){s(n,t,d);var b,_,x=function(e){if(!y&&e in O)return O[e];switch(e){case"keys":case"values":return function(){return new n(this,e)}}return function(){return new n(this,e)}},k=t+" Iterator",w="values"==v,P=!1,O=e.prototype,S=O[p]||O["@@iterator"]||v&&O[v],j=S||x(v);if(S){var E=f(j.call(new e));l(E,k,!0),!o&&c(O,"@@iterator")&&a(E,p,m),w&&"values"!==S.name&&(P=!0,j=function(){return S.call(this)})}if(o&&!h||!y&&!P&&O[p]||a(O,p,j),u[t]=j,u[k]=m,v)if(b={values:w?j:x("values"),keys:g?j:x("keys"),entries:w?x("entries"):j},h)for(_ in b)_ in O||i(O,_,b[_]);else r(r.P+r.F*(y||P),t,b);return b}},function(e,t){e.exports=!0},function(e,t,n){e.exports=n(12)},function(e,t,n){e.exports=!n(8)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(e,t,n){var o=n(7),r=o["__core-js_shared__"]||(o["__core-js_shared__"]={});e.exports=function(e){return r[e]||(r[e]={})}},function(e,t){var n=0,o=Math.random();e.exports=function(e){return"Symbol(".concat(void 0===e?"":e,")_",(++n+o).toString(36))}},function(e,t,n){var o=n(29);e.exports=function(e){if(!o(e))throw TypeError(e+" is not an object!");return e}},function(e,t){e.exports=function(e){return"object"==typeof e?null!==e:"function"==typeof e}},function(e,t,n){e.exports={default:n(37),__esModule:!0}},function(e,t,n){e.exports={default:n(44),__esModule:!0}},function(e,t,n){"use strict";t.__esModule=!0,t.default=function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}},function(e,t,n){"use strict";t.__esModule=!0;var o=function(e){return e&&e.__esModule?e:{default:e}}(n(46));t.default=function(){function e(e,t){for(var n=0;n<t.length;n++){var r=t[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),(0,o.default)(e,r.key,r)}}return function(t,n,o){return n&&e(t.prototype,n),o&&e(t,o),t}}()},function(e,t,n){"use strict";t.__esModule=!0;var o=function(e){return e&&e.__esModule?e:{default:e}}(n(21));t.default=function(e,t){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!t||"object"!==(void 0===t?"undefined":(0,o.default)(t))&&"function"!=typeof t?e:t}},function(e,t,n){"use strict";t.__esModule=!0;var o=a(n(66)),r=a(n(70)),i=a(n(21));function a(e){return e&&e.__esModule?e:{default:e}}t.default=function(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Super expression must either be null or a function, not "+(void 0===t?"undefined":(0,i.default)(t)));e.prototype=(0,r.default)(t&&t.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),t&&(o.default?(0,o.default)(e,t):e.__proto__=t)}},function(e,t){e.exports=n},function(e,t,n){n(38),e.exports=n(3).Object.keys},function(e,t,n){var o=n(10);n(17)("keys",function(e){return function(t){return e(o(t))}})},function(e,t){e.exports=function(e){if("function"!=typeof e)throw TypeError(e+" is not a function!");return e}},function(e,t){},function(e,t,n){n(42),e.exports=n(3).Object.assign},function(e,t,n){var o=n(4);o(o.S+o.F,"Object",{assign:n(43)})},function(e,t,n){var o=n(1),r=n(10),i=n(19);e.exports=n(8)(function(){var e=Object.assign,t={},n={},o=Symbol(),r="abcdefghijklmnopqrst";return t[o]=7,r.split("").forEach(function(e){n[e]=e}),7!=e({},t)[o]||Object.keys(e({},n)).join("")!=r})?function(e,t){for(var n=r(e),a=arguments,c=a.length,u=1,s=o.getKeys,l=o.getSymbols,f=o.isEnum;c>u;)for(var p,y=i(a[u++]),m=l?s(y).concat(l(y)):s(y),d=m.length,v=0;d>v;)f.call(y,p=m[v++])&&(n[p]=y[p]);return n}:Object.assign},function(e,t,n){n(45),e.exports=n(3).Object.getPrototypeOf},function(e,t,n){var o=n(10);n(17)("getPrototypeOf",function(e){return function(t){return e(o(t))}})},function(e,t,n){e.exports={default:n(47),__esModule:!0}},function(e,t,n){var o=n(1);e.exports=function(e,t,n){return o.setDesc(e,t,n)}},function(e,t,n){e.exports={default:n(49),__esModule:!0}},function(e,t,n){n(50),n(54),e.exports=n(5)("iterator")},function(e,t,n){"use strict";var o=n(51)(!0);n(22)(String,"String",function(e){this._t=String(e),this._i=0},function(){var e,t=this._t,n=this._i;return n>=t.length?{value:void 0,done:!0}:(e=o(t,n),this._i+=e.length,{value:e,done:!1})})},function(e,t,n){var o=n(52),r=n(11);e.exports=function(e){return function(t,n){var i,a,c=String(r(t)),u=o(n),s=c.length;return u<0||u>=s?e?"":void 0:(i=c.charCodeAt(u))<55296||i>56319||u+1===s||(a=c.charCodeAt(u+1))<56320||a>57343?e?c.charAt(u):i:e?c.slice(u,u+2):a-56320+(i-55296<<10)+65536}}},function(e,t){var n=Math.ceil,o=Math.floor;e.exports=function(e){return isNaN(e=+e)?0:(e>0?o:n)(e)}},function(e,t,n){"use strict";var o=n(1),r=n(13),i=n(16),a={};n(12)(a,n(5)("iterator"),function(){return this}),e.exports=function(e,t,n){e.prototype=o.create(a,{next:r(1,n)}),i(e,t+" Iterator")}},function(e,t,n){n(55);var o=n(15);o.NodeList=o.HTMLCollection=o.Array},function(e,t,n){"use strict";var o=n(56),r=n(57),i=n(15),a=n(9);e.exports=n(22)(Array,"Array",function(e,t){this._t=a(e),this._i=0,this._k=t},function(){var e=this._t,t=this._k,n=this._i++;return!e||n>=e.length?(this._t=void 0,r(1)):r(0,"keys"==t?n:"values"==t?e[n]:[n,e[n]])},"values"),i.Arguments=i.Array,o("keys"),o("values"),o("entries")},function(e,t){e.exports=function(){}},function(e,t){e.exports=function(e,t){return{value:t,done:!!e}}},function(e,t,n){e.exports={default:n(59),__esModule:!0}},function(e,t,n){n(60),n(65),e.exports=n(3).Symbol},function(e,t,n){"use strict";var o=n(1),r=n(7),i=n(14),a=n(25),c=n(4),u=n(24),s=n(8),l=n(26),f=n(16),p=n(27),y=n(5),m=n(61),d=n(62),v=n(63),g=n(64),h=n(28),b=n(9),_=n(13),x=o.getDesc,k=o.setDesc,w=o.create,P=d.get,O=r.Symbol,S=r.JSON,j=S&&S.stringify,E=!1,C=y("_hidden"),N=o.isEnum,M=l("symbol-registry"),B=l("symbols"),D="function"==typeof O,V=Object.prototype,A=a&&s(function(){return 7!=w(k({},"a",{get:function(){return k(this,"a",{value:7}).a}})).a})?function(e,t,n){var o=x(V,t);o&&delete V[t],k(e,t,n),o&&e!==V&&k(V,t,o)}:k,F=function(e){var t=B[e]=w(O.prototype);return t._k=e,a&&E&&A(V,e,{configurable:!0,set:function(t){i(this,C)&&i(this[C],e)&&(this[C][e]=!1),A(this,e,_(1,t))}}),t},I=function(e){return"symbol"==typeof e},T=function(e,t,n){return n&&i(B,t)?(n.enumerable?(i(e,C)&&e[C][t]&&(e[C][t]=!1),n=w(n,{enumerable:_(0,!1)})):(i(e,C)||k(e,C,_(1,{})),e[C][t]=!0),A(e,t,n)):k(e,t,n)},R=function(e,t){h(e);for(var n,o=v(t=b(t)),r=0,i=o.length;i>r;)T(e,n=o[r++],t[n]);return e},q=function(e,t){return void 0===t?w(e):R(w(e),t)},W=function(e){var t=N.call(this,e);return!(t||!i(this,e)||!i(B,e)||i(this,C)&&this[C][e])||t},L=function(e,t){var n=x(e=b(e),t);return!n||!i(B,t)||i(e,C)&&e[C][t]||(n.enumerable=!0),n},H=function(e){for(var t,n=P(b(e)),o=[],r=0;n.length>r;)i(B,t=n[r++])||t==C||o.push(t);return o},J=function(e){for(var t,n=P(b(e)),o=[],r=0;n.length>r;)i(B,t=n[r++])&&o.push(B[t]);return o},K=s(function(){var e=O();return"[null]"!=j([e])||"{}"!=j({a:e})||"{}"!=j(Object(e))});D||(u((O=function(){if(I(this))throw TypeError("Symbol is not a constructor");return F(p(arguments.length>0?arguments[0]:void 0))}).prototype,"toString",function(){return this._k}),I=function(e){return e instanceof O},o.create=q,o.isEnum=W,o.getDesc=L,o.setDesc=T,o.setDescs=R,o.getNames=d.get=H,o.getSymbols=J,a&&!n(23)&&u(V,"propertyIsEnumerable",W,!0));var G={for:function(e){return i(M,e+="")?M[e]:M[e]=O(e)},keyFor:function(e){return m(M,e)},useSetter:function(){E=!0},useSimple:function(){E=!1}};o.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(e){var t=y(e);G[e]=D?t:F(t)}),E=!0,c(c.G+c.W,{Symbol:O}),c(c.S,"Symbol",G),c(c.S+c.F*!D,"Object",{create:q,defineProperty:T,defineProperties:R,getOwnPropertyDescriptor:L,getOwnPropertyNames:H,getOwnPropertySymbols:J}),S&&c(c.S+c.F*(!D||K),"JSON",{stringify:function(e){if(void 0!==e&&!I(e)){for(var t,n,o=[e],r=1,i=arguments;i.length>r;)o.push(i[r++]);return"function"==typeof(t=o[1])&&(n=t),!n&&g(t)||(t=function(e,t){if(n&&(t=n.call(this,e,t)),!I(t))return t}),o[1]=t,j.apply(S,o)}}}),f(O,"Symbol"),f(Math,"Math",!0),f(r.JSON,"JSON",!0)},function(e,t,n){var o=n(1),r=n(9);e.exports=function(e,t){for(var n,i=r(e),a=o.getKeys(i),c=a.length,u=0;c>u;)if(i[n=a[u++]]===t)return n}},function(e,t,n){var o=n(9),r=n(1).getNames,i={}.toString,a="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[];e.exports.get=function(e){return a&&"[object Window]"==i.call(e)?function(e){try{return r(e)}catch(e){return a.slice()}}(e):r(o(e))}},function(e,t,n){var o=n(1);e.exports=function(e){var t=o.getKeys(e),n=o.getSymbols;if(n)for(var r,i=n(e),a=o.isEnum,c=0;i.length>c;)a.call(e,r=i[c++])&&t.push(r);return t}},function(e,t,n){var o=n(20);e.exports=Array.isArray||function(e){return"Array"==o(e)}},function(e,t){},function(e,t,n){e.exports={default:n(67),__esModule:!0}},function(e,t,n){n(68),e.exports=n(3).Object.setPrototypeOf},function(e,t,n){var o=n(4);o(o.S,"Object",{setPrototypeOf:n(69).set})},function(e,t,n){var o=n(1).getDesc,r=n(29),i=n(28),a=function(e,t){if(i(e),!r(t)&&null!==t)throw TypeError(t+": can't set as prototype!")};e.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(e,t,r){try{(r=n(18)(Function.call,o(Object.prototype,"__proto__").set,2))(e,[]),t=!(e instanceof Array)}catch(e){t=!0}return function(e,n){return a(e,n),t?e.__proto__=n:r(e,n),e}}({},!1):void 0),check:a}},function(e,t,n){e.exports={default:n(71),__esModule:!0}},function(e,t,n){var o=n(1);e.exports=function(e,t){return o.create(e,t)}},function(e,t){e.exports=o},function(e,t,n){"use strict";n.r(t);var o,r=n(30),i=n.n(r),a=n(0),c=n.n(a),u=(n(40),n(6)),s=n.n(u),l=n(31),f=n.n(l),p=n(32),y=n.n(p),m=n(33),d=n.n(m),v=n(34),g=n.n(v),h=n(35),b=n.n(h),_=n(2),x=n.n(_),k=(n(72),function(e){function t(e,n){y()(this,t);var r=g()(this,(t.__proto__||f()(t)).call(this,e));w.call(r),r.uikit=n.uikit;var i=e.description,a=e.id?e.id:i&&i.id?i.id:null,u=e.type?e.type:i&&i.type?i.type:"layout";if(a){switch(u){case"view":i=_reg("Views",a);break;case"form":i=_reg("Forms",a);break;case"block":i=_reg("Blocks",a);break;default:i=_reg("Panels",a)}console.log("desc before assig",i,e),i=s()({type:u,id:a},i,e)}r.title=e.title?e.title:i&&i.title?i.title:null,r.closePanel=e.closePanel?e.closePanel:null,console.log("creating panel",i,e,n,r.context);var l=e.className?e.className+" panel ":" panel ";if(i.id&&(l=l+" "+i.id),i.name&&(l=l+" "+i.name),i.className&&(l=l+" "+i.className),r.overlay=e.overlay?e.overlay:null,i&&"string"==typeof i)r.processBlock(i,e);else if(i)switch(i.type){case"view":l+=" view ",r.processView(i,e,n);break;case"entity":l+=" entity ",r.processEntity(i,e,n);break;case"form":l+=" form ",r.processForm(i,e,n);break;case"html":l+=" html ",r.processHtml(i,e,n);break;case"block":l+=" panelblock ",r.processBlock(i,e,n);break;case"layout":l+=" layout ",r.processLayout(i,e,n);break;case"component":if(i.component)r.getView=function(e,t,n){return i.component};else{var p=r.getComponent(i.module,i.componentName,o.req);r.getView=function(e){return function(t,n,o,r){var a={description:i,className:r},u=i.props?s()({},i.props,a):a;return c.a.createElement(e,u)}}(p)}break;default:var m=_reg("PanelTypes",i.type);r.getView=function(e,t,n){return m.getComponent(i,e,t,n)}}return console.log("class name",r.className,l,i),r.className=l,r}return b()(t,e),d()(t,[{key:"getDisplayFunc",value:function(e,t){if(console.log("getting block",e),!e)return null;if("string"==typeof e)return _reg("Blocks",e);var n=e.id,o=_reg("Blocks",n);return o||(o=_reg("Blocks",e.defaultBlock)),o}},{key:"getChildContext",value:function(){return console.log("getting child contextoverlaying component",this.overlay,this.props,this.context),this.overlay?{overlayComponent:this.overlayComponent,closeOverlay:this.closeOverlay}:this.context&&this.context.overlayComponent?{overlayComponent:this.context.overlayComponent,closeOverlay:this.closeOverlay}:null}},{key:"render",value:function(){console.log("rendering panel",this.props,this.getView,this.className);var e=this.overlay&&this.state&&this.state.overlayComponent,t=this.getView?this.getView(this.props,this.context,this.state,this.title?"":this.className):c.a.createElement(this.context.uikit.Block,null);return this.overlay||this.title||this.closePanel?c.a.createElement(this.uikit.Block,{className:"overlaywrapper",title:this.title,closeBlock:this.closePanel},c.a.createElement(this.uikit.Block,{style:{display:e?"none":"block"}},t),e?this.state.overlayComponent:null):t}}],[{key:"setModule",value:function(e){o=e}}]),t}(c.a.Component)),w=function(){var e=this;this.getPanelItems=function(e){return e?e instanceof Array?e.map(function(e){return c.a.createElement(k,{description:e})}):c.a.createElement(k,{description:e}):null},this.cfgPanel=function(t,n){!e.title&&t&&(e.title=t),!e.overlay&&n&&(e.overlay=n)},this.processLayout=function(t,n,o){if(t&&t.layout){e.cfgPanel(t.title,t.overlay);var r=e,i=null,a=function(e){return t[e]?c.a.createElement(r.uikit.Block,{className:e},r.getPanelItems(t[e])):null};e.getView=function(e,n,o,r){switch(t.layout){case"2col":i=c.a.createElement(this.uikit.Block,{className:r+" twocol"},a("header"),c.a.createElement(this.uikit.Block,{className:"row"},a("left"),a("right")),a("footer"));break;case"3col":i=c.a.createElement(this.uikit.Block,{className:r+" threecol"},a("header"),c.a.createElement(this.uikit.Block,{className:"row"},a("left"),a("right")),a("footer"));break;default:i=c.a.createElement(this.uikit.Block,{className:r},a("items"))}return i}}},this.processBlock=function(t,n,o){var r=e.getDisplayFunc(t,n);e.cfgPanel(t.title,t.overlay),console.log("processing block",t,r,n);var i=e;e.getView=r?function(e,n,o,a){return console.log("calling block func",e,n,a),r({data:e.data,parent:e.parent,panel:i,className:a,routeParams:n.routeParams,storage:Storage},t,n.uikit)}:function(e,t,n,o){return c.a.createElement(t.uikit.Block,null)}},this.createMarkup=function(e){return{__html:e}},this.processHtml=function(t,n,o){e.cfgPanel(t.title,t.overlay),t.html?e.getView=function(e,n,o,r){return console.log("rendering html",t.html),c.a.createElement("div",{className:r,dangerouslySetInnerHTML:this.createMarkup(t.html)})}:e.getView=function(e,t,n,o){return c.a.createElement(t.uikit.Block,null)}},this.processForm=function(t,n,r){if(console.log("processing form=",t),t&&t.info){console.log("processing form+++",t,o),e.cfgPanel(t.info.title,t.info.overlay);var i=t.info;e.form||(console.log("getting form",o),e.form=e.getComponent("reactforms","Form",o.req),console.log("got form",o)),e.form?e.getView=function(e,n,o,r){var a=s()({},i,n.routeParams);return console.log("form cfg",a,i,e),c.a.createElement(this.form,{form:t.id,parentFormRef:e.parentFormRef,formContext:{data:e.data,routeParams:n.routeParams,storage:Storage},config:a,inline:e.inline,onChange:e.onChange,trackChanges:e.trackChanges,formData:e.formData,onSubmit:e.onSubmit,subform:e.subform,title:e.title,autoSubmitOnChange:e.autoSubmitOnChange,actions:e.actions,description:t,className:r,id:t.id})}:e.getView=function(e,t,n,o){return c.a.createElement(t.uikit.Block,null)}}},this.processView=function(t,n,r){console.log("processing my view",t,n,o),e.cfgPanel(t.title,t.overlay);var i=t.header?c.a.createElement(k,{description:t.header}):null;e.view||(e.view=e.getComponent("laatooviews","View",o.req)),console.log("processing view",e.view),e.getView=function(e,n,o,r){return console.log("rendering view",this.view,e,t,r),c.a.createElement(this.view,{params:e.params,description:t,getItem:e.getItem,editable:e.editable,className:r,header:i,viewRef:e.viewRef,id:t.id},c.a.createElement(k,{parent:e.parent,description:t.item}))}},this.processEntity=function(t,n,r){e.entity||(e.entity=e.getComponent("laatooviews","Entity",o.req)),e.getView=function(e,t,n,o){var r=e.description,i=r.entityDisplay?r.entityDisplay:"default";console.log("view entity description",r,i,e),this.cfgPanel(r.title,r.overlay);var a,u={type:"block",id:r.entityName+"_"+i,defaultBlock:r.entityName+"_default"},s="";s=t.routeParams&&t.routeParams.entityId?t.routeParams.entityId:r.entityId,a=r.entityName;var l=e.data?e.data:r.data,f=e.index,p="";return e.index&&(p=e.index%2?"oddindex":"evenindex"),console.log("my entity data111",l,f,r,e),c.a.createElement(this.entity,{id:s,name:a,entityDescription:r,data:l,index:f,uikit:t.uikit},c.a.createElement(k,{description:u,parent:e.parent,className:p}))}},this.overlayComponent=function(t){console.log("overlaying component"),e.overlay?e.setState(s()({},{overlayComponent:t})):e.context&&e.context.overlayComponent&&e.context.overlayComponent(t)},this.closeOverlay=function(){e.overlay?e.setState({}):e.context&&e.context.closeOverlay&&e.context.closeOverlay()},this.getComponent=function(e,t,n){var r=e+t,i=o[r];if(!i){var a=n(e);a&&t&&(i=a[t],o[r]=i)}return i}};k.childContextTypes={overlayComponent:x.a.func,closeOverlay:x.a.func},k.contextTypes={overlayComponent:x.a.func,closeOverlay:x.a.func,uikit:x.a.object,routeParams:x.a.object};var P=k,O=n(36);n.d(t,"Initialize",function(){return j}),n.d(t,"ProcessPages",function(){return E}),n.d(t,"Panel",function(){return P});var S;n(2);function j(e,t,n,o,r,i){(S=this).properties=Application.Properties[t],S.settings=o,S.req=i,Window.redirectPage||(Window.redirectPage=function(e,t){var n=_reg("Pages",e);if(console.log("redirect page",n),n){var o=formatUrl(n.url,t);Window.redirect(o)}}),P.setModule(S)}function E(e,t){var n=Application.AllRegItems("Pages");if(n)for(var o in n)try{!function(){var r=n[o],a=C(r),c=r.components;r.component&&(c={main:r.component}),e&&e.PreprocessPageComponents&&(c=e.PreprocessPageComponents(c,r,o,a,t));var u={};i()(c).forEach(function(n){u[n]=function(t,n,o,r,i){return function(a){var c=!0;if(e&&e.IsComponentVisible&&(c=e.IsComponentVisible(compToRender,n,o,a,r,i)),c){var u="function"==typeof t?t(a):t;if(e&&e.RenderPageComponent){var s=e.RenderPageComponent(u,n,o,a,r,i);if(s)return s}return u}return null}}(c[n],n,o,r,t)}),console.log("page comps ....",u);var s={pattern:r.route,components:u,reducer:Object(O.combineReducers)(a)};Application.Register("Routes",o,s),Application.Register("Actions","Page_"+o,{url:s.pattern})}()}catch(e){console.log(e)}}function C(e){var t={};for(var n in e.datasources)try{var o=_reg("Datasources",n),r={};o.type;var i=o.module;if(i){var a=S.req(i);a&&(r=a[o.processor])}r&&(t[n]=r)}catch(e){}return t}}])});
//# sourceMappingURL=index.js.map
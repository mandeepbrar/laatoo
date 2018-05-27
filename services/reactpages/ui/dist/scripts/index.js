define("reactpages",["react","prop-types","redux","reactwebcommon"],function(t,e,n,o){return function(t){var e={};function n(o){if(e[o])return e[o].exports;var r=e[o]={i:o,l:!1,exports:{}};return t[o].call(r.exports,r,r.exports,n),r.l=!0,r.exports}return n.m=t,n.c=e,n.d=function(t,e,o){n.o(t,e)||Object.defineProperty(t,e,{configurable:!1,enumerable:!0,get:o})},n.r=function(t){Object.defineProperty(t,"__esModule",{value:!0})},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=37)}([function(e,n){e.exports=t},function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(t,n){t.exports=e},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,e,n){var o=n(27)("wks"),r=n(26),i=n(9).Symbol;t.exports=function(t){return o[t]||(o[t]=i&&i[t]||(i||r)("Symbol."+t))}},function(t,e,n){var o=n(9),r=n(3),i=n(33),a=function(t,e,n){var c,s,u,l=t&a.F,f=t&a.G,p=t&a.S,y=t&a.P,d=t&a.B,m=t&a.W,v=f?r:r[e]||(r[e]={}),g=f?o:p?o[e]:(o[e]||{}).prototype;for(c in f&&(n=e),n)(s=!l&&g&&c in g)&&c in v||(u=s?g[c]:n[c],v[c]=f&&"function"!=typeof g[c]?n[c]:d&&s?i(u,o):m&&g[c]==u?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(u):y&&"function"==typeof u?i(Function.call,u):u,y&&((v.prototype||(v.prototype={}))[c]=u))};a.F=1,a.G=2,a.S=4,a.P=8,a.B=16,a.W=32,t.exports=a},function(t,e,n){t.exports={default:n(41),__esModule:!0}},function(t,e,n){var o=n(25),r=n(20);t.exports=function(t){return o(r(t))}},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e,n){"use strict";e.__esModule=!0;var o=a(n(50)),r=a(n(46)),i=a(n(32));function a(t){return t&&t.__esModule?t:{default:t}}e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,i.default)(e)));t.prototype=(0,r.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(o.default?(0,o.default)(t,e):t.__proto__=e)}},function(t,e,n){"use strict";e.__esModule=!0;var o,r=n(32),i=(o=r)&&o.__esModule?o:{default:o};e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,i.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){"use strict";e.__esModule=!0;var o,r=n(70),i=(o=r)&&o.__esModule?o:{default:o};e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var o=e[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),(0,i.default)(t,o.key,o)}}return function(e,n,o){return n&&t(e.prototype,n),o&&t(e,o),e}}()},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){t.exports={default:n(73),__esModule:!0}},function(t,e,n){var o=n(1).setDesc,r=n(17),i=n(4)("toStringTag");t.exports=function(t,e,n){t&&!r(t=n?t:t.prototype,i)&&o(t,i,{configurable:!0,value:e})}},function(t,e){t.exports={}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e,n){var o=n(1),r=n(18);t.exports=n(28)?function(t,e,n){return o.setDesc(t,e,r(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e,n){var o=n(20);t.exports=function(t){return Object(o(t))}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){var o=n(22);t.exports=function(t){if(!o(t))throw TypeError(t+" is not an object!");return t}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){var o=n(24);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==o(t)?t.split(""):Object(t)}},function(t,e){var n=0,o=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+o).toString(36))}},function(t,e,n){var o=n(9),r=o["__core-js_shared__"]||(o["__core-js_shared__"]={});t.exports=function(t){return r[t]||(r[t]={})}},function(t,e,n){t.exports=!n(8)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){t.exports=n(19)},function(t,e){t.exports=!0},function(t,e,n){"use strict";var o=n(30),r=n(5),i=n(29),a=n(19),c=n(17),s=n(16),u=n(63),l=n(15),f=n(1).getProto,p=n(4)("iterator"),y=!([].keys&&"next"in[].keys()),d=function(){return this};t.exports=function(t,e,n,m,v,g,h){u(n,e,m);var b,_,x=function(t){if(!y&&t in O)return O[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},k=e+" Iterator",w="values"==v,P=!1,O=t.prototype,S=O[p]||O["@@iterator"]||v&&O[v],E=S||x(v);if(S){var j=f(E.call(new t));l(j,k,!0),!o&&c(O,"@@iterator")&&a(j,p,d),w&&"values"!==S.name&&(P=!0,E=function(){return S.call(this)})}if(o&&!h||!y&&!P&&O[p]||a(O,p,E),s[e]=E,s[k]=d,v)if(b={values:w?E:x("values"),keys:g?E:x("keys"),entries:w?x("entries"):E},h)for(_ in b)_ in O||i(O,_,b[_]);else r(r.P+r.F*(y||P),e,b);return b}},function(t,e,n){"use strict";e.__esModule=!0;var o=a(n(68)),r=a(n(58)),i="function"==typeof r.default&&"symbol"==typeof o.default?function(t){return typeof t}:function(t){return t&&"function"==typeof r.default&&t.constructor===r.default&&t!==r.default.prototype?"symbol":typeof t};function a(t){return t&&t.__esModule?t:{default:t}}e.default="function"==typeof r.default&&"symbol"===i(o.default)?function(t){return void 0===t?"undefined":i(t)}:function(t){return t&&"function"==typeof r.default&&t.constructor===r.default&&t!==r.default.prototype?"symbol":void 0===t?"undefined":i(t)}},function(t,e,n){var o=n(71);t.exports=function(t,e,n){if(o(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,o){return t.call(e,n,o)};case 3:return function(n,o,r){return t.call(e,n,o,r)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){var o=n(5),r=n(3),i=n(8);t.exports=function(t,e){var n=(r.Object||{})[t]||Object[t],a={};a[t]=e(n),o(o.S+o.F*i(function(){n(1)}),"Object",a)}},function(t,e){t.exports=n},function(t,e,n){t.exports={default:n(44),__esModule:!0}},function(t,e,n){"use strict";n.r(e);var o,r=n(14),i=n.n(r),a=n(13),c=n.n(a),s=n(12),u=n.n(s),l=n(11),f=n.n(l),p=n(10),y=n.n(p),d=n(36),m=n.n(d),v=n(0),g=n.n(v),h=(n(42),n(6)),b=n.n(h),_=n(2),x=n.n(_),k=(n(38),function(t){function e(t,n){c()(this,e);var r=f()(this,(e.__proto__||i()(e)).call(this,t));w.call(r),r.uikit=n.uikit;var a=t.description,s=t.id?t.id:a&&a.id?a.id:null,u=t.type?t.type:a&&a.type?a.type:"layout";if(s){switch(u){case"view":a=_reg("Views",s);break;case"form":a=_reg("Forms",s);break;case"block":a=_reg("Blocks",s);break;default:a=_reg("Panels",s)}console.log("desc before assig",a,t),a=b()({type:u,id:s},a,t)}r.title=t.title?t.title:a&&a.title?a.title:null,r.closePanel=t.closePanel?t.closePanel:null,console.log("creating panel",a,t,n,r.context);var l=t.className?t.className+" panel ":" panel ";if(a.id&&(l=l+" "+a.id),a.name&&(l=l+" "+a.name),a.className&&(l=l+" "+a.className),r.overlay=t.overlay?t.overlay:null,a&&"string"==typeof a)r.processBlock(a,t);else if(a)switch(a.type){case"view":l+=" view ",r.processView(a,t,n);break;case"entity":l+=" entity ",r.processEntity(a,t,n);break;case"form":l+=" form ",r.processForm(a,t,n);break;case"html":l+=" html ",r.processHtml(a,t,n);break;case"block":l+=" panelblock ",r.processBlock(a,t,n);break;case"layout":l+=" layout ",r.processLayout(a,t,n);break;case"component":if(a.component)r.getView=function(t,e,n){return a.component};else{var p=r.getComponent(a.module,a.componentName,o.req);r.getView=function(t){return function(e,n,o,r){var i={description:a,className:r},c=a.props?b()({},a.props,i):i;return g.a.createElement(t,c)}}(p)}break;default:var y=_reg("PanelTypes",a.type);r.getView=function(t,e,n){return y.getComponent(a,t,e,n)}}return console.log("class name",r.className,l,a),r.className=l,r}return y()(e,t),u()(e,[{key:"getDisplayFunc",value:function(t,e){if(console.log("getting block",t),!t)return null;if("string"==typeof t)return _reg("Blocks",t);var n=t.id,o=_reg("Blocks",n);return o||(o=_reg("Blocks",t.defaultBlock)),o}},{key:"getChildContext",value:function(){return console.log("getting child contextoverlaying component",this.overlay,this.props,this.context),this.overlay?{overlayComponent:this.overlayComponent,closeOverlay:this.closeOverlay}:this.context&&this.context.overlayComponent?{overlayComponent:this.context.overlayComponent,closeOverlay:this.closeOverlay}:null}},{key:"render",value:function(){console.log("rendering panel",this.props,this.getView,this.className);var t=this.overlay&&this.state&&this.state.overlayComponent,e=this.getView?this.getView(this.props,this.context,this.state,this.title?"":this.className):g.a.createElement(this.context.uikit.Block,null);return this.overlay||this.title||this.closePanel?g.a.createElement(this.uikit.Block,{className:"overlaywrapper",title:this.title,closeBlock:this.closePanel},g.a.createElement(this.uikit.Block,{style:{display:t?"none":"block"}},e),t?this.state.overlayComponent:null):e}}],[{key:"setModule",value:function(t){o=t}}]),e}(g.a.Component)),w=function(){var t=this;this.getPanelItems=function(t){return t?t instanceof Array?t.map(function(t){return g.a.createElement(k,{description:t})}):g.a.createElement(k,{description:t}):null},this.cfgPanel=function(e,n){!t.title&&e&&(t.title=e),!t.overlay&&n&&(t.overlay=n)},this.processLayout=function(e,n,o){if(e&&e.layout){t.cfgPanel(e.title,e.overlay);var r=t,i=null,a=function(t){return e[t]?g.a.createElement(r.uikit.Block,{className:t},r.getPanelItems(e[t])):null};t.getView=function(t,n,o,r){switch(e.layout){case"2col":i=g.a.createElement(this.uikit.Block,{className:r+" twocol"},a("header"),g.a.createElement(this.uikit.Block,{className:"row"},a("left"),a("right")),a("footer"));break;case"3col":i=g.a.createElement(this.uikit.Block,{className:r+" threecol"},a("header"),g.a.createElement(this.uikit.Block,{className:"row"},a("left"),a("right")),a("footer"));break;default:i=g.a.createElement(this.uikit.Block,{className:r},a("items"))}return i}}},this.processBlock=function(e,n,o){var r=t.getDisplayFunc(e,n);t.cfgPanel(e.title,e.overlay),console.log("processing block",e,r,n);var i=t;t.getView=r?function(t,n,o,a){return console.log("calling block func",t,n),r({data:t.data,parent:t.parent,panel:i,className:a,routeParams:n.routeParams,storage:Storage},e,n.uikit)}:function(t,e,n,o){return g.a.createElement(e.uikit.Block,null)}},this.createMarkup=function(t){return{__html:t}},this.processHtml=function(e,n,o){t.cfgPanel(e.title,e.overlay),e.html?t.getView=function(t,n,o,r){return console.log("rendering html",e.html),g.a.createElement("div",{className:r,dangerouslySetInnerHTML:this.createMarkup(e.html)})}:t.getView=function(t,e,n,o){return g.a.createElement(e.uikit.Block,null)}},this.processForm=function(e,n,r){if(console.log("processing form=",e),e&&e.info){console.log("processing form+++",e,o),t.cfgPanel(e.info.title,e.info.overlay);var i=e.info;t.form||(console.log("getting form",o),t.form=t.getComponent("reactforms","Form",o.req),console.log("got form",o)),t.form?t.getView=function(t,n,o,r){var a=b()({},i,n.routeParams);return console.log("form cfg",a,i,t),g.a.createElement(this.form,{form:e.id,parentFormRef:t.parentFormRef,formContext:{data:t.data,routeParams:n.routeParams,storage:Storage},config:a,inline:t.inline,onChange:t.onChange,trackChanges:t.trackChanges,formData:t.formData,onSubmit:t.onSubmit,subform:t.subform,title:t.title,autoSubmitOnChange:t.autoSubmitOnChange,actions:t.actions,description:e,className:r,id:e.id})}:t.getView=function(t,e,n,o){return g.a.createElement(e.uikit.Block,null)}}},this.processView=function(e,n,r){console.log("processing my view",e,n,o),t.cfgPanel(e.title,e.overlay);var i=e.header?g.a.createElement(k,{description:e.header}):null;t.view||(t.view=t.getComponent("laatooviews","View",o.req)),console.log("processing view",t.view),t.getView=function(t,n,o,r){return console.log("rendering view",this.view,t,e,r),g.a.createElement(this.view,{params:t.params,description:e,getItem:t.getItem,editable:t.editable,className:r,header:i,viewRef:t.viewRef,id:e.id},g.a.createElement(k,{parent:t.parent,description:e.item}))}},this.processEntity=function(e,n,r){t.entity||(t.entity=t.getComponent("laatooviews","Entity",o.req)),t.getView=function(t,e,n,o){var r=t.description,i=r.entityDisplay?r.entityDisplay:"default";console.log("view entity description",r,i,t),this.cfgPanel(r.title,r.overlay);var a,c={type:"block",id:r.entityName+"_"+i,defaultBlock:r.entityName+"_default"},s="";s=e.routeParams&&e.routeParams.entityId?e.routeParams.entityId:r.entityId,a=r.entityName;var u=t.data?t.data:r.data,l=t.index,f="";return t.index&&(f=t.index%2?"oddindex":"evenindex"),console.log("my entity data111",u,l,r,t),g.a.createElement(this.entity,{id:s,name:a,entityDescription:r,data:u,index:l,uikit:e.uikit},g.a.createElement(k,{description:c,parent:t.parent,className:f}))}},this.overlayComponent=function(e){console.log("overlaying component"),t.overlay?t.setState(b()({},{overlayComponent:e})):t.context&&t.context.overlayComponent&&t.context.overlayComponent(e)},this.closeOverlay=function(){t.overlay?t.setState({}):t.context&&t.context.closeOverlay&&t.context.closeOverlay()},this.getComponent=function(t,e,n){var r=t+e,i=o[r];if(!i){var a=n(t);a&&e&&(i=a[e],o[r]=i)}return i}};k.childContextTypes={overlayComponent:x.a.func,closeOverlay:x.a.func},k.contextTypes={overlayComponent:x.a.func,closeOverlay:x.a.func,uikit:x.a.object,routeParams:x.a.object};var P=k,O=n(35);n.d(e,"Initialize",function(){return j}),n.d(e,"ProcessPages",function(){return C}),n.d(e,"Panel",function(){return P});var S,E=n(2);function j(t,e,n,o,r,i){(S=this).properties=Application.Properties[e],S.settings=o,S.req=i,Window.redirectPage||(Window.redirectPage=function(t,e){var n=_reg("Pages",t);if(console.log("redirect page",n),n){var o=formatUrl(n.url,e);Window.redirect(o)}}),P.setModule(S)}function C(t,e){var n=Application.AllRegItems("Pages");if(n)for(var o in n)try{!function(){var r=n[o],i=N(r),a=r.components;r.component&&(a={main:r.component});var c={};m()(a).forEach(function(t){var e,n;c[t]=(e=a[t],n=o,function(o){return g.a.createElement(M,{pageId:n,placeholder:t,routerState:o,description:e})})});var s={pattern:r.route,components:c,reducer:Object(O.combineReducers)(i)},u=s;t&&t.ProcessRoute&&(u=t.ProcessRoute(s,e)),Application.Register("Routes",o,u),Application.Register("Actions","Page_"+o,{url:u.pattern})}()}catch(t){console.log(t)}}function N(t){var e={};for(var n in t.datasources)try{var o=_reg("Datasources",n),r={};o.type;var i=o.module;if(i){var a=S.req(i);a&&(r=a[o.processor])}r&&(e[n]=r)}catch(t){}return e}var M=function(t){function e(){return c()(this,e),f()(this,(e.__proto__||i()(e)).apply(this,arguments))}return y()(e,t),u()(e,[{key:"getChildContext",value:function(){return{routeParams:this.props.routerState.params}}},{key:"render",value:function(){var t=this.props.pageId+this.props.placeholder;return g.a.createElement(P,{key:t,description:this.props.description})}}]),e}(g.a.Component);M.childContextTypes={routeParams:E.object}},function(t,e){t.exports=o},function(t,e,n){var o=n(1),r=n(21),i=n(25);t.exports=n(8)(function(){var t=Object.assign,e={},n={},o=Symbol(),r="abcdefghijklmnopqrst";return e[o]=7,r.split("").forEach(function(t){n[t]=t}),7!=t({},e)[o]||Object.keys(t({},n)).join("")!=r})?function(t,e){for(var n=r(t),a=arguments,c=a.length,s=1,u=o.getKeys,l=o.getSymbols,f=o.isEnum;c>s;)for(var p,y=i(a[s++]),d=l?u(y).concat(l(y)):u(y),m=d.length,v=0;m>v;)f.call(y,p=d[v++])&&(n[p]=y[p]);return n}:Object.assign},function(t,e,n){var o=n(5);o(o.S+o.F,"Object",{assign:n(39)})},function(t,e,n){n(40),t.exports=n(3).Object.assign},function(t,e){},function(t,e,n){var o=n(21);n(34)("keys",function(t){return function(e){return t(o(e))}})},function(t,e,n){n(43),t.exports=n(3).Object.keys},function(t,e,n){var o=n(1);t.exports=function(t,e){return o.create(t,e)}},function(t,e,n){t.exports={default:n(45),__esModule:!0}},function(t,e,n){var o=n(1).getDesc,r=n(22),i=n(23),a=function(t,e){if(i(t),!r(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,r){try{(r=n(33)(Function.call,o(Object.prototype,"__proto__").set,2))(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return a(t,n),e?t.__proto__=n:r(t,n),t}}({},!1):void 0),check:a}},function(t,e,n){var o=n(5);o(o.S,"Object",{setPrototypeOf:n(47).set})},function(t,e,n){n(48),t.exports=n(3).Object.setPrototypeOf},function(t,e,n){t.exports={default:n(49),__esModule:!0}},function(t,e){},function(t,e,n){var o=n(24);t.exports=Array.isArray||function(t){return"Array"==o(t)}},function(t,e,n){var o=n(1);t.exports=function(t){var e=o.getKeys(t),n=o.getSymbols;if(n)for(var r,i=n(t),a=o.isEnum,c=0;i.length>c;)a.call(t,r=i[c++])&&e.push(r);return e}},function(t,e,n){var o=n(7),r=n(1).getNames,i={}.toString,a="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[];t.exports.get=function(t){return a&&"[object Window]"==i.call(t)?function(t){try{return r(t)}catch(t){return a.slice()}}(t):r(o(t))}},function(t,e,n){var o=n(1),r=n(7);t.exports=function(t,e){for(var n,i=r(t),a=o.getKeys(i),c=a.length,s=0;c>s;)if(i[n=a[s++]]===e)return n}},function(t,e,n){"use strict";var o=n(1),r=n(9),i=n(17),a=n(28),c=n(5),s=n(29),u=n(8),l=n(27),f=n(15),p=n(26),y=n(4),d=n(55),m=n(54),v=n(53),g=n(52),h=n(23),b=n(7),_=n(18),x=o.getDesc,k=o.setDesc,w=o.create,P=m.get,O=r.Symbol,S=r.JSON,E=S&&S.stringify,j=!1,C=y("_hidden"),N=o.isEnum,M=l("symbol-registry"),B=l("symbols"),D="function"==typeof O,A=Object.prototype,F=a&&u(function(){return 7!=w(k({},"a",{get:function(){return k(this,"a",{value:7}).a}})).a})?function(t,e,n){var o=x(A,e);o&&delete A[e],k(t,e,n),o&&t!==A&&k(A,e,o)}:k,I=function(t){var e=B[t]=w(O.prototype);return e._k=t,a&&j&&F(A,t,{configurable:!0,set:function(e){i(this,C)&&i(this[C],t)&&(this[C][t]=!1),F(this,t,_(1,e))}}),e},V=function(t){return"symbol"==typeof t},T=function(t,e,n){return n&&i(B,e)?(n.enumerable?(i(t,C)&&t[C][e]&&(t[C][e]=!1),n=w(n,{enumerable:_(0,!1)})):(i(t,C)||k(t,C,_(1,{})),t[C][e]=!0),F(t,e,n)):k(t,e,n)},R=function(t,e){h(t);for(var n,o=v(e=b(e)),r=0,i=o.length;i>r;)T(t,n=o[r++],e[n]);return t},q=function(t,e){return void 0===e?w(t):R(w(t),e)},W=function(t){var e=N.call(this,t);return!(e||!i(this,t)||!i(B,t)||i(this,C)&&this[C][t])||e},L=function(t,e){var n=x(t=b(t),e);return!n||!i(B,e)||i(t,C)&&t[C][e]||(n.enumerable=!0),n},H=function(t){for(var e,n=P(b(t)),o=[],r=0;n.length>r;)i(B,e=n[r++])||e==C||o.push(e);return o},J=function(t){for(var e,n=P(b(t)),o=[],r=0;n.length>r;)i(B,e=n[r++])&&o.push(B[e]);return o},K=u(function(){var t=O();return"[null]"!=E([t])||"{}"!=E({a:t})||"{}"!=E(Object(t))});D||(s((O=function(){if(V(this))throw TypeError("Symbol is not a constructor");return I(p(arguments.length>0?arguments[0]:void 0))}).prototype,"toString",function(){return this._k}),V=function(t){return t instanceof O},o.create=q,o.isEnum=W,o.getDesc=L,o.setDesc=T,o.setDescs=R,o.getNames=m.get=H,o.getSymbols=J,a&&!n(30)&&s(A,"propertyIsEnumerable",W,!0));var G={for:function(t){return i(M,t+="")?M[t]:M[t]=O(t)},keyFor:function(t){return d(M,t)},useSetter:function(){j=!0},useSimple:function(){j=!1}};o.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=y(t);G[t]=D?e:I(e)}),j=!0,c(c.G+c.W,{Symbol:O}),c(c.S,"Symbol",G),c(c.S+c.F*!D,"Object",{create:q,defineProperty:T,defineProperties:R,getOwnPropertyDescriptor:L,getOwnPropertyNames:H,getOwnPropertySymbols:J}),S&&c(c.S+c.F*(!D||K),"JSON",{stringify:function(t){if(void 0!==t&&!V(t)){for(var e,n,o=[t],r=1,i=arguments;i.length>r;)o.push(i[r++]);return"function"==typeof(e=o[1])&&(n=e),!n&&g(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!V(e))return e}),o[1]=e,E.apply(S,o)}}}),f(O,"Symbol"),f(Math,"Math",!0),f(r.JSON,"JSON",!0)},function(t,e,n){n(56),n(51),t.exports=n(3).Symbol},function(t,e,n){t.exports={default:n(57),__esModule:!0}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e){t.exports=function(){}},function(t,e,n){"use strict";var o=n(60),r=n(59),i=n(16),a=n(7);t.exports=n(31)(Array,"Array",function(t,e){this._t=a(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,r(1)):r(0,"keys"==e?n:"values"==e?t[n]:[n,t[n]])},"values"),i.Arguments=i.Array,o("keys"),o("values"),o("entries")},function(t,e,n){n(61);var o=n(16);o.NodeList=o.HTMLCollection=o.Array},function(t,e,n){"use strict";var o=n(1),r=n(18),i=n(15),a={};n(19)(a,n(4)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=o.create(a,{next:r(1,n)}),i(t,e+" Iterator")}},function(t,e){var n=Math.ceil,o=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?o:n)(t)}},function(t,e,n){var o=n(64),r=n(20);t.exports=function(t){return function(e,n){var i,a,c=String(r(e)),s=o(n),u=c.length;return s<0||s>=u?t?"":void 0:(i=c.charCodeAt(s))<55296||i>56319||s+1===u||(a=c.charCodeAt(s+1))<56320||a>57343?t?c.charAt(s):i:t?c.slice(s,s+2):a-56320+(i-55296<<10)+65536}}},function(t,e,n){"use strict";var o=n(65)(!0);n(31)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=o(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){n(66),n(62),t.exports=n(4)("iterator")},function(t,e,n){t.exports={default:n(67),__esModule:!0}},function(t,e,n){var o=n(1);t.exports=function(t,e,n){return o.setDesc(t,e,n)}},function(t,e,n){t.exports={default:n(69),__esModule:!0}},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){var o=n(21);n(34)("getPrototypeOf",function(t){return function(e){return t(o(e))}})},function(t,e,n){n(72),t.exports=n(3).Object.getPrototypeOf}])});
//# sourceMappingURL=index.js.map
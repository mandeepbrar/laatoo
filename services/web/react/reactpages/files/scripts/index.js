define("reactpages",["react","prop-types","redux","reactwebcommon"],(function(e,t,o,n){return function(e){var t={};function o(n){if(t[n])return t[n].exports;var r=t[n]={i:n,l:!1,exports:{}};return e[n].call(r.exports,r,r.exports,o),r.l=!0,r.exports}return o.m=e,o.c=t,o.d=function(e,t,n){o.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},o.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},o.t=function(e,t){if(1&t&&(e=o(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(o.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var r in e)o.d(n,r,function(t){return e[t]}.bind(null,r));return n},o.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return o.d(t,"a",t),t},o.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},o.p="/",o(o.s=14)}([function(t,o){t.exports=e},function(e,t){e.exports=function(e){if(void 0===e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return e}},function(e,t){e.exports=function(e,t,o){return t in e?Object.defineProperty(e,t,{value:o,enumerable:!0,configurable:!0,writable:!0}):e[t]=o,e}},function(e,o){e.exports=t},function(e,t){e.exports=function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}},function(e,t){function o(e,t){for(var o=0;o<t.length;o++){var n=t[o];n.enumerable=n.enumerable||!1,n.configurable=!0,"value"in n&&(n.writable=!0),Object.defineProperty(e,n.key,n)}}e.exports=function(e,t,n){return t&&o(e.prototype,t),n&&o(e,n),e}},function(e,t,o){var n=o(11),r=o(1);e.exports=function(e,t){return!t||"object"!==n(t)&&"function"!=typeof t?r(e):t}},function(e,t){function o(t){return e.exports=o=Object.setPrototypeOf?Object.getPrototypeOf:function(e){return e.__proto__||Object.getPrototypeOf(e)},o(t)}e.exports=o},function(e,t,o){var n=o(12);e.exports=function(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Super expression must either be null or a function");e.prototype=Object.create(t&&t.prototype,{constructor:{value:e,writable:!0,configurable:!0}}),t&&n(e,t)}},function(e,t){e.exports=o},function(e,t){},function(e,t){function o(e){return(o="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e})(e)}function n(t){return"function"==typeof Symbol&&"symbol"===o(Symbol.iterator)?e.exports=n=function(e){return o(e)}:e.exports=n=function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":o(e)},n(t)}e.exports=n},function(e,t){function o(t,n){return e.exports=o=Object.setPrototypeOf||function(e,t){return e.__proto__=t,e},o(t,n)}e.exports=o},function(e,t){e.exports=n},function(e,t,o){"use strict";o.r(t);var n,r=o(0),a=o.n(r),i=(o(10),o(4)),l=o.n(i),c=o(5),s=o.n(c),u=o(6),f=o.n(u),p=o(7),m=o.n(p),y=o(1),g=o.n(y),d=o(8),v=o.n(d),b=o(2),h=o.n(b),k=o(3),w=o.n(k),P=(o(13),function(e){function t(e,o){var r;l()(this,t),r=f()(this,m()(t).call(this,e)),h()(g()(r),"getPanelItems",(function(e){return e?e instanceof Array?e.map((function(e){return console.log("rendering panel array element",e),a.a.createElement(t,{description:e})})):a.a.createElement(t,{description:e}):null})),h()(g()(r),"cfgPanel",(function(e,t){!r.title&&e&&(r.title=e),!r.overlay&&t&&(r.overlay=t)})),h()(g()(r),"processLayout",(function(e,t,o){if(e&&e.layout){r.cfgPanel(e.title,e.overlay);var n=g()(r),i=null,l=function(t){return console.log("looking for panel component in layout",e,t),e[t]?a.a.createElement(_uikit.Block,{className:t},n.getPanelItems(e[t])):a.a.createElement(_uikit.Block,null)};r.getView=function(t,o,n,r){switch(console.log("getting layout",e,t),e.layout){case"2col":i=a.a.createElement(_uikit.Block,{className:r+" fd fdcol w100 twocol"},l("header"),a.a.createElement(_uikit.Block,{className:" fd fdgrow fdrow "},l("left"),l("right")),l("footer"));break;case"3col":i=a.a.createElement(_uikit.Block,{className:r+" fd fdcol w100 threecol"},l("header"),a.a.createElement(_uikit.Block,{className:"row"},l("left"),l("right")),l("footer"));break;default:i=a.a.createElement(_uikit.Block,{className:r+" fd fdcol w100 "},l("items"))}return i}}})),h()(g()(r),"processBlock",(function(e,t,o){var n=r.getDisplayFunc(e,t);r.cfgPanel(e.title,e.overlay),console.log("processing block",e,n,t);var i=g()(r);r.getView=n?function(t,o,r,a){console.log("calling block func",t,o,a);var l=n({data:t.data,parent:t.parent,panel:i,className:a,routeParams:o.routeParams,storage:Storage},e);return console.log("returning block retval",l),l}:function(e,t,o,n){return console.log("rendering empty block func",e,t,n),a.a.createElement(_uikit.Block,null)}})),h()(g()(r),"createMarkup",(function(e){return{__html:e}})),h()(g()(r),"processHtml",(function(e,t,o){r.cfgPanel(e.title,e.overlay),e.html?r.getView=function(t,o,n,r){return console.log("rendering html",e.html),a.a.createElement("div",{className:r,dangerouslySetInnerHTML:this.createMarkup(e.html)})}:r.getView=function(e,t,o,n){return a.a.createElement(_uikit.Block,null)}})),h()(g()(r),"processForm",(function(e,t,o){if(console.log("processing form=",e,t),e&&e.info){r.cfgPanel(e.info.title,e.info.overlay);var i=e.formName?e.formName:e.id;console.log("processing form+++",e,i,t,e.formName,e.id);var l=e.info;r.form||(console.log("getting form",n),r.form=_res("reactforms","Form"),console.log("got form",n)),r.form?r.getView=function(t,o,n,r){var c=l;return t.subform||(c=Object.assign({},l,o.routeParams)),console.log("form cfg",c,l,t,i),a.a.createElement(this.form,{form:i,parentFormRef:t.parentFormRef,formContext:{data:t.data,routeParams:o.routeParams,storage:Storage},info:c,inline:t.inline,onChange:t.onChange,trackChanges:t.trackChanges,formData:t.formData,subform:t.subform,onFormSubmit:t.onFormSubmit,title:t.title,actions:t.actions,description:e,className:r,id:e.id})}:r.getView=function(e,t,o,n){return a.a.createElement(_uikit.Block,null)}}})),h()(g()(r),"processView",(function(e,o,i){console.log("processing my view",e,o,n),r.cfgPanel(e.title,e.overlay);var l=e.header?a.a.createElement(t,{id:e.header,type:"block"}):null;r.view||(r.view=_res("laatooviews","View")),console.log("processing view",r.view),r.getView=function(o,n,r,i){return console.log("rendering view",this.view,o,e,i),a.a.createElement(this.view,{params:o.params,description:e,getItem:o.getItem,editable:o.editable,className:i,header:l,viewRef:o.viewRef,postArgs:o.postArgs,urlParams:o.urlParams,id:e.id},a.a.createElement(t,{parent:o.parent,description:e.item}))}})),h()(g()(r),"processEntity",(function(e,o,n){r.entity||(r.entity=_res("laatooviews","Entity")),r.getView=function(e,o,n,r){var i=e.description,l=i.entityDisplay?i.entityDisplay:"default";console.log("view entity description",i,l,e,o),this.cfgPanel(i.title,i.overlay);var c,s={type:"block",id:i.entityName+"_"+l,defaultBlock:i.entityName+"_default"},u="";u=o.routeParams&&o.routeParams.entityId?o.routeParams.entityId:i.entityId,c=i.entityName;var f=e.data?e.data:i.data,p=e.index,m="";return e.index&&(m=e.index%2?"oddindex":"evenindex"),console.log("my entity data",f,p,i,e,u,o),a.a.createElement(this.entity,{id:u,name:c,entityDescription:i,data:f,index:p},a.a.createElement(t,{description:s,parent:e.parent,className:m}))}})),h()(g()(r),"overlayComponent",(function(e){console.log("overlaying component"),r.overlay?r.setState(Object.assign({},{overlayComponent:e})):r.context&&r.context.overlayComponent&&r.context.overlayComponent(e)})),h()(g()(r),"closeOverlay",(function(){r.overlay?r.setState({}):r.context&&r.context.closeOverlay&&r.context.closeOverlay()}));var i=e.description,c=e.id?e.id:i&&i.id?i.id:null,s=e.type?e.type:i&&i.type?i.type:"layout";if(c){var u=i?i.info:null;switch(s){case"view":i=_reg("Views",c);break;case"form":i=_reg("Forms",c);break;case"block":i=_reg("Blocks",c);break;case"component":break;default:i=_reg("Panels",c)}console.log("desc before assig",i,e,c,s,Application),i=Object.assign({type:s,id:c},i),console.log("desc after assign",i,i.info,u),i.info=Object.assign({},i.info,u)}r.title=e.title?e.title:i&&i.title?i.title:null,r.closePanel=e.closePanel?e.closePanel:null,console.log("creating panel",i,e,o,r.context);var p,y=e.className?e.className+" panel ":" panel ";if(i.id&&(y=y+" "+i.id),i.name&&(y=y+" "+i.name),i.className&&(y=y+" "+i.className),r.overlay=e.overlay?e.overlay:null,i&&"string"==typeof i)r.processBlock(i,e);else if(i)switch(i.type){case"view":y+=" view ",r.processView(i,e,o);break;case"entity":y+=" entity ",r.processEntity(i,e,o);break;case"form":y+=" form ",r.processForm(i,e,o);break;case"html":y+=" html ",r.processHtml(i,e,o);break;case"block":y+=" panelblock ",r.processBlock(i,e,o);break;case"layout":y+=" layout ",r.processLayout(i,e,o);break;case"component":if(i.component)r.getView=function(e,t,o){return i.component};else{var d=_res(i.module,i.componentName);console.log("component from module",d,i),r.getView=(p=d,function(e,t,o,n){var r={description:i,className:n},l=i.props?Object.assign({},i.props,r):Object.assign({},i,r);return a.a.createElement(p,l)})}break;default:var v=_reg("PanelTypes",i.type);r.getView=function(e,t,o){return v.getComponent(i,e,t,o)}}return console.log("class name",r.className,y,i),r.className=y,r}return v()(t,e),s()(t,[{key:"getDisplayFunc",value:function(e,t){if(console.log("getting display func",e),!e)return console.log("returning null display func",e),null;if("string"==typeof e)return _reg("Blocks",e);var o=e.id,n=_reg("Blocks",o);return n||(n=_reg("Blocks",e.defaultBlock)),console.log("returning display func",n,Application),n}},{key:"getChildContext",value:function(){return console.log("getting child contextoverlaying component",this.overlay,this.props,this.context),this.overlay?{overlayComponent:this.overlayComponent,closeOverlay:this.closeOverlay}:this.context&&this.context.overlayComponent?{overlayComponent:this.context.overlayComponent,closeOverlay:this.closeOverlay}:null}},{key:"render",value:function(){var e=this.overlay&&this.state&&this.state.overlayComponent,t=this.getView?this.getView(this.props,this.context,this.state,this.title?"":this.className):a.a.createElement(_uikit.Block,null);return console.log("Rendering panel***************",this.getView,this.props,t,this.overlay,this.title,this.closePanel,this.className),this.overlay||this.title||this.closePanel?a.a.createElement(_uikit.Block,{className:"overlaywrapper",title:this.title,closeBlock:this.closePanel},a.a.createElement(_uikit.Block,{style:{display:e?"none":"block"}},t),e?this.state.overlayComponent:null):t}}],[{key:"setModule",value:function(e){n=e}}]),t}(a.a.Component));P.childContextTypes={overlayComponent:w.a.func,closeOverlay:w.a.func},P.contextTypes={overlayComponent:w.a.func,closeOverlay:w.a.func,routeParams:w.a.object};var x=P,_=o(9);o.d(t,"Initialize",(function(){return E})),o.d(t,"ProcessPages",(function(){return C})),o.d(t,"Panel",(function(){return x}));var O;o(3);function E(e,t,o,n,r,a){(O=this).properties=Application.Properties[t],O.settings=n,O.req=a,Window.redirectPage||(Window.redirectPage=function(e,t){var o=_reg("Pages",e);if(console.log("redirect page",o),o){var n=formatUrl(o.url,t);Window.redirect(n)}}),x.setModule(O)}function C(e){var t=Application.AllRegItems("Pages");if(t)for(var o in t)try{!function(){var n=t[o],r=N(n),a=n.components;n.component&&(a={main:n.component}),e&&e.PreprocessPageComponents&&(a=e.PreprocessPageComponents(a,n,o,r));var i={};console.log("page components ",o,n,a),Object.keys(a).forEach((function(t){i[t]=function(t,o,n,r){return function(a){var i=!0;if(console.log("Page components ",a,t,o,n,r),e&&e.IsComponentVisible&&(i=e.IsComponentVisible(compToRender,o,n,a,r)),i){var l="function"==typeof t?t(a):t;if(e&&e.RenderPageComponent){var c=e.RenderPageComponent(l,o,n,a,r);if(c)return c}return l}return null}}(a[t],t,o,n)}));var l={pattern:n.route,components:i,reducer:Object(_.combineReducers)(r)};console.log("page ....",l),Application.Register("Routes",o,l),Application.Register("Actions","Page_"+o,{url:l.pattern})}()}catch(e){console.log(e)}}function N(e){var t={};for(var o in e.datasources)try{var n=_reg("Datasources",o),r={};n.type;var a=n.module;if(a){var i=O.req(a);i&&(r=i[n.processor])}r&&(t[o]=r)}catch(e){}return t}Window.resolvePanel=function(e,t){return a.a.createElement(x,{type:e,id:t})}}])}));
//# sourceMappingURL=index.js.map
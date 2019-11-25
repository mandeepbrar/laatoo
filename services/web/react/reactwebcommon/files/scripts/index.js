define("reactwebcommon",["react","uicommon","sanitize-html","react-redux"],(function(t,e,n,o){return function(t){var e={};function n(o){if(e[o])return e[o].exports;var r=e[o]={i:o,l:!1,exports:{}};return t[o].call(r.exports,r,r.exports,n),r.l=!0,r.exports}return n.m=t,n.c=e,n.d=function(t,e,o){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:o})},n.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var o=Object.create(null);if(n.r(o),Object.defineProperty(o,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var r in t)n.d(o,r,function(e){return t[e]}.bind(null,r));return o},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=20)}([function(e,n){e.exports=t},function(t,e,n){t.exports=n(15)()},function(t,e){t.exports=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e){function n(t,e){for(var n=0;n<e.length;n++){var o=e[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),Object.defineProperty(t,o.key,o)}}t.exports=function(t,e,o){return e&&n(t.prototype,e),o&&n(t,o),t}},function(t,e,n){var o=n(13),r=n(7);t.exports=function(t,e){return!e||"object"!==o(e)&&"function"!=typeof e?r(t):e}},function(t,e){function n(e){return t.exports=n=Object.setPrototypeOf?Object.getPrototypeOf:function(t){return t.__proto__||Object.getPrototypeOf(t)},n(e)}t.exports=n},function(t,e,n){var o=n(14);t.exports=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function");t.prototype=Object.create(e&&e.prototype,{constructor:{value:t,writable:!0,configurable:!0}}),e&&o(t,e)}},function(t,e){t.exports=function(t){if(void 0===t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return t}},function(t,n){t.exports=e},function(t,e){t.exports=function(t,e,n){return e in t?Object.defineProperty(t,e,{value:n,enumerable:!0,configurable:!0,writable:!0}):t[e]=n,t}},function(t,e){function n(){return t.exports=n=Object.assign||function(t){for(var e=1;e<arguments.length;e++){var n=arguments[e];for(var o in n)Object.prototype.hasOwnProperty.call(n,o)&&(t[o]=n[o])}return t},n.apply(this,arguments)}t.exports=n},function(t,e){t.exports=n},function(t,e){t.exports=o},function(t,e){function n(t){return(n="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t})(t)}function o(e){return"function"==typeof Symbol&&"symbol"===n(Symbol.iterator)?t.exports=o=function(t){return n(t)}:t.exports=o=function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":n(t)},o(e)}t.exports=o},function(t,e){function n(e,o){return t.exports=n=Object.setPrototypeOf||function(t,e){return t.__proto__=e,t},n(e,o)}t.exports=n},function(t,e,n){"use strict";var o=n(16);function r(){}function i(){}i.resetWarningCache=r,t.exports=function(){function t(t,e,n,r,i,s){if(s!==o){var a=new Error("Calling PropTypes validators directly is not supported by the `prop-types` package. Use PropTypes.checkPropTypes() to call them. Read more at http://fb.me/use-check-prop-types");throw a.name="Invariant Violation",a}}function e(){return t}t.isRequired=t;var n={array:t,bool:t,func:t,number:t,object:t,string:t,symbol:t,any:t,arrayOf:e,element:t,elementType:t,instanceOf:e,node:t,objectOf:e,oneOf:e,oneOfType:e,shape:e,exact:e,checkPropTypes:i,resetWarningCache:r};return n.PropTypes=n,n}},function(t,e,n){"use strict";t.exports="SECRET_DO_NOT_PASS_THIS_OR_YOU_WILL_BE_FIRED"},function(t,e){},function(t,e){},function(t,e){},function(t,e,n){"use strict";n.r(e);var o=n(10),r=n.n(o),i=n(2),s=n.n(i),a=n(3),c=n.n(a),l=n(4),u=n.n(l),p=n(5),h=n.n(p),f=n(6),d=n.n(f),m=n(0),y=n.n(m),b=function(t){function e(t){var n;return s()(this,e),n=u()(this,h()(e).call(this,t)),t.skipPrefix||(t.skipPrefix=!1),n}return d()(e,t),c()(e,[{key:"render",value:function(){var t=this.props.src;if(!t||0==t.length)return this.props.children?this.props.children:(console.log("No src tag provieded for image",this.props),null);!this.props.prefix||this.props.skipPrefix||this.props.src.startsWith("http")||(t=this.props.prefix+t);var e=y.a.createElement("img",r()({src:t},this.props.modifier,{className:this.props.className,style:this.props.style}));return this.props.link?y.a.createElement("a",{target:this.props.target,href:this.props.link},e):e}}]),e}(y.a.Component),v=n(11),g=n.n(v);var w=function(t){return y.a.createElement("div",{className:t.className,style:t.style,dangerouslySetInnerHTML:(e=t.children,n=t.sanitize,o=e,n&&(o=g()(e)),{__html:o})});var e,n,o},k=n(7),S=n.n(k),_=function(t){function e(t){var n;if(s()(this,e),(n=u()(this,h()(e).call(this,t))).handleScroll=n.handleScroll.bind(S()(n)),t.windowScroll){var o=Math.floor(window.scrollY/window.innerHeight);n.state={windowNumber:o}}return n.state={scrolledOut:!1,scrolledIn:!1},n}return d()(e,t),c()(e,[{key:"handleScroll",value:function(t){if(this.props.windowScroll){var e=Math.floor(window.scrollY/window.innerHeight);e!=this.state.windowNumber&&(this.props.windowScroll(e),this.setState({windowNumber:e}))}if(this.props.onScrollEnd||this.props.onScrollIn){var n=this.refs.scrollListener,o=window.innerHeight;if(this.props.scrollEndPos&&(o=this.props.scrollEndPos),null!=n){var r=n.getBoundingClientRect();r.bottom<=o&&!this.state.scrolledOut&&(!this.state.scrolledOut&&this.props.onScrollEnd&&this.props.onScrollEnd(r.bottom),this.setState({scrolledOut:!0,scrolledIn:!1})),this.state.scrolledOut&&r.bottom>o&&(!this.state.scrolledIn&&this.props.onScrollIn&&this.props.onScrollIn(r.bottom),this.setState({scrolledOut:!1,scrolledIn:!0}))}}}},{key:"componentDidMount",value:function(){window.addEventListener("scroll",this.handleScroll)}},{key:"componentWillUnmount",value:function(){window.removeEventListener("scroll",this.handleScroll)}},{key:"render",value:function(){return y.a.createElement("div",{ref:"scrollListener",key:this.props.key,style:this.props.style,className:this.props.className},this.props.children)}}]),e}(y.a.Component),O=n(1),E=n.n(O),x=function(t,e){return _uikit.ActionButton?y.a.createElement(_uikit.ActionButton,{flat:t.flat,className:t.className+" actionbutton",onClick:t.actionFunc,btnProps:t},t.actionchildren):y.a.createElement("a",{className:t.className+" actionbutton",onClick:t.actionFunc,role:"button"},t.actionchildren)};x.propTypes={actionFunc:E.a.func.isRequired,actionchildren:E.a.oneOfType([E.a.array,E.a.string])};var N=x,P=n(12),C=function(t){return y.a.createElement("a",{className:t.className+" actionlink",href:"javascript:void(0)",onClick:t.actionFunc},t.actionchildren)};C.propTypes={actionFunc:E.a.func.isRequired,actionchildren:E.a.oneOfType([E.a.array,E.a.string])};var j=C,M=n(8),T=function(t){function e(t){var n;return s()(this,e),n=u()(this,h()(e).call(this,t)),console.log("action comp creation",t),n.renderView=n.renderView.bind(S()(n)),n.dispatchAction=n.dispatchAction.bind(S()(n)),n.actionFunc=n.actionFunc.bind(S()(n)),n.hasPermission=!1,null!=t.action?n.action=t.action:n.action=_reg("Actions",t.name),console.log("action",n.action),n.action&&"method"==n.action.actiontype&&(n.props.method?n.actionMethod=n.props.method?n.props.method:_reg("Methods",n.action.method):n.action.method&&"function"==typeof n.action.method?n.actionMethod=n.action.method:n.actionMethod=_reg("Methods",n.action.method)),n.action&&(n.hasPermission=Object(M.hasPermission)(n.action.permission)),n}return d()(e,t),c()(e,[{key:"dispatchAction",value:function(){var t={};this.props.params&&(t=this.props.params),this.props.dispatch(Object(M.createAction)(this.action.action,t,{successCallback:this.props.successCallback,failureCallback:this.props.failureCallback}))}},{key:"actionFunc",value:function(t){if(console.log("action executed",this.props.name,this.props,this.action),t.preventDefault(),this.props.confirm&&!this.props.confirm(this.props))return!1;switch(this.action.actiontype){case"dispatchaction":return this.dispatchAction(),!1;case"method":var e=this.props.params?this.props.params:this.action.params;return this.actionMethod(e),!1;case"interaction":if(!this.action.interactiontype)return!1;var n=Window.resolvePanel("block",this.action.blockid),o=this.props.onClose?this.props.onClose:_reg("Methods",this.action.onClose);return Window.showInteraction(this.action.interactiontype,this.action.title,n,o,this.action.actions,this.action.contentStyle,this.action.titleStyle),!1;case"newwindow":if(this.action.url){var r=Object(M.formatUrl)(this.action.url,this.props.params);return console.log(r),window.open(r),!1}default:if(this.action.url){var i=Object(M.formatUrl)(this.action.url,this.props.params);console.log(i),Window.redirect(i,this.action.newpage)}return!1}}},{key:"renderView",value:function(){if(!this.hasPermission)return null;var t=this.props.children?this.props.children:this.props.label,e=this.actionFunc;switch(_tn(this.props.widget,this.action.widget)){case"button":return y.a.createElement(N,{flat:_tn(this.props.flat,this.action.flat),className:this.props.className,actionFunc:e,key:this.props.name+"_comp",actionchildren:t});case"component":return y.a.createElement(this.props.component,{actionFunc:e,key:this.props.name+"_comp",actionchildren:t});default:return y.a.createElement(j,{className:this.props.className,actionFunc:e,key:this.props.name+"_comp",actionchildren:t})}}},{key:"render",value:function(){return this.renderView()}}]),e}(y.a.Component);T.propTypes={name:E.a.string.isRequired};var A=Object(P.connect)()(T),B=function(t){var e=t.className?t.className:"",n=t.contentClass?t.contentClass:"";return t.title||t.titleBarActions||t.closeBlock?y.a.createElement("div",{style:t.style,className:"block "+e},y.a.createElement("div",{className:"titlebar"},y.a.createElement("div",{className:"title left"},t.title),t.titleBarActions?y.a.createElement("div",{className:"right"},t.titleBarActions):t.closeBlock?y.a.createElement(_uikit.Icon,{className:"right close fa fa-close",onClick:t.closeBlock}):null),y.a.createElement("div",{style:t.contentStyle,className:"blockcontent "+n},t.children)):y.a.createElement("div",{style:Object.assign({},t.contentStyle,t.style),onClick:t.onClick,className:"block "+n+" "+e},t.children)},I=function(t){function e(t){var n;return s()(this,e),n=u()(this,h()(e).call(this,t)),console.log("menu props",t),n.menu=_reg("Menus",t.id),n.menu&&n.menu.items?(console.log("iterating menu items ",n.menu.items),n.menu.items.forEach((function(t){t.page&&(t.action="Page_"+t.page)})),n.menuitems=n.menu.items):n.menuitems=[],n}return d()(e,t),c()(e,[{key:"render",value:function(){return y.a.createElement(_uikit.Navbar,{items:this.menuitems,vertical:this.props.vertical})}}]),e}(y.a.Component),L=function(t){function e(t){var n;if(s()(this,e),(n=u()(this,h()(e).call(this,t))).actions=[],console.log("creating action bar",t),t.children)n.className=" actionbar "+_tn(t.className,""),n.actions=t.children;else{var o=t.description;if(t.id||o&&o.id){var r=t.id?t.id:o.id;o=_reg("ActionBar",r)}if(o){console.log("action bar",t,o),n.description=o,n.className=" actionbar "+_tn(t.className,_tn(o.className,""));var i=S()(n);o.actions&&o.actions.forEach((function(t){i.actions.push(y.a.createElement(A,{name:t.name,label:t.label,widget:t.widget,className:" action "}))}))}}return n}return d()(e,t),c()(e,[{key:"render",value:function(){return y.a.createElement(_uikit.Block,{className:this.className},this.actions)}}]),e}(y.a.Component),F=n(9),R=n.n(F),D=function(t){function e(t){var n;return s()(this,e),n=u()(this,h()(e).call(this,t)),R()(S()(n),"errorMethod",(function(t){console.log("could not load data",t)})),R()(S()(n),"response",(function(t){console.log("loadable component:---------response",t),t&&t.data&&n.dataLoaded(t.data)})),t.loader&&(n.method=_reg("Methods",t.loader)),t.loadData&&(n.loadData=t.loadData),n}return d()(e,t),c()(e,[{key:"componentWillMount",value:function(){console.log("loadable component:---------",this.method,this.props);var t=this.props;if(this.loadData)if(this.method)this.method(t,this.getLoadContext?this.getLoadContext():{},this.dataLoaded);else if(t.dataService){var e=M.RequestBuilder.DefaultRequest(null,t.dataServiceParams);M.DataSource.ExecuteService(t.dataService,e).then(this.response,this.errorMethod)}else t.entity&&M.EntityData.ListEntities(t.entity).then(this.response,this.errorMethod)}}]),e}(y.a.Component);n(17),n(18),n(19);n.d(e,"ScrollListener",(function(){return _})),n.d(e,"Action",(function(){return A})),n.d(e,"Block",(function(){return B})),n.d(e,"Html",(function(){return w})),n.d(e,"Menu",(function(){return I})),n.d(e,"ActionBar",(function(){return L})),n.d(e,"LoadableComponent",(function(){return D})),n.d(e,"Image",(function(){return b}))}])}));
//# sourceMappingURL=index.js.map
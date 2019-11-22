define("laatooviews",["uicommon","react","react-redux","prop-types"],(function(e,t,r,n){return function(e){var t={};function r(n){if(t[n])return t[n].exports;var a=t[n]={i:n,l:!1,exports:{}};return e[n].call(a.exports,a,a.exports,r),a.l=!0,a.exports}return r.m=e,r.c=t,r.d=function(e,t,n){r.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},r.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},r.t=function(e,t){if(1&t&&(e=r(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(r.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var a in e)r.d(n,a,function(t){return e[t]}.bind(null,a));return n},r.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return r.d(t,"a",t),t},r.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},r.p="/",r(r.s=18)}([function(t,r){t.exports=e},function(e,t){e.exports=function(e){if(void 0===e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return e}},function(e,r){e.exports=t},function(e,t){e.exports=function(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}},function(e,t,r){e.exports=r(16)},function(e,t){function r(t){return e.exports=r=Object.setPrototypeOf?Object.getPrototypeOf:function(e){return e.__proto__||Object.getPrototypeOf(e)},r(t)}e.exports=r},function(e,t){e.exports=function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}},function(e,t){function r(e,t){for(var r=0;r<t.length;r++){var n=t[r];n.enumerable=n.enumerable||!1,n.configurable=!0,"value"in n&&(n.writable=!0),Object.defineProperty(e,n.key,n)}}e.exports=function(e,t,n){return t&&r(e.prototype,t),n&&r(e,n),e}},function(e,t,r){var n=r(12),a=r(1);e.exports=function(e,t){return!t||"object"!==n(t)&&"function"!=typeof t?a(e):t}},function(e,t,r){var n=r(14);e.exports=function(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Super expression must either be null or a function");e.prototype=Object.create(t&&t.prototype,{constructor:{value:e,writable:!0,configurable:!0}}),t&&n(e,t)}},function(e,t){e.exports=r},function(e,t,r){var n=r(13);function a(t,r,o){return"undefined"!=typeof Reflect&&Reflect.get?e.exports=a=Reflect.get:e.exports=a=function(e,t,r){var a=n(e,t);if(a){var o=Object.getOwnPropertyDescriptor(a,t);return o.get?o.get.call(r):o.value}},a(t,r,o||t)}e.exports=a},function(e,t){function r(e){return(r="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e})(e)}function n(t){return"function"==typeof Symbol&&"symbol"===r(Symbol.iterator)?e.exports=n=function(e){return r(e)}:e.exports=n=function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":r(e)},n(t)}e.exports=n},function(e,t,r){var n=r(5);e.exports=function(e,t){for(;!Object.prototype.hasOwnProperty.call(e,t)&&null!==(e=n(e)););return e}},function(e,t){function r(t,n){return e.exports=r=Object.setPrototypeOf||function(e,t){return e.__proto__=t,e},r(t,n)}e.exports=r},function(e,t){e.exports=n},function(e,t,r){var n=function(e){"use strict";var t,r=Object.prototype,n=r.hasOwnProperty,a="function"==typeof Symbol?Symbol:{},o=a.iterator||"@@iterator",i=a.asyncIterator||"@@asyncIterator",s=a.toStringTag||"@@toStringTag";function c(e,t,r,n){var a=t&&t.prototype instanceof m?t:m,o=Object.create(a.prototype),i=new x(n||[]);return o._invoke=function(e,t,r){var n=u;return function(a,o){if(n===d)throw new Error("Generator is already running");if(n===f){if("throw"===a)throw o;return L()}for(r.method=a,r.arg=o;;){var i=r.delegate;if(i){var s=T(i,r);if(s){if(s===h)continue;return s}}if("next"===r.method)r.sent=r._sent=r.arg;else if("throw"===r.method){if(n===u)throw n=f,r.arg;r.dispatchException(r.arg)}else"return"===r.method&&r.abrupt("return",r.arg);n=d;var c=l(e,t,r);if("normal"===c.type){if(n=r.done?f:p,c.arg===h)continue;return{value:c.arg,done:r.done}}"throw"===c.type&&(n=f,r.method="throw",r.arg=c.arg)}}}(e,r,i),o}function l(e,t,r){try{return{type:"normal",arg:e.call(t,r)}}catch(e){return{type:"throw",arg:e}}}e.wrap=c;var u="suspendedStart",p="suspendedYield",d="executing",f="completed",h={};function m(){}function g(){}function y(){}var v={};v[o]=function(){return this};var E=Object.getPrototypeOf,b=E&&E(E(C([])));b&&b!==r&&n.call(b,o)&&(v=b);var w=y.prototype=m.prototype=Object.create(v);function I(e){["next","throw","return"].forEach((function(t){e[t]=function(e){return this._invoke(t,e)}}))}function _(e){var t;this._invoke=function(r,a){function o(){return new Promise((function(t,o){!function t(r,a,o,i){var s=l(e[r],e,a);if("throw"!==s.type){var c=s.arg,u=c.value;return u&&"object"==typeof u&&n.call(u,"__await")?Promise.resolve(u.__await).then((function(e){t("next",e,o,i)}),(function(e){t("throw",e,o,i)})):Promise.resolve(u).then((function(e){c.value=e,o(c)}),(function(e){return t("throw",e,o,i)}))}i(s.arg)}(r,a,t,o)}))}return t=t?t.then(o,o):o()}}function T(e,r){var n=e.iterator[r.method];if(n===t){if(r.delegate=null,"throw"===r.method){if(e.iterator.return&&(r.method="return",r.arg=t,T(e,r),"throw"===r.method))return h;r.method="throw",r.arg=new TypeError("The iterator does not provide a 'throw' method")}return h}var a=l(n,e.iterator,r.arg);if("throw"===a.type)return r.method="throw",r.arg=a.arg,r.delegate=null,h;var o=a.arg;return o?o.done?(r[e.resultName]=o.value,r.next=e.nextLoc,"return"!==r.method&&(r.method="next",r.arg=t),r.delegate=null,h):o:(r.method="throw",r.arg=new TypeError("iterator result is not an object"),r.delegate=null,h)}function S(e){var t={tryLoc:e[0]};1 in e&&(t.catchLoc=e[1]),2 in e&&(t.finallyLoc=e[2],t.afterLoc=e[3]),this.tryEntries.push(t)}function O(e){var t=e.completion||{};t.type="normal",delete t.arg,e.completion=t}function x(e){this.tryEntries=[{tryLoc:"root"}],e.forEach(S,this),this.reset(!0)}function C(e){if(e){var r=e[o];if(r)return r.call(e);if("function"==typeof e.next)return e;if(!isNaN(e.length)){var a=-1,i=function r(){for(;++a<e.length;)if(n.call(e,a))return r.value=e[a],r.done=!1,r;return r.value=t,r.done=!0,r};return i.next=i}}return{next:L}}function L(){return{value:t,done:!0}}return g.prototype=w.constructor=y,y.constructor=g,y[s]=g.displayName="GeneratorFunction",e.isGeneratorFunction=function(e){var t="function"==typeof e&&e.constructor;return!!t&&(t===g||"GeneratorFunction"===(t.displayName||t.name))},e.mark=function(e){return Object.setPrototypeOf?Object.setPrototypeOf(e,y):(e.__proto__=y,s in e||(e[s]="GeneratorFunction")),e.prototype=Object.create(w),e},e.awrap=function(e){return{__await:e}},I(_.prototype),_.prototype[i]=function(){return this},e.AsyncIterator=_,e.async=function(t,r,n,a){var o=new _(c(t,r,n,a));return e.isGeneratorFunction(r)?o:o.next().then((function(e){return e.done?e.value:o.next()}))},I(w),w[s]="Generator",w[o]=function(){return this},w.toString=function(){return"[object Generator]"},e.keys=function(e){var t=[];for(var r in e)t.push(r);return t.reverse(),function r(){for(;t.length;){var n=t.pop();if(n in e)return r.value=n,r.done=!1,r}return r.done=!0,r}},e.values=C,x.prototype={constructor:x,reset:function(e){if(this.prev=0,this.next=0,this.sent=this._sent=t,this.done=!1,this.delegate=null,this.method="next",this.arg=t,this.tryEntries.forEach(O),!e)for(var r in this)"t"===r.charAt(0)&&n.call(this,r)&&!isNaN(+r.slice(1))&&(this[r]=t)},stop:function(){this.done=!0;var e=this.tryEntries[0].completion;if("throw"===e.type)throw e.arg;return this.rval},dispatchException:function(e){if(this.done)throw e;var r=this;function a(n,a){return s.type="throw",s.arg=e,r.next=n,a&&(r.method="next",r.arg=t),!!a}for(var o=this.tryEntries.length-1;o>=0;--o){var i=this.tryEntries[o],s=i.completion;if("root"===i.tryLoc)return a("end");if(i.tryLoc<=this.prev){var c=n.call(i,"catchLoc"),l=n.call(i,"finallyLoc");if(c&&l){if(this.prev<i.catchLoc)return a(i.catchLoc,!0);if(this.prev<i.finallyLoc)return a(i.finallyLoc)}else if(c){if(this.prev<i.catchLoc)return a(i.catchLoc,!0)}else{if(!l)throw new Error("try statement without catch or finally");if(this.prev<i.finallyLoc)return a(i.finallyLoc)}}}},abrupt:function(e,t){for(var r=this.tryEntries.length-1;r>=0;--r){var a=this.tryEntries[r];if(a.tryLoc<=this.prev&&n.call(a,"finallyLoc")&&this.prev<a.finallyLoc){var o=a;break}}o&&("break"===e||"continue"===e)&&o.tryLoc<=t&&t<=o.finallyLoc&&(o=null);var i=o?o.completion:{};return i.type=e,i.arg=t,o?(this.method="next",this.next=o.finallyLoc,h):this.complete(i)},complete:function(e,t){if("throw"===e.type)throw e.arg;return"break"===e.type||"continue"===e.type?this.next=e.arg:"return"===e.type?(this.rval=this.arg=e.arg,this.method="return",this.next="end"):"normal"===e.type&&t&&(this.next=t),h},finish:function(e){for(var t=this.tryEntries.length-1;t>=0;--t){var r=this.tryEntries[t];if(r.finallyLoc===e)return this.complete(r.completion,r.afterLoc),O(r),h}},catch:function(e){for(var t=this.tryEntries.length-1;t>=0;--t){var r=this.tryEntries[t];if(r.tryLoc===e){var n=r.completion;if("throw"===n.type){var a=n.arg;O(r)}return a}}throw new Error("illegal catch attempt")},delegateYield:function(e,r,n){return this.delegate={iterator:C(e),resultName:r,nextLoc:n},"next"===this.method&&(this.arg=t),h}},e}(e.exports);try{regeneratorRuntime=n}catch(e){Function("r","regeneratorRuntime = r")(n)}},function(e,t){},function(e,t,r){"use strict";r.r(t);var n=r(3),a=r.n(n),o=r(10),i={CONTAINER_REFRESH:"CONTAINER_REFRESH",VIEW_FETCH:"VIEW_FETCH",VIEW_FETCHING:"VIEW_FETCHING",VIEW_FETCH_SUCCESS:"VIEW_FETCH_SUCCESS",VIEW_FETCH_FAILED:"VIEW_FETCH_FAILED",VIEW_ITEM_REMOVE:"VIEW_ITEM_REMOVE",VIEW_ITEM_RELOAD:"VIEW_ITEM_RELOAD",ENTITY_VIEW_FETCH:"ENTITY_VIEW_FETCH",ENTITY_VIEW_FETCHING:"ENTITY_VIEW_FETCHING",ENTITY_VIEW_FETCH_SUCCESS:"ENTITY_VIEW_FETCH_SUCCESS",ENTITY_VIEW_FETCH_FAILED:"ENTITY_VIEW_FETCH_FAILED",PAGE_CHANGE:"@@reduxdirector/LOCATION_CHANGE",GROUP_LOAD:"GROUP_LOAD"},s=r(0),c=r(6),l=r.n(c),u=r(7),p=r.n(u),d=r(8),f=r.n(d),h=r(1),m=r.n(h),g=r(5),y=r.n(g),v=r(11),E=r.n(v),b=r(9),w=r.n(b),I=r(2),_=r.n(I),T=function e(){l()(this,e),this.index=-1,this.data=null,this.renderedItem=null},S=function(e){function t(e){var r;return l()(this,t),r=f()(this,y()(t).call(this,e)),a()(m()(r),"itemSelectionChange",(function(e,t){var n=r.viewitems[e];if(n){var a=n.renderedItem;r.setSelection(a,t)}r.props.singleSelection&&(r.selectedItem&&r.setSelection(r.selectedItem.renderedItem,!1),r.selectedItem=n)})),r.setPage=r.setPage.bind(m()(r)),r.selectedItems=r.selectedItems.bind(m()(r)),r.itemCount=r.itemCount.bind(m()(r)),r.reload=r.reload.bind(m()(r)),r.setFilter=r.setFilter.bind(m()(r)),r.loadMore=r.loadMore.bind(m()(r)),r.canLoadMore=r.canLoadMore.bind(m()(r)),r.methods={reload:r.reload,canLoadMore:r.canLoadMore,loadMore:r.loadMore,setFilter:r.setFilter,itemCount:r.itemCount,itemSelectionChange:r.itemSelectionChange,selectedItems:r.selectedItems,setPage:r.setPage},r.addMethod=r.addMethod.bind(m()(r)),r.state={lastLoadTime:-1},r.pushItem=r.pushItem.bind(m()(r)),r.viewitems=new Array,console.log("this items",r.viewitems),r}return w()(t,e),p()(t,[{key:"componentWillMount",value:function(){this.filter=this.props.defaultFilter}},{key:"componentDidMount",value:function(){this.props.load&&!this.props.externalLoad&&this.props.loadView(this.props.currentPage,this.filter)}},{key:"componentWillReceiveProps",value:function(e){e.load&&e.loadView(e.currentPage,this.filter)}},{key:"shouldComponentUpdate",value:function(e,t){if(!e.forceUpdate&&this.lastRenderTime){if(!e.lastUpdateTime)return!1;if(this.lastRenderTime>=e.lastUpdateTime)return!1}return!0}},{key:"addMethod",value:function(e,t){this.methods[e]=t}},{key:"reload",value:function(){this.props.loadView(this.props.currentPage,this.filter)}},{key:"canLoadMore",value:function(){return this.props.currentPage<this.props.totalPages}},{key:"pushItem",value:function(e){console.log("pushing item",this.viewitems,this),this.viewitems.push(e)}},{key:"itemCount",value:function(){return this.viewitems.length}},{key:"setSelection",value:function(e,t){e.setSelected?e.setSelected(t):e.selected=t}},{key:"selectedItems",value:function(){if(this.props.singleSelection)return this.selectedItem.data;var e=[],t=this.itemCount();console.log("num items ",t);for(var r=0;r<t;r++){var n=this.viewitems[r];console.log("vitem",n,r);var a=n.renderedItem;(a.getSelected&&a.getSelected()||a.selected)&&e.push(n.data)}return e}},{key:"setPage",value:function(e){this.props.loadView(e,this.filter)}},{key:"setFilter",value:function(e){this.filter=e,this.props.loadView(1,this.filter)}},{key:"loadMore",value:function(){return!(this.props.currentPage>=this.props.totalPages)&&(this.props.currentPage?(this.props.loadIncrementally(this.props.currentPage+1,this.filter),!0):void 0)}},{key:"render",value:function(){this.lastRenderTime=this.props.lastUpdateTime;var e=this.renderView(this.props.items,this.props.currentPage,this.props.totalPages);return console.log("rendering view data",e,this.props.items),this.items=this.props.items,e}}]),t}(_.a.Component),O=(r(15),function(e){function t(e,r){var n;return l()(this,t),n=f()(this,y()(t).call(this,e)),a()(m()(n),"getRenderedItem",(function(e,t){return console.log("get rendered item",n.props.children,n.props),_.a.Children.map(n.props.children,(function(r){return _.a.cloneElement(r,{data:e,index:t})}))})),n.getItemGroup=n.getItemGroup.bind(m()(n)),n.getView=n.getView.bind(m()(n)),n.renderView=n.renderView.bind(m()(n)),n.getItem=n.getItem.bind(m()(n)),n.getHeader=n.getHeader.bind(m()(n)),n.getPagination=n.getPagination.bind(m()(n)),n.getFilter=n.getFilter.bind(m()(n)),n.onScrollEnd=n.onScrollEnd.bind(m()(n)),n}return w()(t,e),p()(t,[{key:"onScrollEnd",value:function(){this.methods().loadMore()}},{key:"getView",value:function(e,t,r,n){return console.log("view ui getView",this.props),this.props.contentOnly?t:this.props.getView?this.props.getView(this,e,t,r,n,this.props):this.props.incrementalLoad?_.a.createElement(_uikit.scroll,{key:this.props.key,className:this.props.className,onScrollEnd:this.onScrollEnd},n,e,t,r):_.a.createElement(_uikit.Block,{key:this.props.key,className:this.props.className,style:this.props.style},n,e,t,r)}},{key:"getFilter",value:function(){return this.props.getFilter?this.props.getFilter(this,this.props.defaultFilter):null}},{key:"getItemGroup",value:function(e){return this.props.getItemGroup?this.props.getItemGroup(this,e):_.a.createElement(_uikit.Block,{className:"group"},"x")}},{key:"getItem",value:function(e,r){var n=null;n=this.props.getItem?this.props.getItem(this,e,r):this.getRenderedItem(e,r);var a=new T;return a.index=r,a.data=e,a.renderedItem=n,console.log("pushing item",a),E()(y()(t.prototype),"pushItem",this).call(this,a),n}},{key:"getHeader",value:function(){return this.props.getHeader?this.props.getHeader(this):this.props.header?this.props.header:null}},{key:"getPagination",value:function(){if(this.props.paginate&&this.props.getPagination){var e=this.props.totalPages,t=this.props.currentPage;return this.props.getPagination(this,e,t)}return null}},{key:"renderView",value:function(e,t,r){var n=[];if(e){var a=Object.keys(e);for(var o in a){var i=e[a[o]];if(i){var s=this.getItem(i,a[o]);n.push(s)}}}else this.props.loader&&n.push(this.props.loader);var c=this.getHeader(),l=this.getFilter(),u=this.getPagination();return this.getView(c,n,u,l)}}]),t}(S));function x(e){return e.serviceName?e.serviceName:e.name}function C(e,t,r,n,a){var o=x(t);console.log("service.....",o,t),r||(r=1);var c={};t.paginate&&(c.pagesize=t.pageSize,c.pagenum=r);var l=t.dataurl?{url:t.dataurl,method:"POST"}:null,u={queryParams:c,postArgs:Object.assign({},t.postArgs,n)},p={serviceName:o,global:t.global,viewname:t.name,serviceObject:l,incrementalLoad:a};e(Object(s.createAction)(i.VIEW_FETCH,u,p))}var L=Object(o.connect)((function(e,t){var r,n=x(t),o=t.name,i=(t.reducer&&t.reducer,r={name:o,global:t.global,paginate:t.paginate,pageSize:t.pageSize,header:t.header,getView:t.getView,getItem:t.getItem,getHeader:t.getHeader,defaultFilter:t.defaultFilter,singleSelection:t.singleSelection,externalLoad:t.externalLoad,urlParams:t.urlParams,postArgs:t.postArgs},a()(r,"getView",t.getView),a()(r,"incrementalLoad",t.incrementalLoad),a()(r,"currentPage",t.currentPage),a()(r,"className",t.className),a()(r,"ref",t.viewRef),a()(r,"loader",t.loader),a()(r,"getPagination",t.incrementalLoad||t.hidePaginationControl?null:t.getPagination),a()(r,"style",t.style),a()(r,"totalPages",1),a()(r,"load",!1),a()(r,"items",null),r),s=e.views;if(s&&o){var c=s.views[o];c?("Loaded"==c.status&&(i.items=c.data,i.currentPage=c.currentPage,i.totalPages=c.totalPages,i.lastUpdateTime=c.lastUpdateTime,i.latestPageData=c.latestPageData),"NotLoaded"==c.status&&(i.load=!0)):i.load=!0}return i}),(function(e,t){return console.log("view comp properties",t),{loadView:function(r,n){C(e,t,r,n,!1)},loadIncrementally:function(r,n){C(e,t,r,n,!0)}}}))(O),k=function(e){console.log("props of the view...",e);var t=e.description;if(!t&&e.id&&(t=_reg("Views",e.id)),t){console.log("view.....",t);var r=e.postArgs?e.postArgs:t.postArgs,n=e.urlparams?e.urlparams:t.urlparams,a=e.instance?e.instance:t.name?t.name:e.id,o=e.className+" view_"+a+" ",i=t.singleSelection?t.singleSelection:e.singleSelection;t.className&&(o+=t.className);var s=e.children;return _.a.createElement(L,{serviceObject:t.service,serviceName:t.serviceName,name:a,global:t.global,singleSelection:i,className:o,incrementalLoad:t.incrementalLoad,paginate:t.paginate,header:e.header,getHeader:e.getHeader,editable:e.editable,contentOnly:e.contentOnly,getView:e.getView,getItem:e.getItem,urlparams:n,postArgs:r,viewRef:e.viewRef},s)}return null},P=function(e){function t(){return l()(this,t),f()(this,y()(t).apply(this,arguments))}return w()(t,e),p()(t,[{key:"componentDidMount",value:function(){this.props.loadGroup()}},{key:"render",value:function(){return null}}]),t}(_.a.Component),A=Object(o.connect)((function(e,t){return{}}),(function(e,t){return{loadGroup:function(){console.log("load group",t);var r=t.Data,n={serviceName:t.service};e(Object(s.createAction)(i.GROUP_LOAD,r,n))}}}))(P),j={views:{}},F={status:"NotLoaded",data:{},global:!1,currentPage:1,totalPages:1,pagesize:-1};Application.Register("Reducers","views",(function(e,t){if(!t||!t.meta||!t.meta.viewname){if(!e)return j;if(t.type==i.LOGOUT)return j;if(t.type==i.PAGE_CHANGE){var r=[];return Object.keys(e.views).forEach((function(t){var n=e.views[t];n.global&&r.push(n)})),Object.assign({},e,{views:r})}return e}if(t.type&&t.meta&&t.meta.viewname){var n=t.meta.viewname,a=e.views[n],o={};switch(t.type){case i.VIEW_FETCHING:var s=1,c=-1;t.payload.queryParams&&t.payload.queryParams.pagenum&&(s=t.payload.queryParams.pagenum,c=t.payload.queryParams.pagesize);var l={status:"Fetching",currentPage:s,pagesize:c};o=a?Object.assign({},a,l):Object.assign({},F,l);break;case i.VIEW_FETCH_SUCCESS:var u=1;if(t.meta.info&&t.meta.info.totalrecords){var p=t.meta.info.totalrecords;p>0&&a.pagesize>0&&(u=Math.ceil(p/a.pagesize))}var d=null,f=a.data;f&&t.meta.incrementalLoad?t.payload&&(d=Array.isArray(t.payload)?f.concat(t.payload):Object.assign(f,t.payload)):d=t.payload,o=Object.assign({},a,{status:"Loaded",data:d,lastUpdateTime:(new Date).getTime(),totalPages:u});break;case i.VIEW_FETCH_FAILED:o=Object.assign({},F,{status:"LoadingFailed"});break;case i.VIEW_ITEM_RELOAD:var h=t.meta.Index;if(null==h)return e;var m=a.data,g=null;if(Array.isArray(m)){var y=-1;y="string"==typeof h?parseInt(h):h,(g=m.slice(0))[y]=t.payload}else(g=Object.assign({},m))[h]=t.payload;o=Object.assign({},a,{data:g,lastUpdateTime:(new Date).getTime()});break;case i.VIEW_ITEM_REMOVE:var v=t.payload.Index;if(null==v)return e;var E=a.data,b=null;if(Array.isArray(E)){var w=-1;w="string"==typeof v?parseInt(v):v,(b=E.slice(0)).splice(w,1)}else delete(b=Object.assign({},E))[v],null==b&&(b={});o=Object.assign({},a,{data:b,lastUpdateTime:(new Date).getTime()})}var I={};I[n]=o;var _=Object.assign({},e.views,I);return Object.assign({},e,{views:_})}}));var N="entityview",V={entities:{}},H={status:"NotLoaded",data:{},global:!1,entityId:"",entityName:""};Application.Register("Reducers",N,(function(e,t){if(!t||!t.meta||t.meta.reducer!=N){if(!e)return V;if(t.type==i.LOGOUT)return V;if(t.type==i.PAGE_CHANGE){var r=[];return Object.keys(e.entities).forEach((function(t){var n=e.entities[t];n.global&&r.push(n)})),Object.assign({},e,{entities:r})}return e}if(t.type&&t.meta&&t.meta.reducer==N){var n=t.meta.entityId,a=e.entities[n],o={};switch(t.type){case i.ENTITY_VIEW_FETCHING:var s={status:"Fetching",global:t.meta.global,entityId:n,entityName:t.payload.entityName};o[n]=a?Object.assign({},a,s):Object.assign({},H,s);break;case i.ENTITY_VIEW_FETCH_SUCCESS:o[n]=Object.assign({},a,{status:"Loaded",data:t.payload.data,lastUpdateTime:(new Date).getTime()});break;case i.ENTITY_VIEW_FETCH_FAILED:o[n]=Object.assign({},H,{status:"LoadingFailed"})}var c=Object.assign({},e.entities,o);return Object.assign({},e,{entities:c})}}));var W=r(4),R=r.n(W),G=function(e){return"@@redux-saga/"+e},U=G("IO"),D=G("MULTICAST");var M=function(e){return null==e},Y=function(e){return null!=e},q=function(e){return"function"==typeof e},B=function(e){return"string"==typeof e},z=Array.isArray,K=function e(t){return t&&(B(t)||X(t)||q(t)||z(t)&&t.every(e))},J=function(e){return e&&q(e.take)&&q(e.close)},Q=function(e){return q(e)&&e.hasOwnProperty("toString")},X=function(e){return Boolean(e)&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype},Z=function(e){return J(e)&&e[D]};"function"==typeof Symbol&&Symbol.asyncIterator&&Symbol.asyncIterator;var $=function(e){throw e},ee=function(e){return{value:e,done:!0}};var te="TAKE",re="PUT",ne="CALL",ae="FORK",oe=function(e,t){var r;return(r={})[U]=!0,r.combinator=!1,r.type=e,r.payload=t,r};function ie(e,t){return void 0===e&&(e="*"),K(e)?oe(te,{pattern:e}):Z(e)&&Y(t)&&K(t)?oe(te,{channel:e,pattern:t}):J(e)?oe(te,{channel:e}):void 0}function se(e,t){return M(t)&&(t=e,e=void 0),oe(re,{channel:e,action:t})}function ce(e,t){var r,n=null;return q(e)?r=e:(z(e)?(n=e[0],r=e[1]):(n=e.context,r=e.fn),n&&B(r)&&q(n[r])&&(r=n[r])),{context:n,fn:r,args:t}}function le(e){for(var t=arguments.length,r=new Array(t>1?t-1:0),n=1;n<t;n++)r[n-1]=arguments[n];return oe(ne,ce(e,r))}function ue(e){for(var t=arguments.length,r=new Array(t>1?t-1:0),n=1;n<t;n++)r[n-1]=arguments[n];return oe(ae,ce(e,r))}var pe=function(e){return{done:!0,value:e}},de={};function fe(e){return J(e)?"channel":Q(e)?String(e):q(e)?e.name:String(e)}function he(e,t,r){var n,a,o,i=t;function s(t,r){if(i===de)return pe(t);if(r&&!a)throw i=de,r;n&&n(t);var s=r?e[a](r):e[i]();return i=s.nextState,o=s.effect,n=s.stateUpdater,a=s.errorState,i===de?pe(t):o}return function(e,t,r){void 0===t&&(t=$),void 0===r&&(r="iterator");var n={meta:{name:r},next:e,throw:t,return:ee,isSagaIterator:!0};return"undefined"!=typeof Symbol&&(n[Symbol.iterator]=function(){return n}),n}(s,(function(e){return s(null,e)}),r)}function me(e,t){for(var r=arguments.length,n=new Array(r>2?r-2:0),a=2;a<r;a++)n[a-2]=arguments[a];var o,i={done:!1,value:ie(e)},s=function(e){return o=e};return he({q1:function(){return{nextState:"q2",effect:i,stateUpdater:s}},q2:function(){return{nextState:"q1",effect:(e=o,{done:!1,value:ue.apply(void 0,[t].concat(n,[e]))})};var e}},"q1","takeEvery("+fe(e)+", "+t.name+")")}function ge(e,t){for(var r=arguments.length,n=new Array(r>2?r-2:0),a=2;a<r;a++)n[a-2]=arguments[a];return ue.apply(void 0,[me,e,t].concat(n))}var ye=R.a.mark(Ee),ve=R.a.mark(be);function Ee(e){var t,r;return R.a.wrap((function(n){for(;;)switch(n.prev=n.next){case 0:return n.prev=0,n.next=3,se(Object(s.createAction)(i.VIEW_FETCHING,e.payload,{global:e.meta.global,viewname:e.meta.viewname}));case 3:if(t=s.RequestBuilder.DefaultRequest(e.payload.queryParams,e.payload.postArgs,e.payload.headers),r=null,!e.meta.serviceObject){n.next=11;break}return n.next=8,le(s.DataSource.ExecuteServiceObject,e.meta.serviceObject,t);case 8:r=n.sent,n.next=14;break;case 11:return n.next=13,le(s.DataSource.ExecuteService,e.meta.serviceName,t);case 13:r=n.sent;case 14:return n.next=16,se(Object(s.createAction)(i.VIEW_FETCH_SUCCESS,r.data,{info:r.info,incrementalLoad:e.meta.incrementalLoad,global:e.meta.global,viewname:e.meta.viewname}));case 16:n.next=23;break;case 18:return n.prev=18,n.t0=n.catch(0),n.next=22,se(Object(s.createAction)(i.VIEW_FETCH_FAILED,n.t0,e.meta));case 22:e.meta.failureCallback?e.meta.failureCallback(n.t0):Window.handleError&&Window.handleError(n.t0);case 23:case"end":return n.stop()}}),ye,null,[[0,18]])}function be(){return R.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return console.log("views saga start"),e.next=3,ge(i.VIEW_FETCH,Ee);case 3:case"end":return e.stop()}}),ve)}Application.Register("Sagas","views",be);var we=R.a.mark(_e),Ie=R.a.mark(Te);function _e(e){var t,r;return R.a.wrap((function(n){for(;;)switch(n.prev=n.next){case 0:return t={reducer:"entityview",global:e.meta.global,entityId:e.payload.entityId},n.prev=1,n.next=4,se(Object(s.createAction)(i.ENTITY_VIEW_FETCHING,e.payload,t));case 4:return n.next=6,le(s.EntityData.GetEntity,e.payload.entityName,e.payload.entityId,e.payload.headers,e.payload.svc);case 6:return(r=n.sent).data.isOwner=r.data.CreatedBy===Storage.user,n.next=10,se(Object(s.createAction)(i.ENTITY_VIEW_FETCH_SUCCESS,r,t));case 10:e.meta.successCallback&&e.meta.successCallback({resp:r,payload:e.payload}),n.next=18;break;case 13:return n.prev=13,n.t0=n.catch(1),n.next=17,se(Object(s.createAction)(i.ENTITY_VIEW_FETCH_FAILED,n.t0,t));case 17:e.meta.failureCallback?e.meta.failureCallback(n.t0):Window.handleError&&Window.handleError(n.t0);case 18:case"end":return n.stop()}}),we,null,[[1,13]])}function Te(){return R.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return console.log("entity view saga start"),e.next=3,ge(i.ENTITY_VIEW_FETCH,_e);case 3:case"end":return e.stop()}}),Ie)}Application.Register("Sagas","entityViewSaga",Te);var Se=R.a.mark(xe),Oe=R.a.mark(Ce);function xe(e){var t,r,n,a,o,c,l;return R.a.wrap((function(u){for(;;)switch(u.prev=u.next){case 0:return u.prev=0,t={},console.log("load group saga",e.payload),Object.keys(e.payload).forEach((function(r){var n=e.payload[r];console.log("load group ",n,"key",r),"entity"==n.type&&(t[r]={Params:{id:n.entityId},Body:{}}),"view"==n.type&&(t[r]={Params:n.queryParams,Body:n.postArgs})})),console.log("created request",t),r=s.RequestBuilder.DefaultRequest(null,t),u.next=8,le(s.DataSource.ExecuteService,e.meta.serviceName,r);case 8:n=u.sent,console.log("resp",n),a=new Array,Object.keys(e.payload).forEach((function(t){var r=e.payload[t],o=n.data[t];console.log("service ",r,o);var c={data:o.Data,statuscode:o.Status,info:o.Info};"entity"==r.type&&(200==o.Status?(c.data.isOwner=c.data.CreatedBy===Storage.user,a.push(Object(s.createAction)(i.ENTITY_GET_SUCCESS,c,{reducer:r.meta.reducer}))):a.push(Object(s.createAction)(i.ENTITY_GET_FAILED,c,{reducer:r.meta.reducer}))),"view"==r.type&&(200==o.Status?a.push(Object(s.createAction)(i.VIEW_FETCH_SUCCESS,c.data,{info:c.info,incrementalLoad:e.meta.incrementalLoad,reducer:r.meta.reducer})):a.push(Object(s.createAction)(i.VIEW_FETCH_FAILED,c.data,{reducer:r.reducer})))})),e.meta.successCallback&&e.meta.successCallback({resp:n,payload:e.payload}),console.log("actions i",a),o=0;case 15:if(!(o<a.length)){u.next=22;break}return console.log("actions i",o,a[o]),u.next=19,se(a[o]);case 19:o++,u.next=15;break;case 22:u.next=38;break;case 24:if(u.prev=24,u.t0=u.catch(0),!e.meta.services){u.next=37;break}c=new Array,Object.keys(e.meta.services).forEach((function(t){var r=e.meta.services[t];"entity"==r.type&&c.push(Object(s.createAction)(i.ENTITY_GET_FAILED,u.t0,{reducer:r.reducer})),"view"==r.type&&c.push(Object(s.createAction)(i.VIEW_FETCH_FAILED,u.t0,{reducer:r.reducer}))})),l=0;case 30:if(!(l<c.length)){u.next=37;break}return console.log("actions i",l,c[l]),u.next=34,se(c[l]);case 34:l++,u.next=30;break;case 37:e.meta.failureCallback?e.meta.failureCallback(u.t0):Window.handleError&&Window.handleError(u.t0);case 38:case"end":return u.stop()}}),Se,null,[[0,24]])}function Ce(){return R.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return console.log("group load saga start"),e.next=3,ge(i.GROUP_LOAD,xe);case 3:case"end":return e.stop()}}),Oe)}Application.Register("Sagas","groupLoadSaga",Ce);r(17);var Le=function(e){function t(e){return l()(this,t),f()(this,y()(t).call(this,e))}return w()(t,e),p()(t,[{key:"componentDidMount",value:function(){this.props.load&&!this.props.externalLoad&&this.props.loadEntity()}},{key:"componentWillReceiveProps",value:function(e){e.load&&this.props.loadEntity()}},{key:"shouldComponentUpdate",value:function(e,t){if(!e.forceUpdate&&this.lastRenderTime){if(!e.lastUpdateTime)return!1;if(this.lastRenderTime>=e.lastUpdateTime)return!1}return!0}},{key:"render",value:function(){console.log("laatoo views ",this.props);return this.lastRenderTime=this.props.lastUpdateTime,this.props.display&&this.props.status&&"Loading"==this.props.status?this.props.loader:this.props.children?_.a.cloneElement(_.a.Children.only(this.props.children),{data:this.props.data}):this.props.data?this.props.display(this.props.data,this.props.desc,this.props.lastUpdateTime):_.a.createElement(_uikit.Block,null)}}]),t}(_.a.Component),ke=Object(o.connect)((function(e,t){var r={name:t.name,id:t.id,desc:t.desc,params:t.params,loader:t.loader,reducer:t.reducer,forceUpdate:t.forceUpdate,externalLoad:t.externalLoad,display:t.display,load:!1};if(t.data)r.data=t.data,r.status="Loaded";else{var n=e.entityview;if(n&&t.id){var a=n.entities[t.id];a?(r.status=a.status,r.data=a.data,"Loaded"==a.status&&(r.lastUpdateTime=a.lastUpdateTime),"NotLoaded"==a.status&&(r.load=!0)):r.load=!0}}return console.log("entity....",t,r),r}),(function(e,t){return{loadEntity:function(){var r={entityName:t.name,entityId:t.id,headers:t.headers,svc:t.svc},n={global:t.global};e(Object(s.createAction)(i.ENTITY_VIEW_FETCH,r,n))}}}))(Le);r.d(t,"View",(function(){return k})),r.d(t,"ViewComponent",(function(){return L})),r.d(t,"Entity",(function(){return ke})),r.d(t,"GroupLoad",(function(){return A}))}])}));
//# sourceMappingURL=index.js.map
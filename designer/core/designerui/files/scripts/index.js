define("designerui",["uicommon","react","redux-saga"],function(t,e,n){return function(t){var e={};function n(r){if(e[r])return e[r].exports;var o=e[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=t,n.c=e,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},n.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var o in t)n.d(r,o,function(e){return t[e]}.bind(null,o));return r},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=7)}([function(e,n){e.exports=t},function(t,n){t.exports=e},function(t,e,n){t.exports=n(5)},function(t,e){t.exports=n},function(t,e){},function(t,e,n){var r=function(){return this}()||Function("return this")(),o=r.regeneratorRuntime&&Object.getOwnPropertyNames(r).indexOf("regeneratorRuntime")>=0,i=o&&r.regeneratorRuntime;if(r.regeneratorRuntime=void 0,t.exports=n(6),o)r.regeneratorRuntime=i;else try{delete r.regeneratorRuntime}catch(t){r.regeneratorRuntime=void 0}},function(t,e){!function(e){"use strict";var n,r=Object.prototype,o=r.hasOwnProperty,i="function"==typeof Symbol?Symbol:{},a=i.iterator||"@@iterator",u=i.asyncIterator||"@@asyncIterator",c=i.toStringTag||"@@toStringTag",l="object"==typeof t,f=e.regeneratorRuntime;if(f)l&&(t.exports=f);else{(f=e.regeneratorRuntime=l?t.exports:{}).wrap=w;var s="suspendedStart",h="suspendedYield",p="executing",d="completed",y={},g={};g[a]=function(){return this};var v=Object.getPrototypeOf,m=v&&v(v(P([])));m&&m!==r&&o.call(m,a)&&(g=m);var b=_.prototype=E.prototype=Object.create(g);S.prototype=b.constructor=_,_.constructor=S,_[c]=S.displayName="GeneratorFunction",f.isGeneratorFunction=function(t){var e="function"==typeof t&&t.constructor;return!!e&&(e===S||"GeneratorFunction"===(e.displayName||e.name))},f.mark=function(t){return Object.setPrototypeOf?Object.setPrototypeOf(t,_):(t.__proto__=_,c in t||(t[c]="GeneratorFunction")),t.prototype=Object.create(b),t},f.awrap=function(t){return{__await:t}},O(L.prototype),L.prototype[u]=function(){return this},f.AsyncIterator=L,f.async=function(t,e,n,r){var o=new L(w(t,e,n,r));return f.isGeneratorFunction(e)?o:o.next().then(function(t){return t.done?t.value:o.next()})},O(b),b[c]="Generator",b[a]=function(){return this},b.toString=function(){return"[object Generator]"},f.keys=function(t){var e=[];for(var n in t)e.push(n);return e.reverse(),function n(){for(;e.length;){var r=e.pop();if(r in t)return n.value=r,n.done=!1,n}return n.done=!0,n}},f.values=P,M.prototype={constructor:M,reset:function(t){if(this.prev=0,this.next=0,this.sent=this._sent=n,this.done=!1,this.delegate=null,this.method="next",this.arg=n,this.tryEntries.forEach(A),!t)for(var e in this)"t"===e.charAt(0)&&o.call(this,e)&&!isNaN(+e.slice(1))&&(this[e]=n)},stop:function(){this.done=!0;var t=this.tryEntries[0].completion;if("throw"===t.type)throw t.arg;return this.rval},dispatchException:function(t){if(this.done)throw t;var e=this;function r(r,o){return u.type="throw",u.arg=t,e.next=r,o&&(e.method="next",e.arg=n),!!o}for(var i=this.tryEntries.length-1;i>=0;--i){var a=this.tryEntries[i],u=a.completion;if("root"===a.tryLoc)return r("end");if(a.tryLoc<=this.prev){var c=o.call(a,"catchLoc"),l=o.call(a,"finallyLoc");if(c&&l){if(this.prev<a.catchLoc)return r(a.catchLoc,!0);if(this.prev<a.finallyLoc)return r(a.finallyLoc)}else if(c){if(this.prev<a.catchLoc)return r(a.catchLoc,!0)}else{if(!l)throw new Error("try statement without catch or finally");if(this.prev<a.finallyLoc)return r(a.finallyLoc)}}}},abrupt:function(t,e){for(var n=this.tryEntries.length-1;n>=0;--n){var r=this.tryEntries[n];if(r.tryLoc<=this.prev&&o.call(r,"finallyLoc")&&this.prev<r.finallyLoc){var i=r;break}}i&&("break"===t||"continue"===t)&&i.tryLoc<=e&&e<=i.finallyLoc&&(i=null);var a=i?i.completion:{};return a.type=t,a.arg=e,i?(this.method="next",this.next=i.finallyLoc,y):this.complete(a)},complete:function(t,e){if("throw"===t.type)throw t.arg;return"break"===t.type||"continue"===t.type?this.next=t.arg:"return"===t.type?(this.rval=this.arg=t.arg,this.method="return",this.next="end"):"normal"===t.type&&e&&(this.next=e),y},finish:function(t){for(var e=this.tryEntries.length-1;e>=0;--e){var n=this.tryEntries[e];if(n.finallyLoc===t)return this.complete(n.completion,n.afterLoc),A(n),y}},catch:function(t){for(var e=this.tryEntries.length-1;e>=0;--e){var n=this.tryEntries[e];if(n.tryLoc===t){var r=n.completion;if("throw"===r.type){var o=r.arg;A(n)}return o}}throw new Error("illegal catch attempt")},delegateYield:function(t,e,r){return this.delegate={iterator:P(t),resultName:e,nextLoc:r},"next"===this.method&&(this.arg=n),y}}}function w(t,e,n,r){var o=e&&e.prototype instanceof E?e:E,i=Object.create(o.prototype),a=new M(r||[]);return i._invoke=function(t,e,n){var r=s;return function(o,i){if(r===p)throw new Error("Generator is already running");if(r===d){if("throw"===o)throw i;return C()}for(n.method=o,n.arg=i;;){var a=n.delegate;if(a){var u=j(a,n);if(u){if(u===y)continue;return u}}if("next"===n.method)n.sent=n._sent=n.arg;else if("throw"===n.method){if(r===s)throw r=d,n.arg;n.dispatchException(n.arg)}else"return"===n.method&&n.abrupt("return",n.arg);r=p;var c=x(t,e,n);if("normal"===c.type){if(r=n.done?d:h,c.arg===y)continue;return{value:c.arg,done:n.done}}"throw"===c.type&&(r=d,n.method="throw",n.arg=c.arg)}}}(t,n,a),i}function x(t,e,n){try{return{type:"normal",arg:t.call(e,n)}}catch(t){return{type:"throw",arg:t}}}function E(){}function S(){}function _(){}function O(t){["next","throw","return"].forEach(function(e){t[e]=function(t){return this._invoke(e,t)}})}function L(t){var e;this._invoke=function(n,r){function i(){return new Promise(function(e,i){!function e(n,r,i,a){var u=x(t[n],t,r);if("throw"!==u.type){var c=u.arg,l=c.value;return l&&"object"==typeof l&&o.call(l,"__await")?Promise.resolve(l.__await).then(function(t){e("next",t,i,a)},function(t){e("throw",t,i,a)}):Promise.resolve(l).then(function(t){c.value=t,i(c)},a)}a(u.arg)}(n,r,e,i)})}return e=e?e.then(i,i):i()}}function j(t,e){var r=t.iterator[e.method];if(r===n){if(e.delegate=null,"throw"===e.method){if(t.iterator.return&&(e.method="return",e.arg=n,j(t,e),"throw"===e.method))return y;e.method="throw",e.arg=new TypeError("The iterator does not provide a 'throw' method")}return y}var o=x(r,t.iterator,e.arg);if("throw"===o.type)return e.method="throw",e.arg=o.arg,e.delegate=null,y;var i=o.arg;return i?i.done?(e[t.resultName]=i.value,e.next=t.nextLoc,"return"!==e.method&&(e.method="next",e.arg=n),e.delegate=null,y):i:(e.method="throw",e.arg=new TypeError("iterator result is not an object"),e.delegate=null,y)}function k(t){var e={tryLoc:t[0]};1 in t&&(e.catchLoc=t[1]),2 in t&&(e.finallyLoc=t[2],e.afterLoc=t[3]),this.tryEntries.push(e)}function A(t){var e=t.completion||{};e.type="normal",delete e.arg,t.completion=e}function M(t){this.tryEntries=[{tryLoc:"root"}],t.forEach(k,this),this.reset(!0)}function P(t){if(t){var e=t[a];if(e)return e.call(t);if("function"==typeof t.next)return t;if(!isNaN(t.length)){var r=-1,i=function e(){for(;++r<t.length;)if(o.call(t,r))return e.value=t[r],e.done=!1,e;return e.value=n,e.done=!0,e};return i.next=i}}return{next:C}}function C(){return{value:n,done:!0}}}(function(){return this}()||Function("return this")())},function(t,e,n){"use strict";n.r(e);var r=n(1),o=n.n(r),i=function(t){return o.a.createElement("div",{className:"welcomepage"},"some page")},a=(n(4),{SYNC_OBJECTS:"SYNC_OBJECTS"}),u=n(2),c=n.n(u),l=n(3),f=(Object.assign,"function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t}),s=function(t){return"@@redux-saga/"+t},h=s("TASK"),p=s("HELPER");function d(t,e,n){if(!e(t))throw function(t,e){var n=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"";"undefined"==typeof window?console.log("redux-saga "+t+": "+e+"\n"+(n&&n.stack||n)):console[t](e,n)}("error","uncaught at check",n),new Error(n)}var y=Object.prototype.hasOwnProperty;function g(t,e){return v.notUndef(t)&&y.call(t,e)}var v={undef:function(t){return null==t},notUndef:function(t){return null!=t},func:function(t){return"function"==typeof t},number:function(t){return"number"==typeof t},string:function(t){return"string"==typeof t},array:Array.isArray,object:function(t){return t&&!v.array(t)&&"object"===(void 0===t?"undefined":f(t))},promise:function(t){return t&&v.func(t.then)},iterator:function(t){return t&&v.func(t.next)&&v.func(t.throw)},iterable:function(t){return t&&v.func(Symbol)?v.func(t[Symbol.iterator]):v.array(t)},task:function(t){return t&&t[h]},observable:function(t){return t&&v.func(t.subscribe)},buffer:function(t){return t&&v.func(t.isEmpty)&&v.func(t.take)&&v.func(t.put)},pattern:function(t){return t&&(v.string(t)||"symbol"===(void 0===t?"undefined":f(t))||v.func(t)||v.array(t))},channel:function(t){return t&&v.func(t.take)&&v.func(t.close)},helper:function(t){return t&&t[p]},stringableFunc:function(t){return v.func(t)&&g(t,"toString")}};function m(t,e){return function(){return t.apply(void 0,arguments)}}var b=s("IO"),w="TAKE",x="PUT",E="CALL",S=function(t,e){var n;return(n={})[b]=!0,n[t]=e,n};function _(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"*";if(arguments.length&&d(arguments[0],v.notUndef,"take(patternOrChannel): patternOrChannel is undefined"),v.pattern(t))return S(w,{pattern:t});if(v.channel(t))return S(w,{channel:t});throw new Error("take(patternOrChannel): argument "+String(t)+" is not valid channel or a valid pattern")}_.maybe=function(){var t=_.apply(void 0,arguments);return t[w].maybe=!0,t};_.maybe;function O(t,e){return arguments.length>1?(d(t,v.notUndef,"put(channel, action): argument channel is undefined"),d(t,v.channel,"put(channel, action): argument "+t+" is not a valid channel"),d(e,v.notUndef,"put(channel, action): argument action is undefined")):(d(t,v.notUndef,"put(action): argument action is undefined"),e=t,t=null),S(x,{channel:t,action:e})}function L(t,e,n){d(e,v.notUndef,t+": argument fn is undefined");var r=null;if(v.array(e)){var o=e;r=o[0],e=o[1]}else if(e.fn){var i=e;r=i.context,e=i.fn}return r&&v.string(e)&&v.func(r[e])&&(e=r[e]),d(e,v.func,t+": argument "+e+" is not a function"),{context:r,fn:e,args:n}}function j(t){for(var e=arguments.length,n=Array(e>1?e-1:0),r=1;r<e;r++)n[r-1]=arguments[r];return S(E,L("call",t,n))}O.resolve=function(){var t=O.apply(void 0,arguments);return t[x].resolve=!0,t},O.sync=m(O.resolve);Object.assign;var k=n(0),A=c.a.mark(P),M=c.a.mark(C);function P(t){var e,n;return c.a.wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return console.log("syncing objects",t),r.prev=1,e=k.RequestBuilder.DefaultRequest(null,t.payload),r.next=5,j(k.DataSource.ExecuteService,t.meta.type+"resolver",e);case 5:n=r.sent,console.log("resolver",n),r.next=13;break;case 9:r.prev=9,r.t0=r.catch(1),console.log("sync objects",r.t0),Window.handleError&&Window.handleError(r.t0);case 13:case"end":return r.stop()}},A,this,[[1,9]])}function C(){return c.a.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,[Object(l.takeEvery)(a.SYNC_OBJECTS,P)];case 2:case"end":return t.stop()}},M,this)}function N(t,e,n){console.log("Form_Instance_Modules  called",t,e,n),n(e.parentFormValue&&e.parentFormValue.Modules?e.parentFormValue.Modules:[])}function R(t,e,n,r){var o=t.config.entity.toLowerCase();return console.log("form sync",o),function(t){console.log("data",t),r(Object(k.createAction)(a.SYNC_OBJECTS,t,{type:o,setData:n}))}}function T(t,e,n){console.log("abstract server solution modules--------",t,e);var r=e.formValue.Solution;r&&k.EntityData.GetEntity("Solution",r).then(function(t){return function(e){console.log("received entity data",e),e.data&&e.data.Modules&&t(e.data.Modules)}}(n),function(t){console.log("Error in fetching solution modules",t)})}function F(t,e,n,r,i){return console.log("my action buttons",i),o.a.createElement(_uikit.Block,null,o.a.createElement(_uikit.ActionButton,{onClick:e(),className:"submitBtn"},"Save"),o.a.createElement(_uikit.ActionButton,{onClick:e(R(t,0,r,i)),className:""},"Sync Modules"))}function B(t){console.log("Params******** modulesrepo",t),t.ctx.panel.overlayComponent(o.a.createElement("h2",null,"my module"))}function G(t,e,n,r,o,i){console.log("Initializing ui"),_r("Methods","Form_Instance_Modules",N),_r("Methods","AbstractServer_Actions",F),_r("Methods","AbstractServer_Solution_Modules",T),_r("Actions","ModulesRepo_viewModule",{actiontype:"method",method:B}),console.log("registering method",Application)}Application.Register("Sagas","designerSaga",C),n.d(e,"Initialize",function(){return G}),n.d(e,"UsersView",function(){return i})}])});
//# sourceMappingURL=index.js.map
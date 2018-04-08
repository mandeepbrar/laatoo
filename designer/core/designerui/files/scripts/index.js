define("designerui",["react","uicommon","redux-saga"],function(t,e,n){return function(t){var e={};function n(r){if(e[r])return e[r].exports;var o=e[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=t,n.c=e,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{configurable:!1,enumerable:!0,get:r})},n.r=function(t){Object.defineProperty(t,"__esModule",{value:!0})},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=4)}([function(e,n){e.exports=t},function(t,n){t.exports=e},function(t,e,n){t.exports=n(6)},function(t,e){t.exports=n},function(t,e,n){"use strict";n.r(e);var r=n(0),o=n.n(r),i=function(t){return o.a.createElement("div",{className:"welcomepage"},"some page")},a=(n(7),{SYNC_OBJECTS:"SYNC_OBJECTS"}),u=n(2),c=n.n(u),f=n(3),l=(Object.assign,"function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t}),s=function(t){return"@@redux-saga/"+t},h=s("TASK"),p=s("HELPER");function d(t,e,n){if(!e(t))throw function(t,e){var n=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"";"undefined"==typeof window?console.log("redux-saga "+t+": "+e+"\n"+(n&&n.stack||n)):console[t](e,n)}("error","uncaught at check",n),new Error(n)}var y=Object.prototype.hasOwnProperty;function v(t,e){return g.notUndef(t)&&y.call(t,e)}var g={undef:function(t){return null===t||void 0===t},notUndef:function(t){return null!==t&&void 0!==t},func:function(t){return"function"==typeof t},number:function(t){return"number"==typeof t},string:function(t){return"string"==typeof t},array:Array.isArray,object:function(t){return t&&!g.array(t)&&"object"===(void 0===t?"undefined":l(t))},promise:function(t){return t&&g.func(t.then)},iterator:function(t){return t&&g.func(t.next)&&g.func(t.throw)},iterable:function(t){return t&&g.func(Symbol)?g.func(t[Symbol.iterator]):g.array(t)},task:function(t){return t&&t[h]},observable:function(t){return t&&g.func(t.subscribe)},buffer:function(t){return t&&g.func(t.isEmpty)&&g.func(t.take)&&g.func(t.put)},pattern:function(t){return t&&(g.string(t)||"symbol"===(void 0===t?"undefined":l(t))||g.func(t)||g.array(t))},channel:function(t){return t&&g.func(t.take)&&g.func(t.close)},helper:function(t){return t&&t[p]},stringableFunc:function(t){return g.func(t)&&v(t,"toString")}};function m(t,e){return function(){return t.apply(void 0,arguments)}}Object.assign;var w=s("IO"),b="TAKE",x="PUT",E="CALL",L=function(t,e){var n;return(n={})[w]=!0,n[t]=e,n};function O(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"*";if(arguments.length&&d(arguments[0],g.notUndef,"take(patternOrChannel): patternOrChannel is undefined"),g.pattern(t))return L(b,{pattern:t});if(g.channel(t))return L(b,{channel:t});throw new Error("take(patternOrChannel): argument "+String(t)+" is not valid channel or a valid pattern")}O.maybe=function(){var t=O.apply(void 0,arguments);return t[b].maybe=!0,t};O.maybe;function S(t,e){return arguments.length>1?(d(t,g.notUndef,"put(channel, action): argument channel is undefined"),d(t,g.channel,"put(channel, action): argument "+t+" is not a valid channel"),d(e,g.notUndef,"put(channel, action): argument action is undefined")):(d(t,g.notUndef,"put(action): argument action is undefined"),e=t,t=null),L(x,{channel:t,action:e})}function _(t,e,n){d(e,g.notUndef,t+": argument fn is undefined");var r=null;if(g.array(e)){var o=e;r=o[0],e=o[1]}else if(e.fn){var i=e;r=i.context,e=i.fn}return r&&g.string(e)&&g.func(r[e])&&(e=r[e]),d(e,g.func,t+": argument "+e+" is not a function"),{context:r,fn:e,args:n}}function j(t){for(var e=arguments.length,n=Array(e>1?e-1:0),r=1;r<e;r++)n[r-1]=arguments[r];return L(E,_("call",t,n))}S.resolve=function(){var t=S.apply(void 0,arguments);return t[x].resolve=!0,t},S.sync=m(S.resolve);var k=n(1),N=c.a.mark(C),A=c.a.mark(P);function C(t){var e,n;return c.a.wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return console.log("syncing objects",t),r.prev=1,e=k.RequestBuilder.DefaultRequest(null,t.payload),r.next=5,j(k.DataSource.ExecuteService,t.meta.type+"resolver",e);case 5:n=r.sent,console.log("resolver",n),r.next=13;break;case 9:r.prev=9,r.t0=r.catch(1),console.log("sync objects",r.t0),Window.handleError&&Window.handleError(r.t0);case 13:case"end":return r.stop()}},N,this,[[1,9]])}function P(){return c.a.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,[Object(f.takeEvery)(a.SYNC_OBJECTS,C)];case 2:case"end":return t.stop()}},A,this)}function R(t,e,n,r,o,i,a){console.log("transformer called",o,i,t,n,r),console.log("parent value",o.parentFormValue);var u=o.parentFormValue.Modules,c=[];return u&&u.forEach(function(t){console.log("module.....",t),c.push({text:t.Name,value:t.Name})}),i.additionalProperties.items=c,t}function T(t,e,n,r,i,u){return console.log("my action buttons",u),o.a.createElement(r.Block,null,o.a.createElement(r.ActionButton,{onClick:e(),className:"submitBtn"},"Save"),o.a.createElement(r.ActionButton,{onClick:e(function(t,e,n,r){var o=t.config.entity.toLowerCase();return console.log("form sync",o),function(t){console.log("data",t),r(Object(k.createAction)(a.SYNC_OBJECTS,t,{type:o,setData:n}))}}(t,0,i,u)),className:""},"Sync Modules"))}function F(t){console.log("Params******** modulesrepo",t),t.ctx.panel.overlayComponent(o.a.createElement("h2",null,"my module"))}function B(t,e,n,r,o,i){console.log("Initializing ui"),_r("Methods","Form_Instance_Transform_Modules",R),_r("Methods","AbstractServer_Actions",T),_r("Actions","ModulesRepo_viewModule",{actiontype:"method",method:F}),console.log("registering method",Application)}Application.Register("Sagas","designerSaga",P),n.d(e,"Initialize",function(){return B}),n.d(e,"UsersView",function(){return i})},function(t,e){!function(e){"use strict";var n,r=Object.prototype,o=r.hasOwnProperty,i="function"==typeof Symbol?Symbol:{},a=i.iterator||"@@iterator",u=i.asyncIterator||"@@asyncIterator",c=i.toStringTag||"@@toStringTag",f="object"==typeof t,l=e.regeneratorRuntime;if(l)f&&(t.exports=l);else{(l=e.regeneratorRuntime=f?t.exports:{}).wrap=b;var s="suspendedStart",h="suspendedYield",p="executing",d="completed",y={},v={};v[a]=function(){return this};var g=Object.getPrototypeOf,m=g&&g(g(C([])));m&&m!==r&&o.call(m,a)&&(v=m);var w=O.prototype=E.prototype=Object.create(v);L.prototype=w.constructor=O,O.constructor=L,O[c]=L.displayName="GeneratorFunction",l.isGeneratorFunction=function(t){var e="function"==typeof t&&t.constructor;return!!e&&(e===L||"GeneratorFunction"===(e.displayName||e.name))},l.mark=function(t){return Object.setPrototypeOf?Object.setPrototypeOf(t,O):(t.__proto__=O,c in t||(t[c]="GeneratorFunction")),t.prototype=Object.create(w),t},l.awrap=function(t){return{__await:t}},S(_.prototype),_.prototype[u]=function(){return this},l.AsyncIterator=_,l.async=function(t,e,n,r){var o=new _(b(t,e,n,r));return l.isGeneratorFunction(e)?o:o.next().then(function(t){return t.done?t.value:o.next()})},S(w),w[c]="Generator",w[a]=function(){return this},w.toString=function(){return"[object Generator]"},l.keys=function(t){var e=[];for(var n in t)e.push(n);return e.reverse(),function n(){for(;e.length;){var r=e.pop();if(r in t)return n.value=r,n.done=!1,n}return n.done=!0,n}},l.values=C,A.prototype={constructor:A,reset:function(t){if(this.prev=0,this.next=0,this.sent=this._sent=n,this.done=!1,this.delegate=null,this.method="next",this.arg=n,this.tryEntries.forEach(N),!t)for(var e in this)"t"===e.charAt(0)&&o.call(this,e)&&!isNaN(+e.slice(1))&&(this[e]=n)},stop:function(){this.done=!0;var t=this.tryEntries[0].completion;if("throw"===t.type)throw t.arg;return this.rval},dispatchException:function(t){if(this.done)throw t;var e=this;function r(r,o){return u.type="throw",u.arg=t,e.next=r,o&&(e.method="next",e.arg=n),!!o}for(var i=this.tryEntries.length-1;i>=0;--i){var a=this.tryEntries[i],u=a.completion;if("root"===a.tryLoc)return r("end");if(a.tryLoc<=this.prev){var c=o.call(a,"catchLoc"),f=o.call(a,"finallyLoc");if(c&&f){if(this.prev<a.catchLoc)return r(a.catchLoc,!0);if(this.prev<a.finallyLoc)return r(a.finallyLoc)}else if(c){if(this.prev<a.catchLoc)return r(a.catchLoc,!0)}else{if(!f)throw new Error("try statement without catch or finally");if(this.prev<a.finallyLoc)return r(a.finallyLoc)}}}},abrupt:function(t,e){for(var n=this.tryEntries.length-1;n>=0;--n){var r=this.tryEntries[n];if(r.tryLoc<=this.prev&&o.call(r,"finallyLoc")&&this.prev<r.finallyLoc){var i=r;break}}i&&("break"===t||"continue"===t)&&i.tryLoc<=e&&e<=i.finallyLoc&&(i=null);var a=i?i.completion:{};return a.type=t,a.arg=e,i?(this.method="next",this.next=i.finallyLoc,y):this.complete(a)},complete:function(t,e){if("throw"===t.type)throw t.arg;return"break"===t.type||"continue"===t.type?this.next=t.arg:"return"===t.type?(this.rval=this.arg=t.arg,this.method="return",this.next="end"):"normal"===t.type&&e&&(this.next=e),y},finish:function(t){for(var e=this.tryEntries.length-1;e>=0;--e){var n=this.tryEntries[e];if(n.finallyLoc===t)return this.complete(n.completion,n.afterLoc),N(n),y}},catch:function(t){for(var e=this.tryEntries.length-1;e>=0;--e){var n=this.tryEntries[e];if(n.tryLoc===t){var r=n.completion;if("throw"===r.type){var o=r.arg;N(n)}return o}}throw new Error("illegal catch attempt")},delegateYield:function(t,e,r){return this.delegate={iterator:C(t),resultName:e,nextLoc:r},"next"===this.method&&(this.arg=n),y}}}function b(t,e,n,r){var o=e&&e.prototype instanceof E?e:E,i=Object.create(o.prototype),a=new A(r||[]);return i._invoke=function(t,e,n){var r=s;return function(o,i){if(r===p)throw new Error("Generator is already running");if(r===d){if("throw"===o)throw i;return P()}for(n.method=o,n.arg=i;;){var a=n.delegate;if(a){var u=j(a,n);if(u){if(u===y)continue;return u}}if("next"===n.method)n.sent=n._sent=n.arg;else if("throw"===n.method){if(r===s)throw r=d,n.arg;n.dispatchException(n.arg)}else"return"===n.method&&n.abrupt("return",n.arg);r=p;var c=x(t,e,n);if("normal"===c.type){if(r=n.done?d:h,c.arg===y)continue;return{value:c.arg,done:n.done}}"throw"===c.type&&(r=d,n.method="throw",n.arg=c.arg)}}}(t,n,a),i}function x(t,e,n){try{return{type:"normal",arg:t.call(e,n)}}catch(t){return{type:"throw",arg:t}}}function E(){}function L(){}function O(){}function S(t){["next","throw","return"].forEach(function(e){t[e]=function(t){return this._invoke(e,t)}})}function _(t){var e;this._invoke=function(n,r){function i(){return new Promise(function(e,i){!function e(n,r,i,a){var u=x(t[n],t,r);if("throw"!==u.type){var c=u.arg,f=c.value;return f&&"object"==typeof f&&o.call(f,"__await")?Promise.resolve(f.__await).then(function(t){e("next",t,i,a)},function(t){e("throw",t,i,a)}):Promise.resolve(f).then(function(t){c.value=t,i(c)},a)}a(u.arg)}(n,r,e,i)})}return e=e?e.then(i,i):i()}}function j(t,e){var r=t.iterator[e.method];if(r===n){if(e.delegate=null,"throw"===e.method){if(t.iterator.return&&(e.method="return",e.arg=n,j(t,e),"throw"===e.method))return y;e.method="throw",e.arg=new TypeError("The iterator does not provide a 'throw' method")}return y}var o=x(r,t.iterator,e.arg);if("throw"===o.type)return e.method="throw",e.arg=o.arg,e.delegate=null,y;var i=o.arg;return i?i.done?(e[t.resultName]=i.value,e.next=t.nextLoc,"return"!==e.method&&(e.method="next",e.arg=n),e.delegate=null,y):i:(e.method="throw",e.arg=new TypeError("iterator result is not an object"),e.delegate=null,y)}function k(t){var e={tryLoc:t[0]};1 in t&&(e.catchLoc=t[1]),2 in t&&(e.finallyLoc=t[2],e.afterLoc=t[3]),this.tryEntries.push(e)}function N(t){var e=t.completion||{};e.type="normal",delete e.arg,t.completion=e}function A(t){this.tryEntries=[{tryLoc:"root"}],t.forEach(k,this),this.reset(!0)}function C(t){if(t){var e=t[a];if(e)return e.call(t);if("function"==typeof t.next)return t;if(!isNaN(t.length)){var r=-1,i=function e(){for(;++r<t.length;)if(o.call(t,r))return e.value=t[r],e.done=!1,e;return e.value=n,e.done=!0,e};return i.next=i}}return{next:P}}function P(){return{value:n,done:!0}}}(function(){return this}()||Function("return this")())},function(t,e,n){var r=function(){return this}()||Function("return this")(),o=r.regeneratorRuntime&&Object.getOwnPropertyNames(r).indexOf("regeneratorRuntime")>=0,i=o&&r.regeneratorRuntime;if(r.regeneratorRuntime=void 0,t.exports=n(5),o)r.regeneratorRuntime=i;else try{delete r.regeneratorRuntime}catch(t){r.regeneratorRuntime=void 0}},function(t,e){}])});
//# sourceMappingURL=index.js.map
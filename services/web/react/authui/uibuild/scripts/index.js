define("authui",["react","uicommon","prop-types","react-redux","reactwebcommon","md5","redux"],(function(e,t,n,r,o,a,i){return function(e){var t={};function n(r){if(t[r])return t[r].exports;var o=t[r]={i:r,l:!1,exports:{}};return e[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=e,n.c=t,n.d=function(e,t,r){n.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:r})},n.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},n.t=function(e,t){if(1&t&&(e=n(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var o in e)n.d(r,o,function(t){return e[t]}.bind(null,o));return r},n.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return n.d(t,"a",t),t},n.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},n.p="/",n(n.s=18)}([function(t,n){t.exports=e},function(e,n){e.exports=t},function(e,t,n){e.exports=n(17)},function(e,t){e.exports=n},function(e,t){e.exports=function(e){if(void 0===e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return e}},function(e,t){e.exports=function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}},function(e,t){function n(e,t){for(var n=0;n<t.length;n++){var r=t[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),Object.defineProperty(e,r.key,r)}}e.exports=function(e,t,r){return t&&n(e.prototype,t),r&&n(e,r),e}},function(e,t,n){var r=n(14),o=n(4);e.exports=function(e,t){return!t||"object"!==r(t)&&"function"!=typeof t?o(e):t}},function(e,t){function n(t){return e.exports=n=Object.setPrototypeOf?Object.getPrototypeOf:function(e){return e.__proto__||Object.getPrototypeOf(e)},n(t)}e.exports=n},function(e,t,n){var r=n(15);e.exports=function(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Super expression must either be null or a function");e.prototype=Object.create(t&&t.prototype,{constructor:{value:e,writable:!0,configurable:!0}}),t&&r(e,t)}},function(e,t){e.exports=r},function(e,t){e.exports=o},function(e,t){e.exports=a},function(e,t){},function(e,t){function n(e){return(n="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e})(e)}function r(t){return"function"==typeof Symbol&&"symbol"===n(Symbol.iterator)?e.exports=r=function(e){return n(e)}:e.exports=r=function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":n(e)},r(t)}e.exports=r},function(e,t){function n(t,r){return e.exports=n=Object.setPrototypeOf||function(e,t){return e.__proto__=t,e},n(t,r)}e.exports=n},function(e,t){e.exports=i},function(e,t,n){var r=function(e){"use strict";var t,n=Object.prototype,r=n.hasOwnProperty,o="function"==typeof Symbol?Symbol:{},a=o.iterator||"@@iterator",i=o.asyncIterator||"@@asyncIterator",c=o.toStringTag||"@@toStringTag";function u(e,t,n,r){var o=t&&t.prototype instanceof m?t:m,a=Object.create(o.prototype),i=new O(r||[]);return a._invoke=function(e,t,n){var r=l;return function(o,a){if(r===f)throw new Error("Generator is already running");if(r===d){if("throw"===o)throw a;return _()}for(n.method=o,n.arg=a;;){var i=n.delegate;if(i){var c=L(i,n);if(c){if(c===g)continue;return c}}if("next"===n.method)n.sent=n._sent=n.arg;else if("throw"===n.method){if(r===l)throw r=d,n.arg;n.dispatchException(n.arg)}else"return"===n.method&&n.abrupt("return",n.arg);r=f;var u=s(e,t,n);if("normal"===u.type){if(r=n.done?d:p,u.arg===g)continue;return{value:u.arg,done:n.done}}"throw"===u.type&&(r=d,n.method="throw",n.arg=u.arg)}}}(e,n,i),a}function s(e,t,n){try{return{type:"normal",arg:e.call(t,n)}}catch(e){return{type:"throw",arg:e}}}e.wrap=u;var l="suspendedStart",p="suspendedYield",f="executing",d="completed",g={};function m(){}function h(){}function v(){}var y={};y[a]=function(){return this};var S=Object.getPrototypeOf,b=S&&S(S(A([])));b&&b!==n&&r.call(b,a)&&(y=b);var w=v.prototype=m.prototype=Object.create(y);function x(e){["next","throw","return"].forEach((function(t){e[t]=function(e){return this._invoke(t,e)}}))}function N(e){var t;this._invoke=function(n,o){function a(){return new Promise((function(t,a){!function t(n,o,a,i){var c=s(e[n],e,o);if("throw"!==c.type){var u=c.arg,l=u.value;return l&&"object"==typeof l&&r.call(l,"__await")?Promise.resolve(l.__await).then((function(e){t("next",e,a,i)}),(function(e){t("throw",e,a,i)})):Promise.resolve(l).then((function(e){u.value=e,a(u)}),(function(e){return t("throw",e,a,i)}))}i(c.arg)}(n,o,t,a)}))}return t=t?t.then(a,a):a()}}function L(e,n){var r=e.iterator[n.method];if(r===t){if(n.delegate=null,"throw"===n.method){if(e.iterator.return&&(n.method="return",n.arg=t,L(e,n),"throw"===n.method))return g;n.method="throw",n.arg=new TypeError("The iterator does not provide a 'throw' method")}return g}var o=s(r,e.iterator,n.arg);if("throw"===o.type)return n.method="throw",n.arg=o.arg,n.delegate=null,g;var a=o.arg;return a?a.done?(n[e.resultName]=a.value,n.next=e.nextLoc,"return"!==n.method&&(n.method="next",n.arg=t),n.delegate=null,g):a:(n.method="throw",n.arg=new TypeError("iterator result is not an object"),n.delegate=null,g)}function I(e){var t={tryLoc:e[0]};1 in e&&(t.catchLoc=e[1]),2 in e&&(t.finallyLoc=e[2],t.afterLoc=e[3]),this.tryEntries.push(t)}function E(e){var t=e.completion||{};t.type="normal",delete t.arg,e.completion=t}function O(e){this.tryEntries=[{tryLoc:"root"}],e.forEach(I,this),this.reset(!0)}function A(e){if(e){var n=e[a];if(n)return n.call(e);if("function"==typeof e.next)return e;if(!isNaN(e.length)){var o=-1,i=function n(){for(;++o<e.length;)if(r.call(e,o))return n.value=e[o],n.done=!1,n;return n.value=t,n.done=!0,n};return i.next=i}}return{next:_}}function _(){return{value:t,done:!0}}return h.prototype=w.constructor=v,v.constructor=h,v[c]=h.displayName="GeneratorFunction",e.isGeneratorFunction=function(e){var t="function"==typeof e&&e.constructor;return!!t&&(t===h||"GeneratorFunction"===(t.displayName||t.name))},e.mark=function(e){return Object.setPrototypeOf?Object.setPrototypeOf(e,v):(e.__proto__=v,c in e||(e[c]="GeneratorFunction")),e.prototype=Object.create(w),e},e.awrap=function(e){return{__await:e}},x(N.prototype),N.prototype[i]=function(){return this},e.AsyncIterator=N,e.async=function(t,n,r,o){var a=new N(u(t,n,r,o));return e.isGeneratorFunction(n)?a:a.next().then((function(e){return e.done?e.value:a.next()}))},x(w),w[c]="Generator",w[a]=function(){return this},w.toString=function(){return"[object Generator]"},e.keys=function(e){var t=[];for(var n in e)t.push(n);return t.reverse(),function n(){for(;t.length;){var r=t.pop();if(r in e)return n.value=r,n.done=!1,n}return n.done=!0,n}},e.values=A,O.prototype={constructor:O,reset:function(e){if(this.prev=0,this.next=0,this.sent=this._sent=t,this.done=!1,this.delegate=null,this.method="next",this.arg=t,this.tryEntries.forEach(E),!e)for(var n in this)"t"===n.charAt(0)&&r.call(this,n)&&!isNaN(+n.slice(1))&&(this[n]=t)},stop:function(){this.done=!0;var e=this.tryEntries[0].completion;if("throw"===e.type)throw e.arg;return this.rval},dispatchException:function(e){if(this.done)throw e;var n=this;function o(r,o){return c.type="throw",c.arg=e,n.next=r,o&&(n.method="next",n.arg=t),!!o}for(var a=this.tryEntries.length-1;a>=0;--a){var i=this.tryEntries[a],c=i.completion;if("root"===i.tryLoc)return o("end");if(i.tryLoc<=this.prev){var u=r.call(i,"catchLoc"),s=r.call(i,"finallyLoc");if(u&&s){if(this.prev<i.catchLoc)return o(i.catchLoc,!0);if(this.prev<i.finallyLoc)return o(i.finallyLoc)}else if(u){if(this.prev<i.catchLoc)return o(i.catchLoc,!0)}else{if(!s)throw new Error("try statement without catch or finally");if(this.prev<i.finallyLoc)return o(i.finallyLoc)}}}},abrupt:function(e,t){for(var n=this.tryEntries.length-1;n>=0;--n){var o=this.tryEntries[n];if(o.tryLoc<=this.prev&&r.call(o,"finallyLoc")&&this.prev<o.finallyLoc){var a=o;break}}a&&("break"===e||"continue"===e)&&a.tryLoc<=t&&t<=a.finallyLoc&&(a=null);var i=a?a.completion:{};return i.type=e,i.arg=t,a?(this.method="next",this.next=a.finallyLoc,g):this.complete(i)},complete:function(e,t){if("throw"===e.type)throw e.arg;return"break"===e.type||"continue"===e.type?this.next=e.arg:"return"===e.type?(this.rval=this.arg=e.arg,this.method="return",this.next="end"):"normal"===e.type&&t&&(this.next=t),g},finish:function(e){for(var t=this.tryEntries.length-1;t>=0;--t){var n=this.tryEntries[t];if(n.finallyLoc===e)return this.complete(n.completion,n.afterLoc),E(n),g}},catch:function(e){for(var t=this.tryEntries.length-1;t>=0;--t){var n=this.tryEntries[t];if(n.tryLoc===e){var r=n.completion;if("throw"===r.type){var o=r.arg;E(n)}return o}}throw new Error("illegal catch attempt")},delegateYield:function(e,n,r){return this.delegate={iterator:A(e),resultName:n,nextLoc:r},"next"===this.method&&(this.arg=t),g}},e}(e.exports);try{regeneratorRuntime=r}catch(e){Function("r","regeneratorRuntime = r")(r)}},function(e,t,n){"use strict";n.r(t);var r=n(0),o=n.n(r),a=n(5),i=n.n(a),c=n(6),u=n.n(c),s=n(7),l=n.n(s),p=n(8),f=n.n(p),d=n(4),g=n.n(d),m=n(9),h=n.n(m),v=n(3),y=function(e){function t(e){var n;i()(this,t),n=l()(this,f()(t).call(this,e)),console.log("costructor of login web"),n.state={email:"",password:""},n.handleLogin=n.handleLogin.bind(g()(n)),n.handleChange=n.handleChange.bind(g()(n));return e.realm&&"?Realm="+e.realm,n}return h()(t,e),u()(t,[{key:"handleChange",value:function(e){var t={};t[e.target.name]=e.target.value,this.setState(t)}},{key:"handleLogin",value:function(){this.props.handleLogin(this.state.email,this.state.password)}},{key:"render",value:function(){return console.log("login ui",this.props),this.props.renderLogin(this.state,this.handleChange,this.handleLogin,this.oauthLogin,this.props)}}]),t}(o.a.Component);y.propTypes={handleOauthLogin:v.func.isRequired,handleLogin:v.func.isRequired};var S=n(12),b=n.n(S),w=n(10),x={LOGIN:"LOGIN",LOGGING_IN:"LOGGING_IN",LOGIN_SUCCESS:"LOGIN_SUCCESS",LOGIN_FAILURE:"LOGIN_FAILURE",LOGOUT:"LOGOUT",LOGOUT_SUCCESS:"LOGOUT_SUCCESS",VALIDATIONFAILED:"VALIDATIONFAILED",SIGN_UP:"SIGN_UP",SIGNING_UP:"SIGNING_UP",SIGNUP_SUCCESS:"SIGNUP_SUCCESS",SIGNUP_FAILURE:"SIGNUP_FAILURE"},N=n(1),L=n(3),I=n.n(L),E=Object(w.connect)((function(e,t){return{realm:Application.Security.realm,renderLogin:t.renderLogin,signup:t.signup}}),(function(e,t){console.log("map dispatch of login compoent");var n="";return Application.Security.realm&&(n=Application.Security.realm),{handleLogin:function(t,r){var o={Username:t,Password:b()(r),Realm:n},a={serviceName:Application.Security.loginService};e(Object(N.createAction)(x.LOGIN,o,a))},handleOauthLogin:function(t){e(Object(N.createAction)(x.LOGIN_SUCCESS,{userId:t.id,token:t.token,permissions:t.permissions}))}}}))(y);E.propTypes={loginService:I.a.string.isRequired,successpage:I.a.string,realm:I.a.string,signup:I.a.string};n(16);var O,A=n(3);var _=function(e){function t(e){var n;return i()(this,t),(n=l()(this,f()(t).call(this,e))).validatetoken=n.validatetoken.bind(g()(n)),n.state={loggedIn:e.loggedIn,validation:e.validation},console.log("login component",e),e.validation&&n.validatetoken(),n}return h()(t,e),u()(t,[{key:"componentWillReceiveProps",value:function(e){e.loggedIn==this.state.loggedIn&&e.validation==this.state.validation||this.setState({loggedIn:e.loggedIn,validation:e.validation})}},{key:"validatetoken",value:function(){var e=this,t=this.props.logout,n=this.props.login;console.log("sending validation request");var r=N.RequestBuilder.DefaultRequest({},{});N.DataSource.ExecuteService(this.props.validateService,r).then((function(e){n(e.data.Id,e.data.Permissions)}),(function(n){t(!0),e.setState({loggedIn:!1,validation:!1})}))}},{key:"getChildContext",value:function(){return{loggedIn:this.state.loggedIn}}},{key:"render",value:function(){return this.state.validation?null:this.props.children?o.a.cloneElement(this.props.children,{loggedIn:this.state.loggedIn,validation:this.state.validation}):null}}]),t}(o.a.Component);_.childContextTypes={loggedIn:A.bool,user:A.object};var k=Object(w.connect)((function(e,t){return console.log("Login validator ",e,t,O,Storage),O.settings.cookies?function(e,t){switch(e.Security.status){case"NotLogged":return{validation:!0,loggedIn:!1,validateService:t.validateService};case"LoggedIn":return{validation:!1,loggedIn:!0};default:return{loggedIn:!1}}}(e,t):function(e,t){switch(e.Security.status){case"NotLogged":return console.log("get props non cookie, Storage.auth = ",Storage.auth),Storage.auth?{validation:!0,loggedIn:!1,validateService:t.validateService}:{loggedIn:!1};case"LoggedIn":return{loggedIn:!0};default:return{loggedIn:!1}}}(e,t)}),(function(e,t){return{login:function(t,n){e(Object(N.createAction)(x.LOGIN_SUCCESS,{userId:t,token:Storage.auth,user:Storage.user,permissions:n}))},logout:function(t){e(t?Object(N.createAction)(x.VALIDATIONFAILED,null,null):Object(N.createAction)(x.LOGOUT,null,null))}}}))(_),G=n(11);n(13);function C(e,t){return function(n,r,a,i,c){return console.log("renderLogin",t,"uikit",_uikit,"settigs",e,"props",c),o.a.createElement("div",{className:c.className?c.className:" loginbox "},o.a.createElement("div",{className:"logintext"},t.loginForm.formtext),o.a.createElement("div",{className:"sociallogin"},o.a.createElement(G.Action,{widget:"button",method:function(){i(Application.Security.googleAuthUrl)},name:"googleAuth",className:"googleAuthAction"},t.loginForm.google)),o.a.createElement("div",{className:"separator"},t.loginForm.separator),o.a.createElement("div",{className:"main"},o.a.createElement(_uikit.Form,{role:"form"},o.a.createElement("div",{className:"userfield"},o.a.createElement("label",{htmlFor:"email"},t.loginForm.userlabel),o.a.createElement(_uikit.TextField,{className:"text",name:"email",value:n.email,placeholder:t.loginForm.userplaceholder,onChange:r})),o.a.createElement("div",{className:"passwordfield"},o.a.createElement("label",{htmlFor:"inputPassword"},t.loginForm.passwordlabel),o.a.createElement(_uikit.TextField,{type:"password",className:"text",name:"password",value:n.password,placeholder:t.loginForm.passwordplaceholder,onChange:r})),o.a.createElement("a",{className:"pull-right",href:"#"},"Forgot password?"),o.a.createElement("div",{className:"checkbox"},o.a.createElement("label",null,o.a.createElement("input",{type:"checkbox"}),"Remember me")),o.a.createElement("div",{className:"actionbuttons"},o.a.createElement(G.Action,{widget:"button",className:"loginBtn",name:"loginAction",method:a},t.loginForm.loginBtnText)))))}}function F(e,t){return function(n,r,a,i){return console.log("renderSignup",t,"uikit",_uikit,"settings",e,"props",i),o.a.createElement("div",{className:i.className?i.className:" signupbox "},o.a.createElement("div",{className:"signuptext"},t.signupForm.formtext),o.a.createElement("div",{className:"main"},o.a.createElement(_uikit.Form,{role:"form"},o.a.createElement("div",{className:"userfield"},o.a.createElement("label",{htmlFor:"email"},t.signupForm.userlabel),o.a.createElement(_uikit.TextField,{className:"text",name:"email",value:n.email,placeholder:t.signupForm.userplaceholder,onChange:r})),o.a.createElement("div",{className:"passwordfield"},o.a.createElement("label",{htmlFor:"inputPassword"},t.signupForm.passwordlabel),o.a.createElement(_uikit.TextField,{type:"password",className:"text",name:"password",value:n.password,placeholder:t.signupForm.passwordplaceholder,onChange:r})),o.a.createElement("div",{className:"confirmpasswordfield"},o.a.createElement("label",{htmlFor:"inputConfirmPassword"},t.signupForm.confirmpasswordlabel),o.a.createElement(_uikit.TextField,{type:"password",className:"text",name:"confirmpassword",value:n.confirmpassword,placeholder:t.signupForm.confirmpasswordplaceholder,onChange:r})),o.a.createElement("div",{className:"actionbuttons"},o.a.createElement(G.Action,{widget:"button",className:"signupBtn",name:"signupAction",method:a},t.signupForm.signupBtnText)))))}}var U=n(3),j=function(e){function t(e){var n;i()(this,t),n=l()(this,f()(t).call(this,e)),console.log("costructor of login web"),n.state={email:"",password:"",confirmpassword:""},n.handleSignup=n.handleSignup.bind(g()(n)),n.handleChange=n.handleChange.bind(g()(n));return e.realm&&"?Realm="+e.realm,n}return h()(t,e),u()(t,[{key:"handleChange",value:function(e){var t={};t[e.target.name]=e.target.value,this.setState(t)}},{key:"handleSignup",value:function(){this.props.handleSignup(this.state.email,this.state.password,this.state.confirmpassword)}},{key:"render",value:function(){return console.log("login ui",this.props),this.props.renderSignup(this.state,this.handleChange,this.handleSignup,this.props)}}]),t}(o.a.Component);j.propTypes={handleSignup:U.func.isRequired};var T=Object(w.connect)((function(e,t){return{realm:Application.Security.realm,renderSignup:t.renderSignup,signup:t.signup}}),(function(e,t){console.log("map dispatch of signup component");var n="";return Application.Security.realm&&(n=Application.Security.realm),{handleSignup:function(r,o,a){if(console.log("load",r,o,a),a==o){var i={Username:r,Password:b()(o),Realm:n};console.log(i);var c={serviceName:Application.Security.signupService};e(Object(N.createAction)(x.SIGN_UP,i,c))}else Window.showMessage(t.module.properties.errors.passwordmismatch)}}}))(j);T.propTypes={loginService:I.a.string.isRequired,successpage:I.a.string,realm:I.a.string,signup:I.a.string};var P={status:"NotLogged",token:"",userId:"",permissions:[]};Application.Register("Reducers","Security",(function(e,t){if(t.type)switch(t.type){case x.LOGGING_IN:return Object.assign({},e,{status:"LoggingIn"});case x.VALIDATIONFAILED:return Storage.auth="",Storage.permissions=[],Storage.userId="",Storage.userName="",Storage.userFullName="",Storage.email="",Storage.user=null,Object.assign({},P,{status:"ValidationFailed"});case x.LOGIN_SUCCESS:return e.userId===t.payload.userId?e:(Storage.auth=t.payload.token,Storage.permissions=t.payload.permissions,Storage.userId=t.payload.user.Id,Storage.userFullName=t.payload.user.Name,Storage.userName=t.payload.user.Username,Storage.email=t.payload.user.Email,Storage.user=t.payload.user,Object.assign({},e,{status:"LoggedIn",authToken:t.payload.token,userId:t.payload.userId,permissions:t.payload.permissions}));case x.LOGIN_FAILURE:case x.LOGOUT_SUCCESS:return Storage.auth="",Storage.permissions=[],Storage.userId="",Storage.userName="",Storage.userFullName="",Storage.email="",Storage.user=null,P;default:return e||P}}));var R=n(2),q=n.n(R),D=function(e){return"@@redux-saga/"+e},B=D("IO"),M=D("MULTICAST"),V=D("SELF_CANCELLATION");var W=function(e){return null==e},z=function(e){return null!=e},K=function(e){return"function"==typeof e},Y=function(e){return"string"==typeof e},X=Array.isArray,H=function e(t){return t&&(Y(t)||Z(t)||K(t)||X(t)&&t.every(e))},J=function(e){return e&&K(e.take)&&K(e.close)},Q=function(e){return K(e)&&e.hasOwnProperty("toString")},Z=function(e){return Boolean(e)&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype},$=function(e){return J(e)&&e[M]};"function"==typeof Symbol&&Symbol.asyncIterator&&Symbol.asyncIterator;var ee=function(e){throw e},te=function(e){return{value:e,done:!0}};var ne="TAKE",re="PUT",oe="CALL",ae="FORK",ie="CANCEL",ce=function(e,t){var n;return(n={})[B]=!0,n.combinator=!1,n.type=e,n.payload=t,n};function ue(e,t){return void 0===e&&(e="*"),H(e)?ce(ne,{pattern:e}):$(e)&&z(t)&&H(t)?ce(ne,{channel:e,pattern:t}):J(e)?ce(ne,{channel:e}):void 0}function se(e,t){return W(t)&&(t=e,e=void 0),ce(re,{channel:e,action:t})}function le(e,t){var n,r=null;return K(e)?n=e:(X(e)?(r=e[0],n=e[1]):(r=e.context,n=e.fn),r&&Y(n)&&K(r[n])&&(n=r[n])),{context:r,fn:n,args:t}}function pe(e){for(var t=arguments.length,n=new Array(t>1?t-1:0),r=1;r<t;r++)n[r-1]=arguments[r];return ce(oe,le(e,n))}function fe(e){for(var t=arguments.length,n=new Array(t>1?t-1:0),r=1;r<t;r++)n[r-1]=arguments[r];return ce(ae,le(e,n))}var de=function(e){return{done:!0,value:e}},ge={};function me(e){return J(e)?"channel":Q(e)?String(e):K(e)?e.name:String(e)}function he(e,t,n){var r,o,a,i=t;function c(t,n){if(i===ge)return de(t);if(n&&!o)throw i=ge,n;r&&r(t);var c=n?e[o](n):e[i]();return i=c.nextState,a=c.effect,r=c.stateUpdater,o=c.errorState,i===ge?de(t):a}return function(e,t,n){void 0===t&&(t=ee),void 0===n&&(n="iterator");var r={meta:{name:n},next:e,throw:t,return:te,isSagaIterator:!0};return"undefined"!=typeof Symbol&&(r[Symbol.iterator]=function(){return r}),r}(c,(function(e){return c(null,e)}),n)}function ve(e,t){for(var n=arguments.length,r=new Array(n>2?n-2:0),o=2;o<n;o++)r[o-2]=arguments[o];var a,i,c={done:!1,value:ue(e)},u=function(e){return{done:!1,value:fe.apply(void 0,[t].concat(r,[e]))}},s=function(e){return{done:!1,value:(t=e,void 0===t&&(t=V),ce(ie,t))};var t},l=function(e){return a=e},p=function(e){return i=e};return he({q1:function(){return{nextState:"q2",effect:c,stateUpdater:p}},q2:function(){return a?{nextState:"q3",effect:s(a)}:{nextState:"q1",effect:u(i),stateUpdater:l}},q3:function(){return{nextState:"q1",effect:u(i),stateUpdater:l}}},"q1","takeLatest("+me(e)+", "+t.name+")")}function ye(e,t){for(var n=arguments.length,r=new Array(n>2?n-2:0),o=2;o<n;o++)r[o-2]=arguments[o];return fe.apply(void 0,[ve,e,t].concat(r))}var Se=q.a.mark(Ee),be=q.a.mark(Oe),we=q.a.mark(Ae),xe=q.a.mark(_e),Ne=q.a.mark(ke),Le=q.a.mark(Ge),Ie=q.a.mark(Ce);function Ee(e){var t,n;return q.a.wrap((function(r){for(;;)switch(r.prev=r.next){case 0:return r.prev=0,r.next=3,se(Object(N.createAction)(x.SIGNING_UP));case 3:return t=N.RequestBuilder.DefaultRequest(null,e.payload),r.next=6,pe(N.DataSource.ExecuteService,e.meta.serviceName,t);case 6:return r.sent,n=Object(N.createAction)(x.SIGNUP_SUCCESS,{}),r.next=10,se(n);case 10:console.log("dispatched signup action success"),r.next=18;break;case 13:return r.prev=13,r.t0=r.catch(0),r.next=17,se(Object(N.createAction)(x.SIGNUP_FAILURE,r.t0));case 17:Window.handleError(r.t0);case 18:case"end":return r.stop()}}),Se,null,[[0,13]])}function Oe(e){var t,n,r,o,a,i,c,u;return q.a.wrap((function(s){for(;;)switch(s.prev=s.next){case 0:return s.prev=0,console.log("received login action",e),s.next=4,se(Object(N.createAction)(x.LOGGING_IN));case 4:return t=N.RequestBuilder.DefaultRequest(null,e.payload),s.next=7,pe(N.DataSource.ExecuteService,e.meta.serviceName,t);case 7:return n=s.sent,r=Application.Security.AuthToken.toLowerCase(),o=n.info[r],a=n.data,i=n.data.Id,c=n.data.Permissions,u=Object(N.createAction)(x.LOGIN_SUCCESS,{userId:i,token:o,permissions:c,user:a}),s.next=16,se(u);case 16:console.log("dispatched login action &&&&"),s.next=24;break;case 19:return s.prev=19,s.t0=s.catch(0),s.next=23,se(Object(N.createAction)(x.LOGIN_FAILURE,s.t0));case 23:Window.handleError(s.t0);case 24:case"end":return s.stop()}}),be,null,[[0,19]])}function Ae(e){return q.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,se(Object(N.createAction)(x.LOGOUT_SUCCESS,{}));case 2:case"end":return e.stop()}}),we)}function _e(){return q.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,ye(x.LOGIN,Oe);case 2:case"end":return e.stop()}}),xe)}function ke(){return q.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,ye(x.SIGN_UP,Ee);case 2:case"end":return e.stop()}}),Ne)}function Ge(){return q.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,ye(x.LOGOUT,Ae);case 2:case"end":return e.stop()}}),Le)}function Ce(){return q.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return console.log("take latest in auth saga",fe),e.next=3,fe(_e);case 3:return e.next=5,fe(ke);case 5:return e.next=7,fe(Ge);case 7:case"end":return e.stop()}}),Ie)}Application.Register("Sagas","authSaga",Ce);var Fe=function(e){function t(){return i()(this,t),l()(this,f()(t).apply(this,arguments))}return h()(t,e),u()(t,[{key:"render",value:function(){var e=this.props,t=e.module.properties?e.module.properties:{},n=t.logoutText?t.logoutText:"Logout";return e.loggedIn?o.a.createElement(_uikit.Block,{className:"userblock "+e.className},o.a.createElement(_uikit.Block,{className:"username"},Storage.userFullName?Storage.userFullName:Storage.userName),o.a.createElement(G.Action,{name:"logout",method:e.logout,className:"logout"},n)):null}}]),t}(o.a.Component),Ue=Object(w.connect)((function(e,t){return{loggedIn:"LoggedIn"==e.Security.status}}),(function(e,t){return{logout:function(){e(Object(N.createAction)(x.LOGOUT,null,null))}}}))(Fe);n.d(t,"Initialize",(function(){return Te})),n.d(t,"SignupForm",(function(){return Re})),n.d(t,"WebLoginForm",(function(){return Pe})),n.d(t,"LoginComponent",(function(){return E})),n.d(t,"renderWebLogin",(function(){return C})),n.d(t,"renderSignup",(function(){return F})),n.d(t,"SignupComponent",(function(){return T})),n.d(t,"LoginValidator",(function(){return k}));var je;n(3);function Te(e,t,n,r,o,a){(je=this).properties=Application.Properties[t],console.log("authui initialization",Application,t),je.settings=r,function(e){O=e}(je),Application.Security=Object.assign({loginService:"login",signupService:"signup",validateService:"validate",AuthToken:"X-Auth-Token",realm:""},r)}var Pe=function(e,t){return console.log("render logiform",E),o.a.createElement(E,{className:e.className,renderLogin:C(je.settings,je.properties),loginService:e.loginService,googleAuthUrl:e.googleAuthUrl})},Re=function(e,t){return console.log("render signup form",T),o.a.createElement(T,{className:e.className,renderSignup:F(je.settings,je.properties),module:je})};Application.Register("Blocks","userBlock",(function(e,t,n){return o.a.createElement(Ue,{className:e.className,module:je})}))}])}));
//# sourceMappingURL=index.js.map
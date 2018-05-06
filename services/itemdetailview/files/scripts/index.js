define("itemdetailview",["react","reactpages","reactwebcommon","laatooviews","prop-types"],function(t,e,n,r,o){return function(t){var e={};function n(r){if(e[r])return e[r].exports;var o=e[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=t,n.c=e,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{configurable:!1,enumerable:!0,get:r})},n.r=function(t){Object.defineProperty(t,"__esModule",{value:!0})},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=71)}([function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(e,n){e.exports=t},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,e,n){var r=n(19)("wks"),o=n(18),i=n(7).Symbol;t.exports=function(t){return r[t]||(r[t]=i&&i[t]||(i||o)("Symbol."+t))}},function(t,e,n){var r=n(7),o=n(2),i=n(28),u=function(t,e,n){var c,a,s,f=t&u.F,l=t&u.G,p=t&u.S,d=t&u.P,y=t&u.B,v=t&u.W,h=l?o:o[e]||(o[e]={}),m=l?r:p?r[e]:(r[e]||{}).prototype;for(c in l&&(n=e),n)(a=!f&&m&&c in m)&&c in h||(s=a?m[c]:n[c],h[c]=l&&"function"!=typeof m[c]?n[c]:y&&a?i(s,r):v&&m[c]==s?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(s):d&&"function"==typeof s?i(Function.call,s):s,d&&((h.prototype||(h.prototype={}))[c]=s))};u.F=1,u.G=2,u.S=4,u.P=8,u.B=16,u.W=32,t.exports=u},function(t,e,n){var r=n(26),o=n(13);t.exports=function(t){return r(o(t))}},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e,n){var r=n(0).setDesc,o=n(10),i=n(3)("toStringTag");t.exports=function(t,e,n){t&&!o(t=n?t:t.prototype,i)&&r(t,i,{configurable:!0,value:e})}},function(t,e){t.exports={}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e,n){var r=n(0),o=n(11);t.exports=n(20)?function(t,e,n){return r.setDesc(t,e,o(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,n){t.exports=e},function(t,e,n){t.exports={default:n(70),__esModule:!0}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){var r=n(16);t.exports=function(t){if(!r(t))throw TypeError(t+" is not an object!");return t}},function(t,e){var n=0,r=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+r).toString(36))}},function(t,e,n){var r=n(7),o=r["__core-js_shared__"]||(r["__core-js_shared__"]={});t.exports=function(t){return o[t]||(o[t]={})}},function(t,e,n){t.exports=!n(6)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){t.exports=n(12)},function(t,e){t.exports=!0},function(t,e,n){"use strict";var r=n(22),o=n(4),i=n(21),u=n(12),c=n(10),a=n(9),s=n(56),f=n(8),l=n(0).getProto,p=n(3)("iterator"),d=!([].keys&&"next"in[].keys()),y=function(){return this};t.exports=function(t,e,n,v,h,m,g){s(n,e,v);var b,_,x=function(t){if(!d&&t in j)return j[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},w=e+" Iterator",O="values"==h,S=!1,j=t.prototype,P=j[p]||j["@@iterator"]||h&&j[h],M=P||x(h);if(P){var N=l(M.call(new t));f(N,w,!0),!r&&c(j,"@@iterator")&&u(N,p,y),O&&"values"!==P.name&&(S=!0,M=function(){return P.call(this)})}if(r&&!g||!d&&!S&&j[p]||u(j,p,M),a[e]=M,a[w]=y,h)if(b={values:O?M:x("values"),keys:m?M:x("keys"),entries:O?x("entries"):M},g)for(_ in b)_ in j||i(j,_,b[_]);else o(o.P+o.F*(d||S),e,b);return b}},function(t,e,n){"use strict";e.__esModule=!0;var r=u(n(61)),o=u(n(51)),i="function"==typeof o.default&&"symbol"==typeof r.default?function(t){return typeof t}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":typeof t};function u(t){return t&&t.__esModule?t:{default:t}}e.default="function"==typeof o.default&&"symbol"===i(r.default)?function(t){return void 0===t?"undefined":i(t)}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":void 0===t?"undefined":i(t)}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){var r=n(25);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==r(t)?t.split(""):Object(t)}},function(t,e,n){var r=n(13);t.exports=function(t){return Object(r(t))}},function(t,e,n){var r=n(68);t.exports=function(t,e,n){if(r(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,r){return t.call(e,n,r)};case 3:return function(n,r,o){return t.call(e,n,r,o)}}return function(){return t.apply(e,arguments)}}},function(t,e){t.exports=n},function(t,e){t.exports=r},function(t,e,n){"use strict";e.__esModule=!0;var r=u(n(43)),o=u(n(39)),i=u(n(24));function u(t){return t&&t.__esModule?t:{default:t}}e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,i.default)(e)));t.prototype=(0,o.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(r.default?(0,r.default)(t,e):t.__proto__=e)}},function(t,e,n){"use strict";e.__esModule=!0;var r,o=n(24),i=(r=o)&&r.__esModule?r:{default:r};e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,i.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){"use strict";e.__esModule=!0;var r,o=n(63),i=(r=o)&&r.__esModule?r:{default:r};e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var r=e[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),(0,i.default)(t,r.key,r)}}return function(e,n,r){return n&&t(e.prototype,n),r&&t(e,r),e}}()},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){t.exports={default:n(66),__esModule:!0}},function(t,e){},function(t,e){t.exports=o},function(t,e,n){var r=n(0);t.exports=function(t,e){return r.create(t,e)}},function(t,e,n){t.exports={default:n(38),__esModule:!0}},function(t,e,n){var r=n(0).getDesc,o=n(16),i=n(17),u=function(t,e){if(i(t),!o(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,o){try{(o=n(28)(Function.call,r(Object.prototype,"__proto__").set,2))(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return u(t,n),e?t.__proto__=n:o(t,n),t}}({},!1):void 0),check:u}},function(t,e,n){var r=n(4);r(r.S,"Object",{setPrototypeOf:n(40).set})},function(t,e,n){n(41),t.exports=n(2).Object.setPrototypeOf},function(t,e,n){t.exports={default:n(42),__esModule:!0}},function(t,e){},function(t,e,n){var r=n(25);t.exports=Array.isArray||function(t){return"Array"==r(t)}},function(t,e,n){var r=n(0);t.exports=function(t){var e=r.getKeys(t),n=r.getSymbols;if(n)for(var o,i=n(t),u=r.isEnum,c=0;i.length>c;)u.call(t,o=i[c++])&&e.push(o);return e}},function(t,e,n){var r=n(5),o=n(0).getNames,i={}.toString,u="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[];t.exports.get=function(t){return u&&"[object Window]"==i.call(t)?function(t){try{return o(t)}catch(t){return u.slice()}}(t):o(r(t))}},function(t,e,n){var r=n(0),o=n(5);t.exports=function(t,e){for(var n,i=o(t),u=r.getKeys(i),c=u.length,a=0;c>a;)if(i[n=u[a++]]===e)return n}},function(t,e,n){"use strict";var r=n(0),o=n(7),i=n(10),u=n(20),c=n(4),a=n(21),s=n(6),f=n(19),l=n(8),p=n(18),d=n(3),y=n(48),v=n(47),h=n(46),m=n(45),g=n(17),b=n(5),_=n(11),x=r.getDesc,w=r.setDesc,O=r.create,S=v.get,j=o.Symbol,P=o.JSON,M=P&&P.stringify,N=!1,k=d("_hidden"),E=r.isEnum,D=f("symbol-registry"),A=f("symbols"),I="function"==typeof j,F=Object.prototype,T=u&&s(function(){return 7!=O(w({},"a",{get:function(){return w(this,"a",{value:7}).a}})).a})?function(t,e,n){var r=x(F,e);r&&delete F[e],w(t,e,n),r&&t!==F&&w(F,e,r)}:w,C=function(t){var e=A[t]=O(j.prototype);return e._k=t,u&&N&&T(F,t,{configurable:!0,set:function(e){i(this,k)&&i(this[k],t)&&(this[k][t]=!1),T(this,t,_(1,e))}}),e},R=function(t){return"symbol"==typeof t},B=function(t,e,n){return n&&i(A,e)?(n.enumerable?(i(t,k)&&t[k][e]&&(t[k][e]=!1),n=O(n,{enumerable:_(0,!1)})):(i(t,k)||w(t,k,_(1,{})),t[k][e]=!0),T(t,e,n)):w(t,e,n)},J=function(t,e){g(t);for(var n,r=h(e=b(e)),o=0,i=r.length;i>o;)B(t,n=r[o++],e[n]);return t},K=function(t,e){return void 0===e?O(t):J(O(t),e)},L=function(t){var e=E.call(this,t);return!(e||!i(this,t)||!i(A,t)||i(this,k)&&this[k][t])||e},V=function(t,e){var n=x(t=b(t),e);return!n||!i(A,e)||i(t,k)&&t[k][e]||(n.enumerable=!0),n},W=function(t){for(var e,n=S(b(t)),r=[],o=0;n.length>o;)i(A,e=n[o++])||e==k||r.push(e);return r},G=function(t){for(var e,n=S(b(t)),r=[],o=0;n.length>o;)i(A,e=n[o++])&&r.push(A[e]);return r},H=s(function(){var t=j();return"[null]"!=M([t])||"{}"!=M({a:t})||"{}"!=M(Object(t))});I||(a((j=function(){if(R(this))throw TypeError("Symbol is not a constructor");return C(p(arguments.length>0?arguments[0]:void 0))}).prototype,"toString",function(){return this._k}),R=function(t){return t instanceof j},r.create=K,r.isEnum=L,r.getDesc=V,r.setDesc=B,r.setDescs=J,r.getNames=v.get=W,r.getSymbols=G,u&&!n(22)&&a(F,"propertyIsEnumerable",L,!0));var q={for:function(t){return i(D,t+="")?D[t]:D[t]=j(t)},keyFor:function(t){return y(D,t)},useSetter:function(){N=!0},useSimple:function(){N=!1}};r.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=d(t);q[t]=I?e:C(e)}),N=!0,c(c.G+c.W,{Symbol:j}),c(c.S,"Symbol",q),c(c.S+c.F*!I,"Object",{create:K,defineProperty:B,defineProperties:J,getOwnPropertyDescriptor:V,getOwnPropertyNames:W,getOwnPropertySymbols:G}),P&&c(c.S+c.F*(!I||H),"JSON",{stringify:function(t){if(void 0!==t&&!R(t)){for(var e,n,r=[t],o=1,i=arguments;i.length>o;)r.push(i[o++]);return"function"==typeof(e=r[1])&&(n=e),!n&&m(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!R(e))return e}),r[1]=e,M.apply(P,r)}}}),l(j,"Symbol"),l(Math,"Math",!0),l(o.JSON,"JSON",!0)},function(t,e,n){n(49),n(44),t.exports=n(2).Symbol},function(t,e,n){t.exports={default:n(50),__esModule:!0}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e){t.exports=function(){}},function(t,e,n){"use strict";var r=n(53),o=n(52),i=n(9),u=n(5);t.exports=n(23)(Array,"Array",function(t,e){this._t=u(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,o(1)):o(0,"keys"==e?n:"values"==e?t[n]:[n,t[n]])},"values"),i.Arguments=i.Array,r("keys"),r("values"),r("entries")},function(t,e,n){n(54);var r=n(9);r.NodeList=r.HTMLCollection=r.Array},function(t,e,n){"use strict";var r=n(0),o=n(11),i=n(8),u={};n(12)(u,n(3)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=r.create(u,{next:o(1,n)}),i(t,e+" Iterator")}},function(t,e){var n=Math.ceil,r=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?r:n)(t)}},function(t,e,n){var r=n(57),o=n(13);t.exports=function(t){return function(e,n){var i,u,c=String(o(e)),a=r(n),s=c.length;return a<0||a>=s?t?"":void 0:(i=c.charCodeAt(a))<55296||i>56319||a+1===s||(u=c.charCodeAt(a+1))<56320||u>57343?t?c.charAt(a):i:t?c.slice(a,a+2):u-56320+(i-55296<<10)+65536}}},function(t,e,n){"use strict";var r=n(58)(!0);n(23)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=r(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){n(59),n(55),t.exports=n(3)("iterator")},function(t,e,n){t.exports={default:n(60),__esModule:!0}},function(t,e,n){var r=n(0);t.exports=function(t,e,n){return r.setDesc(t,e,n)}},function(t,e,n){t.exports={default:n(62),__esModule:!0}},function(t,e,n){var r=n(4),o=n(2),i=n(6);t.exports=function(t,e){var n=(o.Object||{})[t]||Object[t],u={};u[t]=e(n),r(r.S+r.F*i(function(){n(1)}),"Object",u)}},function(t,e,n){var r=n(27);n(64)("getPrototypeOf",function(t){return function(e){return t(r(e))}})},function(t,e,n){n(65),t.exports=n(2).Object.getPrototypeOf},function(t,e,n){var r=n(0),o=n(27),i=n(26);t.exports=n(6)(function(){var t=Object.assign,e={},n={},r=Symbol(),o="abcdefghijklmnopqrst";return e[r]=7,o.split("").forEach(function(t){n[t]=t}),7!=t({},e)[r]||Object.keys(t({},n)).join("")!=o})?function(t,e){for(var n=o(t),u=arguments,c=u.length,a=1,s=r.getKeys,f=r.getSymbols,l=r.isEnum;c>a;)for(var p,d=i(u[a++]),y=f?s(d).concat(f(d)):s(d),v=y.length,h=0;v>h;)l.call(d,p=y[h++])&&(n[p]=d[p]);return n}:Object.assign},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){var r=n(4);r(r.S+r.F,"Object",{assign:n(67)})},function(t,e,n){n(69),t.exports=n(2).Object.assign},function(t,e,n){"use strict";n.r(e),n.d(e,"ItemDetailView",function(){return x});var r=n(15),o=n.n(r),i=n(35),u=n.n(i),c=n(34),a=n.n(c),s=n(33),f=n.n(s),l=n(32),p=n.n(l),d=n(31),y=n.n(d),v=n(1),h=n.n(v),m=n(14),g=n(30),b=n(29),_=(n(36),n(37)),x=function(t){function e(t){a()(this,e);var n=p()(this,(e.__proto__||u()(e)).call(this,t));return n.openDetail=function(t){console.log("open detail",t);t.data;n.setState(o()({},n.state,{open:!0,item:t.data}))},n.hideDetail=function(){n.setState(o()({},n.state,{open:!1,item:null}))},n.getItem=function(t,e,r){var o=t.getRenderedItem(e,r);return n.props.getItem?(t.addMethod("openDetail",n.openDetail),t.addMethod("hideDetail",n.hideDetail),n.props.getItem(t,e,r)):h.a.createElement(b.Action,{action:{actiontype:"method",method:n.openDetail,params:{data:e,index:r}}},o)},console.log("creating item detail view",t),n.state={open:!1,item:null},n}return y()(e,t),f()(e,[{key:"render",value:function(){console.log("rendering item detail view",this.props,this.context);var t=this.props,e=this.context,n=this.state.open?" col-xs-6 ":"";return h.a.createElement(e.uikit.Block,{className:" itemdetailview ma10  row "},h.a.createElement(e.uikit.Block,{className:" view "+n},t.id?h.a.createElement(m.Panel,{className:t.className,editable:t.editable,getItem:this.getItem,viewRef:t.viewRef,type:"view",id:t.id}):h.a.createElement(g.View,{service:t.service,serviceName:t.serviceName,name:t.viewname,global:t.global,editable:t.editable,className:t.className,incrementalLoad:t.incrementalLoad,paginate:t.paginate,header:t.header,viewRef:t.viewRef,getHeader:t.getHeader,getView:t.getView,getItem:this.getItem,urlparams:t.urlparams,postArgs:t.postArgs})),h.a.createElement(e.uikit.Block,{className:" col-xs-6 "},this.state.open&&t.entityName?h.a.createElement(m.Panel,{className:"detail",title:this.state.item.Name,closePanel:this.hideDetail,description:{type:"entity",entityName:t.entityName,data:this.state.item}}):null))}}]),e}(h.a.Component);x.contextTypes={uikit:_.object}}])});
//# sourceMappingURL=index.js.map
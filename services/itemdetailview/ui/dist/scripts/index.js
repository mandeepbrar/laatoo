define("itemdetailview",["react","reactpages","laatooviews","reactwebcommon","prop-types"],function(t,e,n,r,o){return function(t){function e(r){if(n[r])return n[r].exports;var o=n[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,e),o.l=!0,o.exports}var n={};return e.m=t,e.c=n,e.d=function(t,n,r){e.o(t,n)||Object.defineProperty(t,n,{configurable:!1,enumerable:!0,get:r})},e.n=function(t){var n=t&&t.__esModule?function(){return t.default}:function(){return t};return e.d(n,"a",n),n},e.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},e.p="/",e(e.s=26)}([function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,e,n){var r=n(4),o=n(1),i=n(13),u=function(t,e,n){var c,a,s,f=t&u.F,l=t&u.G,p=t&u.S,d=t&u.P,y=t&u.B,v=t&u.W,h=l?o:o[e]||(o[e]={}),m=l?r:p?r[e]:(r[e]||{}).prototype;l&&(n=e);for(c in n)(a=!f&&m&&c in m)&&c in h||(s=a?m[c]:n[c],h[c]=l&&"function"!=typeof m[c]?n[c]:y&&a?i(s,r):v&&m[c]==s?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(s):d&&"function"==typeof s?i(Function.call,s):s,d&&((h.prototype||(h.prototype={}))[c]=s))};u.F=1,u.G=2,u.S=4,u.P=8,u.B=16,u.W=32,t.exports=u},function(t,e,n){var r=n(22)("wks"),o=n(23),i=n(4).Symbol;t.exports=function(t){return r[t]||(r[t]=i&&i[t]||(i||o)("Symbol."+t))}},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e,n){var r=n(15),o=n(7);t.exports=function(t){return r(o(t))}},function(t,e){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e,n){var r=n(0),o=n(9);t.exports=n(21)?function(t,e,n){return r.setDesc(t,e,o(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports={}},function(t,e,n){var r=n(0).setDesc,o=n(10),i=n(3)("toStringTag");t.exports=function(t,e,n){t&&!o(t=n?t:t.prototype,i)&&r(t,i,{configurable:!0,value:e})}},function(t,e,n){var r=n(30);t.exports=function(t,e,n){if(r(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,r){return t.call(e,n,r)};case 3:return function(n,r,o){return t.call(e,n,r,o)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){var r=n(7);t.exports=function(t){return Object(r(t))}},function(t,e,n){var r=n(16);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==r(t)?t.split(""):Object(t)}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){"use strict";function r(t){return t&&t.__esModule?t:{default:t}}e.__esModule=!0;var o=n(41),i=r(o),u=n(51),c=r(u),a="function"==typeof c.default&&"symbol"==typeof i.default?function(t){return typeof t}:function(t){return t&&"function"==typeof c.default&&t.constructor===c.default&&t!==c.default.prototype?"symbol":typeof t};e.default="function"==typeof c.default&&"symbol"===a(i.default)?function(t){return void 0===t?"undefined":a(t)}:function(t){return t&&"function"==typeof c.default&&t.constructor===c.default&&t!==c.default.prototype?"symbol":void 0===t?"undefined":a(t)}},function(t,e,n){"use strict";var r=n(19),o=n(2),i=n(20),u=n(8),c=n(10),a=n(11),s=n(46),f=n(12),l=n(0).getProto,p=n(3)("iterator"),d=!([].keys&&"next"in[].keys()),y=function(){return this};t.exports=function(t,e,n,v,h,m,g){s(n,e,v);var _,b,x=function(t){if(!d&&t in j)return j[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},w=e+" Iterator",O="values"==h,S=!1,j=t.prototype,P=j[p]||j["@@iterator"]||h&&j[h],M=P||x(h);if(P){var N=l(M.call(new t));f(N,w,!0),!r&&c(j,"@@iterator")&&u(N,p,y),O&&"values"!==P.name&&(S=!0,M=function(){return P.call(this)})}if(r&&!g||!d&&!S&&j[p]||u(j,p,M),a[e]=M,a[w]=y,h)if(_={values:O?M:x("values"),keys:m?M:x("keys"),entries:O?x("entries"):M},g)for(b in _)b in j||i(j,b,_[b]);else o(o.P+o.F*(d||S),e,_);return _}},function(t,e){t.exports=!0},function(t,e,n){t.exports=n(8)},function(t,e,n){t.exports=!n(5)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){var r=n(4),o=r["__core-js_shared__"]||(r["__core-js_shared__"]={});t.exports=function(t){return o[t]||(o[t]={})}},function(t,e){var n=0,r=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+r).toString(36))}},function(t,e,n){var r=n(25);t.exports=function(t){if(!r(t))throw TypeError(t+" is not an object!");return t}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),n.d(e,"ItemDetailView",function(){return x});var r=n(27),o=n.n(r),i=n(32),u=n.n(i),c=n(36),a=n.n(c),s=n(37),f=n.n(s),l=n(40),p=n.n(l),d=n(59),y=n.n(d),v=n(66),h=n.n(v),m=n(67),g=(n.n(m),n(68)),_=(n.n(g),n(69)),b=(n.n(_),n(70)),x=function(t){function e(t){a()(this,e);var n=p()(this,(e.__proto__||u()(e)).call(this,t));return n.openDetail=function(t){console.log("open detail",t),n.setState(o()({},n.state,{open:!0,item:t}))},n.closeDetail=function(){n.setState(o()({},n.state,{open:!1}))},n.getItem=function(t,e,r){var o=t.getRenderedItem(e,r);console.log("getitem",o,n.context.uikit.Action);var i=h.a.createElement(_.Action,{action:{actiontype:"method",method:n.openDetail,params:{data:e,index:r}}},o);return console.log("created actio",i,o,e,r),i},n.state={open:!1,item:null},n}return y()(e,t),f()(e,[{key:"render",value:function(){console.log("rendering itemdetail view");var t=this.props,e=this.context;return h.a.createElement(e.uikit.Block,{className:"itemdetailview"},h.a.createElement(e.uikit.Block,{className:"view"},t.id?h.a.createElement(m.Panel,{className:t.className,getItem:this.getItem,type:"view",id:t.id}):h.a.createElement(g.View,{service:t.service,serviceName:t.serviceName,name:t.viewname,global:t.global,className:t.className,incrementalLoad:t.incrementalLoad,paginate:t.paginate,header:t.header,getHeader:t.getHeader,getView:t.getView,getItem:this.getItem,urlparams:t.urlparams,postArgs:t.postArgs})),this.state.open&&t.entityName?h.a.createElement(m.Panel,{className:"detail",description:{type:"entity",entityName:t.entityName,data:this.state.item}}):null)}}]),e}(h.a.Component);x.contextTypes={uikit:b.object}},function(t,e,n){t.exports={default:n(28),__esModule:!0}},function(t,e,n){n(29),t.exports=n(1).Object.assign},function(t,e,n){var r=n(2);r(r.S+r.F,"Object",{assign:n(31)})},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){var r=n(0),o=n(14),i=n(15);t.exports=n(5)(function(){var t=Object.assign,e={},n={},r=Symbol(),o="abcdefghijklmnopqrst";return e[r]=7,o.split("").forEach(function(t){n[t]=t}),7!=t({},e)[r]||Object.keys(t({},n)).join("")!=o})?function(t,e){for(var n=o(t),u=arguments,c=u.length,a=1,s=r.getKeys,f=r.getSymbols,l=r.isEnum;c>a;)for(var p,d=i(u[a++]),y=f?s(d).concat(f(d)):s(d),v=y.length,h=0;v>h;)l.call(d,p=y[h++])&&(n[p]=d[p]);return n}:Object.assign},function(t,e,n){t.exports={default:n(33),__esModule:!0}},function(t,e,n){n(34),t.exports=n(1).Object.getPrototypeOf},function(t,e,n){var r=n(14);n(35)("getPrototypeOf",function(t){return function(e){return t(r(e))}})},function(t,e,n){var r=n(2),o=n(1),i=n(5);t.exports=function(t,e){var n=(o.Object||{})[t]||Object[t],u={};u[t]=e(n),r(r.S+r.F*i(function(){n(1)}),"Object",u)}},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){"use strict";e.__esModule=!0;var r=n(38),o=function(t){return t&&t.__esModule?t:{default:t}}(r);e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var r=e[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),(0,o.default)(t,r.key,r)}}return function(e,n,r){return n&&t(e.prototype,n),r&&t(e,r),e}}()},function(t,e,n){t.exports={default:n(39),__esModule:!0}},function(t,e,n){var r=n(0);t.exports=function(t,e,n){return r.setDesc(t,e,n)}},function(t,e,n){"use strict";e.__esModule=!0;var r=n(17),o=function(t){return t&&t.__esModule?t:{default:t}}(r);e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,o.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){t.exports={default:n(42),__esModule:!0}},function(t,e,n){n(43),n(47),t.exports=n(3)("iterator")},function(t,e,n){"use strict";var r=n(44)(!0);n(18)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=r(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){var r=n(45),o=n(7);t.exports=function(t){return function(e,n){var i,u,c=String(o(e)),a=r(n),s=c.length;return a<0||a>=s?t?"":void 0:(i=c.charCodeAt(a),i<55296||i>56319||a+1===s||(u=c.charCodeAt(a+1))<56320||u>57343?t?c.charAt(a):i:t?c.slice(a,a+2):u-56320+(i-55296<<10)+65536)}}},function(t,e){var n=Math.ceil,r=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?r:n)(t)}},function(t,e,n){"use strict";var r=n(0),o=n(9),i=n(12),u={};n(8)(u,n(3)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=r.create(u,{next:o(1,n)}),i(t,e+" Iterator")}},function(t,e,n){n(48);var r=n(11);r.NodeList=r.HTMLCollection=r.Array},function(t,e,n){"use strict";var r=n(49),o=n(50),i=n(11),u=n(6);t.exports=n(18)(Array,"Array",function(t,e){this._t=u(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,o(1)):"keys"==e?o(0,n):"values"==e?o(0,t[n]):o(0,[n,t[n]])},"values"),i.Arguments=i.Array,r("keys"),r("values"),r("entries")},function(t,e){t.exports=function(){}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e,n){t.exports={default:n(52),__esModule:!0}},function(t,e,n){n(53),n(58),t.exports=n(1).Symbol},function(t,e,n){"use strict";var r=n(0),o=n(4),i=n(10),u=n(21),c=n(2),a=n(20),s=n(5),f=n(22),l=n(12),p=n(23),d=n(3),y=n(54),v=n(55),h=n(56),m=n(57),g=n(24),_=n(6),b=n(9),x=r.getDesc,w=r.setDesc,O=r.create,S=v.get,j=o.Symbol,P=o.JSON,M=P&&P.stringify,N=!1,k=d("_hidden"),E=r.isEnum,D=f("symbol-registry"),A=f("symbols"),I="function"==typeof j,F=Object.prototype,T=u&&s(function(){return 7!=O(w({},"a",{get:function(){return w(this,"a",{value:7}).a}})).a})?function(t,e,n){var r=x(F,e);r&&delete F[e],w(t,e,n),r&&t!==F&&w(F,e,r)}:w,C=function(t){var e=A[t]=O(j.prototype);return e._k=t,u&&N&&T(F,t,{configurable:!0,set:function(e){i(this,k)&&i(this[k],t)&&(this[k][t]=!1),T(this,t,b(1,e))}}),e},B=function(t){return"symbol"==typeof t},J=function(t,e,n){return n&&i(A,e)?(n.enumerable?(i(t,k)&&t[k][e]&&(t[k][e]=!1),n=O(n,{enumerable:b(0,!1)})):(i(t,k)||w(t,k,b(1,{})),t[k][e]=!0),T(t,e,n)):w(t,e,n)},K=function(t,e){g(t);for(var n,r=h(e=_(e)),o=0,i=r.length;i>o;)J(t,n=r[o++],e[n]);return t},L=function(t,e){return void 0===e?O(t):K(O(t),e)},V=function(t){var e=E.call(this,t);return!(e||!i(this,t)||!i(A,t)||i(this,k)&&this[k][t])||e},W=function(t,e){var n=x(t=_(t),e);return!n||!i(A,e)||i(t,k)&&t[k][e]||(n.enumerable=!0),n},G=function(t){for(var e,n=S(_(t)),r=[],o=0;n.length>o;)i(A,e=n[o++])||e==k||r.push(e);return r},H=function(t){for(var e,n=S(_(t)),r=[],o=0;n.length>o;)i(A,e=n[o++])&&r.push(A[e]);return r},R=function(t){if(void 0!==t&&!B(t)){for(var e,n,r=[t],o=1,i=arguments;i.length>o;)r.push(i[o++]);return e=r[1],"function"==typeof e&&(n=e),!n&&m(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!B(e))return e}),r[1]=e,M.apply(P,r)}},q=s(function(){var t=j();return"[null]"!=M([t])||"{}"!=M({a:t})||"{}"!=M(Object(t))});I||(j=function(){if(B(this))throw TypeError("Symbol is not a constructor");return C(p(arguments.length>0?arguments[0]:void 0))},a(j.prototype,"toString",function(){return this._k}),B=function(t){return t instanceof j},r.create=L,r.isEnum=V,r.getDesc=W,r.setDesc=J,r.setDescs=K,r.getNames=v.get=G,r.getSymbols=H,u&&!n(19)&&a(F,"propertyIsEnumerable",V,!0));var z={for:function(t){return i(D,t+="")?D[t]:D[t]=j(t)},keyFor:function(t){return y(D,t)},useSetter:function(){N=!0},useSimple:function(){N=!1}};r.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=d(t);z[t]=I?e:C(e)}),N=!0,c(c.G+c.W,{Symbol:j}),c(c.S,"Symbol",z),c(c.S+c.F*!I,"Object",{create:L,defineProperty:J,defineProperties:K,getOwnPropertyDescriptor:W,getOwnPropertyNames:G,getOwnPropertySymbols:H}),P&&c(c.S+c.F*(!I||q),"JSON",{stringify:R}),l(j,"Symbol"),l(Math,"Math",!0),l(o.JSON,"JSON",!0)},function(t,e,n){var r=n(0),o=n(6);t.exports=function(t,e){for(var n,i=o(t),u=r.getKeys(i),c=u.length,a=0;c>a;)if(i[n=u[a++]]===e)return n}},function(t,e,n){var r=n(6),o=n(0).getNames,i={}.toString,u="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[],c=function(t){try{return o(t)}catch(t){return u.slice()}};t.exports.get=function(t){return u&&"[object Window]"==i.call(t)?c(t):o(r(t))}},function(t,e,n){var r=n(0);t.exports=function(t){var e=r.getKeys(t),n=r.getSymbols;if(n)for(var o,i=n(t),u=r.isEnum,c=0;i.length>c;)u.call(t,o=i[c++])&&e.push(o);return e}},function(t,e,n){var r=n(16);t.exports=Array.isArray||function(t){return"Array"==r(t)}},function(t,e){},function(t,e,n){"use strict";function r(t){return t&&t.__esModule?t:{default:t}}e.__esModule=!0;var o=n(60),i=r(o),u=n(64),c=r(u),a=n(17),s=r(a);e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,s.default)(e)));t.prototype=(0,c.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(i.default?(0,i.default)(t,e):t.__proto__=e)}},function(t,e,n){t.exports={default:n(61),__esModule:!0}},function(t,e,n){n(62),t.exports=n(1).Object.setPrototypeOf},function(t,e,n){var r=n(2);r(r.S,"Object",{setPrototypeOf:n(63).set})},function(t,e,n){var r=n(0).getDesc,o=n(25),i=n(24),u=function(t,e){if(i(t),!o(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,o){try{o=n(13)(Function.call,r(Object.prototype,"__proto__").set,2),o(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return u(t,n),e?t.__proto__=n:o(t,n),t}}({},!1):void 0),check:u}},function(t,e,n){t.exports={default:n(65),__esModule:!0}},function(t,e,n){var r=n(0);t.exports=function(t,e){return r.create(t,e)}},function(e,n){e.exports=t},function(t,n){t.exports=e},function(t,e){t.exports=n},function(t,e){t.exports=r},function(t,e){t.exports=o}])});
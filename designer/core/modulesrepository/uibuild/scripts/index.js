define("modulesrepository",["react","reactpages","reactwebcommon","itemdetailview","uicommon"],function(t,e,n,r,o){return function(t){var e={};function n(r){if(e[r])return e[r].exports;var o=e[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=t,n.c=e,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},n.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var o in t)n.d(r,o,function(e){return t[e]}.bind(null,o));return r},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=76)}([function(e,n){e.exports=t},function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,e,n){t.exports={default:n(42),__esModule:!0}},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){"use strict";e.__esModule=!0;var r,o=n(45),u=(r=o)&&r.__esModule?r:{default:r};e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var r=e[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),(0,u.default)(t,r.key,r)}}return function(e,n,r){return n&&t(e.prototype,n),r&&t(e,r),e}}()},function(t,e,n){"use strict";e.__esModule=!0;var r,o=n(26),u=(r=o)&&r.__esModule?r:{default:r};e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,u.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){"use strict";e.__esModule=!0;var r=i(n(65)),o=i(n(69)),u=i(n(26));function i(t){return t&&t.__esModule?t:{default:t}}e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,u.default)(e)));t.prototype=(0,o.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(r.default?(0,r.default)(t,e):t.__proto__=e)}},function(t,e,n){var r=n(11),o=n(2),u=n(22),i=function(t,e,n){var c,a,s,f=t&i.F,l=t&i.G,p=t&i.S,d=t&i.P,y=t&i.B,m=t&i.W,h=l?o:o[e]||(o[e]={}),v=l?r:p?r[e]:(r[e]||{}).prototype;for(c in l&&(n=e),n)(a=!f&&v&&c in v)&&c in h||(s=a?v[c]:n[c],h[c]=l&&"function"!=typeof v[c]?n[c]:y&&a?u(s,r):m&&v[c]==s?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(s):d&&"function"==typeof s?u(Function.call,s):s,d&&((h.prototype||(h.prototype={}))[c]=s))};i.F=1,i.G=2,i.S=4,i.P=8,i.B=16,i.W=32,t.exports=i},function(t,e,n){var r=n(31)("wks"),o=n(32),u=n(11).Symbol;t.exports=function(t){return r[t]||(r[t]=u&&u[t]||(u||o)("Symbol."+t))}},function(t,n){t.exports=e},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e,n){var r=n(24),o=n(16);t.exports=function(t){return r(o(t))}},function(t,e,n){t.exports={default:n(38),__esModule:!0}},function(t,e){t.exports=n},function(t,e){t.exports=function(t){if(null==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e,n){var r=n(1),o=n(18);t.exports=n(30)?function(t,e,n){return r.setDesc(t,e,o(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports={}},function(t,e,n){var r=n(1).setDesc,o=n(19),u=n(9)("toStringTag");t.exports=function(t,e,n){t&&!o(t=n?t:t.prototype,u)&&r(t,u,{configurable:!0,value:e})}},function(t,e,n){var r=n(40);t.exports=function(t,e,n){if(r(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,r){return t.call(e,n,r)};case 3:return function(n,r,o){return t.call(e,n,r,o)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){var r=n(16);t.exports=function(t){return Object(r(t))}},function(t,e,n){var r=n(25);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==r(t)?t.split(""):Object(t)}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){"use strict";e.__esModule=!0;var r=i(n(47)),o=i(n(57)),u="function"==typeof o.default&&"symbol"==typeof r.default?function(t){return typeof t}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":typeof t};function i(t){return t&&t.__esModule?t:{default:t}}e.default="function"==typeof o.default&&"symbol"===u(r.default)?function(t){return void 0===t?"undefined":u(t)}:function(t){return t&&"function"==typeof o.default&&t.constructor===o.default&&t!==o.default.prototype?"symbol":void 0===t?"undefined":u(t)}},function(t,e,n){"use strict";var r=n(28),o=n(8),u=n(29),i=n(17),c=n(19),a=n(20),s=n(52),f=n(21),l=n(1).getProto,p=n(9)("iterator"),d=!([].keys&&"next"in[].keys()),y=function(){return this};t.exports=function(t,e,n,m,h,v,g){s(n,e,m);var _,b,x=function(t){if(!d&&t in P)return P[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},w=e+" Iterator",O="values"==h,S=!1,P=t.prototype,j=P[p]||P["@@iterator"]||h&&P[h],M=j||x(h);if(j){var k=l(M.call(new t));f(k,w,!0),!r&&c(P,"@@iterator")&&i(k,p,y),O&&"values"!==j.name&&(S=!0,M=function(){return j.call(this)})}if(r&&!g||!d&&!S&&P[p]||i(P,p,M),a[e]=M,a[w]=y,h)if(_={values:O?M:x("values"),keys:v?M:x("keys"),entries:O?x("entries"):M},g)for(b in _)b in P||u(P,b,_[b]);else o(o.P+o.F*(d||S),e,_);return _}},function(t,e){t.exports=!0},function(t,e,n){t.exports=n(17)},function(t,e,n){t.exports=!n(12)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){var r=n(11),o=r["__core-js_shared__"]||(r["__core-js_shared__"]={});t.exports=function(t){return o[t]||(o[t]={})}},function(t,e){var n=0,r=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+r).toString(36))}},function(t,e,n){var r=n(34);t.exports=function(t){if(!r(t))throw TypeError(t+" is not an object!");return t}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){t.exports=n(71)()},function(t,e){t.exports=r},function(t,e){},function(t,e,n){n(39),t.exports=n(2).Object.assign},function(t,e,n){var r=n(8);r(r.S+r.F,"Object",{assign:n(41)})},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){var r=n(1),o=n(23),u=n(24);t.exports=n(12)(function(){var t=Object.assign,e={},n={},r=Symbol(),o="abcdefghijklmnopqrst";return e[r]=7,o.split("").forEach(function(t){n[t]=t}),7!=t({},e)[r]||Object.keys(t({},n)).join("")!=o})?function(t,e){for(var n=o(t),i=arguments,c=i.length,a=1,s=r.getKeys,f=r.getSymbols,l=r.isEnum;c>a;)for(var p,d=u(i[a++]),y=f?s(d).concat(f(d)):s(d),m=y.length,h=0;m>h;)l.call(d,p=y[h++])&&(n[p]=d[p]);return n}:Object.assign},function(t,e,n){n(43),t.exports=n(2).Object.getPrototypeOf},function(t,e,n){var r=n(23);n(44)("getPrototypeOf",function(t){return function(e){return t(r(e))}})},function(t,e,n){var r=n(8),o=n(2),u=n(12);t.exports=function(t,e){var n=(o.Object||{})[t]||Object[t],i={};i[t]=e(n),r(r.S+r.F*u(function(){n(1)}),"Object",i)}},function(t,e,n){t.exports={default:n(46),__esModule:!0}},function(t,e,n){var r=n(1);t.exports=function(t,e,n){return r.setDesc(t,e,n)}},function(t,e,n){t.exports={default:n(48),__esModule:!0}},function(t,e,n){n(49),n(53),t.exports=n(9)("iterator")},function(t,e,n){"use strict";var r=n(50)(!0);n(27)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=r(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){var r=n(51),o=n(16);t.exports=function(t){return function(e,n){var u,i,c=String(o(e)),a=r(n),s=c.length;return a<0||a>=s?t?"":void 0:(u=c.charCodeAt(a))<55296||u>56319||a+1===s||(i=c.charCodeAt(a+1))<56320||i>57343?t?c.charAt(a):u:t?c.slice(a,a+2):i-56320+(u-55296<<10)+65536}}},function(t,e){var n=Math.ceil,r=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?r:n)(t)}},function(t,e,n){"use strict";var r=n(1),o=n(18),u=n(21),i={};n(17)(i,n(9)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=r.create(i,{next:o(1,n)}),u(t,e+" Iterator")}},function(t,e,n){n(54);var r=n(20);r.NodeList=r.HTMLCollection=r.Array},function(t,e,n){"use strict";var r=n(55),o=n(56),u=n(20),i=n(13);t.exports=n(27)(Array,"Array",function(t,e){this._t=i(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,o(1)):o(0,"keys"==e?n:"values"==e?t[n]:[n,t[n]])},"values"),u.Arguments=u.Array,r("keys"),r("values"),r("entries")},function(t,e){t.exports=function(){}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e,n){t.exports={default:n(58),__esModule:!0}},function(t,e,n){n(59),n(64),t.exports=n(2).Symbol},function(t,e,n){"use strict";var r=n(1),o=n(11),u=n(19),i=n(30),c=n(8),a=n(29),s=n(12),f=n(31),l=n(21),p=n(32),d=n(9),y=n(60),m=n(61),h=n(62),v=n(63),g=n(33),_=n(13),b=n(18),x=r.getDesc,w=r.setDesc,O=r.create,S=m.get,P=o.Symbol,j=o.JSON,M=j&&j.stringify,k=!1,E=d("_hidden"),D=r.isEnum,N=f("symbol-registry"),T=f("symbols"),A="function"==typeof P,I=Object.prototype,R=i&&s(function(){return 7!=O(w({},"a",{get:function(){return w(this,"a",{value:7}).a}})).a})?function(t,e,n){var r=x(I,e);r&&delete I[e],w(t,e,n),r&&t!==I&&w(I,e,r)}:w,F=function(t){var e=T[t]=O(P.prototype);return e._k=t,i&&k&&R(I,t,{configurable:!0,set:function(e){u(this,E)&&u(this[E],t)&&(this[E][t]=!1),R(this,t,b(1,e))}}),e},C=function(t){return"symbol"==typeof t},V=function(t,e,n){return n&&u(T,e)?(n.enumerable?(u(t,E)&&t[E][e]&&(t[E][e]=!1),n=O(n,{enumerable:b(0,!1)})):(u(t,E)||w(t,E,b(1,{})),t[E][e]=!0),R(t,e,n)):w(t,e,n)},B=function(t,e){g(t);for(var n,r=h(e=_(e)),o=0,u=r.length;u>o;)V(t,n=r[o++],e[n]);return t},W=function(t,e){return void 0===e?O(t):B(O(t),e)},J=function(t){var e=D.call(this,t);return!(e||!u(this,t)||!u(T,t)||u(this,E)&&this[E][t])||e},K=function(t,e){var n=x(t=_(t),e);return!n||!u(T,e)||u(t,E)&&t[E][e]||(n.enumerable=!0),n},L=function(t){for(var e,n=S(_(t)),r=[],o=0;n.length>o;)u(T,e=n[o++])||e==E||r.push(e);return r},G=function(t){for(var e,n=S(_(t)),r=[],o=0;n.length>o;)u(T,e=n[o++])&&r.push(T[e]);return r},q=s(function(){var t=P();return"[null]"!=M([t])||"{}"!=M({a:t})||"{}"!=M(Object(t))});A||(a((P=function(){if(C(this))throw TypeError("Symbol is not a constructor");return F(p(arguments.length>0?arguments[0]:void 0))}).prototype,"toString",function(){return this._k}),C=function(t){return t instanceof P},r.create=W,r.isEnum=J,r.getDesc=K,r.setDesc=V,r.setDescs=B,r.getNames=m.get=L,r.getSymbols=G,i&&!n(28)&&a(I,"propertyIsEnumerable",J,!0));var z={for:function(t){return u(N,t+="")?N[t]:N[t]=P(t)},keyFor:function(t){return y(N,t)},useSetter:function(){k=!0},useSimple:function(){k=!1}};r.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=d(t);z[t]=A?e:F(e)}),k=!0,c(c.G+c.W,{Symbol:P}),c(c.S,"Symbol",z),c(c.S+c.F*!A,"Object",{create:W,defineProperty:V,defineProperties:B,getOwnPropertyDescriptor:K,getOwnPropertyNames:L,getOwnPropertySymbols:G}),j&&c(c.S+c.F*(!A||q),"JSON",{stringify:function(t){if(void 0!==t&&!C(t)){for(var e,n,r=[t],o=1,u=arguments;u.length>o;)r.push(u[o++]);return"function"==typeof(e=r[1])&&(n=e),!n&&v(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!C(e))return e}),r[1]=e,M.apply(j,r)}}}),l(P,"Symbol"),l(Math,"Math",!0),l(o.JSON,"JSON",!0)},function(t,e,n){var r=n(1),o=n(13);t.exports=function(t,e){for(var n,u=o(t),i=r.getKeys(u),c=i.length,a=0;c>a;)if(u[n=i[a++]]===e)return n}},function(t,e,n){var r=n(13),o=n(1).getNames,u={}.toString,i="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[];t.exports.get=function(t){return i&&"[object Window]"==u.call(t)?function(t){try{return o(t)}catch(t){return i.slice()}}(t):o(r(t))}},function(t,e,n){var r=n(1);t.exports=function(t){var e=r.getKeys(t),n=r.getSymbols;if(n)for(var o,u=n(t),i=r.isEnum,c=0;u.length>c;)i.call(t,o=u[c++])&&e.push(o);return e}},function(t,e,n){var r=n(25);t.exports=Array.isArray||function(t){return"Array"==r(t)}},function(t,e){},function(t,e,n){t.exports={default:n(66),__esModule:!0}},function(t,e,n){n(67),t.exports=n(2).Object.setPrototypeOf},function(t,e,n){var r=n(8);r(r.S,"Object",{setPrototypeOf:n(68).set})},function(t,e,n){var r=n(1).getDesc,o=n(34),u=n(33),i=function(t,e){if(u(t),!o(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,o){try{(o=n(22)(Function.call,r(Object.prototype,"__proto__").set,2))(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return i(t,n),e?t.__proto__=n:o(t,n),t}}({},!1):void 0),check:i}},function(t,e,n){t.exports={default:n(70),__esModule:!0}},function(t,e,n){var r=n(1);t.exports=function(t,e){return r.create(t,e)}},function(t,e,n){"use strict";var r=n(72),o=n(73),u=n(74);t.exports=function(){function t(t,e,n,r,i,c){c!==u&&o(!1,"Calling PropTypes validators directly is not supported by the `prop-types` package. Use PropTypes.checkPropTypes() to call them. Read more at http://fb.me/use-check-prop-types")}function e(){return t}t.isRequired=t;var n={array:t,bool:t,func:t,number:t,object:t,string:t,symbol:t,any:t,arrayOf:e,element:t,instanceOf:e,node:t,objectOf:e,oneOf:e,oneOfType:e,shape:e,exact:e};return n.checkPropTypes=r,n.PropTypes=n,n}},function(t,e,n){"use strict";function r(t){return function(){return t}}var o=function(){};o.thatReturns=r,o.thatReturnsFalse=r(!1),o.thatReturnsTrue=r(!0),o.thatReturnsNull=r(null),o.thatReturnsThis=function(){return this},o.thatReturnsArgument=function(t){return t},t.exports=o},function(t,e,n){"use strict";var r=function(t){};t.exports=function(t,e,n,o,u,i,c,a){if(r(e),!t){var s;if(void 0===e)s=new Error("Minified exception occurred; use the non-minified dev environment for the full error message and additional helpful warnings.");else{var f=[n,o,u,i,c,a],l=0;(s=new Error(e.replace(/%s/g,function(){return f[l++]}))).name="Invariant Violation"}throw s.framesToPop=1,s}}},function(t,e,n){"use strict";t.exports="SECRET_DO_NOT_PASS_THIS_OR_YOU_WILL_BE_FIRED"},function(t,e){t.exports=o},function(t,e,n){"use strict";n.r(e);var r=n(0),o=n.n(r),u=(n(37),n(14)),i=n.n(u),c=n(3),a=n.n(c),s=n(4),f=n.n(s),l=n(5),p=n.n(l),d=n(6),y=n.n(d),m=n(7),h=n.n(m),v=n(10),g=(n(35),function(t){function e(t){f()(this,e);var n=y()(this,(e.__proto__||a()(e)).call(this,t));console.log("props in module settings view",t);var r={type:"form",info:{}};return t.formValue&&t.formValue.Module&&t.formValue.Module.ParamsForm&&(r=i()(r,t.formValue.Module.ParamsForm)),console.log("form desc",r),n.formDesc=r,n.state={formData:{}},n}return h()(e,t),p()(e,[{key:"render",value:function(){var t=this.props;return console.log("props... module settings ",t,this.formDesc),o.a.createElement(v.Panel,{formData:this.state.formData,subform:!0,parentFormRef:t.formRef,description:this.formDesc})}}]),e}(o.a.Component)),_=(n(75),n(15)),b=n(36),x=(n(35),function(t){function e(t){f()(this,e);var n=y()(this,(e.__proto__||a()(e)).call(this,t));return n.getItem=function(t,e,n){var r=t.methods;return o.a.createElement(_uikit.Block,null,o.a.createElement(_uikit.Block,{className:"row center valigncenter ma10"},e.Name),o.a.createElement(_uikit.Block,{className:"row m10"},o.a.createElement(_.Action,{className:"p10",action:{actiontype:"method",method:function(){console.log("selected ",n),r.itemSelectionChange(n,!0)},params:{}}},"Select"),o.a.createElement(_.Action,{className:"p10",action:{actiontype:"method",method:r.openDetail,params:{data:e,index:n}}},"Details")))},n.submit=function(){var t=n.view.current.selectedItems();n.props.description.add(t,null,null,!0)},n.view=o.a.createRef(),n}return h()(e,t),p()(e,[{key:"render",value:function(){return o.a.createElement(_uikit.Block,null,o.a.createElement(_uikit.Block,{className:" w100 right "},o.a.createElement(_.Action,{widget:"button",className:"p10",action:{actiontype:"method",method:this.submit,params:{}}},"Submit")),o.a.createElement(b.ItemDetailView,{id:"repositoryview",viewRef:this.view,getItem:this.getItem,editable:!0,entityName:"ModuleDefinition"}))}}]),e}(o.a.Component)),w=function(t){function e(t){f()(this,e);var n=y()(this,(e.__proto__||a()(e)).call(this,t));return O.call(n),console.log("module config constructor",t),n.state={module:n.getModuleName(t),value:n.getValue(t)},n}return h()(e,t),p()(e,[{key:"componentWillRecieveProps",value:function(t){var e=this.getModuleName(t),n=this.getValue(t);this.setState(i()({},this.state,{module:e,value:n}))}},{key:"render",value:function(){return o.a.createElement(v.Panel,{id:"module_config_"+this.state.module,formData:this.state.value,type:"form"})}}]),e}(o.a.Component),O=function(){this.getModuleName=function(t){return console.log("props",t),t.formValue.Module.Id},this.getValue=function(t){return t.formValue.Settings}};function S(t,e,n,r,o,u){}n.d(e,"Initialize",function(){return S}),n.d(e,"ModuleSettings",function(){return g}),n.d(e,"ModuleConfig",function(){return w}),n.d(e,"ModuleSelect",function(){return x})}])});
//# sourceMappingURL=index.js.map
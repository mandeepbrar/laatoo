define("subentity",["react","reactwebcommon","reactpages"],function(t,e,n){return function(t){var e={};function n(o){if(e[o])return e[o].exports;var r=e[o]={i:o,l:!1,exports:{}};return t[o].call(r.exports,r,r.exports,n),r.l=!0,r.exports}return n.m=t,n.c=e,n.d=function(t,e,o){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:o})},n.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var o=Object.create(null);if(n.r(o),Object.defineProperty(o,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var r in t)n.d(o,r,function(e){return t[e]}.bind(null,r));return o},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/",n(n.s=76)}([function(e,n){e.exports=t},function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(t,e,n){t.exports={default:n(72),__esModule:!0}},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,n){t.exports=e},function(t,e,n){var o=n(26)("wks"),r=n(25),i=n(9).Symbol;t.exports=function(t){return o[t]||(o[t]=i&&i[t]||(i||r)("Symbol."+t))}},function(t,e,n){var o=n(9),r=n(3),i=n(34),s=function(t,e,n){var a,u,c,l=t&s.F,f=t&s.G,p=t&s.S,d=t&s.P,m=t&s.B,h=t&s.W,y=f?r:r[e]||(r[e]={}),v=f?o:p?o[e]:(o[e]||{}).prototype;for(a in f&&(n=e),n)(u=!l&&v&&a in v)&&a in y||(c=u?v[a]:n[a],y[a]=f&&"function"!=typeof v[a]?n[a]:m&&u?i(c,o):h&&v[a]==c?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(c):d&&"function"==typeof c?i(Function.call,c):c,d&&((y.prototype||(y.prototype={}))[a]=c))};s.F=1,s.G=2,s.S=4,s.P=8,s.B=16,s.W=32,t.exports=s},function(t,e,n){var o=n(33),r=n(21);t.exports=function(t){return o(r(t))}},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e){t.exports=n},function(t,e,n){"use strict";e.__esModule=!0;var o=s(n(47)),r=s(n(43)),i=s(n(31));function s(t){return t&&t.__esModule?t:{default:t}}e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,i.default)(e)));t.prototype=(0,r.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(o.default?(0,o.default)(t,e):t.__proto__=e)}},function(t,e,n){"use strict";e.__esModule=!0;var o,r=n(31),i=(o=r)&&o.__esModule?o:{default:o};e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,i.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){"use strict";e.__esModule=!0;var o,r=n(67),i=(o=r)&&o.__esModule?o:{default:o};e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var o=e[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),(0,i.default)(t,o.key,o)}}return function(e,n,o){return n&&t(e.prototype,n),o&&t(e,o),e}}()},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){t.exports={default:n(69),__esModule:!0}},function(t,e,n){var o=n(1).setDesc,r=n(18),i=n(5)("toStringTag");t.exports=function(t,e,n){t&&!r(t=n?t:t.prototype,i)&&o(t,i,{configurable:!0,value:e})}},function(t,e){t.exports={}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e,n){var o=n(1),r=n(19);t.exports=n(27)?function(t,e,n){return o.setDesc(t,e,r(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e,n){var o=n(21);t.exports=function(t){return Object(o(t))}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){var o=n(23);t.exports=function(t){if(!o(t))throw TypeError(t+" is not an object!");return t}},function(t,e){var n=0,o=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+o).toString(36))}},function(t,e,n){var o=n(9),r=o["__core-js_shared__"]||(o["__core-js_shared__"]={});t.exports=function(t){return r[t]||(r[t]={})}},function(t,e,n){t.exports=!n(8)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){t.exports=n(20)},function(t,e){t.exports=!0},function(t,e,n){"use strict";var o=n(29),r=n(6),i=n(28),s=n(20),a=n(18),u=n(17),c=n(60),l=n(16),f=n(1).getProto,p=n(5)("iterator"),d=!([].keys&&"next"in[].keys()),m=function(){return this};t.exports=function(t,e,n,h,y,v,g){c(n,e,h);var b,_,x=function(t){if(!d&&t in k)return k[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},O=e+" Iterator",S="values"==y,w=!1,k=t.prototype,E=k[p]||k["@@iterator"]||y&&k[y],P=E||x(y);if(E){var F=f(P.call(new t));l(F,O,!0),!o&&a(k,"@@iterator")&&s(F,p,m),S&&"values"!==E.name&&(w=!0,P=function(){return E.call(this)})}if(o&&!g||!d&&!w&&k[p]||s(k,p,P),u[e]=P,u[O]=m,y)if(b={values:S?P:x("values"),keys:v?P:x("keys"),entries:S?x("entries"):P},g)for(_ in b)_ in k||i(k,_,b[_]);else r(r.P+r.F*(d||w),e,b);return b}},function(t,e,n){"use strict";e.__esModule=!0;var o=s(n(65)),r=s(n(55)),i="function"==typeof r.default&&"symbol"==typeof o.default?function(t){return typeof t}:function(t){return t&&"function"==typeof r.default&&t.constructor===r.default&&t!==r.default.prototype?"symbol":typeof t};function s(t){return t&&t.__esModule?t:{default:t}}e.default="function"==typeof r.default&&"symbol"===i(o.default)?function(t){return void 0===t?"undefined":i(t)}:function(t){return t&&"function"==typeof r.default&&t.constructor===r.default&&t!==r.default.prototype?"symbol":void 0===t?"undefined":i(t)}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){var o=n(32);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==o(t)?t.split(""):Object(t)}},function(t,e,n){var o=n(73);t.exports=function(t,e,n){if(o(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,o){return t.call(e,n,o)};case 3:return function(n,o,r){return t.call(e,n,o,r)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){var o=n(6),r=n(3),i=n(8);t.exports=function(t,e){var n=(r.Object||{})[t]||Object[t],s={};s[t]=e(n),o(o.S+o.F*i(function(){n(1)}),"Object",s)}},function(t,e,n){t.exports={default:n(75),__esModule:!0}},function(t,e,n){"use strict";t.exports="SECRET_DO_NOT_PASS_THIS_OR_YOU_WILL_BE_FIRED"},function(t,e,n){"use strict";var o=function(t){};t.exports=function(t,e,n,r,i,s,a,u){if(o(e),!t){var c;if(void 0===e)c=new Error("Minified exception occurred; use the non-minified dev environment for the full error message and additional helpful warnings.");else{var l=[n,r,i,s,a,u],f=0;(c=new Error(e.replace(/%s/g,function(){return l[f++]}))).name="Invariant Violation"}throw c.framesToPop=1,c}}},function(t,e,n){"use strict";function o(t){return function(){return t}}var r=function(){};r.thatReturns=o,r.thatReturnsFalse=o(!1),r.thatReturnsTrue=o(!0),r.thatReturnsNull=o(null),r.thatReturnsThis=function(){return this},r.thatReturnsArgument=function(t){return t},t.exports=r},function(t,e,n){"use strict";var o=n(39),r=n(38),i=n(37);t.exports=function(){function t(t,e,n,o,s,a){a!==i&&r(!1,"Calling PropTypes validators directly is not supported by the `prop-types` package. Use PropTypes.checkPropTypes() to call them. Read more at http://fb.me/use-check-prop-types")}function e(){return t}t.isRequired=t;var n={array:t,bool:t,func:t,number:t,object:t,string:t,symbol:t,any:t,arrayOf:e,element:t,instanceOf:e,node:t,objectOf:e,oneOf:e,oneOfType:e,shape:e,exact:e};return n.checkPropTypes=o,n.PropTypes=n,n}},function(t,e,n){t.exports=n(40)()},function(t,e,n){var o=n(1);t.exports=function(t,e){return o.create(t,e)}},function(t,e,n){t.exports={default:n(42),__esModule:!0}},function(t,e,n){var o=n(1).getDesc,r=n(23),i=n(24),s=function(t,e){if(i(t),!r(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,r){try{(r=n(34)(Function.call,o(Object.prototype,"__proto__").set,2))(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return s(t,n),e?t.__proto__=n:r(t,n),t}}({},!1):void 0),check:s}},function(t,e,n){var o=n(6);o(o.S,"Object",{setPrototypeOf:n(44).set})},function(t,e,n){n(45),t.exports=n(3).Object.setPrototypeOf},function(t,e,n){t.exports={default:n(46),__esModule:!0}},function(t,e){},function(t,e,n){var o=n(32);t.exports=Array.isArray||function(t){return"Array"==o(t)}},function(t,e,n){var o=n(1);t.exports=function(t){var e=o.getKeys(t),n=o.getSymbols;if(n)for(var r,i=n(t),s=o.isEnum,a=0;i.length>a;)s.call(t,r=i[a++])&&e.push(r);return e}},function(t,e,n){var o=n(7),r=n(1).getNames,i={}.toString,s="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[];t.exports.get=function(t){return s&&"[object Window]"==i.call(t)?function(t){try{return r(t)}catch(t){return s.slice()}}(t):r(o(t))}},function(t,e,n){var o=n(1),r=n(7);t.exports=function(t,e){for(var n,i=r(t),s=o.getKeys(i),a=s.length,u=0;a>u;)if(i[n=s[u++]]===e)return n}},function(t,e,n){"use strict";var o=n(1),r=n(9),i=n(18),s=n(27),a=n(6),u=n(28),c=n(8),l=n(26),f=n(16),p=n(25),d=n(5),m=n(52),h=n(51),y=n(50),v=n(49),g=n(24),b=n(7),_=n(19),x=o.getDesc,O=o.setDesc,S=o.create,w=h.get,k=r.Symbol,E=r.JSON,P=E&&E.stringify,F=!1,j=d("_hidden"),C=o.isEnum,N=l("symbol-registry"),R=l("symbols"),M="function"==typeof k,D=Object.prototype,A=s&&c(function(){return 7!=S(O({},"a",{get:function(){return O(this,"a",{value:7}).a}})).a})?function(t,e,n){var o=x(D,e);o&&delete D[e],O(t,e,n),o&&t!==D&&O(D,e,o)}:O,T=function(t){var e=R[t]=S(k.prototype);return e._k=t,s&&F&&A(D,t,{configurable:!0,set:function(e){i(this,j)&&i(this[j],t)&&(this[j][t]=!1),A(this,t,_(1,e))}}),e},I=function(t){return"symbol"==typeof t},B=function(t,e,n){return n&&i(R,e)?(n.enumerable?(i(t,j)&&t[j][e]&&(t[j][e]=!1),n=S(n,{enumerable:_(0,!1)})):(i(t,j)||O(t,j,_(1,{})),t[j][e]=!0),A(t,e,n)):O(t,e,n)},V=function(t,e){g(t);for(var n,o=y(e=b(e)),r=0,i=o.length;i>r;)B(t,n=o[r++],e[n]);return t},W=function(t,e){return void 0===e?S(t):V(S(t),e)},L=function(t){var e=C.call(this,t);return!(e||!i(this,t)||!i(R,t)||i(this,j)&&this[j][t])||e},J=function(t,e){var n=x(t=b(t),e);return!n||!i(R,e)||i(t,j)&&t[j][e]||(n.enumerable=!0),n},K=function(t){for(var e,n=w(b(t)),o=[],r=0;n.length>r;)i(R,e=n[r++])||e==j||o.push(e);return o},G=function(t){for(var e,n=w(b(t)),o=[],r=0;n.length>r;)i(R,e=n[r++])&&o.push(R[e]);return o},q=c(function(){var t=k();return"[null]"!=P([t])||"{}"!=P({a:t})||"{}"!=P(Object(t))});M||(u((k=function(){if(I(this))throw TypeError("Symbol is not a constructor");return T(p(arguments.length>0?arguments[0]:void 0))}).prototype,"toString",function(){return this._k}),I=function(t){return t instanceof k},o.create=W,o.isEnum=L,o.getDesc=J,o.setDesc=B,o.setDescs=V,o.getNames=h.get=K,o.getSymbols=G,s&&!n(29)&&u(D,"propertyIsEnumerable",L,!0));var H={for:function(t){return i(N,t+="")?N[t]:N[t]=k(t)},keyFor:function(t){return m(N,t)},useSetter:function(){F=!0},useSimple:function(){F=!1}};o.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=d(t);H[t]=M?e:T(e)}),F=!0,a(a.G+a.W,{Symbol:k}),a(a.S,"Symbol",H),a(a.S+a.F*!M,"Object",{create:W,defineProperty:B,defineProperties:V,getOwnPropertyDescriptor:J,getOwnPropertyNames:K,getOwnPropertySymbols:G}),E&&a(a.S+a.F*(!M||q),"JSON",{stringify:function(t){if(void 0!==t&&!I(t)){for(var e,n,o=[t],r=1,i=arguments;i.length>r;)o.push(i[r++]);return"function"==typeof(e=o[1])&&(n=e),!n&&v(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!I(e))return e}),o[1]=e,P.apply(E,o)}}}),f(k,"Symbol"),f(Math,"Math",!0),f(r.JSON,"JSON",!0)},function(t,e,n){n(53),n(48),t.exports=n(3).Symbol},function(t,e,n){t.exports={default:n(54),__esModule:!0}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e){t.exports=function(){}},function(t,e,n){"use strict";var o=n(57),r=n(56),i=n(17),s=n(7);t.exports=n(30)(Array,"Array",function(t,e){this._t=s(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,r(1)):r(0,"keys"==e?n:"values"==e?t[n]:[n,t[n]])},"values"),i.Arguments=i.Array,o("keys"),o("values"),o("entries")},function(t,e,n){n(58);var o=n(17);o.NodeList=o.HTMLCollection=o.Array},function(t,e,n){"use strict";var o=n(1),r=n(19),i=n(16),s={};n(20)(s,n(5)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=o.create(s,{next:r(1,n)}),i(t,e+" Iterator")}},function(t,e){var n=Math.ceil,o=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?o:n)(t)}},function(t,e,n){var o=n(61),r=n(21);t.exports=function(t){return function(e,n){var i,s,a=String(r(e)),u=o(n),c=a.length;return u<0||u>=c?t?"":void 0:(i=a.charCodeAt(u))<55296||i>56319||u+1===c||(s=a.charCodeAt(u+1))<56320||s>57343?t?a.charAt(u):i:t?a.slice(u,u+2):s-56320+(i-55296<<10)+65536}}},function(t,e,n){"use strict";var o=n(62)(!0);n(30)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=o(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){n(63),n(59),t.exports=n(5)("iterator")},function(t,e,n){t.exports={default:n(64),__esModule:!0}},function(t,e,n){var o=n(1);t.exports=function(t,e,n){return o.setDesc(t,e,n)}},function(t,e,n){t.exports={default:n(66),__esModule:!0}},function(t,e,n){var o=n(22);n(35)("getPrototypeOf",function(t){return function(e){return t(o(e))}})},function(t,e,n){n(68),t.exports=n(3).Object.getPrototypeOf},function(t,e,n){var o=n(1),r=n(22),i=n(33);t.exports=n(8)(function(){var t=Object.assign,e={},n={},o=Symbol(),r="abcdefghijklmnopqrst";return e[o]=7,r.split("").forEach(function(t){n[t]=t}),7!=t({},e)[o]||Object.keys(t({},n)).join("")!=r})?function(t,e){for(var n=r(t),s=arguments,a=s.length,u=1,c=o.getKeys,l=o.getSymbols,f=o.isEnum;a>u;)for(var p,d=i(s[u++]),m=l?c(d).concat(l(d)):c(d),h=m.length,y=0;h>y;)f.call(d,p=m[y++])&&(n[p]=d[p]);return n}:Object.assign},function(t,e,n){var o=n(6);o(o.S+o.F,"Object",{assign:n(70)})},function(t,e,n){n(71),t.exports=n(3).Object.assign},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){var o=n(22);n(35)("keys",function(t){return function(e){return t(o(e))}})},function(t,e,n){n(74),t.exports=n(3).Object.keys},function(t,e,n){"use strict";n.r(e),n.d(e,"SubEntity",function(){return w});var o=n(36),r=n.n(o),i=n(2),s=n.n(i),a=n(15),u=n.n(a),c=n(14),l=n.n(c),f=n(13),p=n.n(f),d=n(12),m=n.n(d),h=n(11),y=n.n(h),v=n(0),g=n.n(v),b=n(4),_=n(10),x=n(41),O=function(t){function e(t){l()(this,e);var n=m()(this,(e.__proto__||u()(e)).call(this,t));return n.onChange=function(t){console.log("on change of select entity--",t),n.setState(s()({},n.state,{value:t}))},n.saveValue=function(){console.log("svaing value",n.state,n.props),n.props.submit(n.state.value)},n.state={value:t.value,items:t.items},n}return y()(e,t),p()(e,[{key:"componentWillReceiveProps",value:function(t){console.log("on change of select entity--",this.state,t),this.setState(s()({},this.state,{value:t.value,items:t.items}))}},{key:"render",value:function(){console.log("rendering select entity",this.state);var t=this.props,e=t.fld,n=t.uikit,o={label:e.label,name:e.name,widget:"Select",selectItem:!0,type:"entity",items:this.state.items};return g.a.createElement(n.Block,{className:"row between-xs"},g.a.createElement(n.Block,{className:"left col-xs-10"},g.a.createElement(n.Forms.FieldWidget,{className:"w100",field:o,fieldChange:this.onChange,value:this.state.value})),g.a.createElement(n.Block,{className:"right"},g.a.createElement(b.Action,{action:{actiontype:"method"},className:"edit p10",method:this.saveValue},g.a.createElement(n.Icons.EditIcon,null)),g.a.createElement(b.Action,{action:{actiontype:"method"},className:"remove p10",method:this.props.close},g.a.createElement(n.Icons.DeleteIcon,null))))}}]),e}(g.a.Component),S=function(t){function e(t,n){var o=this;l()(this,e);var r=m()(this,(e.__proto__||u()(e)).call(this,t));return r.closeForm=function(){switch(console.log("close form",r.props.field.mode),r.props.field.mode){case"select":case"inline":console.log("inline row close form"),r.inlineRow=null;break;case"dialog":Window.closeDialog();break;case"overlay":default:r.props.overlayComponent&&r.props.overlayComponent(null)}r.setState(s()({},r.state,{formOpen:!1}))},r.actions=function(t,e,n){return console.log("actios returned",t,e,n),g.a.createElement(o.uikit.Block,{className:"right p20"},r.props.field.inline?null:g.a.createElement(o.uikit.ActionButton,{onClick:n},"Reset"),g.a.createElement(o.uikit.ActionButton,{onClick:e},"Save"))},r.edit=function(t,e,n,o){var i=r.state.value;i&&i.length>e&&(i[e]=t,console.log(" items in edit",i[e],e,t,r.state),r.props.onChange(i),r.closeForm())},r.add=function(t,e,n,o){console.log("adding subentity ",t);var i=r.state.value.slice();o&&t&&Array.isArray(t)?t.forEach(function(t){i.push(t)}):i.push(t),console.log(" items in add",i,t,r.state),r.props.onChange(i),r.closeForm()},r.removeItem=function(t,e){var n=r.state.value.slice();e>-1&&n.splice(e,1),r.props.onChange(n)},r.getFormValue=function(){var t=r.props.getFormValue();return console.log("parent form value",r.props,t),t},r.openForm=function(t,e){console.log("opened form",r.props,r.context);var n=r,o=r.props.field,i=t?function(t,o,r){return n.edit(t,e,o,r)}:r.add,a=o.addwidget?g.a.createElement(_.Panel,{title:"Add "+r.props.label,description:{type:"component",componentName:o.addwidget,module:o.addwidgetmodule,add:r.add},parentFormRef:r,subform:!0,closePanel:r.closeForm,autoSubmitOnChange:r.props.autoSubmitOnChange}):g.a.createElement(_.Panel,{actions:r.actions,inline:!0,formData:t,title:"Add "+r.props.label,parentFormRef:r,subform:!0,closePanel:r.closeForm,onSubmit:i,description:r.props.formDesc,autoSubmitOnChange:r.props.autoSubmitOnChange});switch(r.props.field.mode){case"inline":r.inlineRow=a;break;case"select":r.inlineRow=g.a.createElement(O,{fld:o,uikit:r.uikit,submit:i,items:r.state.selectOptions,entity:t,index:e,close:r.closeForm});break;case"dialog":console.log("show subentity dialog",a),Window.showDialog(null,a,r.closeForm);break;case"overlay":default:r.props.overlayComponent&&r.props.overlayComponent(a)}r.setState(s()({},r.state,{formOpen:!0}))},console.log("entity list field ",t),r.state={value:t.value,formOpen:!1,selectOptions:t.selectOptions},r.uikit=t.uikit,r}return y()(e,t),p()(e,[{key:"componentWillReceiveProps",value:function(t){console.log("entity list field : componentWillReceiveProps",t);var e={};this.state.value!=t.value&&(e.value=t.value),this.state.selectOptions!=t.selectOptions&&(e.selectOptions=t.selectOptions),r()(e).length>0&&this.setState(s()({},this.state,e))}},{key:"render",value:function(){var t=[];console.log("rendering items in entity list",this.props,this.state);var e=this.props.field,n=this;this.state.value.forEach(function(o,r){if(o){console.log("entity list ",o,e);var i=e.textField?e.textField:"Name",s=o[i];s=s||o.Title,console.log("entity text ",s,i),t.push(g.a.createElement(n.uikit.Block,{className:"row between-xs"},g.a.createElement(n.uikit.Block,{className:"left"},s),g.a.createElement(n.uikit.Block,{className:"right"},g.a.createElement(b.Action,{action:{actiontype:"method"},className:"edit p10",method:function(){n.openForm(o,r)}},g.a.createElement(n.uikit.Icons.EditIcon,null)),g.a.createElement(b.Action,{action:{actiontype:"method"},className:"remove p10",method:function(){n.removeItem(o,r)}},g.a.createElement(n.uikit.Icons.DeleteIcon,null)))))}});this.state.formOpen&&this.inlineRow&&t.push(this.inlineRow),0==t.length&&t.push("No data"),console.log("subentity items ",t);var o=[g.a.createElement(b.Action,{action:{actiontype:"method"},className:"p10",method:this.openForm}," ",g.a.createElement(this.uikit.Icons.NewIcon,null)," ")];return g.a.createElement(this.uikit.Block,{className:"entitylistfield ",title:this.props.title,titleBarActions:o},t)}}]),e}(g.a.Component),w=function(t){function e(t,n){l()(this,e);var o=m()(this,(e.__proto__||u()(e)).call(this,t));k.call(o),o.list=!!t.field.list,o.label=t.field.label?t.field.label:t.field.entity;var r=t.field.form?t.field.form:"new_form_"+t.field.entity.toLowerCase();o.formDesc={type:"form",id:r},o.uikit=n.uikit;var i=t.input.value?t.input.value:o.list?[]:{};return o.state={value:i},console.log("show subentity",o.formDesc,t,n),o}return y()(e,t),p()(e,[{key:"componentWillReceiveProps",value:function(t){console.log("componentWillReceiveProps  for SubEntity",t);var e=t.input.value?t.input.value:this.list?[]:{};this.state.value!=e&&this.setState(s()({},this.state,{value:e}))}},{key:"render",value:function(){console.log("subentity ",this.state,this.props);var t=this.props.field,e=t.skipLabel?null:this.label;return g.a.createElement(this.uikit.Block,{className:"subentity "+this.label},this.list?g.a.createElement(S,{uikit:this.uikit,getFormValue:this.context.getFormValue,field:this.props.field,onChange:this.change,label:this.label,form:this.props.form,formRef:this.props.formRef,autoSubmitOnChange:this.props.autoSubmitOnChange,selectOptions:this.state.selectOptions,overlayComponent:this.context.overlayComponent,parentFormRef:this.props.parentFormRef,formDesc:this.formDesc,title:e,value:this.state.value}):"select"==t.mode?this.selectSubEntity():g.a.createElement(_.Panel,{actions:function(){},formData:this.state.value,title:e,autoSubmitOnChange:!0,onChange:this.change,trackChanges:!0,subform:this.props.subform,formRef:this.props.formRef,parentFormRef:this.props.parentFormRef,description:this.formDesc}))}}]),e}(b.LoadableComponent),k=function(){var t=this;this.dataLoaded=function(e){"select"==t.props.field.mode&&(console.log("data loaded for SubEntity",e),t.setState(s()({},t.state,{selectOptions:e})))},this.getLoadContext=function(){console.log("get load context called",t.context);var e={formValue:t.context.getFormValue()};return t.context.getParentFormValue&&(e.parentFormValue=t.context.getParentFormValue()),console.log("get load context called",e),e},this.selectSubEntity=function(){var e=t.props.field,n={label:e.label,name:e.name,widget:"Select",selectItem:!0,type:"entity"};return g.a.createElement(t.uikit.Forms.FieldWidget,{className:"w100",field:n,fieldChange:t.change,items:t.state.selectOptions,value:t.state.value})},this.change=function(e){console.log("charnging subentity",e,t.props),t.props.fieldChange(e,t.props.name),t.setState(s()({},t.state,{value:e}))}};w.contextTypes={uikit:x.object,getFormValue:x.func,getParentFormValue:x.func,overlayComponent:x.func}}])});
//# sourceMappingURL=index.js.map
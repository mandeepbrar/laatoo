define("entitylistfield",["react","reactwebcommon","reactpages"],function(t,e,n){return function(t){function e(o){if(n[o])return n[o].exports;var r=n[o]={i:o,l:!1,exports:{}};return t[o].call(r.exports,r,r.exports,e),r.l=!0,r.exports}var n={};return e.m=t,e.c=n,e.d=function(t,n,o){e.o(t,n)||Object.defineProperty(t,n,{configurable:!1,enumerable:!0,get:o})},e.n=function(t){var n=t&&t.__esModule?function(){return t.default}:function(){return t};return e.d(n,"a",n),n},e.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},e.p="/",e(e.s=26)}([function(t,e){var n=Object;t.exports={create:n.create,getProto:n.getPrototypeOf,isEnum:{}.propertyIsEnumerable,getDesc:n.getOwnPropertyDescriptor,setDesc:n.defineProperty,setDescs:n.defineProperties,getKeys:n.keys,getNames:n.getOwnPropertyNames,getSymbols:n.getOwnPropertySymbols,each:[].forEach}},function(t,e){var n=t.exports={version:"1.2.6"};"number"==typeof __e&&(__e=n)},function(t,e,n){var o=n(4),r=n(1),i=n(13),u=function(t,e,n){var c,s,a,f=t&u.F,l=t&u.G,p=t&u.S,d=t&u.P,y=t&u.B,h=t&u.W,m=l?r:r[e]||(r[e]={}),v=l?o:p?o[e]:(o[e]||{}).prototype;l&&(n=e);for(c in n)(s=!f&&v&&c in v)&&c in m||(a=s?v[c]:n[c],m[c]=l&&"function"!=typeof v[c]?n[c]:y&&s?i(a,o):h&&v[c]==a?function(t){var e=function(e){return this instanceof t?new t(e):t(e)};return e.prototype=t.prototype,e}(a):d&&"function"==typeof a?i(Function.call,a):a,d&&((m.prototype||(m.prototype={}))[c]=a))};u.F=1,u.G=2,u.S=4,u.P=8,u.B=16,u.W=32,t.exports=u},function(t,e,n){var o=n(22)("wks"),r=n(23),i=n(4).Symbol;t.exports=function(t){return o[t]||(o[t]=i&&i[t]||(i||r)("Symbol."+t))}},function(t,e){var n=t.exports="undefined"!=typeof window&&window.Math==Math?window:"undefined"!=typeof self&&self.Math==Math?self:Function("return this")();"number"==typeof __g&&(__g=n)},function(t,e){t.exports=function(t){try{return!!t()}catch(t){return!0}}},function(t,e,n){var o=n(15),r=n(7);t.exports=function(t){return o(r(t))}},function(t,e){t.exports=function(t){if(void 0==t)throw TypeError("Can't call method on  "+t);return t}},function(t,e,n){var o=n(0),r=n(9);t.exports=n(21)?function(t,e,n){return o.setDesc(t,e,r(1,n))}:function(t,e,n){return t[e]=n,t}},function(t,e){t.exports=function(t,e){return{enumerable:!(1&t),configurable:!(2&t),writable:!(4&t),value:e}}},function(t,e){var n={}.hasOwnProperty;t.exports=function(t,e){return n.call(t,e)}},function(t,e){t.exports={}},function(t,e,n){var o=n(0).setDesc,r=n(10),i=n(3)("toStringTag");t.exports=function(t,e,n){t&&!r(t=n?t:t.prototype,i)&&o(t,i,{configurable:!0,value:e})}},function(t,e,n){var o=n(30);t.exports=function(t,e,n){if(o(t),void 0===e)return t;switch(n){case 1:return function(n){return t.call(e,n)};case 2:return function(n,o){return t.call(e,n,o)};case 3:return function(n,o,r){return t.call(e,n,o,r)}}return function(){return t.apply(e,arguments)}}},function(t,e,n){var o=n(7);t.exports=function(t){return Object(o(t))}},function(t,e,n){var o=n(16);t.exports=Object("z").propertyIsEnumerable(0)?Object:function(t){return"String"==o(t)?t.split(""):Object(t)}},function(t,e){var n={}.toString;t.exports=function(t){return n.call(t).slice(8,-1)}},function(t,e,n){"use strict";function o(t){return t&&t.__esModule?t:{default:t}}e.__esModule=!0;var r=n(41),i=o(r),u=n(51),c=o(u),s="function"==typeof c.default&&"symbol"==typeof i.default?function(t){return typeof t}:function(t){return t&&"function"==typeof c.default&&t.constructor===c.default&&t!==c.default.prototype?"symbol":typeof t};e.default="function"==typeof c.default&&"symbol"===s(i.default)?function(t){return void 0===t?"undefined":s(t)}:function(t){return t&&"function"==typeof c.default&&t.constructor===c.default&&t!==c.default.prototype?"symbol":void 0===t?"undefined":s(t)}},function(t,e,n){"use strict";var o=n(19),r=n(2),i=n(20),u=n(8),c=n(10),s=n(11),a=n(46),f=n(12),l=n(0).getProto,p=n(3)("iterator"),d=!([].keys&&"next"in[].keys()),y=function(){return this};t.exports=function(t,e,n,h,m,v,g){a(n,e,h);var b,_,x=function(t){if(!d&&t in S)return S[t];switch(t){case"keys":case"values":return function(){return new n(this,t)}}return function(){return new n(this,t)}},w=e+" Iterator",O="values"==m,k=!1,S=t.prototype,E=S[p]||S["@@iterator"]||m&&S[m],j=E||x(m);if(E){var P=l(j.call(new t));f(P,w,!0),!o&&c(S,"@@iterator")&&u(P,p,y),O&&"values"!==E.name&&(k=!0,j=function(){return E.call(this)})}if(o&&!g||!d&&!k&&S[p]||u(S,p,j),s[e]=j,s[w]=y,m)if(b={values:O?j:x("values"),keys:v?j:x("keys"),entries:O?x("entries"):j},g)for(_ in b)_ in S||i(S,_,b[_]);else r(r.P+r.F*(d||k),e,b);return b}},function(t,e){t.exports=!0},function(t,e,n){t.exports=n(8)},function(t,e,n){t.exports=!n(5)(function(){return 7!=Object.defineProperty({},"a",{get:function(){return 7}}).a})},function(t,e,n){var o=n(4),r=o["__core-js_shared__"]||(o["__core-js_shared__"]={});t.exports=function(t){return r[t]||(r[t]={})}},function(t,e){var n=0,o=Math.random();t.exports=function(t){return"Symbol(".concat(void 0===t?"":t,")_",(++n+o).toString(36))}},function(t,e,n){var o=n(25);t.exports=function(t){if(!o(t))throw TypeError(t+" is not an object!");return t}},function(t,e){t.exports=function(t){return"object"==typeof t?null!==t:"function"==typeof t}},function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),n.d(e,"EntityList",function(){return _});var o=n(27),r=n.n(o),i=n(32),u=n.n(i),c=n(36),s=n.n(c),a=n(37),f=n.n(a),l=n(40),p=n.n(l),d=n(59),y=n.n(d),h=n(66),m=n.n(h),v=n(67),g=(n.n(v),n(68)),b=(n.n(g),n(69)),_=function(t){function e(t,n){s()(this,e);var o=p()(this,(e.__proto__||u()(e)).call(this,t));x.call(o);var r=t.input.value?t.input.value:[];console.log("props...i entity list",t),o.label=t.field.label?t.field.label:t.field.entity;var i=t.field.form?t.field.form:"new_form_"+t.field.entity.toLowerCase();o.formDesc={type:"form",id:i},o.uikit=n.uikit;return o.state={items:r,formOpen:!1},o}return y()(e,t),f()(e,[{key:"componentWillReceiveProps",value:function(t){var e=t.input.value?t.input.value:[];this.setState(r()({},this.state,{items:e}))}},{key:"render",value:function(){var t=[];console.log("rendering items in entity list",this.props,this.state);var e=this;this.state.items.forEach(function(n,o){var r=function(){e.removeItem(n,o)},i=function(){e.editItem(n,o)},u=e.props.entityText?e.props.entityText:"Name",c=n[u];c=c||n.Title,t.push(m.a.createElement(e.uikit.Block,{className:"row between-xs"},m.a.createElement(e.uikit.Block,{className:"left"},c),m.a.createElement(e.uikit.Block,{className:"right"},m.a.createElement(v.Action,{action:{actiontype:"method"},className:"edit p10",method:i},m.a.createElement(e.uikit.Icons.EditIcon,null)),m.a.createElement(v.Action,{action:{actiontype:"method"},className:"remove p10",method:r},m.a.createElement(e.uikit.Icons.DeleteIcon,null)))))});return this.state.formOpen&&this.inlineRow&&t.push(this.inlineRow),0==t.length&&t.push("No data"),m.a.createElement(this.uikit.Block,{className:"entitylistfield "+this.label},this.props.field.skipLabel?null:m.a.createElement(this.uikit.Block,{className:"title"},this.label),m.a.createElement(this.uikit.Block,{className:"right tb10"},m.a.createElement(v.Action,{name:"listfield_new_entity",className:"p10",method:this.openForm},m.a.createElement(this.uikit.Icons.NewIcon,null))),t)}}]),e}(m.a.Component),x=function(){var t=this;this.closeForm=function(){switch(console.log("closing form"),t.props.field.mode){case"inline":t.inlineRow=null;break;case"dialog":Window.closeDialog();break;case"overlay":default:t.context.overlayComponent&&t.context.overlayComponent(null)}t.setState(r()({},t.state,{formOpen:!1}))},this.actions=function(e,n,o){return console.log("actios returned",e,n,o),m.a.createElement(t.uikit.Block,{className:"right p20"},t.props.field.inline?null:m.a.createElement(t.uikit.ActionButton,{onClick:o},"Reset"),m.a.createElement(t.uikit.ActionButton,{onClick:n},"Add"))},this.edit=function(e,n,o,r){var i=t.state.items;i&&i.length>n&&(i[n]=e.data,console.log(" items",i),t.props.input.onChange(i),t.closeForm())},this.add=function(e,n,o){var r=t.addItem(e.data);console.log(" items",r),t.props.input.onChange(r),t.closeForm()},this.openForm=function(e,n){console.log("opened form",t.props,t.context);var o=t,i=e?function(t,e,r){return o.edit(t,n,e,r)}:t.add,u=m.a.createElement(g.Panel,{actions:t.actions,inline:!0,formData:e,title:"Add "+t.label,closePanel:t.closeForm,onSubmit:i,description:t.formDesc});switch(t.props.field.mode){case"inline":t.inlineRow=u;break;case"dialog":Window.showDialog(null,u,t.closeForm);break;case"overlay":default:t.context.overlayComponent&&t.context.overlayComponent(u)}t.setState(r()({},t.state,{formOpen:!0}))},this.addItem=function(e){var n=t.state.items.slice();return n.push(e),n},this.editItem=function(e,n){t.openForm(e,n)},this.removeItem=function(e,n){var o=t.state.items.slice();n>-1&&o.splice(n,1),t.props.input.onChange(o)}};_.contextTypes={uikit:b.object,overlayComponent:b.func}},function(t,e,n){t.exports={default:n(28),__esModule:!0}},function(t,e,n){n(29),t.exports=n(1).Object.assign},function(t,e,n){var o=n(2);o(o.S+o.F,"Object",{assign:n(31)})},function(t,e){t.exports=function(t){if("function"!=typeof t)throw TypeError(t+" is not a function!");return t}},function(t,e,n){var o=n(0),r=n(14),i=n(15);t.exports=n(5)(function(){var t=Object.assign,e={},n={},o=Symbol(),r="abcdefghijklmnopqrst";return e[o]=7,r.split("").forEach(function(t){n[t]=t}),7!=t({},e)[o]||Object.keys(t({},n)).join("")!=r})?function(t,e){for(var n=r(t),u=arguments,c=u.length,s=1,a=o.getKeys,f=o.getSymbols,l=o.isEnum;c>s;)for(var p,d=i(u[s++]),y=f?a(d).concat(f(d)):a(d),h=y.length,m=0;h>m;)l.call(d,p=y[m++])&&(n[p]=d[p]);return n}:Object.assign},function(t,e,n){t.exports={default:n(33),__esModule:!0}},function(t,e,n){n(34),t.exports=n(1).Object.getPrototypeOf},function(t,e,n){var o=n(14);n(35)("getPrototypeOf",function(t){return function(e){return t(o(e))}})},function(t,e,n){var o=n(2),r=n(1),i=n(5);t.exports=function(t,e){var n=(r.Object||{})[t]||Object[t],u={};u[t]=e(n),o(o.S+o.F*i(function(){n(1)}),"Object",u)}},function(t,e,n){"use strict";e.__esModule=!0,e.default=function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}},function(t,e,n){"use strict";e.__esModule=!0;var o=n(38),r=function(t){return t&&t.__esModule?t:{default:t}}(o);e.default=function(){function t(t,e){for(var n=0;n<e.length;n++){var o=e[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),(0,r.default)(t,o.key,o)}}return function(e,n,o){return n&&t(e.prototype,n),o&&t(e,o),e}}()},function(t,e,n){t.exports={default:n(39),__esModule:!0}},function(t,e,n){var o=n(0);t.exports=function(t,e,n){return o.setDesc(t,e,n)}},function(t,e,n){"use strict";e.__esModule=!0;var o=n(17),r=function(t){return t&&t.__esModule?t:{default:t}}(o);e.default=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!==(void 0===e?"undefined":(0,r.default)(e))&&"function"!=typeof e?t:e}},function(t,e,n){t.exports={default:n(42),__esModule:!0}},function(t,e,n){n(43),n(47),t.exports=n(3)("iterator")},function(t,e,n){"use strict";var o=n(44)(!0);n(18)(String,"String",function(t){this._t=String(t),this._i=0},function(){var t,e=this._t,n=this._i;return n>=e.length?{value:void 0,done:!0}:(t=o(e,n),this._i+=t.length,{value:t,done:!1})})},function(t,e,n){var o=n(45),r=n(7);t.exports=function(t){return function(e,n){var i,u,c=String(r(e)),s=o(n),a=c.length;return s<0||s>=a?t?"":void 0:(i=c.charCodeAt(s),i<55296||i>56319||s+1===a||(u=c.charCodeAt(s+1))<56320||u>57343?t?c.charAt(s):i:t?c.slice(s,s+2):u-56320+(i-55296<<10)+65536)}}},function(t,e){var n=Math.ceil,o=Math.floor;t.exports=function(t){return isNaN(t=+t)?0:(t>0?o:n)(t)}},function(t,e,n){"use strict";var o=n(0),r=n(9),i=n(12),u={};n(8)(u,n(3)("iterator"),function(){return this}),t.exports=function(t,e,n){t.prototype=o.create(u,{next:r(1,n)}),i(t,e+" Iterator")}},function(t,e,n){n(48);var o=n(11);o.NodeList=o.HTMLCollection=o.Array},function(t,e,n){"use strict";var o=n(49),r=n(50),i=n(11),u=n(6);t.exports=n(18)(Array,"Array",function(t,e){this._t=u(t),this._i=0,this._k=e},function(){var t=this._t,e=this._k,n=this._i++;return!t||n>=t.length?(this._t=void 0,r(1)):"keys"==e?r(0,n):"values"==e?r(0,t[n]):r(0,[n,t[n]])},"values"),i.Arguments=i.Array,o("keys"),o("values"),o("entries")},function(t,e){t.exports=function(){}},function(t,e){t.exports=function(t,e){return{value:e,done:!!t}}},function(t,e,n){t.exports={default:n(52),__esModule:!0}},function(t,e,n){n(53),n(58),t.exports=n(1).Symbol},function(t,e,n){"use strict";var o=n(0),r=n(4),i=n(10),u=n(21),c=n(2),s=n(20),a=n(5),f=n(22),l=n(12),p=n(23),d=n(3),y=n(54),h=n(55),m=n(56),v=n(57),g=n(24),b=n(6),_=n(9),x=o.getDesc,w=o.setDesc,O=o.create,k=h.get,S=r.Symbol,E=r.JSON,j=E&&E.stringify,P=!1,M=d("_hidden"),N=o.isEnum,T=f("symbol-registry"),D=f("symbols"),I="function"==typeof S,A=Object.prototype,F=u&&a(function(){return 7!=O(w({},"a",{get:function(){return w(this,"a",{value:7}).a}})).a})?function(t,e,n){var o=x(A,e);o&&delete A[e],w(t,e,n),o&&t!==A&&w(A,e,o)}:w,C=function(t){var e=D[t]=O(S.prototype);return e._k=t,u&&P&&F(A,t,{configurable:!0,set:function(e){i(this,M)&&i(this[M],t)&&(this[M][t]=!1),F(this,t,_(1,e))}}),e},R=function(t){return"symbol"==typeof t},B=function(t,e,n){return n&&i(D,e)?(n.enumerable?(i(t,M)&&t[M][e]&&(t[M][e]=!1),n=O(n,{enumerable:_(0,!1)})):(i(t,M)||w(t,M,_(1,{})),t[M][e]=!0),F(t,e,n)):w(t,e,n)},W=function(t,e){g(t);for(var n,o=m(e=b(e)),r=0,i=o.length;i>r;)B(t,n=o[r++],e[n]);return t},L=function(t,e){return void 0===e?O(t):W(O(t),e)},J=function(t){var e=N.call(this,t);return!(e||!i(this,t)||!i(D,t)||i(this,M)&&this[M][t])||e},K=function(t,e){var n=x(t=b(t),e);return!n||!i(D,e)||i(t,M)&&t[M][e]||(n.enumerable=!0),n},G=function(t){for(var e,n=k(b(t)),o=[],r=0;n.length>r;)i(D,e=n[r++])||e==M||o.push(e);return o},q=function(t){for(var e,n=k(b(t)),o=[],r=0;n.length>r;)i(D,e=n[r++])&&o.push(D[e]);return o},H=function(t){if(void 0!==t&&!R(t)){for(var e,n,o=[t],r=1,i=arguments;i.length>r;)o.push(i[r++]);return e=o[1],"function"==typeof e&&(n=e),!n&&v(e)||(e=function(t,e){if(n&&(e=n.call(this,t,e)),!R(e))return e}),o[1]=e,j.apply(E,o)}},U=a(function(){var t=S();return"[null]"!=j([t])||"{}"!=j({a:t})||"{}"!=j(Object(t))});I||(S=function(){if(R(this))throw TypeError("Symbol is not a constructor");return C(p(arguments.length>0?arguments[0]:void 0))},s(S.prototype,"toString",function(){return this._k}),R=function(t){return t instanceof S},o.create=L,o.isEnum=J,o.getDesc=K,o.setDesc=B,o.setDescs=W,o.getNames=h.get=G,o.getSymbols=q,u&&!n(19)&&s(A,"propertyIsEnumerable",J,!0));var z={for:function(t){return i(T,t+="")?T[t]:T[t]=S(t)},keyFor:function(t){return y(T,t)},useSetter:function(){P=!0},useSimple:function(){P=!1}};o.each.call("hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","),function(t){var e=d(t);z[t]=I?e:C(e)}),P=!0,c(c.G+c.W,{Symbol:S}),c(c.S,"Symbol",z),c(c.S+c.F*!I,"Object",{create:L,defineProperty:B,defineProperties:W,getOwnPropertyDescriptor:K,getOwnPropertyNames:G,getOwnPropertySymbols:q}),E&&c(c.S+c.F*(!I||U),"JSON",{stringify:H}),l(S,"Symbol"),l(Math,"Math",!0),l(r.JSON,"JSON",!0)},function(t,e,n){var o=n(0),r=n(6);t.exports=function(t,e){for(var n,i=r(t),u=o.getKeys(i),c=u.length,s=0;c>s;)if(i[n=u[s++]]===e)return n}},function(t,e,n){var o=n(6),r=n(0).getNames,i={}.toString,u="object"==typeof window&&Object.getOwnPropertyNames?Object.getOwnPropertyNames(window):[],c=function(t){try{return r(t)}catch(t){return u.slice()}};t.exports.get=function(t){return u&&"[object Window]"==i.call(t)?c(t):r(o(t))}},function(t,e,n){var o=n(0);t.exports=function(t){var e=o.getKeys(t),n=o.getSymbols;if(n)for(var r,i=n(t),u=o.isEnum,c=0;i.length>c;)u.call(t,r=i[c++])&&e.push(r);return e}},function(t,e,n){var o=n(16);t.exports=Array.isArray||function(t){return"Array"==o(t)}},function(t,e){},function(t,e,n){"use strict";function o(t){return t&&t.__esModule?t:{default:t}}e.__esModule=!0;var r=n(60),i=o(r),u=n(64),c=o(u),s=n(17),a=o(s);e.default=function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+(void 0===e?"undefined":(0,a.default)(e)));t.prototype=(0,c.default)(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(i.default?(0,i.default)(t,e):t.__proto__=e)}},function(t,e,n){t.exports={default:n(61),__esModule:!0}},function(t,e,n){n(62),t.exports=n(1).Object.setPrototypeOf},function(t,e,n){var o=n(2);o(o.S,"Object",{setPrototypeOf:n(63).set})},function(t,e,n){var o=n(0).getDesc,r=n(25),i=n(24),u=function(t,e){if(i(t),!r(e)&&null!==e)throw TypeError(e+": can't set as prototype!")};t.exports={set:Object.setPrototypeOf||("__proto__"in{}?function(t,e,r){try{r=n(13)(Function.call,o(Object.prototype,"__proto__").set,2),r(t,[]),e=!(t instanceof Array)}catch(t){e=!0}return function(t,n){return u(t,n),e?t.__proto__=n:r(t,n),t}}({},!1):void 0),check:u}},function(t,e,n){t.exports={default:n(65),__esModule:!0}},function(t,e,n){var o=n(0);t.exports=function(t,e){return o.create(t,e)}},function(e,n){e.exports=t},function(t,n){t.exports=e},function(t,e){t.exports=n},function(t,e,n){t.exports=n(70)()},function(t,e,n){"use strict";var o=n(71),r=n(72),i=n(73);t.exports=function(){function t(t,e,n,o,u,c){c!==i&&r(!1,"Calling PropTypes validators directly is not supported by the `prop-types` package. Use PropTypes.checkPropTypes() to call them. Read more at http://fb.me/use-check-prop-types")}function e(){return t}t.isRequired=t;var n={array:t,bool:t,func:t,number:t,object:t,string:t,symbol:t,any:t,arrayOf:e,element:t,instanceOf:e,node:t,objectOf:e,oneOf:e,oneOfType:e,shape:e,exact:e};return n.checkPropTypes=o,n.PropTypes=n,n}},function(t,e,n){"use strict";function o(t){return function(){return t}}var r=function(){};r.thatReturns=o,r.thatReturnsFalse=o(!1),r.thatReturnsTrue=o(!0),r.thatReturnsNull=o(null),r.thatReturnsThis=function(){return this},r.thatReturnsArgument=function(t){return t},t.exports=r},function(t,e,n){"use strict";function o(t,e,n,o,i,u,c,s){if(r(e),!t){var a;if(void 0===e)a=new Error("Minified exception occurred; use the non-minified dev environment for the full error message and additional helpful warnings.");else{var f=[n,o,i,u,c,s],l=0;a=new Error(e.replace(/%s/g,function(){return f[l++]})),a.name="Invariant Violation"}throw a.framesToPop=1,a}}var r=function(t){};t.exports=o},function(t,e,n){"use strict";t.exports="SECRET_DO_NOT_PASS_THIS_OR_YOU_WILL_BE_FIRED"}])});
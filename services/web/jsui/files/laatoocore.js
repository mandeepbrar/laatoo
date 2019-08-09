let n= (typeof(document)==='undefined');
var Storage = n ? {} : localStorage;
var Application = n ? { Native:n} : document.InitConfig;
Application.Registry={};
Application.Modules={};
var Window = n ? {} : window;
Application.Route={Location: window.location.href, Params: _routeParams};
Application.RegisterModule = function(mod) {
  console.log("register module", mod)
  try {
    let m = Application.Modules[mod] = require(mod);
    return m;
  } catch(ex) {
    console.log("Module could not be registered", ex);
  }
}
Application.Register = function(regName,id,data) {
  let reg=Application.Registry[regName];
  if(reg==null){
    reg={};
    Application.Registry[regName]=reg;
  }
  reg[id]=data;
}
Application.AllRegItems = function(regName) {
  return Application.Registry[regName];
}
Application.GetRegistry = function(regName, id) {
  if(id) {
    let reg = Application.Registry[regName];
    if(reg) {
      return reg[id];
    }
  }
  return null;
}
Application.Resolve = function(mod, comp) {
  let moduleObj
  if(!mod) {
    moduleObj = Application.uikit
  } else {
    moduleObj = _$[mod];
  }
  if(moduleObj && comp) {
    return moduleObj[comp]
  }
}


var _$=Application.Modules;
var _rm = Application.RegisterModule;
var _r = Application.Register;
var _reg= Application.GetRegistry;
var _res = Application.Resolve;
var _re = null;
var _ce = null;
var _uikit = null;

Application.setUikit = function(uikit) {
  Application.uikit = uikit;
  _uikit = uikit;
}
Application.setRouter = function(router) {
  Application.router = router;
}
function createElement(elem, props, children, name) {
  if(elem) {
    return _re.createElement(elem, props, children);
  }
  else {
    console.log("Element is null ", (name? name: ""));
  }
  return null
}
function Base64Decoder(){
  var chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
  // Use a lookup table to find the index.
  var lookup = new Uint8Array(256);
  for (var i = 0; i < chars.length; i++) {
    lookup[chars.charCodeAt(i)] = i;
  }
  return function(base64) {
    var bufferLength = base64.length * 0.75,
    len = base64.length, i, p = 0,
    encoded1, encoded2, encoded3, encoded4;

    if (base64[base64.length - 1] === "=") {
      bufferLength--;
      if (base64[base64.length - 2] === "=") {
        bufferLength--;
      }
    }

    var arraybuffer = new ArrayBuffer(bufferLength),
    bytes = new Uint8Array(arraybuffer);

    for (i = 0; i < len; i+=4) {
      encoded1 = lookup[base64.charCodeAt(i)];
      encoded2 = lookup[base64.charCodeAt(i+1)];
      encoded3 = lookup[base64.charCodeAt(i+2)];
      encoded4 = lookup[base64.charCodeAt(i+3)];

      bytes[p++] = (encoded1 << 2) | (encoded2 >> 4);
      bytes[p++] = ((encoded2 & 15) << 4) | (encoded3 >> 2);
      bytes[p++] = ((encoded3 & 3) << 6) | (encoded4 & 63);
    }

    return arraybuffer;
  };
};

Application.Base64Decoder = Base64Decoder();

Application.LoadWasm = function(mod, str, memory) {
  try {
    var decodedMod = Application.Base64Decoder(str);
    let jsMod = wasmBGImports(mod);
    Application.Modules[mod+"_wasm"] = jsMod;//wasmModule.instance.exports;
    let wasmJSImports = {};
    wasmJSImports['./'+mod] = jsMod;
    console.log("wasm imports", mod, wasmJSImports, jsMod);
    let importObject = Object.assign({}, Application.Modules, wasmJSImports, {env: {memory: memory}});
    return WebAssembly.instantiate(decodedMod, importObject).then(wasmModule => {
      jsMod.__wasmInit(wasmModule.instance.exports);
    });
  }catch(ex) {
    console.log("exception in instantiating wasm", mod, ex);
  }
}

console.log("Application", Application);
var _val = function() {
  var obj;
  var arr = arguments;
  for(var i=0; i< arr.length; i++) {
    if(i==0) {
      obj=arguments[i]
    } else {
      let ind = arr[i];
      if(!obj[ind]) {
        return "";
      } else {
        if(i<arr.length -1) {
          obj = obj[ind]
        } else {
          return obj[ind]
        }
      }
    }
  }
  return "";
}

function modDef(appname, ins, mod, settings) {
  define(ins, [mod], function (m) {
    if(m.Initialize) {
      m.Initialize(appname, ins, mod,  settings, define, require);
    }
    return m;
  });
}

function appLoadingComplete(appname, propsurl, modsToInitialize, wasmURL) {
  console.log("appLoadingComplete", wasmURL);
  if(wasmURL) {
    const memory = new WebAssembly.Memory({initial: 20});
    fetch(wasmURL).then(function(resp) {
      resp.json().then(function(wasmURLArr) {
        let promisArr = new Array();
        wasmURLArr.forEach(function(modItem){
          console.log("loading wasm ", modItem);
          let p = Application.LoadWasm(modItem.Name, modItem.Data, memory);
          promisArr.push(p);
        });  
        Promise.all(promisArr).then(values => {
          initMods(appname, propsurl, modsToInitialize);
        });
      });
    });
  } else {
    initMods(appname, propsurl, modsToInitialize);
  }
}
function initMods(appname, propsurl, modsToInitialize) {
  var init = function() {
    console.log("Initializing application", modsToInitialize);
    _re=require('react');
    _ce=createElement;
    if(modsToInitialize!=null) {
      for(var i=0;i<modsToInitialize.length;i++) {
        var row = modsToInitialize[i];
        console.log("Initializing module", row);
        if(row[0]!=row[1]){
          modDef(appname, row[0], row[1], row[2]);
        }
        let k = _rm(row[1]);
        if(!k){
          console.log("Could not Initialize module ", row[0]);
        } else {
          if(k.Initialize){
            k.Initialize(appname,row[0], row[1], row[2], define, require);
          }
        }
      }
    }
    console.log("Initialized modules", _$);
    Window.InitializeApplication();
  }
  
  if(propsurl) {
    propsurl = window.location.origin + propsurl;
    wasmTest(appname, propsurl, modsToInitialize);
    fetch(propsurl).then(function(resp) {
      resp.json().then(function(data) {
        console.log("props fetched", propsurl, data);
        Application.Properties=data;
        init();
      });
    });
  } else {
    init();
  }
}

function wasmTest(appname, propsurl, modsToInitialize) {
  /*console.log("all modules", _$);
  let lb = _$["laatoobrowser_wasm"];
  console.log("wasm test ", _$, lb);
  let app = lb.initialize();
  console.log("app", app);
  let ws = _$["websocket_wasm"];
  console.log("websocket", ws);
  ws.init(app);*/
}

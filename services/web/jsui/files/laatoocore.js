let n= (typeof(document)==='undefined');
var Storage = n ? {} : localStorage;
var Application = n ? { Native:n} : document.InitConfig;
Application.Registry={};
Application.Modules={};
var Window = n ? {} : window;
Application.RegisterModule = function(mod) {
  console.log("register module", mod)
  try {
    let m = Application.Modules[mod] = require(mod);
    console.log("returned module", mod, m)
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

Application.LoadWasm = function(mod, str) {
  try {
    var decodedMod = Application.Base64Decoder(str);
    let wasmImports = {}
    let wasmExp = wasmBGImports();
    wasmImports['./'+mod] = wasmExp;
    let importObject = Object.assign({}, Application.Modules, wasmImports);
    WebAssembly.instantiate(decodedMod, importObject).then(wasmModule => {
      Application.Modules[mod] = wasmModule.instance.exports;
      wasmExp.__wasmInit(wasmModule.instance.exports);
    });
  }catch(ex) {
    console.log("exception in instantiating wasm", mod, ex);
  }
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
var _$=Application.Modules;
var _rm = Application.RegisterModule;
var _r = Application.Register;
var _reg= Application.GetRegistry;
var _re= null;
var _ce= null;
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
  var init = function() {
    console.log("Initializing application", modsToInitialize);
    _re=require('react');
    _ce=_re.createElement;
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
  
  if(wasmURL) {
    fetch(wasmURL).then(function(resp) {
      resp.json().then(function(wasmURLArr) {
        wasmURLArr.forEach(function(modItem){
          Application.LoadWasm(modItem.Name, modItem.Data);
        });  
      });
    });
  }

  if(propsurl) {
    propsurl = window.location.origin + propsurl;
    fetch(propsurl).then(function(resp) {
      resp.json().then(function(data) {
        console.log("props fetched", propsurl, data);
        Application.Properties=data;
        init();
      });
    });
  } else {
    init()
  }
}

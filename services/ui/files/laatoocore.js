let n= (typeof(document)==='undefined');
var Storage = n ? {} : localStorage;
var Application = n ? { Native:n} : document.InitConfig;
Application.Registry={};
Application.Modules={};
var Window = n ? {} : window;
Application.RegisterModule = function(mod) {
  console.log("register module", mod)
  let m = Application.Modules[mod] = require(mod);
  console.log("returned module", mod, m)
  return m;
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
function appLoadingComplete(appname, propsurl, modsToInitialize) {
  console.log("appLoadingComplete");
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

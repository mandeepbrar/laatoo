let n= (typeof(document)==='undefined');
var Storage = n ? {} : localStorage;
var Application = n ? { Native:n} : document.InitConfig;
Application.Registry={};
var Window = n ? {} : window;
Application.Register = function(regName,id,data) {
  let reg=Application.Registry[regName];
  if(reg==null){
    reg={};
    Application.Registry[regName]=reg;
  }
  reg[id]=data;
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
  var init = function() {
    if(modsToInitialize!=null) {
      for(var i=0;i<modsToInitialize.length;i++) {
        var row = modsToInitialize[i];
        modDef(appname, row[0], row[1], row[2]);
      }
    }
    Window.InitializeApplication();
  }
  if(propsurl) {
    propsurl = window.location.origin + propsurl;
    fetch(propsurl).then(function(resp) {
      resp.json().then(function(data) {
        document.InitConfig.Properties=data;
        init();
      });
    });
  } else {
    init()
  }
}

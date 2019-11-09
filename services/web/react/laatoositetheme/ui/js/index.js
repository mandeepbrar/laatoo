import React from 'react'
import {SiteTheme, PreprocessPageComponents, setModule as initTheme} from './Theme'
import {RenderPageComponent, setModule as initPage} from './SitePage'
import Welcome from './Welcome'
import './styles/app.scss'

var module;

function Initialize(appName, ins, mod, settings, def, req) {
  console.log("appname = ", appName, "ins ", ins, "mod", mod, "settings", settings)
  module =  this;
  initPage(module);
  initTheme(module);
  module.properties = Application.Properties[ins];
  module.settings = settings;
  module.authenticate = settings.authenticate
  let homePage = _reg('Pages', 'home')
  if(!homePage) {
    let homePage={id:"home", route:"/", components: {"main": {type:"component", component: <Welcome modProps={module.properties}/>}}}
    Application.Register('Pages', "home", homePage);
    Application.Register('Actions','Page_home', {url:'/'})
  }
  let loginModule = "authui"
  let loginComp = "WebLoginForm"
  if(settings && settings.loginModule) {
      loginModule = settings.loginModule
      loginComp = settings.loginComp
  }
  let loginMod = req(loginModule)
  if(loginMod) {
    module.logInComp = loginMod[loginComp]
  }
  if(module.settings.showMenu) {
    if(module.settings.menu) {
      module.menu = module.settings.menu
    } else {
      module.menu = Application.Properties.menu
    }
  }
}

export {
  Initialize ,
  PreprocessPageComponents,
  RenderPageComponent,
  SiteTheme as Theme
}

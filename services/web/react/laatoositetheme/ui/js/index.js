import React from 'react'
import {SiteTheme, PreprocessPageComponents, setModule as initTheme} from './Theme'
import {RenderPageComponent, setModule as initPage} from './SitePage'
import Welcome from './Welcome'
import {Panel} from 'reactpages'
import './styles/app.scss'

var module;

function Initialize(appName, ins, mod, settings, def, req) {
  console.log("appname = ", appName, "ins ", ins, "mod", mod, "settings", settings)
  module =  this;
  initPage(module);
  initTheme(module);
  module.properties = Application.Properties[ins];
  module.settings = settings;
  module.skipAuth = settings.skipAuth
  let homePage = _reg('Pages', 'home')
  if(!homePage) {
    let homePage={id:"home", route:"/", components: {"main": {type:"component", component: <Welcome modProps={module.properties}/>}}}
    Application.Register('Pages', "home", homePage);
    Application.Register('Actions','Page_home', {url:'/'})
  }
  let loginModule = "authui"
  let loginComp = "WebLoginForm"
  if(settings && !settings.skipAuth && settings.loginModule) {
      loginModule = settings.loginModule
      loginComp = settings.loginComp
  }
  let loginMod = req(loginModule)
  if(loginMod) {
    module.logInComp = loginMod[loginComp]
  }
  if(module.settings.showMenu) {
    processMenu()
  }
}

function processMenu(){
  var menuItems=[]
  let menu = []
  if(module.settings && module.settings.menu) {
    menu = module.settings.menu
  } else {
    menu = Application.Properties.menu
  }
  console.log("dashboard menu", menu)
  if(menu) {
    menu.forEach(function(menuItem){
      //let menuItem=menuConfig[menuItem]
      if(menuItem.page) {
        menuItems.push({title:menuItem.title, action: "Page_" + menuItem.page})
      } else {
        menuItems.push({title:menuItem.title, action: menuItem.action})
      }
    })
  } else {
    menuItems.push({title:'Home', action:'Page_home'})
  }
  module.menu = menuItems
}


export {
  Initialize ,
  PreprocessPageComponents,
  RenderPageComponent,
  SiteTheme as Theme
}

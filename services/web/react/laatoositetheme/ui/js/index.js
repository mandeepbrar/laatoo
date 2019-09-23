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
    processMenu()
  }
}

function processMenu(){
  //var menuItems=[]
  let menuName
  let menu
  console.log("Menu process", module.settings, Application)
  if(module.settings && module.settings.menu) {
    menuName = module.settings.menu
  } else {
    menuName = Application.Properties.menu
  }
  console.log("dashboard menu", menuName)
  if(menuName) {
    menu = _reg("Menus", menuName)
    console.log("received menu ", menuName, menu)
    if(menu && menu.items) {
      console.log("iterating menu items ", menu.items)
      menu.items.forEach(function(menuItem){
        console.log("iterating menu ", menuItem)
        //let menuItem=menuConfig[menuItem]
        if(menuItem.page) {
          menuItem.action = "Page_" + menuItem.page
        }
      })  
    }
  } else {
    menu = {items: []}
    menu.items.push({title:'Home', action:'Page_home'})
  }
  console.log("setting module menu", module, menu)
  module.menu = menu
}


export {
  Initialize ,
  PreprocessPageComponents,
  RenderPageComponent,
  SiteTheme as Theme
}

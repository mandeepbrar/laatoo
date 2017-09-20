import React from 'react'
import Header from './Header'
import Welcome from './Welcome'
import './styles/app.scss'
import {LoginValidator, LoginForm} from 'authui';

var module = this;

function Initialize(appName, ins, mod, settings, def, req) {
  module.properties = Application.Properties[ins];
  module.settings = settings;
  console.log("Initialize dashboardtheme", module, Application);
  module.skipAuth = settings.skipAuth
  if(!Application.Registry.Pages || !Application.Registry.Page['home']) {
    let homePage={id:"home", pattern:"", components: {"main": <Welcome modProps={module.properties}/>}}
    Application.Register('Pages', "home", homePage);
    Application.Register('Actions','Page_home', {url:''})
  }
  let loginModule = "reactforms"
  let loginComp = "LoginForm"
  if(settings && !settings.skipAuth && settings.loginModule) {
      loginModule = settings.loginModule
      loginComp = settings.loginComp
  }
  let loginMod = req(loginModule)
  if(loginMod) {
    module.logInComp = loginMod[loginComp]
  }
  console.log("login comp", module.logInComp, "mod", loginMod)
}

function processPage(menus, page, uikit){
  let pageRoute = {pattern: page.pattern, reducer: page.reducer}
  pageRoute.components = {
    "menu":(routerState) => {
      return (
        <uikit.Navbar items={menus} vertical={module.settings.vertical}/>
      )
    },
    "main":(routerState) => {
      console.log(page.components);
      return (
        <div className="main"  vertical={module.settings.vertical}>{page.components["main"]}</div>
      )
    }
  }
  Application.Register('Routes', page.id, pageRoute)
}

function Start(appName, uikit){
  let menus=[]
  let menuConfig = {}
  if(module.settings && module.settings.menu) {
    menuConfig = module.settings.menu
  } else {
    menuConfig = Application.Properties.menu
  }
  if(menuConfig) {
    Object.keys(menuConfig).forEach(function(key){
      let menuItem=menuConfig[key]
      menus.push({title:menuItem.title, action: "Page_" + menuItem.page})
    })
  } else {
    menus.push({title:'Home', action:'Page_home'})
  }
  let pages=Application.Registry.Pages
  if(pages!=null) {
    for(var pageId in pages) {
      console.log("processing page", pageId)
      processPage(menus, pages[pageId], uikit)
    }
  }
  console.log("started app", Application.Registry.Routes)
}

const MainView = (props) => {
  let vertical = module.settings.vertical?true:false
  let logInComp = module.logInComp
  let loggedInComp = (
    <div className="body">
      <div className={vertical?"vertmenu":"horizmenu"}>
        <props.router.View name="menu"  />
      </div>
      <div className={vertical?"vertbody":"horizbody"}>
        <props.router.View name="main"/>
      </div>
    </div>
  )
  let retval= (props.loggedIn || module.skipAuth)?loggedInComp:<div className="dashlogin"><module.logInComp/></div>;
  console.log("main view of dashboard theme", props, retval)
  return retval
}

const DashboardTheme = (props) => {
  console.log("uikit initializer", props.uikit, props.uikit.UIWrapper)
  let loggedIn = false
  return (
    <props.uikit.UIWrapper>
      <div className={module.settings.className?module.settings.className + ' dashboard':'dashboard'}>
        <Header headerProps={module.properties.header} />
        <LoginValidator  validateService="validatetoken">
          <MainView router={props.router}/>
        </LoginValidator>
        <div className="footer">
        </div>
      </div>
    </props.uikit.UIWrapper>
  )
}

export {
  Initialize ,
  Start,
  DashboardTheme as Theme
}

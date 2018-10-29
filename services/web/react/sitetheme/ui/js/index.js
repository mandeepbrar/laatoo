import React from 'react'
import Header from './Header'
import Welcome from './Welcome'
import './styles/app.scss'
import {LoginValidator, LoginForm} from 'authui';

var module;

function Initialize(appName, ins, mod, settings, def, req) {
  console.log("appname = ", appName, "ins ", ins, "mod", mod, "settings", settings)
  module =  this;
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
  processMenu()
}

/*function processPage(menus, page, uikit){
  let pageRoute = {pattern: page.pattern, reducer: page.reducer}
  pageRoute.components = {
    "menu":,
    "main":(routerState) => {
      console.log(page.components);
      return (
        <div className="main"  vertical={module.settings.vertical}>{page.components["main"]}</div>
      )
    }
  }
  Application.Register('Routes', page.id, pageRoute)
}*/

function processMenu(){
  var menuItems=[]
  let menu = []
  if(module.settings && module.settings.menu) {
    menu = module.settings.menu
  } else {
    menu = Application.Properties.menu
  }
  console.log("website menu", menu)
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

function ProcessRoute(route, uikit) {
  route.components.menu = (routerState) => {
    return (
      <uikit.Navbar items={module.menu} vertical={module.settings.vertical}/>
    )
  }
  return route
}

const MainView = (props) => {
  let vertical = module.settings.vertical?true:false
  let logInComp = module.logInComp
  let loggedInComp = (
    <props.uikit.Block className="body">
      <props.uikit.Block className={vertical?"vertmenu":"horizmenu"}>
        <props.router.View name="menu"  />
      </props.uikit.Block>
      <props.uikit.Block className={vertical?"vertbody":"horizbody"}>
        <props.router.View name="main"/>
      </props.uikit.Block>
    </props.uikit.Block>
  )
  let retval= (props.loggedIn || module.skipAuth)?loggedInComp:<props.uikit.Block className="dashlogin"><module.logInComp/></props.uikit.Block>;
  return retval
}

const WebsiteTheme = (props) => {
  let loggedIn = false
  return (
    <props.uikit.UIWrapper>
      <props.uikit.Block className={module.settings.className?module.settings.className + ' dashboard':'dashboard'}>
        <Header uikit={props.uikit} module={module} />
        <LoginValidator  validateService="validate">
          <MainView uikit={props.uikit} router={props.router}/>
        </LoginValidator>
        <props.uikit.Block className="footer">
        {props.children}
        </props.uikit.Block>
      </props.uikit.Block>
    </props.uikit.UIWrapper>
  )
}

export {
  Initialize ,
  ProcessRoute,
  WebsiteTheme as Theme
}

import React from 'react'
import Header from './Header'
import Welcome from './Welcome'
import './styles/app.scss'
import {LoginValidator, LoginForm} from 'authui';

var module = this;

function Initialize(appName, ins, mod, settings, def, req) {
  console.log("appname", appName, "ins ", ins, "mod", mod, "settings", settings)
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
  console.log("dashboard menu", menu)
  if(menu) {
    menu.forEach(function(menuItem){
      console.log("found menu item", menuItem)
      //let menuItem=menuConfig[menuItem]
      console.log("found menu", menuItem.title)
      if(menuItem.page) {
        menuItems.push({title:menuItem.title, action: "Page_" + menuItem.page})
      } else {
        menuItems.push({title:menuItem.title, action: menuItem.action})
      }
    })
  } else {
    menuItems.push({title:'Home', action:'Page_home'})
  }
  console.log("menu items", menuItems)
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
  return retval
}

const DashboardTheme = (props) => {
  let loggedIn = false
  return (
    <props.uikit.UIWrapper>
      <div className={module.settings.className?module.settings.className + ' dashboard':'dashboard'}>
        <Header headerProps={module.properties.header} />
        <LoginValidator  validateService="validate">
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
  ProcessRoute,
  DashboardTheme as Theme
}

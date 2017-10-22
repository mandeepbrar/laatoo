import React from 'react'
import Header from './Header'
import Welcome from './Welcome'
import './styles/app.scss'
import {LoginValidator, LoginForm} from 'authui';

var module = this;

function Initialize(appName, ins, mod, settings, def, req) {
  module.properties = Application.Properties[ins];
  module.settings = settings;
  module.skipAuth = settings.skipAuth
  if(!Application.Registry.Pages || !Application.Registry.Pages['home']) {
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
  let menu=[]
  let menuConfig = {}
  if(module.settings && module.settings.menu) {
    menuConfig = module.settings.menu
  } else {
    menuConfig = Application.Properties.menu
  }
  if(menuConfig) {
    Object.keys(menuConfig).forEach(function(key){
      let menuItem=menuConfig[key]
      console.log("found menu", key, menuItem)
      menu.push({title:menuItem.title, action: "Page_" + menuItem.page})
    })
  } else {
    menu.push({title:'Home', action:'Page_home'})
  }
  module.menu = menu
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

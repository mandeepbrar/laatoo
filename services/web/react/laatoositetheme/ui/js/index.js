import React from 'react'
import Header from './Header'
import Welcome from './Welcome'
import {Panel} from 'reactpages'
import './styles/app.scss'
import {LoginValidator, LoginForm} from 'authui';
const PropTypes = require('prop-types');

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

function PreprocessPageComponents(components, page, pageId, reducers, uikit) {
  console.log("pre processing.......", components);
  components.menu = (routerState) => {
    console.log("menu items ", routerState, module.menu)
    return (
      <uikit.Navbar items={module.menu} vertical={module.settings.vertical}/>
    )
  }
  return components;
}

const SiteMenu  = (props, context) => {
  console.log("rendering site menu", props, context)
  if(context.loggedIn) {
    return props.menu
  } else {
    return null
  }
}

SiteMenu.contextTypes = {
  routeParams: PropTypes.object,
  routerState: PropTypes.object,
  loggedIn: PropTypes.bool,
  user: PropTypes.object  
};

/*const SitePage  = (props) => {
  console.log("site props", props, module)
  let uikit = props.uikit;
  return (props.loggedIn || module.skipAuth)? props.pageComp:<props.uikit.Block className="dashlogin"><module.logInComp/></props.uikit.Block>;
}*/
//pageComp={comp} pageKey={key} pageId={pageId} routerState={routerState} page={page} uikit={uikit}
//<SitePage uikit={props.uikit} router={props.router}/>
const SitePage  = (props, context) => {
  console.log("rendering site page", props, context)
  if(context.loggedIn) {
    return <PageComponent pageId={props.pageId} placeholder={props.pageKey} routerState={props.routerState} description={props.pageComp} />
  } else {
    return <props.uikit.Block className="dashlogin"><module.logInComp/></props.uikit.Block>
  }
}

SitePage.contextTypes = {
  routeParams: PropTypes.object,
  routerState: PropTypes.object,
  loggedIn: PropTypes.bool,
  user: PropTypes.object  
};

const SiteTheme = (props, st, c) => {
  let vertical = module.settings.vertical?true:false
  return (
    <props.uikit.UIWrapper>
      <LoginValidator  validateService="validate">
          <props.uikit.Block className={module.settings.className?module.settings.className + ' dashboard':'dashboard'}>
            <Header uikit={props.uikit} module={module} />
            <props.uikit.Block className="body">
              <props.uikit.Block className={vertical?"vertmenu":"horizmenu"}>
                <props.router.View name="menu"  />
              </props.uikit.Block>
              <props.uikit.Block className={vertical?"vertbody":"horizbody"}>
                <props.router.View name="main"/>
              </props.uikit.Block>
            </props.uikit.Block>
            <props.uikit.Block className="footer">
            {props.children}
            </props.uikit.Block>
          </props.uikit.Block>
      </LoginValidator>
    </props.uikit.UIWrapper>
  )
}



function RenderPageComponent(comp, key, pageId, routerState, page, uikit) {
  console.log("RenderPageComponent", routerState, key, comp, page)
  if(key=="main") {
    return <SitePage pageComp={comp} pageKey={key} pageId={pageId} routerState={routerState} page={page} uikit={uikit}/>
  //  return null
  }
  /*if(key=="menu") {
//    return null
    return comp //<SiteMenu menu={comp} pageKey={key} pageId={pageId} routerState={routerState} page={page} uikit={uikit}/>
  }*/
}


class PageComponent extends React.Component {
  getChildContext() {
    return {routeParams: this.props.routerState.params, routerState: this.props.routerState};
  }
  render() {
    console.log("render page component**************", this.props);
    let compKey = this.props.pageId + this.props.placeholder
    return <Panel key={compKey}  description={this.props.description} />
  }
}

PageComponent.childContextTypes = {
  routeParams: PropTypes.object,
  routerState: PropTypes.object
};

//return <PageComponent pageId={pageId} placeholder={key} routerState={routerState} description={pagecomp} />


export {
  Initialize ,
  PreprocessPageComponents,
  RenderPageComponent,
  SiteTheme as Theme
}

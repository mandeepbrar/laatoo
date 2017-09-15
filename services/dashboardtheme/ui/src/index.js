import React from 'react'
import Header from './Header'
import Welcome from './Welcome'
import './styles/app.scss'

var module = this;

function Initialize(appName, ins, mod, settings) {
  module.properties = Application.Properties["dashboardtheme"]
  module.settings = settings;

  if(!Application.Registry.Pages || !Application.Registry.Page['home']) {
    let homePage={id:"home", pattern:"", component: <Welcome modProps={module.properties}/>}
    Application.Register('Pages', "home", homePage);
    Application.Register('Actions','Page_home', {url:''})
  }
}

function processPage(menus, page, uikit){
  let pageRoute = {pattern: page.pattern, reducer: page.reducer}
  pageRoute.components = {
    "menu":(routerState) => {
      return (
        <uikit.Navbar items={menus} vertical/>
      )
    },
    "main":(routerState) => {
      console.log(page.component);
      return (
        <div className="main">{page.component}</div>
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

const DashboardTheme = (props) => {
  console.log("uikit initializer", props.uikit, props.uikit.UIWrapper)
  return (
    <props.uikit.UIWrapper>
      <div className={module.properties.className?module.properties.className:'dashboard'}>
        <Header headerProps={module.properties.header} />
        <div className="body">
          <div className="menu">
            <props.router.View name="menu"  />
          </div>
          <div className="page">
            <props.router.View name="main"/>
          </div>
        </div>
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

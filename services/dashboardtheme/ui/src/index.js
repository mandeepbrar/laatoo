import React from 'react'
import {View} from 'redux-director'
import Header from './Header'
import Home from './Home'
import Menu from './Menu'
import './styles/app.scss'

var module = this;

function Initialize(appName, settings) {
  console.log(document.InitConfig);
  let dashProps = document.InitConfig.Properties["dashboardtheme"]
  module.properties = settings.propertiesOverrider ? Object.assign({}, dashProps, document.InitConfig.Properties[settings.propertiesOverrider]) : dashProps;
  console.log("Initializing dashboard theme with settings ", module.properties)
  let homePage={id:"home", pattern:"", component: <Home/>}
  Application.Register('Pages', "home", homePage)
}
function processPage(page){
  let pageRoute = {pattern: page.pattern, reducer: page.reducer}
  pageRoute.components = {
    "menu":(routerState) => {
      return (
        <Menu/>
      )
    },
    "main":(routerState) => {
      console.log(page.component);
      return (
        <div className="main">{page.component}</div>
      )
    }
  }
  console.log("Registered route", pageRoute)
  Application.Register('Routes', page.id, pageRoute)
}
function Start(appName){
  let pages=Application.Registry.Pages
  if(pages!=null) {
    for(var pageId in pages) {
      console.log("processing page", pageId)
      processPage(pages[pageId])
    }
  }
  console.log("started app", Application.Registry.Routes)
}
/*pattern:"/newmenu/:type/:id",
    components:  {
      "main": (routerState) => {
        return (
          <PermissionCheck>
            <Designer type="Mehfil" itemId={routerState.params.id}/>
          </PermissionCheck>        )
      },
      "designersettings": (routerState) => {
        return (
          <MenuEdit type={routerState.params.type} mehfilId={routerState.params.id}/>
        )
      }
    },
    reducer: combineReducers({Menu:  EntityReducer("Menu"), UserPages: ViewReducer("UserPages")})*/
class DashboardTheme extends React.Component {
  render() {
    console.log("props of dashboard theme", this.props, module.properties)
    return (
      <div className={module.properties.className?module.properties.className:'dashboard'}>
        <Header headerProps={module.properties.header} />
        <div className="body">
          <div className="menu">
            <View name="menu"  />
          </div>
          <div className="page">
            <View name="main"/>
          </div>
        </div>
        <div className="footer">
        </div>
      </div>
    )
  }
}

export {
  Initialize ,
  Start,
  DashboardTheme as Theme
}

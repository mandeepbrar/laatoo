import React from 'react'
import './styles/app.scss'
import Panel from './Panel'
import { combineReducers } from 'redux';
const PropTypes = require('prop-types');
//import {ViewReducer, View} from 'laatooviews';
var module = this;

function Initialize(appName, ins, mod, settings, def, req) {
  module.properties = Application.Properties[ins]
  module.settings = settings;
  module.req = req;
  if(!Window.redirectPage) {
    Window.redirectPage= (pageId, params) => {
      let page = _reg('Pages', pageId)
      console.log("redirect page", page)
      if(page) {
        let formattedUrl = formatUrl(page.url, params);
        Window.redirect(formattedUrl);
      }
    }
  }
  Panel.setModule(module)
}

function ProcessPages(theme, uikit) {
  let pages = Application.AllRegItems("Pages")
  if(pages) {
    for(var pageId in pages) {
      try {
        let page = pages[pageId]
        let reducers = GetPageReducers( page)
        let components = page.components
        if(page.component) {
          components = {"main":page.component}
        }
        let pageComps={}
        Object.keys(components).forEach(function(key){
          pageComps[key] = function(comp, page) {
            return (routerState) => {
              return <PageComponent pageId={page} placeholder={key} routerState={routerState} description={comp} />
            }
          }(components[key], pageId)
        });
        let route = {pattern: page.route, components: pageComps, reducer: combineReducers(reducers)}
        let newRoute = route
        if(theme && theme.ProcessRoute) {
          newRoute = theme.ProcessRoute(route, uikit)
        }
        Application.Register('Routes', pageId, newRoute)
        Application.Register('Actions','Page_'+pageId, {url: newRoute.pattern})
      }catch(ex) {
        console.log(ex)
      }
      //processPage(settings.pages[pageId], req)
    }
  }
}


function GetPageReducers(page) {
  let reducers = {}
  for(var datasourceId in page.datasources) {
    try {
      let datasource = _reg("Datasources", datasourceId)
      let obj= {}
      switch(datasource.type) {
        default:
          let mod = datasource.module
          if(mod) {
            let moduleObj = module.req(mod);
            if(moduleObj) {
              obj=moduleObj[datasource.processor]
            }
          }
      }
      if(obj) {
        reducers[datasourceId] = obj;
      }
    }catch(ex){}
  }
  return reducers
}

class PageComponent extends React.Component {
  getChildContext() {
    return {routeParams: this.props.routerState.params};
  }
  render() {
    let compKey = this.props.pageId + this.props.placeholder
    return <Panel key={compKey}  description={this.props.description} />
  }
}

PageComponent.childContextTypes = {
  routeParams: PropTypes.object
};

export {
  Initialize,
  Panel,
  ProcessPages
}

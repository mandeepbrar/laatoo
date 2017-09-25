import React from 'react'
import './styles/app.scss'
import Block from './Block'
import { combineReducers } from 'redux';
//import {ViewReducer, View} from 'laatooviews';
var module = this;

function Initialize(appName, ins, mod, settings, def, req) {
  module.properties = Application.Properties[ins]
  module.settings = settings;
  module.req = req;
  let viewsMod = module.req("laatooviews")
  if(viewsMod) {
    module.view = viewsMod["View"]
  }
  Block.setModule(module)
}

function ProcessPages(theme, uikit) {
  let pages = Application.Registry.Pages
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
              return <Block key={page+key}  blockDescription={comp} />
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
      let datasource = Application.Registry.Datasources[datasourceId]
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

export {
  Initialize,
  Block,
  ProcessPages
}

import React from 'react'
import './styles/app.scss'
import Panel from './Panel'
import { combineReducers } from 'redux';
const PropTypes = require('prop-types');
//import {ViewReducer, View} from 'laatooview';

var module;
function Initialize(appName, ins, mod, settings, def, req) {
  module=this;
  module.properties = Application.Properties[ins]
  module.settings = settings;
  module.req = req;
  Panel.setModule(module)
}

Window.resolvePanel = (panelType, id) => {
  return <Panel type={panelType} id={id}/>
}

function ProcessPages(theme) {
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
        if(theme && theme.PreprocessPageComponents) {
          components = theme.PreprocessPageComponents(components, page, pageId, reducers)
        }
        let pageComps={}
        console.log("page components ", pageId, page, components)
        Object.keys(components).forEach(function(key){
          pageComps[key] = function(pagecomp, key, pageId, page) {
            return (routerState) => {
              let visible = true
              console.log("Page components ", routerState, pagecomp, key, pageId, page)
              if(theme && theme.IsComponentVisible) {
                visible = theme.IsComponentVisible(compToRender, key, pageId, routerState, page)
              }
              if(visible) {
                let compToRender = typeof(pagecomp) == 'function'? pagecomp(routerState): pagecomp
                if(theme && theme.RenderPageComponent) {
                  let retval = theme.RenderPageComponent(compToRender, key, pageId, routerState, page)
                  if(retval) {
                    return retval
                  }
                }
                return compToRender  
              }
              return null
            }
          }(components[key], key, pageId, page)
        });
        let route = {pattern: page.route, components: pageComps, reducer: combineReducers(reducers)}
        console.log("page ....", route)
        Application.Register('Routes', pageId, route)
        Application.Register('Actions','Page_'+pageId, {url: route.pattern})
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

export {
  Initialize,
  Panel,
  ProcessPages
}

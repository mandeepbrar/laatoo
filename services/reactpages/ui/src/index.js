import React from 'react'
import Header from './Header'
import Welcome from './Welcome'
import './styles/app.scss'
import { combineReducers } from 'redux';
import {ViewReducer, View} from 'laatooviews';
var module = this;

function Initialize(appName, ins, mod, settings, def, req) {
  module.properties = Application.Properties[ins]
  module.settings = settings;

  if(settings.pages) {
    for(var pageId in settings.pages) {
      console.log("processing page", pageId)
      let page = settings.pages[pageId]
      let reducers = processReducers(pageId, page, settings, req)
      let components =  processComponents(pageId, page, settings, req)
      Application.Register('Pages', pageId, {pattern: page.route, components: components, reducer: combineReducers(reducers)})
      //processPage(settings.pages[pageId], req)
    }
  }
}

function processComponents(pageId, page, settings, req) {
  let components = {}
  let pagecomponents = page.components
  if(page.component) {
    pagecomponents = {"main":pagecomponents}
  }
  if(pagecomponents) {
    for(var viewId in pagecomponents) {
      let compdesc = pagecomponents[viewId]
      try {
        let obj= null
        switch compdesc.type {
          case "view":
            let viewname = compdesc.name
            obj = <View name={viewname} contentClass="row" paginate={false} getView={getLeaderRow}
              defaultFilter={compdesc.filter} global={compdesc.global?compdesc.global:false}>
              <ViewItem style={{padding:0}}/>
            </View>
            break;
          /*case "entity":
            obj = EntityReducer(datasourceId)
            break;*/
          default:
            let mod = compdesc.module
            if(mod) {
              let module = req(mod);
              if(module) {
                obj=module[compdesc.component]
              }
            }
        }
        if(obj) {
          components[viewId] = obj;
        }
      }catch(ex){}
    }
  }
  return components
}

function processReducers(pageId, page, settings, req) {
  let reducers = {}
  for(var datasourceId in page.datasources) {
    try {
      let datasource = Application.Registry.Datasources[datasourceId]
      let obj= {}
      switch datasource.type {
        case "view":
          obj = ViewReducer(datasourceId)
          break;
        /*case "entity":
          obj = EntityReducer(datasourceId)
          break;*/
        default:
          let mod = datasource.module
          if(mod) {
            let module = req(mod);
            if(module) {
              obj=module[datasource.processor]
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
  Initialize
}

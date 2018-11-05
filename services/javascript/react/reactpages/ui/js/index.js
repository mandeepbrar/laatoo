import React from 'react'
import './styles/app.scss'
import Panel from './Panel'
import { combineReducers } from 'redux';
const PropTypes = require('prop-types');
//import {ViewReducer, View} from 'laatooviews';

var module;
function Initialize(appName, ins, mod, settings, def, req) {
  module=this;
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
  console.log("jhkjhjkhjkhkjhh = ", pages)
  if(pages) {
    for(var pageId in pages) {
      try {
        let page = pages[pageId]
        let reducers = GetPageReducers( page)
        let components = page.components
        let pageComp = function(comp, page) {
          return (routerState) => {
            console.log("page change-------------------------------------", routerState)
            return <PageComponent pageId={page} placeholder={key} routerState={routerState} description={comp} />
          }
        }
        if(page.component) {
          console.log("changed ==============")
          components = {"main": pageComp(page.component, pageId)}
        }
        let pageComps={}
        Object.keys(components).forEach(function(key){
          pageComps[key] = pageComp(components[key], pageId)
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
    return {routeParams: this.props.routerState.params, routerState: this.props.routerState};
  }
  render() {
    console.log("page component render---", this.props)
    let compKey = this.props.pageId + this.props.placeholder
    return <Panel key={compKey}  description={this.props.description} />
  }
}

PageComponent.childContextTypes = {
  routeParams: PropTypes.object,
  routerState: PropTypes.object
};

export {
  Initialize,
  Panel,
  ProcessPages
}

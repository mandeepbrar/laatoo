import React from 'react'
const PropTypes = require('prop-types');
import {Panel} from 'reactpages';

var module;

function setModule(mod) {
    module = mod;
}

function RenderPageComponent(comp, key, pageId, routerState, page, uikit) {
    console.log("RenderPageComponent", routerState, key, comp, page)
    if(key=="main") {
      return <SitePage pageComp={comp} pageKey={key} pageId={pageId} routerState={routerState} page={page} uikit={uikit}/>
    }
    if(key=="menu") {
  //    return null
      return <SiteMenu menu={comp} pageKey={key} pageId={pageId} routerState={routerState} page={page} uikit={uikit}/>
    }
}

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
  


export {
    RenderPageComponent, setModule    
}
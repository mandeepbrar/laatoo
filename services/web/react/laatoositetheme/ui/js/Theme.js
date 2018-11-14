import React from 'react';
import {Panel} from 'reactpages';
import {connect} from 'react-redux';
import {LoginValidator, LoginForm} from 'authui';
import Header from './Header'

var module;

function setModule(mod) {
    module = mod;
}

class ThemeUI extends React.Component {
  constructor(props) {
    super(props);
  }
  render() {
      console.log("rendering theme ui", this.props);
    let props = this.props;
    let vertical = module.settings.vertical?true:false
    return (
      <props.uikit.Block className={module.settings.className?module.settings.className + ' dashboard':'dashboard'}>
        <Header uikit={props.uikit} module={module} />
        <props.uikit.Block className="body">
        {
        module.settings.showMenu?
          <props.uikit.Block className={vertical?"vertmenu":"horizmenu"}>
            <props.router.View name="menu" loggedIn={props.loggedIn} />
          </props.uikit.Block>
        :null
        }
          <props.uikit.Block className={vertical?"vertbody":"horizbody"}>
            <props.router.View name="main" loggedIn={props.loggedIn}/>
          </props.uikit.Block>
        </props.uikit.Block>
        <props.uikit.Block className="footer">
        {props.children}
        </props.uikit.Block>
      </props.uikit.Block>)
  }
}



const SiteTheme = (props, st, c) => {
    let vertical = module.settings.vertical?true:false
    return (
      <props.uikit.UIWrapper>
        <LoginValidator  validateService="validate">
          <ThemeUI uikit={props.uikit} router={props.router}/>
        </LoginValidator>
      </props.uikit.UIWrapper>
    )
}
  

function PreprocessPageComponents(components, page, pageId, reducers, uikit) {
    if(module.settings.showMenu) {
        console.log("pre processing.......", components);
        components.menu = (routerState) => {
          console.log("menu items ", routerState, module.menu)
          return (
            <uikit.Navbar items={module.menu} vertical={module.settings.vertical}/>
          )
        }    
    }
    return components;
}

  

export {
    setModule,
    SiteTheme,
    PreprocessPageComponents
} 

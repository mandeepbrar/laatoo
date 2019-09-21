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
      <_uikit.Block className={module.settings.className?module.settings.className + ' dashboard':'dashboard'}>
        <Header module={module} />
        <_uikit.Block className={"body " + (vertical?"fdrow":"fdcol")}>
          <_uikit.Block className={vertical?"vertmenu":"horizmenu"}>
            <Application.router.View name="menu" loggedIn={props.loggedIn} />
          </_uikit.Block>
          <_uikit.Block className={vertical?"vertbody":"horizbody"}>
            <Application.router.View name="main" loggedIn={props.loggedIn}/>
          </_uikit.Block>
        </_uikit.Block>
        <_uikit.Block className="footer">
        {props.children}
        </_uikit.Block>
      </_uikit.Block>)
  }
}



const SiteTheme = (props, st, c) => {
    let vertical = module.settings.vertical?true:false
    return (
      <_uikit.UIWrapper>
        <LoginValidator  validateService="validate">
          <ThemeUI />
        </LoginValidator>
      </_uikit.UIWrapper>
    )
}
  

function PreprocessPageComponents(components, page, pageId, reducers) {
    if(module.settings.showMenu && !page.hideMenu) {
      let comp = module.settings.menuLagger      
      console.log("pre processing.......", components, page);
      components.menu = (routerState) => {
        console.log("menu items ", routerState, module.menu)
        return (
          <_uikit.Block className="menu">
              <_uikit.Navbar items={module.menu} vertical={module.settings.vertical}/>
              {module.settings.menuEnd?
              <Panel id={module.settings.menuEnd} type="block" className="menuEnd"/>
              :null
              }              
          </_uikit.Block>
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

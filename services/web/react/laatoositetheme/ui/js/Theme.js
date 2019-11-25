import React from 'react';
import {Panel} from 'reactpages';
import {connect} from 'react-redux';
import {Menu} from 'reactwebcommon';
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
          <_uikit.Block className={vertical?" fd fdcenter vertbody":" fd fdcenter horizbody"}>
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
  console.log("props of theme", props)
    let vertical = module.settings.vertical?true:false
    return (
      <_uikit.UIWrapper>
        <_uikit.Block className="interactions">
        {props.interactionsComponent}
        </_uikit.Block>
        <LoginValidator  validateService="validate">
          <ThemeUI />
        </LoginValidator>
      </_uikit.UIWrapper>
    )
}
  

function PreprocessPageComponents(components, page, pageId, reducers) {
    if(module.menu && module.settings.showMenu && !page.hideMenu) {
      let comp = module.settings.menuLagger      
      console.log("pre processing.......", components, page);
      components.menu = (routerState) => {
        return (
          <_uikit.Block className="menu">
            {
              module.menu?
              <Menu id={module.menu} vertical={module.settings.vertical}/>
              :null
            }            
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

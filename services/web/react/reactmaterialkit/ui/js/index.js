import React from 'react';
import {render } from 'react-dom'
import Dialogs from './components/Dialogs'
import Navbar from './components/Navbar'
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles';
import {Button} from '@material-ui/core';
import {TextField} from './components/TextField';
import {Select} from './components/Select';
import './styles/app.scss';
import {FieldWidget} from './forms/FieldWidget';
import {Tabset, Tab} from './components/Tabs';
import {Block} from './components/Block';
import {Icons} from './components/Icons';
import Icon from '@material-ui/core/Icon';
import {ScrollListener, Action, Html, ActionBar, LoadableComponent, Image} from 'reactwebcommon';
//import injectTapEventPlugin from "react-tap-event-plugin";


function Initialize(appName, ins, mod, settings, def, req) {
  //injectTapEventPlugin();
}


const ActionButton=(props)=> (
  <Button variant="contained" onClick={props.onClick} {...props.btnProps} className={props.className} style={props.style}>{props.children}</Button>
)


class UIWrapper extends React.Component {
  constructor(props) {
    super(props)
    this.theme = createMuiTheme()
  }
  /*componentDidMount() {
    var injectTapEventPlugin = require("react-tap-event-plugin");
    console.log("injectd.........", injectTapEventPlugin)
//    injectTapEventPlugin();
    injectTapEventPlugin();
  }*/
  render() {
    return (
      <MuiThemeProvider><Dialogs/>{this.props.children}</MuiThemeProvider>
    )
  }
}

class Form extends React.Component {
  constructor(props) {
    super(props)
    this.reset = this.reset.bind(this)
  }
  reset() {
    this.form.reset();
  }
  render() {
    return <form ref={(form) => this.form = form } {...this.props}></form>
  }
}

export {
  Initialize,
  render,
  FieldWidget as Field,
  Icon,
  Icons,
  ActionButton,
  Tabset,
  Tab,
  Dialogs,
  Select,
  Block,
  TextField,
  Form,
  ScrollListener as Scroll,
  Action, 
  Html, 
  ActionBar, 
  LoadableComponent, 
  Image,
  UIWrapper,
  Navbar
}

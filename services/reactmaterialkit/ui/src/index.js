import React from 'react';
import {render } from 'react-dom'
import Dialogs from './components/Dialogs'
import Navbar from './components/Navbar'
import { MuiThemeProvider, createMuiTheme } from 'material-ui/styles';
import {Button, FontIcon} from 'material-ui';
import {TextField} from './components/TextField';
import {ScrollListener} from 'reactwebcommon';
import './styles/app.scss';
import Forms from './forms';
import { FormControl } from 'material-ui/Form';
import Select from 'material-ui/Select';
import { MenuItem } from 'material-ui/Menu';
import {Tabset, Tab} from './components/Tabs';
import {Block} from './components/Block';
import {Icons} from './components/Icons';
//import injectTapEventPlugin from "react-tap-event-plugin";


function Initialize(appName, ins, mod, settings, def, req) {
  //injectTapEventPlugin();
}


const ActionButton=(props)=> (
  <Button raised onClick={props.onClick} {...props.btnProps} className={props.className} style={props.style}>{props.children}</Button>
)


class UIWrapper extends React.Component {
  constructor(props) {
    super(props)
    this.theme = createMuiTheme()
  }
  componentDidMount() {
    var injectTapEventPlugin = require("react-tap-event-plugin");
    console.log("injectd.........", injectTapEventPlugin)
//    injectTapEventPlugin();
    injectTapEventPlugin();
  }
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
  Icon,
  Icons,
  ActionButton,
  Tabset,
  Tab,
  Forms,
  Dialogs,
  Block,
  TextField,
  Form,
  ScrollListener as Scroll,
  UIWrapper,
  Navbar
}

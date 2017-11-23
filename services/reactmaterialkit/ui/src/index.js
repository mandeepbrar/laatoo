import {render } from 'react-dom'
import Dialogs from './components/Dialogs'
import Navbar from './components/Navbar'
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import {Button} from 'material-ui';
import {TextField} from './components/TextField';
import {ScrollListener} from 'reactwebcommon';
import React from 'react';
import './styles/app.scss';
import Forms from './forms';
import { FormControl } from 'material-ui/Form';
import Select from 'material-ui/Select';
import { MenuItem } from 'material-ui/Menu';
//import injectTapEventPlugin from "react-tap-event-plugin";

console.log("react material kit... forms", Forms)

function Initialize(appName, ins, mod, settings, def, req) {
  //injectTapEventPlugin();
}

const Block=(props) => (
  <div style={props.style} className={props.className}>
   {props.children}</div>
)


const ActionButton=(props)=> (
  <Button raised onClick={props.onClick} {...props.btnProps} className={props.className} style={props.style}>{props.children}</Button>
)

class UIWrapper extends React.Component {
  componentDidMount() {
    var injectTapEventPlugin = require("react-tap-event-plugin");
    console.log("injectd.........", injectTapEventPlugin)
//    injectTapEventPlugin();
    injectTapEventPlugin();
  }
  render() {
    return (
      <MuiThemeProvider>{this.props.children}</MuiThemeProvider>
    )
  }
}

const Form=(props)=>(
  <form {...props}></form>
)

export {
  Initialize,
  render,
  ActionButton,
  Forms,
  Dialogs,
  Block,
  TextField,
  Form,
  ScrollListener as Scroll,
  UIWrapper,
  Navbar
}

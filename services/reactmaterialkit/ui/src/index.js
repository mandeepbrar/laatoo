import {render } from 'react-dom'
import Dialogs from './components/Dialogs'
import Navbar from './components/Navbar'
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import {RaisedButton} from 'material-ui';
import {TextField} from './components/TextField';
import {ScrollListener} from 'reactwebcommon';
import React from 'react';
import './styles/app.scss';

const Block=(props) => (
  <div style={props.style} className={props.className}>{props.children}</div>
)


const ActionButton=(props)=> (
  <RaisedButton onClick={props.onClick} {...props.btnProps} className={props.className} style={props.style}>{props.children}</RaisedButton>
)

const Form=(props)=> (
  <form>{props.children}</form>
)

const UIWrapper=(props)=>(
  <MuiThemeProvider>{props.children}</MuiThemeProvider>
)

export {
  render,
  ActionButton,
  Dialogs,
  Block,
  TextField,
  Form,
  ScrollListener as Scroll,
  UIWrapper,
  Navbar
}

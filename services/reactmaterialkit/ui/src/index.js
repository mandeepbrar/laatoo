import {render } from 'react-dom'
import Dialogs from './components/Dialogs'
import Navbar from './components/Navbar'
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import React from 'react';
import './styles/app.scss';

const UIWrapper=(props)=>(
  <MuiThemeProvider>{props.children}</MuiThemeProvider>
)

export {
  render,
  Dialogs,
  UIWrapper,
  Navbar
}

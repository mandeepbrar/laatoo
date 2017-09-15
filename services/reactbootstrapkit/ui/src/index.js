import {render } from 'react-dom'
import Dialogs from './components/Dialogs'
import Navbar from './components/Navbar'
import React from 'react';

const UIWrapper=(props)=>(
  <div>{props.children}</div>
)

export {
  render,
  Dialogs,
  UIWrapper,
  Navbar
}

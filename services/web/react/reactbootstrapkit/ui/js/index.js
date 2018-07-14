import {render } from 'react-dom'
import Dialogs from './components/Dialogs'
import Navbar from './components/Navbar'
import React from 'react';
import {TextField} from './components/TextField';
import {ScrollListener} from 'reactwebcommon';

const UIWrapper=(props)=>(
  <div>{props.children}</div>
)

const ActionButton=(props)=> (
  <RaisedButton onClick={props.onClick} {...props.btnProps} className={props.className} style={props.style}>{props.children}</RaisedButton>
)

const Block=(props) => (
  <div style={props.style} className={props.className}>{props.children}</div>
)

const Form=(props)=> (
  <form>{props.children}</form>
)

/*const Icon=(props)=> (
    <Icon className={props.className} style={props.style}/>
)*/


export {
  render,
  ActionButton,
  Dialogs,
  TextField,
  Form,
  ScrollListener as Scroll,
  Block,
  UIWrapper,
  Navbar
}

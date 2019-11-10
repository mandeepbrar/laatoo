import {render } from 'react-dom'
import Dialogs from './components/Dialogs'
import Navbar from './components/Navbar'
import React from 'react';
import {TextField} from './components/TextField';
import {FieldWidget} from './forms/FieldWidget';
import {Icons} from './components/Icons';
import {Block, ScrollListener, Action, Html, ActionBar, LoadableComponent, Image} from 'reactwebcommon';
import {Tabset, Tab} from './components/Tabs';
import Button from 'react-bootstrap/Button'
import './styles/app.scss';
import {Form as BSForm} from 'react-bootstrap';
import {Select} from './components/Select';
import 'bootstrap/dist/css/bootstrap.min.css';

function Initialize(appName, ins, mod, settings, def, req) {
  //injectTapEventPlugin();
}

const UIWrapper=(props)=>(
  props.children
)

const ActionButton=(props)=> (
  //props.flat?
  <Button onClick={props.onClick} {...props.btnProps} className={props.className} style={props.style}>{props.children}</Button>
)

const Form=(props)=> (
  <BSForm>{props.children}</BSForm>
)

const Icon=(props)=> (
    <i {...props}></i>
)


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


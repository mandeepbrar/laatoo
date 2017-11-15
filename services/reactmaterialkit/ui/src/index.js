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

class SelectFieldExampleSimple extends React.Component {
  state = {
    value: 1,
  };

  handleChange = (event, index, value) => {console.log("on change called", event); this.setState({value});} ;

  render() {
    return (
        <Select
          floatingLabelText="Frequency"
          value={this.state.value}
          onChange={this.handleChange} >
          <MenuItem key={"ssss1"} value={1} primaryText="Never" />
          <MenuItem key={"ssss2"} value={2} primaryText="Every Night" />
          <MenuItem key={"ssss3"} value={3} primaryText="Weeknights" />
          <MenuItem key={"ssss4"} value={4} primaryText="Weekends" />
          <MenuItem key={"ssss5"} value={5} primaryText="Weekly" />
        </Select>
      )
  }
}

const Block=(props) => (
  <div style={props.style} className={props.className}>
    <SelectFieldExampleSimple/>
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

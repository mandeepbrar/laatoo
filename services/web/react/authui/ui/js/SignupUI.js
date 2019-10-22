'use strict';

import React from 'react';

const PropTypes = require('prop-types');

class SignupUI extends React.Component {
    constructor(props) {
      super(props);
      console.log("costructor of login web")
      this.state = {
        "name": "",
        "email":"",
        "password":"",
        "confirmpassword":""
      };
      this.handleSignup = this.handleSignup.bind(this);
      this.handleChange = this.handleChange.bind(this);
      let loginSite = "";
      let realm = ""
      if(props.realm) {
        realm = "?Realm=" + props.realm
      }
    }
    handleChange(e) {
        var nextState = {};
        nextState[e.target.name] = e.target.value;
        this.setState(nextState);
    }
    handleSignup() {
        this.props.handleSignup(this.state.name, this.state.email, this.state.password, this.state.confirmpassword);
    }
    render() {
      console.log("login ui", this.props);
      return this.props.renderSignup(this.state, this.handleChange, this.handleSignup, this.props)
    }
  }
  
 
  // Uncomment properties you need
  SignupUI.propTypes = {
    handleSignup: PropTypes.func.isRequired
  };
  

  export {SignupUI} ;


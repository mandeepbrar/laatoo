'use strict';

import React from 'react';

const PropTypes = require('prop-types');

class LoginUI extends React.Component {
  constructor(props) {
    super(props);
    console.log("costructor of login web")
    this.state = {
      "email":"",
      "password":""
    };
    this.handleLogin = this.handleLogin.bind(this);
    this.handleChange = this.handleChange.bind(this);
/*
    let getLocation = function(href) {
        var l = document.createElement("a");
        l.href = href;
        return l;
    };*/
    let loginSite = "";
    let realm = ""
    if(props.realm) {
      realm = "?Realm=" + props.realm
    }
    /*this.oauthLogin = function(location) {
      let instance = window.open(location+realm, '_blank','height=500,width=400,toolbar=no,resizable=yes,menubar=no,location=0')
      var location = getLocation(location);
      loginSite = location.protocol+'//'+location.hostname+(location.port ? ':'+location.port: '');
    }

    window.addEventListener("message", function(ev) {
      if(ev.origin === loginSite && ev.data.message=="LoginSuccess") {
        props.handleOauthLogin(ev.data)
      }
    });*/
  }
  handleChange(e) {
      var nextState = {};
      nextState[e.target.name] = e.target.value;
      this.setState(nextState);
  }
  handleLogin() {
      this.props.handleLogin(this.state.email, this.state.password);
  }
  render() {
    console.log("login ui", this.props);
    return this.props.renderLogin(this.state, this.handleChange, this.handleLogin, this.oauthLogin, this.props)
  }
}

//LoginUI.displayName = 'LoginComponent';

// Uncomment properties you need
LoginUI.propTypes = {
  handleOauthLogin: PropTypes.func.isRequired,
  handleLogin: PropTypes.func.isRequired
};

LoginUI.contextTypes = {
  uikit: PropTypes.object
};

export {LoginUI} ;

'use strict';

import React from 'react';

class LoginWeb extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      "email":"",
      "password":""
    };
    this.handleLogin = this.handleLogin.bind(this);
    this.handleChange = this.handleChange.bind(this);

    let getLocation = function(href) {
        var l = document.createElement("a");
        l.href = href;
        return l;
    };
    let loginSite = "";
    let realm = ""
    if(props.realm) {
      realm = "?Realm=" + props.realm
    }
    this.oauthLogin = function(location) {
      let instance = window.open(location+realm, '_blank','height=500,width=400,toolbar=no,resizable=yes,menubar=no,location=0')
      var location = getLocation(location);
      loginSite = location.protocol+'//'+location.hostname+(location.port ? ':'+location.port: '');
    }

    window.addEventListener("message", function(ev) {
      if(ev.origin === loginSite && ev.data.message=="LoginSuccess") {
        props.handleOauthLogin(ev.data)
      }
    });
  }
  handleChange(e) {
      console.log("handle change of login ", e)
      var nextState = {};
      nextState[e.target.name] = e.target.value;
      this.setState(nextState);
  }
  handleLogin() {
      this.props.handleLogin(this.state.email, this.state.password);
  }
  render() {
    return this.props.renderLogin(this.state, this.handleChange, this.handleLogin, this.oauthLogin)
  }
}

LoginWeb.displayName = 'LoginComponent';

// Uncomment properties you need
LoginWeb.propTypes = {
  handleOauthLogin: React.PropTypes.func.isRequired,
  handleLogin: React.PropTypes.func.isRequired
};

export default LoginWeb;

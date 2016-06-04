'use strict';

import React from 'react';
require('./Login.scss');

class LoginWeb extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      "email":"",
      "password":""
    };
    this.handleLogin = this.handleLogin.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.facebook = this.facebook.bind(this);
  }
  handleChange(e) {
      var nextState = {};
      nextState[e.target.name] = e.target.value;
      this.setState(nextState);
  }
  handleLogin() {
      this.props.handleLogin(this.state.email, this.state.password);
  }
  signup() {
    if(this.props.signup && this.props.signup === 'true') {
        return(
            <div>, or <a href="#">Sign Up</a></div>
        )
    }
  }
  facebook() {
      if(this.props.facebook && this.props.facebook === 'true') {
          return(
            <div className="col-xs-6 col-sm-6 col-md-6">
              <a href={"javascript:window.open('"+ this.props.facebookAuthUrl +"', '_blank','height=500,width=400,toolbar=no,resizable=yes,menubar=no,location=0')"} className="btn btn-lg btn-info btn-block">Facebook</a>
            </div>
          )
      }
  }
  google() {
      if(this.props.google && this.props.google === 'true') {
          return(
            <div className="col-xs-6 col-sm-6 col-md-6">
              <a href={"javascript:window.open('"+ this.props.googleAuthUrl +"', '_blank','height=500,width=400,toolbar=no,resizable=yes,menubar=no,location=0')"} className="btn btn-lg btn-info btn-block">Google</a>
            </div>
          )
      }
  }
  social() {
      if((this.props.google && this.props.google === 'true') ||
         (this.props.facebook && this.props.facebook === 'true')
        ){
          return(
              <div className="login-or">
                <hr className="hr-or"/>
                <span className="span-or">or</span>
              </div>
          )
      }
  }
  render() {
    return (
        <div className="container loginbox">
          <div className="row">

            <div className="main">

              <h3>Please Log In {this.signup()}</h3>
              <div className="row">
                {this.facebook()}
                {this.google()}
              </div>
                {this.social()}

              <form role="form">
                <div className="form-group">
                  <label htmlFor="email">Username or email</label>
                  <input type="text" className="form-control" name="email" value={this.state.email} placeholder="Email" onChange={this.handleChange} />
                </div>
                <div className="form-group">
                  <a className="pull-right" href="#">Forgot password?</a>
                  <label htmlFor="inputPassword">Password</label>
                  <input type="password" className="form-control" name="password" value={this.state.password} placeholder="Password" onChange={this.handleChange} />
                </div>
                <div className="checkbox">
                    <label>
                        <input type="checkbox"/>
                        Remember me
                    </label>
                </div>
                <button type="button"  className="btn btn btn-primary  pull-right" onClick={this.handleLogin}>Login</button>
              </form>
            </div>
          </div>
        </div>
    );
  }
}

LoginWeb.displayName = 'LoginComponent';

// Uncomment properties you need
LoginWeb.propTypes = {
  handleLogin: React.PropTypes.func.isRequired,
  facebook:  React.PropTypes.string,
  facebookAuthUrl: React.PropTypes.string,
  google:  React.PropTypes.string,
  googleAuthUrl: React.PropTypes.string,
  signup:  React.PropTypes.string
};
// LoginComponent.defaultProps = {};

export default LoginWeb;

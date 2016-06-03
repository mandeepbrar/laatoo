'use strict';

import React from 'react';
import LoginWeb from './LoginWeb';
import md5 from 'md5';
import { connect } from 'react-redux';
import {ActionNames} from '../../actions/ActionNames';
import {createAction} from '../../utils';

const mapStateToProps = (state, ownProps) => {
  return {
    facebook: ownProps.facebook,
    google: ownProps.google,
    signup: ownProps.signup
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    handleLogin: (email, password) => {
      let loginPayload = {"Id": email, "Password": md5(password)};
      let loginMeta = {serviceName: ownProps.loginService};
      dispatch(createAction(ActionNames.LOGIN, loginPayload, loginMeta));
    }
  }
}

const LoginComponent = connect(
  mapStateToProps,
  mapDispatchToProps
)(LoginWeb);

// Uncomment properties you need
LoginComponent.propTypes = {
  loginService: React.PropTypes.string.isRequired,
  successpage: React.PropTypes.string,
  facebook: React.PropTypes.string,
  google: React.PropTypes.string,
  signup: React.PropTypes.string
};
// LoginComponent.defaultProps = {};

export default LoginComponent;

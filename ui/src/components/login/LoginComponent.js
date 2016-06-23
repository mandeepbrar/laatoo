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
    facebookAuthUrl: ownProps.facebookAuthUrl,
    google: ownProps.google,
    googleAuthUrl: ownProps.googleAuthUrl,
    realm : ownProps.realm,
    signup: ownProps.signup
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  let realm = "";
  if(ownProps.realm) {
	   realm = ownProps.realm
  }
  return {
    handleLogin: (email, password) => {
      let loginPayload = {"Username": email, "Password": md5(password), "Realm" : realm };
      let loginMeta = {serviceName: ownProps.loginService};
      dispatch(createAction(ActionNames.LOGIN, loginPayload, loginMeta));
    },
    handleOauthLogin: (data) => {
      dispatch(createAction(ActionNames.LOGIN_SUCCESS, {userId: data.id, token: data.token, permissions: data.permissions}));
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
  realm: React.PropTypes.string,
  facebook: React.PropTypes.string,
  facebookAuthUrl: React.PropTypes.string,
  google: React.PropTypes.string,
  googleAuthUrl: React.PropTypes.string,
  signup: React.PropTypes.string
};
// LoginComponent.defaultProps = {};

export default LoginComponent;

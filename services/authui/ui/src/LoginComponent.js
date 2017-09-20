'use strict';

import React from 'react';
import {LoginUI} from './LoginUI';
import md5 from 'md5';
import { connect } from 'react-redux';
import {ActionNames} from './actions';
import {createAction} from 'uicommon';

const mapStateToProps = (state, ownProps) => {
  return {
    realm : Application.Security.realm,
    renderLogin: ownProps.renderLogin,
    signup: ownProps.signup
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  console.log("map dispatch of login compoent")
  let realm = "";
  if(Application.Security.realm) {
	   realm = Application.Security.realm
  }
  return {
    handleLogin: (email, password) => {
      let loginPayload = {"Username": email, "Password": md5(password), "Realm" : realm };
      let loginMeta = {serviceName: Application.Security.loginService};
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
)(LoginUI);

// Uncomment properties you need
LoginComponent.propTypes = {
  loginService: React.PropTypes.string.isRequired,
  successpage: React.PropTypes.string,
  realm: React.PropTypes.string,
  signup: React.PropTypes.string
};
// LoginComponent.defaultProps = {};

export {LoginComponent} ;

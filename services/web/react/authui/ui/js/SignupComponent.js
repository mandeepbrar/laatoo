'use strict';

import React from 'react';
import {SignupUI} from './SignupUI';
import md5 from 'md5';
import { connect } from 'react-redux';
import {ActionNames} from './actions';
import {createAction} from 'uicommon';
import PropTypes from 'prop-types';

const mapStateToProps = (state, ownProps) => {
  return {
    realm : Application.Security.realm,
    renderSignup: ownProps.renderSignup,
    signup: ownProps.signup
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  console.log("map dispatch of signup component")
  let realm = "";
  if(Application.Security.realm) {
	   realm = Application.Security.realm
  }
  return {
    handleSignup: (name, email, password, confirmpassword) => {
      console.log("load", email, password, confirmpassword)
      if (confirmpassword == password) {
        let signupPayload = {"Name": name, "Username": email, "Password": md5(password), "Realm" : realm };
        console.log(signupPayload)
        
        let signupMeta = {serviceName: Application.Security.signupService};
        dispatch(createAction(ActionNames.SIGN_UP, signupPayload, signupMeta));
      } else {
        Window.showMessage(ownProps.module.properties.errors.passwordmismatch)
      }
    }
  }
}

const SignupComponent = connect(
  mapStateToProps,
  mapDispatchToProps
)(SignupUI);

// Uncomment properties you need
SignupComponent.propTypes = {
  loginService: PropTypes.string.isRequired,
  successpage: PropTypes.string,
  realm: PropTypes.string,
  signup: PropTypes.string
};
// LoginComponent.defaultProps = {};

export {SignupComponent} ;

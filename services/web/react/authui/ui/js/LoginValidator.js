import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import {ActionNames} from './actions';
import {createAction, DataSource, RequestBuilder } from 'uicommon';
const PropTypes = require('prop-types');

var module;

function setModule(mod) {
  module = mod
}

/* Populated by react-webpack-redux:reducer */
class Login extends React.Component {
  constructor(props) {
    super(props)
    this.validatetoken = this.validatetoken.bind(this)
    this.state= {loggedIn: props.loggedIn, validation: props.validation}
    console.log("login component", props)
  }
  componentWillReceiveProps(nextprops) {
    console.log("Login validator componentWillReceiveProps", nextprops)
    if(nextprops.loggedIn != this.state.loggedIn || nextprops.validation != this.state.validation) {
      this.setState({loggedIn: nextprops.loggedIn, validation: nextprops.validation})
    }
  }

  validatetoken() {
    let logout = this.props.logout
    let login = this.props.login
    let failure = (resp) => {
      logout(true)
      this.setState({loggedIn: false, validation: false})
    }
    let success=(resp) => {
      login(resp.data.Id, resp.data.Permissions);
    }
    console.log("sending validation request")
    let req = RequestBuilder.DefaultRequest({},{} )
    DataSource.ExecuteService(this.props.validateService, req).then(success, failure);
  }
  getChildContext() {
    return {loggedIn: this.state.loggedIn};
  }
  render() {
    console.log("rendering login validator", this.state)
    if(this.state.validation) {
      this.validatetoken()
      return null
    }
    return this.props.children? React.cloneElement(this.props.children, {loggedIn: this.state.loggedIn, validation: this.state.validation}) : null
  }
}

Login.childContextTypes = {
  loggedIn: PropTypes.bool,
  user: PropTypes.object
};

function getPropsCookieMode(state, ownProps) {
  switch (state.Security.status) {
    case "NotLogged":
    return {
      validation: true,
      loggedIn: false,
      validateService: ownProps.validateService
    }
    case "LoggedIn": 
    return {
      validation: false,
      loggedIn: true
    }
    default:
    return {
      validation: false,
      loggedIn: false
    }
  }
}

function getPropsNonCookieMode(state, ownProps) {
  switch (state.Security.status) {
    case "NotLogged":
    console.log("get props non cookie, Storage.auth = ", Storage.auth)
    if(Storage.auth) {
      return {
        validation: true,
        loggedIn: false,
        validateService: ownProps.validateService
      }  
    } else {
      return {loggedIn: false}
    }
    case "LoggedIn": 
    return {
      validation: false,
      loggedIn: true
    }
    default:
    return {
      validation: false,
      loggedIn: false
    }
  }
}


const mapStateToProps = (state, ownProps) => {
  if (module.settings.cookies) {
    return getPropsCookieMode(state, ownProps)
  } else {
    return getPropsNonCookieMode(state, ownProps)
  }  
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    login: (userId, permissions) => {
      dispatch(createAction(ActionNames.LOGIN_SUCCESS, {userId: userId, token: Storage.auth, user: Storage.user, permissions: permissions}));
    },
    logout: (validationFailed) => {
      if(validationFailed) {
        dispatch(createAction(ActionNames.VALIDATIONFAILED, null, null));
      } else {
        dispatch(createAction(ActionNames.LOGOUT, null, null));
      }
    }
  };
}

const LoginValidator = connect(mapStateToProps, mapDispatchToProps)(Login);


export {LoginValidator, setModule}

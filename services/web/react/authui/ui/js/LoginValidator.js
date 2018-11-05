import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import {ActionNames} from './actions';
import {createAction, DataSource, RequestBuilder } from 'uicommon';
const PropTypes = require('prop-types');

/* Populated by react-webpack-redux:reducer */
class Login extends React.Component {
  constructor(props) {
    super(props)
    this.validatetoken = this.validatetoken.bind(this)
    this.state= {loggedIn: props.loggedIn, validation: props.validation}
    if(props.validation) {
      this.validatetoken()
    }
  }
  componentWillReceiveProps(nextprops) {
    if(nextprops.loggedIn != this.state.loggedIn || nextprops.validation != this.state.validation) {
      this.setState({loggedIn: nextprops.loggedIn, validation: nextprops.validation})
    }
  }

  validatetoken() {
    let logout = this.props.logout
    let login = this.props.login
    let failure = (resp) => {
      logout()
      this.setState({loggedIn: false, validation: false})
    }
    let success=(resp) => {
      login(resp.data.Id, resp.data.Permissions);
    }
    let req = RequestBuilder.DefaultRequest({},{} )
    DataSource.ExecuteService(this.props.validateService, req).then(success, failure);
  }
  getChildContext() {
    return {loggedIn: this.state.loggedIn};
  }
  render() {
    if(this.state.validation) {
      return null
    }
    return this.props.children? React.cloneElement(this.props.children, {loggedIn: this.state.loggedIn, validation: this.state.validation}) : null
  }
}

Login.childContextTypes = {
  loggedIn: PropTypes.bool,
  user: PropTypes.object
};

const mapStateToProps = (state, ownProps) => {
  if(Storage.auth == null) {
    return {
      validation: false,
      loggedIn: false,
      validateService: ownProps.validateService
    }
  }
  else if (Storage.auth != "") {
    if(state.Security.status != "LoggedIn") {
      return {
        validation: true,
        loggedIn: false,
        validateService: ownProps.validateService
      }
    } else {
      return {
        validation: false,
        loggedIn: true,
        validateService: ownProps.validateService
      }
    }
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    login: (userId, permissions) => {
      dispatch(createAction(ActionNames.LOGIN_SUCCESS, {userId: userId, token: Storage.auth, user: Storage.user, permissions: permissions}));
    },
    logout: () => {
      dispatch(createAction(ActionNames.LOGOUT, null, null));
    }
  };
}

const LoginValidator = connect(mapStateToProps, mapDispatchToProps)(Login);


export {LoginValidator}

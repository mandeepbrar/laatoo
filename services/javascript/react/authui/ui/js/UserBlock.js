import React from 'react';
import {connect} from 'react-redux';
import {Action} from 'reactwebcommon';
import {createAction } from 'uicommon';
import {ActionNames} from './actions';

class UserBlockUI extends React.Component {
  render() {
    let props = this.props
    let modprops = props.module.properties? props.module.properties: {}
    let logout = modprops.logoutText?  modprops.logoutText: "Logout";
    return props.loggedIn?
      <props.uikit.Block className={"userblock " + props.className}>
        <props.uikit.Block className="username">
        {Storage.userFullName? Storage.userFullName: Storage.userName}
        </props.uikit.Block>
        <Action name="logout" method={props.logout} className="logout">{logout}</Action>
      </props.uikit.Block>
    : null
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    loggedIn: state.Security.status == "LoggedIn"
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    logout: () => {
      dispatch(createAction(ActionNames.LOGOUT, null, null));
    }
  };
}

const UserBlock = connect(mapStateToProps, mapDispatchToProps)(UserBlockUI)

export {UserBlock}
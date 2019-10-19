import React from 'react';
const PropTypes = require('prop-types');
import ActionNames from './actions'
import {Action} from 'reactwebcommon';
import {connect} from 'react-redux';

function Initialize(appName, ins, mod, settings, def, req) {
}

function getLocation(href) {
    var l = document.createElement("a");
    l.href = href;
    return l;
}

function oauthLogin(cfg) {
  //oauthurl = Application.Security.googleAuthUrl
  let instance = window.open(cfg.oauthurl + cfg.realm, '_blank','height=500,width=400,toolbar=no,resizable=yes,menubar=no,location=0')
  var location = getLocation(cfg.oauthurl);
  loginSite = location.protocol+'//'+location.hostname+(location.port ? ':'+location.port: '');
}

const OauthBtnUI = (props) => {
  let btnClick = function() {
    oauthLogin(props)
  }
  window.addEventListener("message", function(ev) {
    if(ev.origin === loginSite && ev.data.message=="LoginSuccess") {
      props.handleOauthLogin(ev.data)
    }
  });
  console.log("returning Action from oauth btn", Action)
  return  (
    <Action widget="button" method={btnClick} action={{ actiontype: "method"}} className={"oauthlogin " + _tn(props.className, "")}>{props.children}</Action>
  )
}



const mapStateToProps = (state, ownProps) => {
  return {
    realm : Application.Security.realm
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  console.log("map dispatch of oauth login compoent")
  return {
    handleOauthLogin: (data) => {
      dispatch(createAction(ActionNames.LOGIN_SUCCESS, {userId: data.id, token: data.token, permissions: data.permissions}));
    }
  }
}

const OauthButton = connect(
  mapStateToProps,
  mapDispatchToProps
)(OauthBtnUI);


export {
  Initialize,
  OauthButton,
  oauthLogin
}

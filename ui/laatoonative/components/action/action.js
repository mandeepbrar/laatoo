'use strict';

import React from 'react';
import ActionButton from './actionbutton';
import { connect } from 'react-redux';
import ActionLink from './actionlink';
import {Application, Storage, createAction, hasPermission } from 'laatoocommon';
import { NavigationActions } from 'react-navigation'

class ActionComp extends React.Component {
  constructor(props) {
    super(props);
    this.renderView = this.renderView.bind(this);
    this.dispatchAction = this.dispatchAction.bind(this);
    this.actionFunc = this.actionFunc.bind(this);
    this.hasPermission = false
    console.log("Action", props)
    let action = Application.Actions[props.name];
    if(action) {
      this.action = action;
      this.hasPermission =  hasPermission(action.permission);
    }
  }

  dispatchAction() {
    let payload={};
    if(this.props.params) {
      payload = this.props.params;
    }
    this.props.dispatch(createAction(this.action.action, payload, {successCallback: this.props.successCallback, failureCallback: this.props.failureCallback}));
  }
  actionFunc() {
    if(this.props.confirm) {
      if(!this.props.confirm(this.props)) {
        return false
      }
    }
    console.log("exceuting action", this.action)
    switch(this.action.actiontype) {
      case "dispatchaction":
        this.dispatchAction();
      return false;
      case "method":
        let params = this.props.params
        let method = this.props.method
        method(params);
      return false;
/*      case "newwindow":
      if(this.action.url) {
        let formattedUrl = formatUrl(this.action.url, this.props.params);
        console.log(formattedUrl);
        //browserHistory.push({pathname: formattedUrl});
        window.open(formattedUrl);
        return false
      }*/
      default:
      if(this.action.target) {
//        let formattedUrl = formatUrl(this.action.url, this.props.params);
        console.log(this.action.target);
        //browserHistory.push({pathname: formattedUrl});
        this.props.dispatch(NavigationActions.navigate({ routeName: this.action.target, params: this.props.params }));
        //this.navigate(this.action.target, this.props.params);
      //  Router.redirect(formattedUrl);
      }
      return false;
    }
  }

  renderView() {
    if (!this.hasPermission) {
      return null;
    }
    let actionF = this.actionFunc;
    console.log("action props", this.props)
    switch(this.props.widget) {
      case 'button': {
        return (
          <ActionButton style={this.props.style} actionFunc={actionF} key={this.props.name +"_comp"}  transparent={this.props.transparent}
            iconRight={this.props.iconRight} iconLeft={this.props.iconLeft} actionchildren={this.props.children}>
          </ActionButton>
        )
      }
      default: {
        return (
          <ActionLink  style={this.props.style} actionFunc={actionF}  key={this.props.name +"_comp"} actionchildren={this.props.children}>
          </ActionLink>
        )
      }
    }
  }
  render() {
    return this.renderView();
  }
}
// Uncomment properties you need
ActionComp.propTypes = {
  name:  React.PropTypes.string.isRequired
};
// View.defaultProps = {};

const Action = connect()(ActionComp);
export {Action as Action} ;

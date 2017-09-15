'use strict';

import React from 'react';
import ActionButton from './ActionButton';
import { connect } from 'react-redux';
import ActionLink from './ActionLink';
import {createAction, formatUrl, hasPermission } from 'uicommon';

class ActionComp extends React.Component {
  constructor(props) {
    super(props);
    this.renderView = this.renderView.bind(this);
    this.dispatchAction = this.dispatchAction.bind(this);
    this.actionFunc = this.actionFunc.bind(this);
    this.hasPermission = false
    if(Application.Registry.Actions) {
      let action = Application.Registry.Actions[props.name];
      if(action) {
        this.action = action;
        this.hasPermission =  hasPermission(action.permission);
      }
    }
  }

  dispatchAction() {
    let payload={};
    if(this.props.params) {
      payload = this.props.params;
    }
    this.props.dispatch(createAction(this.action.action, payload, {successCallback: this.props.successCallback, failureCallback: this.props.failureCallback}));
  }
  actionFunc(evt) {
    evt.preventDefault();
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
      case "newwindow":
      if(this.action.url) {
        let formattedUrl = formatUrl(this.action.url, this.props.params);
        console.log(formattedUrl);
        //browserHistory.push({pathname: formattedUrl});
        window.open(formattedUrl);
        return false
      }
      default:
      if(this.action.url) {
        let formattedUrl = formatUrl(this.action.url, this.props.params);
        console.log(formattedUrl);
        //browserHistory.push({pathname: formattedUrl});
        Window.redirect(formattedUrl); //Router.redirect(formattedUrl);
      }
      return false;
    }
  }

  renderView() {
    console.log("checking permission", this.hasPermission)
    if (!this.hasPermission) {
      return null;
    }
    console.log("checking permission")
    let actionF = this.actionFunc;
    switch(this.props.widget) {
      case 'button': {
        return (
          <ActionButton className={this.props.className} actionFunc={actionF} key={this.props.name +"_comp"} actionchildren={this.props.children}>
          </ActionButton>
        )
      }
      case 'component':{
        console.log("component in ction", this.props)
        return (
          <this.props.component actionFunc={actionF}  key={this.props.name +"_comp"} actionchildren={this.props.children}/>
        )
      }
      default: {
        return (
          <ActionLink  className={this.props.className} actionFunc={actionF}  key={this.props.name +"_comp"} actionchildren={this.props.children}>
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
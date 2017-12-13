'use strict';

import React from 'react';
import ActionButton from './ActionButton';
import { connect } from 'react-redux';
import ActionLink from './ActionLink';
import {createAction, formatUrl, hasPermission } from 'uicommon';
import PropTypes from 'prop-types';

class ActionComp extends React.Component {
  constructor(props) {
    super(props);
    console.log("action comp creation", props)
    this.renderView = this.renderView.bind(this);
    this.dispatchAction = this.dispatchAction.bind(this);
    this.actionFunc = this.actionFunc.bind(this);
    this.hasPermission = false
    //let action = null
    if(props.action!=null) {
      this.action = props.action
    } else {
      this.action = _reg('Actions', props.name)
    }
    if(this.action) {
      this.hasPermission =  hasPermission(this.action.permission);
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
    console.log("action executed", this.props.name, this.props)
    evt.preventDefault();
    if(this.props.confirm) {
      if(!this.props.confirm(this.props)) {
        return false
      }
    }
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
    if (!this.hasPermission) {
      return null;
    }
    let children= this.props.children? this.props.children: this.props.label
    let actionF = this.actionFunc;
    switch(this.props.widget) {
      case 'button': {
        return (
          <ActionButton className={this.props.className} actionFunc={actionF} key={this.props.name +"_comp"} actionchildren={children}>
          </ActionButton>
        )
      }
      case 'component':{
        return (
          <this.props.component actionFunc={actionF}  key={this.props.name +"_comp"} actionchildren={children}/>
        )
      }
      default: {
        console.log("ffffffffffff")
        return (
          <ActionLink  className={this.props.className} actionFunc={actionF}  key={this.props.name +"_comp"} actionchildren={children}>
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
  name:  PropTypes.string.isRequired
};
// View.defaultProps = {};
console.log("react-redux connect in reactwebcommon", require('react-redux'));

const Action = connect()(ActionComp);
export {Action as Action} ;

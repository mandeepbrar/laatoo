'use strict';

import React from 'react';
import ActionButton from './ActionButton';
import { connect } from 'react-redux';
import ActionLink from './ActionLink';
import {Router} from 'redux-director';
import { createAction, formatUrl, hasPermission } from '../../utils';

class ActionComp extends React.Component {
  constructor(props) {
    super(props);
    this.renderView = this.renderView.bind(this);
    this.dispatchAction = this.dispatchAction.bind(this);
    this.actionFunc = this.actionFunc.bind(this);
    this.hasPermission = false
    let action = document.Actions[props.name];
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
    this.props.dispatch(createAction(this.action.action, payload));
  }
  actionFunc(evt) {
    console.log("clicked", this.props.params)
    evt.preventDefault();
    switch(this.action.actiontype) {
      case "dispatchaction":
        this.dispatchAction();
      return;
      case "method":
        let params = this.props.params
        let method = this.props.method
        console.log("method and params", method, params)
        method(params);
      return;
      default:
      if(this.action.url) {
        let formattedUrl = formatUrl(this.action.url, this.props.params);
        console.log(formattedUrl);
        //browserHistory.push({pathname: formattedUrl});
        Router.redirect(formattedUrl);
      }
      return;
    }
  }

  renderView() {
    if (!this.hasPermission) {
      return null;
    }
    let actionF = this.actionFunc;
    switch(this.props.widget) {
      case 'button': {
        return (
          <ActionButton className={this.props.className} actionFunc={actionF} key={this.props.name +"_comp"} actionchildren={this.props.children}>
          </ActionButton>
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

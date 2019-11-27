'use strict';

import React from 'react';
import ActionButton from './ActionButton';
import { connect } from 'react-redux';
import ActionLink from './ActionLink';
import {createAction, hasPermission } from 'uicommon';
import PropTypes from 'prop-types';

class ActionComp extends React.Component {
  constructor(props) {
    super(props);
    console.log("action comp creation", props)
    //this.renderView = this.renderView.bind(this);
    //this.dispatchAction = this.dispatchAction.bind(this);
    //this.actionFunc = this.actionFunc.bind(this);
    this.hasPermission = false
    //let action = null
    if(props.action!=null) {
      this.action = props.action
    } else {
      this.action = _reg('Actions', props.name)
    }
    if(this.props.method) {
      if(!this.action) {
        this.action = {actiontype: "method", method: this.props.method}
      } else {
        this.action.method = this.props.method
      }
    }
    console.log("action", this.action)
    if(this.action) {
      this.hasPermission =  hasPermission(this.action.permission);
    }
  }

  actionFunc = (evt) =>{

    console.log("action executed", this.props.name, this.props, this.action)

    if(this.props.confirm) {
      if(!this.props.confirm(this.props)) {
        return false
      }
    }
    let params = _tn(this.props.params, this.action.params)

    evt.preventDefault();
    this.props.dispatch(createAction("LAATOO_ACTION", params, this.action));
    return false
/*
    switch(this.action.actiontype) {
      case "dispatchaction":
        this.dispatchAction();
      return false;
      case "method":
        this.actionMethod(params);
      return false;
      case "interaction":
        if(!this.action.interactiontype) {
          return false
        }
        console.log("resolving panel", params, Window, "method", Window.resolvePanel)
        let comp = Window.resolvePanel("block", this.action.blockid, params)
        let onClose = this.props.onClose? this.props.onClose: _reg("Methods", this.action.onClose)
        Window.showInteraction(this.action.interactiontype, this.action.title, comp, onClose, this.action.actions, this.action.contentStyle, this.action.titleStyle)
        return false;
      case "newwindow":
      if(this.action.url) {
        let formattedUrl = formatUrl(this.action.url, params);
        console.log(formattedUrl);
        //browserHistory.push({pathname: formattedUrl});
        window.open(formattedUrl);
        return false
      }
      default:
      if(this.action.url) {
        let formattedUrl = formatUrl(this.action.url, params);
        console.log(formattedUrl);
        //browserHistory.push({pathname: formattedUrl});
        Window.redirect(formattedUrl, this.action.newpage); //Router.redirect(formattedUrl);
      }
      return false;
    }*/
  }

  render() {
    if (!this.hasPermission) {
      return null;
    }
    let children= this.props.children? this.props.children: this.props.label
    let widget = _tn(this.props.widget, this.action.widget)
    switch(widget) {
      case 'button': {
        return (
          <ActionButton flat={_tn(this.props.flat, this.action.flat)} className={this.props.className} actionFunc={this.actionFunc} key={this.props.name +"_comp"} actionchildren={children}>
          </ActionButton>
        )
      }
      case 'component':{
        return (
          <this.props.component actionFunc={this.actionFunc}  key={this.props.name +"_comp"} actionchildren={children}/>
        )
      }
      default: {
        return (
          <ActionLink  className={this.props.className} actionFunc={this.actionFunc}  key={this.props.name +"_comp"} actionchildren={children}>
          </ActionLink>
        )
      }
    }
  }
}
// Uncomment properties you need
ActionComp.propTypes = {
  name:  PropTypes.string.isRequired
};
// View.defaultProps = {};

const Action = connect()(ActionComp);
export {Action as Action} ;

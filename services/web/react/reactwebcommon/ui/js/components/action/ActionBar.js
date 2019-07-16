'use strict';

import React from 'react';
import {Action} from './Action';
import PropTypes from 'prop-types';

class ActionBar extends React.Component {
  constructor(props) {
    super(props)
    this.actions = []
    var desc = props.description;
    console.log("action bar", props);
    this.description = desc;
    this.className = props.className? props.className: desc.className;
    var comp = this;
    if(desc && desc.actions) {
        desc.actions.forEach(function(action){
          comp.actions.push(<Action name={action.name} label={action.label} widget={action.widget} className=" action "/>)
        })
    }
  }
  render() {
    return (
      <this.context.uikit.Block className={" actionbar " + this.className}>
      {this.actions}
      </this.context.uikit.Block>
    )
  }
}

ActionBar.contextTypes = {
  uikit: PropTypes.object
};

export {ActionBar as ActionBar}

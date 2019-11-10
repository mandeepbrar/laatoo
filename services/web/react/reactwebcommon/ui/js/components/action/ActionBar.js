'use strict';

import React from 'react';
import {Action} from './Action';
import PropTypes from 'prop-types';

class ActionBar extends React.Component {
  constructor(props) {
    super(props)
    this.actions = []
    console.log("creating action bar", props)
    if(props.children) {
      this.className = " actionbar " + _tn(props.className, "");
      this.actions = props.children
    } else {
      var desc = props.description;
      if(props.id || (desc && desc.id)) {
        let id = props.id? props.id : desc.id
        desc = _reg('ActionBar', id)
      }
      if(desc) {
        console.log("action bar", props, desc);
        this.description = desc;
        this.className = " actionbar " + _tn(props.className, _tn(desc.className,""));
        var comp = this;
        if(desc.actions) {
            desc.actions.forEach(function(action){
              comp.actions.push(<Action name={action.name} label={action.label} widget={action.widget} className=" action "/>)
            })
        }    
      }
    }
  }
  render() {
    return (
      <_uikit.Block className={this.className}>
      {this.actions}
      </_uikit.Block>
    )
  }
}

export {ActionBar as ActionBar}

'use strict';

import React from 'react';
const PropTypes = require('prop-types');

const ActionButton = (props, context) => {
  if(context.uikit && context.uikit.ActionButton) {
    return (
      <context.uikit.ActionButton className={props.className + " actionbutton"} onClick={props.actionFunc} btnProps={props}>
      {props.actionchildren}
      </context.uikit.ActionButton>
    )
  }
  return (
    <a className={props.className + " actionbutton"} onClick={props.actionFunc} role="button">
      {props.actionchildren}
    </a>
  )
}


ActionButton.contextTypes = {
  uikit: PropTypes.object
};
// Uncomment properties you need
ActionButton.propTypes = {
  actionFunc:  React.PropTypes.func.isRequired,
  actionchildren: React.PropTypes.oneOfType([
    React.PropTypes.array,
    React.PropTypes.string
  ])

};
// View.defaultProps = {};

export default ActionButton;

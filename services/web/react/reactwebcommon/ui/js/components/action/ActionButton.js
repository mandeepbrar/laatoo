'use strict';

import React from 'react';
import PropTypes from 'prop-types';

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
  actionFunc:  PropTypes.func.isRequired,
  actionchildren: PropTypes.oneOfType([
    PropTypes.array,
    PropTypes.string
  ])

};
// View.defaultProps = {};

export default ActionButton;

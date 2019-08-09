'use strict';

import React from 'react';
import PropTypes from 'prop-types';

const ActionButton = (props, context) => {
  if(_uikit.ActionButton) {
    return (
      <_uikit.ActionButton className={props.className + " actionbutton"} onClick={props.actionFunc} btnProps={props}>
      {props.actionchildren}
      </_uikit.ActionButton>
    )
  }
  return (
    <a className={props.className + " actionbutton"} onClick={props.actionFunc} role="button">
      {props.actionchildren}
    </a>
  )
}

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

'use strict';

import React from 'react';

const ActionButton = (props) =>(
  <a className={props.className + " actionbutton"} onClick={props.actionFunc} role="button">
    {props.actionchildren}
  </a>
)

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

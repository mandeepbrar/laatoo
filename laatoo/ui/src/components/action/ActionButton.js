'use strict';

import React from 'react';

const ActionButton = (props) =>(
  <a className={props.className + " btn btn-default"} onClick={props.actionFunc} role="button">
    {props.actionchildren}
  </a>
)

// Uncomment properties you need
ActionButton.propTypes = {
  actionFunc:  React.PropTypes.func.isRequired,
  actionchildren: React.PropTypes.array
};
// View.defaultProps = {};

export default ActionButton;

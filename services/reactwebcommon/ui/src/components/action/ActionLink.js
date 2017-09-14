'use strict';

import React from 'react';

const ActionLink = (props) =>(
  <a className={props.className +" actionlink"} href="javascript:void(0)" onClick={props.actionFunc}>
    {props.actionchildren}
  </a>
)

// Uncomment properties you need
ActionLink.propTypes = {
  actionFunc:  React.PropTypes.func.isRequired,
  actionchildren: React.PropTypes.oneOfType([
    React.PropTypes.array,
    React.PropTypes.string
  ])
};
// View.defaultProps = {};

export default ActionLink;

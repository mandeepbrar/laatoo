'use strict';

import React from 'react';

const ActionLink = (props) =>(
  <a className={props.className} href="#" onClick={props.actionFunc}>
    {props.actionchildren}
  </a>
)

// Uncomment properties you need
ActionLink.propTypes = {
  actionFunc:  React.PropTypes.func.isRequired,
  actionchildren: React.PropTypes.array
};
// View.defaultProps = {};

export default ActionLink;

'use strict';

import React from 'react';
import PropTypes from 'prop-types';

const ActionLink = (props) =>(
  <a className={props.className +" actionlink"} href="javascript:void(0)" onClick={props.actionFunc}>
    {props.actionchildren}
  </a>
)

// Uncomment properties you need
ActionLink.propTypes = {
  actionFunc:  PropTypes.func.isRequired,
  actionchildren: PropTypes.oneOfType([
    PropTypes.array,
    PropTypes.string
  ])
};
// View.defaultProps = {};

export default ActionLink;

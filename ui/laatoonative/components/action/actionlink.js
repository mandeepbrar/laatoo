'use strict';

import React from 'react';
import {Button} from 'native-base'

const ActionLink = (props) =>(
  <Button style={props.style} onPress={props.actionFunc}>
  {props.actionchildren}
  </Button>
)

// Uncomment properties you need
ActionLink.propTypes = {
  actionFunc:  React.PropTypes.func.isRequired,
  actionchildren: React.PropTypes.oneOfType([
    React.PropTypes.array,
    React.PropTypes.string
  ])
}

export default ActionLink;

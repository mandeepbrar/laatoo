'use strict';

import React from 'react';
import {Button} from 'native-base'

const ActionButton = (props) =>{
  console.log("props of action button", props)
  return (
    <Button style={props.style}  transparent={props.transparent} iconRight={props.iconRight} iconLeft={props.iconLeft} onPress={props.actionFunc}>
    {props.actionchildren}
    </Button>

  )
}

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

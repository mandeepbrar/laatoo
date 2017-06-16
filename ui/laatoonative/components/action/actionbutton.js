'use strict';

import React from 'react';
import {Button} from 'native-base'

const ActionButton = (props) =>{
  return (
    <Button style={props.style}  transparent={props.transparent} iconRight={props.iconRight} iconLeft={props.iconLeft} onPress={props.actionFunc}>
    {props.actionchildren}
    </Button>
  )
}

export default ActionButton;

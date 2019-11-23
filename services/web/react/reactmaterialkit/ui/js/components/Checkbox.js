import {Checkbox} from '@material-ui/core';
import React from 'react';

const CheckboxComp = (props)=> {
  console.log("props of checkbox", props)
  return(
    <Checkbox  {...props} hintText={props.placeholder} />
  )
}

export {CheckboxComp as Checkbox}

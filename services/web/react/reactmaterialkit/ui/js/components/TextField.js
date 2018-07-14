import {TextField} from 'material-ui';
import React from 'react';

const TextFieldComp = (props)=> {
  return(
    <TextField  {...props} hintText={props.placeholder} />
  )
}

export {TextFieldComp as TextField}

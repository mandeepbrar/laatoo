import MUIcon from '@material-ui/core/Icon';
import React from 'react';

const Icons = {
  CloseIcon: (props)=>(<MUIcon className="fa fa-close"/>),
  NewIcon: (props)=>(<MUIcon className="fa fa-plus-circle"/>),
  DeleteIcon:  (props)=>(<MUIcon className="fa fa-trash"/>),
  EditIcon: (props)=>(<MUIcon className="fa fa-edit"/>)
}

export {Icons}

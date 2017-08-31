'use strict';

import React from 'react';
import {Button} from 'native-base'

const ActionLink = (props) =>(
    <Button style={props.style}  transparent iconLeft onPress={props.actionFunc}>
    {props.actionchildren}
    </Button>
)

export default ActionLink;

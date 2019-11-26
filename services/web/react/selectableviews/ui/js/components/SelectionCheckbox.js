import React, { useState } from 'react';
import {SelectableViewContext} from './SelectableViewContext';

class SelectionCheckbox extends React.Component {
    constructor(props, context) {
        super(props)
        //props.id
        console.log("context in selection checkbox", context, this.context)
        let {register, setItemSelection, isSelected} = context
        this.setItemSelection = setItemSelection
        register(props.id, props.data)
        this.state = {checked: isSelected(props.id)}
    }

    onChange = (evt) => {
        this.setItemSelection(this.props.id, this.props.data, evt.target.checked)
        this.setState({checked: evt.target.checked})
    }

    render() {
        return (
            <_uikit.Checkbox value={this.state.checked} onChange={this.onChange} {...this.props}/>
        )
    }
}

//        return React.cloneElement(this.props.children, {selection: this.state.selection, setItemSelection: this.setItemSelection})

SelectionCheckbox.contextType = SelectableViewContext;

export {
    SelectionCheckbox
}

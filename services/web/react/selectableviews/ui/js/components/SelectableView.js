import React, { useState } from 'react';
import {SelectableViewContext} from './SelectableViewContext';
import {Panel} from 'reactpages';

class SelectableView extends React.Component {
    constructor(props) {
        super(props)
        console.log("selectable view props", props)
        this.allItems = {}
        this.childcontext = {
            register: this.register,
            setItemSelection: this.setItemSelection,
            isSelected: this.isSelected
        }
        this.selectedItems = {}
        if(props.value) {
            props.value.forEach(element => {
                this.selectedItems[element.Id] = element
            });
        }
    }

    register = (id, data) => {
        console.log("registering item", id, data)
        //keeping for later use if all item status is needed
        //this.allItems[id] = data
    }

    isSelected = (id) => {
        console.log("checking id", id, this.selectedItems)
        return (id in this.selectedItems)
    }

    setItemSelection = (id, data, selection) => {
        if(selection) {
            this.selectedItems[id] = data
        } else {
            delete this.selectedItems[id]
        }
        console.log("items selected", this.selectedItems)
        if(this.props.onChange) {
            this.props.onChange(Object.values(this.selectedItems), this.props.name)
        }
    }

    change = (value, name, evt) => {
       /* console.log("changing entity", value, this.props, name, evt)
        this.props.onChange(value, this.props.name, evt)
        this.setState(Object.assign({}, this.state, {value}))*/
    }

    render() {
        console.log("selectable views", this.props, "context", this.childcontext)
        return (
            <SelectableViewContext.Provider value={this.childcontext}>
                {
                    this.props.children? this.props.children:
                    <Panel type="block" id={this.props.blockid} className=" m20 tableview "/>
                }
            </SelectableViewContext.Provider>
        )
    }
}

export {
    SelectableView
}
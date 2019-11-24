import React, {useState} from 'react'
const PropTypes = require('prop-types');
import {View} from 'laatooviews';
import {Panel} from 'reactpages';
import {Action} from 'reactwebcommon';

class TreeView extends React.Component {
    constructor(props) {
        super(props)
        console.log("props of tree view", props, "view", View, "panel", Panel, "action1", Action)
        //let viewheader = props.header
        this.treeHeader = (<Panel id={props.header} type="block"/>)
        //let postArgs = props.postArgs
    }
    render() {
        let props = this.props
        return (
            <View header={this.treeHeader} className={props.className} postArgs={props.postArgs} instance={props.name} description={{serviceName: props.serviceName}}>
                <TreeViewRow rowClassName=" tablerow " postArgs={props.postArgs} idfield={_tn(props.idfield, "Id")} 
                    parentField={props.parentField} row={props.row} serviceName={props.serviceName} viewdepth={0} />
            </View>
        )
    }
}

class TreeViewRow extends React.Component {
    constructor(props) {
        super(props)
        console.log("row props", props)
        this.state = {expanded: false}
    }

    expandedContent = (depth, id) => {
        let props = this.props
        console.log("expanding", depth, "id", id, "props", props)
        let newargs = {}
        newargs[props.parentField] = id
        let postArgs = Object.assign({}, props.postArgs, newargs)
        console.log("expanding", depth, "id", id, "props", props, "postArgs", postArgs)
        return (
            <View contentOnly={true} instance={id} postArgs={postArgs} description={{serviceName: props.serviceName}}>
                <TreeViewRow row={props.row} viewdepth={depth + 1} postArgs={postArgs} idfield={props.idfield} 
                    serviceName={props.serviceName} rowClassName={props.rowClassName} parentField={props.parentField} instance={id}/>
            </View>        
        )
    }
    setExpanded = () => {
        console.log("setting expanded")
        this.setState({expanded: !this.state.expanded})
    }
    expansionComp = () => {
        let props = this.props
        return (
            <Action className="nodecoration" action={{actiontype:"method", method: this.setExpanded}}>
            {props.viewdepth? 
            <_uikit.Block className=" inlineblock " style={{width: props.viewdepth*15}}>&nbsp;</_uikit.Block> 
            : null
            }
            {
            this.state.expanded? 
            <_uikit.Icon className="fa fa-caret-down"/> 
            : <_uikit.Icon className="fa fa-caret-right"/>
            }
            </Action>    
        )
    }
    render() {
        console.log("row expanded", this.state.expanded)
        let props = this.props
        let comps = []
        comps.push(
            <Panel type="block" data={props.data} id={props.row} viewdepth={props.viewdepth} expansionComponent={this.expansionComp()} />
        )
        if(this.state.expanded) {
            comps.push(this.expandedContent(props.viewdepth, props.data[props.idfield]))
        }
        return comps
    }
}


export {
    TreeView as TreeView
}
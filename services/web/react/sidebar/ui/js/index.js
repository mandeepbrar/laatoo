import React from 'react';
import Sidebar from 'react-sidebar';
import {Block} from 'reactwebcommon';

function Initialize(app, ins, mod, s, def, req) {

    Window.showDrawer = function(title, component, onClose, actions, contentStyle, titleStyle) {
        Window.showInteraction("Drawer", title, component, onClose, actions, contentStyle, titleStyle)
    }
    Window.closeDrawer = function() {
        Window.closeInteraction("Drawer")
    }
    let injectdrawer = function(store, Application, Uikit, Theme, Router) {
        console.log("inject drawer called", Uikit)
        Uikit.Drawer = Drawer
    }
    console.log("registering for injectdrawer in boot")
    _r("Bootmethods", "injectdrawer", injectdrawer)
}


class Drawer extends React.Component {
    constructor(props) {
        super(props)
    }
    render() {
        console.log("drawer props", this.props)
        let content = ( 
            <Block title={this.props.title} className="dialog" closeBlock={this.props.onClose} contentStyle={this.props.contentStyle}>
                {this.props.component}
            </Block>
        )
        return (
            <Sidebar
                sidebar={content}
                open={true}
                pullRight={true}
                //onSetOpen={this.onSetSidebarOpen}
                styles={{ sidebar: { background: "white" } }}
            >
            </Sidebar>
        )        
    }
}

export {
    Initialize,   
    Drawer
}

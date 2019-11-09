import React from 'react';

class Menu extends React.Component {
    constructor(props) {
        super(props)
        console.log("menu props", props)
        this.menu = _reg("Menus", props.id)
        if(this.menu && this.menu.items) {
            console.log("iterating menu items ", this.menu.items)
            this.menu.items.forEach(function(menuItem){
                if(menuItem.page) {
                    menuItem.action = "Page_" + menuItem.page
                }
            })  
            this.menuitems = this.menu.items
        } else {
            this.menuitems = []
        }
    }
    render() {
        return <_uikit.Navbar items={this.menuitems} vertical={this.props.vertical}/>
    }
}

export { Menu as Menu}

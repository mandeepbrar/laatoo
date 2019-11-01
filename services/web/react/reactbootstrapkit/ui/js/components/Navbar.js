import React from 'react';
import {Block, Action} from'reactwebcommon';
import {Navbar as BSNavBar, Nav} from 'react-bootstrap'

function getNavButton(isVertical) {
  return function(props) {
    return (
      <Nav className={isVertical?"vertnavbutton":"horiznavbutton"}>
        <Nav.Link onClick={props.actionFunc}>{props.actionchildren}</Nav.Link>
      </Nav>
    )
  }
}


class Navbar extends React.Component {
  render() {
    let isVertical = this.props.vertical
    let className=isVertical?"vertnavitem":"horiznavitem"
    let navClass=isVertical?"vertnavbar":"horiznavbar"
    let items=[]
    console.log("navbar items", items, this.props.items)
    if(this.props.items) {
      this.props.items.forEach(function(item){
        let icon
        let iconStyle = item.iconSize? {fontSize: item.iconSize}: null
        icon = item.icon || item.iconClass?<i className={item.iconClass} style={iconStyle}>{item.icon}</i>: null
        items.push(          
          <Action widget="component" vertical={isVertical} name={item.action} component={getNavButton(isVertical)}><Block className={className}><Block className="icon">{icon?icon:null}</Block><Block>{item.title}</Block></Block></Action>
        )
      });
      console.log("navbar items", items, this.props.items)
    }
    let stacked=this.props.vertical?["stacked"]:[""]
    return(
      <BSNavBar bsStyle="pills" className={navClass} {...stacked}>
        {items}
      </BSNavBar>
    )
  }
}

export default Navbar
import React from 'react';
import {Block, Action} from'reactwebcommon';
import {Nav, NavItem} from 'react-bootstrap'

function getNavButton(isVertical) {
  return function(props) {
    return (
      <NavItem className={isVertical?"vertnavbutton":"horiznavbutton"} onClick={props.actionFunc}>{props.actionchildren}</NavItem>
    )
  }
}


class Navbar extends React.Component {
  render() {
    let isVertical = this.props.vertical
    let className=isVertical?"vertnavitem":"horiznavitem"
    let navClass=isVertical?"vertnavbar":"horiznavbar"
    let items=[]
    if(this.props.items) {
      this.props.items.forEach(function(item){
        let icon
        let iconStyle = item.iconSize? {fontSize: item.iconSize}: null
        icon = item.icon || item.iconClass?<i className={item.iconClass} style={iconStyle}>{item.icon}</i>: null
        items.push(          
          <Action widget="component" vertical={isVertical} name={item.action} component={getNavButton(isVertical)}><Block className={className}><Block className="icon">{icon?icon:null}</Block><Block>{item.title}</Block></Block></Action>
        )
      });
    }
    let stacked=this.props.vertical?["stacked"]:[""]
    return(
      <Nav bsStyle="pills" className={navClass} {...stacked}>
        {items}
      </Nav>
    )
  }
}

export default Navbar
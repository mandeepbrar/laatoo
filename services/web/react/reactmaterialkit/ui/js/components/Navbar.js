import React from 'react';
import {Action, Block} from'reactwebcommon';
import Button from '@material-ui/core/Button';
import Icon from '@material-ui/core/Icon';

function getNavButton(isVertical) {
  return function(props) {
    return (
      <Button className={isVertical?"vertnavbutton":"horiznavbutton"} onClick={props.actionFunc}>{props.actionchildren}</Button>
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
        icon = item.icon || item.iconClass?<Icon className={item.iconClass} style={iconStyle}>{item.icon}</Icon>: null
        items.push(          
          <Action widget="component" vertical={isVertical} name={item.action} component={getNavButton(isVertical)}><div className={className}><div className="icon">{icon?icon:null}</div><div>{item.title}</div></div></Action>
        )
      });
    }
    return(
      <Block className={navClass}>
        {items}
      </Block>
    )
  }
}

export default Navbar

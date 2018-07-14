import React from 'react';
import {Nav, NavItem} from 'react-bootstrap'
import {Action} from'reactwebcommon';

const Navbutton=(props)=>(
  <NavItem onClick={props.actionFunc}>{props.actionchildren}</NavItem>
)

class Navbar extends React.Component {
  render() {
    let items=[]
    if(props.items) {
      props.items.forEach(function(item){
        items.push(
          <Action widget="component" name={item.action} component={Navbutton}>{item.title}</Action>
        )
      });
    }
    let stacked=this.props.vertical?["stacked"]:[""]
    return(
      <Nav bsStyle="pills" {...stacked}>
        {items}
      </Nav>
    )
  }
}

export default Navbar

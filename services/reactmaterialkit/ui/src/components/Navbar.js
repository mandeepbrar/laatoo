import React from 'react';
import {Action} from'reactwebcommon';
import FlatButton from 'material-ui/FlatButton';

const Navbutton=(props)=>(
    <FlatButton className={props.vertical?"vertnavbutton":"horiznavbutton"} onClick={props.actionFunc}>{props.actionchildren}</FlatButton>
  )


class Navbar extends React.Component {
  render() {
    let isVertical = this.props.vertical
    let className=isVertical?"vertnavitem":"horiznavitem"
    let items=[]
    if(this.props.items) {
      this.props.items.forEach(function(item){
        items.push(
          <div className={className}>
            <Action widget="component" vertical={isVertical} name={item.action} component={Navbutton}>{item.title}</Action>
          </div>
        )
      });
    }
    return(
      <div className='navbar'>
        {items}
      </div>
    )
  }
}

export default Navbar

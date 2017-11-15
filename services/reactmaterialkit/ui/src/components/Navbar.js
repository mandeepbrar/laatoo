import React from 'react';
import {Action} from'reactwebcommon';
import Button from 'material-ui/Button';

const Navbutton=(props)=>(
    <Button className={props.vertical?"vertnavbutton":"horiznavbutton"} onClick={props.actionFunc}>{props.actionchildren}</Button>
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

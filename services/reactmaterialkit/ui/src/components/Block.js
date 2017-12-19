import React from 'react';
import {Icons} from './Icons';
import IconButton from 'material-ui/IconButton';

const Block=(props) =>  {
  let cl = props.className? props.className:""
  let contentClass = props.contentClass?props.contentClass:""
  if(props.title) {
    return (
      <div style={props.style} className={"block "+cl}>
        <div className="titlebar">
          <div className="title left">
          {props.title}
          </div>
          {
            props.closeBlock?  <IconButton className="right close fa fa-close" onClick={props.closeBlock}/>  : null
          }
        </div>
        <div style={props.contentStyle} className={"blockcontent "+contentClass}>
        {props.children}
        </div>
      </div>
    )
  }
  else {
    return (
      <div style={Object.assign({}, props.contentStyle, props.style)} className={"block "+contentClass+" "+cl}>
       {props.children}
      </div>
    )
  }
}

export {Block}

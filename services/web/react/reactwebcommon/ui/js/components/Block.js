import React from 'react';

const Block=(props) =>  {
  let cl = props.className? props.className:""
  let contentClass = props.contentClass?props.contentClass:""
  if(props.title || props.titleBarActions || props.closeBlock) {
    return (
      <div style={props.style} className={"block "+cl}>
        <div className="titlebar">
          <div className="title left">
          {props.title}
          </div>
          {
            props.titleBarActions? <div className="right">{props.titleBarActions}</div>: (props.closeBlock?  <_uikit.Icon className="right close fa fa-close" onClick={props.closeBlock}/>  : null)
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
      <div style={Object.assign({}, props.contentStyle, props.style)} onClick={props.onClick} className={"block "+contentClass+" "+cl}>
       {props.children}
      </div>
    )
  }
}

export {Block}

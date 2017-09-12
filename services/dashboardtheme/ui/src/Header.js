import React from 'react';
import {Image} from 'reactwebcommon';

const Header = (props) => {
  let hs = props.headerProps;
  console.log("header properties", hs, props)
  return (
    <div className={hs.className?hs.className:'header'}>
      <div className="logo">
        {hs.image?<div className="image"><Image src={hs.image}/></div>:null}
        {hs.title?<div className="title">{hs.title}</div>:null}
      </div>
    </div>
  )
}

export default Header

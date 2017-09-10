import React from 'react';
import {Image} from 'reactwebcommon';

const Header = (props) => (
  <div className={props.headerclass?props.headerclass:'header'}>
    <div className="logo">
    {
      props.image?
      <Image src={props.image} />
      :<div className="title">{props.title}</div>
    }
    </div>
  </div>
)

export default Header

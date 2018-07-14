import React from 'react';

const Welcome = (props) => (
    <div className='welcomepage'>
    {props.modProps.welcome.text}
    </div>
  )
export default Welcome

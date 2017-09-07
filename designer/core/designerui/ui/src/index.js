import React from 'react';
import ReactDOM from 'react-dom';
import App from './app'

function StartApplication () {
  console.log("starting applicaiton")
  // Render the main component into the dom
  ReactDOM.render(<App/>, document.getElementById('app'));
}

console.log("defined start applicaiton", StartApplication);

export {
  StartApplication
}

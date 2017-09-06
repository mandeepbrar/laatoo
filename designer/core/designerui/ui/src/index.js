import React from 'react';
import ReactDOM from 'react-dom';
import App from './app'

window.StartApplication = function() {
  console.log("starting applicaiton")
  // Render the main component into the dom
  ReactDOM.render(<App/>, document.getElementById('app'));
}

console.log("defined start applicaiton", window.StartApplication);

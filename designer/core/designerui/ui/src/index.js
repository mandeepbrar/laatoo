import React from 'react';
import ReactDOM from 'react-dom';
import App from './app'

window.StartApplication = function() {
  // Render the main component into the dom
  ReactDOM.render(<App/>, document.getElementById('app'));
}

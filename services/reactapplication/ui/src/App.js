import React from 'react';
const PropTypes = require('prop-types');

class App extends React.Component {
  getChildContext() {
   return {uikit: this.props.uikit};
  }
  render() {
    return (
      <this.props.theme uikit={this.props.uikit}/>
    )
  }
}

App.childContextTypes = {
  uikit: PropTypes.string
};

export {
  App
}

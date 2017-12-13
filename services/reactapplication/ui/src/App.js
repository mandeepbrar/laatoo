import React from 'react';
const PropTypes = require('prop-types');

class App extends React.Component {
  getChildContext() {
   return {uikit: this.props.uikit, router: this.props.router};
  }
  render() {
    return (
      <this.props.theme router={this.props.router} uikit={this.props.uikit}/>
    )
  }
}


/*

<Dialogs uikit={this.props.uikit} />
<Messages uikit={this.props.uikit} />
*/


App.childContextTypes = {
  uikit: PropTypes.object,
  router: PropTypes.object
};

export {
  App
}

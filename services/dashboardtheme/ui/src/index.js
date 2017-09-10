import React from 'react'
import {View} from 'redux-director'
import Header from './Header'
import './styles/app.scss'

var module = this;

function Initialize(appName, settings) {
  module.settings = settings
  console.log("Initializing dashboard theme with settings ","module", module.settings)
}

class DashboardTheme extends React.Component {
  render() {
    console.log("props of dashboard theme", this.props, module.settings)
    return (
      <div className="dashboard ">
        <Header image={module.settings.headerimage} title={module.settings.headertitle} headerclass={module.settings.headerclass}/>
        <div className="body">
          <div className="col-sm-4">
            <View name="menu"  />
          </div>
          <div className="col-sm-8">
            <View name="main"  />
          </div>
        </div>
        <div className="row">
        </div>
      </div>
    )
  }
}

export {
  Initialize ,
  DashboardTheme as Theme
}

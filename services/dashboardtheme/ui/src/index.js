import React from 'react'
import {View} from 'redux-director'
import Header from './Header'
import './styles/app.scss'

var module = this;

function Initialize(appName, settings) {
  console.log(document.InitConfig);
  let dashProps = document.InitConfig.Properties["dashboardtheme"]
  module.properties = settings.propertiesOverrider ? Object.assign({}, dashProps, document.InitConfig.Properties[settings.propertiesOverrider]) : dashProps;
  console.log("Initializing dashboard theme with settings ", module.properties)
}

class DashboardTheme extends React.Component {
  render() {
    console.log("props of dashboard theme", this.props, module.properties)
    return (
      <div className="dashboard ">
        <Header headerProps={module.properties.header} />
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

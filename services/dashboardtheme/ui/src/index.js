import React from 'react'
import {View} from 'redux-director'

function Initialize(appName, settings) {
  console.log("Initializing dashboard theme with settings ", settings)
}

class DashboardTheme extends React.Component {
  render() {
    return (
      <div className="dashboard ">
        <div className="row">
        </div>
        <div className="row">
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

export default {
  Initialize : Initialize,
  Theme : DashboardTheme
}

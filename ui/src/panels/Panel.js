import React from 'react';
import PanelsRegistry from './Registry'


class Panel extends React.Component {
  render() {
    let panelConf = this.props.conf
    let widget = null
    if(panelConf && panelConf.type) {
      let panelRenderer = PanelsRegistry.GetPanel(panelConf.type)
      widget = panelRenderer(panelConf.name, panelConf.params, panelConf.styleAttributes)
    }
    return (
      <div>
      {widget}
      </div>
    )
  }
}

export default Panel

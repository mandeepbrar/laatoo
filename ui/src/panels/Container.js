import React from 'react';
import PanelsRegistry from './Registry'
import Panel from './Panel'

class Container extends React.Component {
  render() {
    let children = new Array()
    for (i=0;i<this.props.children.length;i++) {
      let childConf = children[i]
      children.push(
        <Panel conf={childConf}/>
      )
    }
    return (
      <div className={this.props.styleAttributes.className}>
        {children}
      </div>
    )
  }
}
/*
{
  name:some,
  type: Container
  contentparams: {
    children: [

    ]
  }
  styleattributes : {

  }
}
*/
//renderFunc = function (panelname, contentparams, styleattribs )
function renderContainer(panelname, contentparams, styleattribs ) {
  if(!contentparams || !contentparams.children) {
    return null
  }
  let children = contentparams.children
  return (
    <Container children={children} name={panelname} styleAttributes={styleattribs}/>
  )
}

PanelsRegistry.RegisterPanel("Container", renderContainer)

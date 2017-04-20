import React from 'react';
import Container from './Container'
import PanelsRegistry from './Registry'

const ColumnLayout = (props) => {
  let children = new Array()
  for (i=0;i<this.props.columns.length;i++) {
    let childConf = this.props.columns[i]
    children.push(
      <div className={this.props.styleAttributes.columns[i].className}>
        <Container conf={childConf}/>
      </div>
    )
  }
  return (
    <div className={this.props.styleAttributes.className}>
      {children}
    </div>
  )
}

/*
{
  name:some,
  type: ColumnLayout
  contentparams: {
    columns: [
      {

      },

    ]
  }
  styleAttributes: {
    columns: [
      {
        "className":
      }
    ]
  }
}
*/
//renderFunc = function (panelname, contentparams, styleattribs )
function renderColumnLayout(panelname, contentparams, styleattribs ) {
  if(!contentparams || !contentparams.columns) {
    return null
  }
  let columns = contentparams.columns
  return (
    <ColumnLayout columns={columns} name={panelname} styleAttributes={styleattribs}/>
  )
}

PanelsRegistry.RegisterPanel("ColumnLayout", renderColumnLayout)

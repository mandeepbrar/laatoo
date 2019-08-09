import 'styles/app.scss'

import React from 'react';
const PropTypes = require('prop-types');

function Initialize(appName, ins, mod, settings, def, req) {
  module.properties = Application.Properties[ins]
  module.settings = settings;
  console.log("entity initialize", settings, settings.object)
  if(settings.object) {
    let defDisp = settings.object + "_default"
    let disp = _reg("Blocks", defDisp)
    if(!disp) {
      _r('Blocks', defDisp, function(ctx, desc) {
        console.log("rendering default display", defDisp, ctx, desc)
        return <h1>default display</h1>
      })
    }
  }
}

Window.displayDefaultEntity = function(ctx, desc) {
  return <DefaultEntityDisplay desc={desc} ctx={ctx} />
}

class DefaultEntityDisplay extends React.Component {
  createField = (fieldVal, field,  level, ctx, desc) => {
    let fldDisp = this.createObjFields(fieldVal, level+1, ctx, desc)
    return (<_uikit.Block className={"field " + field}>
       <_uikit.Block className="name">
       {field}
       </_uikit.Block>
       <_uikit.Block className="value">
       {fldDisp}
       </_uikit.Block>
     </_uikit.Block>)
  }
  createObjFields = (obj, level, ctx, desc) => {
    if(obj==null) return null;
    if(obj instanceof Array) {
      let fields = new Array()
      for(var i=0;i<obj.length;i++) {
        fields.push(<_uikit.Block className="entityarrayitem">{this.createObjFields(obj[i], level+1, ctx, desc)}</_uikit.Block>)
      }
      return fields
    } else if(typeof(obj) == "object") {
      let fields = new Array()
      let tabs = new Array()
      let dispobj = this
      Object.keys(obj).forEach(function(field) {
        let fieldVal = obj[field]
        let dispElems = dispobj.createField(fieldVal, field, level, ctx, desc)
        console.log("field", field, "fieldVal", fieldVal," level ", level)
        if ((fieldVal instanceof Array) && (level == 0)) {
          tabs.push(<_uikit.Tab label={field}>{dispElems}</_uikit.Tab>)
        } else {
          fields.push(dispElems)
        }
      })
      return level!=0?fields:<_uikit.Tabset><_uikit.Tab label="General">{fields}</_uikit.Tab>{tabs}</_uikit.Tabset>
    } else {
      return obj
    }
  }
  render() {
    let {ctx, desc} = this.props
    console.log(ctx, desc, _uikit)
    return <_uikit.Block className="entity ">
      {this.createObjFields(ctx.data, 0, ctx, desc)}
    </_uikit.Block>
  }
}

export {
  Initialize
}

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
      _r('Blocks', defDisp, function(ctx, desc, uikit) {
        console.log("rendering default display", defDisp, ctx, desc, uikit)
        return <h1>default display</h1>
      })
    }
  }
}

Window.displayDefaultEntity = function(ctx, desc, uikit) {
  return <DefaultEntityDisplay desc={desc} uikit={uikit} ctx={ctx} />
}

class DefaultEntityDisplay extends React.Component {
  createField = (fieldVal, field,  level, ctx, desc, uikit) => {
    let fldDisp = this.createObjFields(fieldVal, level+1, ctx, desc, uikit)
    return (<uikit.Block className={"field " + field}>
       <uikit.Block className="name">
       {field}
       </uikit.Block>
       <uikit.Block className="value">
       {fldDisp}
       </uikit.Block>
     </uikit.Block>)
  }
  createObjFields = (obj, level, ctx, desc, uikit) => {
    if(obj==null) return null;
    if(obj instanceof Array) {
      let fields = new Array()
      for(var i=0;i<obj.length;i++) {
        fields.push(<uikit.Block className="entityarrayitem">{this.createObjFields(obj[i], level+1, ctx, desc, uikit)}</uikit.Block>)
      }
      return fields
    } else if(typeof(obj) == "object") {
      let fields = new Array()
      let tabs = new Array()
      let dispobj = this
      Object.keys(obj).forEach(function(field) {
        let fieldVal = obj[field]
        let dispElems = dispobj.createField(fieldVal, field, level, ctx, desc, uikit)
        console.log("field", field, "fieldVal", fieldVal," level ", level)
        if ((fieldVal instanceof Array) && (level == 0)) {
          tabs.push(<uikit.Tab label={field}>{dispElems}</uikit.Tab>)
        } else {
          fields.push(dispElems)
        }
      })
      return level!=0?fields:<uikit.Tabset><uikit.Tab label="General">{fields}</uikit.Tab>{tabs}</uikit.Tabset>
    } else {
      return obj
    }
  }
  render() {
    let {ctx, desc, uikit} = this.props
    console.log(ctx, desc, uikit)
    return <uikit.Block className="entity ">
      {this.createObjFields(ctx.data, 0, ctx, desc, uikit)}
    </uikit.Block>
  }
}

export {
  Initialize
}

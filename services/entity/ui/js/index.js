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
/*

  layoutFields = (fldToDisp, flds, className) => {
    let fieldsArr = new Array()
    let comp = this
    fldToDisp.forEach(function(k) {
      let fd = flds[k]
      let cl = className? className + " m10": "m10"
      fieldsArr.push(  <Field key={fd.name} name={fd.name} formValue={comp.state.formValue} {...comp.parentFormProps} time={comp.state.time} className={cl}/>      )
    })
    return fieldsArr
  }

  fields = () => {
    let desc = this.props.description
    console.log("desc of form ", desc)
    let comp = this
    if(desc && desc.fields) {
      let flds = desc.fields
      if(flds) {
        if(desc.info && desc.info.tabs) {
          let tabs = new Array()
          let tabsToDisp = desc.info && desc.info.tabs? desc.info.layout: Object.keys(desc.info.tabs)
          tabsToDisp.forEach(function(k) {
            let tabFlds = desc.info.tabs[k];
            if(tabFlds) {
              let tabArr = comp.layoutFields(tabFlds, flds, "tabfield formfield")
              tabs.push(
                <comp.uikit.Tab label={k} time={comp.state.time} value={k}>
                  {tabArr}
                </comp.uikit.Tab>
              )
            }
          })
          let vertical = desc.info.verticaltabs? true: false
          return (
            <this.uikit.Tabset vertical={vertical} time={comp.state.time}>
              {tabs}
            </this.uikit.Tabset>
          )
        } else {
          let fldToDisp = desc.info && desc.info.layout? desc.info.layout: Object.keys(flds)
          let className=comp.props.inline?"inline formfield":"formfield"
          return this.layoutFields(fldToDisp, flds, className)
        }
      }
    }
    return null
  }*/

export {
  Initialize
}

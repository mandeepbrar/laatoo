import React from 'react'
const PropTypes = require('prop-types');

var module;

class Panel extends React.Component {
  constructor(props, ctx) {
    super(props)
    this.uikit = ctx.uikit
    let desc = props.description
    if(!desc && props.id) {
      desc = _reg('Panels', props.id)
    }
    //console.log("print id ", desc.id)
    if(!desc) {
      console.trace()
      console.log("returning without description", props, desc)
      return
    }
    console.log("creating panel", desc, props)
    let className = "panel "
    if(desc.id) {
      className = className + " " +desc.id
    }
    if(desc.name) {
      className = className + " " +desc.name
    }
    if(desc && (typeof(desc) == 'string')) {
      this.processBlock(desc, props, className)
    } else if(desc){
      switch(desc.type) {
        case "view":
          className = className + " view "
          this.processView(desc,  props, className)
          break;
        case "entity":
          className = className + " entity "
          this.processEntity(desc, props, className)
          break;
        case "form":
          className = className + " form "
          this.processForm(desc, props, className)
          break;
        case "block":
          className = className + " block "
          this.processBlock(desc, props, className)
          break;
        case "layout":
          className = className + " panel "
          this.processLayout(desc, props, className)
          break;
        case "component":
          if(desc.component) {
            this.getView = function(props, context, state) {
              return desc.component
            }
          } else {
            let comp = this.getComponent(desc.module, desc.component, module.req)
            this.getView = function(props, context, state) {
              return React.createElement(comp, {className: className})
            }
          }
          break;
        default:
          let processor = _reg("PanelTypes", desc.type)
          //let comp = processor.getComponent(desc)
          this.getView = function(props, context, state) {
            return processor.getComponent(desc, props, context, state)
          }
          break;
      }
    }
  }

  getPanelItems = (desc) => {
    if(!desc) {
      return null
    }
    if(desc instanceof Array) {
      let items = desc.map(function(item) {
        return <Panel description={item}/>
      });
      return items
    } else {
      return <Panel description={desc}/>
    }
    return null
  }

  processLayout = (desc, props, className) => {
    let panelDesc = desc
    if(desc.id) {
      panelDesc = _reg("Panels", desc.id)
    }
    if(!panelDesc.layout) {
      return
    }
    var block = this
    var layout = null
    var panelComp = (id)=> {
      return panelDesc[id]?
      <block.uikit.Block className={id}>
        {block.getPanelItems(panelDesc[id])}
      </block.uikit.Block>
      :null
    }

    switch(panelDesc.layout) {
      case "2col": {
        layout=( <this.uikit.Block className={className + " twocol"}>
            {panelComp("header")}
          <this.uikit.Block className="row">
            {panelComp("left")}
            {panelComp("right")}
          </this.uikit.Block>
          {panelComp("footer")}
        </this.uikit.Block>
        )
      }
      break;
      case "3col": {
        layout=( <this.uikit.Block className={className + " threecol"}>
            {panelComp("header")}
          <this.uikit.Block className="row">
            {panelComp("left")}
            {panelComp("right")}
          </this.uikit.Block>
          {panelComp("footer")}
        </this.uikit.Block>
        )
      }
      break;
      default: {
        layout=( <this.uikit.Block className={className}>
            {panelComp("items")}
            </this.uikit.Block>
        )
      }
    }
    this.getView = function (props, ctx, state) {
      return layout
    }
  }

  processBlock = (desc, props) => {
    console.log("processing block", desc)
    var display = this.getDisplayFunc(desc, props)
    console.log("processing block", desc, display, props)
    if(display) {
      this.getView = function(props, ctx, state) {
        let retval = display(props.data, desc, ctx.uikit)
        return retval
      }
    } else {
      this.getView = function(props, ctx, state) {
        return <ctx.uikit.Block></ctx.uikit.Block>
      }
    }
  }

  processForm = (desc, props) => {
    console.log("processing form", desc)
    let formdesc = desc
    if( desc.id) {
      formdesc = _reg("Forms", desc.id)
    }

    if(!formdesc) {
      return
    }

    if(!this.form) {
      this.form = this.getComponent("reactforms", "Form", module.req)
    }

    if(this.form) {
      this.getView = function(props, ctx, state) {
        return <this.form description={formdesc} id={desc.id}></this.form>
      }
    } else {
      this.getView = function(props, ctx, state) {
        return <ctx.uikit.Block></ctx.uikit.Block>
      }
    }
  }

  getDisplayFunc(item, props) {
    if(!item) {
      return null
    }
    if (typeof(item) == "string") {
      return _reg('Blocks', item)
    } else {
      let display = _reg('Blocks', item.block)
      if(!display) {
        display = _reg('Blocks', item.defaultBlock)
      }
      return display
    }
  }

  processView = (desc, props) => {
    let viewid = desc.viewid
    let viewdesc = desc
    if(viewid) {
      viewdesc = _reg('Views', viewid)
    }
    let description = Object.assign({}, viewdesc, desc)
    let viewHeader = description.header? <Panel description={description.header}/> :null

    if(!this.view) {
      this.view = this.getComponent("laatooviews", "View", module.req)
    }
    this.getView = function(props, context, state) {
      return <this.view params={props.params} header={viewHeader} id={viewid}>
        <Panel description={description.item} />
      </this.view>
    }
  }

  processEntity = (desc, props) => {
    let displayMode = desc.entityDisplay? desc.entityDisplay :"default"
    let id = "", name = ""
    if(props.params && props.params.id) {
      id = props.params.id
    } else {
      id = desc.entityId
    }
    name = desc.entityName
    if(!this.entity) {
      this.entity = this.getComponent("laatooviews", "Entity", module.req)
    }
    console.log("processing entity", props, " entity id ", id, "data",props.data)
    let entityDisplay={type:"block", block: desc.entityName+"_" + displayMode, defaultBlock: desc.entityName+"_default"}
    this.getView = function(props, ctx, state) {
      return <this.entity id={id} name={name} entityDescription={desc} data={props.data} index={props.index} uikit={ctx.uikit}>
        <Panel description={entityDisplay} />
      </this.entity>
    }
  }

  static setModule(mod) {
    module = mod;
  }
  getComponent = (mod, comp, req) => {
    let key = mod + comp
    let retval = module[key]
    if(!retval) {
      let moduleObj = req(mod);
      if(moduleObj && comp) {
        retval = moduleObj[comp]
        module[key] = retval
      }
    }
    return retval
  }

  render() {
    return this.getView? this.getView(this.props, this.context, this.state): <this.context.uikit.Block/>
  }
}

Panel.contextTypes = {
  uikit: PropTypes.object
};
export default Panel

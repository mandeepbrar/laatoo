import React from 'react'
const PropTypes = require('prop-types');

var module;

class Block extends React.Component {
  constructor(props, ctx) {
    super(props)
    let desc = props.blockDescription
    if(!desc && props.blockId) {
      desc = Application.Registry.Blocks[props.blockId]
    }
    if(desc && (typeof(desc) == 'string')) {
      this.processXMLComponent(desc, props)
    } else if(desc){
      switch(desc.type) {
        case "view":
          this.processView(desc,  props)
          break;
        case "entity":
          this.processEntity(desc, props)
          break;
        case "xmlcomponent":
          this.processXMLComponent(desc, props)
          break;
        case "component":
          this.getBlock = function(props, context, state) {
            return desc.component
          }
          break;
        default:
          let comp = this.getComponent(desc.module, desc.component, module.req)
          this.getBlock = function(props, context, state) {
            return React.createElement(comp)
          }
          break;
      }
    }
  }

  processXMLComponent = (desc, props) => {
    let display = this.getDisplayFunc(desc, props)
    if(display) {
      this.getBlock = function(props, ctx, state) {
        let retval = display(props.data, desc, ctx.uikit)
        return retval
      }
    } else {
      this.getBlock = function(props, ctx, state) {
        return <ctx.uikit.Block></ctx.uikit.Block>
      }
    }
  }

  getDisplayFunc(item, props) {
    if(!item || !Application.Registry.XMLDisplay) {
      return null
    }
    if (typeof(item) == "string") {
      return Application.Registry.XMLDisplay[item]
    } else {
      let display = Application.Registry.XMLDisplay[item.xmlName]
      if(!display) {
        display = Application.Registry.XMLDisplay[item.defaultXML]
      }
      return display
    }
  }

  processView = (desc, props) => {
    let viewid = desc.viewid
    let viewdesc = desc
    if(viewid && Application.Registry.Views) {
      viewdesc = Application.Registry.Views[viewid]
    }
    let viewItem = desc.item? desc.item: viewdesc.item
    let viewHeader = viewdesc.header? <Block blockDescription={viewdesc.header}/> :null

    if(!this.view) {
      this.view = this.getComponent("laatooviews", "View", module.req)
    }
    this.getBlock = function(props, context, state) {
      return <this.view params={props.params} header={viewHeader} id={viewid}>
        <Block blockDescription={viewItem} />
      </this.view>
    }
  }

  processEntity = (desc, props) => {
    let displayMode = desc.entityDisplay? desc.entityDisplay :"default"
    let id = "", name = ""
    if(props.params && props.params.id) {
      id = props.params.id
      name = props.params.name
    } else {
      id = desc.entityId
      name = desc.entityName
    }
    if(!this.entity) {
      this.entity = this.getComponent("laatooviews", "Entity", module.req)
    }

    let entityDisplay={type:"xmlcomponent", xmlName: desc.entityName+"_" + displayMode, defaultXML: desc.entityName+"_default"}
    this.getBlock = function(props, ctx, state) {
      return <this.entity id={id} name={name} entityDescription={desc} data={props.data} index={props.index} uikit={ctx.uikit}>
        <Block blockDescription={entityDisplay} ></Block>
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
    /*if(item.layout) {
      let blk = Application.Registry.Blocks[item.layout]
      if(blk) {
        return blk(item, comps, uikit)
      }
      return <uikit.Block/>
    }*/
    return this.getBlock? this.getBlock(this.props, this.context, this.state): <this.context.uikit.Block/>
  }
}

Block.contextTypes = {
  uikit: PropTypes.object
};
export default Block

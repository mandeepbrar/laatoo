import React from 'react'
const PropTypes = require('prop-types');

var module;

class Block extends React.Component {
  constructor(props, ctx) {
    super(props)
    let desc = props.blockDescription
    if(props.blockId) {
      desc = Application.Registry.Blocks[props.blockId]
    }
    console.log("creating block", desc)
    if(desc){
      switch(desc.type) {
        case "view":
          let viewname = desc.name
          let viewItem = desc.item
          let view = this.getComponent("laatooviews", "View", module.req)
          this.block = <view params={props.params} name={desc.name} id={desc.id}>
            <Block blockDescription={desc.item} />
          </view>
          break;
        case "entity":
          let displayMode = desc.entityDisplay? desc.entityDisplay :"default"
          let id = "", name = ""
          if(props.params && props.params.id) {
            id = props.params.id
            name = props.params.name
          } else {
            id = desc.entityId
            name = desc.entityName
          }
          let entity = this.getComponent("laatooviews", "Entity", module.req)
          var display;
          if(Application.Registry.EntityDisplay) {
            display = Application.Registry.EntityDisplay[desc.entityName+"_"+displayMode]
            if(!display) {
              display = Application.Registry.EntityDisplay[desc.entityName+"_default"]
            }
          }
          if(!display) {
            display=function(data, desc, uikit, time) {
              console.log("rendering ", data, time)
              if(data) {
                return <uikit.Block>{JSON.stringify(data)}</uikit.Block>
              }
              return <uikit.Block></uikit.Block>
            }
          }
          console.log("entity comp", entity, display)
          this.block = React.createElement(entity, {id: id, name: name, entityDescription:desc, display: display, uikit:ctx.uikit})
          break;
        case "component":
          this.block = desc.component
          break;
        default:
          let comp = this.getComponent(desc.module, desc.component, module.req)
          this.block = React.createElement(comp)
          break;
      }
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
    return this.block? this.block: <this.context.uikit.Block/>
  }
}

Block.contextTypes = {
  uikit: PropTypes.object
};
export default Block

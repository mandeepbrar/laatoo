import React from 'react'
const PropTypes = require('prop-types');

var module;

class Block extends React.Component {
  constructor(props) {
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
          this.block = <module.view name={desc.name} id={desc.id}>
            <Block blockDescription={desc.item} />
          </module.view>
          break;
        case "entity":
          let displayMode = desc.entityDisplay
          let display = Application.Registry.EntityDisplay[desc.entityName+"_"+displayMode]
          if(!display) {
            display = Application.Registry.EntityDisplay[desc.entityName+"_default"]
          }
          this.block = display(desc, uikit)
          break;
        case "component":
          this.block = desc.component
          break;
        default:
          this.block = this.getComponent(desc, module.req)
          break;
      }
    }
  }
  static setModule(mod) {
    module = mod;
  }
  getComponent = (compdesc, req) => {
    if(compdesc) {
      let mod = compdesc.module
      if(mod) {
        let moduleObj = req(mod);
        if(moduleObj && compdesc.component) {
          return React.createElement(moduleObj[compdesc.component])
        }
      }
    }
    return null
  }

  render() {
    /*if(item.layout) {
      let blk = Application.Registry.Blocks[item.layout]
      if(blk) {
        return blk(item, comps, uikit)
      }
      return <uikit.Block/>
    }*/
    return this.block? this.block: <uikit.Block/>
  }
}

Block.contextTypes = {
  uikit: PropTypes.object
};
export default Block

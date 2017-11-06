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
    console.log("creating panel", desc, props, ctx, this.context)
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
          this.processView(desc,  props, ctx, className)
          break;
        case "entity":
          className = className + " entity "
          this.processEntity(desc, props,  ctx, className)
          break;
        case "form":
          className = className + " form "
          this.processForm(desc, props, ctx,  className)
          break;
        case "html":
          className = className + " html "
          this.processHtml(desc, props, ctx,  className)
          break;
        case "block":
          className = className + " block "
          this.processBlock(desc, props,  ctx, className)
          break;
        case "layout":
          className = className + " panel "
          this.processLayout(desc, props,  ctx, className)
          break;
        case "component":
          if(desc.component) {
            this.getView = function(props, context, state) {
              return desc.component
            }
          } else {
            var comp = this.getComponent(desc.module, desc.componentName, module.req)
            console.log("rendering component", desc, comp)
            let cl = { className: className}
            var compProps = desc.props? Object.assign({}, desc.props, cl): cl
            this.getView = function(compProps, comp) {
              return function(props, context, state) {
                console.log("rendering comp", comp, compProps)
                return React.createElement(comp, compProps)
              }
            }(compProps, comp)
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

  processLayout = (desc, props,  ctx, className) => {
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

  processBlock = (desc, props, ctx, className) => {
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

  createMarkup = (text) => { return {__html: text}; };
  processHtml = (desc, props, ctx, className) => {
    if(desc.html) {
      this.getView = function(props, ctx, state) {
        console.log("rendering html", desc.html)
        return <div className={className} dangerouslySetInnerHTML={this.createMarkup(desc.html)} />
      }
    } else {
      this.getView = function(props, ctx, state) {
        return <ctx.uikit.Block></ctx.uikit.Block>
      }
    }
  }

  processForm = (desc, props,  ctx, className) => {
    console.log("processing form", desc)
    let formdesc = desc
    if( desc.id) {
      formdesc = _reg("Forms", desc.id)
    }

    if(!formdesc) {
      return
    }

    var cfg = formdesc.config
    if(!this.form) {
      this.form = this.getComponent("reactforms", "Form", module.req)
    }

    if(this.form) {
      this.getView = function(props, ctx, state) {
        let formCfg = Object.assign({}, cfg, ctx.routeParams)
        console.log("form cfg", formCfg, cfg)
        return <this.form form={desc.id} config={formCfg} description={formdesc} id={desc.id}></this.form>
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

  processView = (desc, props, ctx, className) => {
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

  processEntity = (desc, props, ctx, className) => {
    let displayMode = desc.entityDisplay? desc.entityDisplay :"default"
    console.log("view entity description", desc, displayMode)
    let id = "", name = ""
    if(ctx.routeParams && ctx.routeParams.entityId) {
      id = ctx.routeParams.entityId
    } else {
      id = desc.entityId
    }
    name = desc.entityName
    if(!this.entity) {
      this.entity = this.getComponent("laatooviews", "Entity", module.req)
    }
    console.log("processing entity", props, " entity id ", id, "data", props.data)
    let entityDisplay={type:"block", block: desc.entityName+"_" + displayMode, defaultBlock: desc.entityName+"_default"}
    console.log("entity display", entityDisplay)
    this.getView = function(props, ctx, state) {
      console.log("entity display in get view", entityDisplay)
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
  uikit: PropTypes.object,
  routeParams: PropTypes.object
};
export default Panel

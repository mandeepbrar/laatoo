import React from 'react'
import PropTypes from 'prop-types';
import {Action} from 'reactwebcommon';

var module;

class Panel extends React.Component {
  constructor(props, ctx) {
    super(props)
    this.uikit = ctx.uikit
    let desc = props.description

    if(!desc && props.id) {
      desc = _reg('Panels', props.id)
    }
    if(!desc && props.id) {
      let type = props.type? props.type: "block";
      desc = Object.assign({type: type}, props)
    }

    this.title = props.title? props.title: (desc.title? desc.title:null)
    this.closePanel = props.closePanel? props.closePanel: null

    //console.log("print id ", desc.id)
    console.log("creating panel", desc, props, ctx, this.context)
    let className = props.className? props.className + " panel " : " panel "
    if(desc.id) {
      className = className + " " +desc.id
    }
    if(desc.name) {
      className = className + " " +desc.name
    }
    if(desc.className) {
      className = className + " " +desc.className
    }
    this.overlay = props.overlay? props.overlay: null
    if(desc && (typeof(desc) == 'string')) {
      this.processBlock(desc, props)
    } else if(desc){
      switch(desc.type) {
        case "view":
          className = className + " view "
          this.processView(desc,  props, ctx)
          break;
        case "entity":
          className = className + " entity "
          this.processEntity(desc, props,  ctx)
          break;
        case "form":
          className = className + " form "
          this.processForm(desc, props, ctx)
          break;
        case "html":
          className = className + " html "
          this.processHtml(desc, props, ctx)
          break;
        case "block":
          className = className + " panelblock "
          this.processBlock(desc, props,  ctx)
          break;
        case "layout":
          className = className + " layout "
          this.processLayout(desc, props,  ctx)
          break;
        case "component":
          if(desc.component) {
            this.getView = function(props, context, state) {
              return desc.component
            }
          } else {
            var comp = this.getComponent(desc.module, desc.componentName, module.req)
            this.getView = function(comp) {
              return function(props, context, state, className) {
                let cl = { description: desc, className: className};
                var compProps = desc.props? Object.assign({}, desc.props, cl): cl
                return React.createElement(comp, compProps)
              }
            }(comp)
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
    console.log("class name", this.className, className, desc)
    this.className = className
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

  cfgPanel = (title, overlay) => {
    if(!this.title && title) {
      this.title = title
    }
    if(!this.overlay && overlay) {
      this.overlay = overlay
    }
  }

  processLayout = (desc, props,  ctx) => {
    let panelDesc = desc
    if(desc.id) {
      panelDesc = _reg("Panels", desc.id)
    }
    if(!panelDesc.layout) {
      return
    }
    this.cfgPanel(panelDesc.title, panelDesc.overlay)
    var block = this
    var layout = null
    var panelComp = (id)=> {
      return panelDesc[id]?
      <block.uikit.Block className={id}>
        {block.getPanelItems(panelDesc[id])}
      </block.uikit.Block>
      :null
    }

    this.getView = function (props, ctx, state, className) {
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
      return layout
    }
  }

  processBlock = (desc, props, ctx) => {
    var display = this.getDisplayFunc(desc, props)
    this.cfgPanel(desc.title, desc.overlay)
    console.log("processing block", desc, display, props)
    if(display) {
      this.getView = function(props, ctx, state, className) {
        console.log("calling block func", props, ctx)
        let retval = display({data: props.data, parent: props.parent, className: className, routeParams: ctx.routeParams, storage: Storage}, desc, ctx.uikit)
        return retval
      }
    } else {
      this.getView = function(props, ctx, state, className) {
        return <ctx.uikit.Block></ctx.uikit.Block>
      }
    }
  }

  createMarkup = (text) => { return {__html: text}; };
  processHtml = (desc, props, ctx) => {
    this.cfgPanel(desc.title, desc.overlay)
    if(desc.html) {
      this.getView = function(props, ctx, state, className) {
        console.log("rendering html", desc.html)
        return <div className={className} dangerouslySetInnerHTML={this.createMarkup(desc.html)} />
      }
    } else {
      this.getView = function(props, ctx, state, className) {
        return <ctx.uikit.Block></ctx.uikit.Block>
      }
    }
  }

  processForm = (desc, props,  ctx) => {
    console.log("processing form", desc)
    let formdesc = desc
    if( desc.id) {
      formdesc = _reg("Forms", desc.id)
    }

    if(!formdesc) {
      return
    }
    console.log("processing form", desc)
    this.cfgPanel(formdesc.info.title, formdesc.info.overlay)

    var cfg = formdesc.info
    if(!this.form) {
      this.form = this.getComponent("reactforms", "Form", module.req)
    }

    if(this.form) {
      this.getView = function(props, ctx, state, className) {
        let formCfg = Object.assign({}, cfg, ctx.routeParams)
        console.log("form cfg", formCfg, cfg)
        return <this.form form={desc.id} parent={props.parent} formContext={{data: props.data, routeParams: ctx.routeParams, storage: Storage}} config={formCfg} inline={props.inline}
          onChange={props.onChange} trackChanges={props.trackChanges} formData={props.formData} onSubmit={props.onSubmit} subform={props.subform} title={props.title}
          actions={props.actions} description={formdesc} className={className} id={desc.id}/>
      }
    } else {
      this.getView = function(props, ctx, state, className) {
        return <ctx.uikit.Block></ctx.uikit.Block>
      }
    }
  }

  getDisplayFunc(item, props) {
    console.log("getting block", item)
    if(!item) {
      return null
    }
    if (typeof(item) == "string") {
      return _reg('Blocks', item)
    } else {
      let bid = item.block? item.block: item.id;
      let display = _reg('Blocks', bid)
      if(!display) {
        display = _reg('Blocks', item.defaultBlock)
      }
      return display
    }
  }

  processView = (desc, props, ctx) => {
    let viewid = desc.viewid? desc.viewid: desc.id
    var viewdesc = desc
    if(viewid) {
      viewdesc = _reg('Views', viewid)
    }
    this.cfgPanel(viewdesc.title, viewdesc.overlay)
    let description = Object.assign({}, viewdesc, desc)
    let viewHeader = description.header? <Panel description={description.header}/> :null

    if(!this.view) {
      this.view = this.getComponent("laatooviews", "View", module.req)
    }

    this.getView = function(props, context, state, className) {
      return <this.view params={props.params} description={viewdesc} getItem={props.getItem} className={className} header={viewHeader} 
        id={viewid}>
        <Panel parent={props.parent} description={description.item} />
      </this.view>
    }
  }

  processEntity = (desc, props, ctx) => {
    let displayMode = desc.entityDisplay? desc.entityDisplay :"default"
    console.log("view entity description", desc, displayMode, props)
    let id = "", name = ""
    if(ctx.routeParams && ctx.routeParams.entityId) {
      id = ctx.routeParams.entityId
    } else {
      id = desc.entityId
    }
    name = desc.entityName
    this.cfgPanel(desc.title, desc.overlay)
    if(!this.entity) {
      this.entity = this.getComponent("laatooviews", "Entity", module.req)
    }
    var itemClass = ""
    if(props.index) {
      itemClass = props.index%2 ? "oddindex": "evenindex"
    }
    var entityDisplay={type:"block", block: desc.entityName+"_" + displayMode, defaultBlock: desc.entityName+"_default"}
    this.getView = function(props, ctx, state, className) {
      return <this.entity id={id} name={name} entityDescription={desc} data={props.data} index={props.index} uikit={ctx.uikit}>
        <Panel description={entityDisplay} parent={props.parent} className={itemClass} />
      </this.entity>
    }
  }

  static setModule(mod) {
    module = mod;
  }

  getChildContext() {
    console.log("getting child contextoverlaying component", this.overlay, this.props, this.context)
    if(this.overlay) {
      return {overlayComponent: this.overlayComponent, closeOverlay: this.closeOverlay}
    } else {
      return this.context && this.context.overlayComponent? {overlayComponent: this.context.overlayComponent, closeOverlay: this.closeOverlay}: null
    }
  }

  overlayComponent = (comp) => {
    console.log("overlaying component")
    this.setState(Object.assign({}, {overlayComponent: comp}))
  }

  closeOverlay = () => {
    this.setState({})
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
    console.log("rendering panel", this.props);
    let showOverlay = this.overlay && this.state && this.state.overlayComponent // ? "block": "none"
    let comp = this.getView? this.getView(this.props, this.context, this.state, (this.title? "": this.className)): <this.context.uikit.Block/>
    if(this.overlay || this.title) {
      return <this.uikit.Block className={this.className} title={this.title} closeBlock={this.closePanel}>
        <this.uikit.Block style={{display:( showOverlay?"none":"block")}}>
        {comp}
        </this.uikit.Block>
        {showOverlay?this.state.overlayComponent:null}
      </this.uikit.Block>
    } else {
      return comp
    }
  }
}

Panel.childContextTypes = {
  overlayComponent: PropTypes.func,
  closeOverlay: PropTypes.func
};

Panel.contextTypes = {
  overlayComponent: PropTypes.func,
  closeOverlay: PropTypes.func,
  uikit: PropTypes.object,
  routeParams: PropTypes.object
};
export default Panel

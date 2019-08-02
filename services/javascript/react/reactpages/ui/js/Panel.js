import React from 'react'
import PropTypes from 'prop-types';
import {Action} from 'reactwebcommon';

var module;

class Panel extends React.Component {
  constructor(props, ctx) {
    super(props)
    this.uikit = ctx.uikit
    let desc = props.description

    /*if(!desc && props.id) {
      desc = _reg('Panels', props.id)
    }*/
    let id = props.id? props.id : ((desc && desc.id)? desc.id: null);
    let type = props.type? props.type: ((desc && desc.type)? desc.type: "layout");
    if(id) {
      switch(type) {
        case "view":
          desc = _reg('Views', id)
          break;
        case "form":
          desc = _reg('Forms', id)
          break;
        case "block":
          desc = _reg('Blocks', id)
          break;
        default:
          desc = _reg('Panels', id)
      }
      console.log("desc before assig", desc, props)
      desc = Object.assign({type: type, id: id}, desc, props)
    }

    this.title = props.title? props.title: (desc && desc.title? desc.title:null)
    this.closePanel = props.closePanel? props.closePanel: null

    //console.log("print id ", desc.id)
    console.log("creating panel test*$$$$$$$$$$$$$$$$", desc, props, ctx, this.context)
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
    if(!desc || !desc.layout) {
      return
    }
    this.cfgPanel(desc.title, desc.overlay)
    var block = this
    var layout = null
    var panelComp = (id)=> {
      return desc[id]?
      <block.uikit.Block className={id}>
        {block.getPanelItems(desc[id])}
      </block.uikit.Block>
      :null
    }

    this.getView = function (props, ctx, state, className) {
      switch(desc.layout) {
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
    let panel=this
    if(display) {
      this.getView = function(props, ctx, state, className) {
        console.log("calling block func", props, ctx)
        let retval = display({data: props.data, parent: props.parent, panel: panel, className: className, routeParams: ctx.routeParams, storage: Storage}, desc, ctx.uikit)
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
    if(!desc || !desc.info) {
      return
    }
    this.cfgPanel(desc.info.title, desc.info.overlay)

    var cfg = desc.info
    if(!this.form) {
      this.form = this.getComponent("reactforms", "Form", module.req)
    }

    if(this.form) {
      this.getView = function(props, ctx, state, className) {
        let formCfg = Object.assign({}, cfg, ctx.routeParams)
        console.log("form cfg", formCfg, cfg, props)
        return <this.form form={desc.id} parentFormRef={props.parentFormRef} formContext={{data: props.data, routeParams: ctx.routeParams, storage: Storage}} config={formCfg} inline={props.inline}
          onChange={props.onChange} trackChanges={props.trackChanges} formData={props.formData} onSubmit={props.onSubmit} subform={props.subform} title={props.title}
          autoSubmitOnChange={props.autoSubmitOnChange} actions={props.actions} description={desc} className={className} id={desc.id}/>
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
      let bid =  item.id;
      let display = _reg('Blocks', bid)
      if(!display) {
        display = _reg('Blocks', item.defaultBlock)
      }
      return display
    }
  }

  processView = (desc, props, ctx) => {
    console.log("processing my view", desc, props, module)
    this.cfgPanel(desc.title, desc.overlay)
    let viewHeader = desc.header? <Panel description={desc.header}/> :null

    if(!this.view) {
      this.view = this.getComponent("laatooviews", "View", module.req)
    }
    console.log("processing view", this.view)

    this.getView = function(props, context, state, className) {
      console.log("rendering view", this.view, props, desc, className)
      return <this.view params={props.params} description={desc} getItem={props.getItem} editable={props.editable} className={className} header={viewHeader}
        viewRef={props.viewRef} id={desc.id}>
        <Panel parent={props.parent} description={desc.item} />
      </this.view>
    }
  }

  processEntity = (desc, props, ctx) => {
    if(!this.entity) {
      this.entity = this.getComponent("laatooviews", "Entity", module.req)
    }
    this.getView = function(props, ctx, state, className) {
      var desc = props.description
      let displayMode = desc.entityDisplay? desc.entityDisplay :"default"
      console.log("view entity description", desc, displayMode, props)
      this.cfgPanel(desc.title, desc.overlay)
      var entityDisplay={type:"block", id: desc.entityName+"_" + displayMode, defaultBlock: desc.entityName+"_default"}
      let id = "", name = ""
      if(ctx.routeParams && ctx.routeParams.entityId) {
        id = ctx.routeParams.entityId
      } else {
        id = desc.entityId
      }
      name = desc.entityName
      var entityData = props.data? props.data: desc.data
      var entityIndex = props.index
      var itemClass = ""
      if(props.index) {
        itemClass = props.index%2 ? "oddindex": "evenindex"
      }
      console.log("my entity data111", entityData, entityIndex, desc, props)
      return <this.entity id={id} name={name} entityDescription={desc} data={entityData} index={entityIndex} uikit={ctx.uikit}>
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
    if(this.overlay) {
      this.setState(Object.assign({}, {overlayComponent: comp}))
    } else {
      if(this.context && this.context.overlayComponent) {
        this.context.overlayComponent(comp)
      }
    }
  }

  closeOverlay = () => {
    if(this.overlay) {
      this.setState({})
    } else {
      if(this.context && this.context.closeOverlay) {
        this.context.closeOverlay()
      }
    }
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
    console.log("rendering panel", this.props, this.getView, this.className);
    let showOverlay = this.overlay && this.state && this.state.overlayComponent // ? "block": "none"
    let comp = this.getView? this.getView(this.props, this.context, this.state, (this.title? "": this.className)): <this.context.uikit.Block/>
    if(this.overlay || this.title || this.closePanel) {
      return <this.uikit.Block className="overlaywrapper" title={this.title} closeBlock={this.closePanel}>
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

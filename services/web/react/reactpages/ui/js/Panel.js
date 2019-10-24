import React from 'react'
import PropTypes from 'prop-types';
import {Action} from 'reactwebcommon';

var module;

class Panel extends React.Component {
  constructor(props, ctx) {
    super(props)
    let desc = props.description

    /*if(!desc && props.id) {
      desc = _reg('Panels', props.id)
    }*/
    let id = props.id? props.id : ((desc && desc.id)? desc.id: null);
    let type = props.type? props.type: ((desc && desc.type)? desc.type: "layout");
    if(id) {
      let panelInfo = desc? desc.info: null;
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
        case "component":
          break;
        default:
          desc = _reg('Panels', id)
      }
      console.log("desc before assig", desc, props, id, type, Application)
      desc = Object.assign({type: type, id: id}, desc)
      console.log("desc after assign", desc, desc.info, panelInfo)
      desc.info = Object.assign({}, desc.info, panelInfo)
    }

    this.title = props.title? props.title: (desc && desc.title? desc.title:null)
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
            var comp = _res(desc.module, desc.componentName)
            console.log("component from module", comp, desc)
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
        console.log("rendering panel array element", item)
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
      console.log("looking for panel component in layout", desc, id)
      return desc[id]?
      <_uikit.Block className={id}>
        {block.getPanelItems(desc[id])}
      </_uikit.Block>
      :null
    }

    this.getView = function (props, ctx, state, className) {
      console.log("getting layout", desc, props)
      switch(desc.layout) {
        case "2col": {
          layout=( <_uikit.Block className={className + " twocol"}>
              {panelComp("header")}
            <_uikit.Block className="row">
              {panelComp("left")}
              {panelComp("right")}
            </_uikit.Block>
            {panelComp("footer")}
          </_uikit.Block>
          )
        }
        break;
        case "3col": {
          layout=( <_uikit.Block className={className + " threecol"}>
              {panelComp("header")}
            <_uikit.Block className="row">
              {panelComp("left")}
              {panelComp("right")}
            </_uikit.Block>
            {panelComp("footer")}
          </_uikit.Block>
          )
        }
        break;
        default: {
          layout=( <_uikit.Block className={className}>
              {panelComp("items")}
              </_uikit.Block>
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
        console.log("calling block func", props, ctx, className)
        let retval = display({data: props.data, parent: props.parent, panel: panel, className: className, routeParams: ctx.routeParams, storage: Storage}, desc)
        console.log("returning block retval", retval)
        return retval
      }
    } else {
      this.getView = function(props, ctx, state, className) {
        console.log("rendering empty block func", props, ctx, className)
        return <_uikit.Block></_uikit.Block>
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
        return <_uikit.Block></_uikit.Block>
      }
    }
  }

  processForm = (desc, props,  ctx) => {
    console.log("processing form=", desc, props)
    if(!desc || !desc.info) {
      return
    }
    this.cfgPanel(desc.info.title, desc.info.overlay)

    let formName=desc.formName? desc.formName : desc.id
    console.log("processing form+++", desc, formName, props, desc.formName, desc.id)

    var cfg = desc.info
    if(!this.form) {
      console.log("getting form",  module)
      this.form = _res("reactforms", "Form")
      console.log("got form",  module)
    }

    if(this.form) {
      this.getView = function(props, ctx, state, className) {
        let formCfg = cfg
        if(!props.subform) {
          formCfg = Object.assign({}, cfg, ctx.routeParams)
        }
        console.log("form cfg", formCfg, cfg, props, formName)
        return <this.form form={formName} parentFormRef={props.parentFormRef} formContext={{data: props.data, routeParams: ctx.routeParams, storage: Storage}}
           info={formCfg} inline={props.inline} onChange={props.onChange} trackChanges={props.trackChanges} formData={props.formData} subform={props.subform}
           onFormSubmit={props.onFormSubmit} title={props.title} actions={props.actions} description={desc} className={className} id={desc.id}/>
      }
    } else {
      this.getView = function(props, ctx, state, className) {
        return <_uikit.Block></_uikit.Block>
      }
    }
  }

  getDisplayFunc(item, props) {
    console.log("getting display func", item);
    if(!item) {
      console.log("returning null display func", item)
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
      console.log("returning display func", display, Application);
      return display
    }
  }

  processView = (desc, props, ctx) => {
    console.log("processing my view", desc, props, module)
    this.cfgPanel(desc.title, desc.overlay)
    let viewHeader = desc.header? <Panel id={desc.header} type="block"/> :null

    if(!this.view) {
      this.view = _res("laatooviews", "View")
    }
    console.log("processing view", this.view)

    this.getView = function(props, context, state, className) {
      console.log("rendering view", this.view, props, desc, className)
      return <this.view params={props.params} description={desc} getItem={props.getItem} editable={props.editable} className={className} header={viewHeader}
        viewRef={props.viewRef} postArgs={props.postArgs} urlParams={props.urlParams} id={desc.id}>
        <Panel parent={props.parent} description={desc.item} />
      </this.view>
    }
  }

  processEntity = (desc, props, ctx) => {
    if(!this.entity) {
      this.entity = _res("laatooviews", "Entity")
    }
    this.getView = function(props, ctx, state, className) {
      var desc = props.description
      let displayMode = desc.entityDisplay? desc.entityDisplay :"default"
      console.log("view entity description", desc, displayMode, props, ctx)
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
      console.log("my entity data", entityData, entityIndex, desc, props, id, ctx)
      return <this.entity id={id} name={name} entityDescription={desc} data={entityData} index={entityIndex}>
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

  /*getComponent = (mod, comp, req) => {
    let key = mod + comp
    let retval = module[key]
    if(!retval) {
      let moduleObj = req(mod);
      if(moduleObj && comp) {
        retval = moduleObj[comp]
        module[key] = retval
      }
    }
    return retval;
  }*/

  render() {
    let showOverlay = this.overlay && this.state && this.state.overlayComponent // ? "block": "none"
    let comp = this.getView? this.getView(this.props, this.context, this.state, (this.title? "": this.className)): <_uikit.Block/>
    console.log("Rendering panel***************", this.getView, this.props, comp, this.overlay, this.title, this.closePanel, this.className);
    if(this.overlay || this.title || this.closePanel) {
      return <_uikit.Block className="overlaywrapper" title={this.title} closeBlock={this.closePanel}>
        <_uikit.Block style={{display:( showOverlay?"none":"block")}}>
        {comp}
        </_uikit.Block>
        {showOverlay?this.state.overlayComponent:null}
      </_uikit.Block>
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
  routeParams: PropTypes.object
};
export default Panel

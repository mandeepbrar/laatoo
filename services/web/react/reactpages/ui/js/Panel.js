import React, {useState} from 'react'
import PropTypes from 'prop-types';
import {Action, Menu} from 'reactwebcommon';

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
          //desc = _reg('Blocks', id)          
          break;
        case "menu":
          desc = _reg('Menus', id)
          break;
        case "component":
          break;
        default:
          desc = _reg('Panels', id)
      }
      console.log("desc before assigne", props.description, props, id, type, Application)
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
        case "menu":
          className = className + " menu "
          this.processMenu(desc, props,  ctx)
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
                var compProps = desc.props? Object.assign({}, desc.props, cl): Object.assign({}, desc, cl)
                return React.createElement(comp, compProps)
              }
            }(comp)
          }
          break;
        default:
          let processor = _reg("PanelTypes", desc.type)
          //let comp = processor.getComponent(desc)
          this.getView = function(props, context, state) {
            console.log("unknown comp type", desc.type, desc)
            if(processor != null) {              
              return processor.getComponent(desc, props, context, state)
            } else {
              return <_uikit.Block/>
            }
            
          }
          break;
      }
    }
    console.log("class name", this.className, className, desc)
    this.state = {expanded: false, overlayComponent:null}
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
      :<_uikit.Block/>
    }

    this.getView = function (props, ctx, state, className) {
      console.log("getting layout", desc, props)
      switch(desc.layout) {
        case "2col": {
          layout=( <_uikit.Block className={className + " fd fdcol w100 twocol"}>
              {panelComp("headerrow")}
            <_uikit.Block className=" fd fdgrow fdrow ">
              {panelComp("leftcol")}
              {panelComp("rightcol")}
            </_uikit.Block>
            {panelComp("footerrow")}
          </_uikit.Block>
          )
        }
        break;
        case "3col": {
          layout=( <_uikit.Block className={className + " fd fdcol w100 threecol"}>
              {panelComp("headerrow")}
            <_uikit.Block className=" fd fdgrow fdrow ">
              {panelComp("leftcol")}
              {panelComp("rightcol")}
            </_uikit.Block>
            {panelComp("footerrow")}
          </_uikit.Block>
          )
        }
        break;
        default: {
          layout=( <_uikit.Block className={className + " fd fdcol w100 "}>
              {panelComp("items")}
              </_uikit.Block>
          )
        }
      }
      return layout
    }
  }

  processMenu = (desc, props, ctx) => {
    this.getView = function(props, ctx, state, className) {
      console.log("rendering menu", desc, props, ctx, className)
      return <Menu id={desc.id} vertical={desc.vertical} className={className}/>
    }
  }

  processBlock = (desc, props, ctx) => {
    var display = this.getDisplayFunc(desc, props)
    var expansionDisplay = null
    var expansionDesc = null
    if(desc.expansionBlock){
      expansionDesc = {id: desc.expansionBlock, type: "block"}
      expansionDisplay = this.getDisplayFunc(expansionDesc)
    }
    this.cfgPanel(desc.title, desc.overlay)
    console.log("processing block", desc, display, props)
    //this.state = 
    //const [expanded, expand] = useState(false);
    let panel=this
    let expand=function(expand) {
      console.log("expanding", panel.state)
      panel.setState(Object.assign({}, panel.state, {expanded: !panel.state.expanded}))
    }
    if(display) {
      this.getView = function(props, ctx, state, className) {
        console.log("use stat", useState, state, expand);
        console.log("calling block func", props, ctx, state, className, state.expanded)
        let expandedComp = props.expanded
        if ((expandedComp != null) && (expandedComp == false)) {
          return null
        }
        let blockCtx = {data: props.data, expanded: state.expanded, expand: expand, props: props, panel: panel, className: className, routeParams: ctx.routeParams, storage: Storage, state: state}
        console.log("block blockCtx", blockCtx, "desc", desc, "expansionDesc", expansionDesc)
        let retval = display(blockCtx, desc)
        if(state.expanded && expansionDisplay){
          return [retval, expansionDisplay(blockCtx, expansionDesc)]
        }
        if(retval && props.contentOnly) {
          return retval.children
        }
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

    var cfg = Object.assign({}, desc.info, props.info)
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
    let postArgs = _tn(props.postArgs, desc.postArgs)
    console.log("processing view", this.view)

    this.getView = function(props, context, state, className) {
      console.log("rendering view", this.view, props, desc, className, postArgs)      
      return <this.view params={props.params} description={desc} getItem={props.getItem} editable={props.editable} className={className} header={viewHeader}
        viewRef={props.viewRef} postArgs={postArgs} urlParams={props.urlParams} contentOnly={props.contentOnly} id={desc.id} instance={props.instance}>
        <Panel parent={props.parent} description={desc.item} viewdepth={props.viewdepth} />
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
      this.setState(Object.assign({}, this.state, {overlayComponent: comp}))
    } else {
      if(this.context && this.context.overlayComponent) {
        this.context.overlayComponent(comp)
      }
    }
  }

  closeOverlay = () => {
    if(this.overlay) {
      this.setState({}, this.state, {overlayComponent: null})
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
    if((!this.props.contentOnly) && (this.overlay || this.title || this.closePanel)) {
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

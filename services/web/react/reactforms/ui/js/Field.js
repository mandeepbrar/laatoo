import React from 'react';
import { Field as RFField } from 'redux-form'
import {Entity} from './entity';
import {List} from './list';
import PropTypes from 'prop-types';


class FieldWidget extends React.Component {
  constructor(props) {
    super(props)
    //injectTapEventPlugin()
    console.log("creating field widget", props)
    let cfg = null
    let field = props.field
    if(field) {
      cfg = this.processEntityField(field)
    } 
    cfg = _tn(cfg, {})
    this.processProps(props, cfg)    

    this.state = {formValue: props.formValue}
    this.getWidget(cfg)
    this.cfg = cfg
    console.log("constructing material kit field", this.cfg, props)
    /*if(field && field.widget) {
    }*/
  }

  componentWillReceiveProps(nextProps, nextState) {
    console.log("field component of react forms", nextProps)
    this.setState({formValue: nextProps.formValue})
  }    


  getWidget=(cfg)=> {
    console.log("get widget for cfg", cfg.name, cfg)
    if(cfg.widgetModule) {
      console.log("chosing custom component for ", cfg.name)
      this.widgetComp = _res(cfg.widgetModule, cfg.widgetName);
    } else if(cfg.list) {
      console.log("chosing listfor ", cfg.name)
      this.widgetComp = List
    } else if(cfg.entity) {
      console.log("chosing entityfor ", cfg.name)
      this.widgetComp = Entity
    } else {
      this.processTypes(cfg)
//      console.log("chosing uikit field for ", cfg.name)
    }
  }

  processTypes = (cfg) => {
    this.widgetComp = _uikit.Field
    if(!cfg.widgetName) {
      switch(cfg.type) {
        case "string":
          cfg.widgetName = "TextField"
        break;
        case "int":
        case "float":
          cfg.widgetName = "NumberField"
        break;
        case "storableref": 
          cfg.isRef = true
          break
        case "stringmap":
          this.widgetComp = List
          cfg.itemForm = "list_add_keyvalue"
          cfg.titleField = "mapkey"
          cfg.isMap = true
        break;
        case "bool":
          cfg.widgetName = "Checkbox"
        break;
        case "date":
          cfg.widgetName = "DatePicker"
        break;
      }
    }
  }

  processEntityField = (fld) => {
    let cfg = {}
    if(fld.widget) {
      let widgetProps = fld.widget.props
      cfg = Object.assign({}, widgetProps)
      cfg.widgetProps = widgetProps
      cfg.label = _tn(cfg.label, fld.name)
      cfg.widgetName = fld.widget.name
      cfg.widgetModule = fld.widget.module
/*      cfg.className = widgetProps.className
      cfg.selectItem = widgetProps.selectItem
      cfg.dataServiceParams = widgetProps.dataServiceParams
      cfg.loader = widgetProps.loader
      cfg.loadData = widgetProps.loadData
      cfg.dataService = widgetProps.dataService
      cfg.controlClassName = widgetProps.controlClassName
      cfg.items = widgetProps.items
      cfg.itemClass = widgetProps.itemClass
      */

      //text and value used by select
      cfg.textField = _tn(cfg.textField, "text")
      cfg.valueField = _tn(cfg.valueField, "value")
    }
    cfg.name = fld.name
    cfg.type = fld.type
    if(fld.list) {
      cfg.list = fld.list
    }
    if(fld.entity) {
      cfg.entity = fld.entity
    }
    return cfg
  }

  processProps= (props, cfg) => {
    let fldProps = ["name", "label", "items", "textField", "valueField", "widgetName", "widgetModule", "selectItem", "itemClass", 
      "dataServiceParams", "loader", "loadData", "dataService", "type", "list", "mode" ]
    fldProps.map((item)=>{
      cfg[item] = _tn(props[item], cfg[item])
    })
    cfg.titleField = _tn(cfg.titleField, "Name")
    cfg.className = _tn(props.className, "") + " " + cfg.name +" " + _tn(cfg.className, " ") 
    cfg.controlClassName= cfg.name + " " + _tn(cfg.controlClassName, "")
  }

  fieldChange = (onChange) => {
    let cfg = this.cfg
    return (data, name, evt) => {
      console.log("field change", onChange, data, name)
      if(onChange) {
        if(cfg.isRef) {
          let myRefObj = {}
          myRefObj[cfg.name] = {"Id": data, "Type": cfg.entity}
          data = myRefObj
          console.log("set ref value data", myRefObj, data)
        }
        if(cfg.isMap) {
          if(data) {
            let changedObj = {}
            data.forEach((item)=> {
              changedObj[item.mapkey] = item.mapvalue  
            })
            console.log("changed obj for map", changedObj, data)
            data = changedObj
          }
        }
        console.log(" field change ", data, name)
        onChange(data, name, evt)
      }
    }
  }


  component = (fieldProps) => {
    let {input, meta, className} = fieldProps
    let errorText = meta.touched && meta.error
    let cfg = this.cfg
    let value = null
    if(cfg.isRef) {
      console.log("changing value of storable ref", value, input )
      value = input.value.Id
    } else {
      value = input.value
    }

    if(value && cfg.isMap) {
      value = Object.keys(value).map((item)=>{
          return {mapkey: item, mapvalue: value[item]}
      })
    }

    let rfieldProps ={onChange: this.fieldChange(input.onChange), errorText: errorText, formValue: this.state.formValue, 
      onFocus: input.onFocus, onBlur: input.onBlur, value: value}            

    console.log("field component", this.state, rfieldProps, cfg)
    let newProps = Object.assign({}, cfg, rfieldProps)
    if(this.cfg.transformer) {
        let transformerMethod = _reg("Methods", cfg.transformer)
        newProps = transformerMethod(newProps, cfg, this.props)
    }
    console.log("creating widget for field:", this.cfg.name, " newProps:", newProps, " widget:", this.widgetComp)
    return React.createElement(this.widgetComp, newProps, null)
/*
    if(field) {
        let widget = this.fieldWidgets[fieldName]
        let fprops = this.fieldProps[fieldName]
        let fieldChange= this.fieldChange(input.onChange)
        let errorText = meta.touched && meta.error
        let cl = className + (fprops.className? fprops.className: "")
        
        let newProps = Object.assign({}, fprops, {onChange: fieldChange, errorText: errorText, formValue: this.state.formValue, className: cl,
            onFocus: input.onFocus, onBlur: input.onBlur, value: input.value})            
        if(field.transformer) {
            let transformerMethod = _reg("Methods", field.transformer)
            newProps = transformerMethod(newProps, this.props.formValue, field, this.fields, this.props, this.state,  this)
        }
        if(field.type == "storableref") {
            let isRef = true
            let ref = input.value[fieldName]
            if(ref) {
              newprops.value = input.value = ref.Id
            }
            console.log("ref props", newProps)
        }
        console.log("creating widget, state:", this.state, " props:", newProps, " widget:", widget)
        return React.createElement(widget, newProps, null)
    }
    return null*/
  }

  render() {
    return <RFField name={this.cfg.name} component={this.component}/>
  }
}

FieldWidget.propTypes = {
  classes: PropTypes.object.isRequired,
};

export {
  FieldWidget as Field
}

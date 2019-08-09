import React from 'react';
const PropTypes = require('prop-types');
//import {Field} from './Field';
import { Field } from 'redux-form'
import {FldList} from './FldList';

class FieldsPanel extends React.Component {
    constructor(props, context) {
        super(props)
        let desc = props.description
        this.fields = desc? desc.fields:{}
        this.state={formValue: props.formValue}       
        console.log("fields panel constructor", desc, props)
        this.configureFields(props)
    }
    componentWillReceiveProps(nextProps, nextState) {
        this.setState(Object.assign({}, this.state, {formValue: nextProps.formValue}))
    }    

    configureFields = (props) => {
        this.fieldWidgets = {}
        this.fieldProps = {}
        console.log("configure fields", props, this)
        for(let [fieldName, field] of Object.entries(this.fields)) {
            if(field.widget && field.widget.module) {
                let fldWidget = _res(field.widget.module, field.widget.name);
                this.fieldWidgets[fieldName] = fldWidget
            } else if(field.list) {
                this.fieldWidgets[fieldName] = FldList
            } else {
                this.fieldWidgets[fieldName] = _uikit.Forms.FieldWidget
            }

            let cl= props.inline?"inline formfield m10":"formfield m10"

            if(!field.label && !field.skipLabel) {
                field.label = fieldName
            }
            let fieldProperties = {name: fieldName, entity: field.entity, className: cl, field: field, formRef: props.formRef}
            if(field.widget) {
                fieldProperties = Object.assign(fieldProperties, field.widget.props)
            }
            
            this.fieldProps[fieldName] = fieldProperties
        }
    }

    layoutFields = (fldToDisp, flds, className, state) => {
        let fieldsArr = new Array()
        let fldpanel = this
        fldToDisp.forEach(function(k) {
            let field = flds[k]    
            console.log("layout field in fld panel ", k)
            fieldsArr.push( <Field key={field.name} name={field.name} className={className} component={fldpanel.component}/>)
            //fieldsArr.push(  <Field key={fd.name} name={fd.name} formValue={state.formValue} autoSubmitOnChange={props.autoSubmitOnChange} fields={flds}
              //   className={cl} formRef={props.formRef} />  )
        })
        return fieldsArr
    }

    

//    <Field key={this.props.name} name={this.props.name} time={this.state.time} component={this.component}/>


    fieldChange = (onChange) => {
        let comp=this
        return (data, name, evt)=> {
            console.log("fld change", data, name, evt, this.props, this.context, onChange, comp.isRef)
            if(onChange) {
                if(comp.isRef) {
                    let myRefObj = {}
                    myRefObj[comp.field.name] = {"Id": data, "Type": comp.field.entity}
                    data = myRefObj
                    console.log("set ref value data", myRefObj, data)
                }
                console.log("setting fld value", comp, data, name)
                onChange(data, name)
            }
        }
    }

    component = (fieldProps) => {
        console.log("field panel component", this.state, fieldProps, this.props)
        let {input, meta, className} = fieldProps
        let fieldName = fieldProps.input.name
        let field = this.fields[fieldName]
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
        return null
      }
    

    render() {
        let desc = this.props.description
        console.log("render fields panel ", this.props, this.state)
        let comp = this
        if(desc && desc.fields) {
            let flds = desc.fields
            if(flds) {
                if(desc.info && desc.info.tabs) {
                    let tabs = new Array()
                    let tabsToDisp = desc.info && desc.info.tabs? desc.info.layout: Object.keys(desc.info.tabs)
                    tabsToDisp.forEach(function(k) {
                        let tabFlds = desc.info.tabs[k];
                        if(tabFlds) {
                            let tabArr = comp.layoutFields(tabFlds, flds, "tabfield formfield", comp.state)
                            tabs.push(
                                <_uikit.Tab label={k} value={k}>
                                {tabArr}
                                </_uikit.Tab>
                            )
                        }
                    })
                    let vertical = desc.info.verticaltabs? true: false
                    return (
                        <_uikit.Tabset vertical={vertical}>
                        {tabs}
                        </_uikit.Tabset>
                    )
                } else {
                    let fldToDisp = desc.info && desc.info.layout? desc.info.layout: Object.keys(flds)
                    let className=comp.props.inline?"inline formfield":"formfield"
                    return this.layoutFields(fldToDisp, flds, className, comp.state)
                }
            }
        }
        return null
    }
}

export { FieldsPanel as FieldsPanel}
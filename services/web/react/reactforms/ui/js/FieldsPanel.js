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
        this.uikit = context.uikit        
        this.configureFields(props, this.uikit)
    }
    componentWillReceiveProps(nextProps, nextState) {
        this.setState(Object.assign({}, this.state, {formValue: nextProps.formValue}))
    }    

    configureFields = (props, uikit) => {
        this.fieldWidgets = {}
        this.fieldProps = {}
        for(let [fieldName, field] of Object.entries(this.fields)) {
            if(field.module) {
                let fldWidget = _res(field.module, field.widget);
                this.fieldWidgets[fieldName] = fldWidget
            } else if(field.list) {
                this.fieldWidgets[fieldName] = FldList
            } else {
                this.fieldWidgets[fieldName] = uikit.Forms.FieldWidget
            }

            let cl= props.inline?"inline formfield m10":"formfield m10"

            if(!field.label && !field.skipLabel) {
                field.label = fieldName
            }
            let fieldProperties = {name: fieldName, entity: field.entity, dataService: field.dataService, dataServiceParams: field.dataServiceParams, 
                loader: field.loader, skipDataLoad: field.skipDataLoad, className: cl, field: field, autoSubmitOnChange: props.autoSubmitOnChange,
                formRef: props.formRef}
            
            this.fieldProps[fieldName] = fieldProperties
        }
    }

    layoutFields = (fldToDisp, flds, className, state) => {
        let fieldsArr = new Array()
        let fldpanel = this
        fldToDisp.forEach(function(k) {
            let field = flds[k]    
            fieldsArr.push( <Field key={field.name} name={field.name} component={fldpanel.component}/>)
            //fieldsArr.push(  <Field key={fd.name} name={fd.name} formValue={state.formValue} autoSubmitOnChange={props.autoSubmitOnChange} fields={flds}
              //   className={cl} formRef={props.formRef} />  )
        })
        return fieldsArr
    }

    

//    <Field key={this.props.name} name={this.props.name} time={this.state.time} component={this.component}/>


    fieldChange = (fldProps) => {
        let comp=this
        return (data, name, evt)=> {
            console.log("fld change", data, name, evt, this.props, this.context, fldProps, fldProps.input.onChange, comp.isRef)
            if(fldProps.input.onChange) {
                if(comp.isRef) {
                    let myRefObj = {}
                    myRefObj[comp.field.name] = {"Id": data, "Type": comp.field.entity}
                    data = myRefObj
                    console.log("set ref value data", myRefObj, data)
                }
                console.log("setting fld value", comp, data, name)
                fldProps.input.onChange(data, name)
            }
        }
    }

    component = (fieldProps) => {
        console.log("component", this.state, fieldProps, this.props)
        let fieldName = fieldProps.input.name
        let field = this.fields[fieldName]
        if(field) {
            let widget = this.fieldWidgets[fieldName]
            let fprops = this.fieldProps[fieldName]
            let fieldChange= this.fieldChange(fieldProps)
            let newProps = Object.assign({}, fieldProps, fprops, {fieldChange, formValue: this.state.formValue})            
            if(field.transformer) {
                let transformerMethod = _reg("Methods", field.transformer)
                newProps = transformerMethod(newProps, this.props.formValue, field, this.fields, this.props, this.state,  this)
            }
            if(field.type == "storableref") {
                let isRef = true
                let ref = newProps.input.value[fieldName]
                if(ref) {
                  newProps.input.value = ref.Id
                }
                console.log("ref props", newProps)
            }
            return React.createElement(widget, newProps, null)
              /*let newProps = fieldProps
            if(this.transformer) {
              newProps = this.transformer(fieldProps, this.props.formValue, this.field, this.context.fields, this.props, this.state,  this)
            }
            if(this.isRef) {
              let ref = newProps.input.value[this.field.name]
              if(ref) {
                newProps.input.value = ref.Id
              }
              console.log("ref props", newProps)
            }
    
    
            let comp = null
    
            let baseComp = null
            if(this.fldWidget) {
              return <this.fldWidget name={this.props.name} className={this.props.className} {...this.state.additionalProperties} time={this.state.time}
                  formValue={this.props.formValue} field={this.field} fieldChange={this.fieldChange(fieldProps)} subFormChange={this.props.subFormChange} autoSubmitOnChange={this.props.autoSubmitOnChange}
                  subform={this.props.subform} formRef={this.props.formRef} parentFormRef={this.props.parentFormRef} {...newProps}/>
            } else {
              if(this.field.list) {
                return <FldList name={this.props.name} baseComponent={this.context.uikit.Forms.FieldWidget} className={this.props.className} ap={this.state.additionalProperties}
                 time={this.state.time}  formValue={this.props.formValue} field={this.field}  fieldChange={this.fieldChange(fieldProps)} autoSubmitOnChange={this.props.autoSubmitOnChange}
                 subFormChange={this.props.subFormChange} subform={this.props.subform} formRef={this.props.formRef} parentFormRef={this.props.parentFormRef} baseProps={newProps}/>
              } else {
                return <this.context.uikit.Forms.FieldWidget  name={this.props.name} className={this.props.className} {...this.state.additionalProperties}
                  time={this.state.time}  formValue={this.props.formValue} field={this.field}  fieldChange={this.fieldChange(fieldProps)} subFormChange={this.props.subFormChange} autoSubmitOnChange={this.props.autoSubmitOnChange}
                  subform={this.props.subform} formRef={this.props.formRef} parentFormRef={this.props.parentFormRef} {...newProps}/>
              }
            }
            */    
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
                                <comp.uikit.Tab label={k} value={k}>
                                {tabArr}
                                </comp.uikit.Tab>
                            )
                        }
                    })
                    let vertical = desc.info.verticaltabs? true: false
                    return (
                        <this.uikit.Tabset vertical={vertical}>
                        {tabs}
                        </this.uikit.Tabset>
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

FieldsPanel.contextTypes = {
    uikit:  PropTypes.object
  };
  
export { FieldsPanel as FieldsPanel}
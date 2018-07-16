import React from 'react';
const PropTypes = require('prop-types');
import {Field} from './Field';

class FieldsPanel extends React.Component {
    constructor(props, context) {
        super(props)
        let desc = props.description
        this.state={formValue: props.formValue, time: props.time}       
        this.uikit = context.uikit        
    }
    componentWillReceiveProps(nextProps, nextState) {
        this.setState(Object.assign({}, this.state, {formValue: nextProps.formValue, time: nextProps.time}))
    }    
    layoutFields = (fldToDisp, flds, className) => {
        let fieldsArr = new Array()
        let comp = this
        let props = this.props
        console.log("layout fields =========", this.state, this.props, flds, fldToDisp)
        fldToDisp.forEach(function(k) {
        let fd = flds[k]
        let cl = className? className + " m10": "m10"
        fieldsArr.push(  <Field key={fd.name} name={fd.name} formValue={comp.state.formValue} formRef={props.formRef} subform={props.subform}
            autoSubmitOnChange={props.autoSubmitOnChange} parentFormRef={props.parentFormRef} parentFormValue={props.parentFormValue} time={comp.state.time} className={cl} />      )
        })
        return fieldsArr
    }

    render() {
        console.log("changing state", this.state)
        let desc = this.props.description
        console.log("desc of form ", desc)
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
                        let tabArr = comp.layoutFields(tabFlds, flds, "tabfield formfield")
                        tabs.push(
                            <comp.uikit.Tab label={k} time={comp.state.time} value={k}>
                            {tabArr}
                            </comp.uikit.Tab>
                        )
                        }
                    })
                    let vertical = desc.info.verticaltabs? true: false
                    return (
                        <this.uikit.Tabset vertical={vertical} time={comp.state.time}>
                        {tabs}
                        </this.uikit.Tabset>
                    )
                } else {
                    let fldToDisp = desc.info && desc.info.layout? desc.info.layout: Object.keys(flds)
                    let className=comp.props.inline?"inline formfield":"formfield"
                    return this.layoutFields(fldToDisp, flds, className)
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
import React from 'react';
const PropTypes = require('prop-types');
import {Field} from './Field';
//import { Field } from 'redux-form'
import {Entity} from './entity';
import {List} from './list';

class FieldsPanel extends React.Component {
    constructor(props, context) {
        super(props)
        let desc = props.description
        this.fields = desc? desc.fields:{}
        this.state={formValue: props.formValue}       
        console.log("fields panel constructor", desc, props)
       // this.configureFields(props)
    }
    componentWillReceiveProps(nextProps, nextState) {
        this.setState(Object.assign({}, this.state, {formValue: nextProps.formValue}))
    }    


    layoutFields = (fldToDisp, flds, className, state) => {
        let fieldsArr = new Array()
        let fldpanel = this
        fldToDisp.forEach(function(k) {
            let field = flds[k]    
            console.log("layout field in fld panel ", k, field)
            //fieldsArr.push( <Field key={field.name} name={field.name} className={className} component={fldpanel.component}/>)
            fieldsArr.push( <Field key={k} field={field} className=" formfield m10 " formValue={fldpanel.state.formValue}/>)
            //fieldsArr.push(  <Field key={fd.name} name={fd.name} formValue={state.formValue} autoSubmitOnChange={props.autoSubmitOnChange} fields={flds}
              //   className={cl} formRef={props.formRef} />  )
        })
        return fieldsArr
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
                            let tabArr = comp.layoutFields(tabFlds, flds, " tabfield formfield ", comp.state)
                            console.log("rendered tab", k, tabFlds, tabArr)
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
                    let className=comp.props.inline?" inline formfield ":" formfield "
                    return this.layoutFields(fldToDisp, flds, className, comp.state)
                }
            }
        }
        return null
    }
}

export { FieldsPanel as FieldsPanel}
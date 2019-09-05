import React from 'react';
import {Panel} from 'reactpages';
const PropTypes = require('prop-types');

class ModuleSettings extends React.Component {
    constructor(props) {
        super(props)
        console.log("props in module settings view", props)
        let formDesc = {type: "form", info: {}}
        if(props.formValue && props.formValue.Module && props.formValue.Module.ParamsForm) {
            formDesc = Object.assign(formDesc, props.formValue.Module.ParamsForm)
        }
        console.log("form desc", formDesc)
        this.formDesc = formDesc
        //Object.assign({}, props.formValue)
        this.state={formData: {}}
    }
    render() {
        let props = this.props
        console.log("props... module settings ", props, this.formDesc)
       // return <Panel actions={this.actions} inline={true} formData={this.state.formData} parentFormRef={this}  subform={true} closePanel={this.closeForm} onSubmit={submit} description={this.props.formDesc} autoSubmitOnChange={this.props.autoSubmitOnChange}/> //, actions, contentStyle)
       return <Panel formData={this.state.formData} subform={true} parentFormRef={props.formRef} description={this.formDesc} /> 
    }
}
export {
    ModuleSettings
}

/*
config:
  className: w100
  submit: MySubmit
  loaderService: formload
  submissionService: test
fields:
  Name:
    type: string
    label: Name
    className: w100
  Email:
    type: string
    label: Email
    className: w100
  Username:
    type: string
    label: User Name
    className: w100
    */
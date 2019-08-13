import React from 'react';
import {Panel} from 'reactpages';
import {  DataSource,  RequestBuilder } from 'uicommon';

class ModuleConfig extends React.Component {
    constructor(props) {
        super(props)
        console.log("module config constructor", props)
        let moduleId = this.getModuleName(props)
        this.loadModuleForm(moduleId)
        this.state = {"module": moduleId, "desc":null, "value": this.getValue(props)}
    }
    componentWillRecieveProps(nextprops) {
        console.log("received next props", nextprops)
        let mod = this.getModuleName(nextprops)   
        if(this.state.moduleId != mod) {
            this.loadModuleForm(mod)
        }
        let val = this.getValue(nextprops)     
        this.setState(Object.assign({}, this.state, {"module": mod, "value": val}))
    }
    onChange = (val, fld, evt) => {
        console.log("on change of module config", val, fld, evt)
        this.props.onChange(val, fld, evt)
    }
    formLoaded=(data)=> {
        console.log("form loaded",data)
        if(data.data) {
            this.setState(Object.assign({}, this.state, {"desc": Object.assign({type:"form", formName: "instance_settings"}, JSON.parse(data.data.Form))}))
        }
    }
    loadModuleForm=(moduleId)=> {
        if(moduleId != "") {
            let req = RequestBuilder.URLParamsRequest({id: moduleId})
            console.log("load module form request", req)
            let prom = DataSource.ExecuteService("moduleconfig", req, {})
            prom.then(this.formLoaded, function(err) {
                console.log(err)
            })
        }
    }
    getModuleName=(props)=> {
        console.log("props", props)
        if(props.formValue.Module) {
            return props.formValue.Module.Id
        } else {
            return ""
        }
    }
    getValue=(props)=> {
        return props.formValue.Settings?props.formValue.Settings: {}
    }
    render() {
        console.log("render of confi", this.state)
        if(this.state.desc) {
            let panelDesc = Object.assign({type:"form"}, this.state.desc)
            console.log(" rerendering config form", panelDesc)
            return (
                <Panel formData={this.state.value} description={panelDesc} onChange={this.onChange} name="Settings"/>
            )    
        }
        return null
    }
}

export {ModuleConfig} 

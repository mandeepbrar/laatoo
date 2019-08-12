import React from 'react';
import {Panel} from 'reactpages';

class ModuleConfig extends React.Component {
    constructor(props) {
        super(props)
        console.log("module config constructor", props)
        this.state = {"module":this.getModuleName(props), "value": this.getValue(props)}
    }
    componentWillRecieveProps(nextprops) {
        let mod = this.getModuleName(nextprops)   
        let val = this.getValue(nextprops)     
        this.setState(Object.assign({}, this.state, {"module": mod, "value": val}))
    }
    getModuleName=(props)=> {
        console.log("props", props)
        return props.formValue.Module.Id
    }
    getValue=(props)=> {
        return props.formValue.Settings
    }
    render() {
        return (
            <Panel id={"module_config_" + this.state.module} formData={this.state.value} type="form"/>
        )
    }
}

export {ModuleConfig} 

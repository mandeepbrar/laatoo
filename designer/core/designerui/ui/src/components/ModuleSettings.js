import React from 'react';

class ModuleSettings extends React.Component {
    constructor(props) {
        super(props)
        console.log("props in module settings view", props)
    }
    render() {
        return <div>These are my settings</div>
    }
}
export {
    ModuleSettings
}
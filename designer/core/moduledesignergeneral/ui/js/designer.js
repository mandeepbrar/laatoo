import React from 'react'
const PropTypes = require('prop-types');
import {Action} from 'reactwebcommon';


class Designer extends React.Component {
    render() {
        return (
            <_uikit.Block>
                <_uikit.Tabset >
                    <_uikit.Tab label="Tab 1" value="tab1">
                        <h1 module="html">my</h1>
                    </_uikit.Tab>
                    <_uikit.Tab label="Tab 2" value="tab2">
                        <h1 module="html">my form2</h1>
                    </_uikit.Tab>
                </_uikit.Tabset>
            </_uikit.Block>
        )
    }
}

export default Designer
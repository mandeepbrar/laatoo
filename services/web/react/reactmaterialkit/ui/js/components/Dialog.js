import React from 'react';
import {Button, Dialog,  DialogActions,  DialogContent,  DialogContentText,  DialogTitle} from '@material-ui/core';
import {Block} from 'reactwebcommon';

class MDialog extends React.Component {
    constructor(props) {
        super(props)
    }
    render() {
        return (
            <Dialog modal={this.props.Modal} open={true} onClose={this.props.onClose}>
                <DialogContent>
                    <Block title={this.props.title} className="dialog" closeBlock={this.props.onClose} contentStyle={this.props.contentStyle}>
                    {this.props.component}
                    </Block>
                </DialogContent>
                <DialogActions>
                {this.props.actions}
                </DialogActions>
            </Dialog>
        )        
    }
}

export {MDialog as Dialog}

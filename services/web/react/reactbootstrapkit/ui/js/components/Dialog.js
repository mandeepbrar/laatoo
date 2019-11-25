import React from 'react';
import {Modal, Button} from 'react-bootstrap';
import {Block} from 'reactwebcommon';

class Dialog extends React.Component {
    constructor(props) {
        super(props)
    }
    render() {
        return (
            <Modal titleClassName="errorTitle" contentStyle={this.props.contentStyle} show={true}  onHide={this.props.onClose}>
                <Modal.Header closeButton> <Modal.Title className="dialogTitle">{this.props.title}</Modal.Title>  </Modal.Header>
                <Modal.Body className="dialogComponent">
                {this.props.component}
                </Modal.Body>
                <Modal.Footer>{this.props.actions}</Modal.Footer>
            </Modal>
        )        
    }
}

export {Dialog}

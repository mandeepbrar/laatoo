'use strict';

import React from 'react';
import Modal from 'react-bootstrap/lib/Modal'
import OverlayTrigger from 'react-bootstrap/lib/OverlayTrigger'
import Button from 'react-bootstrap/lib/Button'
import t from 'tcomb-form';

class ViewFilter extends React.Component {
  constructor(props) {
    super(props)
    this.state = {showModal: false }
    this.openFilter = this.openFilter.bind(this)
    this.cancel = this.cancel.bind(this)
    this.close = this.close.bind(this)
    this.modalDialog = this.modalDialog.bind(this)
  }
  cancel() {
    this.setState({ showModal: false });
  }
  close() {
    if(this.props.setFilter) {
      let data = this.refs.form.getValue()
      if (!data) {
        return;
      }
      data = Object.assign({}, data);
      this.props.setFilter(data);
    }
    this.setState({ showModal: false });
  }

  modalDialog() {
    return (
      <Modal show={this.state.showModal} onHide={this.cancel}>
        <Modal.Header closeButton>
          <Modal.Title>{this.props.title}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <t.form.Form ref="form" type={this.props.schema} value={this.props.formdata} options={this.props.schemaOptions}/>
        </Modal.Body>
        <Modal.Footer>
          <Button onClick={this.close}>{this.props.goBtn}</Button>
        </Modal.Footer>
      </Modal>
    )
  }

  openFilter() {
    this.setState({ showModal: true });
  }
  render() {
    return(
      <div onClick={this.openFilter}>
        {this.props.children}
        {this.modalDialog()}
      </div>
    )
  }
}

export {
  ViewFilter as ViewFilter
}

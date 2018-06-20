import React from 'react';
import { connect } from 'react-redux';
import {Modal, Button} from 'react-bootstrap';

class DialogHandler extends React.Component {
  constructor(props) {
    super(props)
    this.state = {message: props.message, type: props.type, open: false, time: props.time}
    this.handleClose = this.handleClose.bind(this)
  }
  handleClose() {
    this.setState({message: "", open:false})
  }
  componentWillReceiveProps(nextprops) {
    if(nextprops.time != this.state.time) {
      let open = false
      if( nextprops.type === "Message" ||  nextprops.type === "Error" ){
        open = nextprops.message != null && nextprops.message !=""
      } else {
        open = (nextprops.component!=null)
      }
      this.setState({message: nextprops.message, type: nextprops.type, open: open, time: nextprops.time})
    }
  }
  render() {
    let contentStyle = this.props.contentStyle
    if(!contentStyle) {
      contentStyle = {minWidth:400, maxWidth: 450}
    }
    if(!this.state.open) {
      return <div></div>
    }
    switch (this.state.type) {
      case "Error":
        return  (
          <Modal titleClassName="errorTitle" contentStyle={{minWidth:300, maxWidth: 350}} show={this.state.open}  onHide={this.handleClose}>
            <Modal.Header closeButton> <Modal.Title>Error</Modal.Title>  </Modal.Header>
            <Modal.Body>{this.state.message}</Modal.Body>
            <Modal.Footer> <Button onTouchTap={this.handleClose}>Close</Button> </Modal.Footer>
          </Modal>
        )
      break;
      case "Message":
        return  (
          <Modal titleClassName="msgTitle" contentStyle={{minWidth:300, maxWidth: 350}} show={this.state.open}  onHide={this.handleClose}>
            <Modal.Body>{this.state.message}</Modal.Body>
            <Modal.Footer> <Button onTouchTap={this.handleClose}>Close</Button> </Modal.Footer>
          </Modal>
        )
        break
      default:
        return (
          <Modal titleClassName="errorTitle" contentStyle={{minWidth:300, maxWidth: 350}} show={this.state.open}  onHide={this.handleClose}>
            <Modal.Header closeButton> <Modal.Title className="dialogTitle">{this.props.title}</Modal.Title>  </Modal.Header>
            <Modal.Body className="dialogComponent">
              {this.props.component}
            </Modal.Body>
            <Modal.Footer>{this.props.actions}</Modal.Footer>
          </Modal>
        )
    }
  }

}

const mapStateToProps = (state, ownProps) => {
  if(!state.Dialogs.Content) {
    return {
      time: state.Dialogs.Time
    }
  }
  return {
    component: state.Dialogs.Content.Component,
    actions: state.Dialogs.Content.Actions,
    contentStyle: state.Dialogs.Content.ContentStyle,
    message: state.Dialogs.Content.Message,
    title: state.Dialogs.Content.Title,
    type: state.Dialogs.Type,
    time: state.Dialogs.Time
  }
}

const Dialogs = connect(
  mapStateToProps,
  null
)(DialogHandler);

export default Dialogs;

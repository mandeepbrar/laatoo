import React from 'react';
import { connect } from 'react-redux';
import {Button, Snackbar, Dialog,  DialogActions,  DialogContent,  DialogContentText,  DialogTitle} from '@material-ui/core';
import {Block} from 'reactwebcommon';


class DialogHandler extends React.Component {
  constructor(props) {
    super(props)
    this.state = {message: props.message, type: props.type, open: false, time: props.time}
    this.handleClose = this.handleClose.bind(this)
  }
  handleClose() {
    this.setState({message: "", open:false})
    if(this.props.onClose) {
      this.props.onClose()
    }
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
      contentStyle = {minWidth:400, backgroundColor: 'white'}
    }
    if(!this.state.open) {
      return <div></div>
    }

    switch (this.state.type) {
      case "Error":
        return  <Dialog actions={<Button raised label="Close" onTouchTap={this.handleClose}/>} title="Error" titleClassName="errorTitle" modal={true}
            contentStyle={{minWidth:300, maxWidth: 350}} open={this.state.open} onRequestClose={this.handleClose} >
            <Block className="errorMessage">{this.state.message}</Block>
          </Dialog>
      break;
      case "Message":
        return <Snackbar open={this.state.open} message={this.state.message} autoHideDuration={4000}/>
      break
      default:
        console.log("rendering dialog...........", this.props)
        return (
         <Dialog open={this.state.open} modal={true} onRequestClose={this.handleClose}>
           <Block title={this.props.title} className="dialog" closeBlock={this.handleClose} contentStyle={contentStyle}>
           {this.props.component}
           </Block>
           <DialogActions>
            {this.props.actions}
           </DialogActions>
         </Dialog>
        )
    }
  }

}

const mapStateToProps = (state, ownProps) => {
  console.log("state dialog", state, state.Dialogs, ownProps)
  if(!state.Dialogs ) {
    return {}
  }
  if(!state.Dialogs.Content) {
    return {
      time: state.Dialogs.Time
    }
  }
  return {
    component: state.Dialogs.Content.Component,
    onClose: state.Dialogs.Content.OnClose,
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

import React from 'react';
import { connect } from 'react-redux';
import FlatButton from 'material-ui/FlatButton';
import Snackbar from 'material-ui/Snackbar';
import Dialog from 'material-ui/Dialog';
import RaisedButton from 'material-ui/RaisedButton';

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
        return  <Dialog actions={<RaisedButton label="Close" onTouchTap={this.handleClose}/>} title="Error" titleClassName="errorTitle" modal={true}
            contentStyle={{minWidth:300, maxWidth: 350}}open={this.state.open} onRequestClose={this.handleClose} >
            <div className="errorMessage">{this.state.message}</div>
          </Dialog>
      break;
      case "Message":
        return <Snackbar open={this.state.open} message={this.state.message} autoHideDuration={4000}/>
      break
      default:
        return <Dialog actions={this.props.actions} title={<div className="dialogTitle">{this.props.title}<FlatButton className="closeButton" style={{minWidth:25}} label="x" onTouchTap={this.handleClose}/></div>}
           modal={true} contentStyle={contentStyle} open={this.state.open} onRequestClose={this.handleClose} >
          <div className="dialogComponent">
            {this.props.component}
          </div>
        </Dialog>
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

import React from 'react';
import { connect } from 'react-redux';
import RaisedButton from 'material-ui/RaisedButton';

class MessageHandler extends React.Component {
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
      let open = nextprops.message != null && nextprops.message !=""
      this.setState({message: nextprops.message, type: nextprops.type, open: open, time: nextprops.time})
    }
  }
  render() {
    return (
      <_uikit.Block >
      {
        this.state.open?
          <_uikit.Block>
          {
            (this.state.type =="Error")?
            <_uikit.Dialog actions={<_uikit.ActionButton label="Close" onTouchTap={this.handleClose}/>} title="Error" titleClassName="primaryBGColor1 white p10" modal={true}
                contentStyle={{minWidth:300, maxWidth: 350}}
                open={this.state.open} onRequestClose={this.handleClose} >
                <_uikit.Block className="ptb10">{this.state.message}</_uikit.Block>
            </_uikit.Dialog>
            :
              <_uikit.Message open={this.state.open} message={this.state.message} autoHideDuration={4000}/>
          }
          </_uikit.Block>
        :null
      }

      </_uikit.Block>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    message: state.Messages.Message,
    type: state.Messages.Type,
    time: state.Messages.Time
  }
}

const Messages = connect(
  mapStateToProps,
  null
)(MessageHandler);

export default Messages;

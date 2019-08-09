import React from 'react';
import { connect } from 'react-redux';

class DialogHandler extends React.Component {
  constructor(props) {
    super(props)
    this.state = {open: false, time: props.time}
    this.handleClose = this.handleClose.bind(this)
  }
  handleClose() {
    this.setState({open:false})
  }
  componentWillReceiveProps(nextprops) {
    if(nextprops.time != this.state.time) {
      this.setState({open: (nextprops.component!=null), time: nextprops.time})
    }
  }
  render() {
    let widget = null
    if(!this.state.open) {
      return <div></div>
    }
    console.log("opening dialog", this.props)
    let contentStyle = this.props.contentStyle
    if(!contentStyle) {
      contentStyle = {minWidth:400, maxWidth: 450}
    }
    return (
        <_uikit.Dialog actions={this.props.actions}
          title={<_uikit.Block className="primaryBGColor1 white row col-xs-between">
            <b className="col-xs-10">{this.props.title}</b>
            <_uikit.ActionButton className="white" style={{minWidth:25}} label="x" onClick={this.handleClose}/>
          </_uikit.Block>}
           modal={true} contentStyle={contentStyle} open={this.state.open} onRequestClose={this.handleClose} >
          <_uikit.Block className="p10">
            {this.props.component}
          </_uikit.Block>
        </_uikit.Dialog>
    )
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
    title: state.Dialogs.Content.Title,
    time: state.Dialogs.Time
  }
}

const Dialogs = connect(
  mapStateToProps,
  null
)(DialogHandler);

export default Dialogs;

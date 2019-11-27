import React from 'react';
import { connect } from 'react-redux';

class InteractionHandler extends React.Component {
  constructor(props) {
    super(props)
    this.state = {open: (props.component!=null), time: props.time}
    this.handleClose = this.handleClose.bind(this)
  }
  handleClose() {
    console.log("closing interaction")
    this.setState({open:false})
    Window.closeInteraction(this.props.interactiontype)
  }
  componentWillReceiveProps(nextprops) {
    if(nextprops.time != this.state.time) {
      this.setState({open: (nextprops.component!=null), time: nextprops.time})
    }
  }
  render() {
    let widget = null
    if(!this.state.open) {
      return <_uikit.Block />
    }
    let interactionComp = _uikit[this.props.interactiontype]
    console.log("opening interaction", this.props, interactionComp)
    let contentStyle = this.props.contentStyle
    if(!contentStyle) {
      contentStyle = {minWidth:400, maxWidth: 450}
    }
    let propsForComp = {actions: this.props.actions, title: this.props.title, onClose: this.handleClose, modal: true,
      contentStyle: contentStyle, titleStyle: this.props.titleStyle, open: this.state.open, component:this.props.component}
    console.log("about to create element", interactionComp)
    let comp = React.createElement(interactionComp, propsForComp)
    console.log("comp to render", comp)
    return comp
  }
}


/*

        <this.interactionComp actions={this.props.actions}
          title={<_uikit.Block className="primaryBGColor1 white row col-xs-between">
            <b className="col-xs-10">{this.props.title}</b>
            <_uikit.ActionButton className="white" style={{minWidth:25}} label="x" onClick={this.handleClose}/>
          </_uikit.Block>}
           modal={true} contentStyle={contentStyle} open={this.state.open} onRequestClose={this.handleClose} >
          <_uikit.Block className="p10">
            {this.props.component}
          </_uikit.Block>
        </this.interactionComp>
*/
const mapStateToProps = (state, ownProps) => {
  if(!state.Interactions.Content) {
    return {
      time: state.Interactions.Time
    }
  }
  return {
    component: state.Interactions.Content.component,
    interactiontype: state.Interactions.Content.interactiontype,
    actions: state.Interactions.Content.actions,
    contentStyle: state.Interactions.Content.contentStyle,
    titleStyle: state.Interactions.Content.titleStyle,
    title: state.Interactions.Content.title,
    time: state.Interactions.Time
  }
}

const Interactions = connect(
  mapStateToProps,
  null
)(InteractionHandler);

export default Interactions;

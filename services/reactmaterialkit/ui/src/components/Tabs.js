import Tabs, { Tab as MUITab} from 'material-ui/Tabs';
import React from 'react';
import AppBar from 'material-ui/AppBar';
import PropTypes from 'prop-types';

class Tabset extends React.Component {
  constructor(props) {
    super(props)
    this.state = {selectedTab: null, value: null}
    this.childTabs = {}
  }
  addTab = (value, tab) => {
    console.log("adding tab++++++++++=", tab)
    if(!this.state.value) {
      this.setState({ selectedTab: tab, value: value});
    }
    this.childTabs[value] = tab;
  }
  getChildContext() {
    return {tabset: this};
  }
  handleChange = (event, value) => {
    let selTab = this.childTabs[value]
    if(selTab) {
      this.setState({ selectedTab: selTab, value: value});
    }
  };
  render() {
    return (
      <this.context.uikit.Block className={this.props.className}>
        <Tabs value={this.state.value} onChange={this.handleChange}>
        {this.props.children}
        </Tabs>
        {this.state.selectedTab}
      </this.context.uikit.Block>
    )
  }
}

Tabset.contextTypes = {
  uikit: PropTypes.object
};
Tabset.childContextTypes = {
  tabset: PropTypes.object
};


class Tab extends React.Component {
  constructor(props, ctx) {
    super(props)
    this.value = props.value? props.value : props.label
    if(ctx.tabset) {
      ctx.tabset.addTab(this.value, props.children)
    }
  }
  render() {
    console.log("rendering tab", this, this.props, this.context)
    return (
      <MUITab label={this.props.label} value={this.value}  {...this.props} icon={this.props.icon} />
    )
  }
}

Tab.contextTypes = {
  tabset: PropTypes.object
};

export {
  Tabset,
  Tab
}

import {AppBar, Button, Tabs, Tab as MUITab} from '@material-ui/core';
import React from 'react';
import PropTypes from 'prop-types';

class Tabset extends React.Component {
  constructor(props) {
    super(props)
    let value = null
    if(props.children.length>0) {
      value = props.children[0].props.value? props.children[0].props.value: props.children[0].props.label
    }
    this.state = {selectedTab: null, value}
    this.childTabs = {}
    console.log("processing tabset", props)
  }
  isVertical = () => {
    return this.props.vertical? true : false
  }
  addTab = (value, tab) => {
    if(this.state.value == value) {
      this.setState({ selectedTab: tab, value: value});
    }
    this.childTabs[value] = tab;
  }
  tabChanged = (value, tab) => {
    this.childTabs[value] = tab;
    if(this.state.value == value) {
      this.setState({ selectedTab: tab, value: value});
    }
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
        this.isVertical()?
        <this.context.uikit.Block className="row">
          <this.context.uikit.Block className="col-xs-3">
          {this.props.children}
          </this.context.uikit.Block>
          <this.context.uikit.Block  className="col-xs-9">
          {this.state.selectedTab}
          </this.context.uikit.Block>
        </this.context.uikit.Block>
        :
        <this.context.uikit.Block className={this.props.className}>
          <Tabs value={this.state.value} scrollable onChange={this.handleChange}>
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
    console.log("processing tab", props, ctx)
    this.value = props.value? props.value : props.label
    this.children = props.children
    if(ctx.tabset) {
      this.vertical = ctx.tabset.isVertical()
      ctx.tabset.addTab(this.value, props.children)
    }
  }
  componentWillReceiveProps(nextProps, nextState) {
    if(this.children != nextProps.children) {
      this.children = nextProps.children
      if(this.context.tabset) {
        this.context.tabset.tabChanged(this.value, nextProps.children)
      }
    }
  }
  tabClick = () => {
    this.context.tabset.handleChange(null, this.value)
  }
  render() {
    return (
        this.vertical?
        <this.context.uikit.Block className="tab w100">
          <Button onClick={this.tabClick}>{this.props.label}</Button>
        </this.context.uikit.Block>
        :
        <MUITab label={this.props.label} className={this.props.className? this.props.className + " tab " : "tab"} value={this.value}  {...this.props} icon={this.props.icon} />
    )
  }
}

Tab.contextTypes = {
  uikit: PropTypes.object,
  tabset: PropTypes.object
};

export {
  Tabset,
  Tab
}

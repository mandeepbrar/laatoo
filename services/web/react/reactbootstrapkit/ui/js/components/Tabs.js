import {Tabs, Tab as BSTab, Nav} from 'react-bootstrap'
import React from 'react';
import PropTypes from 'prop-types';

class Tabset extends React.Component {
  constructor(props) {
    super(props)
    /*let value = null
    if(props.children.length>0) {
      value = props.children[0].props.value? props.children[0].props.value: props.children[0].props.label
    }*/
    this.state = {tabs:[]}
    //this.childTabs = {}
    //console.log("processing tabset", props)
  }
  isVertical = () => {
    return this.props.vertical? true : false
  }
  addTab = (label, tab) => {
  /*  if(this.state.value == value) {
      this.setState({ selectedTab: tab, value: value});
    }*/
    let tabs = this.state.tabs
    tabs.push(
      <BSTab.Pane eventKey={label}>
        {tab}
      </BSTab.Pane>
    )
    console.log("added tab", label, tab, tabs)
    this.setState({tabs})
  }
  getChildContext() {
    return {tabset: this};
  }
  render() {
    console.log("tabset bootstrap", this.state)
    return (
      //defaultActiveKey="first">
        this.isVertical()?
        <BSTab.Container> 
          <_uikit.Block className="row">
            <_uikit.Block className="col-xs-3">
              <Nav variant="pills" className="flex-column">
              {this.props.children}
              </Nav>
            </_uikit.Block>
            <BSTab.Content className="col-xs-9">
              {this.state.tabs}
            </BSTab.Content>
          </_uikit.Block>
        </BSTab.Container>
        :
        <BSTab.Container> 
          <_uikit.Block>
            <_uikit.Block >
              <Nav variant="pills">
              {this.props.children}
              </Nav>
            </_uikit.Block>
            <BSTab.Content>
              {this.state.tabs}
            </BSTab.Content>
          </_uikit.Block>
        </BSTab.Container>
/*        <Tabs activeKey={this.state.value} id="tab"  className={this.props.className}>
        {this.props.children}
        </Tabs>*/
    )
  }
}

Tabset.childContextTypes = {
  tabset: PropTypes.object
};


class Tab extends React.Component {
  constructor(props, ctx) {
    super(props)
    console.log("processing tab", props, ctx)
  //  this.label = props.label; //props.value? props.value : props.label
    //this.children = props.children
    if(ctx.tabset) {
      this.vertical = ctx.tabset.isVertical()
      ctx.tabset.addTab(this.props.label, props.children)
    }
  }
  /*componentWillReceiveProps(nextProps, nextState) {
    if(this.children != nextProps.children) {
      this.children = nextProps.children
      if(this.context.tabset) {
        this.context.tabset.tabChanged(this.value, nextProps.children)
      }
    }
  }
  tabClick = () => {
    this.context.tabset.handleChange(null, this.value)
  }*/
  render() {
      console.log("bootstrap tab", this.props, this.vertical)
    return (
        <Nav.Item className="tab ">
          <Nav.Link eventKey={this.props.label}>{this.props.label}</Nav.Link>
        </Nav.Item>
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

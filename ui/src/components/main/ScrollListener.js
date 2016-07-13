import React from 'react';

class ScrollListener extends React.Component {
  constructor(props) {
    super(props)
    this.handleScroll = this.handleScroll.bind(this)
    if(props.windowScroll) {
        let windowNumber = Math.floor(window.scrollY / window.innerHeight)
        this.state = {windowNumber}
    }
  }
  handleScroll(evt) {
    if(this.props.windowScroll) {
      let windowNumber = Math.floor(window.scrollY / window.innerHeight)
      if(windowNumber != this.state.windowNumber) {
        this.props.windowScroll(windowNumber)
        this.setState({windowNumber})
      }
    }
    if(this.props.onScrollEnd) {
      var node = this.refs.scrollListener;
      let bottomReached = node.getBoundingClientRect().bottom <= window.innerHeight
      if(bottomReached && this.props.onScrollEnd) {
        this.props.onScrollEnd()
      }
    }
  }
  componentDidMount() {
    window.addEventListener('scroll', this.handleScroll);
  }
  componentWillUnmount() {
    window.removeEventListener('scroll', this.handleScroll);
  }
  render() {
    return (
      <div ref="scrollListener" key={this.props.key} style={this.props.style} className={this.props.className}>
      {this.props.children}
      </div>
    )
  }
}

export {ScrollListener as ScrollListener} ;

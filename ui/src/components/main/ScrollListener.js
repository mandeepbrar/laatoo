import React from 'react';

class ScrollListener extends React.Component {
  constructor(props) {
    super(props)
    this.handleScroll = this.handleScroll.bind(this)
  }
  handleScroll(evt) {
    var node = this.refs.scrollListener;
    let bottomReached = node.getBoundingClientRect().bottom <= window.innerHeight
    if(bottomReached && this.props.onScrollEnd) {
      this.props.onScrollEnd()
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

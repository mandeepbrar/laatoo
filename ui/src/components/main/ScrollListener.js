import React from 'react';

class ScrollListener extends React.Component {
  constructor(props) {
    super(props)
    this.handleScroll = this.handleScroll.bind(this)
    if(props.windowScroll) {
        let windowNumber = Math.floor(window.scrollY / window.innerHeight)
        this.state = {windowNumber}
    }
    this.state = {scrolledOut: false, scrolledIn: false}
  }
  handleScroll(evt) {
    if(this.props.windowScroll) {
      let windowNumber = Math.floor(window.scrollY / window.innerHeight)
      if(windowNumber != this.state.windowNumber) {
        this.props.windowScroll(windowNumber)
        this.setState({windowNumber})
      }
    }
    if(this.props.onScrollEnd || this.props.onScrollIn) {
      var node = this.refs.scrollListener;
      let endPos = window.innerHeight
      if(this.props.scrollEndPos) {
        endPos = this.props.scrollEndPos
      }
      if(node != null) {
        let boundingrect = node.getBoundingClientRect()
        if(boundingrect.bottom <= endPos && !this.state.scrolledOut) {
          if(!this.state.scrolledOut && this.props.onScrollEnd) {
            this.props.onScrollEnd(boundingrect.bottom)
          }
          this.setState({scrolledOut: true, scrolledIn: false})
        }
        if(this.state.scrolledOut && boundingrect.bottom > endPos) {
          if(!this.state.scrolledIn && this.props.onScrollIn) {
            this.props.onScrollIn(boundingrect.bottom)
          }
          this.setState({scrolledOut: false, scrolledIn: true})
        }        
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

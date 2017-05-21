import React from 'react';

class Image extends React.Component {
  constructor(props) {
    super(props)
    if(!props.skipPrefix) {
      props.skipPrefix = false
    }
  }
  render() {
    let source = this.props.src
    if(!source || source.length==0) {
      if(this.props.children) {
        return this.props.children
      }
      return null
    }
    if (!this.props.skipPrefix && !this.props.src.startsWith("http")) {
      source = this.props.prefix + source
    }
    let i = <img src={source} {...this.props.modifier} className={this.props.className} style={this.props.style}/>
    if(this.props.link) {
      return (
        <a target={this.props.target} href={this.props.link}>{i}</a>
      )
    } else {
      return i
    }
  }
}

export {
  Image as Image
}

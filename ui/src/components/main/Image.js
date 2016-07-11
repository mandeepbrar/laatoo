import React from 'react';

class Image extends React.Component {
  constructor(props) {
    super(props)
  }
  render() {
    let source = this.props.src
    if(!source || source.length==0) {
      if(this.props.children) {
        return this.props.children
      }
      return null
    }
    if (!this.props.src.startsWith("http")) {
      source = this.props.prefix + source
    }
    return (
      <img src={source} {...this.props.modifier} className={this.props.className} style={this.props.style}/>
    )
  }
}

export {
  Image as Image
}

import React from 'react';

class Image extends React.Component {
  constructor(props) {
    super(props)
  }
  render() {
    if(!this.props.src || this.props.length == 0) {
      return (
        <div></div>
      )
    }
    let source = this.props.src
    if (source && source.length>0 && !this.props.src.startsWith("http")) {
      source = this.props.prefix + source
    }
    return (
      <img src={source} {...this.props.modifier} style={this.props.style}/>
    )
  }
}

export {
  Image as Image
}

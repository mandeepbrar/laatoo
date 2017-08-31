import React from 'react';
import {Application} from 'laatoocommon'
import { Image, View } from 'react-native';


class PrefixImage extends React.Component {
  constructor(props) {
    super(props)
  }
  render() {
    let skipPrefix = this.props.skipPrefix? this.props.skipPrefix: false
    let source = this.props.src
    if(!source || source.length==0) {
      if(this.props.children) {
        return this.props.children
      }
      return null
    }
    if (!skipPrefix && !this.props.src.startsWith("http")) {
      source = this.props.prefix + source
    }
    console.log("source file ", source, this.props)
    let srcobj = null
    if(this.props.local) {
      srcobj = require(source)
    } else {
      srcobj={uri:source}
    }
    console.log("src obj--", srcobj)

    let style = Object.assign({}, {flex: 1, resizeMode: 'cover', height: Application.MaxHeight, width: null}, this.props.style)
    console.log("style of the image*************", style)
    let i = <Image style={style} {...this.props.modifier} source={srcobj}/>
    if(this.props.link) {
      return i
       /*(
        i //<a target={this.props.target} href={this.props.link}>{i}</a>
      )*/
    } else {
      return i
    }
  }
}

export {
  PrefixImage as Image
}

import React from 'react';
import Dropzone from 'react-dropzone';
//import {Image} from 'reactwebcommon'
import {  Response,  DataSource,  RequestBuilder } from 'uicommon';
const PropTypes = require('prop-types');

var module;

function Initialize(appName, ins, mod, settings, def, req) {
  module =  this;
  module.properties = Application.Properties[ins];
  module.settings = settings;
}


class FileUpload extends React.Component {
  constructor(props, context) {
    super(props)
    console.log("fileupload", props, context)

    this.drop = this.drop.bind(this)
    this.clear = this.clear.bind(this)
    this.getComponent = this.getComponent.bind(this)
    this.getDropZoneComponent = this.getDropZoneComponent.bind(this)
    this.getValueComponent = this.getValueComponent.bind(this)
    this.state = {value: props.value}
  }
  drop(files) {
    console.log("files dropped", files)
    if(files.length >0) {
      let data = new FormData()
      if(this.props.multiple) {
        files.forEach((file) => {
          data.append("Data", file, file.name);
        })
      } else {
        data.append("Data", files[0], files[0].name);
      }
      var config = {
        progress: function(progressEvent) {
          var percentCompleted = progressEvent.loaded / progressEvent.total;
        }
      };
      console.log(data, files[0])
      let comp = this;
      let req = RequestBuilder.DefaultRequest(null, data)
      let prom = DataSource.ExecuteService(this.props.uploadService, req, config)
      prom.then(
        function (res) {
          if(comp.props.multiple) {
            //let val = comp.props.value? comp.props.value.split(","):[]
            let newVal = res.data.join(",")
            let val = comp.state.value? [comp.state.value, newVal].join(","): newVal
            if(comp.props.onChange) {
              comp.props.onChange(val)
            }
            console.log("response of uploading images", val)
          } else {
            comp.setState({value: files[0].name})
            if(comp.props.onChange) {
              comp.props.onChange(res.data[0], files[0].name)
            }
          }
          if(comp.props.successCallback) {
            comp.props.successCallback(res)
          }
        },
        function (res) {
          console.log("res", res, "failureCallback", comp.props.failureCallback, "props", comp.props)
          if(comp.props.failureCallback) {
            comp.props.failureCallback(res)
          }
          console.log(res);
        });
    }
  }
  getValueComponent() {
    let items=[]
    if(this.props.multiple) {
      console.log("value of files ", this.state.value)
      let val = this.state.value? this.state.value.split(","): []
      val.forEach((item)=> {
        items.push(
          <_uikit.Block className="w100">{item}</_uikit.Block>
        )
      })
    } else {
      items.push(
        <_uikit.Block className="w100">{this.state.value}</_uikit.Block>
      )
    }
    return(
      <_uikit.Block className="w100">
        {module.properties.success.text}
        {items}
      </_uikit.Block>
    )
  }
//  {this.props.clear? this.props.clear(this.clear):<_uikit.ActionButton onClick={this.clear}>Clear</_uikit.ActionButton>}

  clear(evt) {
    evt.preventDefault();
    this.setState({value: null})
    if(this.props.onChange) {
      this.props.onChange(null)
    }
  }
  getDropZoneComponent() {
    let multiple = this.props.multiple? this.props.multiple: false
    return(
      <_uikit.Block className="upload" style={{display:"block"}}>
        <Dropzone onDrop={this.drop} className={this.props.dropzoneClass} multiple={multiple}>
          {
            this.props.text ?
            this.props.text
            :
            <_uikit.Block>{module.properties.dropzone.text}</_uikit.Block>
          }
        </Dropzone>
      </_uikit.Block>
    )
  }
  getComponent() {
  /*    let _uikit = this.context._uikit
    if (choice == "url") {
      return(
        <_uikit.Block style={{display:"block"}} className="url">
          <input className="ma30" type="text" ref="imageediturl" placeholder="link"/>
          <button className="ma10 rightalign" role="button" onClick={this.uselink}>Use Image</button>
        </_uikit.Block>
      )
    } else {*/
      return this.getDropZoneComponent()
    //}
  }
  
  render () {
    let style = null
    if(this.props.style) {
      style = this.props.style
    } else {
      style = {minHeight:"300"}
    }
//    let showChoice = !this.props.multiple && !this.props.hideChoice
    return (
      <_uikit.Block className="fileupload" style={style}>
        {this.getComponent()}
        {this.getValueComponent()}
      </_uikit.Block>
    )
    /*if(this.props.value) {
      )
    } else {
      let component = this.getComponent()
      if(showChoice) {
        return (
          <_uikit.Block style={style} className="imageedit">
            {component}
          </_uikit.Block>
        )
      } else {
        return component
      }
    }*/
  }
}

export {
  FileUpload,
  Initialize
};

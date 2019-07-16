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
    this.uikit = context.uikit

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
    let uikit = this.context.uikit
    let items=[]
    if(this.props.multiple) {
      console.log("value of files ", this.state.value)
      let val = this.state.value? this.state.value.split(","): []
      val.forEach((item)=> {
        items.push(
          <uikit.Block className="w100">{item}</uikit.Block>
        )
      })
    } else {
      items.push(
        <uikit.Block className="w100">{this.state.value}</uikit.Block>
      )
    }
    return(
      <uikit.Block className="w100">
        {module.properties.success.text}
        {items}
      </uikit.Block>
    )
  }
//  {this.props.clear? this.props.clear(this.clear):<uikit.ActionButton onClick={this.clear}>Clear</uikit.ActionButton>}

  clear(evt) {
    evt.preventDefault();
    this.setState({value: null})
    if(this.props.onChange) {
      this.props.onChange(null)
    }
  }
  getDropZoneComponent() {
    let multiple = this.props.multiple? this.props.multiple: false
    let uikit = this.context.uikit
    return(
      <uikit.Block className="upload" style={{display:"block"}}>
        <Dropzone onDrop={this.drop} className={this.props.dropzoneClass} multiple={multiple}>
          {
            this.props.text ?
            this.props.text
            :
            <uikit.Block>{module.properties.dropzone.text}</uikit.Block>
          }
        </Dropzone>
      </uikit.Block>
    )
  }
  getComponent() {
  /*    let uikit = this.context.uikit
    if (choice == "url") {
      return(
        <uikit.Block style={{display:"block"}} className="url">
          <input className="ma30" type="text" ref="imageediturl" placeholder="link"/>
          <button className="ma10 rightalign" role="button" onClick={this.uselink}>Use Image</button>
        </uikit.Block>
      )
    } else {*/
      return this.getDropZoneComponent()
    //}
  }
  
  render () {
    let uikit = this.context.uikit
    let style = null
    if(this.props.style) {
      style = this.props.style
    } else {
      style = {minHeight:"300"}
    }
//    let showChoice = !this.props.multiple && !this.props.hideChoice
    return (
      <uikit.Block className="fileupload" style={style}>
        {this.getComponent()}
        {this.getValueComponent()}
      </uikit.Block>
    )
    /*if(this.props.value) {
      )
    } else {
      let component = this.getComponent()
      if(showChoice) {
        return (
          <uikit.Block style={style} className="imageedit">
            {component}
          </uikit.Block>
        )
      } else {
        return component
      }
    }*/
  }
}

FileUpload.contextTypes = {
  uikit: PropTypes.object
};

export {
  FileUpload,
  Initialize
};

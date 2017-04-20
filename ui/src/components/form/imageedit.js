import t from 'tcomb-form';
import React from 'react';
import Dropzone from 'react-dropzone';
import {Image} from '../main/Image'
import {  Response,  DataSource,  RequestBuilder } from '../../sources/DataSource';


class ImageChooser extends React.Component {
  constructor(props) {
    super(props)
    this.drop = this.drop.bind(this)
    this.uselink = this.uselink.bind(this)
    this.clear = this.clear.bind(this)
    this.getComponent = this.getComponent.bind(this)
    this.getChoiceComponent = this.getChoiceComponent.bind(this)
    this.getDropZoneComponent = this.getDropZoneComponent.bind(this)
    this.getValueComponent = this.getValueComponent.bind(this)
    this.handleChoiceChange = this.handleChoiceChange.bind(this)
    this.state = {imagechooser: "upload"}
    if(this.props.imageStyle) {
      this.imageStyle = this.props.imageStyle
    } else {
      this.imageStyle = {height:"220"}
    }
  }
  drop(files) {
    console.log("files dropped", files)
    if(files.length >0) {
      let data = new FormData()
      if(this.props.multiple) {
        files.forEach((file) => {
          data.append(file.name, file);
        })
      } else {
        data.append(files[0].name, files[0]);
      }
      var config = {
        progress: function(progressEvent) {
          var percentCompleted = progressEvent.loaded / progressEvent.total;
        }
      };
      let comp = this;
      let req = RequestBuilder.DefaultRequest(null, data)
      let prom = DataSource.ExecuteService(this.props.uploadService, req, config)
      prom.then(
        function (res) {
          if(comp.props.multiple) {
            //let val = comp.props.value? comp.props.value.split(","):[]
            let newVal = res.data.join(",")
            let val = comp.props.value? [comp.props.value, newVal].join(","): newVal
            comp.props.onChange(val)
            console.log("response of uploading images", val)
          } else {
            comp.props.onChange(res.data[0], files[0].name)
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
  uselink(evt) {
    evt.preventDefault();
    this.props.onChange(this.refs.imageediturl.value)
  }
  clear(evt) {
    evt.preventDefault();
    this.props.onChange(null)
  }
  getChoiceComponent() {
    return (
      <div className="w100">
        <input type="radio" name="imagechooser" value="upload" checked={this.state.imagechooser=="upload"} onChange={this.handleChoiceChange}/> Upload
        &nbsp;&nbsp;
        <input type="radio" name="imagechooser" value="url" checked={this.state.imagechooser=="url"} onChange={this.handleChoiceChange}/> URL
      </div>
    )
  }
  getValueComponent() {
    let items=[]
    if(this.props.multiple) {
      console.log("value of image ", this.props.value)
      let val = this.props.value? this.props.value.split(","): []
      val.forEach((item)=> {
        items.push(
          <Image src={item} modifier={{thumbnail:true}} prefix={this.props.prefix} style={this.imageStyle} />
        )
      })
    } else {
      items.push(
        <Image src={this.props.value} modifier={{thumbnail:true}} prefix={this.props.prefix} style={this.imageStyle} />
      )
    }
    return(
      <div className="w100">
        {items}
        {this.props.clear? this.props.clear(this.clear):<button className="btn" role="button" onClick={this.clear}>Clear</button>}
      </div>
    )
  }
  getDropZoneComponent() {
    let multiple = this.props.multiple? this.props.multiple: false
    return(
      <div className="upload" style={{display:"block"}}>
        <Dropzone onDrop={this.drop} className={this.props.dropzoneClass} multiple={multiple}>
          {
            this.props.text ?
            this.props.text
            :
            <div>Try dropping some files here, or click to select files to upload.</div>
          }
        </Dropzone>
      </div>
    )
  }
  getComponent(choice) {
    if (choice == "url") {
      return(
        <div style={{display:"block"}} className="url">
          <input className="ma30" type="text" ref="imageediturl" placeholder="link"/>
          <button className="ma10 rightalign" role="button" onClick={this.uselink}>Use Image</button>
        </div>
      )
    } else {
      return this.getDropZoneComponent()
    }
  }
  handleChoiceChange(e) {
    this.setState({imagechooser: e.target.value});
  }
  render () {
    let style = null
    if(this.props.style) {
      style = this.props.style
    } else {
      style = {minHeight:"300"}
    }
    let showChoice = !this.props.multiple && !this.props.hideChoice

    if(this.props.value) {
      return (
        <div className="imageedit" style={style}>
          {!this.props.hideURL?<p><strong>URL: </strong>{this.props.value}</p>:null}
          {this.props.multiple? this.getDropZoneComponent(): null}
          {this.getValueComponent()}
        </div>
      )
    } else {
      let component = this.getComponent(this.state.imagechooser)
      if(showChoice) {
        let choiceComponent = this.getChoiceComponent()
        return (
          <div style={style} className="imageedit">
            {choiceComponent}
            {component}
          </div>
        )
      } else {
        return component
      }
    }
  }
}

class ImageEdit extends t.form.Component { // extend the base class
  getTemplate() {
    return (locals) => {
      let config = {}
      if(locals.config) {
        config = locals.config
      }
      let choose = (file) => {
        console.log(file);
      }
      console.log("image config", config)
      return (
          <ImageChooser value={locals.value} style={config.style} hideURL={config.hideURL} hideChoice={config.hideChoice} clear={config.clear}
              successCallback={config.successCallback} failureCallback={config.failureCallback} text={config.text} dropzoneClass={config.dropzoneClass}
              multiple={config.multiple} onChange={locals.onChange} prefix={locals.config.prefix} imageStyle={config.imageStyle} uploadService={config.service}/>
      );
    };
  }
}


export {
  ImageEdit as ImageEdit,
  ImageChooser as ImageChooser
};

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
    this.handleChoiceChange = this.handleChoiceChange.bind(this)
    this.state = {imagechooser: "upload"}
  }
  drop(files) {
    if(files.length >0) {
      let data = new FormData()
      data.append(files[0].name, files[0]);
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
          let img = res.data[0];
          comp.props.onChange(img)
        },
        function (res) {
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
  getComponent(choice) {
    if (choice == "url") {
      return(
        <div style={{display:"block"}} className="url">
          <input className="ma30" type="text" ref="imageediturl" placeholder="link"/>
          <button className="ma10 rightalign" role="button" onClick={this.uselink}>Use Image</button>
        </div>
      )
    } else {
      return(
        <div className="upload" style={{display:"block"}}>
          <Dropzone onDrop={this.drop} multiple={false}>
            <div>Try dropping some files here, or click to select files to upload.</div>
          </Dropzone>
        </div>
      )
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
      style = {height:"300"}
    }
    let imageStyle = {height:"220"}
    if(this.props.imageStyle) {
      imageStyle = this.props.imageStyle
    } else {
      imageStyle = {height:"220"}
    }
    if(this.props.value) {
      return (
        <div className="imageedit" style={style}>
          {!this.props.hideURL?<p><strong>URL: </strong>{this.props.value}</p>:null}
          <Image src={this.props.value} modifier={{thumbnail:true}} prefix={this.props.prefix} style={imageStyle} />
          {this.props.clear? this.props.clear(this.clear):<button className="btn" role="button" onClick={this.clear}>Clear</button>}
        </div>
      )
    } else {
      let choice = this.state.imagechooser
      let component = this.getComponent(choice)
      if(!this.props.hideChoice) {
        return (
          <div style={style} className="imageedit">
            <input type="radio" name="imagechooser" value="upload" checked={this.state.imagechooser=="upload"} onChange={this.handleChoiceChange}/> Upload
            &nbsp;&nbsp;
            <input type="radio" name="imagechooser" value="url" checked={this.state.imagechooser=="url"} onChange={this.handleChoiceChange}/> URL
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
      return (
          <ImageChooser value={locals.value} style={config.style} hideURL={config.hideURL} hideChoice={config.hideChoice} clear={config.clear}
              onChange={locals.onChange} prefix={locals.config.prefix} imageStyle={config.imageStyle} uploadService={config.service}/>
      );
    };
  }
}


export {
  ImageEdit as ImageEdit,
  ImageChooser as ImageChooser
};

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
          console.log("complete", percentCompleted);
        }
      };
      let comp = this;
      let req = RequestBuilder.DefaultRequest(null, data)
      let prom = DataSource.ExecuteService(this.props.uploadService, req, config)
      prom.then(
        function (res) {
          let img = res.data[0];
          console.log(img);
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
    if(this.props.value) {
      return (
        <div className="imageedit" style={{height:"300px"}}>
          <p><strong>URL: </strong>{this.props.value}</p>
          <Image src={this.props.value} modifier={{thumbnail:true}} prefix={this.props.prefix} style={{width: 'auto', height: 220}} />
          <button className="btn" role="button" onClick={this.clear}>Clear</button>
        </div>
      )
    } else {
      let choice = this.state.imagechooser
      let component = this.getComponent(choice)
      return (
        <div style={{height:"300px"}} className="imageedit">
          <input type="radio" name="imagechooser" value="upload" checked={this.state.imagechooser=="upload"} onChange={this.handleChoiceChange}/> Upload
          &nbsp;&nbsp;
          <input type="radio" name="imagechooser" value="url" checked={this.state.imagechooser=="url"} onChange={this.handleChoiceChange}/> URL
          {component}
        </div>
      )
    }
  }
}

class ImageEdit extends t.form.Component { // extend the base class
  getTemplate() {
    return (locals) => {
      console.log("locals....", locals)
      let choose = (file) => {
        console.log(file);
      }
      return (
        <div>
          <label>{locals.label}</label>
          <ImageChooser value={locals.value} onChange={locals.onChange} prefix={locals.config.prefix} uploadService={locals.config.service}/>
        </div>
      );
    };
  }
}


export {
  ImageEdit as ImageEdit
};

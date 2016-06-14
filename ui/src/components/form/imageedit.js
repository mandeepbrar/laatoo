import t from 'tcomb-form';
import React from 'react';
import Dropzone from 'react-dropzone';
import Tab from 'react-bootstrap/lib/Tab'
import Image from 'react-bootstrap/lib/Image'
import Tabs from 'react-bootstrap/lib/Tabs'
import {  Response,  DataSource,  RequestBuilder } from '../../sources/DataSource';


class ImageChooser extends React.Component {
  constructor(props) {
    super(props)
    this.drop = this.drop.bind(this)
    this.uselink = this.uselink.bind(this)
    this.clear = this.clear.bind(this)
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
          if(comp.props.prefix) {
            img = comp.props.prefix + img
          }
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
  render () {
    console.log("value ", this.props.value)
    if(this.props.value) {
      return (
        <div class="container" style={{height:"300px"}} className="m20">
          <p><strong>URL: </strong>{this.props.value}</p>
          <Image src={this.props.value} thumbnail style={{width: 'auto', height: 220}} />
          <button className="btn" role="button" onClick={this.clear}>Clear</button>
        </div>
      )
    } else {
      return (
        <div style={{height:"300px"}}  className="m20">
          <Tabs id="imagechooser">
              <Tab eventKey={1} title="URL">
                <input className="form-control ma30" type="text" ref="imageediturl" placeholder="link"/>
                <button className="btn col-xs-offset-1 ma10 pull-right" role="button" onClick={this.uselink}>Use Image</button>
              </Tab>
              <Tab eventKey={2} title="Upload">
                <div className="m20">
                  <Dropzone onDrop={this.drop} multiple={false}>
                    <div>Try dropping some files here, or click to select files to upload.</div>
                  </Dropzone>
                </div>
              </Tab>
          </Tabs>
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

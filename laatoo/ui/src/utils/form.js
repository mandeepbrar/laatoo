import t from 'tcomb-form';
import React from 'react';
import TinyMCEInput from 'react-tinymce-input';
import ReactPlayer from 'react-player';
import Dropzone from 'react-dropzone';
import Tab from 'react-bootstrap/lib/Tab'
import Image from 'react-bootstrap/lib/Image'
import Tabs from 'react-bootstrap/lib/Tabs'

const TINYMCE_CONFIG = {
  'language'  : 'en',
  'theme'     : 'modern',
  'toolbar'   : 'bold italic underline strikethrough hr | bullist numlist | link unlink | undo redo | alignleft aligncenter alignright | spellchecker code',
  'menubar'   : false,
  'statusbar' : true,
  'resize'    : true,
  'plugins'   : 'link,spellchecker,paste',
  'theme_modern_toolbar_location' : 'top',
  'theme_modern_toolbar_align': 'left'
};

class TextEdit extends t.form.Component { // extend the base class
  getTemplate() {
    return (locals) => {
      let data = locals.value
      if(!data) {
        data = '';
      }
      return (
        <div>
          <label>{locals.label}</label>
          <TinyMCEInput  value={data}  tinymceConfig={TINYMCE_CONFIG} onChange={locals.onChange} />
        </div>
      );
    };
  }
}

class VideoEdit extends t.form.Component { // extend the base class
  getTemplate() {
    return (locals) => {
      return (
        <div>
          <label>{locals.label}</label>
          <t.form.Form type={t.Str} value={locals.value} onChange={locals.onChange}/>
          <ReactPlayer url={locals.value} />
        </div>
      );
    };
  }
}

class ImageChooser extends React.Component {
  constructor(props) {
    super(props)
    this.drop = this.drop.bind(this)
    this.uselink = this.uselink.bind(this)
    this.clear = this.clear.bind(this)
  }
  drop(files) {
    console.log('Received files: ', files);
/*    var uploader = new FileUploader({
        url: document.Application.Backend + document.Services.upload.url,
        queueLimit: 1
    });
    uploader.onCompleteItem = function(fileItem, response, status, headers) {
        if (status == 200 && response.length > 0) {
            $uibModalInstance.close(response[0]);
        }
    };*/
    if(files.length >0) {
      this.props.onChange(files[0].name)
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
      let choose = (file) => {
        console.log(file);
      }
      return (
        <div>
          <label>{locals.label}</label>
          <ImageChooser value={locals.value} onChange={locals.onChange}/>
        </div>
      );
    };
  }
}


export {
  TextEdit as TextEdit,
  VideoEdit as VideoEdit,
  ImageEdit as ImageEdit
};

'use strict';

import React from 'react';
import { Field } from 'redux-form'
import {RequestBuilder, DataSource, EntityData} from 'uicommon';
const PropTypes = require('prop-types');

var modrequire = null;

function Initialize(appName, ins, mod, settings, def, req) {
  modrequire = req;
}

class FieldWrapper extends React.Component {
  constructor(props, context) {
    super(props)
    this.field = context.fields[props.name]
    console.log("fields created", props, this.field, context)
    if(!this.field.label && !this.field.skipLabel) {
      this.field.label = props.name
    }
    this.state = { additionalProperties:{} }
    let errorMethod = function (resp) {
      console.log("could not load data", resp)
    };
    if(this.field.transformer) {
      let method = _reg("Methods", this.field.transformer)
      this.transformer = method;
    }
    if(this.field.module) {
      let mod = modrequire(this.field.module);
      this.fldWidget = mod[this.field.widget];
    } else {
      if (this.field.widget == "Select") {
        if(this.field.items){
          this.state.additionalProperties.items = this.field.items
        }
        if(this.field.dataService) {
          let req = RequestBuilder.DefaultRequest(null, this.field.dataServiceParams);
          DataSource.ExecuteService(this.field.dataService, req).then(this.selectOptionsLoaded, errorMethod);
        }
        if(this.field.type == "entity") {
          EntityData.ListEntities(this.field.name).then(this.selectOptionsLoaded, errorMethod);
        }
      }
    }
  }

  selectOptionsLoaded = (resp) => {
    if(resp && resp.data && resp.data.length >0) {
      this.loadedData = resp.data
      let data = {}
      let items=[]
      let textField = this.field.textField? this.field.textField: "Name"
      let valueField = this.field.valueField? this.field.valueField: "Id"
      resp.data.forEach(function(item) {
        let text = item[textField];
        text = text? text: item["Title"]
        text = text? text: item[valueField]
        items.push({text: text, value: item[valueField]})
      })
      /*let imgField = this.props.config? this.props.config.imgField: null
      resp.data.forEach((item)=> {
        if(this.props.qualifier) {
          if (!this.props.qualifier(item))  {
            return
          }
        }
        if(imgField) {
          data[item.Id] = {text: item.Title, image: item[imgField]}
          console.log("item ", item, data[item.Id])
        } else {
          data[item.Id] = item.Title
        }
      })
      let options = this.getItems(this.props, data)*/

      this.setState(Object.assign({}, this.state, {additionalProperties: {items: items} }))
    }
  }
  change = (name, val) => {
    console.log("field changed", name, val)
  }

  component = (props) => {
    if(this.transformer) {
      props = this.transformer(props)
    }
    console.log("Wrapped form field", props, this.transformer)
    if(this.fldWidget) {
      return <this.fldWidget {...props}/>
    } else {
      return <this.context.uikit.Forms.FieldWidget {...props}/>
    }
  }

  render() {
    console.log("rendering field+", this.props, this.props.name, this.field, this.fldWidget);
    return (
      <Field key={this.props.name} name={this.props.name} className={this.props.className} {...this.state.additionalProperties} field={this.field} component={this.component}/>
    )
  }

}

FieldWrapper.contextTypes = {
  fields: PropTypes.object,
  uikit:  PropTypes.object
};

export { FieldWrapper as Field, Initialize}

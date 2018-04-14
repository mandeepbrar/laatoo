'use strict';

import React from 'react';
import { Field } from 'redux-form'
import {RequestBuilder, DataSource, EntityData} from 'uicommon';
const PropTypes = require('prop-types');
import {FldList} from './FldList';

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
    this.state = {time: props.time? props.time: Date.now(), additionalProperties:{}}
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
        if(!this.field.skipDataLoad && this.field.type == "entity") {
          EntityData.ListEntities(this.field.entity).then(this.selectOptionsLoaded, errorMethod);
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
      console.log(" setting items", items)
      this.setState(Object.assign({}, this.state, {time: Date.now(), additionalProperties: {items: items} }))
    }
  }

  componentWillReceiveProps(nextProps, nextState) {
    if(nextProps.time > this.state.time) {
      this.setState(Object.assign({}, this.state, {time: nextProps.time}))
    }
  }

  component = (fieldProps) => {
    console.log("component", this.state)
    let newProps = fieldProps
    if(this.transformer) {
      newProps = this.transformer(fieldProps, this.props.formValue, this.field, this.context.fields, this.props, this.state,  this)
    }
    let comp = null
    let baseComp = null
    if(this.fldWidget) {
      return <this.fldWidget name={this.props.name} className={this.props.className} {...this.state.additionalProperties} time={this.state.time} field={this.field} {...newProps}/>
    } else {
      if(this.field.list) {
        return <FldList name={this.props.name} baseComponent={this.context.uikit.Forms.FieldWidget} className={this.props.className} ap={this.state.additionalProperties} time={this.state.time} field={this.field} baseProps={newProps}/>
      } else {
        return <this.context.uikit.Forms.FieldWidget  name={this.props.name} className={this.props.className} {...this.state.additionalProperties} time={this.state.time} field={this.field} {...newProps}/>
      }
    }
  }

  render() {
    console.log("changing state", this.state)
    return (
      <Field key={this.props.name} name={this.props.name} time={this.state.time} component={this.component}/>
    )
  }

}

FieldWrapper.contextTypes = {
  fields: PropTypes.object,
  uikit:  PropTypes.object
};

export { FieldWrapper as Field, Initialize}

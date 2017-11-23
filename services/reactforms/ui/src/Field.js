'use strict';

import React from 'react';
import { Field } from 'redux-form'
const PropTypes = require('prop-types');

class FieldWrapper extends React.Component {
  constructor(props, context) {
    super(props)
    console.log("fields created", props, this.field)
    this.field = context.fields[props.name]
    if(!this.field.label && !this.field.skipLabel) {
      this.field.label = props.name
    }
    this.state = { additionalProperties:{} }
    if (this.field.widget == "Select") {
      if(this.field.items){
        this.state.additionalProperties.items = this.field.items
      } else if(this.field.dataService) {
        let errorMethod = function (resp) {
          console.log("could not load data", resp)
        };
        let req = RequestBuilder.DefaultRequest(null, this.field.dataServiceParams);
        DataSource.ExecuteService(this.field.dataService, req).then(this.selectOptionsLoaded, errorMethod);
      }
    }
  }

  selectOptionsLoaded = (resp) => {
    if(resp && resp.data && resp.data.length >0) {
      this.loadedData = resp.data
      let data = {}
      let items=[]
      let textField = this.field.textField? this.field.textField: "Title"
      let valueField = this.field.valueField? this.field.valueField: "Id"
      resp.data.forEach(function(item) {
        items.push({text: item[textField], value: item[valueField]})
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
  render() {
    return (
      <Field name={this.props.name} className={this.props.className} {...this.state.additionalProperties} field={this.field} component={this.context.uikit.Forms.FieldWidget}/>
    )
  }

}

FieldWrapper.contextTypes = {
  fields: PropTypes.object,
  uikit:  PropTypes.object
};

export { FieldWrapper as Field}

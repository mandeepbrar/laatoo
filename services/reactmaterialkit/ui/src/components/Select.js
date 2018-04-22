import React from 'react';
import { Select} from 'material-ui';
import { MenuItem } from 'material-ui/Menu';
import { InputLabel } from 'material-ui/Input';
import { FormControl} from 'material-ui/Form';
import {RequestBuilder, DataSource, EntityData} from 'uicommon';
import PropTypes from 'prop-types';
import {LoadableComponent} from 'reactwebcommon'

class SelectComp extends LoadableComponent {
  constructor(props) {
    super(props)
    let items = null
    this.valueField = props.valueField? props.valueField: "value"
    this.textField = props.textField? props.textField: "text"
    if(props.items) {
      items = this.setItems(props.items)
    }
    this.state = {items: items, value: this.getValue(props)}
  }

  componentWillReceiveProps(nextProps) {
    console.log("next props", nextProps)
    let st = {}
    if(this.props.items != nextProps.items) {
        let indexedItems = this.setItems(nextProps.items)
        st.items = indexedItems
    }
    if(this.props.value != nextProps.value) {
      st.value = this.getValue(nextProps)
    }
    console.log("next props st", this.state, st)
    if(st.value || st.items) {
      this.setState(Object.assign({}, this.state, st))
    }
  }


  getValue = (props) => {
    if(props.selectItem && props.value && typeof(props.value) == 'object') {
      return props.value[this.valueField]
    } else {
      return props.value
    }
  }

  optionChanged = (evt) => {
    console.log("evt", evt, evt.target.value)
    let p = this.props
    this.setState(Object.assign({}, this.state, {value: evt.target.value}))
    if(p.onChange) {
      let val = p.selectItem? this.state.items[evt.target.value]: evt.target.value
      console.log("sendin value*********", p.selectItem, val, event.target.value, this.state)
      p.onChange(val, evt.target.name, evt)
    }
  }

  dataLoaded = (data) => {
    this.loadedData = data
    /*let comp = this
    let props = this.props
    //let data = {}
    let items=[]
    data.forEach(function(item) {
      let text = item[comp.textField];
      text = text? text: item["Title"]
      text = text? text: item[comp.valueField]
      items.push({text: text, value: item[comp.valueField]})
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
    let indexedItems = this.setItems(data)
    this.setState(Object.assign({}, this.state, {items: indexedItems}))
  }

  setItems = (items) => {
    console.log("setting new items")
    let comp = this
    let props = this.props
    comp.items=[]
    let indexedItems = {}
    if(items) {
      items.forEach(function(item) {
        let val = item[comp.valueField]
        let text = item[comp.textField]
        console.log("setting new items val===============================", props.selectItem, val, text, item, comp.textField, comp.valueField)
        if(props.selectItem) {

          indexedItems[val] = item
        }
        comp.items.push(
          <MenuItem className={props.itemClass} value={val}>{text}</MenuItem>
        )
      })
    }
    return indexedItems
  }

  render() {
    console.log("render select ********", this.state)
      let p = this.props
      let v = this.state.value? this.state.value: ""
      return (
        <FormControl className={(p.className?p.className + " ":"") + p.name + " formcontrol "}>
          <InputLabel htmlFor={p.name}>{p.label}</InputLabel>
          <Select name={p.name} floatingLabelText={p.label} label={p.label} errorText={p.errorText}
            onChange={this.optionChanged} value={v} className={p.name + " select " + (p.controlClassName?p.controlClassName:"")}>
          {this.items}
          </Select>
        </FormControl>
      )
  }
}

SelectComp.contextTypes = {
  getFormValue: PropTypes.func
};

export {SelectComp as Select}

import React from 'react';
import PropTypes from 'prop-types';
import {LoadableComponent} from 'reactwebcommon'
import Form from 'react-bootstrap/Form'

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
    console.log("evt", evt, evt.target.value, this.state, this.props)
    let p = this.props
    this.setState(Object.assign({}, this.state, {value: evt.target.value}))
    if(p.onChange) {
    //  let val = p.selectItem? this.state.items[evt.target.value]: evt.target.value
      //console.log("sendin value*********", p.selectItem, val, event.target.value, this.state)
      p.onChange(evt)
    }
  }

  dataLoaded = (data) => {
    this.loadedData = data
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
        if(item) {
          let val = item[comp.valueField]
          let text = item[comp.textField]
          console.log("setting new items val===============================", props.selectItem, val, text, item, comp.textField, comp.valueField)
          if(props.selectItem) {
            indexedItems[val] = item
          }
          comp.items.push(
            <option className={props.itemClass} value={val}>{text}</option>
          )
        }
      })
    }
    return indexedItems
  }

  render() {
    console.log("render select ********", this.state)
      let p = this.props
      let v = this.state.value? this.state.value: ""
      return (
        <Form.Group controlId={p.name} className={(p.className?p.className + " ":"") + p.name + " formcontrol "}>
          <Form.Label>{p.label}</Form.Label>
          <Form.Control as="select" onChange={this.optionChanged} value={v}  name={p.name} 
            className={p.name + " select " + (p.controlClassName?p.controlClassName:"")}>
            {this.items}
          </Form.Control>
          {p.errorText?<Form.Text className="text-muted">{p.errorText}</Form.Text>:null}
        </Form.Group>
      )
  }
}

SelectComp.contextTypes = {
  getFormValue: PropTypes.func
};

export {SelectComp as Select}

import React from 'react'
import {Action} from 'reactwebcommon'
const PropTypes = require('prop-types');
import {Form} from '../Form'
import {Field} from '../Field'

class ListSummary extends React.Component {
  constructor(props, ctx) {
    super(props)
    let value = props.value? props.value: [] 
    this.state = {value}
  }

  componentWillReceiveProps(nextProps) {
    console.log("componentWillReceiveProps  for list", nextProps)
    let value = nextProps.value? nextProps.value: [] 
    this.setState({value})
  }

  formsubmit = (data)=> { 
    console.log("summary form submit ", data)
    this.props.overlayComponent(null)
    this.props.onChange(data, this.props.name)
  }

  beforeValueSet = (data) => {
    return { listvalue: data}
  }

  preSubmit = (data) => {
    return _tn(data["listvalue"], null)
  }

  editList = () => {
    let props = this.props
    props.overlayComponent(
      <Form form="list_edit" name={props.name} onFormSubmit={this.formsubmit} formData={this.state.value} beforeValueSet={this.beforeValueSet} preSubmit={this.preSubmit} >
        <Field name="listvalue" label={props.name} list={true} mode="dialog" className={props.className + " w100"}/>    
      </Form>
    )
  }
  render() {
    let props = this.props
    console.log("list render ", this.state, props)
    let listToJoin = this.state.value
    if(listToJoin.length > 4) {
        listToJoin = listToJoin.slice(0,4)
    }
    
    let retStr = listToJoin.map((item)=> {
        return _tn(item[props.titleField], item)
    }).join(",")

    //field should not be used in any of the forms items to keep it generic
    return <_uikit.Block  className={" row stringlist"}>
          <_uikit.Block className="col-xs-12 label">{props.label}</_uikit.Block>
          <_uikit.Block className="value col-xs-10">
            {_tn(retStr, "<No Data>")}
          </_uikit.Block>
          <Action action={{actiontype:"method"}} className=" col-xs-2" method={this.editList}>
            <_uikit.Icons.EditIcon/>
          </Action>
    </_uikit.Block>
  }
}

ListSummary.contextTypes = {
  overlayComponent: PropTypes.func
};

export {
    ListSummary as ListSummary
}

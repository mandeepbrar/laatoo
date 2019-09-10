import React from 'react'
import {Action, LoadableComponent} from 'reactwebcommon'

class SelectEntity extends React.Component {
  constructor(props) {
    super(props)
    this.state={value: props.value, items: props.items}
  }
  componentWillReceiveProps(nextProps) {
    console.log("on change of select entity--", this.state, nextProps)
    this.setState(Object.assign({}, this.state, {value: nextProps.value, items: nextProps.items}))
  }
  onChange=(value)=> {
    console.log("on change of select entity--", value)
    this.setState(Object.assign({}, this.state, {value}))
  }
  saveValue=()=> {
    console.log("svaing value", this.state, this.props)
    this.props.submit(this.state.value)
  }
  render() {
    console.log("rendering select entity", this.state)
    let {fld} = this.props
    let fldDesc = {label: fld.label, name: fld.name, widget: { "name": "Select"}, type: "entity", items: this.state.items}
    return <_uikit.Block  className="row between-xs">
      <_uikit.Block className="left col-xs-10" >
        <_uikit.Field className="w100" field={fldDesc} onChange={this.onChange} selectItem={true} value={this.state.value}/>
      </_uikit.Block>
      <_uikit.Block className="right">
        <Action action={{actiontype:"method"}} className="edit p10" method={this.saveValue}>
          <_uikit.Icons.EditIcon />
        </Action>
        <Action action={{actiontype:"method"}} className="remove p10" method={this.props.close}>
          <_uikit.Icons.DeleteIcon />
        </Action>
      </_uikit.Block>
    </_uikit.Block>
  }
}

export default SelectEntity
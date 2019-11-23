import React from 'react';
import Form from 'react-bootstrap/Form'

const CheckboxComp = (props)=> {

  /*let onChange = (evt) => {
    console.log("on change of text field", evt)
    props.onChange(evt)
  }
  console.log("props of text field", props)*/
  return(
    <Form.Group controlId={this.name} className={props.className + " checkbox "} >
        <Form.Check type="checkbox" label={props.label}  onCheck={props.onChange} />
    </Form.Group>    
  )
}

export {CheckboxComp as Checkbox}

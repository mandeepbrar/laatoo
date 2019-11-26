import React from 'react';
import Form from 'react-bootstrap/Form'

const CheckboxComp = (props)=> {

  /*let onChange = (evt) => {
    console.log("on change of text field", evt)
    props.onChange(evt)
  }*/
  console.log("props of checkbox", props)
  return(
    <Form.Group controlId={props.name} className={props.className + " checkbox "} >
        <Form.Check type="checkbox" label={props.label} value={props.value} onCheck={props.onChange} />
    </Form.Group>    
  )
}

export {CheckboxComp as Checkbox}

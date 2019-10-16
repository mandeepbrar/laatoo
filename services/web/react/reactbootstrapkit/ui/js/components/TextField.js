import React from 'react';
import Form from 'react-bootstrap/Form'

const TextFieldComp = (props)=> {

  let onChange = (evt) => {
    console.log("on change of text field", evt)
    props.onChange(evt)
  }
  console.log("props of text field", props)
  return(
    <Form.Group controlId={props.name} className={props.className + " textfield " }>
      {props.label?<Form.Label>{props.label}</Form.Label>:null}
      <Form.Control name={props.name} type={props.type} placeholder={props.placeholder} onChange={onChange} 
        onBlur={props.onBlur} onFocus={props.onFocus} value={props.value}/>
      {props.errorText?<Form.Text className="text-muted">{props.errorText}</Form.Text>:null}
    </Form.Group>
  )
}

export {TextFieldComp as TextField}

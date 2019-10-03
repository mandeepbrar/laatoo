import React from 'react';
import Form from 'react-bootstrap/Form'

const TextFieldComp = (props)=> {
  console.log("props of text field", props)
  return(
    <Form.Group controlId={this.name} className={this.className + " textfield " }>
      <Form.Label>{this.label}</Form.Label>
      <Form.Control type="text" placeholder={props.placeholder} onChange={this.change} 
        onBlur={props.onBlur} onFocus={props.onFocus} value={this.state.value}/>
      {props.errorText?<Form.Text className="text-muted">{props.errorText}</Form.Text>:null}
    </Form.Group>
  )
}

export {TextFieldComp as TextField}

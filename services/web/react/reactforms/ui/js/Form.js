'use strict';

import React from 'react';
import { connect } from 'react-redux';
import {ActionNames} from './Actions';
import {createAction} from 'uicommon';
import { reduxForm } from 'redux-form';
import {FieldsPanel} from './FieldsPanel';

const PropTypes = require('prop-types');

class BaseForm extends React.Component {
  constructor(props, context) {
    super(props)
    console.log('base form', props, context)
    this.desc = props.description
    this.info = props.info? props.info: null
    if(!this.info) {
      this.info = this.desc && this.desc.info ? this.desc.info : {}
    }
    this.loader = null
    this.formSubmit = null
    let comp = this
    if(props.onFormSubmit) {
      this.formSubmit = (data, successCallback, failureCallback) => {
        props.onFormSubmit(data, successCallback, failureCallback);
      }
    }   
    this.state={formValue: props.formVal}
    this.trackChanges = this.info.trackChanges || props.trackChanges

    this.className = "webform " + (this.info.className? this.info.className :props.id)
    this.formName="myform"

    this.actions = props.actions
    if(!this.actions && this.info.actions) {
      this.actions = _reg('Methods', this.info.actions)
    }
    console.log("webform ", props, context, this.state)
  }

  componentWillMount() {
    this.configureForm(this.props.dispatch, this.props)
    console.log("base form mount", this)    
    if(this.loader) {
      console.log('base form: executing loader', this.loader, this.props, this.context)
      this.loader(this.context.routeParams, this.dataLoaded, this.failureCallback)
    } else {
      this.setData(this.props.formData)      
    }
  }

  /*componentWillReceiveProps(nextProps, nextState) {
    this.dataLoading = false
    let formValue = nextProps.formVal
    let oldFormValue = this.state.formValue
    console.log("webform: next props of componentWillReceiveProps", oldFormValue, formValue, this.state)
    if(formValue != oldFormValue) {
      this.setState(Object.assign({}, this.state, {formValue: nextProps.formVal}))
      if(this.props.onChange) {
        this.props.onChange(formValue, oldFormValue)
      }
    }
    console.log("webform: next props of componentWillReceiveProps", this.props.form, nextProps, nextState, this.state)
  }*/

  getDescription = () => {
    return this.props.description
  }

  configureForm = (dispatch, props) => {
    console.log("configure base form : props ", props)
  }
  preSubmit= (data) => {
    let presub
    let methodName = this.info && this.info.preSubmit? this.info.preSubmit: ""
    if(this.props.preSubmit) {
      if(typeof this.props.preSubmit === 'string') {
        methodName = this.props.preSubmit
      } else {
        presub = this.props.preSubmit
      }
    }
    if(methodName) {
      presub = _reg('Methods', methodName)
    }
    console.log("presubmit called ", presub, data, this.info, this.props, methodName)
    return presub? presub(data): data
  }

  reset = () => {
    this.props.reset()
  }

  getChildContext() {
    return {getFormValue: this.getFormValue, getParentFormValue: this.getParentFormValue};
  }

  failureCallback = () => {

  }

  submitSuccessCallback = (data) => {
    if(this.info.successRedirect) {
      Window.redirect(this.info.successRedirect);
    }
    if(this.info.successRedirectPage) {
      Window.redirect(this.info.successRedirectPage);
    }
  }

  preprocessData = (data) => {
    let beforeValueSet
    let methodName = this.info && this.info.beforeValueSet? this.info.beforeValueSet: ""
    if(this.props.beforeValueSet) {
      if(typeof this.props.beforeValueSet === 'string') {
        methodName = this.props.beforeValueSet
      } else {
        beforeValueSet = this.props.beforeValueSet
      }
    }
    if(methodName) {
      beforeValueSet = _reg('Methods', methodName)
    }
    console.log("beforeValueSet called ", beforeValueSet, data)
    return beforeValueSet? beforeValueSet(data): data
  }

  setData = (formData) => {
    if(this.info.preAssigned) {
      formData = Object.assign({}, formData, this.info.preAssigned)
    }
    formData = this.preprocessData(formData)
    console.log("setData", this.props.form, formData)
    if(formData) {
      let x = this.props.initialize( formData)
      this.props.dispatch(x)  
    }
  }

  dataLoaded = (data) => {
    this.dataLoading = true
    //let formData = Object.assign({}, data.resp.data)
    this.setData(data.resp.data)
  }

  getFormValue = () => {
    return this.state.formValue
  }

  submitFunc = (customFunc) => {
    console.log("invoked", customFunc)
    let comp = this
    let mySubmitFunc = (data) => {
      console.log("ui form submit called", data)
      data = comp.preSubmit(data)
      //let subToCall = _tn(method, comp.formSubmit)    
      console.log("data submit", data)
      if(customFunc) {
        return customFunc(data, comp.submitSuccessCallback, comp.failureCallback)
      } else {
        return comp.formSubmit(data, comp.submitSuccessCallback, comp.failureCallback)
      }
    };
    return this.props.handleSubmit(mySubmitFunc);
  }

  render() {
    let props = this.props
    let cfg = this.info
    console.log("**********************rendering web form****************", props.form, props, this.state, cfg)
    if(_uikit.Form) {
      let formComp = null
      if(cfg && cfg.layout && (typeof(cfg.layout) == "string") ) {
        let display = _reg('Blocks', cfg.layout)
        if(display) {
          console.log("form context", this.props.formContext);
          let root = display(props.formContext, props.description)
          formComp = React.cloneElement(root, { formValue: this.state.formValue, onSubmit: this.submitFunc, className: this.className})
        }
      } else {
        formComp = (
          <_uikit.Form onSubmit={this.submitFunc}>
            {
              props.children?
              props.children
              :
              <FieldsPanel description={props.description} formRef={this} formValue={this.state.formValue} />
            }
          </_uikit.Form>
        )
      }
      return (
        <_uikit.Block className={this.className}>
            {formComp}
            <_uikit.Block className="actionbar p20 right">
              {
                this.actions?
                this.actions(this, this.submitFunc, this.reset, this.setData, this.props.dispatch):
                <_uikit.ActionButton onClick={this.submitFunc()} className="submitBtn">{cfg.submit? cfg.submit: "Submit"}</_uikit.ActionButton>
              }
            </_uikit.Block>
        </_uikit.Block>
      )
    } else {
      return <_uikit.Block/>
    }
  }  
}


BaseForm.childContextTypes = {
  getParentFormValue: PropTypes.func,
  getFormValue: PropTypes.func
};


/** entity needs to be saved to database */
class EntityForm extends BaseForm {
  constructor(props, ctx) {
    super(props, ctx)
    console.log("creating entity form : props ", props,ctx)
  }
  configureForm = (dispatch, props) => {
    console.log("configure entity form : props ", props)
    if(this.info.entity) {
      let entityId = this.info.entityId
      let entityName = this.info.entity
      let svc = this.info.entityService
      let form = this
          //console.log("desc....", entityId, "name", entityName, entityFormCfg)
      if(entityId) {
        this.loader = (routeParams, dataLoaded, failureCallback) => {
          dispatch(createAction(ActionNames.ENTITY_GET, { entityId, entityName}, {successCallback:  form.dataLoaded, failureCallback: failureCallback}));
        } 
      }
      if(!this.formSubmit) {
        if(entityId) {
          if(this.info.put) {
            this.formSubmit = (data, successCallback, failureCallback) => {
              dispatch(createAction(ActionNames.ENTITY_PUT, {data, entityId, entityName}, {reload: form.info.reloadOnUpdate, successCallback, failureCallback}));
            }
          } else {
            this.formSubmit = (data, successCallback, failureCallback) => {
              dispatch(createAction(ActionNames.ENTITY_UPDATE, {data, entityId, entityName}, {reload: form.info.reloadOnUpdate, successCallback, failureCallback}));
            }
          }
        } else {
          this.formSubmit = (data, successCallback, failureCallback) => {
            dispatch(createAction(ActionNames.ENTITY_SAVE, {data, entityName}, {successCallback, failureCallback}));
          }
        }
      }
    }
    console.log("after entity form configuration", this)
  }
}
EntityForm.contextTypes = {
  getFormValue: PropTypes.func,
  routeParams: PropTypes.object
};

/** custom form loading and submission */
class CustomForm extends BaseForm {
  constructor(props, context) {
    super(props, context)
    console.log("creating custom form", props)
  }
  configureForm = (dispatch, props) => {
    let form = this
    if(this.info.loaderService) {
      let loaderServiceParams = {}
      let loaderService = ""
      if(typeof(this.info.loaderService) == "string") {
        loaderService = this.info.loaderService
      } else {
        loaderService = this.info.loaderService.name
        loaderServiceParams = this.info.loaderService.params
      }
      if(loaderService) {
        this.loader = (routeParams, dataLoaded, failureCallback) => {
          dispatch(createAction(ActionNames.LOAD_DATA, Object.assign({}, loaderServiceParams, routeParams), {serviceName: loaderService, successCallback:  form.dataLoaded, failureCallback: failureCallback}));
        }
      }
    }
    if(!this.formSubmit) {
      this.formSubmit = (data, successCallback, failureCallback) => {
        console.log("form submit custom form", data)
        if(form.info) {
          successCallback = form.info.submitSuccess? _reg('Methods', form.info.submitSuccess) : successCallback
          failureCallback = form.info.submitFailure? _reg('Methods', form.info.submitFailure) : failureCallback
        }
        dispatch(createAction(ActionNames.SUBMIT_FORM, data, {serviceName: form.info.submissionService, successCallback: successCallback, failureCallback: failureCallback}));
      }
    }
  }
}
CustomForm.contextTypes = {
  getFormValue: PropTypes.func,
  routeParams: PropTypes.object
};

class WebFormUI extends React.Component {
  constructor(props) {
    super(props)
    let desc = props.description
    let info = props.info? props.info: null
    info = desc && desc.info ? desc.info : {}
    if (info.entity){
      this.formType = reduxForm({form: props.form})(EntityForm)
    } else {
      this.formType = reduxForm({form: props.form})(CustomForm)
    }
  }

  componentWillReceiveProps(nextProps, nextState) {
    console.log("webform ui: next props ", nextProps)
  }


  render() {
    console.log("creating web form", this.formType, this.props, this.props.children)
    return React.createElement(this.formType, this.props, this.props.children)
  }
}

/*
  <t.form.Form ref="form" key={state.key} type={state.schema} value={state.formValue} options={state.so} onChange={this.onChange}/>
  {this.actionButtons}*/

const ReduxForm = reduxForm({})(WebFormUI)
const empty = {}


const mapStateToProps = (state, ownProps) => {
  let desc = ownProps.description
  console.log("redux form state...........=======", state, ownProps, desc)

  let formData = ownProps.formData
  let form  = state.form[ownProps.form]
  console.log("form val", form, ownProps.form, state)
  let formVal = form? form.values: empty
  console.log("form val", formVal, formData)
  return { formVal, formData}
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return { dispatch }
}


const Form = connect(
  mapStateToProps,
  mapDispatchToProps
)(WebFormUI);


export {Form } ;

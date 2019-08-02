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
    this.config = props.config? props.config: desc.config
    this.config = this.config? this.config :{}
    this.loader = null
    this.formSubmit = null
    let comp = this
    if(props.onSubmit) {
      this.formSubmit = (data, successCallback, failureCallback) => {
        data = comp.preSubmit(data)
        console.log("data ", data)
        props.onSubmit(data, {successCallback, failureCallback});
      }
    }   
    this.state={formValue: props.formVal}
    this.trackChanges = this.config.trackChanges || props.trackChanges
    this.uikit = context.uikit

    this.className = "webform " + (this.config.className? this.config.className :"")
    this.formName="myform"

    this.actions = props.actions
    if(!this.actions && this.desc.info.actions) {
      this.actions = _reg('Methods', this.desc.info.actions)
    }
    console.log("webform ", props, context, this.state)
    
    let f = this.uiformSubmit(this.submitSuccessCallback, this.failureCallback)
    this.submitFunc = (customFunc) => { return customFunc? props.handleSubmit(customFunc): props.handleSubmit(f) }
  }

  componentWillMount() {
    this.configureForm(this.props.dispatch, this.props)
    console.log("base form mount", this)
    if(this.loader) {
      console.log('base form: executing loader', this.loader, this.props, this.context)
      this.loader(this.context.routeParams, this.dataLoaded, this.failureCallback)
    }
  }

  componentWillReceiveProps(nextProps, nextState) {
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
  }

  getDescription = () => {
    return this.props.description
  }

  configureForm = (dispatch, props) => {
    console.log("configure base form : props ", props)
  }
  preSubmit= (data) => {
    let presub
    if(this.config) {
      presub = _reg('Method', this.config.preSubmit)
    }
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
    if(this.config) {
      if(this.config.successRedirect) {
        Window.redirect(this.config.successRedirect);
      }
      if(this.config.successRedirectPage) {
        Window.redirect(this.config.successRedirectPage);
      }
    }
  }

  setData = (formData) => {
    if(this.config.dataMapper) {
      let mapper = _reg('Method', this.config.dataMapper)
      formData = mapper(formData)
    }
    let x = this.props.initialize( formData)
    this.props.dispatch(x)
  }

  dataLoaded = (data) => {
    this.dataLoading = true
    //let formData = Object.assign({}, data.resp.data)
    this.setData(data.resp.data)
  }

  getFormValue = () => {
    return this.state.formValue
  }

  uiformSubmit = (success, failure) => {
    let comp = this
    return (data) => {
      comp.formSubmit(data, success, failure)
    }
  }

  render() {
    let props = this.props
    console.log("**********************rendering web form****************", props.form, props, this.state)
    let cfg = this.config
    if(this.uikit.Form) {
      if(cfg.layout && (typeof(cfg.layout) == "string") ) {
        let display = _reg('Blocks', cfg.layout)
        if(display) {
          console.log("form context", this.props.formContext);
          let root = display(props.formContext, props.description, this.uikit)
          return React.cloneElement(root, { formValue: this.state.formValue, onSubmit: this.submitFunc, className: this.className})
        }
      } else {
        return (
          <this.uikit.Form onSubmit={this.submitFunc} className={this.className}>
            <FieldsPanel description={props.description} formRef={this} autoSubmitOnChange={props.autoSubmitOnChange} formValue={this.state.formValue} />
            <this.uikit.Block className="actionbar p20 right">
              {
                this.actions?
                this.actions(this, this.submitFunc, this.reset, this.uikit, this.setData, this.props.dispatch):
                <this.uikit.ActionButton onClick={this.submitFunc()} className="submitBtn">{cfg.submit? cfg.submit: "Submit"}</this.uikit.ActionButton>
              }
            </this.uikit.Block>
          </this.uikit.Form>
        )
      }
    } else {
      return <this.uikit.Block/>
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
    if(this.config.entity) {
      let entityId = this.config.entityId
      let entityName = this.config.entity
      let svc = this.config.entityService
      let form = this
          //console.log("desc....", entityId, "name", entityName, entityFormCfg)
      this.loader = (routeParams, dataLoaded, failureCallback) => {
        if(entityId) {
          dispatch(createAction(ActionNames.ENTITY_GET, { entityId, entityName}, {successCallback:  form.dataLoaded, failureCallback: failureCallback}));
        }
      }
      if(!this.formSubmit) {
        if(entityId) {
          if(this.config.put) {
            this.formSubmit = (data, successCallback, failureCallback) => {
              data = form.preSubmit(data)
              console.log("form submit put", data)
              dispatch(createAction(ActionNames.ENTITY_PUT, {data, entityId, entityName}, {reload: form.config.reloadOnUpdate, successCallback, failureCallback}));
            }
          } else {
            this.formSubmit = (data, successCallback, failureCallback) => {
              data = form.preSubmit(data)
              console.log("form submit update", data)
              dispatch(createAction(ActionNames.ENTITY_UPDATE, {data, entityId, entityName}, {reload: form.config.reloadOnUpdate, successCallback, failureCallback}));
            }
          }
        } else {
          this.formSubmit = (data, successCallback, failureCallback) => {
            data = form.preSubmit(data)
            console.log("form submit save", data)
            dispatch(createAction(ActionNames.ENTITY_SAVE, {data, entityName}, {successCallback, failureCallback}));
          }
        }
      }
    }
  }
}
EntityForm.contextTypes = {
  uikit: PropTypes.object,
  getFormValue: PropTypes.func,
  routeParams: PropTypes.object
};


/** entity needs to be provided to its parent component */
class SubEntityForm extends BaseForm {
  constructor(props, context) {
    super(props, context)
  }

  getParentForm = () => {
    return this.parentFormRef
  }

  getParentFormValue = () => {
    return this.parentFormValue
  }

  configureForm = (dispatch, props) => {
    this.setData(this.props.formData)
  }
}
SubEntityForm.contextTypes = {
  uikit: PropTypes.object,
  getFormValue: PropTypes.func,
  routeParams: PropTypes.object
};

/** custom form loading and submission */
class CustomForm extends BaseForm {
  constructor(props, context) {
    super(props, context)
  }
  configureForm = (dispatch, props) => {
    let form = this
    if(this.config.loaderService) {
      let loaderServiceParams = {}
      let loaderService = ""
      if(typeof(this.config.loaderService) == "string") {
        loaderService = this.config.loaderService
      } else {
        loaderService = this.config.loaderService.name
        loaderServiceParams = this.config.loaderService.params
      }
      if(loaderService) {
        this.loader = (routeParams, dataLoaded, failureCallback) => {
          dispatch(createAction(ActionNames.LOAD_DATA, Object.assign({}, loaderServiceParams, routeParams), {serviceName: loaderService, successCallback:  form.dataLoaded, failureCallback: failureCallback}));
        }
      }
      if(!this.formSubmit) {
        this.formSubmit = (data, successCallback, failureCallback) => {
          data = form.preSubmit(data)
          console.log("form submit submit form", data)
          if(form.config) {
            successCallback = form.config.submitSuccess? _reg('Method', form.config.submitSuccess) : successCallback
            failureCallback = form.config.submitFailure? _reg('Method', form.config.submitFailure) : failureCallback
          }
          dispatch(createAction(ActionNames.SUBMIT_FORM, data, {serviceName: form.config.submissionService, successCallback: successCallback, failureCallback: failureCallback}));
        }
      }
    }
  }
}
CustomForm.contextTypes = {
  uikit: PropTypes.object,
  getFormValue: PropTypes.func,
  routeParams: PropTypes.object
};

class WebFormUI extends React.Component {
  constructor(props) {
    super(props)
    let desc = props.description
    let config = props.config? props.config: desc.config
    config = config? config :{}
    if (config.entity && config.entityId){
      this.formType = EntityForm
    } else if(props.subform || desc.subform) {
      this.formType = SubEntityForm
    } else {
      this.formType = CustomForm
    }
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
  console.log("redux form state...........=======", state, ownProps)

  let formData = ownProps.formData
  if (desc.info && desc.info.preAssigned){
    formData = Object.assign({}, formData, desc.info.preAssigned)
  }


  let formVal  = state.form[ownProps.form]
  formVal = formVal? formVal: empty
  return { formVal, formData}
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return { dispatch }
}


const Form = connect(
  mapStateToProps,
  mapDispatchToProps
)(ReduxForm);


export {Form } ;

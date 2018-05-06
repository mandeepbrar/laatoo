'use strict';

import React from 'react';
import { connect } from 'react-redux';
import {ActionNames} from './Actions';
import {createAction} from 'uicommon';
import { reduxForm } from 'redux-form'
import {Field} from './Field';
const PropTypes = require('prop-types');

class WebFormUI extends React.Component {
  constructor(props, context) {
    super(props);
    let p = this.configureForm(props.dispatch, props)
    this.formSubmit = p.formSubmit
    this.uikit = context.uikit
    if(p.loader) {
      p.loader(context.routeParams, this.dataLoaded, this.failureCallback)
    }
    this.className = "webform " + ((p.config && p.config.className)? p.config.className :"")
    this.formName="myform"
    this.config = p.config? p.config :{}
    this.actions = props.actions
    let desc = props.description
    if(!this.actions && desc.info.actions) {
      this.actions = _reg('Methods', desc.info.actions)
    }
    console.log("webform ", props, context)
    this.state={formValue: this.getFormValue(props), time: Date.now()}
    let layoutFunc = null
    if(props.formData) {
      this.setData(props.formData)
    }
    this.trackChanges = this.config.trackChanges || props.trackChanges
    this.parentFormProps = {}
    if(props.subform || desc.subform) {
      this.parentFormValue = props.parentFormRef.getFormValue()
      console.log("received parent form value", this.parentFormValue)
      this.parentFormProps = {parentFormValue: this.parentFormValue}
    }
    let f = this.uiformSubmit(this.submitSuccessCallback, this.failureCallback)
    this.submitFunc = (customFunc) => { return customFunc? props.handleSubmit(customFunc): props.handleSubmit(f) }
  }

  configureForm(dispatch, ownProps) {
    var desc = ownProps.description
    var config = null
    var loader = null
    var formSubmit = null
    console.log("ownprops ", ownProps)
    if(desc) {
      config = ownProps.config? ownProps.config: desc.config
      if(config) {
        if(config.entity) {
          let entityId = config.entityId
          let entityName = config.entity
          let svc = config.entityService
          //console.log("desc....", entityId, "name", entityName, entityFormCfg)
          loader = (routeParams, dataLoaded, failureCallback) => {
            if(entityId) {
              dispatch(createAction(ActionNames.ENTITY_GET, { entityId, entityName}, {successCallback:  dataLoaded, failureCallback: failureCallback}));
            }
          }
          if(ownProps.onSubmit) {
            formSubmit = (data, successCallback, failureCallback) => {
              console.log("form submit", data)
              ownProps.onSubmit({data, entityId, entityName}, {reload: config.reloadOnUpdate, successCallback, failureCallback});
            }
          } else {
            if(entityId) {
              if(config.put) {
                formSubmit = (data, successCallback, failureCallback) => {
                  console.log("form submit put", data)
                  dispatch(createAction(ActionNames.ENTITY_PUT, {data, entityId, entityName}, {reload: config.reloadOnUpdate, successCallback, failureCallback}));
                }
              } else {
                formSubmit = (data, successCallback, failureCallback) => {
                  console.log("form submit update", data)
                  dispatch(createAction(ActionNames.ENTITY_UPDATE, {data, entityId, entityName}, {reload: config.reloadOnUpdate, successCallback, failureCallback}));
                }
              }
            } else {
              formSubmit = (data, successCallback, failureCallback) => {
                console.log("form submit save", data)
                dispatch(createAction(ActionNames.ENTITY_SAVE, {data, entityName}, {successCallback, failureCallback}));
              }
            }
          }
        } else {
          if(config.loaderService) {
            let loaderServiceParams = {}
            let loaderService = ""
            if(typeof(config.loaderService) == "string") {
              loaderService = config.loaderService
            } else {
              loaderService = config.loaderService.name
              loaderServiceParams = config.loaderService.params
            }
            if(loaderService) {
              loader = (routeParams, dataLoaded, failureCallback) => {
                dispatch(createAction(ActionNames.LOAD_DATA, Object.assign({}, loaderServiceParams, routeParams), {serviceName: loaderService, successCallback:  dataLoaded, failureCallback: failureCallback}));
              }
            }
          }
          if(ownProps.onSubmit) {
            formSubmit = (data, successCallback, failureCallback) => {
              console.log("data ", data, ownProps)
              ownProps.onSubmit(data, {successCallback, failureCallback});
            }
          } else {
            formSubmit = (data, successCallback, failureCallback) => {
              console.log("form submit submit form", data)
              if(config) {
                let preSubmit = _reg('Method', config.preSubmit)
                if(preSubmit) {
                  data = preSubmit(data)
                }
                successCallback = config.submitSuccess? _reg('Method', config.submitSuccess) : successCallback
                failureCallback = config.submitFailure? _reg('Method', config.submitFailure) : failureCallback
              }
              dispatch(createAction(ActionNames.SUBMIT_FORM, data, {serviceName: config.submissionService, successCallback: successCallback, failureCallback: failureCallback}));
            }
          }
        }
      }
    }

    return {loader, config, formSubmit }
  }



  componentWillReceiveProps(nextProps, nextState) {
    console.log("webform: next props of componentWillReceiveProps", this.props.form, nextProps, nextState)
    this.dataLoading = false
    let formValue = this.getFormValue(nextProps)
    let oldFormValue = this.state.formValue
    console.log("on change of form", formValue, oldFormValue)
    if(formValue != oldFormValue) {
      this.setState(Object.assign({}, this.state, {formValue, time: Date.now()}))
      if(this.props.onChange) {
        console.log("on change of form", formValue)
        this.props.onChange(formValue, oldFormValue)
      }
    }
  }

  reset = () => {
    this.props.reset()
  }

  subFormChange = (field, data) => {
    if(this.props.subform) {
      console.log("subform has changed ", field, data)
    }
  }

  getChildContext() {
    return {fields: this.props.description.fields, getFormValue: this.getFormValue, getParentFormValue: this.getParentFormValue};
  }

  failureCallback = () => {

  }

  submitSuccessCallback = (data) => {
    let cfg = this.props.config
    if(cfg) {
      if(cfg.successRedirect) {
        Window.redirect(cfg.successRedirect);
      }
      if(cfg.successRedirectPage) {
        Window.redirect(cfg.successRedirectPage);
      }
    }
  }

  setData = (formData) => {

    let x = this.props.initialize( formData)
    this.props.dispatch(x)
  }

  dataLoaded = (data) => {
    this.dataLoading = true
    let formData = Object.assign({}, data.resp.data)
    if(this.config.dataMapper) {
      let mapper = _reg('Method', this.config.dataMapper)
      formData = mapper(formData)
    }
    this.setData(formData)
  }

  getParentFormValue = () => {
    return this.parentFormValue
  }

  getFormValue = (props) => {
    if(!props) {
      props = this.props
    }
    /*console.log("get form value", this.props)
    if(!props) {
      props = this.props
    }
    let formVal  = props.state.form[props.form]
    return formVal? formVal.values: {}*/
    return props.formVal
  }

  layoutFields = (fldToDisp, flds, className) => {
    let fieldsArr = new Array()
    let comp = this
    console.log("layout fields =========", this.state, this.props)
    fldToDisp.forEach(function(k) {
      let fd = flds[k]
      let cl = className? className + " m10": "m10"
      fieldsArr.push(  <Field key={fd.name} name={fd.name} formValue={comp.state.formValue} formRef={comp} subFormChange={comp.subFormChange} subform={comp.props.subform}
        autoSubmitOnChange={comp.props.autoSubmitOnChange} parentFormRef={comp.props.parentFormRef} {...comp.parentFormProps} time={comp.state.time} className={cl}/>      )
    })
    return fieldsArr
  }

  fields = () => {
    let desc = this.props.description
    console.log("desc of form ", desc)
    let comp = this
    if(desc && desc.fields) {
      let flds = desc.fields
      if(flds) {
        if(desc.info && desc.info.tabs) {
          let tabs = new Array()
          let tabsToDisp = desc.info && desc.info.tabs? desc.info.layout: Object.keys(desc.info.tabs)
          tabsToDisp.forEach(function(k) {
            let tabFlds = desc.info.tabs[k];
            if(tabFlds) {
              let tabArr = comp.layoutFields(tabFlds, flds, "tabfield formfield")
              tabs.push(
                <comp.uikit.Tab label={k} time={comp.state.time} value={k}>
                  {tabArr}
                </comp.uikit.Tab>
              )
            }
          })
          let vertical = desc.info.verticaltabs? true: false
          return (
            <this.uikit.Tabset vertical={vertical} time={comp.state.time}>
              {tabs}
            </this.uikit.Tabset>
          )
        } else {
          let fldToDisp = desc.info && desc.info.layout? desc.info.layout: Object.keys(flds)
          let className=comp.props.inline?"inline formfield":"formfield"
          return this.layoutFields(fldToDisp, flds, className)
        }
      }
    }
    return null
  }

  uiformSubmit = (success, failure) => {
      var formSubmit = this.formSubmit
      return (data) => {
        formSubmit(data, success, failure)
      }
  }

  render() {
    console.log("**********************rendering web form****************", this.props.form, this.props)
    let cfg = this.config
    if(this.uikit.Form) {
      if(cfg.layout && (typeof(cfg.layout) == "string") ) {
        let display = _reg('Blocks', cfg.layout)
        if(display) {
          console.log("form context", this.props.formContext);
          let root = display(this.props.formContext, this.props.description, this.uikit)
          return React.cloneElement(root, { time: this.state.time, onSubmit: this.submitFunc, className: this.className})
        }
      } else {
        return (
          <this.uikit.Form time={this.state.time} onSubmit={this.submitFunc} className={this.className}>
            {this.fields()}
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

WebFormUI.childContextTypes = {
  fields: PropTypes.object,
  getParentFormValue: PropTypes.func,
  getFormValue: PropTypes.func
};

WebFormUI.contextTypes = {
  uikit: PropTypes.object,
  getFormValue: PropTypes.func,
  routeParams: PropTypes.object
};
/*
  <t.form.Form ref="form" key={state.key} type={state.schema} value={state.formValue} options={state.so} onChange={this.onChange}/>
  {this.actionButtons}*/

const ReduxForm = reduxForm({})(WebFormUI)
const empty = {}
const mapStateToProps = (state, ownProps) => {
  let desc = ownProps.description
  console.log("redux form state...........=======", state, ownProps)
  let formVal  = state.form[ownProps.form]
  formVal = formVal? formVal.values: empty

  return { formVal }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return { dispatch }
}


const Form = connect(
  mapStateToProps,
  mapDispatchToProps
)(ReduxForm);


export {Form } ;

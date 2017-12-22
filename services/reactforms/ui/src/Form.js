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
    this.uikit = context.uikit
    if(props.loader) {
      props.loader(context.routeParams, this.dataLoaded, this.failureCallback)
    }
    this.className = "webform " + ((props.config && props.config.className)? props.config.className :"")
    this.formName="myform"
    this.config = this.props.config? this.props.config :{}
    this.actions = props.actions
    let desc = props.description
    if(!this.actions && desc.actions) {
      this.actions = _reg('Method', desc.actions)
    }
    console.log("webform ", props)
    this.state={formValue: this.getFormValue(props), time: Date.now()}
    let layoutFunc = null
    if(props.formData) {
      this.setData(props.formData)
    }
    this.parentFormProps = {}
    if(props.subform || desc.subform) {
      let parentFormValue = props.parent.getFormValue()
      console.log("received parent form value", parentFormValue)
      this.parentFormProps = {parentFormValue}
    }
  }

  componentWillReceiveProps(nextProps, nextState) {
    if(this.config.dynamicFields) {
      let formValue = this.getFormValue(nextProps)
      if(formValue != this.state.formValue) {
        this.setState(Object.assign({}, this.state, {formValue, time: Date.now()}))
      }
    }
  }

  reset = () => {
    this.props.reset()
  }

  getChildContext() {
    return {fields: this.props.description.fields, getFormValue: this.getFormValue};
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
    let formData = data
    if(this.config.dataMapper) {
      let mapper = _reg('Method', this.config.dataMapper)
      formData = mapper(data)
    }
    this.setData(data.resp.data)
  }

  getFormValue = (props) => {
    if(!props) {
      props = this.props
    }
    let formVal  = props.state.form[props.form]
    return formVal? formVal.values: {}
  }

  layoutFields = (fldToDisp, flds, className) => {
    let fieldsArr = new Array()
    let comp = this
    fldToDisp.forEach(function(k) {
      let fd = flds[k]
      let cl = className? className + " m10": "m10"
      fieldsArr.push(  <Field key={fd.name} name={fd.name} formValue={comp.state.formValue} {...comp.parentFormProps} time={comp.state.time} className={cl}/>      )
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
      var formSubmit = this.props.formSubmit
      return (data) => {
        formSubmit(data, success, failure)
      }
  }

  render() {
    let {handleSubmit, actions} = this.props
    let f = this.uiformSubmit(this.submitSuccessCallback, this.failureCallback)
    let submitFunc = handleSubmit(f)
    let cfg = this.config
    if(this.uikit.Form) {
      if(cfg.layout && (typeof(cfg.layout) == "string") ) {
        let display = _reg('Blocks', cfg.layout)
        if(display) {
          console.log("form context", this.props.formContext);
          let root = display(this.props.formContext, this.props.description, this.uikit)
          return React.cloneElement(root, { time: this.state.time, onSubmit: submitFunc, className: this.className})
        }
      } else {
        return (
          <this.uikit.Form time={this.state.time} onSubmit={submitFunc} className={this.className}>
            {this.fields()}
            <this.uikit.Block className="actionbar p20 right">
              {
                this.actions?
                this.actions(this, submitFunc, this.reset):
                <this.uikit.ActionButton onClick={() => submitFunc()} className="submitBtn">{cfg.submit? cfg.submit: "Submit"}</this.uikit.ActionButton>
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
  getFormValue: PropTypes.func
};

WebFormUI.contextTypes = {
  uikit: PropTypes.object,
  routeParams: PropTypes.object
};
/*
  <t.form.Form ref="form" key={state.key} type={state.schema} value={state.formValue} options={state.so} onChange={this.onChange}/>
  {this.actionButtons}*/

const ReduxForm = reduxForm({})(WebFormUI)

const mapStateToProps = (state, ownProps) => {
  let desc = ownProps.description
  return { state }
}

const mapDispatchToProps = (dispatch, ownProps) => {
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
            console.log("form submit")
            ownProps.onSubmit({data, entityId, entityName}, {reload: config.reloadOnUpdate, successCallback, failureCallback});
          }
        } else {
          if(entityId) {
            if(config.put) {
              formSubmit = (data, successCallback, failureCallback) => {
                dispatch(createAction(ActionNames.ENTITY_PUT, {data, entityId, entityName}, {reload: config.reloadOnUpdate, successCallback, failureCallback}));
              }
            } else {
              formSubmit = (data, successCallback, failureCallback) => {
                dispatch(createAction(ActionNames.ENTITY_UPDATE, {data, entityId, entityName}, {reload: config.reloadOnUpdate, successCallback, failureCallback}));
              }
            }
          } else {
            formSubmit = (data, successCallback, failureCallback) => {
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
            ownProps.onSubmit({data}, {successCallback, failureCallback});
          }
        } else {
          formSubmit = (data, successCallback, failureCallback) => {
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


const Form = connect(
  mapStateToProps,
  mapDispatchToProps
)(ReduxForm);


export {Form } ;

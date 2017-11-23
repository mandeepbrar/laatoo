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

    let layoutFunc = null
    if(props.description.layout) {

    } else {

    }
  }

  getChildContext() {
    return {fields: this.props.description.fields};
  }

  failureCallback = () => {

  }

  submitSuccessCallback = (data) => {
    let cfg = this.props.config
    if(cfg) {
      console.log("submit callback", cfg);
      if(cfg.successRedirect) {
        Window.redirect(cfg.successRedirect);
      }
      if(cfg.successRedirectPage) {
        Window.redirect(cfg.successRedirectPage);
      }
    }
  }

  dataLoaded = (data) => {
    let formData = data
    let cfg = this.props.config
    if(cfg & cfg.dataMapper) {
      let mapper = _reg('Method', cfg.dataMapper)
      formData = mapper(data)
    }
    let x = this.props.initialize(data.resp.data, 'myform')
    this.props.dispatch(x)
  }

  fields = () => {
    let fieldsArr = new Array()
    let desc = this.props.description
    console.log("desc of form ", desc)
    if(desc && desc.fields) {
      let flds = desc.fields
      let fldToDisp = desc.info && desc.info.fieldsLayout? desc.info.fieldsLayout: Object.keys(flds)
      console.log("fldToDisp", fldToDisp)
      if(flds) {
        fldToDisp.forEach(function(k) {
          let fd = flds[k]
          fieldsArr.push(  <Field name={fd.name}/>      )
        })
      }
    }
    return fieldsArr
  }

  uiformSubmit = (success, failure) => {
      var formSubmit = this.props.formSubmit
      return (data) => {
        formSubmit(data, success, failure)
      }
  }

  render() {
    let {handleSubmit} = this.props
    let f = this.uiformSubmit(this.submitSuccessCallback, this.failureCallback)
    let cfg = this.props.config? this.props.config :{}
    console.log('render form', this.props)
    if(this.uikit.Form) {
      if(cfg.layout) {
        let display = _reg('Blocks', cfg.layout)
        if(display) {
          console.log("form context", this.props.formContext);
          let root = display(this.props.formContext, this.props.description, this.uikit)
          return React.cloneElement(root, {onSubmit: handleSubmit(f), className: this.className})
        }
      } else {
        return (
          <this.uikit.Form onSubmit={handleSubmit(f)} className={this.className}>
            {this.fields()}
            <this.uikit.Block className="actionbar">
              <button type="submit" className="submitBtn">{cfg.submit? cfg.submit: "Submit"}</button>
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
  fields: PropTypes.object
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
  return {  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  var desc = ownProps.description
  var config = null
  var loader = null
  var formSubmit = null


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

  return {loader, config, formSubmit }
}


const Form = connect(
  mapStateToProps,
  mapDispatchToProps
)(ReduxForm);


export {Form } ;

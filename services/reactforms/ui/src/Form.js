'use strict';

import React from 'react';
import { connect } from 'react-redux';
import {ActionNames} from './Actions';
import {createAction} from 'uicommon';
import { Field, reduxForm } from 'redux-form'
const PropTypes = require('prop-types');

class WebFormUI extends React.Component {
  constructor(props, context) {
    super(props);
    this.uikit = context.uikit
    props.loader(context.routeParams, this.dataLoaded, this.failureCallback)
    ///const { handleSubmit } = props
  /*  this.submitForm = this.submitForm.bind(this);
    this.setValue = this.setValue.bind(this);
    this.submit = this.submit.bind(this);
    this.getValue = this.getValue.bind(this);
    this.onChange = this.onChange.bind(this);
    if(this.props.actionButtons) {
      this.actionButtons = this.props.actionButtons
    } else {

    }
    this.lookupSchemaOptions = desc.lookupSchemaOptions
    let so = desc.lookupSchemaOptions? desc.lookupSchemaOptions(this) : desc.schemaOptions

    this.state = {formValue: props.entityData, so : so, key: "entityform" + (new Date())}
    if(props.refCallback) {
      props.refCallback(this)
    }*/
  }

  failureCallback = () => {

  }

  submitSuccessCallback = (data) => {
    var redirect = this.config && this.config.successRedirect ? this.config.successRedirect: null
    if(redirect) {
      Window.redirect(redirect);
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
    let f = this.props.description.fields
    let uikit = this.uikit
    if(uikit.Forms.FieldWidget && f) {
      Object.keys(f).forEach(function(k) {
        let fd = f[k]
        //let component = uikit.Forms.GetFieldWidget(fd)
        fieldsArr.push(
         <Field name={fd.name} field={fd} component={uikit.Forms.FieldWidget}/>
        )
      })
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
    if(this.uikit.Form) {
      return (
        <this.uikit.Form onSubmit={handleSubmit(f)} className={"webform " + ((cfg && cfg.className)? cfg.className :"")}>
        {this.fields()}
        <button type="submit">{cfg.submit? cfg.submit: "Submit"}</button>
        </this.uikit.Form>
      )
    } else {
      return <this.uikit.Block/>
    }
  }
}

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
  //var
  /*if(ownProps.onSubmit) {
      onSubmit = ownProps.onSubmit
  }
  else {
    let formDesc = ownProps.description
    if(formDesc.entityName) {
      switch(formDesc.formType) {
        case "Create":
          onSubmit = (data) => {
            dispatch(createAction(ActionNames.ENTITY_SAVE, {data:data, entityName: ownProps.name}, {reducer: ownProps.reducer, successCallback: ownProps.postSave, failureCallback: ownProps.failureCallback}));
          }
          break;
        case "Put":
          onSubmit = (data) => {
            dispatch(createAction(ActionNames.ENTITY_PUT, {data:data, entityId: ownProps.id, entityName: ownProps.name}, {reducer: ownProps.reducer, reload: ownProps.reloadOnUpdate, successCallback: ownProps.postSave, failureCallback: ownProps.failureCallback}));
          }
          break;
        case "Update":
          onSubmit = (data) => {
            dispatch(createAction(ActionNames.ENTITY_UPDATE, {data:data, entityId: ownProps.id, entityName: ownProps.name}, {reducer: ownProps.reducer, reload: ownProps.reloadOnUpdate, successCallback: ownProps.postSave, failureCallback: ownProps.failureCallback}));
          }
          break;
      }
    } else {
      let svc = formDesc.serviceName
      onSubmit = (data) => {
        dispatch(createAction(ActionNames.SUBMIT_FORM, data, {successCallback: ownProps.postSave, failureCallback: ownProps.failureCallback}));
      }
    }
  }*/
  return {loader, config, formSubmit }
}


const Form = connect(
  mapStateToProps,
  mapDispatchToProps
)(ReduxForm);



export {Form } ;

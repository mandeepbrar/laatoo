import {Response,  DisplayEntity, EntityData, getCreateReducer, getViewReducer, getUpdateReducer, getDisplayReducer, ActionNames} from 'laatoocommon';
import {EntityView} from './EntityWebView';
import {CreateEntity} from './EntityCreate';
import {UpdateEntity} from './EntityUpdate';
import React from 'react';
import {Action} from '../action/Action';

class Entity {
    constructor(name, entityProperties) {
      this.name = name;
      this.entityProperties = entityProperties;
      this.ViewComponent = this.ViewComponent.bind(this)
      this.CreateComponent = this.CreateComponent.bind(this)
      this.UpdateComponent = this.UpdateComponent.bind(this)
      this.ViewReducer = this.ViewReducer.bind(this)
      this.CreateReducer = this.CreateReducer.bind(this)
      this.UpdateReducer = this.UpdateReducer.bind(this)
      this.DisplayReducer = this.DisplayReducer.bind(this)
      this.DisplayComponent = this.DisplayComponent.bind(this)
      //schema, schemaoptions,
    }

    ViewComponent() {
      let reducer = this.name.toUpperCase()+"_View";
      let idField = this.entityProperties.idField;
      let titleField = this.entityProperties.titleField;
      let actions = this.entityProperties.actions
      let header = this.entityProperties.viewHeader;
      let paginate = this.entityProperties.paginate;
      let pageSize = this.entityProperties.pageSize;
      let loader = this.entityProperties.loader;
      let row = this.entityProperties.viewRow;
      let currentPage = 1;
      let viewService = this.entityProperties.viewService;
      let urlParams = this.entityProperties.urlParams;
      let viewArgs = this.entityProperties.viewArgs;
      let filterForm = this.entityProperties.filterForm;
      let defaultFilter = this.entityProperties.defaultFilter;
      return (props) => {
          let params = null
          if(props && props.params) {
            params=  props.params
          }
          return (
            <EntityView key={reducer} name={this.name} filterForm={filterForm} idField={idField} paginate={paginate} loader={loader}
              pageSize={pageSize} getHeader={header} getItem={row} reducer={reducer} postArgs={viewArgs} defaultFilter={defaultFilter}
              actions={actions} titleField={titleField} params={params} viewService={viewService} urlParams={urlParams} currentPage={currentPage}>
            </EntityView>
          )
        }
    }
    CreateComponent() {
      let reducer = this.name.toUpperCase()+"_Form";
      let schema = this.entityProperties.schema;
      let data = this.entityProperties.data;
      let mountForm = this.entityProperties.mountForm;
      let postSave = this.entityProperties.postSave;
      let preSave = this.entityProperties.preSave;
      let refCallback = this.entityProperties.refCallback;
      let schemaOptions = this.entityProperties.schemaOptions;
      return (props) =>  {
        let idToDuplicate = null
        let params = null
        if(props && props.params && props.params.idToDuplicate) {
          idToDuplicate = props.params.idToDuplicate
        }
        if(props && props.params) {
          params = props.params
        }
        return (
            <CreateEntity name={this.name} idToDuplicate={idToDuplicate} data={data} reducer={reducer} refCallback={refCallback} schema={schema}
             params={params} mountForm={mountForm} postSave={postSave} preSave={preSave} schemaOptions={schemaOptions}></CreateEntity>
          )
      }
    }
    UpdateComponent() {
      let reducer = this.name.toUpperCase()+"_Form";
      let schema = this.entityProperties.schema;
      let mountForm = this.entityProperties.mountForm;
      let postSave = this.entityProperties.postSave;
      let preSave = this.entityProperties.preSave;
      let usePut = this.entityProperties.usePut;
      let refCallback = this.entityProperties.refCallback;
      let schemaOptions = this.entityProperties.schemaOptions;
      return (props) => (
          <UpdateEntity name={this.name} id={props.params.id} refCallback={refCallback} usePut={usePut} reducer={reducer} schema={schema} mountForm={mountForm}
            postSave={postSave} preSave={preSave} schemaOptions={schemaOptions}></UpdateEntity>
        )
    }
    DisplayComponent() {
      let reducer = this.name.toUpperCase()+"_Display";
      let display = this.entityProperties.display;
      let loader = this.entityProperties.loader;
      return (props) =>  {
        let params = null
        if(props && props.params) {
          params=  props.params
        }
        return (
          <DisplayEntity name={this.name} id={props.params.id} loader={loader} params={params} reducer={reducer} display={display}></DisplayEntity>
        )
      }
    }
    ViewReducer() {
      return getViewReducer(this.name);
    }
    CreateReducer() {
      return getCreateReducer(this.name);
    }
    UpdateReducer() {
      return getUpdateReducer(this.name);
    }
    DisplayReducer() {
      return getDisplayReducer(this.name);
    }
}

export {Entity as Entity};

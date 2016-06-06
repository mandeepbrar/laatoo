import {getCreateReducer, getViewReducer, getUpdateReducer} from './Reducers';
import {EntityView} from './EntityWebView';
import {CreateEntity} from './EntityCreate';
import {UpdateEntity} from './EntityUpdate';
import React from 'react';

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
      //schema, schemaoptions,
    }
    ViewComponent() {
      let reducer = this.name.toUpperCase()+"_View";
      let idField = this.entityProperties.idField;
      let titleField = this.entityProperties.titleField;
      let header = this.entityProperties.viewHeader;
      let row = this.entityProperties.viewRow;
      let viewService = this.entityProperties.viewService;
      return () => (
          <EntityView name={this.name} idField={idField} header={header} row={row} reducer={reducer} titleField={titleField} viewService={viewService} ></EntityView>
        )
    }
    CreateComponent() {
      let reducer = this.name.toUpperCase()+"_Form";
      let schema = this.entityProperties.schema;
      let mountForm = this.entityProperties.mountForm;
      let postSave = this.entityProperties.postSave;
      let preSave = this.entityProperties.preSave;
      let schemaOptions = this.entityProperties.schemaOptions;
      return () => (
          <CreateEntity name={this.name} reducer={reducer} schema={schema} mountForm={mountForm} postSave={postSave} preSave={preSave} schemaOptions={schemaOptions}></CreateEntity>
        )

    }
    UpdateComponent() {
      let reducer = this.name.toUpperCase()+"_Form";
      let schema = this.entityProperties.schema;
      let mountForm = this.entityProperties.mountForm;
      let postSave = this.entityProperties.postSave;
      let preSave = this.entityProperties.preSave;
      let schemaOptions = this.entityProperties.schemaOptions;
      return (props) => (
          <UpdateEntity name={this.name} id={props.params.id} reducer={reducer} schema={schema} mountForm={mountForm} postSave={postSave} preSave={preSave} schemaOptions={schemaOptions}></UpdateEntity>
        )
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
}

export {Entity as Entity};

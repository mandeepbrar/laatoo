'use strict';

import React from 'react';
import { createAction } from 'uicommon';
import {ActionNames} from '../Actions';
import { connect } from 'react-redux';

class ViewEntity extends React.Component {
  constructor(props) {
    super(props);
  }
  componentDidMount() {
    if(this.props.load && !this.props.externalLoad) {
      this.props.loadEntity();
    }
  }
  componentWillReceiveProps(nextprops) {
    if(nextprops.load) {
      this.props.loadEntity();
    }
  }
  shouldComponentUpdate(nextProps, nextState) {
    console.log("shoudl component update", nextProps, nextState)
    if(!nextProps.forceUpdate && this.lastRenderTime) {
      if(nextProps.lastUpdateTime) {
        if(this.lastRenderTime >= nextProps.lastUpdateTime) {
          console.log("update false", nextProps.name)
          return false
        }
      } else {
        console.log("update false", nextProps.name)
        return false
      }
    }
    console.log("update true", nextProps.name)
    return true;
  }
  render() {
    console.log("render",this.props.name)
    let display = null
    this.lastRenderTime = this.props.lastUpdateTime
    if(this.props.display && this.props.status && this.props.status == "Loading") {
      display = this.props.loader
    } else {
      display = this.props.display(this.props.data, this.props.desc, this.props.uikit, this.props.lastUpdateTime)
    }
    return display
  }
}

const mapStateToProps = (state, ownProps) => {
  console.log("map state for entity", ownProps, state)
  let props = {
    name: ownProps.name,
    id: ownProps.id,
    desc: ownProps.desc,
    uikit: ownProps.uikit,
    params: ownProps.params,
    loader: ownProps.loader,
    reducer: ownProps.reducer,
    forceUpdate: ownProps.forceUpdate,
    externalLoad: ownProps.externalLoad,
    display: ownProps.display,
    load: false
  };
  let entityViewReducer = state["entityview"];
/*  if(!ownProps.globalReducer) {
    if(state.router && state.router.routeStore) {
      entity = state.router.routeStore[ownProps.reducer];
    }
  } else {
    entity = state[ownProps.reducer];
  }*/
  console.log("entityViewReducer", entityViewReducer, ownProps.id)
  if(entityViewReducer && ownProps.id) {
    let entity = entityViewReducer.entities[ownProps.id]
    if(entity) {
      console.log("entity ", entity, ownProps.id)
      props.status = entity.status
      props.data = entity.data
      if(entity.status == "Loaded") {
        props.lastUpdateTime = entity.lastUpdateTime
      }
      if(entity.status == "NotLoaded") {
          props.load = true
      }
    } else {
      props.load = true
    }
  }
  return props;
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    loadEntity: () => {
      let payload = {entityName: ownProps.name, entityId: ownProps.id, headers: ownProps.headers, svc: ownProps.svc};
      let meta = { global: ownProps.global};
      dispatch(createAction(ActionNames.ENTITY_VIEW_FETCH, payload, meta));
    }
  }
}


const Entity = connect(
  mapStateToProps,
  mapDispatchToProps
)(ViewEntity);

export {Entity as Entity}

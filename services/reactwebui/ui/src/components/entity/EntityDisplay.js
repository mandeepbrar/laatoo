'use strict';
/*
import React from 'react';
import {  Response,  DataSource,  RequestBuilder, createAction, ActionNames } from 'laatoouibase';
import { connect } from 'react-redux';

class Display extends React.Component {
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
      display = this.props.display(this.props.data)
    }
    return (
      <div>
        {display}
      </div>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  let props = {
    name: ownProps.name,
    id: ownProps.id,
    params: ownProps.params,
    loader: ownProps.loader,
    reducer: ownProps.reducer,
    forceUpdate: ownProps.forceUpdate,
    externalLoad: ownProps.externalLoad,
    display: ownProps.display,
    load: false
  };
  let entity = null;
  if(!ownProps.globalReducer) {
    if(state.router && state.router.routeStore) {
      entity = state.router.routeStore[ownProps.reducer];
    }
  } else {
    entity = state[ownProps.reducer];
  }
  if(entity) {
    props.status = entity.status
    props.data = entity.data
    if(entity.status == "Loaded") {
      props.lastUpdateTime = entity.lastUpdateTime
    }
    if(entity.status == "NotLoaded") {
        props.load = true
    }
  }
  return props;
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    loadEntity: () => {
      let payload = {entityName: ownProps.name, entityId: ownProps.id, headers: ownProps.headers, svc: ownProps.svc};
      let meta = {reducer: ownProps.reducer};
      dispatch(createAction(ActionNames.ENTITY_GET, payload, meta));
    }
  }
}

const DisplayEntity = connect(
  mapStateToProps,
  mapDispatchToProps
)(Display);

export {DisplayEntity as DisplayEntity}
*/

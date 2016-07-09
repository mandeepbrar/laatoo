'use strict';

import React from 'react';
import {  Response,  DataSource,  RequestBuilder } from '../../sources/DataSource';
import { connect } from 'react-redux';
import { createAction } from '../../utils';
import  {ActionNames} from '../../actions/ActionNames';

class Display extends React.Component {
  constructor(props) {
    super(props);
  }
  componentDidMount() {
    if(this.props.load) {
      this.props.loadEntity();
    }
  }
  componentWillReceiveProps(nextprops) {
    if(nextprops.load) {
      this.props.loadEntity();
    }
  }
  shouldComponentUpdate(nextProps, nextState) {
    if(this.lastRenderTime) {
      if(nextProps.lastUpdateTime) {
        if(this.lastRenderTime >= nextProps.lastUpdateTime) {
          return false
        }
      } else {
        return false
      }
    }
    return true;
  }
  render() {
    let display = null
    this.lastRenderTime = this.props.lastUpdateTime
    if(this.props.display && this.props.status && this.props.status == "Loaded") {
      display = this.props.display(this.props.data)
    } else {
      display = this.props.loader
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
    loader: ownProps.loader,
    reducer: ownProps.reducer,
    display: ownProps.display,
    load: false
  };
  if(state.router && state.router.routeStore) {
    let entity = state.router.routeStore[ownProps.reducer];
    if(entity) {
      props.status = entity.status
      props.data = entity.data
      props.lastUpdateTime = entity.lastUpdateTime
      if(entity.status == "NotLoaded") {
          props.load = true
      }
    }
  }
  return props;
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    loadEntity: () => {
      let payload = {entityName: ownProps.name, entityId: ownProps.id};
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

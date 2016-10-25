'use strict';

import React from 'react';
import { connect } from 'react-redux';
import { createAction } from '../../utils';
import  {ActionNames} from '../../actions/ActionNames';

class GroupLoadView extends React.Component {
  componentDidMount() {
    this.props.loadGroup();
  }
  render() {
    return null
  }
}

const mapStateToProps = (state, ownProps) => {
  return {}
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    loadGroup: () => {
      console.log("load group", ownProps)
      let payload = ownProps.Data //{entityName: ownProps.name, entityId: ownProps.id};
      let meta = {serviceName: ownProps.service};
      dispatch(createAction(ActionNames.GROUP_LOAD, payload, meta));
    }
  }
}

const GroupLoad = connect(
  mapStateToProps,
  mapDispatchToProps
)(GroupLoadView);

export {GroupLoad as GroupLoad}

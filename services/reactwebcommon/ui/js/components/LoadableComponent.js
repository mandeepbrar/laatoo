import React from 'react';
import {RequestBuilder, DataSource, EntityData} from 'uicommon';
import PropTypes from 'prop-types';

class LoadableComponent extends React.Component {
  constructor(props) {
    super(props)
    if(props.loader) {
      this.method = _reg("Methods", props.loader)
    }
  }

  componentWillMount() {
    console.log("loadable component:---------", this.method, this.props)
    let props = this.props
    if(this.method) {
      this.method(props, this.getLoadContext? this.getLoadContext(): {}, this.dataLoaded)
    } else if(!props.skipDataLoad && props.dataService) {
      let req = RequestBuilder.DefaultRequest(null, props.dataServiceParams);
      DataSource.ExecuteService(props.dataService, req).then(this.response, this.errorMethod);
    } else if(!props.skipDataLoad && props.entity) {
      EntityData.ListEntities(props.entity).then(this.response, this.errorMethod);
    }
  }

  errorMethod = (resp) => {
    console.log("could not load data", resp)
  };

  response = (resp) => {
    console.log("loadable component:---------response", resp)
    if(resp && resp.data) {
      this.dataLoaded(resp.data)
    }
  }
}

export {LoadableComponent as LoadableComponent}

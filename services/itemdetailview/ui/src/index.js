import React from 'react';
import {Panel} from 'reactpages';
import {View} from 'laatooviews';
import {Action} from 'reactwebcommon';
const PropTypes = require('prop-types');

class ItemDetailView extends React.Component {
  constructor(props) {
    super(props)
    this.state={open:false, item: null}
  }
  openDetail = (item) => {
    console.log("open detail", item)
    this.setState(Object.assign({}, this.state, {open: true, item: item}))
  }
  closeDetail = () => {
    this.setState(Object.assign({}, this.state, {open: false}))
  }
  getItem = (view, x, i) => {
    let item = view.getRenderedItem(x,i)
    console.log("getitem", item, this.context.uikit.Action)
    let action = <Action action={{actiontype:"method", method: this.openDetail, params:{data: x, index: i}}}>{item}</Action>
    console.log("created actio", action, item, x, i)
    return action
  }
  render() {
    console.log("rendering itemdetail view")
    let props = this.props
    let ctx = this.context
    return (
      <ctx.uikit.Block className="itemdetailview">
        <ctx.uikit.Block className="view">
          {props.id?
            <Panel className={props.className} getItem={this.getItem} type="view" id={props.id}></Panel>
            :
            <View service={props.service} serviceName={props.serviceName} name={props.viewname} global={props.global}
              className={props.className} incrementalLoad={props.incrementalLoad} paginate={props.paginate} header={props.header}
              getHeader={props.getHeader} getView={props.getView} getItem={this.getItem} urlparams={props.urlparams} postArgs={props.postArgs}>
            </View>
          }
        </ctx.uikit.Block>
        {
          this.state.open?(
            props.entityName?
            <Panel className="detail" description={{type:"entity", entityName: props.entityName, data: this.state.item}}/>
            :null
            )
          :null
        }
      </ctx.uikit.Block>
    )
  }
}
/*<ViewComponent serviceObject={view.service} serviceName={view.serviceName} name={viewname} global={view.global}
  className={className} incrementalLoad={view.incrementalLoad} paginate={view.paginate} header={props.header} getHeader={props.getHeader}
   getView={props.getView} getItem={props.getItem} urlparams={params} postArgs={args}>
   {item}
   </ViewComponent>*/
ItemDetailView.contextTypes = {
  uikit: PropTypes.object
};

export {
  ItemDetailView
}

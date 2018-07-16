import React from 'react';
import {Panel} from 'reactpages';
import {View} from 'laatooviews';
import {Action} from 'reactwebcommon';
const PropTypes = require('prop-types');
import './styles/app.scss'

class ItemDetailView extends React.Component {
  constructor(props) {
    super(props)
    console.log("creating item detail view", props)
    this.state={open:false, item: null}
  }
  openDetail = (item) => {
    console.log("open detail", item)
    let entity = item.data
    //Window.showDialog(entity.Id, <Panel className="detail" description={{type:"entity", entityName: this.props.entityName, data: entity}}/>);
    this.setState(Object.assign({}, this.state, {open: true, item: item.data}))
  }
  hideDetail = () => {
    this.setState(Object.assign({}, this.state, {open:false, item: null}))
  }
  getItem = (view, x, i) => {
    let item = view.getRenderedItem(x,i)
    if(this.props.getItem) {
      view.addMethod('openDetail', this.openDetail)
      view.addMethod('hideDetail', this.hideDetail)
      return this.props.getItem(view, x, i)
    } else {
      let action = <Action action={{actiontype:"method", method: this.openDetail, params:{data: x, index: i}}}>{item}</Action>
      return action
    }
  }
  render() {
    console.log("rendering item detail view", this.props, this.context)
    let props = this.props
    let ctx = this.context
    let itemDetailViewClass = " row "
    let detailPanelClass = " col-xs-6 "
    let viewClass = this.state.open? " col-xs-6 ": ""
    return (
      <ctx.uikit.Block className={" itemdetailview ma10 " + itemDetailViewClass}>
        <ctx.uikit.Block className={" view " + viewClass}>
          {props.id?
            <Panel className={props.className} editable={props.editable} getItem={this.getItem} viewRef={props.viewRef} type="view" id={props.id}></Panel>
            :
            <View service={props.service} serviceName={props.serviceName} name={props.viewname} global={props.global}  editable={props.editable}
              className={props.className} incrementalLoad={props.incrementalLoad} paginate={props.paginate} header={props.header} viewRef={props.viewRef}
              getHeader={props.getHeader} getView={props.getView} getItem={this.getItem} urlparams={props.urlparams} postArgs={props.postArgs}>
            </View>
          }
        </ctx.uikit.Block>
        <ctx.uikit.Block className={detailPanelClass}>
        {
          this.state.open?(
            props.entityName?
            <Panel className="detail" title={this.state.item.Name} closePanel={this.hideDetail} description={{type:"entity", entityName: props.entityName, data: this.state.item}}/>
            :null
            )
          :null
        }
        </ctx.uikit.Block>
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

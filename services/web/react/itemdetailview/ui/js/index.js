import React from 'react';
import {Panel} from 'reactpages';
import {View} from 'laatooviews';
import {Action} from 'reactwebcommon';
import {  Response,  DataSource,  RequestBuilder } from 'uicommon';
const PropTypes = require('prop-types');
import './styles/app.scss'
 
class ItemDetailView extends React.Component {
  constructor(props) {
    super(props)
    console.log("creating item detail view", props)
    this.state={open:false, item: null}
    this.view = props.viewRef? props.viewRef: React.createRef()

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
      if(this.props.form) {
        var methods = view.methods;
        console.log("methods ", methods)
        let select = ()=> {
          console.log("selected ", i)
          methods.itemSelectionChange(i, true)
        }
        let uikit = this.context.uikit;
        return <uikit.Block>
          <uikit.Block className="row center valigncenter ma10">
          {x.Name}
          </uikit.Block>
          <uikit.Block className="row m10">
            <Action className="p10" action={{actiontype:"method", method: select, params:{}}}>Select</Action>
            <Action className="p10" action={{actiontype:"method", method: this.openDetail, params:{data: x, index: i}}}>Details</Action>
          </uikit.Block>
        </uikit.Block>  
      } else { 
        return <Action action={{actiontype:"method", method: this.openDetail, params:{data: x, index: i}}}>{item}</Action>
      }
    }
  }

  submit = () => {
    let retData = this.view.current.selectedItems()
    let data;
    console.log("selected items", retData)
    if(Array.isArray(retData)) {
      if(this.props.postIds) {
        data = retData.forEach((item)=> {
          return item.Id
        })
      } else {
        data = retData
      }
    } else {
      if(this.props.postIds) {
        data = retData.Id
      } else {
        data = retData
      }
    }
    console.log("post data", data)
    if(this.props.submitService) {
      let req = RequestBuilder.DefaultRequest(null, data)
      let prom = DataSource.ExecuteService(this.props.submitService, req, {})
      let comp = this
      prom.then(
        function (res) {
          if(comp.props.successCallback) {
            comp.props.successCallback(res)
          }
        },
        function (res) {
          if(comp.props.failureCallback) {
            comp.props.failureCallback(res)
          }
          console.log(res);
        });
    } else {
      if(this.props.add) {
        this.props.add(data, null, null, true)
      }
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
            <Panel className={props.className} editable={props.editable} getItem={this.getItem} viewRef={this.view} type="view" id={props.id}></Panel>
            :
            <View service={props.service} serviceName={props.serviceName} name={props.viewname} global={props.global}  editable={props.editable} singleSelection={props.singleSelection}
              className={props.className} incrementalLoad={props.incrementalLoad} paginate={props.paginate} header={props.header} viewRef={this.view}
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
        {
          props.form?
          <ctx.uikit.Block className=" w100 right ">
            <Action widget="button" className="p10" action={{actiontype:"method", method: this.submit, params:{}}}>Submit</Action>
          </ctx.uikit.Block>
          : null
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

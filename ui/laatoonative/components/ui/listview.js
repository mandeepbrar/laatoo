'use strict';

import React from 'react'
import {Text, View, FlatList} from 'react-native'
import {ViewData} from 'laatoocommon'

class ListView extends React.Component {
  constructor(props) {
    super(props)
    this.getView=this.getView.bind(this);
    this.renderView=this.renderView.bind(this);
    this.renderItem=this.renderItem.bind(this);
    this.getItem=this.getItem.bind(this);
    this.getHeader=this.getHeader.bind(this);
    this.getPagination=this.getPagination.bind(this);
    this.getFilter=this.getFilter.bind(this);
    this.onScrollEnd=this.onScrollEnd.bind(this);
    this.addMethod=this.addMethod.bind(this);
    this.style = Object.assign({}, {flex:1}, this.props.style)
    this.contentStyle = Object.assign({}, {flex:1}, this.props.contentStyle)
    this.numItems=0;
  }

  addMethod(name,method){
    return this.viewdata.addMethod(name,method);
  }
  methods(){
    return this.viewdata.methods;
  }

  onScrollEnd() {
    var methods=this.methods();
    methods.loadMore();
  }

  getView(header,groups,pagination,filter){
    console.log("getting the view ", this.style, this.contentStyle);
    if(this.props.getView){
      return this.props.getView(this,header,groups,pagination,filter);
    }
    return(
      <View style={this.style}>
        <View style={this.props.headerStyle}>
          {header}
        </View>
        {filter?filter:null}
        <View style={this.contentStyle}>
          {groups}
        </View>
        {pagination?pagination:null}
      </View>
    )
  }
  getFilter(filterTitle,filterForm,filterGo){
    if(this.props.getFilter){
      return this.props.getFilter(this,filterTitle,filterForm,filterGo);
    }
    return null;
  }
  renderItem(itemInfo){
    console.log("rendering item*****",itemInfo);
    return this.getItem(itemInfo.item,itemInfo.index);
  }
  getItem(x,i){
    if(this.props.getItem){
      return this.props.getItem(this,x,i);
    }
    let child = React.Children.only(this.props.children);
    console.log("child", child)
    return React.cloneElement(child, {item:x,index:i} );
  }
  getHeader() {
    if(this.props.getHeader){
      return this.props.getHeader(this);
    }
    return null;
  }
  getPagination() {
    if(this.props.paginate&&this.props.getPagination){
      var pages=this.props.totalPages;
      var page=this.props.currentPage;
      return this.props.getPagination(this,pages,page);
    }
    return null;
  }
  renderView(viewdata,items,currentPage,totalPages) {
    this.viewdata=viewdata;
    var body=[];
    if(items){
      if(this.props.incrementalLoad){
        body.push(<FlatList data={items} onEndReached={this.onScrollEnd} renderItem={this.renderItem} horizontal={this.props.horizontal} numColumns={this.props.numColumns} columnWrapperStyle={this.props.columnWrapperStyle} onPress={this.props.onPress}/>);
      } else{
        body.push(<FlatList data={items} horizontal={this.props.horizontal} numColumns={this.props.numColumns}  renderItem={this.renderItem}/>);
      }
    } else{
      if(this.props.loader){
        body.push(this.props.loader);
      }
    }
    var header=this.getHeader();
    var filterCtrl=this.getFilter(this.props.filterTitle,this.props.filterForm,this.props.filterGo,this.filter);
    var pagination=this.getPagination();
    return this.getView(header,body,pagination,filterCtrl);
  }

  render() {
    return(
      <ViewData
        getView={this.renderView}
        key={this.props.key}
        reducer={this.props.reducer}
        paginate={this.props.paginate}
        pageSize={this.props.pageSize}
        viewService={this.props.viewService}
        urlParams={this.props.urlParams}
        postArgs={this.props.postArgs}
        defaultFilter={this.props.defaultFilter}
        currentPage={this.props.currentPage}
        style={this.props.style}
        className={this.props.className}
        incrementalLoad={this.props.incrementalLoad}
        globalReducer={this.props.globalReducer}   />
    );
  }
}

export {
  ListView as ListView
};

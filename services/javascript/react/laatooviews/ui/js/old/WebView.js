/*'use strict';

import React from 'react';
import {View} from 'reactuibase';
import {ViewFilter} from './Filter';
const PropTypes = require('prop-types');*/
/*
import Paginator from 'react-pagify';
import pagifyBootstrapPreset from 'react-pagify-preset-bootstrap';
import segmentize from 'segmentize';
*/
/*class WebView extends React.Component {
  constructor(props, context) {
    super(props);
    //this.getPagination = this.getPagination.bind(this);
    this.getView = this.getView.bind(this);
    this.getItemGroup = this.getItemGroup.bind(this);
    this.getItem = this.getItem.bind(this);
    this.getHeader = this.getHeader.bind(this);
    this.getFilter = this.getFilter.bind(this);
    this.addMethod = this.addMethod.bind(this);
    this.methods = this.methods.bind(this);
    this.uikit = context.uikit;
    this.div = this.uikit.Block;
    this.onScrollEnd = this.onScrollEnd.bind(this);
  }

  addMethod(name, method) {
    return this.refs.view.addMethod(name, method)
  }
  methods() {
    return this.refs.view.methods
  }*/ /*
  getFilter(view, filterTitle, filterForm, filterGo, filter) {
    if(this.props.getFilter) {
      return this.props.getFilter(filterTitle, filterForm, filterGo, filter)
    }
    if(!filterForm) {
      return null
    }
    if(!filterTitle) {
      filterTitle="Search"
    }
    if(!filterGo) {
      filterGo = "Go"
    }
    return (
      <ViewFilter title={filterTitle} schema={filterForm} defaultFilter={filter} setFilter={view.methods.setFilter} goBtn={filterGo} >
        <div className="row m20">
          <i className="fa fa-search pull-right"></i>
        </div>
      </ViewFilter>
    )
  }*/
/*
  onItemCheckboxChange(view){
    return (evt)=> {
      let cb = evt.target
      let item = view.refs[cb.value]
      if (cb.checked) {
        item.selected = true
      } else {
        item.selected = false
      }
    }
  }


  onScrollEnd() {
    let methods = this.methods();
    methods.loadMore();
  }*/

/*
  getView(view, header, groups, pagination, filter) {
    let viewComp = null
    if(this.props.getView) {
        viewComp = this.props.getView(view, header, groups)
    }
    return (
      <div key={this.props.key} className={this.props.className} style={this.props.style}>
        {filter}
        {viewComp}
        {pagination}
      </div>
    )
  }
*//*
  getView(view, header, groups, pagination, filter) {
    if(this.props.editable) {
      view.addMethod('onItemCheckboxChange', this.onItemCheckboxChange(view))
    }
    if(this.props.getView) {
        return this.props.getView(view, header, groups, pagination, filter)
    }
    let viewComp = null
    if(this.props.incrementalLoad) {
      viewComp = (
        <ScrollListener key={this.props.key} className={this.props.className} style={this.props.style} onScrollEnd={this.onScrollEnd}>
          <this.div className={this.props.headerClass}>
            {header}
          </this.div>
          <this.div className={this.props.contentClass}>
            {groups}
          </this.div>
        </ScrollListener>
      )
    } else {
      viewComp = (
        <this.div key={this.props.key} >
          <this.div className="viewheader">
            {header}
          </this.div>
          <this.div className="viewcontent">
            {groups}
          </this.div>
        </this.div>
      )
    }

    return (
      <this.div key={this.props.key} className={this.props.className} style={this.props.style}>
        {filter}
        {viewComp}
        {pagination}
      </this.div>
    )
  }
  getItemGroup(view, x) {
    return x
  }
  getHeader(view) {
    if(this.props.getHeader) {
      return this.props.getHeader(view)
    }
    return null
  }
  getItem(view, x, i) {
    if(this.props.getItem) {
      return this.props.getItem(view, x, i)
    }
    return React.Children.map(this.props.children, (child) => React.cloneElement(child, { item: x, index: i }) );
  }*/

  /*
  getPagination(view, pages, page) {
    if(pages == 1) {
      return null
    }
    return (
      <Paginator.Context
        {...pagifyBootstrapPreset}
        segments={segmentize({
            pages,
            page,
            beginPages: 1,
            endPages: 1,
            sidePages: 3
        })}
        onSelect={(newPage, event) => {
            event.preventDefault();
            view.setPage(newPage);
        }}
      >
        <Paginator.Button page={page - 1}>{'<'}</Paginator.Button>
        <Paginator.Segment field="beginPages" />
        <Paginator.Ellipsis previousField="beginPages" nextField="previousPages" />
        <Paginator.Segment field="previousPages" />
        <Paginator.Segment field="centerPage" className="active" />
        <Paginator.Segment field="nextPages" />
        <Paginator.Ellipsis previousField="nextPages" nextField="endPages" />
        <Paginator.Segment field="endPages" />
        <Paginator.Button page={page + 1}>{'>'}</Paginator.Button>
      </Paginator.Context>
    )
  }*/
  /*render() {
    return (
      <View
        ref="view"
        key={this.props.key}
        reducer={this.props.reducer}
        paginate={this.props.paginate}
        pageSize={this.props.pageSize}
        viewService={this.props.viewService}
        urlParams = {this.props.urlParams}
        postArgs = {this.props.postArgs}
        defaultFilter = {this.props.defaultFilter}
        currentPage={this.props.currentPage}
        filterTitle= {this.props.filterTitle}
        filterForm={this.props.filterForm}
        filterGo={ this.props.filterGo}
        loader = {this.props.loader}
        getView={this.getView}
        getItem={this.props.getItem}
        globalReducer={this.props.globalReducer}
        getFilter={this.getFilter}
        getItemGroup={this.getItemGroup}
        getHeader={this.getHeader}
        style={this.props.style}
        className={this.props.className}
        incrementalLoad={this.props.incrementalLoad}
        getPagination={this.props.incrementalLoad || this.props.hidePaginationControl ? null : this.props.getPagination} >
      </View>
    )
  }
}

WebView.contextTypes = {
  uikit: PropTypes.object
};
export {
  WebView
}
*/

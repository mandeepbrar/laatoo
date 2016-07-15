'use strict';

import React from 'react';
import Paginator from 'react-pagify';
import pagifyBootstrapPreset from 'react-pagify-preset-bootstrap';
import segmentize from 'segmentize';
import {View} from './View';
import {ViewFilter} from './Filter';

class WebView extends React.Component {
  constructor(props) {
    super(props);
    this.getPagination = this.getPagination.bind(this);
    this.getView = this.getView.bind(this);
    this.getFilter = this.getFilter.bind(this);
    this.addMethod = this.addMethod.bind(this);
    this.methods = this.methods.bind(this);
  }

  addMethod(name, method) {
    return this.refs.view.getWrappedInstance().addMethod(name, method)
  }
  methods() {
    return this.refs.view.getWrappedInstance().methods
  }
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
  }

  getView(view, header, groups, pagination, filter) {
    view.addMethod('onItemCheckboxChange', this.onItemCheckboxChange(view))
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

  getPagination(view, pages, page) {
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
  }
  render() {

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
        getItemGroup={this.props.getItemGroup}
        getHeader={this.props.getHeader}
        style={this.props.style}
        className={this.props.className}
        incrementalLoad={this.props.incrementalLoad}
        getPagination={this.props.incrementalLoad ? null : this.getPagination} >
      </View>
    )
  }
}

export {
  WebView
}

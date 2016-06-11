'use strict';

import React from 'react';
import Paginator from 'react-pagify';
import pagifyBootstrapPreset from 'react-pagify-preset-bootstrap';
import segmentize from 'segmentize';
import {View} from './View'

class WebView extends React.Component {
  constructor(props) {
    super(props);
    this.getPagination = this.getPagination.bind(this);
    this.getView = this.getView.bind(this);
    this.selectedItems = this.selectedItems.bind(this);
  }

  selectedItems() {
    return this.refs.view.getWrappedInstance().selectedItems()
  }

  getView(view, header, groups, pagination) {
    view.onItemCheckboxChange = this.onItemCheckboxChange(view)
    let viewComp = null
    if(this.props.getView) {
        viewComp = this.props.getView(view, header, groups)
    }
    return (
      <div className={this.props.className}>
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
        <Paginator.Button page={page - 1}>Previous</Paginator.Button>
        <Paginator.Segment field="beginPages" />
        <Paginator.Ellipsis previousField="beginPages" nextField="previousPages" />
        <Paginator.Segment field="previousPages" />
        <Paginator.Segment field="centerPage" className="active" />
        <Paginator.Segment field="nextPages" />
        <Paginator.Ellipsis previousField="nextPages" nextField="endPages" />
        <Paginator.Segment field="endPages" />
        <Paginator.Button page={page + 1}>Next</Paginator.Button>
      </Paginator.Context>
    )
  }
  render() {
    return (
      <View
        ref="view"
        reducer={this.props.reducer}
        paginate={this.props.paginate}
        pageSize={this.props.pageSize}
        viewService={this.props.viewService}
        currentPage={this.props.currentPage}
        getView={this.getView}
        getItem={this.props.getItem}
        getItemGroup={this.props.getItemGroup}
        getHeader={this.props.getHeader}
        getPagination={this.getPagination} >
      </View>
    )
  }
}

export {
  WebView
}

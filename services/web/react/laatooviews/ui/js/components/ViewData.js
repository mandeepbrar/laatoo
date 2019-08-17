import React from 'react';

class ViewItem {
  constructor() {
    this.index = -1;
    this.data = null;
    this.renderedItem = null;
  }
}

class ViewData extends React.Component {
  constructor(props) {
    super(props);
    this.setPage = this.setPage.bind(this);
    this.selectedItems = this.selectedItems.bind(this);
    this.itemCount = this.itemCount.bind(this);
    this.reload = this.reload.bind(this);
    this.setFilter = this.setFilter.bind(this);
    this.loadMore = this.loadMore.bind(this);
    this.canLoadMore = this.canLoadMore.bind(this);
    //this.getView = this.getView.bind(this);
    this.methods = {reload: this.reload, canLoadMore: this.canLoadMore, loadMore: this.loadMore, setFilter:this.setFilter,
        itemCount: this.itemCount, itemSelectionChange: this.itemSelectionChange, selectedItems: this.selectedItems, setPage: this.setPage}
    this.addMethod = this.addMethod.bind(this);
    this.state = {lastLoadTime: -1}
  //  this.numItems = 0
    this.pushItem = this.pushItem.bind(this)
    this.viewitems = new Array()
    console.log("this items", this.viewitems)
  }
  componentWillMount() {
    this.filter = this.props.defaultFilter
  }
  componentDidMount() {
    if(this.props.load && !this.props.externalLoad) {
      this.props.loadView(this.props.currentPage, this.filter);
    }
  }
  componentWillReceiveProps(nextprops) {
    if(nextprops.load) {
      nextprops.loadView(nextprops.currentPage, this.filter);
    }
  }
  shouldComponentUpdate(nextProps, nextState) {
    if(!nextProps.forceUpdate && this.lastRenderTime) {
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
  addMethod(name, method) {
    this.methods[name] = method
  }
  reload() {
    this.props.loadView(this.props.currentPage, this.filter);
  }
  canLoadMore() {
    return this.props.currentPage < this.props.totalPages
  }
  pushItem(item) {
    console.log("pushing item", this.viewitems, this)
    this.viewitems.push(item)
  }
  itemCount() {
    return this.viewitems.length
  }
  setSelection(item, val){
    if(item.setSelected) {
      item.setSelected(val)
    } else {
      item.selected = val
    }
  }
  itemSelectionChange = (i, val) => {
    let viewItem = this.viewitems[i]
    if(viewItem) {
      let renderedItem = viewItem.renderedItem
      this.setSelection(renderedItem, val)
    }
    if(this.props.singleSelection) {
      if(this.selectedItem) {
        this.setSelection(this.selectedItem.renderedItem, false)
      }
      this.selectedItem = viewItem
    }
  }
  selectedItems() {
    if(this.props.singleSelection) {
      return this.selectedItem.data
    }
    let selectedItems = []
    let numItems = this.itemCount()
    console.log("num items ", numItems)
    for(var i=0; i<numItems;i++) {
      let vitem = this.viewitems[i]
      console.log("vitem", vitem, i)
      let renderedItem = vitem.renderedItem
      if((renderedItem.getSelected && renderedItem.getSelected()) || renderedItem.selected) {
        selectedItems.push(vitem.data)
      }
    }
    return selectedItems
  }
  /*itemStatus() {
    let items = {}
    for(var i=0; i<this.numItems;i++) {
      let refName = "item"+i
      let item = this.refs[refName]
      items[item.id] = item.selected
    }
    return items
  }*/
  setPage(newPage) {
    this.props.loadView(newPage, this.filter)
  }
  setFilter(filter) {
    this.filter = filter
    this.props.loadView(1, this.filter)
  }
/*  getView(items, currentPage, totalPages) {
    if(this.props.getView) {
      return this.props.getView(this, items, currentPage, totalPages)
    }
    return null
  }*/
  loadMore() {
    if(this.props.currentPage>=this.props.totalPages) {
      return false
    } else {
      if(this.props.currentPage) {
        this.props.loadIncrementally(this.props.currentPage + 1, this.filter)
        return true
      }
    }
  }
  render() {
    this.lastRenderTime = this.props.lastUpdateTime
    let view = this.renderView(this.props.items, this.props.currentPage, this.props.totalPages)
    console.log("rendering view data", view, this.props.items)
    this.items=this.props.items
    return view
  }
}


export {ViewData, ViewItem }

import  {ActionNames} from './Actions';

var initialState={
    views:{}
}

var initialViewState = {
  status: "NotLoaded",
  data:{},
  global: false,
  currentPage: 1,
  totalPages: 1,
  pagesize: -1
};

function ViewReducer (state, action) {
  //return state if this is not the correct copy of reducer
  if(!action || !action.meta || !action.meta.viewname) {
    if(!state) {
      return initialState;
    }
    if(action.type == ActionNames.LOGOUT) {
      return initialState;
    }
    if(action.type == ActionNames.PAGE_CHANGE) {
      let views=[];
      Object.keys(state.views).forEach(function(key){
        let view = state.views[key];
        if(view.global) {
          views.push(view);
        }
      });
      return Object.assign({}, state, {views});
    }
    return state;
  }
  if (action.type && action.meta && action.meta.viewname) {
    let name = action.meta.viewname
    let oldViewState = state.views[name]
    let newViewState = {}
    switch (action.type) {

      case ActionNames.VIEW_FETCHING:
        let pagenum = 1
        let pagesize = -1
        if(action.payload.queryParams && action.payload.queryParams.pagenum) {
          pagenum = action.payload.queryParams.pagenum
          pagesize = action.payload.queryParams.pagesize
        }
        let stateChange = {
          status: "Fetching",
          currentPage: pagenum,
          pagesize: pagesize
        }
        if(oldViewState) {
          newViewState = Object.assign({}, oldViewState, stateChange)
        } else {
          newViewState = Object.assign({}, initialViewState, stateChange)
        }
      case ActionNames.VIEW_FETCH_SUCCESS:
        let totalPages = 1
        if (action.meta.info && action.meta.info.totalrecords) {
          let totalrecords = action.meta.info.totalrecords
          if(totalrecords >0 && oldViewState.pagesize > 0) {
            totalPages = Math.ceil(totalrecords / oldViewState.pagesize)
          }
        }
        let newData = null
        let data = oldViewState.data
        if(data && action.meta.incrementalLoad) {
          if(action.payload) {
            if(Array.isArray(action.payload)) {
              newData = data.concat(action.payload)
            } else {
              newData = Object.assign(data, action.payload)
            }
          }
        } else {
          newData = action.payload
        }
        newViewState = Object.assign({}, oldViewState, {
          status:"Loaded",
          data: newData,
          lastUpdateTime: (new Date()).getTime(),
          totalPages: totalPages
        });

      case ActionNames.VIEW_FETCH_FAILED: {
        newViewState = Object.assign({}, initialViewState, {
          status:"LoadingFailed"
        });
      }

      case ActionNames.VIEW_ITEM_RELOAD: {
        let index = action.meta.Index
        if (index==null) {
          return state
        }
        let data = oldViewState.data
        let newData = null
        //remove by index for arrays and keys for map
        if(Array.isArray(data)) {
          let ind = -1
          if((typeof index) == "string") {
            ind = parseInt(index)
          } else {
            ind = index
          }
          newData = data.slice(0)
          newData[ind] = action.payload;
        } else {
          newData = Object.assign({}, data);
          newData[index] = action.payload;
        }
        newViewState = Object.assign({}, oldViewState, {
          data: newData,
          lastUpdateTime: (new Date()).getTime(),
        });
      }

      case ActionNames.VIEW_ITEM_REMOVE: {
        let index = action.payload.Index
        if (index==null) {
          return state
        }
        let data = oldViewState.data
        let newData = null
        //remove by index for arrays and keys for map
        if(Array.isArray(data)) {
          let ind = -1
          if((typeof index) == "string") {
            ind = parseInt(index)
          } else {
            ind = index
          }
          newData = data.slice(0)
          newData.splice(ind, 1)
        } else {
          newData = Object.assign({}, data);
          delete newData[index]
          if(newData == null) {
            newData = {}
          }
        }
        newViewState = Object.assign({}, oldViewState, {
          data: newData,
          lastUpdateTime: (new Date()).getTime(),
        });
      }
    }
    state.views = Object.assign({}, state.views, {name: newViewState});
    return state
  }
}



Application.Register('Reducers', "views", ViewReducer)

/*
export {
  ViewReducer as ViewReducer
};
*/

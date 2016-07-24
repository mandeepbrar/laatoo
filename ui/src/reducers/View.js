import  {ActionNames} from '../actions/ActionNames';

function ViewReducer(reducerName) {
  var initialState = {
    status: "NotLoaded",
    data:{},
    currentPage: 1,
    totalPages: 1,
    pagesize: -1
  };

  return (state, action) => {
    //return state if this is not the correct copy of reducer
    if(!action || !action.meta || !action.meta.reducer || (reducerName != action.meta.reducer)) {
      if(!state) {
        return initialState;
      }
      return state;
    }
    if (action.type) {
      switch (action.type) {
        case ActionNames.VIEW_FETCHING:
          let pagenum = 1
          let pagesize = -1
          if(action.payload.queryParams && action.payload.queryParams.pagenum) {
            pagenum = action.payload.queryParams.pagenum
            pagesize = action.payload.queryParams.pagesize
          }
          return Object.assign({}, state, {
            status: "Fetching",
            currentPage: pagenum,
            pagesize: pagesize
          });


        case ActionNames.VIEW_FETCH_SUCCESS:
          let totalPages = 1
          if (action.meta.info && action.meta.info.totalrecords) {
            let totalrecords = action.meta.info.totalrecords
            if(totalrecords >0 && state.pagesize > 0) {
              totalPages = Math.ceil(totalrecords / state.pagesize)
            }
          }
          let newData = null
          let data = state.data
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
          return Object.assign({}, state, {
            status:"Loaded",
            data: newData,
            lastUpdateTime: (new Date()).getTime(),
            totalPages: totalPages
          });

        case ActionNames.VIEW_FETCH_FAILED: {
          return Object.assign({}, initialState, {
            status:"LoadingFailed"
          });
        }

        case ActionNames.VIEW_ITEM_RELOAD: {
          let index = action.meta.Index
          if (index==null) {
            return state
          }
          let data = state.data
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
            console.log("newData ", newData, " index", ind, "  payload ", action.payload)
            newData[ind] = action.payload;
          } else {
            newData = Object.assign({}, data);
            newData[index] = action.payload;
          }
          return Object.assign({}, state, {
            data: newData,
            lastUpdateTime: (new Date()).getTime(),
          });
        }

        case ActionNames.VIEW_ITEM_REMOVE: {
          let index = action.payload.Index
          if (index==null) {
            return state
          }
          let data = state.data
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
          return Object.assign({}, state, {
            data: newData,
            lastUpdateTime: (new Date()).getTime(),
          });
        }

        default:
          if (!state) {
            return initialState;
          }
          return state;
      }
    }
  }

}

export {
  ViewReducer as ViewReducer
};

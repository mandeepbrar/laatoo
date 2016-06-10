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
          return Object.assign({}, state, {
            status:"Loaded",
            data: action.payload,
            totalPages: totalPages
          });

        case ActionNames.VIEW_FETCH_FAILURE: {
          return Object.assign({}, initialState, {
            status:"LoadingFailed"
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

import  {ActionNames} from '../actions/ActionNames';

function ViewReducer(reducerName) {
  var initialState = {
    status: "NotLoaded",
    data:{},
    currentPage: 1,
    totalPages: 1,
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
          return Object.assign({}, state, {
            status: "Fetching"
          });


        case ActionNames.VIEW_FETCH_SUCCESS:
          return Object.assign({}, state, {
            status:"Loaded",
            data: action.payload
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

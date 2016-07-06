import  {ActionNames} from '../actions/ActionNames';

function EntityReducer(reducerName) {
  var initialState = {
    status: "NotLoaded",
    data:{}
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
        case ActionNames.ENTITY_GETTING:
          return Object.assign({}, state, {
            status: "Fetching"
          });

        case ActionNames.ENTITY_GET_SUCCESS:
          return Object.assign({}, state, {
            status:"Loaded",
            data: action.payload
          });

        case ActionNames.ENTITY_GET_FAILED: {
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
  EntityReducer as EntityReducer
};

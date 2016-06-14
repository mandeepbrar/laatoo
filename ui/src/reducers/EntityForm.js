import  {ActionNames} from '../actions/ActionNames';

function EntityFormReducer(reducerName) {
  var initialState = {
    status: "Empty",
    entityId: "",
    entityName:"",
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
            status: "Fetching Entity",
            entityName: action.payload.entityName,
            entityId: action.payload.entityId
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

        case ActionNames.ENTITY_SAVING: {
          return Object.assign({}, initialState, {
            status:"Saving",
            entityName: action.payload.entityName,
            data: action.payload.data
          });
        }

        case ActionNames.ENTITY_SAVE_SUCCESS: {
          return Object.assign({}, initialState, {
            status:"Saved"
          });
        }

        case ActionNames.ENTITY_SAVE_FAILURE: {
          return Object.assign({}, initialState, {
            status:"SavingFailed",
            data: action.payload
          });
        }

        case ActionNames.ENTITY_UPDATING: {
          return Object.assign({}, initialState, {
            status:"Updating",
            entityName: action.payload.entityName,
            entityId: action.payload.entityId,
            data: action.payload.data
          });
        }

        case ActionNames.ENTITY_UPDATE_SUCCESS: {
          return Object.assign({}, initialState, {
            status:"Updated"
          });
        }

        case ActionNames.ENTITY_UPDATE_FAILURE: {
          return Object.assign({}, initialState, {
            status:"UpdateFailed",
            data: action.payload
          });
        }

        case ActionNames.ENTITY_PUTTING: {
          return Object.assign({}, initialState, {
            status:"Updating",
            entityName: action.payload.entityName,
            entityId: action.payload.entityId,
            data: action.payload.data
          });
        }

        case ActionNames.ENTITY_PUT_SUCCESS: {
          return Object.assign({}, initialState, {
            status:"Updated"
          });
        }

        case ActionNames.ENTITY_PUT_FAILURE: {
          return Object.assign({}, initialState, {
            status:"UpdateFailed",
            data: action.payload
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
  EntityFormReducer as EntityFormReducer
};

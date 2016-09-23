import  {ActionNames} from '../actions/ActionNames';

function EntityReducer(reducerName) {
  var initialState = {
    status: "NotLoaded",
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
            status: "Loading",
            entityName: action.payload.entityName,
            entityId: action.payload.entityId
          });

        case ActionNames.ENTITY_GET_SUCCESS:
          let st = Object.assign({}, state, {
            status:"Loaded",
            lastUpdateTime: (new Date()).getTime(),
            data: action.payload.data
          });
          return st
        case ActionNames.ENTITY_GET_FAILED: {
          return Object.assign({}, state, {
            status:"LoadingFailed",
            data: null
          });
        }

        case ActionNames.ENTITY_SAVING: {
          return Object.assign({}, state, {
            status:"Saving",
            entityName: action.payload.entityName
          });
        }

        case ActionNames.ENTITY_SAVE_SUCCESS: {
          return Object.assign({}, state, {
            status:"Saved"
          });
        }

        case ActionNames.ENTITY_SAVE_FAILURE: {
          return Object.assign({}, state, {
            status:"SavingFailed"
          });
        }

        case ActionNames.ENTITY_UPDATING: {
          return Object.assign({}, state, {
            status:"Updating",
            entityName: action.payload.entityName,
            entityId: action.payload.entityId
          });
        }

        case ActionNames.ENTITY_UPDATE_SUCCESS: {
          return Object.assign({}, state, {
            status:"Updated"
          });
        }

        case ActionNames.ENTITY_UPDATE_FAILURE: {
          return Object.assign({}, state, {
            status:"UpdateFailed"
          });
        }

        case ActionNames.ENTITY_PUTTING: {
          return Object.assign({}, state, {
            status:"Updating",
            entityName: action.payload.entityName,
            entityId: action.payload.entityId
          });
        }

        case ActionNames.ENTITY_PUT_SUCCESS: {
          return Object.assign({}, state, {
            status:"Updated"
          });
        }

        case ActionNames.ENTITY_PUT_FAILURE: {
          return Object.assign({}, state, {
            status:"UpdateFailed"
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

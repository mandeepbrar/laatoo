import  {ActionNames} from './Actions';

let reducerName = "entityview"

var initialState={
    entities:{}
}

var initialEntityState = {
  status: "NotLoaded",
  data:{},
  global: false,
  entityId: "",
  entityName:""
};

function EntityViewReducer (state, action)  {
  //return state if this is not the correct copy of reducer
  if(!action || !action.meta || (action.meta.reducer != reducerName)) {
    if(!state) {
      return initialState;
    }
    if(action.type == ActionNames.LOGOUT) {
      return initialState;
    }
    if(action.type == ActionNames.PAGE_CHANGE) {
      let entities=[];
      Object.keys(state.entities).forEach(function(key){
        let entity = state.entities[key];
        if(entity.global) {
          entities.push(entity);
        }
      });
      return Object.assign({}, state, {entities});
    }
    return state;
  }
  if (action.type && action.meta && (action.meta.reducer == reducerName)) {
    let id = action.meta.entityId
    let oldEntityState = state.entities[id]
    let newEntityState = {}
    switch (action.type) {

      case ActionNames.ENTITY_VIEW_FETCHING:
        let stateChange = {
          status: "Fetching",
          global: action.meta.global,
          entityId: id,
          entityName: action.payload.entityName
        }
        if(oldEntityState) {
          newEntityState[id] = Object.assign({}, oldEntityState, stateChange)
        } else {
          newEntityState[id] = Object.assign({}, initialEntityState, stateChange)
        }
        break;
      case ActionNames.ENTITY_VIEW_FETCH_SUCCESS:
        newEntityState[id] = Object.assign({}, oldEntityState, {
          status:"Loaded",
          data: action.payload.data,
          lastUpdateTime: (new Date()).getTime()
        });
        break;
      case ActionNames.ENTITY_VIEW_FETCH_FAILED: {
        newEntityState[id] = Object.assign({}, initialEntityState, {
          status:"LoadingFailed"
        });
        break;
      }
    }
    let newEntities = Object.assign({}, state.entities, newEntityState)
    return Object.assign({}, state, {entities: newEntities});
    /*console.log("new state of entity view reducer", state)
    return state*/
  }
}



Application.Register('Reducers', reducerName, EntityViewReducer)

/*
export {
  ViewReducer as ViewReducer
};
*/

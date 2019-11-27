import  Actions from '../actions';

var initialState = {
};

const Interactions = (state, action) => {
  if (action.type) {
    switch (action.type) {
      case Actions.PAGE_CHANGE: 
      case Actions.LOGOUT: {
        return initialState
      }
      case Actions.SHOW_INTERACTION_COMP:
      console.log("show interaction ", action)
        return {
          Content: action.payload,
          Type: "Interaction",
          Time: (new Date()).getTime()
        };
      case Actions.CLOSE_INTERACTION_COMP:
        return {
          Content: null,
          Type: "Close",
          Time: (new Date()).getTime()
        };
      default:
        if (!state) {
          return initialState;
        }
        return state;
    }
  }
}

Application.Register('Reducers', "Interactions", Interactions)

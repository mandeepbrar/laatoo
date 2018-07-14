import  Actions from '../actions';

var initialState = {
};

const Dialogs = (state, action) => {
  if (action.type) {
    switch (action.type) {
      case Actions.LOGOUT: {
        return initialState
      }
      case Actions.SHOW_DIALOG:
      console.log("show dialog ", action)
        return {
          Content: action.payload,
          Type: "Dialog",
          Time: (new Date()).getTime()
        };
      case Actions.CLOSE_DIALOG:
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

Application.Register('Reducers', "Dialogs", Dialogs)

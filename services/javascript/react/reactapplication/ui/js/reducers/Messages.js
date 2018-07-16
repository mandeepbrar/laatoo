import  Actions from '../actions';

var initialState = {
};

const Messages = (state, action) => {
  if (action.type) {
    switch (action.type) {
      case Actions.LOGOUT: {
        return initialState
      }
      case Actions.DISPLAY_ERROR:
        return {
          Message: action.payload.message,
          Type: "Error",
          Time: (new Date()).getTime()
        };
      case Actions.SHOW_MESSAGE:
        return {
          Message: action.payload.message,
          Type: "Message",
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

Application.Register('Reducers', "Messages", Messages)

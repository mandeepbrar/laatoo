import  {ActionNames} from '../actions/ActionNames';

var initialSecState = {
  status:"NotLogged",
  token: "",
  userId: "",
  permissions: []
};

const Account = (state, action) => {
  if (action.type) {
    switch (action.type) {
      case ActionNames.LOGGING_IN:
        return Object.assign({}, state, {
          status: "LoggingIn"
        })

      case ActionNames.LOGIN_SUCCESS:
        if (state.authToken === action.payload.token) {
          return state
        }
        localStorage.auth = action.payload.token;
        localStorage.permissions = action.payload.permissions;
        localStorage.user = action.payload.userId;
        return Object.assign({}, state, {
          status: "LoggedIn",
          authToken: action.payload.token,
          userId: action.payload.userId,
          permissions: action.payload.permissions
        })

      case ActionNames.LOGIN_FAILURE:
        localStorage.auth = "";
        localStorage.permissions = [];
        localStorage.user = "";
        return initialSecState;

      case ActionNames.LOGOUT:
        localStorage.auth = "";
        localStorage.permissions = [];
        localStorage.user = "";
        return initialSecState;

      default:
        if (!state) {
          if (localStorage.auth != null && localStorage.auth != "") {
            return {
              status: "LoggedIn",
              authToken: localStorage.auth,
              userId: localStorage.user,
              permissions: localStorage.permissions
            };
          }
          return initialSecState;
        }
        return state;
    }
  }
}

export {
  Account as Account
};

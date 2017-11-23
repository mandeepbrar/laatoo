import  {ActionNames} from '../actions';

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
        Storage.auth = action.payload.token;
        Storage.permissions = action.payload.permissions;
        Storage.userId = action.payload.userId;
        Storage.userName = action.payload.user.Name;
        Storage.email = action.payload.user.email;
        Storage.user = action.payload.user;
        return Object.assign({}, state, {
          status: "LoggedIn",
          authToken: action.payload.token,
          userId: action.payload.userId,
          permissions: action.payload.permissions
        })

      case ActionNames.LOGIN_FAILURE:
        Storage.auth = "";
        Storage.permissions = [];
        Storage.userId = "";
        Storage.userName = "";
        Storage.email = "";
        Storage.user = null;
        return initialSecState;

      case ActionNames.LOGOUT:
        Storage.auth = "";
        Storage.permissions = [];
        Storage.userId = "";
        Storage.userName = "";
        Storage.email = "";
        Storage.user = null;
        return initialSecState;

      default:
        if (!state) {
          return initialSecState;
        }
        return state;
    }
  }
}

Application.Register('Reducers', "Security", Account)
/*
export {
  Account as Account
};*/

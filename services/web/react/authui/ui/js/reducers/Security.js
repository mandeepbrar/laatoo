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

        case ActionNames.VALIDATIONFAILED:
          Storage.auth = "";
          Storage.permissions = [];
          Storage.userId = "";
          Storage.userName = "";
          Storage.userFullName = "";
          Storage.email = "";
          Storage.user = null;
          return Object.assign({}, initialSecState, {
            status: "ValidationFailed"
          });
        case ActionNames.LOGIN_SUCCESS:
          if (state.authToken === action.payload.token) {
            return state
          }
          Storage.auth = action.payload.token;
          Storage.permissions = action.payload.permissions;
          Storage.userId = action.payload.userId;
          Storage.userFullName = action.payload.user.Name;
          Storage.userName = action.payload.user.Username;
          Storage.email = action.payload.user.Email;
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
        Storage.userFullName = "";
        Storage.email = "";
        Storage.user = null;
        return initialSecState;

      case ActionNames.LOGOUT_SUCCESS:
        Storage.auth = "";
        Storage.permissions = [];
        Storage.userId = "";
        Storage.userName = "";
        Storage.userFullName = "";
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

import  {ActionNames} from '../actions/ActionNames';
import {Application, Storage} from '../Globals'

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
        Storage.user = action.payload.userId;
        return Object.assign({}, state, {
          status: "LoggedIn",
          authToken: action.payload.token,
          userId: action.payload.userId,
          permissions: action.payload.permissions
        })

      case ActionNames.LOGIN_FAILURE:
        Storage.auth = "";
        Storage.permissions = [];
        Storage.user = "";
        return initialSecState;

      case ActionNames.LOGOUT:
        Storage.auth = "";
        Storage.permissions = [];
        Storage.user = "";
        return initialSecState;

      default:
        if (!state) {
          if (Storage.auth != null && Storage.auth != "") {
            return {
              status: "LoggedIn",
              authToken: Storage.auth,
              userId: Storage.user,
              permissions: Storage.permissions
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

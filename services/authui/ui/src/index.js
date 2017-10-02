import {LoginComponent} from './LoginComponent';
import {LoginValidator} from './LoginValidator';
import './reducers/Security';
import './sagas/Security';

var module = this;

function Initialize(appName, ins, mod, settings, def, req) {
  module.properties = Application.Properties[ins]
  module.settings = settings;
  if(Object.keys(settings).length != 0 ) {
    Application.Security = {
      googleAuthUrl: settings.googleAuthUrl,
      loginService: settings.loginService,
      validateService: settings.validateService,
      loginServiceURL: settings.loginServiceURL,
      realm: settings.realm
    }
  } else {
    Application.Security = {
      loginService: "login",
      validateService: "validate",
      realm: ""
    }
    if((Application.Registry.Service==null) || (Application.Registry.Service["login"]==null)) {
      Application.Register('Service', 'login', {url:"/login", method:'POST'})
      Application.Register('Service', 'validate', {url:"/validate", method:'POST'})
    }
  }
  if(settings.AuthToken) {
    Application.Security.AuthToken = settings.AuthToken
  } else {
    Application.Security.AuthToken = "x-auth-token"
  }
}

export {
  Initialize,
  LoginComponent,
  LoginValidator
}

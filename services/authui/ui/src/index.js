import {LoginComponent} from './LoginComponent';
import {LoginValidator} from './LoginValidator';
import './reducers/Security';
import './sagas/Security';

var module = this;

function Initialize(appName, ins, mod, settings, def, req) {
  console.log("init authui")
  module.properties = Application.Properties[ins]
  module.settings = settings;
  if(settings) {
    Application.Security = {
      googleAuthUrl: settings.googleAuthUrl,
      loginService: settings.loginService,
      loginServiceURL: settings.loginServiceURL,
      realm: settings.realm
    }
  } else {
    Application.Security = {
      loginService: "login",
      realm: ""
    }
    if(!Application.Services || !Application.Services.loginService) {
      Application.Register('Services', 'login', {url:"/login", method:'POST'})
    }
  }
  console.log("init authui ed", Application)
}

export {
  Initialize,
  LoginComponent,
  LoginValidator
}

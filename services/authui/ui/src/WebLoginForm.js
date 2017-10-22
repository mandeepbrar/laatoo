import {Action} from 'reactwebcommon';
import React from 'react';
import './styles/app.scss'

function renderWebLogin(uikit, settings, properties) {
  return function (state, handleChange, handleLogin, oauthLogin, props) {
    let openGoogleauthWindow = function() {
      oauthLogin(Application.Security.googleAuthUrl)
    };
    console.log("renderLogin", properties,"uikit", uikit,"settigs", settings, "props", props);
    return (
        <div className={ props.className? props.className: " loginbox "}>
          <div className="logintext">
            {properties.loginForm.formtext}
          </div>
          <div className="sociallogin">
            <Action widget="button" method={openGoogleauthWindow} name="googleAuth" className="googleAuthAction">{properties.loginForm.google}</Action>
          </div>
          <div className="separator">
            {properties.loginForm.separator}
          </div>
          <div className="main">
            <uikit.Form role="form">
              <div className="userfield">
                <label htmlFor="email">{properties.loginForm.userlabel}</label>
                <uikit.TextField className="text" name="email" value={state.email} placeholder={properties.loginForm.userplaceholder} onChange={handleChange} />
              </div>
              <div className="passwordfield">
                <label htmlFor="inputPassword">{properties.loginForm.passwordlabel}</label>
                <uikit.TextField type="password" className="text" name="password" value={state.password} placeholder={properties.loginForm.passwordplaceholder} onChange={handleChange} />
              </div>
              <a className="pull-right" href="#">Forgot password?</a>
              <div className="checkbox">
                  <label>
                      <input type="checkbox"/>
                      Remember me
                  </label>
              </div>
              <div className="actionbuttons">
                <Action widget="button" className="loginBtn" name="loginAction" method={handleLogin}>{properties.loginForm.loginBtnText}</Action>
              </div>
            </uikit.Form>
          </div>
        </div>
    );
  }
}


export {
  renderWebLogin
}

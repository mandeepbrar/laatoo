import {Action} from 'reactwebcommon';
import React from 'react';
import './styles/app.scss'

function renderSignup(uikit, settings, properties) {
  return function (state, handleChange, handleSignup, props) {
    console.log("renderSignup", properties,"uikit", uikit,"settings", settings, "props", props);
    return (
        <div className={ props.className? props.className: " signupbox "}>
          <div className="signuptext">
            {properties.signupForm.formtext}
          </div>
          <div className="main">
            <uikit.Form role="form">
              <div className="userfield">
                <label htmlFor="email">{properties.signupForm.userlabel}</label>
                <uikit.TextField className="text" name="email" value={state.email} placeholder={properties.signupForm.userplaceholder} onChange={handleChange} />
              </div>
              <div className="passwordfield">
                <label htmlFor="inputPassword">{properties.signupForm.passwordlabel}</label>
                <uikit.TextField type="password" className="text" name="password" value={state.password} placeholder={properties.signupForm.passwordplaceholder} onChange={handleChange} />
              </div>
              <div className="confirmpasswordfield">
                <label htmlFor="inputConfirmPassword">{properties.signupForm.confirmpasswordlabel}</label>
                <uikit.TextField type="password" className="text" name="confirmpassword" value={state.confirmpassword} placeholder={properties.signupForm.confirmpasswordplaceholder} onChange={handleChange} />
              </div>
              <div className="actionbuttons">
                <Action widget="button" className="signupBtn" name="signupAction" method={handleSignup}>{properties.signupForm.signupBtnText}</Action>
              </div>
            </uikit.Form>
          </div>
        </div>
    );
  }
}


export {
    renderSignup
}

package ginauth_local

import (
	"golang.org/x/crypto/bcrypt"
)

func (localauth *localAuth) configBCryptCost() error {
	bcryptcost, ok := localauth.app.Configuration[BCRYPTCOST]
	if ok {
		localauth.bCryptCost = bcryptcost.(int)
	} else {
		localauth.bCryptCost = bcrypt.DefaultCost
	}
	return nil
}

func (localauth *localAuth) configPaths() error {
	localauth.app.Logger.Debug("localAuthProvider: Configure Paths")
	localauth.loginpath = "/login"
	loginpath, ok := localauth.app.Configuration[LOGINPATH]
	if ok {
		localauth.loginpath = loginpath.(string)
	}
	return nil
}

func (localauth *localAuth) configRegistration() error {
	registrationAllowed, ok := localauth.app.Configuration[REGISTRATIONALLOWED]
	if ok {
		localauth.registrationAllowed = registrationAllowed.(bool)
	} else {
		localauth.registrationAllowed = true
	}
	registrationPath, ok := localauth.app.Configuration[REGISTRATIONPATH]
	if ok {
		localauth.registrationPath = registrationPath.(string)
	} else {
		localauth.registrationPath = "/register"
	}
	return nil
}

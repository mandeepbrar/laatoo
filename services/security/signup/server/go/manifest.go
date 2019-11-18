package main

import (
	"laatoo/sdk/server/core"
	"signup/signup"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Object: signup.RegistrationService{}},
		core.PluginComponent{Object: signup.SignupEmailTask{}},
		core.PluginComponent{Object: signup.VerifyEmailService{}},
	}
}

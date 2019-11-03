module laatoo/cli

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.5.0
	laatoo/sdk v0.0.0
)

replace laatoo/sdk => ../sdk

go 1.13

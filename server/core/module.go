package core

import (
	"laatoo/sdk/core"

	"github.com/blang/semver"
)

type module struct {
	name         string
	version      semver.Version
	dependencies map[string]semver.Range
	svrContext   *serverContext
}

func (mod *module) Reference() core.ServerElement {
	return mod
}

func (mod *module) GetProperty(string) interface{} {
	return nil
}
func (mod *module) GetName() string {
	return mod.name
}
func (mod *module) GetType() core.ServerElementType {
	return core.ServerElementModule
}

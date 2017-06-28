package core

import "laatoo/sdk/core"

type environmentProxy struct {
	env *environment
}

func (proxy *environmentProxy) Reference() core.ServerElement {
	return &environmentProxy{env: proxy.env}
}
func (proxy *environmentProxy) GetProperty(name string) interface{} {
	return nil
}
func (proxy *environmentProxy) GetName() string {
	return proxy.env.name
}
func (proxy *environmentProxy) GetType() core.ServerElementType {
	return core.ServerElementEnvironment
}

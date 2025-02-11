package elements

import "laatoo.io/sdk/server/core"

type InformationBucket interface {
	core.ConfigurableObjectInfo
}

type Agent interface {
	Service
	GetModel() string
	GetVersion() string
	GetInstructions() string
	GetDescription() string
	Information() []InformationBucket
	Tools() []Service
}

package core

import "laatoo.io/sdk/config"

type ActionType string

const (
	Transform         ActionType = "Transform"
	Notify                       = "Notify"
	SaveData                     = "SaveData"
	StartWorkflow                = "StartWorkflow"
	DeleteData                   = "DeleteData"
	UpdateDataset                = "UpdateDataset"
	InvokeScript                 = "InvokeScript"
	InvokeActivity               = "InvokeActivity"
	InvokeService                = "InvokeService"
	ValidateArguments            = "ValidateArguments"
)

type Action struct {
	Type      ActionType
	Condition *GenericExpression
	Config    *config.GenericConfig
}

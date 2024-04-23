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
	ValidateArguments            = "ValidateArguments"
)

type Action struct {
	StepType   ActionType
	Condition  *GenericExpression
	StepConfig *config.GenericConfig
}

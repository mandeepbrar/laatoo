package core

import "laatoo.io/sdk/config"

type ActionType string

const (
	Transform     ActionType = "Transform"
	Notify                   = "Notify"
	SaveData                 = "SaveData"
	StartWorkflow            = "StartWorkflow"
	DeleteData               = "DeleteData"
	UpdateDataset            = "UpdateDataset"
	InvokeScript             = "InvokeScript"
)

type Action interface {
	GetActionType() ActionType
	GetConfig() config.Config
	GetCondtion() string
}

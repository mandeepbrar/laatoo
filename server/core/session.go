package core

import (
	"time"
)

type Session interface {
	GetId() string
	CreationTime() time.Time
	GetUser() string
	GetString(key string) (string, bool)
	GetBool(key string) (bool, bool)
	GetInt(key string) (int, bool)
	GetStringArray(key string) ([]string, bool)
	AllKeys() []string
	GetStringMap(key string) (StringMap, bool)
	GetStringsMap(key string) (StringsMap, bool)
	Set(key string, val interface{})
	SetVals(vals StringMap)
	Save(RequestContext) error
}

package common

import "fmt"

func GetCacheKey(objectType string, variants ...interface{}) string {
	return fmt.Sprintf("%s_%#v", objectType, variants)
}

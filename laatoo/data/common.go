package data

import (
	"fmt"
)

func getCacheKey(storableName string, id string) string {
	return fmt.Sprintf("_%s_%s", storableName, id)
}

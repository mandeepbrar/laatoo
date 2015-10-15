package data

type Cache interface {
	PutObject(ctx interface{}, key string, item interface{}) error
	GetObject(ctx interface{}, key string, val interface{}) error
	GetMulti(ctx interface{}, keys []string, val map[string]interface{}) error
	Delete(ctx interface{}, key string) error
}

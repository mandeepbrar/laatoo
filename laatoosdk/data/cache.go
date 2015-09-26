package data

type Cache interface {
	PutObject(ctx interface{}, key string, item interface{}) error
	GetObject(ctx interface{}, key string) (interface{}, error)
	Delete(ctx interface{}, key string) error
}

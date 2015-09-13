package data

type Cache interface {
	PutObject(key string, item interface{}) error
	GetObject(key string) (interface{}, error)
	Delete(key string) error
}

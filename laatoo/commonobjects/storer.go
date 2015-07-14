package commonobjects

type Storer interface {
	GetName() string
	Put(id string, obj interface{}) error
	Delete(id string) error
	GetById(id string) (interface{}, error)
	Get(interface{}) (interface{}, error)
	GetList() (interface{}, error)
}

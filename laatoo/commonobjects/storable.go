package commonobjects

type Storable interface {
	GetId() string
	SetId(string)
	GetIdField() string
}

type StorableCreator func() interface{}

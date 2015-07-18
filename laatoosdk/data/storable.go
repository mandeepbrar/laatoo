package data

//Object stored by data service
type Storable interface {
	GetId() string
	SetId(string)
	GetIdField() string
}

//Factory function for creating storable
type StorableCreator func() interface{}

package data

//Object stored by data service
type Storable interface {
	GetId() string
	SetId(string)
	PreSave() error
	PostLoad() error
	GetIdField() string
}

//Factory function for creating storable
type StorableCreator func() interface{}

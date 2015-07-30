package data

//Service that provides data from various data sources
//Service interface that needs to be implemented by any data service
type DataService interface {
	GetDataServiceType() string
	//save an object
	Save(objectType string, item interface{}) error
	//Store an object against an id
	Put(objectType string, id string, obj interface{}) error
	//Delete an object by id
	Delete(objectType string, id string) error
	//Get an object by id
	GetById(objectType string, id string) (interface{}, error)
	//Get all object with given conditions
	Get(objectType string, conditions interface{}) (interface{}, error)
	//Get a list of all items
	GetList(objectType string) (interface{}, error)
}

const (
	DATASERVICE_TYPE_NOSQL = "SERVICE_TYPE_NOSQL"
	DATASERVICE_TYPE_SQL   = "SERVICE_TYPE_SQL"
	DATASERVICE_TYPE_CACHE = "SERVICE_TYPE_CACHE"
)

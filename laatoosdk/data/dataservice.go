package data

//Service that provides data from various data sources
//Service interface that needs to be implemented by any data service
type DataService interface {
	GetDataServiceType() string
	//Store an object against an id
	Put(id string, obj interface{}) error
	//Delete an object by id
	Delete(id string) error
	//Get an object by id
	GetById(id string) (interface{}, error)
	//Get all object with given conditions
	Get(interface{}) (interface{}, error)
	//Get a list of all items
	GetList() (interface{}, error)
}

const (
	DATASERVICE_TYPE_NOSQL = "SERVICE_TYPE_NOSQL"
	DATASERVICE_TYPE_SQL   = "SERVICE_TYPE_SQL"
	DATASERVICE_TYPE_CACHE = "SERVICE_TYPE_CACHE"
)

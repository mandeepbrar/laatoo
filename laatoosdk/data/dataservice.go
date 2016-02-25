package data

import (
	"github.com/labstack/echo"
)

//Service that provides data from various data sources
//Service interface that needs to be implemented by any data service
type DataService interface {
	GetDataServiceType() string
	//save an object
	Save(ctx *echo.Context, objectType string, item interface{}) error
	//Store an object against an id
	Put(ctx *echo.Context, objectType string, id string, obj interface{}) error
	//Store multiple objects
	PutMulti(ctx *echo.Context, objectType string, ids []string, items interface{}) error
	//Delete an object by id
	Delete(ctx *echo.Context, objectType string, id string) error
	//Get an object by id
	GetById(ctx *echo.Context, objectType string, id string) (interface{}, error)
	//Get multiple objects by id
	GetMulti(ctx *echo.Context, objectType string, ids []string, orderBy string) (map[string]interface{}, error)
	//Get all object with given conditions
	Get(ctx *echo.Context, objectType string, conditions interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn interface{}, totalrecs int, recsreturned int, err error)
	//Get a list of all items
	GetList(ctx *echo.Context, objectType string, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn interface{}, totalrecs int, recsreturned int, err error)
}

const (
	DATASERVICE_TYPE_NOSQL = "SERVICE_TYPE_NOSQL"
	DATASERVICE_TYPE_SQL   = "SERVICE_TYPE_SQL"
	DATASERVICE_TYPE_CACHE = "SERVICE_TYPE_CACHE"
)

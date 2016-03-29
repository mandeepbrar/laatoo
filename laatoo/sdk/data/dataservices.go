package data

import (
	"laatoo/sdk/core"
)

type Feature int

type ConditionType int

const (
	MATCHMULTIPLEVALUES ConditionType = iota // expects first value as field name and second value as array of values
	MATCHANCESTOR                            //expects collection name and id
	FIELDVALUE                               //expects map of field values
)

const (
	InQueries Feature = iota
	Ancestors
)

const (
	DATASERVICE_TYPE_NOSQL = "SERVICE_TYPE_NOSQL"
	DATASERVICE_TYPE_SQL   = "SERVICE_TYPE_SQL"
	DATASERVICE_TYPE_CACHE = "SERVICE_TYPE_CACHE"
	DATA_PAGENUM           = "pagenum"
	DATA_PAGESIZE          = "pagesize"
	DATA_RECSRETURNED      = "records"
	DATA_TOTALRECS         = "totalrecords"
)

//Service that provides data from various data sources
//Service interface that needs to be implemented by any data service
type DataService interface {
	core.Service
	GetDataServiceType() string
	//supported features
	Supports(Feature) bool
	//create condition for passing to data service
	CreateCondition(ctx core.RequestContext, operation ConditionType, args ...interface{}) (interface{}, error)
	//save an object
	Save(ctx core.RequestContext, item Storable) error
	//Store an object against an id
	Put(ctx core.RequestContext, id string, item Storable) error
	//Store multiple objects
	PutMulti(ctx core.RequestContext, ids []string, items []Storable) error
	//update objects by ids, fields to be updated should be provided as key value pairs
	UpdateMulti(ctx core.RequestContext, ids []string, newVals map[string]interface{}) error
	//update an object by ids, fields to be updated should be provided as key value pairs
	Update(ctx core.RequestContext, id string, newVals map[string]interface{}) error
	//update with condition
	UpdateAll(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error)
	//Delete an object by id
	Delete(ctx core.RequestContext, id string, softdelete bool) error
	//Delete object by ids
	DeleteMulti(ctx core.RequestContext, ids []string, softdelete bool) error
	//delete with condition
	DeleteAll(ctx core.RequestContext, queryCond interface{}, softdelete bool) ([]string, error)
	//Get an object by id
	GetById(ctx core.RequestContext, id string) (Storable, error)
	//Get multiple objects by id
	GetMulti(ctx core.RequestContext, ids []string, orderBy string) (map[string]Storable, error)
	//Get all object with given conditions
	Get(ctx core.RequestContext, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []Storable, totalrecs int, recsreturned int, err error)
	//Get a list of all items
	GetList(ctx core.RequestContext, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []Storable, totalrecs int, recsreturned int, err error)
}

//save an object
type Data_Save func(ctx core.RequestContext, item Storable) error

//Store an object against an id
type Data_Put func(ctx core.RequestContext, id string, item Storable) error

//Store multiple objects
type Data_PutMulti func(ctx core.RequestContext, ids []string, items []Storable) error

//update objects by ids, fields to be updated should be provided as key value pairs
type Data_UpdateMulti func(ctx core.RequestContext, ids []string, newVals map[string]interface{}) error

//update an object by ids, fields to be updated should be provided as key value pairs
type Data_Update func(ctx core.RequestContext, id string, newVals map[string]interface{}) error

//update with condition
type Data_UpdateAll func(ctx core.RequestContext, queryCond interface{}, newVals map[string]interface{}) ([]string, error)

//Delete an object by id
type Data_Delete func(ctx core.RequestContext, id string, softdelete bool) error

//Delete object by ids
type Data_DeleteMulti func(ctx core.RequestContext, ids []string, softdelete bool) error

//delete with condition
type Data_DeleteAll func(ctx core.RequestContext, queryCond interface{}, softdelete bool) ([]string, error)

//Get an object by id
type Data_GetById func(ctx core.RequestContext, id string) (Storable, error)

//Get multiple objects by id
type Data_GetMulti func(ctx core.RequestContext, ids []string, orderBy string) (map[string]Storable, error)

//Get all object with given conditions
type Data_Get func(ctx core.RequestContext, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []Storable, totalrecs int, recsreturned int, err error)

//Get a list of all items
type Data_GetList func(ctx core.RequestContext, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn []Storable, totalrecs int, recsreturned int, err error)

//create condition for passing to data service
type Data_CreateCondition func(ctx core.RequestContext, operation ConditionType, args ...interface{}) (interface{}, error)

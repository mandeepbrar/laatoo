package data

import (
	"laatoosdk/core"
)

type Feature int

type ConditionType int

//Service that provides data from various data sources
//Service interface that needs to be implemented by any data service
type DataService interface {
	GetDataServiceType() string
	//supported features
	Supports(Feature) bool
	//create condition for passing to data service
	CreateCondition(ctx core.Context, operation ConditionType, args ...interface{}) (interface{}, error)
	//save an object
	Save(ctx core.Context, objectType string, item interface{}) error
	//Store an object against an id
	Put(ctx core.Context, objectType string, id string, obj interface{}) error
	//Store multiple objects
	PutMulti(ctx core.Context, objectType string, ids []string, items interface{}) error
	//update objects by ids, fields to be updated should be provided as key value pairs
	UpdateMulti(ctx core.Context, objectType string, ids []string, newVals map[string]interface{}) error
	//update an object by ids, fields to be updated should be provided as key value pairs
	Update(ctx core.Context, objectType string, id string, newVals map[string]interface{}) error
	//update with condition
	UpdateAll(ctx core.Context, objectType string, queryCond interface{}, newVals map[string]interface{}) ([]string, error)
	//Delete an object by id
	Delete(ctx core.Context, objectType string, id string, softdelete bool) error
	//Delete object by ids
	DeleteMulti(ctx core.Context, objectType string, ids []string, softdelete bool) error
	//delete with condition
	DeleteAll(ctx core.Context, objectType string, queryCond interface{}, softdelete bool) ([]string, error)
	//Get an object by id
	GetById(ctx core.Context, objectType string, id string) (interface{}, error)
	//Get multiple objects by id
	GetMulti(ctx core.Context, objectType string, ids []string, orderBy string) (map[string]interface{}, error)
	//Get all object with given conditions
	Get(ctx core.Context, objectType string, conditions interface{}, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn interface{}, totalrecs int, recsreturned int, err error)
	//Get a list of all items
	GetList(ctx core.Context, objectType string, pageSize int, pageNum int, mode string, orderBy string) (dataToReturn interface{}, totalrecs int, recsreturned int, err error)
}

const (
	DATASERVICE_TYPE_NOSQL = "SERVICE_TYPE_NOSQL"
	DATASERVICE_TYPE_SQL   = "SERVICE_TYPE_SQL"
	DATASERVICE_TYPE_CACHE = "SERVICE_TYPE_CACHE"
)

const (
	MATCHMULTIPLEVALUES ConditionType = iota // expects first value as field name and second value as array of values
	MATCHANCESTOR                            //expects collection name and id
	FIELDVALUE                               //expects map of field values
)

const (
	InQueries Feature = iota
	Ancestors
)

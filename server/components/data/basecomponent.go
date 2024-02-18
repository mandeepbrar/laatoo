package data

import (
	"fmt"

	"laatoo.io/sdk/config"
	"laatoo.io/sdk/constants"
	"laatoo.io/sdk/server/components"
	"laatoo.io/sdk/server/core"
	"laatoo.io/sdk/server/elements"
	"laatoo.io/sdk/server/errors"
)

/*
*
Base Component helps create a new data service
*/
type BaseComponent struct {
	core.Service
	impl            DataComponent
	Object          string
	ObjectFactory   core.ObjectFactory
	ObjectConfig    *StorableConfig
	Trackable       bool
	SoftDelete      bool
	PreSave         bool
	PostSave        bool
	PostLoad        bool
	PostUpdate      bool
	Workflow        bool
	Multitenant     bool
	SoftDeleteField string
	EmbeddedSearch  bool
}

func (bc *BaseComponent) SetImpl(impl DataComponent) {
	bc.impl = impl
}

func (bc *BaseComponent) Describe(ctx core.ServerContext) error {
	bc.SetComponent(ctx, true)
	bc.AddStringConfigurations(ctx, []string{CONF_DATA_OBJECT}, nil)
	bc.AddOptionalConfigurations(ctx, map[string]string{CONF_DATA_AUDITABLE: constants.OBJECTTYPE_BOOL, CONF_DATA_POSTUPDATE: constants.OBJECTTYPE_BOOL,
		CONF_DATA_EMBEDDED_DOC_SEARCH: constants.OBJECTTYPE_BOOL, CONF_DATA_POSTSAVE: constants.OBJECTTYPE_BOOL, CONF_DATA_PRESAVE: constants.OBJECTTYPE_BOOL,
		CONF_DATA_POSTLOAD: constants.OBJECTTYPE_BOOL, CONF_DATA_MULTITENANT: constants.OBJECTTYPE_BOOL,
		CONF_DATA_WORKFLOW_ENABLED: constants.OBJECTTYPE_BOOL}, nil)
	return nil
}

func (bc *BaseComponent) Initialize(ctx core.ServerContext, conf config.Config) error {
	bc.Object, _ = bc.GetStringConfiguration(ctx, CONF_DATA_OBJECT)

	fac, ok := ctx.GetObjectFactory(bc.Object)
	if !ok {
		return errors.BadConf(ctx, CONF_DATA_OBJECT)
	} else {
		bc.ObjectFactory = fac
	}

	testObj, err := ctx.CreateObject(bc.Object)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	stor := testObj.(Storable)
	bc.ObjectConfig = stor.Config()

	trackable, ok := bc.GetBoolConfiguration(ctx, CONF_DATA_AUDITABLE)
	if ok {
		bc.Trackable = trackable
	} else {
		bc.Trackable = bc.ObjectConfig.Trackable
	}

	bc.EmbeddedSearch, _ = bc.GetBoolConfiguration(ctx, CONF_DATA_EMBEDDED_DOC_SEARCH)

	postsave, ok := bc.GetBoolConfiguration(ctx, CONF_DATA_POSTSAVE)
	if ok {
		bc.PostSave = postsave
	} else {
		bc.PostSave = bc.ObjectConfig.PostSave
	}
	postupdate, ok := bc.GetBoolConfiguration(ctx, CONF_DATA_POSTUPDATE)
	if ok {
		bc.PostUpdate = postupdate
	} else {
		bc.PostUpdate = bc.ObjectConfig.PostUpdate
	}
	presave, ok := bc.GetBoolConfiguration(ctx, CONF_DATA_PRESAVE)
	if ok {
		bc.PreSave = presave
	} else {
		bc.PreSave = bc.ObjectConfig.PreSave
	}
	postload, ok := bc.GetBoolConfiguration(ctx, CONF_DATA_POSTLOAD)
	if ok {
		bc.PostLoad = postload
	} else {
		bc.PostLoad = bc.ObjectConfig.PostLoad
	}

	multitenant, ok := bc.GetBoolConfiguration(ctx, CONF_DATA_MULTITENANT)
	if ok {
		bc.Multitenant = multitenant
	} else {
		bc.Multitenant = bc.ObjectConfig.Multitenant
	}

	workflow, ok := bc.GetBoolConfiguration(ctx, CONF_DATA_WORKFLOW_ENABLED)
	if ok {
		bc.Workflow = workflow
	} else {
		bc.Workflow = bc.ObjectConfig.Workflow
	}

	return nil
}

func (bc *BaseComponent) GetDataServiceType() string {
	return ""
}

/*
func (bc *BaseComponent) identifyStorableRefs(ctx core.ServerContext, obj interface{}) {

	//objVal := reflect.ValueOf(obj).Elem()
	objTyp := reflect.Indirect(reflect.ValueOf(obj)).Type()

	for i := 0; i < objTyp.NumField(); i++ {
		field := objTyp.Field(i)

		if strings.Contains(field.Type.String(), "data.StorableRef") {
			storableType, ok := field.Tag.Lookup("entity")
			if !ok {
			}
			bc.StorableRefs[field.Name] = storableType
		}
	}
}
*/

func (bc *BaseComponent) GetObject() string {
	return bc.Object
}

func (bc *BaseComponent) GetObjectFactory() core.ObjectFactory {
	return bc.ObjectFactory
}

func (bc *BaseComponent) GetCollection() string {
	return bc.ObjectConfig.Collection
}

// create object
func (bc *BaseComponent) CreateObject(ctx core.RequestContext) interface{} {
	return bc.ObjectFactory.CreateObject(ctx)
}

// create object collection
func (bc *BaseComponent) CreateObjectCollection(ctx core.RequestContext, len int) interface{} {
	return bc.ObjectFactory.CreateObjectCollection(ctx, len)
}

// create object collection
func (bc *BaseComponent) CreateObjectPointersCollection(ctx core.RequestContext, len int) interface{} {
	return bc.ObjectFactory.CreateObjectPointersCollection(ctx, len)
}

// supported features
func (bc *BaseComponent) Supports(Feature) bool {
	return false
}

func (bc *BaseComponent) PreProcessConditionMap(ctx core.RequestContext, operation ConditionType, args core.StringMap) core.StringMap {
	if bc.Multitenant {
		if ctx.GetUser() != nil && ctx.GetUser().GetTenant() != nil {
			args["TenantId"] = ctx.GetUser().GetTenant().GetTenantId()
		} else {
			_, ok := args["TenantId"]
			if ok {
				delete(args, "TenantId")
			}
		}
	}
	if bc.SoftDelete {
		args[bc.SoftDeleteField] = false
	}
	return args
}

// create condition for passing to data service
func (bc *BaseComponent) CreateCondition(ctx core.RequestContext, operation ConditionType, args ...interface{}) (interface{}, error) {
	return nil, errors.NotImplemented(ctx, "CreateCondition")
}

// save an object
func (bc *BaseComponent) Save(ctx core.RequestContext, item Storable) error {
	return errors.NotImplemented(ctx, "Save")
}

// adds an item to an array field
func (bc *BaseComponent) AddToArray(ctx core.RequestContext, id string, fieldName string, item interface{}) error {
	return errors.NotImplemented(ctx, "AddToArray")
}

// execute function
func (bc *BaseComponent) Execute(ctx core.RequestContext, name string, data interface{}, params core.StringMap) (interface{}, error) {
	return nil, errors.NotImplemented(ctx, "Execute")
}

// Store an object against an id
func (bc *BaseComponent) Put(ctx core.RequestContext, id string, item Storable) error {
	return errors.NotImplemented(ctx, "Put")
}

// Store multiple objects
func (bc *BaseComponent) PutMulti(ctx core.RequestContext, items []Storable) error {
	return errors.NotImplemented(ctx, "PutMulti")
}

func (bc *BaseComponent) CreateMulti(ctx core.RequestContext, items []Storable) error {
	return errors.NotImplemented(ctx, "CreateMulti")
}

// upsert an object ...insert if not there... update if there
func (bc *BaseComponent) UpsertId(ctx core.RequestContext, id string, newVals core.StringMap) error {
	return errors.NotImplemented(ctx, "UpsertId")
}

// upsert an object ...insert if not there... update if there
func (bc *BaseComponent) Upsert(ctx core.RequestContext, queryCond interface{}, newVals core.StringMap, getids bool) ([]string, error) {
	return nil, errors.NotImplemented(ctx, "Upsert")
}

// update objects by ids, fields to be updated should be provided as key value pairs
func (bc *BaseComponent) UpdateMulti(ctx core.RequestContext, ids []string, newVals core.StringMap) error {
	return errors.NotImplemented(ctx, "UpdateMulti")
}

// update an object by ids, fields to be updated should be provided as key value pairs
func (bc *BaseComponent) Update(ctx core.RequestContext, id string, newVals core.StringMap) error {
	return errors.NotImplemented(ctx, "Update")
}

// update with condition
func (bc *BaseComponent) UpdateAll(ctx core.RequestContext, queryCond interface{}, newVals core.StringMap, getids bool) ([]string, error) {
	return nil, errors.NotImplemented(ctx, "UpdateAll")
}

// Delete an object by id
func (bc *BaseComponent) Delete(ctx core.RequestContext, id string) error {
	return errors.NotImplemented(ctx, "Delete")
}

// Delete object by ids
func (bc *BaseComponent) DeleteMulti(ctx core.RequestContext, ids []string) error {
	return errors.NotImplemented(ctx, "DeleteMulti")
}

// delete with condition
func (bc *BaseComponent) DeleteAll(ctx core.RequestContext, queryCond interface{}, getids bool) ([]string, error) {
	return nil, errors.NotImplemented(ctx, "DeleteAll")
}

// Get an object by id
func (bc *BaseComponent) GetById(ctx core.RequestContext, id string) (Storable, error) {
	return nil, errors.NotImplemented(ctx, "GetById")
}

// get storables in a hashtable
func (bc *BaseComponent) GetMultiHash(ctx core.RequestContext, ids []string) (map[string]Storable, error) {
	return nil, errors.NotImplemented(ctx, "GetMultiHash")
}

// Get multiple objects by id
func (bc *BaseComponent) GetMulti(ctx core.RequestContext, ids []string, orderBy interface{}) ([]Storable, error) {
	return nil, errors.NotImplemented(ctx, "GetMulti")
}

// Get all object with given conditions
func (bc *BaseComponent) Get(ctx core.RequestContext, queryCond interface{}, pageSize int, pageNum int, mode string, orderBy interface{}) (dataToReturn []Storable, ids []string, totalrecs int, recsreturned int, err error) {
	return nil, nil, -1, -1, errors.NotImplemented(ctx, "Get")
}

// Get one record satisfying condition
func (bc *BaseComponent) GetOne(ctx core.RequestContext, queryCond interface{}) (dataToReturn Storable, err error) {
	if bc.impl != nil {
		data, _, _, recs, err := bc.impl.Get(ctx, queryCond, -1, -1, "", nil)
		if err != nil {
			return nil, err
		}
		if recs < 1 {
			return nil, errors.NotFound(ctx, "Data Record")
		}
		return data[0], nil
	}
	return nil, errors.NotImplemented(ctx, "GetOne")
}

// Get a list of all items
func (bc *BaseComponent) GetList(ctx core.RequestContext, pageSize int, pageNum int, mode string, orderBy interface{}) (dataToReturn []Storable, ids []string, totalrecs int, recsreturned int, err error) {
	return nil, nil, -1, -1, errors.NotImplemented(ctx, "GetList")

}

func (bc *BaseComponent) Count(ctx core.RequestContext, queryCond interface{}) (count int, err error) {
	return -1, errors.NotImplemented(ctx, "Count")
}

func (bc *BaseComponent) CountGroups(ctx core.RequestContext, queryCond interface{}, groupids []string, group string) (res core.StringMap, err error) {
	return nil, errors.NotImplemented(ctx, "CountGroups")
}

func (bc *BaseComponent) CreateDBCollection(ctx core.ServerContext) error {
	return errors.NotImplemented(ctx, "CreateDBCollection")
}

func (bc *BaseComponent) DropDBCollection(ctx core.ServerContext) error {
	return errors.NotImplemented(ctx, "DropDBCollection")
}

func (bc *BaseComponent) DBCollectionExists(ctx core.ServerContext) (bool, error) {
	return false, errors.NotImplemented(ctx, "DBCollectionExists")
}

func (bc *BaseComponent) Transaction(ctx core.RequestContext, callback func(cx core.RequestContext) error) error {
	return errors.NotImplemented(ctx, "StartTransaction")
}

// Gets the value of a key
func (bc *BaseComponent) GetValue(ctx core.RequestContext, key string) (interface{}, error) {
	return nil, errors.NotImplemented(ctx, "GetValue")
	//	return bc.PluginDataComponent.GetById(ctx, key)
}

// Puts the value of a key
func (bc *BaseComponent) PutValue(ctx core.RequestContext, key string, value interface{}) error {
	return errors.NotImplemented(ctx, "PutValue")
}

// Deletes the key
func (bc *BaseComponent) DeleteValue(ctx core.RequestContext, key string) error {
	return errors.NotImplemented(ctx, "DeleteValue")
	//	return bc.PluginDataComponent.Delete(ctx, key)
}

func (bc *BaseComponent) ValidateTenant(ctx core.RequestContext, stor Storable) bool {
	if !bc.Multitenant {
		return true
	}
	if stor != nil {
		mt, ok := stor.(Multitenant)
		if ok && mt.GetTenantInfo().GetTenantId() == ctx.GetUser().GetTenant().GetTenantId() {
			return true
		}
	}
	return false
}

// supported features
func (bc *BaseComponent) StartWorkflow(ctx core.RequestContext, stor Storable, insConf core.StringMap) (components.WorkflowInstance, error) {
	if !bc.Workflow {
		return nil, nil
	}
	workflowManager := ctx.GetServerElement(core.ServerElementWorkflowManager).(elements.WorkflowManager)
	workflowName := fmt.Sprintf("%s_workflow", bc.Object)
	if workflowManager.IsWorkflowRegistered(ctx.ServerContext(), workflowName) {
		ins, err := ctx.StartWorkflow(workflowName, core.StringMap{"data": stor}, insConf)
		if err != nil {
			return nil, errors.WrapError(ctx, err)
		}
		return ins.(components.WorkflowInstance), nil

	}
	return nil, nil
}

/*
//start a transaction.
func (bc *BaseComponent) StartTransaction(ctx core.RequestContext) error {
	return errors.NotImplemented(ctx, "StartTransaction")
}

//commit a transaction
func (bc *BaseComponent) CommitTransaction(ctx core.RequestContext) error {
	return errors.NotImplemented(ctx, "CommitTransaction")
}

//rollback a transaction
func (bc *BaseComponent) RollbackTransaction(ctx core.RequestContext) error {
	return errors.NotImplemented(ctx, "RollbackTransaction")
}
*/

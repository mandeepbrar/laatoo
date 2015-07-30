package service

//Service interface that needs to be implemented by any service of a system
type Service interface {
	//Provides the name of the service
	GetName() string
	//Initialize the service. Consumer of a service passes the data
	Initialize(ctx ServiceContext) error
	//The service starts serving when this method is called
	Serve() error
	//Type of service
	GetServiceType() string
}

//service context object for initializing services
type ServiceContext interface {
	GetService(alias string) (Service, error)
	CreateObject(objName string, confData map[string]interface{}) (interface{}, error)
	CreateEmptyObject(objName string) (interface{}, error)
	CreateCollection(objName string) (interface{}, error)
}

const (
	SERVICE_TYPE_APP  = "SERVICE_TYPE_APP"
	SERVICE_TYPE_DATA = "SERVICE_TYPE_DATA"
	SERVICE_TYPE_WEB  = "SERVICE_TYPE_WEB"
)

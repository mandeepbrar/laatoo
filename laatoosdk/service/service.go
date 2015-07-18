package service

//Service interface that needs to be implemented by any service of a system
type Service interface {
	//Provides the name of the service
	GetName() string
	//Provides the alias of the service
	GetAlias() string
	//Initialize the service. Consumer of a service passes the data
	Initialize(ctx interface{}) error
	//The service starts serving when this method is called
	Serve() error
	//Type of service
	GetServiceType() string
}

const (
	SERVICE_TYPE_APP  = "SERVICE_TYPE_APP"
	SERVICE_TYPE_DATA = "SERVICE_TYPE_DATA"
	SERVICE_TYPE_WEB  = "SERVICE_TYPE_WEB"
)

package core

//Service interface that needs to be implemented by any service of a system
type Service interface {
	//Provides the name of the service
	GetName() string
	//Initialize the service. Consumer of a service passes the data
	Initialize(ctx Context) error
	//The service starts serving when this method is called
	//called on first request
	Serve(ctx Context) error
	//Type of service
	GetServiceType() string
	//Execute method
	Execute(Context, string, map[string]interface{}) (interface{}, error)
}

const (
	SERVICE_TYPE_APP  = "SERVICE_TYPE_APP"
	SERVICE_TYPE_DATA = "SERVICE_TYPE_DATA"
	SERVICE_TYPE_WEB  = "SERVICE_TYPE_WEB"
)

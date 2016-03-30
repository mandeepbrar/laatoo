package core

type ServiceStatus int

const (
	StatusSuccess ServiceStatus = iota
	StatusServeFile
	StatusServeBytes
	StatusUnauthorized
	StatusNotFound
	StatusRedirect
)

type ServiceResponse struct {
	Status ServiceStatus
	Data   interface{}
}

func NewServiceResponse(status ServiceStatus, data interface{}) *ServiceResponse {
	return &ServiceResponse{status, data}
}

package core

const (
	StatusSuccess      = 1
	StatusServeFile    = 2
	StatusServeBytes   = 3
	StatusUnauthorized = 4
	StatusNotFound     = 5
	StatusRedirect     = 6
	StatusNotModified  = 7
)

/***Header****/
const (
	ContentType  = "Content-Type"
	LastModified = "Last-Modified"
)

type ServiceResponse struct {
	Status int
	Data   interface{}
	Info   map[string]interface{}
}

func NewServiceResponse(status int, data interface{}, info map[string]interface{}) *ServiceResponse {
	return &ServiceResponse{status, data, info}
}

var (
	StatusUnauthorizedResponse = NewServiceResponse(StatusUnauthorized, nil, nil)
	StatusNotFoundResponse     = NewServiceResponse(StatusNotFound, nil, nil)
	StatusNotModifiedResponse  = NewServiceResponse(StatusNotModified, nil, nil)
)

type ServiceResponseHandler interface {
	HandleResponse(ctx RequestContext, res *ServiceResponse) error
}

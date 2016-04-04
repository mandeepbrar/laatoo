package core

const (
	StatusSuccess      = 200
	StatusServeFile    = 201
	StatusServeBytes   = 202
	StatusUnauthorized = 401
	StatusNotFound     = 404
	StatusRedirect     = 301
	StatusNotModified  = 305
	StatusBadRequest   = 400
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
	StatusBadRequestResponse   = NewServiceResponse(StatusBadRequest, nil, nil)
	StatusNotModifiedResponse  = NewServiceResponse(StatusNotModified, nil, nil)
)

type ServiceResponseHandler interface {
	HandleResponse(ctx RequestContext) error
}

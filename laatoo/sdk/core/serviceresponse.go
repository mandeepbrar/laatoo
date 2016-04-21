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
	Return bool
}

func NewServiceResponse(status int, data interface{}, info map[string]interface{}) *ServiceResponse {
	return newServiceResponse(status, data, info, false)
}
func newServiceResponse(status int, data interface{}, info map[string]interface{}, ReturnVal bool) *ServiceResponse {
	return &ServiceResponse{status, data, info, ReturnVal}
}

var (
	StatusUnauthorizedResponse = newServiceResponse(StatusUnauthorized, nil, nil, true)
	StatusNotFoundResponse     = newServiceResponse(StatusNotFound, nil, nil, true)
	StatusBadRequestResponse   = newServiceResponse(StatusBadRequest, nil, nil, true)
	StatusNotModifiedResponse  = newServiceResponse(StatusNotModified, nil, nil, true)
)
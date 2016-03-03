package core

type HandlerFunc func(ctx Context) error

type Router interface {
	//Get a sub router
	Group(ctx Context, path string, conf map[string]interface{}) Router
	//Use middleware
	Use(ctx Context, handler interface{})
	Get(ctx Context, path string, conf map[string]interface{}, handler HandlerFunc) error
	Post(ctx Context, path string, conf map[string]interface{}, handler HandlerFunc) error
	Put(ctx Context, path string, conf map[string]interface{}, handler HandlerFunc) error
	Delete(ctx Context, path string, conf map[string]interface{}, handler HandlerFunc) error
	Static(ctx Context, path string, conf map[string]interface{}, dir string) error
	ServeFile(ctx Context, pagePath string, conf map[string]interface{}, dest string)
}

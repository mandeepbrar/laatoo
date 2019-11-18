package gaetasks

import (
	"encoding/json"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"

	"google.golang.org/appengine/taskqueue"
)

const (
	GAE_PATH = "path"
)

type GaeProducer struct {
	core.Service
	path       string
	authHeader string
}

func (svc *GaeProducer) Describe(ctx core.ServerContext) error {
	svc.SetDescription(ctx, "GAE task service producer component")
	svc.AddStringConfigurations(ctx, []string{GAE_PATH}, nil)
	svc.SetComponent(ctx, true)
	return nil
}
func (svc *GaeProducer) Initialize(ctx core.ServerContext, conf config.Config) error {
	sh := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sh != nil {
		shandler := sh.(elements.SecurityHandler)
		ah := shandler.GetProperty(config.AUTHHEADER)
		if ah == nil {
			return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
		}
		svc.authHeader = ah.(string)
	} else {
		return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
	}

	svc.path, _ = svc.GetStringConfiguration(ctx, GAE_PATH)
	return nil
}

func (svc *GaeProducer) PushTask(ctx core.RequestContext, queue string, t *components.Task) error {
	appEngineContext := ctx.GetAppengineContext()
	bytes, err := json.Marshal(t)
	if err != nil {
		return err
	}
	gaetask := taskqueue.NewPOSTTask(svc.path, map[string][]string{})
	gaetask.Payload = bytes
	_, err = taskqueue.Add(appEngineContext, gaetask, "")
	return err
}

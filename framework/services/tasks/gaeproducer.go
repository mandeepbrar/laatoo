package tasks

import (
	"encoding/json"
	"google.golang.org/appengine/taskqueue"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	//	"laatoo/sdk/log"
)

const (
	GAE_PATH = "path"
)

type gaeProducer struct {
	path string
}

func (svc *gaeProducer) Initialize(ctx core.ServerContext, conf config.Config) error {
	path, ok := conf.GetString(GAE_PATH)
	if !ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", GAE_PATH)
	}
	svc.path = path
	return nil
}

func (svc *gaeProducer) PushTask(ctx core.RequestContext, queue string, taskData interface{}) error {
	appEngineContext := ctx.GetAppengineContext()
	data, err := json.Marshal(taskData)
	if err != nil {
		return err
	}
	t := task{Queue: queue, Data: data, User: ctx.GetUser().GetId()}
	bytes, err := json.Marshal(t)
	if err != nil {
		return err
	}
	gaetask := taskqueue.NewPOSTTask(svc.path, map[string][]string{})
	gaetask.Payload = bytes
	_, err = taskqueue.Add(appEngineContext, gaetask, "")
	return err
}

func (svc *gaeProducer) Invoke(ctx core.RequestContext) error {
	return nil
}

func (svc *gaeProducer) Start(ctx core.ServerContext) error {
	return nil
}

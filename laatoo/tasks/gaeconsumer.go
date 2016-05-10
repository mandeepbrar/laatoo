package tasks

import (
	"encoding/json"
	//	"google.golang.org/appengine/taskqueue"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
)

type gaeConsumer struct {
	queues map[string]*taskQueue
}

func (svc *gaeConsumer) Initialize(ctx core.ServerContext, conf config.Config) error {
	queuesConf, ok := conf.GetSubConfig(config.CONF_TASK_QUEUES)
	if ok {
		queueNames := queuesConf.AllConfigurations()
		for _, queueName := range queueNames {
			qCtx := ctx.SubContext("Creating Queue" + queueName)
			queueProcessorName, _ := queuesConf.GetString(queueName)
			processor, err := qCtx.GetService(queueProcessorName)
			if err != nil {
				return errors.WrapError(qCtx, err)
			}
			svc.createQueue(qCtx, queueName, processor)
		}
	} else {
		return errors.ThrowError(ctx, errors.CORE_ERROR_MISSING_CONF, "Conf", config.CONF_TASK_QUEUES)
	}
	return nil
}

func (svc *gaeConsumer) createQueue(ctx core.ServerContext, queue string, lstnr core.Service) error {
	_, ok := svc.queues[queue]
	if ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "queuename", queue)
	}
	tq := &taskQueue{name: queue, qRef: nil, lstnr: lstnr}
	svc.queues[queue] = tq
	return nil
}

func (svc *gaeConsumer) Invoke(ctx core.RequestContext) error {
	//gae header... if an outside request comes, this header would not be there.. gae will remove it
	//_, ok := ctx.GetString("X-AppEngine-TaskName")
	//if ok {
	bytes := ctx.GetRequest().([]byte)
	t := &task{}
	err := json.Unmarshal(bytes, t)
	if err != nil {
		log.Logger.Error(ctx, "Error in background process", "job", string(bytes), "err", err)
		return err
	} else {
		log.Logger.Debug(ctx, "Received job", "task", t)
		req := ctx.SubContext("Task")
		req.SetRequest(t.Data)
		req.Set("User", t.User)
		queueName := t.Queue
		log.Logger.Debug(ctx, "Received job", "task", t, "queue", queueName)
		q, ok := svc.queues[queueName]
		if ok {
			err := q.lstnr.Invoke(req)
			if err != nil {
				log.Logger.Error(ctx, "Error in background process", "err", err)
				return err
			}
		}
	}
	ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, nil, nil))
	/*} else {
		log.Logger.Error(ctx, "Non gae requests to background processor")
		ctx.SetResponse(core.StatusBadRequestResponse)
	}*/
	return nil
}

func (svc *gaeConsumer) Start(ctx core.ServerContext) error {
	return nil
}

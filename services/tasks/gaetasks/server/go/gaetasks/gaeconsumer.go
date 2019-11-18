package gaetasks

import (
	"encoding/json"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/log"
)

type taskQueue struct {
	name  string
	qRef  interface{}
	lstnr core.Service
}
type task struct {
	Queue string
	Data  []byte
	Token string
}

type GaeConsumer struct {
	core.Service
	queues      map[string]*taskQueue
	authHeader  string
	shandler    elements.SecurityHandler
	taskManager elements.TaskManager
}

func (svc *GaeConsumer) Describe(ctx core.ServerContext) error {
	svc.queues = make(map[string]*taskQueue, 10)
	svc.SetDescription(ctx, "GAE task service consumer component")
	return svc.AddParamWithType(ctx, "task", config.OBJECTTYPE_BYTES)
	//svc.SetRequestType(ctx, config.OBJECTTYPE_BYTES, false, false)
}

func (svc *GaeConsumer) Initialize(ctx core.ServerContext, conf config.Config) error {
	svc.taskManager = ctx.GetServerElement(core.ServerElementTaskManager).(elements.TaskManager)
	return nil
}

/*
func (svc *gaeConsumer) createQueue(ctx core.ServerContext, queue string, lstnr core.Service) error {
	_, ok := svc.queues[queue]
	if ok {
		return errors.ThrowError(ctx, errors.CORE_ERROR_BAD_ARG, "queuename", queue)
	}
	tq := &taskQueue{name: queue, qRef: nil, lstnr: lstnr}
	svc.queues[queue] = tq
	return nil
}*/

func (svc *GaeConsumer) Invoke(ctx core.RequestContext) error {
	//gae header... if an outside request comes, this header would not be there.. gae will remove it
	//_, ok := ctx.GetString("X-AppEngine-TaskName")
	//if ok {
	val, _ := ctx.GetParamValue("task")
	bytes := val.([]byte)
	t := &components.Task{}
	err := json.Unmarshal(bytes, t)
	if err != nil {
		log.Error(ctx, "Error in background process", "job", string(bytes), "err", err)
		return err
	} else {
		return svc.taskManager.ProcessTask(ctx, t)

		/*
			req := ctx.SubContext("Gae background task " + t.Queue)
			log.Debug(req, "Received job ")
			req.SetRequest(t.Data)
			req.Set(svc.authHeader, t.Token)
			svc.shandler.AuthenticateRequest(req)
			queueName := t.Queue
			q, ok := svc.queues[queueName]
			if ok {
				err := q.lstnr.Invoke(req)
				if err != nil {
					log.Error(req, "Error in background process", "err", err)
					return err
				}
			}*/
	}
	//ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, nil, nil))
	/*} else {
		log.Error(ctx, "Non gae requests to background processor")
		ctx.SetResponse(core.StatusBadRequestResponse)
	}*/
}
func (svc *GaeConsumer) SubsribeQueue(ctx core.ServerContext, queue string) error {
	return nil
}

func (svc *GaeConsumer) UnsubsribeQueue(ctx core.ServerContext, queue string) error {
	return nil
}

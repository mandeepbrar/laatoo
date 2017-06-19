package main

import (
	"encoding/json"
	"laatoo/sdk/components"
	//	"google.golang.org/appengine/taskqueue"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/log"
	"laatoo/sdk/server"
)

const (
	CONF_TASKS_GAE_PRODUCER = "gaetaskpublisher"
	CONF_TASKS_GAE_CONSUMER = "gaetaskprocessor"
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

type gaeConsumer struct {
	queues      map[string]*taskQueue
	authHeader  string
	shandler    server.SecurityHandler
	taskManager server.TaskManager
}

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_TASKS_GAE_PRODUCER, Object: gaeConsumer{}},
		core.PluginComponent{Name: CONF_TASKS_GAE_CONSUMER, Object: gaeProducer{}}}
}

func (svc *gaeConsumer) Initialize(ctx core.ServerContext, conf config.Config) error {
	/*queuesConf, ok := conf.GetSubConfig(config.CONF_TASK_QUEUES)
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

	sh := ctx.GetServerElement(core.ServerElementSecurityHandler)
	if sh != nil {
		shandler := sh.(server.SecurityHandler)
		svc.shandler = shandler
		ah, ok := shandler.GetString(config.AUTHHEADER)
		if !ok {
			return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
		}
		svc.authHeader = ah
	} else {
		return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Resource", config.AUTHHEADER)
	}

	*/
	svc.taskManager = ctx.GetServerElement(core.ServerElementTaskManager).(server.TaskManager)
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

func (svc *gaeConsumer) Invoke(ctx core.RequestContext) error {
	//gae header... if an outside request comes, this header would not be there.. gae will remove it
	//_, ok := ctx.GetString("X-AppEngine-TaskName")
	//if ok {
	bytes := ctx.GetRequest().([]byte)
	t := &components.Task{}
	err := json.Unmarshal(bytes, t)
	if err != nil {
		log.Logger.Error(ctx, "Error in background process", "job", string(bytes), "err", err)
		return err
	} else {
		return svc.taskManager.ProcessTask(ctx, t)

		/*
			req := ctx.SubContext("Gae background task " + t.Queue)
			log.Logger.Debug(req, "Received job ")
			req.SetRequest(t.Data)
			req.Set(svc.authHeader, t.Token)
			svc.shandler.AuthenticateRequest(req)
			queueName := t.Queue
			q, ok := svc.queues[queueName]
			if ok {
				err := q.lstnr.Invoke(req)
				if err != nil {
					log.Logger.Error(req, "Error in background process", "err", err)
					return err
				}
			}*/
	}
	//ctx.SetResponse(core.NewServiceResponse(core.StatusSuccess, nil, nil))
	/*} else {
		log.Logger.Error(ctx, "Non gae requests to background processor")
		ctx.SetResponse(core.StatusBadRequestResponse)
	}*/
	return nil
}
func (svc *gaeConsumer) SubsribeQueue(ctx core.ServerContext, queue string) error {
	return nil
}
func (svc *gaeConsumer) Start(ctx core.ServerContext) error {
	return nil
}

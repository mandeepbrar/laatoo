package main

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"

	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/worker"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	CONF_CADENCE_WORKER            = "cadence_worker"
	CONF_CADENCE_SIMPLEWORKFLOW    = "cadence_simpleworkflow"
	CONF_CADENCE_WORKFLOWINITIATOR = "cadence_workflowinitiator"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_CADENCE_WORKER, Object: CadenceWorkerService{}},
		core.PluginComponent{Name: CONF_CADENCE_SIMPLEWORKFLOW, Object: SimpleWorkflow{}},
		core.PluginComponent{Name: CONF_CADENCE_WORKFLOWINITIATOR, Object: WorkflowInitiator{}},
	}
}

/*var HostPort = "127.0.0.1:7933"
var Domain = "SimpleDomain"
var TaskListName = "SimpleWorker"
var ClientName = "SimpleWorker"
var CadenceService = "cadence-frontend"*/

type CadenceWorkerService struct {
	core.Service
	Host           string
	Domain         string
	TaskListName   string
	ClientName     string
	CadenceService string
}

func (svc *CadenceWorkerService) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

func (svc *CadenceWorkerService) buildCadenceClient(ctx core.ServerContext) workflowserviceclient.Interface {
	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(svc.ClientName))
	if err != nil {
		panic("Failed to setup tchannel")
	}
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name:      svc.ClientName,
		Outbounds: yarpc.Outbounds{svc.CadenceService: {Unary: ch.NewSingleOutbound(svc.Host)}},
	})
	if err := dispatcher.Start(); err != nil {
		panic("Failed to start dispatcher")
	}

	return workflowserviceclient.New(dispatcher.ClientConfig(svc.CadenceService))
}

func buildLogger() (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zapcore.InfoLevel)

	var err error
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func (svc *CadenceWorkerService) Start(ctx core.ServerContext) error {
	svc.Host, _ = svc.GetStringConfiguration(ctx, "Host")
	svc.Domain, _ = svc.GetStringConfiguration(ctx, "Domain")
	svc.TaskListName, _ = svc.GetStringConfiguration(ctx, "TaskListName")
	svc.ClientName, _ = svc.GetStringConfiguration(ctx, "ClientName")
	svc.CadenceService, _ = svc.GetStringConfiguration(ctx, "CadenceService")

	log.Error(ctx, "config ", "svc", svc, "svc host", svc.Host)

	logger, err := buildLogger()
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	worker := worker.New(
		svc.buildCadenceClient(ctx),
		svc.Domain,
		svc.TaskListName,
		worker.Options{
			Logger: logger,
		})

	err = worker.Start()
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	log.Info(ctx, "Started Worker", "worker", svc.TaskListName)
	return nil
}
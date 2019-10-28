package main

import (
	"context"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"time"

	"github.com/twinj/uuid"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
)

type WorkflowInitiator struct {
	core.Service
	Host           string
	Domain         string
	ClientName     string
	CadenceService string
	dispatcher     *yarpc.Dispatcher
	client         client.Client
}

func (svc *WorkflowInitiator) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

func (svc *WorkflowInitiator) Start(ctx core.ServerContext) error {
	workflowClient, err := svc.buildCadenceClient(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	svc.client = workflowClient
	return nil
}

func (svc *WorkflowInitiator) StartWorkflow(ctx core.RequestContext, workflowName string, initVal interface{}) error {
	return svc.StartWorkflowOnTasklist(ctx, workflowName, workflowName, initVal)
}

func (svc *WorkflowInitiator) StartWorkflowOnTasklist(ctx core.RequestContext, workflowName, tasklistName string, initVal interface{}) error {
	workflowoptions := client.StartWorkflowOptions{
		ID:                              workflowName + "_" + uuid.NewV1().String(),
		TaskList:                        tasklistName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	ctx = ctx.SubContext("Workflow" + workflowName)
	we, err := svc.client.StartWorkflow(context.Background(), workflowoptions, workflowName, initVal)
	if err != nil {
		log.Error(ctx, "Failed to create workflow", "workflow", workflowName)
		return errors.WrapError(ctx, err)
	} else {
		log.Info(ctx, "Started Workflow", "WorkflowID", we.ID, "RunID", we.RunID)
	}
	return nil
}

// BuildCadenceClient builds a client to cadence service
func (svc *WorkflowInitiator) buildCadenceClient(ctx core.ServerContext) (client.Client, error) {
	service, err := svc.buildServiceClient(ctx)
	if err != nil {
		return nil, err
	}

	return client.NewClient(
		service, svc.Domain, &client.Options{}), nil //Identity: b.clientIdentity, MetricsScope: b.metricsScope, DataConverter: b.dataConverter, ContextPropagators: b.ctxProps

}

// BuildServiceClient builds a rpc service client to cadence service
func (svc *WorkflowInitiator) buildServiceClient(ctx core.ServerContext) (workflowserviceclient.Interface, error) {
	if err := svc.build(ctx); err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	if svc.dispatcher == nil {
		log.Error(ctx, "No RPC dispatcher provided to create a connection to Cadence Service")
	}

	return workflowserviceclient.New(svc.dispatcher.ClientConfig(svc.CadenceService)), nil
}

func (svc *WorkflowInitiator) build(ctx core.ServerContext) error {
	if svc.dispatcher != nil {
		return nil
	}

	ch, err := tchannel.NewChannelTransport(
		tchannel.ServiceName(svc.ClientName))
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	log.Debug(ctx, "Creating RPC dispatcher outbound", "ServiceName", svc.CadenceService, "Host", svc.Host)

	svc.dispatcher = yarpc.NewDispatcher(yarpc.Config{
		Name: svc.ClientName,
		Outbounds: yarpc.Outbounds{
			svc.CadenceService: {Unary: ch.NewSingleOutbound(svc.Host)},
		},
	})

	if svc.dispatcher != nil {
		if err := svc.dispatcher.Start(); err != nil {
			log.Error(ctx, "Failed to create outbound transport channel: %v", "err", err)
			return errors.WrapError(ctx, err)
		}
	}

	return nil
}

/*
func (h *SampleHelper) QueryWorkflow(workflowID, runID, queryType string, args ...interface{}) {
	workflowClient, err := h.Builder.BuildCadenceClient()
	if err != nil {
		h.Logger.Error("Failed to build cadence client.", zap.Error(err))
		panic(err)
	}

	resp, err := workflowClient.QueryWorkflow(context.Background(), workflowID, runID, queryType, args...)
	if err != nil {
		h.Logger.Error("Failed to query workflow", zap.Error(err))
		panic("Failed to query workflow.")
	}
	var result interface{}
	if err := resp.Get(&result); err != nil {
		h.Logger.Error("Failed to decode query result", zap.Error(err))
	}
	h.Logger.Info("Received query result", zap.Any("Result", result))
}

func (h *SampleHelper) SignalWorkflow(workflowID, signal string, data interface{}) {
	workflowClient, err := h.Builder.BuildCadenceClient()
	if err != nil {
		h.Logger.Error("Failed to build cadence client.", zap.Error(err))
		panic(err)
	}

	err = workflowClient.SignalWorkflow(context.Background(), workflowID, "", signal, data)
	if err != nil {
		h.Logger.Error("Failed to signal workflow", zap.Error(err))
		panic("Failed to signal workflow.")
	}
}

func (h *SampleHelper) CancelWorkflow(workflowID string) {
	workflowClient, err := h.Builder.BuildCadenceClient()
	if err != nil {
		h.Logger.Error("Failed to build cadence client.", zap.Error(err))
		panic(err)
	}

	err = workflowClient.CancelWorkflow(context.Background(), workflowID, "")
	if err != nil {
		h.Logger.Error("Failed to cancel workflow", zap.Error(err))
		panic("Failed to cancel workflow.")
	}
}
*/

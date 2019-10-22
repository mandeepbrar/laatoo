package main

import (
	"context"
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"time"

	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

type input struct {
	initVal interface{}
	prevVal interface{}
}

type SimpleActivity func(ctx context.Context, value input) (interface{}, error)

type SimpleWorkflow struct {
	core.Service
	TaskListName string
	Tasks        []string
	activities   []SimpleActivity
}

func (svc *SimpleWorkflow) Initialize(ctx core.ServerContext, conf config.Config) error {
	return nil
}

func (svc *SimpleWorkflow) Start(ctx core.ServerContext) error {
	svc.TaskListName, _ = svc.GetStringConfiguration(ctx, "TaskListName")
	svc.Tasks, _ = svc.GetStringArrayConfiguration(ctx, "Tasks")
	workflow.Register(svc.Workflow)
	return nil
}

func actCreator(taskName string) SimpleActivity {
	return func(actctx context.Context, value input) (interface{}, error) {
		activity.GetLogger(actctx).Info("Done", zap.String("taskName", taskName))
		return nil, nil
	}
}

func (svc *SimpleWorkflow) createActivities(ctx core.ServerContext) error {
	svc.activities = make([]SimpleActivity, len(svc.Tasks))
	for idx, taskName := range svc.Tasks {
		act := actCreator(taskName)
		activity.Register(act)
		svc.activities[idx] = act
	}
	return nil
}

func (svc *SimpleWorkflow) Workflow(ctx workflow.Context, value interface{}) error {
	ao := workflow.ActivityOptions{
		TaskList:               svc.TaskListName,
		ScheduleToCloseTimeout: time.Second * 60,
		ScheduleToStartTimeout: time.Second * 60,
		StartToCloseTimeout:    time.Second * 60,
		HeartbeatTimeout:       time.Second * 10,
		WaitForCancellation:    false,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	actLen := len(svc.activities)
	var prevVal interface{}
	for idx := 0; idx < actLen; idx++ {
		act := svc.activities[idx]
		inp := input{initVal: value, prevVal: prevVal}
		future := workflow.ExecuteActivity(ctx, act, inp)
		if err := future.Get(ctx, &prevVal); err != nil {
			return err
		}
		workflow.GetLogger(ctx).Info("Done", zap.String("result", fmt.Sprintf("%s", prevVal)))
	}
	return nil
}

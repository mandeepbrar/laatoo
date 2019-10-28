package main

import (
	"encoding/json"
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"time"

	"laatoo/sdk/modules/cadence"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

type SimpleWorkflow struct {
	core.Service
	taskManager  elements.TaskManager
	TaskListName string
	WorkflowName string
	Tasks        []string
}

func (svc *SimpleWorkflow) Initialize(ctx core.ServerContext, conf config.Config) error {
	svc.taskManager = ctx.GetServerElement(core.ServerElementTaskManager).(elements.TaskManager)
	err := svc.createActivities(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if svc.TaskListName == "" {
		svc.TaskListName = svc.WorkflowName
	}
	cadence.RegisterWorkflow(ctx, svc.WorkflowName, svc.Workflow)
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
	actLen := len(svc.Tasks)
	var prevVal = value

	for idx := 0; idx < actLen; idx++ {

		taskName := svc.Tasks[idx]
		workflow.GetLogger(ctx).Info("Going to execute activity", zap.String("task", fmt.Sprintf("%s", taskName)))

		var data []byte
		var err error
		if prevVal != nil {
			data, err = json.Marshal(prevVal)
			if err != nil {
				return err
			}
		}

		tsk := &components.Task{Queue: taskName, Data: data, Token: ""}

		future := workflow.ExecuteActivity(ctx, taskName, tsk)

		var result interface{}
		if err := future.Get(ctx, &result); err != nil {
			return err
		}
		workflow.GetLogger(ctx).Info("Executed activity", zap.String("task", fmt.Sprintf("%s", taskName)), zap.String("result", fmt.Sprintf("%s", result)))
		prevVal = result

	}
	return nil
}

func (svc *SimpleWorkflow) createActivities(ctx core.ServerContext) error {
	log.Info(ctx, "creating activities", "tasks", svc.Tasks)

	for _, taskName := range svc.Tasks {
		taskActivityCreator(ctx, taskName, svc.WorkflowName, svc.taskManager)
		log.Info(ctx, "registered activity", "taskname", taskName)
	}

	return nil
}

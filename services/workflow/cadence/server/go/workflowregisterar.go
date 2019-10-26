package main

import (
	"context"
	"encoding/json"
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"reflect"
	"time"

	"laatoo/sdk/modules/cadence"

	"go.uber.org/cadence/activity"
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
	return nil
}

func (svc *SimpleWorkflow) Start(ctx core.ServerContext) error {
	svc.taskManager = ctx.GetServerElement(core.ServerElementTaskManager).(elements.TaskManager)
	svc.TaskListName, _ = svc.GetStringConfiguration(ctx, "TaskListName")
	svc.WorkflowName, _ = svc.GetStringConfiguration(ctx, "WorkflowName")
	svc.Tasks, _ = svc.GetStringArrayConfiguration(ctx, "Tasks")
	err := svc.createActivities(ctx)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	svc.RegisterWorkflow(ctx, svc.WorkflowName, svc.Workflow)
	return nil
}

func (svc *SimpleWorkflow) RegisterWorkflow(ctx core.ServerContext, name string, workflowToRegister cadence.CadenceWorkflow) {
	workflow.RegisterWithOptions(workflowToRegister, workflow.RegisterOptions{Name: svc.WorkflowName})
}

func (svc *SimpleWorkflow) taskActivityCreator(ctx core.ServerContext, taskName string) {
	log.Info(ctx, "Creating workflow task", "taskName", taskName)

	myFunc := func(args []reflect.Value) (results []reflect.Value) {

		reqctx := args[0].Interface().(core.RequestContext)
		actCtxInterface, _ := reqctx.Get("WorkflowCtx")
		actctx := actCtxInterface.(context.Context)

		inputTask := args[1].Interface().(*components.Task)

		log.Info(ctx, "Posting workflow task", "input", inputTask)

		var err error

		var result interface{}

		err = svc.taskManager.ProcessTask(reqctx, inputTask)
		if err == nil {
			res := reqctx.GetResponse()
			activity.GetLogger(actctx).Info("Task complete", zap.String("taskName", inputTask.Queue))
			result = res.Data
		}

		var resVal reflect.Value
		if result != nil {
			resVal = reflect.ValueOf(result)
		} else {
			resVal = reflect.Zero(reflect.TypeOf((*interface{})(nil)).Elem())
		}

		var errVal reflect.Value
		if err != nil {
			errVal = reflect.ValueOf(err)
		} else {
			errVal = reflect.Zero(reflect.TypeOf((*error)(nil)).Elem())
		}

		return []reflect.Value{resVal, errVal}
	}

	var k components.TaskProcessor
	act := reflect.MakeFunc(reflect.TypeOf(k), myFunc).Interface().(components.TaskProcessor)

	activity.RegisterWithOptions(act, activity.RegisterOptions{Name: taskName})
}

func (svc *SimpleWorkflow) createActivities(ctx core.ServerContext) error {
	log.Info(ctx, "creating activities", "tasks", svc.Tasks)

	for _, taskName := range svc.Tasks {
		svc.taskActivityCreator(ctx, taskName)
		log.Info(ctx, "registered activity", "taskname", taskName)
	}

	return nil
}

func (svc *SimpleWorkflow) Workflow(ctx workflow.Context, reqCtx core.RequestContext, value interface{}) error {
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

		reqCtx.Set("WorkflowCtx", ctx)

		future := workflow.ExecuteActivity(ctx, taskName, reqCtx, tsk)

		var result interface{}
		if err := future.Get(ctx, &result); err != nil {
			return err
		}
		workflow.GetLogger(ctx).Info("Executed activity", zap.String("task", fmt.Sprintf("%s", taskName)), zap.String("result", fmt.Sprintf("%s", result)))
		prevVal = result

	}
	return nil
}

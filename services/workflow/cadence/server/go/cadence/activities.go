package cadence

import (
	"context"
	"laatoo/sdk/modules/cadence"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/log"
	"reflect"

	"go.uber.org/cadence/activity"
	"go.uber.org/zap"
)

func taskActivityCreator(ctx core.ServerContext, taskName string, workflowName string, taskManager elements.TaskManager) {
	log.Info(ctx, "Creating workflow task", "taskName", taskName)
	errType := reflect.TypeOf((*error)(nil)).Elem()
	interfaceType := reflect.TypeOf((*interface{})(nil)).Elem()
	myFunc := func(args []reflect.Value) (results []reflect.Value) {
		var result interface{}

		log.Error(ctx, "Decoding workflow task", "taskName", taskName)
		actctx := args[0].Interface().(context.Context)
		activity.GetLogger(actctx).Info("Starting task")

		reqctx, err := ctx.CreateNewRequest("Cadence workflow task "+workflowName, nil, nil, "")
		if err == nil {
			inputTask := args[1].Interface().(*components.Task)

			log.Info(ctx, "Posting workflow task", "input", inputTask)

			err = taskManager.ProcessTask(reqctx, inputTask)
			if err == nil {
				res := reqctx.GetResponse()
				activity.GetLogger(actctx).Info("Task complete", zap.String("taskName", inputTask.Queue))
				if res != nil {
					result = res.Data
				}
			}
		}

		var resVal reflect.Value
		if result != nil {
			resVal = reflect.ValueOf(result)
		} else {
			resVal = reflect.Zero(interfaceType)
		}

		var errVal reflect.Value
		if err != nil {
			errVal = reflect.ValueOf(err).Convert(errType)
		} else {
			errVal = reflect.Zero(errType)
		}

		return []reflect.Value{resVal, errVal}
	}

	var k cadence.TaskProcessor
	act := reflect.MakeFunc(reflect.TypeOf(k), myFunc).Interface().(cadence.TaskProcessor)

	activity.RegisterWithOptions(act, activity.RegisterOptions{Name: taskName})
}

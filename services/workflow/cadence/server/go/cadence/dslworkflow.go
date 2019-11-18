package cadence

import (
	"encoding/json"
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/cadence"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/elements"
	"laatoo/sdk/server/errors"
	"time"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

type DSLWorkflow struct {
	core.Service
	TaskListName string
	WorkflowName string
	DSLWorkflow  DSLWorkflowConfig
	wfConfig     config.Config
	taskManager  elements.TaskManager
	//activities   []*activityFunc
}

func (svc *DSLWorkflow) Initialize(ctx core.ServerContext, conf config.Config) error {

	svc.wfConfig, _ = conf.GetSubConfig(ctx, "Workflow")

	//var x map[string]interface{}

	//x = map[string]interface{}(svc.wfConfig)
	var dslworkflow DSLWorkflowConfig

	byts, err := json.Marshal(svc.wfConfig)
	if err != nil {
		return errors.BadConf(ctx, "Workflow", "err", err)
	}
	err = json.Unmarshal(byts, &dslworkflow)
	if err != nil {
		return errors.BadConf(ctx, "Workflow", "err", err)
	}
	svc.DSLWorkflow = dslworkflow

	svc.taskManager = ctx.GetServerElement(core.ServerElementTaskManager).(elements.TaskManager)

	actsCreated := make(map[string]bool)
	//svc.Tasks, _ = svc.GetStringArrayConfiguration(ctx, "Tasks")
	err = svc.createActivities(ctx, svc.wfConfig, actsCreated)
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	if svc.TaskListName == "" {
		svc.TaskListName = svc.WorkflowName
	}

	cadence.RegisterWorkflow(ctx, svc.WorkflowName, svc.Workflow)
	return nil
}

type (
	// Workflow is the type used to express the workflow definition. Variables are a map of valuables. Variables can be
	// used as input to Activity.
	DSLWorkflowConfig struct {
		Variables map[string]string
		Root      Statement
	}

	// Statement is the building block of dsl workflow. A Statement can be a simple ActivityInvocation or it
	// could be a Sequence or Parallel.
	Statement struct {
		Activity *ActivityInvocation
		Sequence *Sequence
		Parallel *Parallel
	}

	// Sequence consist of a collection of Statements that runs in sequential.
	Sequence struct {
		Elements []*Statement
	}

	// Parallel can be a collection of Statements that runs in parallel.
	Parallel struct {
		Branches []*Statement
	}

	// ActivityInvocation is used to express invoking an Activity. The Arguments defined expected arguments as input to
	// the Activity, the result specify the name of variable that it will store the result as which can then be used as
	// arguments to subsequent ActivityInvocation.
	ActivityInvocation struct {
		Name      string
		Arguments []string
		Result    string
	}

	executable interface {
		execute(ctx workflow.Context, bindings map[string]interface{}) error
	}
)

// SimpleDSLWorkflow workflow decider
func (svc *DSLWorkflow) Workflow(ctx workflow.Context, value interface{}) error {
	bindings := make(map[string]interface{})
	for k, v := range svc.DSLWorkflow.Variables {
		bindings[k] = v
	}

	ao := workflow.ActivityOptions{
		TaskList:               svc.TaskListName,
		ScheduleToCloseTimeout: time.Second * 60,
		ScheduleToStartTimeout: time.Second * 60,
		StartToCloseTimeout:    time.Second * 60,
		HeartbeatTimeout:       time.Second * 20,
		WaitForCancellation:    false,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)

	err := svc.DSLWorkflow.Root.execute(ctx, bindings)
	if err != nil {
		logger.Error("DSL Workflow failed.", zap.Error(err))
		return err
	}

	logger.Info("DSL Workflow completed.")
	return err
}

func (svc *DSLWorkflow) createActivities(ctx core.ServerContext, confToProcess config.Config, actsCreated map[string]bool) error {

	actConfig, ok := confToProcess.GetSubConfig(ctx, "activity")
	if ok {
		actName, ok := actConfig.GetString(ctx, "name")
		if ok {
			created := actsCreated[actName]
			if !created {
				taskActivityCreator(ctx, actName, svc.WorkflowName, svc.taskManager)
				actsCreated[actName] = true
			}
		}
	} else {
		allConfs := confToProcess.AllConfigurations(ctx)
		for _, conf := range allConfs {
			subConf, ok := confToProcess.GetSubConfig(ctx, conf)
			if ok {
				err := svc.createActivities(ctx, subConf, actsCreated)
				if err != nil {
					return errors.WrapError(ctx, err)
				}
				continue
			}
			confArr, ok := confToProcess.GetConfigArray(ctx, conf)
			if ok {
				for _, subConf = range confArr {
					err := svc.createActivities(ctx, subConf, actsCreated)
					if err != nil {
						return errors.WrapError(ctx, err)
					}
				}
			}
		}
	}
	return nil
}

func (b *Statement) execute(ctx workflow.Context, bindings map[string]interface{}) error {
	if b.Parallel != nil {
		err := b.Parallel.execute(ctx, bindings)
		if err != nil {
			return err
		}
	}
	if b.Sequence != nil {
		err := b.Sequence.execute(ctx, bindings)
		if err != nil {
			return err
		}
	}
	if b.Activity != nil {
		err := b.Activity.execute(ctx, bindings)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a ActivityInvocation) execute(ctx workflow.Context, bindings map[string]interface{}) error {
	tsk, err := makeInput(a.Name, a.Arguments, bindings)
	if err != nil {
		return err
	}
	var result string
	err = workflow.ExecuteActivity(ctx, a.Name, tsk).Get(ctx, &result)
	if err != nil {
		return err
	}
	if a.Result != "" {
		bindings[a.Result] = result
	}
	return nil
}

func (s Sequence) execute(ctx workflow.Context, bindings map[string]interface{}) error {
	for _, a := range s.Elements {
		err := a.execute(ctx, bindings)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p Parallel) execute(ctx workflow.Context, bindings map[string]interface{}) error {
	//
	// You can use the context passed in to activity as a way to cancel the activity like standard GO way.
	// Cancelling a parent context will cancel all the derived contexts as well.
	//

	// In the parallel block, we want to execute all of them in parallel and wait for all of them.
	// if one activity fails then we want to cancel all the rest of them as well.
	childCtx, cancelHandler := workflow.WithCancel(ctx)
	selector := workflow.NewSelector(ctx)
	var activityErr error
	for _, s := range p.Branches {
		f := executeAsync(s, childCtx, bindings)
		selector.AddFuture(f, func(f workflow.Future) {
			err := f.Get(ctx, nil)
			if err != nil {
				// cancel all pending activities
				cancelHandler()
				activityErr = err
			}
		})
	}

	for i := 0; i < len(p.Branches); i++ {
		selector.Select(ctx) // this will wait for one branch
		if activityErr != nil {
			return activityErr
		}
	}

	return nil
}

func executeAsync(exe executable, ctx workflow.Context, bindings map[string]interface{}) workflow.Future {
	future, settable := workflow.NewFuture(ctx)
	workflow.Go(ctx, func(ctx workflow.Context) {
		err := exe.execute(ctx, bindings)
		settable.Set(nil, err)
	})
	return future
}

func makeInput(taskName string, argNames []string, argsMap map[string]interface{}) (*components.Task, error) {
	var args []interface{}
	for _, arg := range argNames {
		args = append(args, argsMap[arg])
	}

	var data []byte
	var err error
	data, err = json.Marshal(args)
	if err != nil {
		return nil, err
	}

	tsk := &components.Task{Queue: taskName, Data: data, Token: ""}

	return tsk, nil
}

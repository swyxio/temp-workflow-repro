package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func MyWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow MyWorkflow started")

	ao := workflow.ActivityOptions {
		StartToCloseTimeout: 	time.Minute,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)
	var result string
	err := workflow.ExecuteActivity(ctx, MyActivity, name).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed", err)
	}
	return result, nil
}

func MyActivity(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("MyActivity activity started")
	return "hello " + name, nil
}
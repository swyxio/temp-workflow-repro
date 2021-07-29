package app

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// depends on "go.temporal.io/sdk/activity"
func FoodDeliveryWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow FoodDelivery started")

	ao := workflow.ActivityOptions {
		StartToCloseTimeout: 	time.Minute,
    RetryPolicy: 			&temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2,
			MaximumInterval:    time.Minute * 10,
			MaximumAttempts:    5,
		},
	}

	ctx = workflow.WithActivityOptions(ctx, ao)
	var result string
	err := workflow.ExecuteActivity(ctx, ChargeStripeActivity, name).Get(ctx, &result)
	if err != nil {
		logger.Error("ChargeStripeActivity failed", err)
	}
	return result, nil
}

// depends on "go.temporal.io/sdk/activity"
// depends on "context"
func ChargeStripeActivity(ctx context.Context, apikey string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity apikey started")
	return "apikey: " + apikey, nil
}
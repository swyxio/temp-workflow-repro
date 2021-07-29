package main

import (
	"context"
	"hello-world-project-template-go/app"
	"log"

	"go.temporal.io/sdk/client"
)


func main() {
	// The client is a heavyweight object that should be created once
	// options https://pkg.go.dev/go.temporal.io/sdk@v1.8.0/internal#ClientOptions
	serviceClient, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer serviceClient.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:		"food-delivery-workflow",
		TaskQueue:	"my-task-queue",
	}

	workflowExecution, err := serviceClient.ExecuteWorkflow(
		context.Background(), workflowOptions, app.FoodDeliveryWorkflow, "hello world")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	// synchronously start workflow
	var result string
	err = workflowExecution.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get result", err)
	}
	log.Println("Result:", result)
}
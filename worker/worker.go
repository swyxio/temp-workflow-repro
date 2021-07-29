package main

import (
	"hello-world-project-template-go/app"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)


func main() {
	// The client is a heavyweight object that should be created once
	// options https://pkg.go.dev/go.temporal.io/sdk@v1.8.0/internal#ClientOptions
	serviceClient, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer serviceClient.Close()

	// // // // // // // // // // // // // // 
	// // Poll task queue and set Worker options
	// // options https://pkg.go.dev/go.temporal.io/sdk@v1.8.0/internal#WorkerOptions
	// // // // // // // // // // // // // // 
	mainWorker := worker.New(serviceClient, "my-task-queue", worker.Options{
			// Logger: logger,
	})

	// // // // // // // // // // // // // // 
	// // Register Workflows and activities
	// // // // // // // // // // // // // // 
	mainWorker.RegisterWorkflow(app.FoodDeliveryWorkflow)
	log.Printf("Workflow FoodDeliveryWorkflow registered")
	mainWorker.RegisterActivity(app.ChargeStripeActivity)
	log.Printf("Activity ChargeStripeActivity registered")
	// // // // // // // // // // // // // // 
	// // Register Workflows and activities
	// // // // // // // // // // // // // // 

	err = mainWorker.Run(worker.InterruptCh())
	// or err = mainWorker.Start(), but remember to call Close()

	if err != nil {
			log.Fatalln("Unable to start worker", err)
	}

}

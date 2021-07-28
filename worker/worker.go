package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"hello-world-project-template-go/app"
)


func main() {
	// The client is a heavyweight object that should be created once
	// options https://pkg.go.dev/go.temporal.io/sdk@v1.8.0/internal#ClientOptions
	serviceClient, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer serviceClient.Close()

	// options https://pkg.go.dev/go.temporal.io/sdk@v1.8.0/internal#WorkerOptions
	mainWorker := worker.New(serviceClient, "my-queue", worker.Options{
			// Logger: logger,
	})
	mainWorker.RegisterWorkflow(app.MyWorkflow)
	mainWorker.RegisterActivity(app.MyActivity)

	err = mainWorker.Run(worker.InterruptCh())
	// or err = mainWorker.Start(), but remember to call Close()

	if err != nil {
			log.Fatalln("Unable to start worker", err)
	}

}

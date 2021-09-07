package main

import (
	"context"
	"log"
	"os"
	"time"
)

func main() {

	argoUrl := os.Getenv("ARGO_SERVER_API")
	argoToken := os.Getenv("ARGO_TOKEN")

	if len(argoUrl) == 0 || len(argoToken) == 0 {
		log.Fatal("environment variables ARGO_SERVER_API=<server_url>/api/v1 and ARGO_TOKEN='<secret_token>' must be set")
	}

	if len(os.Args) < 2 {
		log.Fatal(usage())
	}

	context, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	argoClient := NewArgoClient(context, argoUrl, argoToken)

	command := os.Args[1]
	switch command {
	case "events":
		if  len(os.Args) < 3 || len(os.Args[2]) == 0 {
			log.Fatalf(usage())
		}
		watchEvents(argoClient, os.Args[2])

	case "synchistory":
		watchRepo(argoClient)

	default:
		usage()
	}
}

func usage() string {
	return `
usage: argo-watch <command> [qualifiers]
	commands:
		events <app-name>
			Watches for events for a given application and prints to stdout.
			Must provide an application name
		synchistory
			Watches for all git sync history in realtime and prints changes to stdout 
`
}

func watchEvents(client ArgoServerClient, appName string) {
	var eventPrinter = NewEventPrinter()
	argoApp := NewApp(client, appName)

	for {
		events := argoApp.Events()

		for _, event := range events {
			eventPrinter.print(event)
		}

		time.Sleep(time.Millisecond * 60)
	}
}

func watchRepo(client ArgoServerClient) {
	receiver := make(chan *Repo)
	var repo = NewRepo(client)
	repo.client.SetPath("stream/applications")
	go repo.ProcessHistory(receiver, repo.client.Get())

	for {
		select {
		case repoResult := <-receiver:
			repoResult.Print(os.Stdout)
		}
	}
}

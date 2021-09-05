package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"
)

func main() {
	var appName string
	flag.StringVar(&appName, "name", "", "an argo application name to watch")
	flag.Parse()

	if appName == "" {
		log.Fatalf("must provide a valid application name -name=<appname>")
	}

	var ctx, cancel = context.WithTimeout(context.Background(), time.Minute * 5)
	defer cancel()
	watchEvents(ctx, appName)
}

// TODO:
//  - Extract only the time stamp without the date and print it [√]
//  - Add `reason` property to the log [√]
//  - Add unit tests to all of them [√] ------> Check how to create a fake server to test clinet.go
//  - Add a flag modifier to specify the application name to watch [√]
// 	- Add ctx so you would be able to time out after 5 minutes [√]
//  - Add flags to allow using iteration events vs streams
// 	- Back pressure the requests in case of a long wait time and restart clock the moment messages start flowing
//  - Add goroutines

func watchEvents(ctx context.Context, appName string) {
	var eventPrinter = NewEventPrinter()
	for {
		url := os.Getenv("ARGOSERVER_API")
		authToken := "Bearer " + os.Getenv("ARGO_TOKEN")
		argoClient := NewClient(ctx, url, authToken)
		argoApp := NewApp(argoClient, appName)
		events := argoApp.Events()

		for _, event := range events {
			eventPrinter.print(event)
		}

		time.Sleep(time.Millisecond * 60)
	}
}

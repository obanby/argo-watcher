package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type ArgoResponse struct {
	Items []ArgoEvent `json:"items"`
}

type ArgoEvent struct {
	ID        int
	Reason    string `json:"reason"`
	Message   string `json:"message"`
	TimeStamp string `json:"firstTimestamp"`
}

type Application struct {
	name   string
	client ArgoServerClient
}

func NewApp(client ArgoServerClient, name string) *Application {
	return &Application{
		name:   name,
		client: client,
	}
}

func (app *Application) Events() []ArgoEvent {
	app.client.SetPath(fmt.Sprintf("applications/%s/events", app.name))

	var argoResponse = &ArgoResponse{}
	var response = app.client.Get()

	events, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("error reading data from response body: %v", string(events))
	}

	if len(events) == 0 {
		return []ArgoEvent{}
	}

	// TODO: find a better way to handle situation for {"items": null}
	if strings.Contains(string(events), "\"item\":null") {
		return []ArgoEvent{}
	}

	err = json.Unmarshal(events, argoResponse)
	if err != nil {
		log.Fatalf("error unmarshalling data from response body: %v", err)
	}

	for idx, event := range argoResponse.Items {
		event.ID = idx
	}

	return argoResponse.Items
}

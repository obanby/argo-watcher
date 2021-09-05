package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	client ArgoClient
}

func NewApp(client ArgoClient, name string) *Application {
	return &Application{
		name:   name,
		client: client,
	}
}

func (app *Application) Events() []ArgoEvent {
	app.client.SetURL(fmt.Sprintf("%s/applications/%s/events", app.client.URL(), app.name))

	var argoResponse = &ArgoResponse{}
	err := json.Unmarshal(app.client.Get(), argoResponse)
	if err != nil {
		log.Fatalf("error unmarshalling data from response body: %v", err)
	}

	for idx, event := range argoResponse.Items {
		event.ID = idx
	}

	return argoResponse.Items
}

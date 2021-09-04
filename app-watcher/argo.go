package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ArgoServer struct {
	url       string
	authToken string
}

type Application struct {
	name       string
	serverInfo *ArgoServer
}

type ArgoEvent struct {
	ID        int
	Reason    string `json:"reason"`
	Message   string `json:"message"`
	TimeStamp string `json:"firstTimestamp"`
}

type ArgoResponse struct {
	Items []ArgoEvent `json:"items"`
}

func NewArgoServer(url string, authToken string) *ArgoServer {
	return &ArgoServer{url: url, authToken: authToken}
}

func NewApp(serverInfo *ArgoServer, name string) *Application {
	return &Application{name: name, serverInfo: serverInfo}
}

func (app *Application) Events() []ArgoEvent {
	var client = &Client{
		url:     fmt.Sprintf("%s/applications/%s/events", app.serverInfo.url, app.name),
		headers: map[string]string{"authorization": "Bearer " + app.serverInfo.authToken},
	}

	var argoResponse = &ArgoResponse{}
	err := json.Unmarshal(client.Get(), argoResponse)
	if err != nil {
		log.Fatalf("error unmarshalling data from response body.")
	}

	for idx, event := range argoResponse.Items {
		event.ID = idx
	}

	return argoResponse.Items
}

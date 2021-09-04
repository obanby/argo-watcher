package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ArgoEvent struct {
	ID        int
	Reason    string `json:"reason"`
	Message   string `json:"message"`
	TimeStamp string `json:"firstTimestamp"`
}

type Response struct {
	Items []ArgoEvent `json:"items"`
}

func main() {
	itterate()
}

func fatal(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

// TODO:
//  - Extract only the time stamp without the date and print it
//  - Add `reason` property to the log
//  - Add unit tests to all of them
//  - Add flags to allow using iteration events vs streams
//  - Add a flag modifier to specify the application name to watch
// 	- Back pressure the requests in case of a long wait time and restart clock the moment messages start flowing
// 	- Add context so you would be able to time out after 5 minutes
//  - Add goroutines


func itterate() {
	var eventPrinter = NewEventPrinter()
	for {
		const ARGOCD_SERVER = "http://127.0.0.1:8080/api/v1"
		url := ARGOCD_SERVER + "/applications/sock-shop-example/events"
		barer := "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzA4NTYwODYsImp0aSI6IjExOGVkYmUwLTAwZDctNDhjZC1iOTQ2LTE0ZjJlZmFiYjk1NCIsImlhdCI6MTYzMDc2OTY4NiwiaXNzIjoiYXJnb2NkIiwibmJmIjoxNjMwNzY5Njg2LCJzdWIiOiJhZG1pbjpsb2dpbiJ9.YmZjZEK2DHGQvKtRiePawveD5TIu8sykNlvY47xykU8"

		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("authorization", barer)

		client := &http.Client{}
		res, err := client.Do(req)
		fatal(err, "error occured while connecting to ARGOCD_SERVER %v")

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("error reading data from response body. %v", data)
		}

		var response = &Response{}
		err = json.Unmarshal(data, response)
		fatal(err, "failed while unmarshalling json")

		for idx, event := range response.Items {
			event.ID = idx
			eventPrinter.printArgoEvent(event)
		}

		time.Sleep(time.Millisecond * 60)
	}
}


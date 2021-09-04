package main

import "fmt"

type EventPrinter struct {
	events []ArgoEvent
}

func NewEventPrinter() *EventPrinter {
	var eventPrinter = EventPrinter{}
	eventPrinter.events = make([]ArgoEvent, 0, 100)
	return &eventPrinter
}

func (ep *EventPrinter) add(response ArgoEvent) {
	ep.events = append(ep.events, response)
}

func (ep *EventPrinter) printArgoEvent(event ArgoEvent) {
	for _, e := range ep.events {
		if e.Message == event.Message && e.ID == event.ID {
			return
		}
	}
	ep.add(event)
	fmt.Printf("[%s] %s\n", event.TimeStamp, event.Message)
}

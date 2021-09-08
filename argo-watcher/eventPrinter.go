package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type EventPrinter struct {
	writer io.Writer
	events  []ArgoEvent
}

func NewEventPrinter() *EventPrinter {
	var eventPrinter = EventPrinter{}
	eventPrinter.events = make([]ArgoEvent, 0, 100)
	eventPrinter.writer = os.Stdout
	return &eventPrinter
}

func (ep *EventPrinter) add(event ArgoEvent) {
	ep.events = append(ep.events, event)
}

func (ep *EventPrinter) print(event ArgoEvent) {
	if ep.isDuplicate(event) {
		return
	}
	ep.add(event)
	fmt.Fprintf(ep.writer, "%s %s: %s\n", extractTime(event.TimeStamp), event.Reason, event.Message)
}

func (ep *EventPrinter) isDuplicate(event ArgoEvent) bool {
	for _, e := range ep.events {
		if e.Message == event.Message && e.ID == event.ID {
			return true
		}
	}
	return false
}

func extractTime(timestamp string) string {
	extractedTime := timestamp
	if strings.Contains(timestamp, "Z") {
		extractedTime = timestamp[:len(timestamp)-1]
	}

	if strings.Contains(extractedTime, "T") {
		tIdx := strings.Index(extractedTime, "T")
		extractedTime = extractedTime[tIdx+1:]
	}

	return extractedTime
}

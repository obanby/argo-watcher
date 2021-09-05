package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestEventPrinter_add(t *testing.T) {
	got := NewEventPrinter()
	want := &EventPrinter{
		writer: os.Stdout,
		events: argoEventsTable,
	}
	for _, event := range argoEventsTable {
		got.add(event)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("failed wanted struct %v \ngot %v", want, got)
	}
}

func TestEventPrinter_isDuplicate(t *testing.T) {
	eventPrinter := NewEventPrinter()
	want := true
	for _, event := range argoEventsTable {
		eventPrinter.add(event)
	}

	for _, event := range argoEventsTable {
		got := eventPrinter.isDuplicate(event)
		if want != got {
			t.Fatalf("failed in eventPrinter.isDuplicate(event) = %v, expected %v", got, want)
		}
	}
}

func TestExtractTime(t *testing.T) {
	arg := "2021-09-05T01:31:18Z"
	want := "01:31:18"
	got := extractTime(arg)

	if want != got {
		t.Fatalf("failed in extractTime(%v) = %v expected %v", arg, got, want)
	}

	got = extractTime(want)
	if want != got {
		t.Fatalf("extractTime(%s) = %s Expected %s", want, got, want)
	}
}

func TestEventPritner_print(t *testing.T) {
	var buffer bytes.Buffer
	eventPrinter := NewEventPrinter()
	eventPrinter.writer = &buffer

	for _, event := range argoEventsTable {
		want := fmt.Sprintf("%s %s: %s\n", event.TimeStamp, event.Reason, event.Message)
		eventPrinter.print(event)
		if want != buffer.String() {
			t.Fatalf("error in eventPrinter.print(%v) = %s Expected %s", event, buffer.String(), want)
		}
		buffer.Reset()
	}
}

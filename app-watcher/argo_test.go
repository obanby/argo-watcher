package main

import (
	"log"
	"reflect"
	"testing"
)

func TestApplication_Events(t *testing.T) {
	mockClient := NewMockClient("localhost:8080", "faketoken")
	argoApp := NewApp(mockClient, "example-app")
	got := argoApp.Events()

	if !reflect.DeepEqual(argoEventsTable, got) {
		log.Fatalf("failed in argoApp.Events() = %v Expected %v", got, argoEventsTable)
	}
}

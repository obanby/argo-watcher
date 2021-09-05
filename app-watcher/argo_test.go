package main

import (
	"context"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestApplication_Events(t *testing.T) {
	var ctx = context.Background()
	mockClient := NewMockClient(ctx, "localhost:8080", "faketoken")
	argoApp := NewApp(mockClient, "example-app")
	got := argoApp.Events()

	if !reflect.DeepEqual(argoEventsTable, got) {
		log.Fatalf("failed in argoApp.Events() = %v Expected %v", got, argoEventsTable)
	}
}

func TestApplication_Events_timeout(t *testing.T) {
	 ctx, close := context.WithTimeout(context.Background(), time.Millisecond * 2)
	 defer close()
	 mockClient := NewMockClient(ctx, "localhost:8080", "faketoken")
	 argoApp := NewApp(mockClient, "example-app")
	 got := argoApp.Events()

	 if len(got) != 0 {
	 	log.Fatalf("error with context timeout: expected timeout to cause empty events list, got %v", got)
	 }
}

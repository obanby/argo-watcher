package main

import (
	"bytes"
	"context"
	"log"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestNewRepo(t *testing.T) {
	mockClient := NewMockClient(context.TODO(), "localhost:8080/rest/api/v1", "faketoken")
	want := Repo{client: mockClient}
	got := NewRepo(mockClient)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("failed in NewRepo(client) = %v wanted = %v", got, want)
	}
}

func TestProcessHistory(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	mockClient := NewMockClient(ctx, "localhost:8080/rest/api/v1", "faketoken")
	mockClient.SetPath("/stream/application")
	repo := NewRepo(mockClient)
	receiver := make(chan *Repo)
	go repo.ProcessHistory(receiver, repo.client.Get())

	var buffer bytes.Buffer

channelHandler:
	for {
		select {
		case repoResult, isOpen := <-receiver:
			if !isOpen {
				receiver = nil
				break channelHandler
			}
			repoResult.Print(&buffer)
		}
	}

	want := []string{
		"commit: [e51eba] app: guestbook has changed for reason -> keep more revisions history",
		"commit: [e40996] app: guestbook has changed for reason -> keep more revisions history",
	}

	repoResults := buffer.String()
	got := strings.Split(repoResults, "\n")

	for idx, gotMsg := range got {
		if gotMsg == "" {
			continue
		}
		if len(want) > idx {
			wantMsg := want[idx]
			if wantMsg != gotMsg {
				log.Fatalf("error while Repo.ProccessHistory() = %s  want %s", got, want)
			}
		}
	}
}

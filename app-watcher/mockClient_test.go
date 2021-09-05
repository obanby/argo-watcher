package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type MockClient struct {
	ctx     context.Context
	url     string
	headers map[string]string
}

func NewMockClient(ctx context.Context, url string, authToken string) ArgoClient {
	return &MockClient{
		ctx:     ctx,
		url:     url,
		headers: map[string]string{"authorization": "Bearer " + authToken},
	}
}

func (c *MockClient) URL() string {
	return c.url
}

func (c *MockClient) Context() context.Context {
	return c.ctx
}

func (c *MockClient) SetURL(url string) {
	c.url = url
}

func (c *MockClient) Get() []byte {
	time.Sleep(time.Millisecond * 3)

	response := ArgoResponse{Items: argoEventsTable}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(response)

	select {
	case <-c.ctx.Done():
		fmt.Println("context is done")
		return []byte{}
	default:
		return buffer.Bytes()
	}
}

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type MockClient struct {
	ctx     context.Context
	url     string
	headers map[string]string
}

func NewMockClient(ctx context.Context, url string, authToken string) ArgoServerClient {
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

func (c *MockClient) SetPath(url string) {
	c.url = url
}

func (c *MockClient) Get() *http.Response {
	time.Sleep(time.Millisecond * 3)

	response := ArgoResponse{Items: argoEventsTable}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(response)

	select {
	case <-c.ctx.Done():
		fmt.Println("done is called")
		return &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(""))}
	default:
		return &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(buffer.String())),
		}
	}
}

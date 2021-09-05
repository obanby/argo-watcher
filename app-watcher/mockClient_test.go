package main

import (
	"bytes"
	"encoding/json"
)

type MockClient struct {
	url     string
	headers map[string]string
}

func NewMockClient(url string, authToken string) ArgoClient {
	return &MockClient{
		url:     url,
		headers: map[string]string{"authorization": "Bearer " + authToken},
	}
}

func (c *MockClient) URL() string {
	return c.url
}

func (c *MockClient) SetURL(url string) {
	c.url = url
}

func (c *MockClient) Get() []byte {
	response := ArgoResponse{Items: argoEventsTable}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(response)
	return buffer.Bytes()
}

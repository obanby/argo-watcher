package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type ArgoServerClient interface {
	SetPath(path string)
	URL() string
	Get() *http.Response
	Context() context.Context
}

type ArgoClient struct {
	ctx     context.Context
	url     string
	path    string
	headers map[string]string
}

func NewArgoClient(ctx context.Context, url string, authToken string) ArgoServerClient {
	return &ArgoClient{
		ctx:     ctx,
		url:     url,
		headers: map[string]string{"authorization": "Bearer " + authToken},
	}
}

func (c *ArgoClient) URL() string {
	return c.url
}

func (c *ArgoClient) SetPath(path string) {
	c.path = path
}

func (c *ArgoClient) Context() context.Context {
	return c.ctx
}

func (c *ArgoClient) Get() *http.Response {
	uri := fmt.Sprintf("%s/%s", c.url, c.path)
	req, err := http.NewRequestWithContext(c.ctx, "GET", uri, nil)
	if err != nil {
		log.Fatalf("error creating request: %v", err)
	}

	for header, value := range c.headers {
		req.Header.Add(header, value)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("error occured while connecting to ARGOCD_SERVER %v", err)
	}

	return response
}

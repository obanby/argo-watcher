package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type ArgoClient interface {
	SetURL(url string)
	URL() string
	Get() []byte
}

type Client struct {
	url     string
	headers map[string]string
}

func NewClient(url string, authToken string) ArgoClient {
	return &Client{
		url:     url,
		headers: map[string]string{"authorization": "Bearer " + authToken},
	}
}

func (c *Client) URL() string {
	return c.url
}

func (c *Client) SetURL(url string) {
	c.url = url
}

func (c *Client) Get() []byte {
	req, err := http.NewRequest("GET", c.url, nil)
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

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("error reading data from response body: %v", string(data))
	}

	return data
}

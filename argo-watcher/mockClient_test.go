package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type MockClient struct {
	ctx     context.Context
	url     string
	path    string
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

func (c *MockClient) SetPath(path string) {
	c.path = path
}

func (c *MockClient) Get() *http.Response {
	time.Sleep(time.Millisecond * 3)
	select {
	case <-c.ctx.Done():
		return &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(""))}
	default:
		return pathHandler(c.path)
	}
}

func pathHandler(path string) *http.Response {
	response := ArgoResponse{Items: argoEventsTable}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(response)

	switch {
	case strings.Contains(path, "events"):
		return &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(buffer.String())),
		}

	case strings.Contains(path, "error/stream/application"):
		return &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(sampleErrorStreamPayLoad())),
		}

	case strings.Contains(path, "stream/application"):
		return &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(sampleStreamPayLoad())),
		}

	case strings.Contains(path, "error/revisions/"):
		return &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(sampleErrorStreamPayLoad())),
		}

	case strings.Contains(path, "revisions/"):
		return &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(sampleRevisionPayLoad())),
		}
	}

	return &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(""))}
}

func sampleStreamPayLoad() string {
	return `
{
  "result": {
    "type": "MODIFIED",
    "application": {
      "status": {
        "history": [
          {
            "revision": "e51eba979260fbf733db58bf01e5bc0becb0511b",
            "deployedAt": "2021-09-05T18:12:12Z",
            "id": 0,
            "source": {
              "repoURL": "https://github.com/obanby/argocd-example-apps.git",
              "path": "guestbook",
              "targetRevision": "HEAD"
            },
            "deployStartedAt": "2021-09-05T18:12:12Z"
          },
          {
            "revision": "e409963badf15dc7a15c2575bb543166fa943efe",
            "deployedAt": "2021-09-05T18:28:02Z",
            "id": 1,
            "source": {
              "repoURL": "https://github.com/obanby/argocd-example-apps.git",
              "path": "guestbook",
              "targetRevision": "HEAD"
            },
            "deployStartedAt": "2021-09-05T18:28:01Z"
          }
        ]
      }
    }
  }
}
`
}

func sampleErrorStreamPayLoad() string {
	return `
`
}

func sampleRevisionPayLoad() string {
	return `
{
    "author": "Omar Elbanby",
    "date": "2021-09-05T16:45:52Z",
    "message": "keep more revisions history"
}
`
}
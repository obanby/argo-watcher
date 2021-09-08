package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Repo struct {
	client ArgoServerClient
	Result struct {
		App struct {
			Status struct {
				History []struct {
					commitMsg string
					Source    struct {
						RepoURL string `json:"repoURL"`
						Path    string `json:"path"`
					} `json:"source"`
					Revision string `json:"revision"`
				} `json:"history"`
			} `json:"status"`
		} `json:"application"`
	} `json:"result"`
}

func NewRepo(client ArgoServerClient) Repo {
	return Repo{
		client: client,
	}
}

func (repo *Repo) ProcessHistory(receiver chan<- *Repo, response *http.Response) {
	for {
		err := json.NewDecoder(response.Body).Decode(repo)
		if err != nil {
			if err == io.EOF {
				close(receiver)
				return
			}
			// we can safely ignore the data here
			// since sometimes we get a big payload that pollutes the buffer and causes the decoding to fail
			// the easiest way is to continue and the server will send each request individually on other iterations
			continue
		}
		repo.addCommitMsg()

		select {
		case receiver <- repo:
		case <-repo.client.Context().Done():
			return
		}
	}

}

func (repo *Repo) addCommitMsg() {
	syncHistory := repo.Result.App.Status.History
	for idx, item := range syncHistory {
		syncHistory[idx].commitMsg = repo.commitMsg(item.Source.Path, item.Revision)
	}
}

func (repo *Repo) commitMsg(appName string, revision string) string {
	repo.client.SetPath(fmt.Sprintf("/applications/%s/revisions/%s/metadata", appName, revision))
	result := repo.client.Get()
	data, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatalf("error while reading response body for commit message")
	}

	commitMsg := struct {
		Message string `json:"message"`
	}{}

	err = json.Unmarshal(data, &commitMsg)
	if err != nil {
		log.Fatalf("error marshaling json for commitMsg() error is: %s ", err)
	}

	return commitMsg.Message
}

func (repo *Repo) Print(writer io.Writer) {
	history := repo.Result.App.Status.History
	for _, h := range history {
		fmt.Fprintf(writer, "commit: [%s] app: %s has changed for reason -> %s\n", trimCommitSha(h.Revision), h.Source.Path, h.commitMsg)
	}
}

func trimCommitSha(sha string) string {
	if len(sha) < 6 {
		return sha
	}
	return sha[:6]
}

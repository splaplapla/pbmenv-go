package commands

import (
	"main/internal"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func AvailableVersions(client HTTPClient) ([]string, error) {
	versions, err := internal.PBMGithubClient(client).AvailableVersions()
	if err != nil {
		return []string{}, err
	}

	return versions, nil
}

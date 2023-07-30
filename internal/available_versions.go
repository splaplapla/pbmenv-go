package internal

import (
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func AvailableVersions(client HTTPClient) ([]string, error) {
	versions, err := PBMGithubClient(client).availableVersions()
	if err != nil {
		return []string{}, err
	}

	return versions, nil
}

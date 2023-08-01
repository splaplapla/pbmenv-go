package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type GithubClient struct {
	Http      HTTPClient
	RepoOwner string
	RepoName  string
}

type GithubTag struct {
	Name string `json:"name"`
}

func (c *GithubClient) tags() ([]GithubTag, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/tags", c.RepoOwner, c.RepoName)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var tags []GithubTag
	if err := json.Unmarshal(body, &tags); err != nil {
		return nil, err
	}
	return tags, nil
}

func (c *GithubClient) AvailableVersions() ([]string, error) {
	tags, err := c.tags()
	if err != nil {
		return nil, err
	}
	return c.extractVersions(tags), nil
}

func (c *GithubClient) extractVersions(tags []GithubTag) []string {
	re := regexp.MustCompile(`v([\d.]+)$`)
	versions := []string{}

	for _, tag := range tags {
		if matches := re.FindStringSubmatch(tag.Name); len(matches) == 2 {
			versions = append(versions, matches[1])
		}
	}
	return versions
}

func PBMGithubClient(client HTTPClient) *GithubClient {
	return &GithubClient{
		Http:      client,
		RepoOwner: "splaplapla",
		RepoName:  "procon_bypass_man",
	}
}

package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func AvailableVersions(client HttpClient) []string {
	url := "https://api.github.com/repos/splaplapla/procon_bypass_man/tags"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response", err)
		return nil
	}

	var jsonResp []map[string]interface{}
	if err := json.Unmarshal(body, &jsonResp); err != nil {
		fmt.Println("Error parsing the response", err)
		return nil
	}
	versions := extractVersions(jsonResp)
	return versions
}

func extractVersions(jsonResp []map[string]interface{}) []string {
	re := regexp.MustCompile(`v([\d.]+)$`)
	versions := []string{}

	for _, tag := range jsonResp {
		if name, ok := tag["name"].(string); ok {
			if matches := re.FindStringSubmatch(name); len(matches) == 2 {
				versions = append(versions, matches[1])
			}
		}
	}

	return versions
}

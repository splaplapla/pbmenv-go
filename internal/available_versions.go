package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func AvailableVersions() []string {
	url := "https://api.github.com/repos/splaplapla/procon_bypass_man/tags"
	resp, err := http.Get(url)
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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		help()
		return
	}

	subCommand := args[0]
	switch subCommand {
	case "available_versions", "av", "a":
		availableVersions := AvailableVersions()
		for _, version := range availableVersions {
			fmt.Println(version)
		}
	case "versions", "v":
		// TODO
	case "install", "i":
		// TODO
	case "use", "u":
		// TODO
	case "uninstall":
		// TODO
	case "clean", "c":
		// TODO
	case "--version", "-v":
		// TODO
	case "--help", "-h", "help", "h":
		help()
	default:
		fmt.Println("pbmenv: '" + subCommand + "' is not a pbmenv command. See 'pbmenv --help'.")
	}
}

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

func help() {
	fmt.Println(`Usage: pbmenv [command]

	Available commands:
	available_versions    Display the available versions of pbmenv
	versions              List the installed versions of pbmenv
	install               Install a specific version of pbmenv
	use                   Set a specific version of pbmenv as the active version
	uninstall             Uninstall a specific version of pbmenv
	clean                 Remove old installed versions of pbmenv`)
	return
}

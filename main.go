package main

import (
	"fmt"
	"log"
	"main/internal"
	"net/http"
	"os"
)

const VERSION = "0.1.0"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		help()
		return
	}

	subCommand := args[0]
	switch subCommand {
	case "available_versions", "av", "a":
		availableVersions, err := internal.AvailableVersions(&http.Client{})
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
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
		fmt.Println(VERSION)
	case "--help", "-h", "help", "h":
		help()
	default:
		fmt.Println("pbmenv: '" + subCommand + "' is not a pbmenv command. See 'pbmenv --help'.")
	}
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

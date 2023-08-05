package main

import (
	"fmt"
	"log"
	"main/internal/commands"
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
		availableVersions, err := commands.AvailableVersions(&http.Client{})
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		for _, version := range availableVersions {
			fmt.Println(version)
		}
	case "versions", "v":
		// TODO
	case "install", "i":
		if len(args) < 2 {
			fmt.Println("pbmenv: 'install' requires a version argument. See 'pbmenv --help'.")
			os.Exit(1)
		}

		subCommandArg := args[1]
		useOption := false
		installBaseDir := "/usr/share/pbm"
		if len(args) > 2 {
			switch args[2] {
			case "--use", "-u":
				useOption = true
			case "--debug_change_install_dir":
				installBaseDir = "./tmp/pbm"
			default:
				fmt.Println("pbmenv: '" + args[2] + "' is not a pbmenv option. See 'pbmenv --help'.")
				os.Exit(1)
			}
		}
		err := commands.InstallVersion(&http.Client{}, subCommandArg, useOption, installBaseDir)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
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
    available_versions    Display the available versions of pbmenv.
    versions              List the installed versions of pbmenv.
    install               Install a specific version of pbmenv. (optionally, use the --use or -u flag to use the installed version)
    use                   Set a specific version of pbmenv as the active version.
    uninstall             Uninstall a specific version of pbmenv.
    clean                 Remove old installed versions of pbmenv.`)
	return
}

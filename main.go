package main

import (
	"flag"
	"fmt"
	"log"
	"main/internal/commands"
	"net/http"
	"os"
)

const VERSION = "0.1.0"

func main() {
	flag.Parse()
	args := flag.Args()[0:]
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
		// NOTE: ex: $ cmd install 0.1.1 --use --debug_change_install_dir
		// NOTE: ex: $ cmd install 0.1.1
		if len(args) < 2 {
			fmt.Println("pbmenv: 'install' requires a version argument. See 'pbmenv --help'.")
			os.Exit(1)
		}

		installCmd := flag.NewFlagSet("install", flag.ExitOnError)
		useOption := installCmd.Bool("use", false, "use installed version")
		debugChangeInstallDir := installCmd.Bool("debug_change_install_dir", false, "Debug change install dir")
		installCmd.Parse(args[2:])
		installBaseDir := "/usr/share/pbm"
		if *debugChangeInstallDir {
			installBaseDir = "./tmp/pbm"
		}
		versionToInstall := args[1]

		err := commands.InstallVersion(&http.Client{}, versionToInstall, *useOption, installBaseDir)
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

package commands

import "fmt"

func downloadVersion(client HTTPClient, version string) error {
	return nil
}

func prepareDirectory(installBaseDir string) error {
	return nil
}

func validateDiretory(installBaseDir string) error {
	return nil
}

func validateVersion(version string) error {
	return nil
}

func InstallVersion(client HTTPClient, version string, useOption bool, installBaseDir string) error {
	fmt.Println(version)
	fmt.Println(useOption)
	fmt.Println(installBaseDir)
	error := validateVersion(version)
	if error != nil {
		return error
	}
	downloadVersion(client, version)

	return nil
}

package commands

import (
	"fmt"
	"main/internal"
)

func downloadVersion(client HTTPClient, version string) error {
	return nil
}

func prepareDirectory(installBaseDir string) error {
	return nil
}

func validateDiretory(installBaseDir string) error {
	return nil
}

func InstallVersion(client HTTPClient, version string, useOption bool, installBaseDir string) error {
	fmt.Println(version)
	fmt.Println(useOption)
	fmt.Println(installBaseDir)

	exists, err := internal.PBMGithubClient(client).ExistsVersion(version)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("Version %s does not exist", version)
	}

	downloadVersion(client, version)

	return nil
}

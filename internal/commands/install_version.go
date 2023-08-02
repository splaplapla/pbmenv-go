package commands

import "fmt"

func downloadVersion(client HTTPClient, version string) error {
	return nil
}

func InstallVersion(client HTTPClient, version string, useOption bool) ([]string, error) {
	fmt.Println(version)
	fmt.Println(useOption)

	return []string{}, nil
}

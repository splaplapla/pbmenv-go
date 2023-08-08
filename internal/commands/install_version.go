package commands

import (
	"fmt"
	"main/internal"
	"os"
	"os/exec"
)

func downloadVersion(client HTTPClient, version string) error {
	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", fmt.Sprintf("curl -L https://github.com/splaplapla/procon_bypass_man/archive/refs/tags/v%s.tar.gz | tar xvz -C /tmp > /dev/null 2>&1", version))
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func prepareDirectory(installBaseDir string) error {
	return nil
}

func existsVersionDirectory(installBaseDir string, version string) (bool, error) {
	pbmVersionPath := installBaseDir + "/v" + version
	if _, err := os.Stat(pbmVersionPath); err == nil {
		// fmt.Printf("ファイル %s は存在します。\n", pbmVersionPath)
		return true, nil
	} else if os.IsNotExist(err) {
		// fmt.Printf("ファイル %s は存在しません。\n", pbmVersionPath)
		return false, nil
	} else {
		return false, err
	}
}

func InstallVersion(client HTTPClient, version string, useOption bool, installBaseDir string) error {
	exists, err := internal.PBMGithubClient(client).ExistsVersion(version)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Version %s does not exist", version)
	}

	exists, err = existsVersionDirectory(installBaseDir, version)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("Version %s already installed", version)
	}

	downloadVersion(client, version)

	return nil
}

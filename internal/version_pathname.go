package internal

type VersionPathname struct {
	InstallBaseDir string
	Version        string
}

func (v *VersionPathname) VersionPath() string {
	return v.InstallBaseDir + "/v" + v.Version
}

func (v *VersionPathname) VersionPathWithoutV() string {
	return v.InstallBaseDir + "/" + v.Version
}

func (v *VersionPathname) AppRbPath() string {
	return v.VersionPath() + "/app.rb"
}

func (v *VersionPathname) AppRbErbPath() string {
	return v.VersionPath() + "/app.rb.erb"
}

func (v *VersionPathname) DeviceIdPathInVersion() string {
	return v.VersionPath() + "/device_id"
}

func (v *VersionPathname) SrcPbmPath() string {
	return "/tmp/procon_bypass_man-v" + v.Version
}

func (v *VersionPathname) ProjectTemplateFilePaths(includeAppErb bool) []string {
	paths := []string{"README.md", "setting.yml"}
	if includeAppErb {
		paths = append(paths, "app.rb.erb")
	} else {
		paths = append(paths, "app.rb")
	}
	return paths
}

func (v *VersionPathname) Shared() string {
	return v.InstallBaseDir + "/shared"
}

func (v *VersionPathname) DeviceIdPathInShared() string {
	return v.Shared() + "/device_id"
}

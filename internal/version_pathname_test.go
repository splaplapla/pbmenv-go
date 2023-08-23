package internal

import (
	"testing"
)

func TestVersionPath(t *testing.T) {
	v := &VersionPathname{
		InstallBaseDir: "/path/to/install",
		Version:        "1.0",
	}
	expected := "/path/to/install/v1.0"
	if v.VersionPath() != expected {
		t.Errorf("Expected %s, but got %s", expected, v.VersionPath())
	}
}

func TestAppRbPath(t *testing.T) {
	v := &VersionPathname{
		InstallBaseDir: "/path/to/install",
		Version:        "1.0",
	}
	expected := "/path/to/install/v1.0/app.rb"
	if v.AppRbPath() != expected {
		t.Errorf("Expected %s, but got %s", expected, v.AppRbPath())
	}
}

func TestVersionPathWithoutV(t *testing.T) {
	v := &VersionPathname{
		InstallBaseDir: "/path/to/install",
		Version:        "1.0",
	}
	expected := "/path/to/install/1.0"
	if v.VersionPathWithoutV() != expected {
		t.Errorf("Expected %s, but got %s", expected, v.VersionPathWithoutV())
	}
}

func TestAppErbPath(t *testing.T) {
	v := &VersionPathname{
		InstallBaseDir: "/path/to/install",
		Version:        "1.0",
	}
	expected := "/path/to/install/v1.0/app.rb.erb"
	if v.AppRbErbPath() != expected {
		t.Errorf("Expected %s, but got %s", expected, v.AppRbErbPath())
	}
}

func TestDeviceIdPathInVersion(t *testing.T) {
	v := &VersionPathname{
		InstallBaseDir: "/path/to/install",
		Version:        "1.0",
	}
	expected := "/path/to/install/v1.0/device_id"
	if v.DeviceIdPathInVersion() != expected {
		t.Errorf("Expected %s, but got %s", expected, v.DeviceIdPathInVersion())
	}
}

func TestSrcPbmPath(t *testing.T) {
	v := &VersionPathname{
		InstallBaseDir: "/path/to/install",
		Version:        "1.0",
	}
	expected := "/tmp/procon_bypass_man-v1.0"
	if v.SrcPbmPath() != expected {
		t.Errorf("Expected %s, but got %s", expected, v.SrcPbmPath())
	}
}

func TestSharedPath(t *testing.T) {
	v := &VersionPathname{
		InstallBaseDir: "/path/to/install",
		Version:        "1.0",
	}
	expected := "/path/to/install/shared"
	if v.Shared() != expected {
		t.Errorf("Expected %s, but got %s", expected, v.Shared())
	}
}

func TestDevicePathInShared(t *testing.T) {
	v := &VersionPathname{
		InstallBaseDir: "/path/to/install",
		Version:        "1.0",
	}
	expected := "/path/to/install/shared/device_id"
	if v.DeviceIdPathInShared() != expected {
		t.Errorf("Expected %s, but got %s", expected, v.DeviceIdPathInShared())
	}
}

//go:build windows
// +build windows

package shortcut

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

func TestCreateShortcut(t *testing.T) {
	var err error
	err = CreateShortcut("google", "https://google.com", "", ".")
	if err == nil {
		_, err = os.Stat("./google.lnk")
		if err == nil {
			t.Log("CreateShortcut pass.")
			return
		}
	}
	t.Errorf("CreateShortcut error: %s", err)
}

func TestCreateDesktopShortcut(t *testing.T) {
	var err error
	err = CreateDesktopShortcut("google", "https://google.com", "")
	if err == nil {
		u, _ := user.Current()
		dst := filepath.Join(u.HomeDir, "Desktop", "google.lnk")
		_, err = os.Stat(dst)
		if err == nil {
			t.Log("CreateDesktopShortcut pass.")
			return
		}
	}
	t.Errorf("CreateDesktopShortcut error: %s", err)
}

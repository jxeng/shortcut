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
	sc := Shortcut{
		ShortcutPath:     "./google.lnk",
		Target:           "https://google.com",
		IconLocation:     "%SystemRoot%\\System32\\SHELL32.dll,0",
		Arguments:        "",
		Description:      "",
		Hotkey:           "",
		WindowStyle:      "1",
		WorkingDirectory: "",
	}
	err = Create(sc)
	if err == nil {
		_, err = os.Stat("./google.lnk")
		if err == nil {
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
			return
		}
	}
	t.Errorf("CreateDesktopShortcut error: %s", err)
}

//go:build windows
// +build windows

package shortcut

import (
	"os/user"
	"path/filepath"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// Shortcut the shortcut (.lnk file) property struct
type Shortcut struct {
	// Shortcut (.lnk file) path
	ShortcutPath string
	// Shortcut target: a file path or a website
	Target string
	// Shortcut icon path, default: "%SystemRoot%\\System32\\SHELL32.dll,0"
	IconLocation string
	// Arguments of shortcut
	Arguments string
	// Description of shortcut
	Description string
	// Hotkey of shortcut
	Hotkey string
	// WindowStyle, "1"(default) for default size and location; "3" for maximized window; "7" for minimized window
	WindowStyle string
	// Working directory of shortcut
	WorkingDirectory string
}

// CreateShortcut create a desktop shortcut with name, target and iconPath
// target is a file or website
// if iconPath is empty string, icon would be "%SystemRoot%\\System32\\SHELL32.dll,0"
func CreateDesktopShortcut(name, target, iconPath string) error {

	u, err := user.Current()
	if err != nil {
		return err
	}
	shortcutPath := filepath.Join(u.HomeDir, "Desktop", name+".lnk")
	shortcut := Shortcut{
		ShortcutPath:     shortcutPath,
		Target:           target,
		IconLocation:     iconPath,
		Arguments:        "",
		Description:      "",
		Hotkey:           "",
		WindowStyle:      "1",
		WorkingDirectory: "",
	}
	return Create(shortcut)
}

// CreateShortcut create with a shortcut object
func Create(shortcut Shortcut) error {
	if shortcut.IconLocation == "" {
		shortcut.IconLocation = "%SystemRoot%\\System32\\SHELL32.dll,0"
	}
	if shortcut.WindowStyle == "" {
		shortcut.WindowStyle = "1"
	}
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		return err
	}
	defer oleShellObject.Release()
	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}
	defer wshell.Release()
	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", shortcut.ShortcutPath)
	if err != nil {
		return err
	}

	idispatch := cs.ToIDispatch()
	_, err = oleutil.PutProperty(idispatch, "IconLocation", shortcut.IconLocation)
	if err != nil {
		return err
	}
	_, err = oleutil.PutProperty(idispatch, "TargetPath", shortcut.Target)
	if err != nil {
		return err
	}
	_, err = oleutil.PutProperty(idispatch, "Arguments", shortcut.Arguments)
	if err != nil {
		return err
	}
	_, err = oleutil.PutProperty(idispatch, "Description", shortcut.Description)
	if err != nil {
		return err
	}
	_, err = oleutil.PutProperty(idispatch, "Hotkey", shortcut.Hotkey)
	if err != nil {
		return err
	}
	_, err = oleutil.PutProperty(idispatch, "WindowStyle", shortcut.WindowStyle)
	if err != nil {
		return err
	}
	_, err = oleutil.PutProperty(idispatch, "WorkingDirectory", shortcut.WorkingDirectory)
	if err != nil {
		return err
	}
	_, err = oleutil.CallMethod(idispatch, "Save")
	if err != nil {
		return err
	}
	return nil
}

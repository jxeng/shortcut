//go:build windows
// +build windows

package shortcut

import (
	"os/user"
	"path/filepath"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// CreateShortcut create a shortcut with name and icon
// target to src file or website in the dir directory
func CreateShortcut(name, src, icon, dir string) error {
	dst := filepath.Join(dir, name+".lnk")
	return createShortcut(src, dst, icon)
}

// CreateShortcut create a desktop shortcut with name and icon
// target to src file or website
func CreateDesktopShortcut(name, src, icon string) error {

	u, err := user.Current()
	if err != nil {
		return err
	}
	dst := filepath.Join(u.HomeDir, "Desktop", name+".lnk")
	return createShortcut(src, dst, icon)
}

func createShortcut(src, dst, icon string) error {
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
	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", dst)
	if err != nil {
		return err
	}
	idispatch := cs.ToIDispatch()
	_, err = oleutil.PutProperty(idispatch, "TargetPath", src)
	if err != nil {
		return err
	}
	if icon != "" {
		_, err = oleutil.PutProperty(idispatch, "IconLocation", icon)
		if err != nil {
			return err
		}
	}
	_, err = oleutil.CallMethod(idispatch, "Save")
	if err != nil {
		return err
	}
	return nil
}

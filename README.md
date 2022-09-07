# Shortcut

Shortcut is Go (golang) lib for create windows shortcut with a specified name, icon and target directory.

Install

```golang
go get github.com/jxeng/shortcut
```


Usage

```golang

// Create a shortcut on desktop
shortcut.CreateDesktopShortcut("google", "https://google.com", "/path/to/logo.ico", ".")

// Advanced useage
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
err := Create(sc)

```

# Shortcut

Shortcut is Go (golang) lib for create windows shortcut with a specified name, icon and target directory.

Install

```golang
go get github.com/jxeng/shortcut
```


Usage

```golang
// create a shortcut in current directory
shortcut.CreateShortcut("google", "https://google.com", "/path/to/logo.ico", ".")

// create a shortcut on desktop
shortcut.CreateDesktopShortcut("google", "https://google.com", "/path/to/logo.ico", ".")

// create a shortcut for file
shortcut.CreateDesktopShortcut("google", "/path/to/foo.exe", "/path/to/logo.ico")

// icon path could be empty
shortcut.CreateDesktopShortcut("google", "/path/to/foo.exe", "")

```
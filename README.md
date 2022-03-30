# Shortcut

Shortcut is Go (golang) lib for create windows shortcut with a specified name, icon and target directory.

Usage

```golang
// create a shortcut in current directory
CreateShortcut("google", "https://google.com", "/path/to/logo.ico", ".")

// create a shortcut on desktop
CreateDesktopShortcut("google", "https://google.com", "/path/to/logo.ico", ".")

// create a shortcut for file
CreateDesktopShortcut("google", "/path/to/foo.exe", "/path/to/logo.ico")

// icon path could be empty
CreateDesktopShortcut("google", "/path/to/foo.exe", "")

```
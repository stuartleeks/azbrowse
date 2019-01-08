package keybindings

import (
	"os/exec"
)

// OpenEditor opens the specified file in the default editor
func OpenEditor(filename string) error {
	cmd := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", filename)
	return cmd.Run()
}

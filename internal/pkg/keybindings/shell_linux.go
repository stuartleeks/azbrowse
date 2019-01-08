package keybindings

import (
	"os/exec"
)

// OpenEditor opens the specified file in the default editor
func OpenEditor(filename string) error {
	cmd := exec.Command("xdg-open", filename)
	return cmd.Run()
}

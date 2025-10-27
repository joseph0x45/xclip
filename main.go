// provides functions to copy and paste from the clipboard
package xclip

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

var Unsupported bool = false

func init() {
	if _, err := exec.LookPath("xclip"); err != nil {
		log.Println("xclip not installed")
		Unsupported = true
	}
}

func CopyToClipboard(text string) error {
	if Unsupported {
		return errors.New("xclip is not installed")
	}
	cmd := exec.Command("xclip", "-selection", "clipboard")
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

func ReadFromClipboard() (string, error) {
	if Unsupported {
		return "", errors.New("xclip is not installed")
	}
	out, err := exec.Command("xclip", "-selection", "clipboard", "-o").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

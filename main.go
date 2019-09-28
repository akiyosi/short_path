package shortpath

import (
	"fmt"
	"runtime"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

func Minimum(path string) (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	// Shorten homedir
	path = strings.Replace(path, home, "~", 1)

	var out []string
	if runtime.GOOS == "windows" {
		out = strings.Split(path, "\\")
	} else {
		out = strings.Split(path, "/")
	}

	for i := 0; i < len(out)-1; i++ {
		runeout := []rune(out[i])
		if len(runeout) > 0 {
			out[i] = string(runeout[0])
		}
	}

	ret := fmt.Sprintf("%s", strings.Join(out, "/"))
	return ret, nil
}

func PrettyMinimum(path string) (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	// Shorten homedir
	path = strings.Replace(path, home, "~", 1)

	var out []string
	if runtime.GOOS == "windows" {
		out = strings.Split(path, "\\")
	} else {
		out = strings.Split(path, "/")
	}

	if len(out) <= 2 {
		return path, nil
	}

	for i := 0; i < len(out)-1; i++ {
		runeout := []rune(out[i])
		if len(runeout) > 0 {
			out[i] = string(runeout[0])
		}
	}

	return fmt.Sprintf("%s", strings.Join(out, "/")), nil
}

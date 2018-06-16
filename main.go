package shortpath

import (
	"fmt"
	"runtime"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

func Minimum(cur string) (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	// Shorten homedir
	cur = strings.Replace(cur, home, "~", 1)

	var out []string
	if runtime.GOOS == "windows" {
		out = strings.Split(cur, "\\")
	} else {
		out = strings.Split(cur, "/")
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

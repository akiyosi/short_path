package shortpath

import (
	"fmt"
	"strings"
        "runtime"

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
		if len(out[i]) > 0 {
			out[i] = string(out[i][0])
		}
	}

	ret := fmt.Sprintf("%s", strings.Join(out, "/"))
	return ret, nil
}

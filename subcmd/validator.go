package subcmd

import (
	"errors"
	"strings"
)

func validate(args []string, prefixes ...string) error {
	for _, prefix := range prefixes {
		var count int
		for _, v := range args {
			if strings.HasPrefix(v, prefix+"=") {
				count++
			}
			if 1 < count {
				return errors.New(prefix + "は1つまでしかセットできない")
			}
		}
	}
	return nil
}

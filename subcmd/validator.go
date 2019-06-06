package subcmd

import (
	"errors"
	"fmt"
	"strings"
)

func validate(args []string, prefixes ...string) error {
	var found int
	for _, prefix := range prefixes {
		var count int
		for _, v := range args {
			if strings.HasPrefix(v, prefix+"=") {
				found++
				count++
			}
			if 1 < count {
				return errors.New(prefix + "は1つまでしかセットできない")
			}
		}
	}
	if found != len(prefixes) {
		msg := fmt.Sprintf("パラメータが不足しています。 args=%v, prefixes=%v", args, prefixes)
		return errors.New(msg)
	}
	return nil
}

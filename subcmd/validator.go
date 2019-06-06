package subcmd

import "strings"

func validate(args []string, prefixes ...string) {
	for _, prefix := range prefixes {
		var count int
		for _, v := range args {
			if strings.HasPrefix(v, prefix+"=") {
				count++
			}
			if 1 < count {
				panic("widthは1つまで")
			}
		}
	}
}

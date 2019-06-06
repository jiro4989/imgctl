package subcmd

import (
	"strconv"
	"strings"
)

func getParams(args []string, prefixes ...string) (result map[string]int) {
	var err error
	for _, arg := range args {
		for _, prefix := range prefixes {
			kv := strings.Split(arg, "=")
			if kv[0] == prefix {
				result[prefix], err = strconv.Atoi(kv[1])
				if err != nil {
					panic(err)
				}
				break
			}
		}
	}
	return
}

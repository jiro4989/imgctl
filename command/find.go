package command

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
)

func CmdFind(c *cli.Context) {
	target := c.String("d")
	if target == "" {
		log.Fatal("need path.")
	}

	// 検索対象の拡張子を取得
	ext := strings.ToLower(c.String("e"))
	filepath.Walk(target, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		f := strings.ToLower(info.Name())
		if strings.HasSuffix(f, ext) {
			fmt.Println(p)
		}

		return nil
	})
}

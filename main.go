package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jiro4989/imgctl/subcmd"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}

func main() {
	if err := subcmd.RootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

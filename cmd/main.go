package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "jiro4989"
	app.Email = ""
	app.Usage = "Utilitiy to process images for RPG Maker"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

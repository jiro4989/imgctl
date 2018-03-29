package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/jiro4989/tkimgutil/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "scale",
		Usage:  "Scaling images",
		Action: command.CmdScale,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "trim",
		Usage:  "Scaling images",
		Action: command.CmdTrim,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "flip",
		Usage:  "Scaling images",
		Action: command.CmdFlip,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "paste",
		Usage:  "Scaling images",
		Action: command.CmdPaste,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

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
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "size,s",
				Value: 100,
				Usage: "scale size",
			},
			cli.StringFlag{
				Name:  "dist,d",
				Value: "dist/scale",
				Usage: "output dist",
			},
		},
	},
	{
		Name:   "trim",
		Usage:  "Crop images",
		Action: command.CmdTrim,
		Flags: []cli.Flag{
			cli.IntFlag{Name: "x", Value: 0, Usage: "Point"},
			cli.IntFlag{Name: "y", Value: 0, Usage: "Point"},
			cli.IntFlag{Name: "width", Value: 144, Usage: "width"},
			cli.IntFlag{Name: "height", Value: 144, Usage: "height"},
			cli.StringFlag{
				Name:  "dist,d",
				Value: "dist/trim",
				Usage: "output dist",
			},
		},
	},
	{
		Name:   "flip",
		Usage:  "Scaling images",
		Action: command.CmdFlip,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "dist,d",
				Value: "dist/flip",
				Usage: "output dist",
			},
		},
	},
	{
		Name:   "paste",
		Usage:  "Scaling images",
		Action: command.CmdPaste,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "row,r",
				Value: 2,
				Usage: "row count",
			},
			cli.IntFlag{
				Name:  "column,c",
				Value: 4,
				Usage: "column count",
			},
			cli.IntFlag{Name: "width", Value: 144, Usage: "width"},
			cli.IntFlag{Name: "height", Value: 144, Usage: "height"},
			cli.StringFlag{
				Name:  "out-filename-prefix,o",
				Value: "actor",
				Usage: "out filename prefix",
			},
			cli.StringFlag{
				Name:  "padding,p",
				Value: "%03d",
				Usage: "number padding format",
			},
			cli.StringFlag{
				Name:  "dist,d",
				Value: "dist/paste",
				Usage: "output dist",
			},
		},
	},
	{
		Name:   "generate",
		Usage:  "Generate images",
		Action: command.CmdGenerate,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "dist,d",
				Value: "dist/generate",
				Usage: "output dist",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

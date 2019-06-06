package subcmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jiro4989/imgctl/paste"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(pasteCommand)
	pasteCommand.Flags().StringP("outdir", "d", "paste", "Output directory")
	pasteCommand.Flags().StringP("filename-format", "f", "%03d.png", "Out filename format")
}

var pasteCommand = &cobra.Command{
	Use:   "paste",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		f := cmd.Flags()

		outDir, err := f.GetString("outdir")
		if err != nil {
			panic(err)
		}

		filenameFormat, err := f.GetString("filename-format")
		if err != nil {
			panic(err)
		}

		pasted, err := paste.PasteImages(outDir, filenameFormat, args, r, c, w, h)
		if err != nil {
			panic(err)
		}
		for _, file := range pasted {
			fmt.Println(file)
		}
	},
}

func validatePastePatams(args []string) {
	validate(args, "r", "c", "w", "h")
}

func getTileParams(args []string) (x, y, w, h int) {
	var err error
	for _, arg := range args {
		kv := strings.Split(arg, "=")
		switch kv[0] {
		case "x":
			x, err = strconv.Atoi(kv[1])
			if err != nil {
				panic(err)
			}
		case "y":
			y, err = strconv.Atoi(kv[1])
			if err != nil {
				panic(err)
			}
		case "w":
			w, err = strconv.Atoi(kv[1])
			if err != nil {
				panic(err)
			}
		case "h":
			h, err = strconv.Atoi(kv[1])
			if err != nil {
				panic(err)
			}
		}
	}
	return
}

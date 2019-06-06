package subcmd

import (
	"fmt"

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
		validatePasteParams(args)
		r, c, w, h := getTileParams(args)
		l := len(getKeyValueParams(args, "r", "c", "w", "h"))
		args = args[l:]

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

func validatePasteParams(args []string) {
	validate(args, "r", "c", "w", "h")
}

func getTileParams(args []string) (r, c, w, h int) {
	ret := getParams(args, "r", "c", "w", "h")
	r = ret["r"]
	c = ret["c"]
	w = ret["w"]
	h = ret["h"]
	return
}

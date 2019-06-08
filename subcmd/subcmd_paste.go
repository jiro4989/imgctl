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
	pasteCommand.Flags().IntP("row", "r", 0, "Tile row")
	pasteCommand.Flags().IntP("col", "c", 0, "Tile column")
	pasteCommand.Flags().IntP("width", "W", 0, "Width")
	pasteCommand.Flags().IntP("height", "H", 0, "Height")
}

var pasteCommand = &cobra.Command{
	Use:   "paste",
	Short: "Pasting image files",
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

		var (
			r, c, w, h int
		)
		r, err = f.GetInt("row")
		if err != nil {
			panic(err)
		}
		c, err = f.GetInt("col")
		if err != nil {
			panic(err)
		}
		w, err = f.GetInt("width")
		if err != nil {
			panic(err)
		}
		h, err = f.GetInt("height")
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

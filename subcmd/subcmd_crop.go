package subcmd

import (
	"fmt"

	"github.com/jiro4989/imgctl/crop"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(cropCommand)
	cropCommand.Flags().StringP("outdir", "d", "crop", "Output directory")
	cropCommand.Flags().IntP("axis-x", "x", 0, "Axis X")
	cropCommand.Flags().IntP("axis-y", "y", 0, "Axis Y")
	cropCommand.Flags().IntP("width", "W", 0, "Width")
	cropCommand.Flags().IntP("height", "H", 0, "Height")
}

var cropCommand = &cobra.Command{
	Use:   "crop",
	Short: "Croping image files",
	Run: func(cmd *cobra.Command, args []string) {
		f := cmd.Flags()

		outDir, err := f.GetString("outdir")
		if err != nil {
			panic(err)
		}

		var (
			x, y, w, h int
		)
		x, err = f.GetInt("axis-x")
		if err != nil {
			panic(err)
		}
		y, err = f.GetInt("axis-y")
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

		if w <= 0 || h <= 0 {
			msg := fmt.Sprintf("width or height is over 0. w = %d, h = %d", w, h)
			panic(msg)
		}

		cropd, err := crop.CropImages(outDir, args, x, y, w, h)
		if err != nil {
			panic(err)
		}
		for _, file := range cropd {
			fmt.Println(file)
		}
	},
}

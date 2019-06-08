package subcmd

import (
	"fmt"

	"github.com/jiro4989/imgctl/scale"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(scaleCommand)
	scaleCommand.Flags().StringP("outdir", "d", "scale", "Output directory")
	scaleCommand.Flags().IntP("size", "s", 100, "Scale size")
}

var scaleCommand = &cobra.Command{
	Use:   "scale",
	Short: "Scaling image files",
	Run: func(cmd *cobra.Command, args []string) {
		f := cmd.Flags()

		outDir, err := f.GetString("outdir")
		if err != nil {
			panic(err)
		}

		s, err := f.GetInt("size")
		if err != nil {
			panic(err)
		}

		scaled, err := scale.ScaleImages(outDir, args, s)
		if err != nil {
			panic(err)
		}
		for _, file := range scaled {
			fmt.Println(file)
		}
	},
}

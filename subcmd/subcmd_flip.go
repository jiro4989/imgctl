package subcmd

import (
	"fmt"

	"github.com/jiro4989/imgctl/flip"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(flipCommand)
	flipCommand.Flags().StringP("outdir", "d", "flip", "Output directory")
}

var flipCommand = &cobra.Command{
	Use:   "flip",
	Short: "Fliping image files",
	Run: func(cmd *cobra.Command, args []string) {
		f := cmd.Flags()

		outDir, err := f.GetString("outdir")
		if err != nil {
			panic(err)
		}

		flipd, err := flip.FlipImages(outDir, args)
		if err != nil {
			panic(err)
		}
		for _, file := range flipd {
			fmt.Println(file)
		}
	},
}

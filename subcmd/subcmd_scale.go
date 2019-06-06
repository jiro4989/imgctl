package subcmd

import (
	"fmt"

	"github.com/jiro4989/imgctl/scale"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(scaleCommand)
	scaleCommand.Flags().StringP("outdir", "d", "scale", "Output directory")
}

var scaleCommand = &cobra.Command{
	Use:   "scale",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		validateScaleParams(args)
		s := getScaleParams(args)
		f := cmd.Flags()

		outDir, err := f.GetString("outdir")
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

func validateScaleParams(args []string) {
	validate(args, "s")
}

func getScaleParams(args []string) (s int) {
	ret := getParams(args, "s")
	s = ret["s"]
	return
}

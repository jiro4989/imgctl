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
	Short: "Scaling image files",
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateScaleParams(args); err != nil {
			panic(err)
		}
		s := getScaleParams(args)
		l := len(getKeyValueParams(args, "s"))
		args = args[l:]

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

func validateScaleParams(args []string) error {
	return validate(args, "s")
}

func getScaleParams(args []string) (s int) {
	ret, err := getParams(args, "s")
	if err != nil {
		panic(err)
	}
	s = ret["s"]
	return
}

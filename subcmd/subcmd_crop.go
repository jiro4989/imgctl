package subcmd

import (
	"fmt"

	"github.com/jiro4989/imgctl/crop"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(cropCommand)
	cropCommand.Flags().StringP("outdir", "d", "crop", "Output directory")
}

var cropCommand = &cobra.Command{
	Use:   "crop",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateCropParams(args); err != nil {
			panic(err)
		}
		x, y, w, h := getRectParams(args)
		l := len(getKeyValueParams(args, "x", "y", "w", "h"))
		args = args[l:]

		f := cmd.Flags()

		outDir, err := f.GetString("outdir")
		if err != nil {
			panic(err)
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

func validateCropParams(args []string) error {
	return validate(args, "w", "h")
}

func getRectParams(args []string) (x, y, w, h int) {
	ret, err := getParams(args, "x", "y", "w", "h")
	if err != nil {
		panic(err)
	}
	x = ret["x"]
	y = ret["y"]
	w = ret["w"]
	h = ret["h"]
	return
}

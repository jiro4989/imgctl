package subcmd

import (
	"fmt"
	"strconv"
	"strings"

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
		validateCropPatams(args)
		x, y, w, h := getRect(args)

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

func validateCropPatams(args []string) {
	validate(args, "w", "h")
}

func getRect(args []string) (x, y, w, h int) {
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

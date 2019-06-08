package subcmd

import (
	"fmt"
	"strings"

	"github.com/jiro4989/imgctl/generate"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(generateCommand)
	generateCommand.Flags().StringP("outdir", "d", "generate", "Output directory")
	generateCommand.Flags().StringP("filename-format", "f", "%03d.png", "Out filename format")
}

var generateCommand = &cobra.Command{
	Use:   "generate",
	Short: "Generating images from multiple image files",
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

		var patterns [][]string
		for _, arg := range args {
			files := strings.Split(arg, ":")
			patterns = append(patterns, files)
		}

		generated, err := generate.GenerateImages(outDir, filenameFormat, patterns)
		if err != nil {
			panic(err)
		}
		for _, file := range generated {
			fmt.Println(file)
		}
	},
}

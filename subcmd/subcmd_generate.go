package subcmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(generateCommand)
	generateCommand.Flags().StringP("pad", "p", " ", "Padding string")
	generateCommand.Flags().IntP("length", "n", -1, "Padding length")
	generateCommand.Flags().BoolP("write", "w", false, "Overwrite file")
	generateCommand.Flags().StringP("linefeed", "l", "\n", "Line feed")
	generateCommand.Flags().BoolP("termwidth", "t", false, "Terminal width")
}

var generateCommand = &cobra.Command{
	Use:   "generate",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

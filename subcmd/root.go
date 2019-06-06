package subcmd

import (
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

var RootCommand = &cobra.Command{
	Use:   "imgctl",
	Short: "imgctl controls image files.",
	Version: Version + `
Copyright (C) 2019, jiro4989
Released under the MIT License.
https://github.com/jiro4989/imgctl`,
}

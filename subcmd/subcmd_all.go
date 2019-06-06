package subcmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jiro4989/imgctl/all"
	"github.com/jiro4989/imgctl/config"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(allCommand)
}

var allCommand = &cobra.Command{
	Use:   "all",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		for _, configFile := range args {
			b, err := ioutil.ReadFile(configFile)
			if err != nil {
				panic(err)
			}

			var conf config.Config
			if err := json.Unmarshal(b, &conf); err != nil {
				panic(err)
			}

			generated, err := all.RunAll(conf)
			if err != nil {
				panic(err)
			}
			for _, f := range generated {
				fmt.Println(f)
			}
		}
	},
}

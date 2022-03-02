package run

import (
	"github.com/awisu2/gin-study/tryFunctions/app"
	"github.com/spf13/cobra"
)

var params = struct {
	host       string
	port       int
	production bool
}{}

// create command
var Cmd = &cobra.Command{
	Use:   "run",
	Short: "app running",
	Args:  cobra.ArbitraryArgs, // arguments setting.(ArbitraryArgs: any value ok)
	// can chooses `Run()` or `RunE()`
	// args has cli parameters without has -{-str} option.
	//
	// get arg: cmd.Flags().GetInt("num")
	// get persistent arg: cmd.PersistentFlags().GetString("str")
	//
	// show help: `cmd.Help()`
	Run: func(cmd *cobra.Command, args []string) {
		opt := app.AppOptions{
			Host:       params.host,
			Port:       params.port,
			Production: params.production,
		}

		app.Run(&opt)
	},
}

// run initialize like viper
func initConfig() {
}

// execute at running.
func init() {
	// run between parse arguments and run Command.Run()
	cobra.OnInitialize(initConfig)

	// arguments
	flags := Cmd.Flags()
	flags.StringVarP(&params.host, "host", "", "", "host")
	flags.IntVarP(&params.port, "port", "p", 8080, "port")
	flags.BoolVar(&params.production, "production", false, "production")

	// persistent arguments
	// pFlags := Cmd.PersistentFlags()
	// pFlags.StringP("str", "s", "abc", "arg2 string")

	// require settings
	requireds := []string{}
	persistentRequireds := []string{}
	for _, required := range requireds {
		Cmd.MarkFlagRequired(required)
	}
	for _, required := range persistentRequireds {
		Cmd.MarkPersistentFlagRequired(required)
	}

	// add command
	Cmd.AddCommand()
}

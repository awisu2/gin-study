package cmd

import (
	"github.com/awisu2/gin-study/tryFunctions/cmd/run"
	"github.com/spf13/cobra"
)

var params = struct {
}{}

// create command
var Cmd = &cobra.Command{
	Use:   "tryFunctions",
	Short: "gin server",
	Args:  cobra.ArbitraryArgs, // arguments setting.(ArbitraryArgs: any value ok)
	// can chooses `Run()` or `RunE()`
	// args has cli parameters without has -{-str} option.
	//
	// get arg: cmd.Flags().GetInt("num")
	// get persistent arg: cmd.PersistentFlags().GetString("str")
	//
	// show help: `cmd.Help()`
	// Run: func(cmd *cobra.Command, args []string) {
	// 	cmd.Help()
	// },
}

// run initialize like viper
func initConfig() {
}

// execute at running.
func init() {
	// run between parse arguments and run Command.Run()
	cobra.OnInitialize(initConfig)

	// arguments
	// flags := Cmd.Flags()
	// flags.IntVarP(&params.arg1, "num", "n", 99, "arg1 number")

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
	Cmd.AddCommand(run.Cmd)
}

// Execute(): Tipycaly function name
func Execute() error {
	return Cmd.Execute()
}

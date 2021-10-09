package cmd

import (
	"github.com/gaitr/goprobe/internal"
	"github.com/gaitr/goprobe/internal/request"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var client http.Client
var flagPool request.FlagPool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "goprobe",
	Version: internal.VERSION,
	Short:   "goprobe is a CLI app to check the Status of accessible URLs",
	Example: "Enter a URL as below to check Status Code:\ngoprobe https://example.com/",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			internal.Router(client, args, &flagPool)
		}
		os.Exit(1)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolVarP(&flagPool.IsGet, "get", "G", false, "to send get request")
}

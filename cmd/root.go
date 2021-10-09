package cmd

import (
	"fmt"
	"github.com/gaitr/goprobe/internal/request"
	"github.com/spf13/cobra"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string
var client http.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "goprobe",
	Short:   "goprobe is a CLI app to check the Status of accessible URLs",
	Example: "Enter a URL as below to check Status Code:\ngoprobe https://example.com/",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			url := args[0]
			request.DefaultResponse(request.Head(client, url), url)
			//request.CompleteHeadResponse(request.Head(client, url))
		}
		os.Exit(1)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goprobe.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".goprobe" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".goprobe")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

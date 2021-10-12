package cmd

import (
	"errors"
	"github.com/gaitr/goprobe/internal"
	"github.com/gaitr/goprobe/internal/request"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "goprobe",
	Version: internal.VERSION,
	Short:   "goprobe is a CLI app to check the Status of accessible URLs",
	Example: "Enter a URL as below to check Status Code:\ngoprobe https://example.com/",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runCommand(client, args)
		if errors.Is(err, notFoundCommand) {
			cmd.Help()
		}
		return err
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolVarP(&flagPool.IsGet, "get", "G", false, "to send get request")
	rootCmd.Flags().StringVarP(
		&flags.filepath,
		flagsName.file,
		flagsName.fileShort,
		"", "path to the file")
	rootCmd.PersistentFlags().BoolVarP(
		&flags.verbose,
		flagsName.verbose,
		flagsName.verboseShort,
		false, "log verbose output")
}

func runCommand(client http.Client, args []string) error {
	print = logNoop
	if flags.verbose {
		print = logOut
	}

	if isInputFromPipe() {
		return scanList(os.Stdin)
	} else if len(args) > 0 {
		return directly(client, args)
	} else if flags.filepath != "" {
		return fileRead()
	} else {
		return notFoundCommand
	}
}

func directly(client http.Client, args []string) error {
	if len(args) > 0 {
		return internal.Router(client, args[0], &flagPool)
	}
	return errors.New("args not found")
}

var client http.Client
var flagPool request.FlagPool

var flags struct {
	filepath string
	verbose  bool
}
var flagsName = struct {
	file, fileShort       string
	verbose, verboseShort string
}{
	"file", "f",
	"verbose", "v",
}

var print func(s string)

func logNoop(s string) {}

func logOut(s string) { log.Println(s) }

var notFoundCommand = errors.New("command not found")

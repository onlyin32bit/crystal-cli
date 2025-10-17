package cmd

import (
	"crystal-cli/console"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crystal-cli",
	Short: "A brief description of your application",
	Long: strings.TrimSpace(`
	A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.
`),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		console.PrintErrorAndExit(err, 1)
	}
}

func init() {
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true
}

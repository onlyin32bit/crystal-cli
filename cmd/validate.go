package cmd

import (
	"crystal-cli/wd"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:     "validate",
	Aliases: []string{"val"},
	Short:   "Validate Crystal Schema without side effects",
	Long:    ``,
	RunE:    runValidate,
}

func runValidate(cmd *cobra.Command, args []string) error {
	if err := wd.RequireSchemaFile(); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(validateCmd)
}

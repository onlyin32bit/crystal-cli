package db

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db",
		Short: "Sync schema file into database",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	return nil
}

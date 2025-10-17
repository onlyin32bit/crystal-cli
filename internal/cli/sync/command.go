package sync

import (
	"crystal-cli/internal/cli/sync/db"
	"crystal-cli/internal/cli/sync/diff"
	"crystal-cli/internal/cli/sync/schema"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "sync",
	}

	cmd.AddCommand(db.NewCommand())
	cmd.AddCommand(schema.NewCommand())
	cmd.AddCommand(diff.NewCommand())

	return cmd
}

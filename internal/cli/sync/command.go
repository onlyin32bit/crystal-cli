package sync

import (
	"crystal-cli/internal/application/workdir"
	"crystal-cli/internal/cli/sync/db"
	"crystal-cli/internal/cli/sync/diff"
	"crystal-cli/internal/cli/sync/schema"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "sync",
		PersistentPreRunE: persistentPreRun,
	}

	cmd.AddCommand(db.NewCommand())
	cmd.AddCommand(schema.NewCommand())
	cmd.AddCommand(diff.NewCommand())

	return cmd
}

func persistentPreRun(cmd *cobra.Command, args []string) error {
	wd, err := workdir.Explore()
	if err != nil {
		return err
	}
	wd.Load()
	return nil
}

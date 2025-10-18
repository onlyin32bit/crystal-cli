package diff

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "diff",
	}
	return cmd
}

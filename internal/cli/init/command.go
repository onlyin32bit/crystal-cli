package init

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "",
		Long:  "",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	return nil
}

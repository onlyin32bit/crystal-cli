package cmd

import (
	"crystal-cli/console"
	"crystal-cli/input"
	"crystal-cli/wd"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	force bool
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Crystal",
	Long:  `Setup Crystal ORM for your project`,
	RunE:  runInit,
}

func runInit(cmd *cobra.Command, args []string) error {
	if !force && wd.ExistCrystal() {
		return errors.New("your Crystal Schema already exists, you can use `--force` to force re-init (not safe)")
	} else if force && input.YesOrNo(fmt.Sprintf("Re-init Crystal? %s", color.YellowString("(WARNING: All data will be destroy!)")), input.No) == input.No {
		return errors.New("operation cancelled")
	}

	if err := wd.InitCrystal(); err != nil { // TODO: Implement
		console.PrintErrorAndExit(err, 1)
	}
	return nil
}

func init() {
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Force re-init Crystal")
	rootCmd.AddCommand(initCmd)
}

package cmd

import (
	"crystal-cli/console"
	"crystal-cli/input"
	"crystal-cli/wd"

	"github.com/spf13/cobra"
)

var (
	force bool
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Crystal",
	Long:  `Setup Crystal ORM for your Go project`,
	Run:   runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	if !force && wd.ExistCrystal() {
		console.PrintErrorAndExit("Your Crystal Schema already exists")
	} else if !input.YesOrNo("Re-init Crystal? (WARNING: All data will be destroy!)", input.No) {
		return
	}

	if err := wd.InitCrystal(); err != nil { // TODO: Implement
		console.PrintErrorAndExit(err)
	}
}

func init() {
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Force re-init Crystal")
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

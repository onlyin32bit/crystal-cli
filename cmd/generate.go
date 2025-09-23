/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crystal-cli/console"
	"crystal-cli/schema"
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate Crystal artifacts",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRunE: runValidate,
	Run:     runGenerate,
}

func runGenerate(cmd *cobra.Command, args []string) {
	// if err := wd.RequireSchemaFile(); err != nil {
	// 	console.PrintErrorAndExit(err.Error())
	// }
	s, err := schema.LoadSchema("./crystal/schema.yaml")
	if err != nil {
		console.PrintErrorAndExit(err, 1)
	}

	fmt.Print(s.Models["User"].Fields["id"].Type)
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

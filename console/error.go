package console

import (
	"os"

	"github.com/fatih/color"
)

var (
	errorLabelPrinter = color.New(color.FgBlue, color.BgRed).PrintFunc()
	errorTextPrinter  = color.New(color.FgHiRed).PrintfFunc()
)

func PrintError(err interface{}) {
	errorLabelPrinter(" ERROR ")
	errorTextPrinter(" %s\n", err)
}

func PrintErrorAndExit(err interface{}) {
	PrintError(err)
	os.Exit(1)
}

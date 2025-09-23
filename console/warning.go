package console

import (
	"os"

	"github.com/fatih/color"
)

var (
	warningLabelPrinter = color.New(color.FgBlack, color.BgYellow).PrintFunc()
	warningTextPrinter  = color.New(color.FgHiYellow).PrintfFunc()
)

func PrintWarning(err interface{}) {
	warningLabelPrinter(" WARNING ")
	warningTextPrinter(" %s\n", err)
}

func PrintWarningAndExit(err interface{}) {
	PrintWarning(err)
	os.Exit(1)
}

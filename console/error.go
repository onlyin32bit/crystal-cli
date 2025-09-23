package console

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	ErrorLabelPrinter = color.New(color.FgBlue, color.BgRed, color.Bold)
	errorTextPrinter  = color.New(color.FgHiRed)
)

func PrintError(err error) {
	fmt.Printf("%s %s\n", ErrorLabelPrinter.Sprint(" ERROR "), errorTextPrinter.Sprint(err.Error()))
}

func PrintErrorAndExit(err error, code int) {
	PrintError(err)
	os.Exit(code)
}

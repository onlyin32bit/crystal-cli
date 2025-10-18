package console

import (
	"fmt"

	"github.com/fatih/color"
)

func print() {}

func Println(content ...interface{}) {
	fmt.Printf("%s\n", content...)
}

func Info(content ...interface{}) {
	fmt.Printf("%s %s\n",
		color.New(color.FgHiWhite, color.BgCyan, color.Bold).Sprint(" INFO "),
		color.New(color.FgHiRed).Sprint(content...),
	)
}

func Warning(content ...interface{}) {
	fmt.Printf("%s %s\n",
		color.New(color.FgBlack, color.BgYellow, color.Bold).Sprint(" WARNING "),
		color.New(color.FgHiRed).Sprint(content...),
	)
}

func Error(content ...interface{}) {
	fmt.Printf("%s %s\n",
		color.New(color.FgBlack, color.BgRed, color.Bold).Sprint(" ERROR "),
		color.New(color.FgHiRed).Sprint(content...),
	)
}

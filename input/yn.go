package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type YesNo bool

const (
	Yes YesNo = true
	No  YesNo = false
)

func YesOrNo(prompt string, defaultOption YesNo) YesNo {
	reader := bufio.NewReader(os.Stdin)

	s := "y/N"
	if defaultOption == Yes {
		s = "Y/n"
	}
	for {
		fmt.Printf("%s [%s]: ", prompt, s)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			return defaultOption
		}

		switch strings.ToLower(input) {
		case "y", "yes":
			return Yes
		case "n", "no":
			return No
		default:
			fmt.Println("Please answer y/yes or n/no.")
		}
	}
}

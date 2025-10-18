package console

import (
	"github.com/manifoldco/promptui"
)

func InputYesNoSelect(label string) (*bool, error) {
	prompt := promptui.Select{
		Label: label,
		Items: []string{"Yes", "No"},
	}

	_, res, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	choice := true
	if res == "No" {
		choice = false
	}
	return &choice, nil
}

func InputString(label string) (*string, error) {
	prompt := promptui.Prompt{
		Label: label,
	}
	res, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	return &res, nil
}

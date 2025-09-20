package wd

import (
	"crystal-cli/utils"
	"errors"
)

func RequireSchemaFile() error {
	if exists := ExistCrystal(); !exists {
		return errors.New("`crystal` folder is not found, you can try running `crystal-cli init` in your working directory")
	}
	if exists := utils.FileExists("crystal/schema.yaml"); !exists {
		return errors.New("`crystal` folder is available but no `schema.yaml` found")
	}
	return nil
}

func ExistCrystal() bool {
	return utils.FolderExists("crystal")
}

package schema

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Schema struct {
	Provider   string                      `yaml:"provider"`
	Source     string                      `yaml:"source"`
	Enums      map[string][]string         `yaml:"enums"`
	Composites map[string]string           `yaml:"composites"`
	Models     map[string]map[string]Model `yaml:"models"`
}

type Model struct {
	Fields struct {
		Type    string `yaml:"$type"`
		ID      bool   `yaml:"$id,omitempty"`
		Unique  bool   `yaml:"$unique,omitempty"`
		Default string `yaml:"$default,omitempty"`
	} `yaml:",inline"`
}

func LoadSchema(path string) (*Schema, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema: %w", err)
	}

	var s Schema
	if err := yaml.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("failed to parse schema file: %w", err)
	}
	return &s, nil
}

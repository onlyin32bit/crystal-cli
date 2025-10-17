package schema

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Schema struct {
	Provider   string              `yaml:"provider"`
	Source     string              `yaml:"source"`
	Enums      map[string][]string `yaml:"enums,omitempty"`
	Composites map[string]Model    `yaml:"composites,omitempty"`
	Models     map[string]Model    `yaml:"models"`
}

type Model struct {
	Fields map[string]Field `yaml:",inline"`
}

type Field struct {
	Type      string    `yaml:"$type"`
	ID        bool      `yaml:"$id,omitempty"`
	Unique    bool      `yaml:"$unique,omitempty"`
	Default   string    `yaml:"$default,omitempty"`
	UpdatedAt bool      `yaml:"$updatedAt,omitempty"`
	Relation  *Relation `yaml:"$relation,omitempty"`
}

type Relation struct {
	Fields     []string `yaml:"fields,omitempty"`
	References []string `yaml:"references,omitempty"`
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

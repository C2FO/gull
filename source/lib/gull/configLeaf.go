package gull

import "fmt"

type MigrationTarget interface {
	Set(path string, value string) error
	GetEnvironment() string
	GetAll() map[string]string
	Debug()
}

type ConfigLeaf struct {
	Path  string
	Value string
}

type ConfigLeaves struct {
	Entries []ConfigLeaf
}

func NewConfigLeaves() (*ConfigLeaves, error) {
	return &ConfigLeaves{
		Entries: []ConfigLeaf{},
	}, nil
}

func (c *ConfigLeaves) Apply(target MigrationTarget) error {
	if c == nil || c.Entries == nil || len(c.Entries) == 0 {
		return fmt.Errorf("No leaves were found within this migration. Unable to perform an Apply().")
	}
	for _, leaf := range c.Entries {
		err := target.Set(leaf.Path, leaf.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ConfigLeaves) GetValue(path string) (string, error) {
	for _, leaf := range c.Entries {
		if leaf.Path == path {
			return leaf.Value, nil
		}
	}
	return "", fmt.Errorf("No value found at path [%v]", path)
}

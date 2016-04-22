package gull

import "fmt"

type MigrationTarget interface {
	Set(path string, value string) error
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

func (c *ConfigLeaves) Apply(dryRun bool, target MigrationTarget) error {
	if len(c.Entries) == 0 {
		return fmt.Errorf("No leaves were found within this migration. Unable to perform an Apply().")
	}
	for index, leaf := range c.Entries {
		if dryRun {
			fmt.Printf("Applying leaf #%v with contents %+v\n", index, leaf)
		} else {
			err := target.Set(leaf.Path, leaf.Value)
			if err != nil {
				return err
			}
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

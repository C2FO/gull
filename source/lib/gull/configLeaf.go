package gull

import "fmt"

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

func (c *ConfigLeaves) GetValue(path string) (string, error) {
	for _, leaf := range c.Entries {
		if leaf.Path == path {
			return leaf.Value, nil
		}
	}
	return "", fmt.Errorf("No value found at path [%v]", path)
}

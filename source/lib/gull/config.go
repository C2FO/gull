package gull

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ConfigLeaf struct {
	Path  string
	Value string
}

type Config struct {
	Raw          map[string]interface{}
	ConfigLeaves []ConfigLeaf
}

func NewConfigFromJson(source string) (*Config, error) {
	reader := strings.NewReader(source)
	decoder := json.NewDecoder(reader)
	var dest map[string]interface{}
	if err := decoder.Decode(&dest); err != nil {
		return nil, err
	}

	config := &Config{Raw: dest}
	config.recurseReadConfig(config.Raw, "")
	return config, nil
}

func (c *Config) GetPath(path string) (string, error) {
	for _, leaf := range c.ConfigLeaves {
		if leaf.Path == path {
			return leaf.Value, nil
		}
	}
	return "", fmt.Errorf("No value found at path [%v]", path)
}

func (c *Config) recurseReadConfig(node interface{}, path string) {
	mapped, ok := node.(map[string]interface{})
	if !ok {
		stringed := fmt.Sprintf("%v", node)
		c.ConfigLeaves = append(c.ConfigLeaves, ConfigLeaf{Path: path, Value: stringed})
	} else {
		for key, value := range mapped {
			targetPath := path + "/" + key
			//Special case to handle legacy C2FO configs
			if targetPath == "/*" {
				targetPath = "/default"
			}
			c.recurseReadConfig(value, targetPath)
		}
	}
}

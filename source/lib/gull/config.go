package gull

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Config struct {
	Raw    map[string]interface{}
	Leaves *ConfigLeaves
}

func NewConfigFromJson(source string) (*Config, error) {
	reader := strings.NewReader(source)
	decoder := json.NewDecoder(reader)
	var dest map[string]interface{}
	if err := decoder.Decode(&dest); err != nil {
		return nil, err
	}

	leaves, err := NewConfigLeaves()
	if err != nil {
		return nil, err
	}
	config := &Config{
		Raw:    dest,
		Leaves: leaves,
	}
	config.recurseReadConfig(config.Raw, "")
	return config, nil
}

func (c *Config) recurseReadConfig(node interface{}, path string) {
	mapped, ok := node.(map[string]interface{})
	if !ok {
		stringed := fmt.Sprintf("%v", node)
		c.Leaves.Entries = append(c.Leaves.Entries, ConfigLeaf{Path: path, Value: stringed})
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

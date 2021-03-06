package gull

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Config struct {
	Raw    map[string]interface{}
	Leaves *ConfigLeaves
}

func NewConfigFromJson(source string, jsonEncode bool) (*Config, error) {
	reader := strings.NewReader(source)
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
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
	err = config.recurseReadConfig(config.Raw, "", jsonEncode)
	return config, err
}

func (c *Config) recurseReadConfig(node interface{}, path string, jsonEncode bool) error {
	mapped, ok := node.(map[string]interface{})
	if !ok || len(mapped) == 0 {
		value := fmt.Sprintf("%v", node)
		if jsonEncode {
			var jsonBytes bytes.Buffer
			encoder := json.NewEncoder(&jsonBytes)
			err := encoder.Encode(&node)
			if err != nil {
				return err
			}
			value = jsonBytes.String()
		}
		c.Leaves.Entries = append(c.Leaves.Entries, ConfigLeaf{Path: path, Value: value})

	} else {

		for key, value := range mapped {
			targetPath := path + "/" + key
			err := c.recurseReadConfig(value, targetPath, jsonEncode)
			if err != nil {
				return err
			}
			//Special case to handle legacy C2FO configs
			if targetPath == "/*" {
				err := c.recurseReadConfig(value, "/default", jsonEncode)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

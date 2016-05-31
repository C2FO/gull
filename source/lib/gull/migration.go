package gull

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-yaml/yaml"
)

type Migration struct {
	Content    *ConfigLeaves
	Source     string
	Id         string
	NextId     string
	PreviousId string
	Name       string
}

func NewMigrationFromGull(name string, source string) (*Migration, error) {
	migration := newMigration(name)

	err := yaml.Unmarshal([]byte(source), &migration.Content)

	return migration, err
}

func NewMigrationFromGullFile(filePath string) (*Migration, error) {
	name := migrationNameFromPath(filePath)
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return NewMigrationFromGull(name, string(bytes))
}

func NewMigrationFromConfig(name string, config *Config) (*Migration, error) {
	migration := newMigration(name)
	if config != nil && config.Leaves != nil && len(config.Leaves.Entries) > 0 {
		migration.Content = config.Leaves
	}
	return migration, nil
}

func NewMigrationFromConfigFile(filePath string, fileNameIsEnvironment bool, jsonEncode bool) (*Migration, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	config, err := NewConfigFromJson(string(bytes), jsonEncode)
	if err != nil {
		return nil, err
	}
	migration, err := NewMigrationFromConfig("", config)
	if err != nil {
		return nil, err
	}
	if fileNameIsEnvironment {
		environment := strings.Split(filepath.Base(filePath), ".")[0]
		for ii, _ := range migration.Content.Entries {
			migration.Content.Entries[ii].Path = fmt.Sprintf("/%v%v", environment, migration.Content.Entries[ii].Path)
		}
	}
	return migration, nil
}

func GetMigrationNameFromConfigName(filePath string) string {
	id := createId()
	name := migrationNameFromPath(filePath)
	name = strings.Replace(name, ".json", "", 1)
	name = strings.Replace(name, " ", "-", -1)
	return fmt.Sprintf("%v-%v.%v", id, name, "yaml")
}

func (m *Migration) ConvertToYaml() (string, error) {
	yamlBytes, err := yaml.Marshal(m.Content)
	return string(yamlBytes), err

}

func (m *Migration) WriteToFile(filePath string) error {
	rawYaml, err := m.ConvertToYaml()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, []byte(rawYaml), 0644)
}

func migrationNameFromPath(filePath string) string {
	return strings.Replace(strings.Replace(path.Base(filePath), ".yaml", "", 1), ".yml", "", 1)
}

func newMigration(name string) *Migration {
	id := createId()
	if name != "" {
		id = strings.Split(name, "-")[0]
	}
	return &Migration{
		Id:   id,
		Name: name,
	}
}

func createId() string {
	return fmt.Sprintf("%v", time.Now().UnixNano())
}

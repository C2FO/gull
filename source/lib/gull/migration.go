package gull

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/cloudfoundry-incubator/candiedyaml"
)

type Migration struct {
	Content    *ConfigLeaves
	Source     string
	Id         string
	NextId     string
	PreviousId string
}

func NewMigrationFromGull(name string, source string) (*Migration, error) {
	migration := newMigration(name)
	reader := strings.NewReader(source)
	decoder := candiedyaml.NewDecoder(reader)
	err := decoder.Decode(migration)
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

func NewMigrationFromConfigFile(filePath string) (*Migration, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	config, err := NewConfigFromJson(string(bytes))
	if err != nil {
		return nil, err
	}
	return NewMigrationFromConfig("", config)
}

func GetMigrationNameFromConfigName(filePath string) string {
	id := createId()
	name := migrationNameFromPath(filePath)
	return fmt.Sprintf("%v-%v.%v", id, strings.Replace(name, ".json", "", 1), "yaml")
}

func (m *Migration) WriteToFile(filePath string) error {
	target, err := os.Create(filePath)
	if target != nil {
		defer func() { _ = target.Close() }()
	}
	if err != nil {
		return err
	}

	encoder := candiedyaml.NewEncoder(target)
	return encoder.Encode(m.Content)
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
		Id: id,
	}
}

func createId() string {
	return fmt.Sprintf("%v", time.Now().UnixNano())
}

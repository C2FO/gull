package gull

import (
	"strings"

	"github.com/cloudfoundry-incubator/candiedyaml"
)

type Migration struct {
	Up   *ConfigLeaves
	Down *ConfigLeaves
}

func NewMigrationFromYaml(source string) (*Migration, error) {
	migration := &Migration{}
	reader := strings.NewReader(source)
	decoder := candiedyaml.NewDecoder(reader)
	err := decoder.Decode(migration)
	return migration, err
}

func NewMigrationFromConfig(up *Config, down *Config) (*Migration, error) {
	migration := &Migration{}
	if up != nil && up.Leaves != nil && len(up.Leaves.Entries) > 0 {
		migration.Up = up.Leaves
	}
	if down != nil && down.Leaves != nil && len(down.Leaves.Entries) > 0 {
		migration.Down = down.Leaves
	}
	return migration, nil
}

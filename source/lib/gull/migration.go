package gull

import (
	"fmt"
	"strings"

	"github.com/cloudfoundry-incubator/candiedyaml"
)

type Migration struct {
	Up     []ConfigLeaf
	Down   []ConfigLeaf
	DryRun bool
}

func NewMigrationFromYaml(source string, dryRun bool) (*Migration, error) {
	migration := &Migration{
		DryRun: dryRun,
	}
	reader := strings.NewReader(source)
	decoder := candiedyaml.NewDecoder(reader)
	err := decoder.Decode(migration)
	return migration, err
}

func (m *Migration) ApplyUp() error {
	return m.apply(m.Up)
}

func (m *Migration) ApplyDown() error {
	return m.apply(m.Down)
}

func (m *Migration) apply(leaves []ConfigLeaf) error {
	for index, leaf := range leaves {
		if m.DryRun {
			fmt.Printf("Applying leaf #%v with contents %+v\n", index, leaf)
		}
	}
	return nil
}

package gull

import (
	"fmt"
	"sort"
)

type Migrations struct {
	entries []*Migration
	lookup  map[string]*Migration
}

func NewMigrations() *Migrations {
	return &Migrations{
		entries: []*Migration{},
		lookup:  map[string]*Migration{},
	}
}

func (m *Migrations) Add(migration *Migration) error {
	m.entries = append(m.entries, migration)
	if existing, ok := m.lookup[migration.Id]; ok {
		return fmt.Errorf("Unable to add migration [%v].[%+v] conflicts with [%+v]", migration.Id, migration, existing)
	}
	m.lookup[migration.Id] = migration
	m.Sort()
	return nil
}

func (m *Migrations) First() (*Migration, error) {
	if len(m.entries) == 0 {
		return nil, fmt.Errorf("No migrations exist")
	}
	return m.entries[0], nil
}

func (m *Migrations) Last() (*Migration, error) {
	if len(m.entries) == 0 {
		return nil, fmt.Errorf("No migrations exist")
	}
	return m.entries[len(m.entries)-1], nil
}

func (m *Migrations) Get(id string) (*Migration, error) {
	migration, ok := m.lookup[id]
	if !ok {
		return nil, fmt.Errorf("No migration was found with ID [%v]", id)
	}
	return migration, nil
}

func (m *Migrations) Apply(target MigrationTarget) error {
	for _, entry := range m.entries {
		err := entry.Content.Apply(target)
		if err != nil {
			return err
		}
	}
	return nil
}

// Implement the Sort interface
func (m *Migrations) Len() int {
	return len(m.entries)
}
func (m *Migrations) Swap(i, j int) {
	m.entries[i], m.entries[j] = m.entries[j], m.entries[i]
}
func (m *Migrations) Less(i, j int) bool {
	return m.entries[i].Id < m.entries[j].Id
}
func (m *Migrations) Sort() {
	sort.Sort(m)
	previousId := ""
	for ii, entry := range m.entries {
		if ii > 0 {
			previous := m.entries[ii-1]
			previous.NextId = entry.Id
			previousId = previous.Id
		}
		entry.PreviousId = previousId
	}
}

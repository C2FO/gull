package gull

import (
	"fmt"
	"sort"
	"strings"
)

type Migrations struct {
	Entries []*Migration
	Lookup  map[string]*Migration
	logger  ILogger
}

func NewMigrations() *Migrations {
	return &Migrations{
		Entries: []*Migration{},
		Lookup:  map[string]*Migration{},
	}
}

func (m *Migrations) Add(migration *Migration) error {
	m.Entries = append(m.Entries, migration)
	if existing, ok := m.Lookup[migration.Id]; ok {
		return fmt.Errorf("Unable to add migration [%v].[%+v] conflicts with [%+v]", migration.Id, migration, existing)
	}
	m.Lookup[migration.Id] = migration
	m.Sort()
	return nil
}

func (m *Migrations) First() (*Migration, error) {
	if len(m.Entries) == 0 {
		return nil, fmt.Errorf("No migrations exist")
	}
	return m.Entries[0], nil
}

func (m *Migrations) Last() (*Migration, error) {
	if len(m.Entries) == 0 {
		return nil, fmt.Errorf("No migrations exist")
	}
	return m.Entries[len(m.Entries)-1], nil
}

func (m *Migrations) Get(id string) (*Migration, error) {
	migration, ok := m.Lookup[id]
	if !ok {
		return nil, fmt.Errorf("No migration was found with ID [%v]", id)
	}
	return migration, nil
}

func (m *Migrations) Apply(target MigrationTarget) error {
	environments := []string{"default"}
	if target.GetEnvironment() != "default" && target.GetEnvironment() != "" {
		environments = append(environments, target.GetEnvironment())
	}
	tip, err := target.GetMigrationTip()
	tipId := ""
	if err == nil && tip != nil {
		tipId = tip.Id
		target.GetLogger().Info("[%v]/[%v] is currently migrated to [%v]", target.GetApplication(), target.GetEnvironment(), tip.Name)
	}
	last, err := m.Last()
	if err != nil {
		return err
	}
	if last.Id == tipId {
		target.GetLogger().Info("There are no new migrations to apply\n")
		return nil
	}
	for _, environment := range environments {
		err := m.apply(target, environment, tipId)
		if err != nil {
			return err
		}
	}
	migrationState := NewMigrationState(m)
	return target.SetStatus(migrationState)
}

func (m *Migrations) apply(target MigrationTarget, source string, tipId string) error {
	src := fmt.Sprintf("/%v/", source)
	dest := fmt.Sprintf("/%v/", target.GetEnvironment())
	active := target.IsPerformingFullMigration() || tipId == ""
	for _, entry := range m.Entries {
		if active {
			target.GetLogger().Debug("Applying migration [%v] for environment [%v]", entry.Name, source)
			for _, leaf := range entry.Content.Entries {
				if strings.Contains(leaf.Path, src) {
					path := strings.Replace(leaf.Path, src, dest, 1)
					err := target.Set(path, leaf.Value)
					if err != nil {
						return err
					}
				}
			}
		} else {
			if entry.Id == tipId {
				active = true
			}
		}
	}
	if !active {
		return fmt.Errorf("The existing migration tip of [%v] was not found in the local migrations. Unable to migrate.\n", tipId)
	}
	return nil
}

func (m *Migrations) Pop() (*Migration, error) {
	if m.Len() <= 0 {
		return nil, fmt.Errorf("No migrations found, unable to remove the last element.")
	}
	popped := m.Entries[m.Len()-1]
	m.Entries = m.Entries[:m.Len()-1]
	return popped, nil
}

// Implement the Sort interface
func (m *Migrations) Len() int {
	return len(m.Entries)
}

func (m *Migrations) Swap(i, j int) {
	m.Entries[i], m.Entries[j] = m.Entries[j], m.Entries[i]
}

func (m *Migrations) Less(i, j int) bool {
	return m.Entries[i].Id < m.Entries[j].Id
}

func (m *Migrations) Sort() {
	sort.Sort(m)
	previousId := ""
	for ii, entry := range m.Entries {
		if ii > 0 {
			previous := m.Entries[ii-1]
			previous.NextId = entry.Id
			previousId = previous.Id
		}
		entry.PreviousId = previousId
	}
}

package gull

import (
	"os"
	"path/filepath"
)

var migrationTemplate = `
---
entries:
- path: '/default/newkey'
  value: 'first'
- path: '/default/secondkey'
  value: '[second third]'
`

func CreateNewMigrationFile(name string, destination string) (string, error) {
	err := os.MkdirAll(destination, 0755)
	if err != nil {
		return "", err
	}
	migrationPath := filepath.Join(destination, GetMigrationNameFromConfigName(name))
	migration, err := NewMigrationFromGull(name, migrationTemplate)
	if err != nil {
		return "", err
	}
	return migrationPath, migration.WriteToFile(migrationPath)
}

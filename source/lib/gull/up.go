package gull

import (
	"os"
	"path/filepath"
)

type Up struct {
	Environment     string
	SourceDirectory string
	MigrateTarget   MigrationTarget
	Migrations      *Migrations
}

func NewUp(source string, environment string, target MigrationTarget) *Up {
	return &Up{
		Environment:     environment,
		MigrateTarget:   target,
		SourceDirectory: source,
		Migrations:      NewMigrations(),
	}
}

func (u *Up) Migrate() error {
	err := filepath.Walk(u.SourceDirectory, u.IngestFile)
	if err != nil {
		return err
	}
	return u.Migrations.Apply(u.MigrateTarget)
}

func (u *Up) IngestFile(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !f.IsDir() {
		migration, err := NewMigrationFromGullFile(path)
		if err != nil {
			return err
		}
		u.Migrations.Add(migration)
	}
	return nil
}

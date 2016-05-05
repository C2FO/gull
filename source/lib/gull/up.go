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

func NewUp(source string, target MigrationTarget) *Up {
	return &Up{
		Environment:     target.GetEnvironment(),
		MigrateTarget:   target,
		SourceDirectory: source,
		Migrations:      NewMigrations(),
	}
}

func (u *Up) Migrate() error {
	err := u.Ingest()
	if err != nil {
		return err
	}
	return u.Migrations.Apply(u.MigrateTarget)
}

func (u *Up) Ingest() error {
	return filepath.Walk(u.SourceDirectory, u.IngestFile)
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
		err = u.Migrations.Add(migration)
		if err != nil {
			return err
		}
	}
	return nil
}

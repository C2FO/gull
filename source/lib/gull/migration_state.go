package gull

import "time"

type MigrationsSerializable struct {
}

type MigrationState struct {
	Created    time.Time
	Updated    time.Time
	Migrations *Migrations
}

func NewMigrationState(migrations *Migrations) *MigrationState {
	return &MigrationState{
		Created:    time.Now(),
		Updated:    time.Now(),
		Migrations: migrations,
	}
}

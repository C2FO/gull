package gull

import "time"

type MigrationsSerializable struct {
}

type MigrationState struct {
	Created    time.Time
	Migrations *Migrations
}

func NewMigrationState(migrations *Migrations) *MigrationState {
	return &MigrationState{
		Created:    time.Now(),
		Migrations: migrations,
	}
}

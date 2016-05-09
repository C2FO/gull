package gull

import (
	"fmt"
	"strings"
)

type MigrationTarget interface {
	Set(path string, value string) error
	Get(path string) (string, error)
	GetEnvironment() string
	GetApplication() string
	GetAll() map[string]string
	Debug()
	GetStatus() (*MigrationState, error)
	SetStatus(state *MigrationState) error
	DeleteEnvironment() error
	DeleteApplication() error
	IsPerformingFullMigration() bool
	GetMigrationTip() (*Migration, error)
}

type MockMigrationTarget struct {
	Storage        map[string]string
	Application    string
	Environment    string
	MigrationState *MigrationState
}

func NewMockMigrationTarget(application string, environment string) *MockMigrationTarget {
	return &MockMigrationTarget{
		Storage:     map[string]string{},
		Environment: environment,
		Application: application,
	}
}

func (mmt *MockMigrationTarget) Set(path string, value string) error {
	mmt.Storage[mmt.getAppPath(path)] = value
	return nil
}

func (mmt *MockMigrationTarget) Get(path string) (string, error) {
	val, ok := mmt.Storage[mmt.getAppPath(path)]
	if !ok {
		return "", fmt.Errorf("Unable to find path %v in the mock store", mmt.getAppPath(path))
	}
	return val, nil
}

func (mmt *MockMigrationTarget) getAppPath(path string) string {
	return "/" + mmt.Application + path
}

func (mmt *MockMigrationTarget) DeleteEnvironment() error {
	return mmt.remove(mmt.GetApplication() + "/" + mmt.GetEnvironment())
}

func (mmt *MockMigrationTarget) DeleteApplication() error {
	return mmt.remove(mmt.GetApplication())
}

func (mmt *MockMigrationTarget) remove(root string) error {
	targets := []string{}
	for k, _ := range mmt.Storage {
		if strings.Contains(k, root) {
			targets = append(targets, k)
		}
	}
	for _, target := range targets {
		delete(mmt.Storage, target)
	}
	return nil
}

func (mmt *MockMigrationTarget) GetEnvironment() string {
	return mmt.Environment
}

func (mmt *MockMigrationTarget) GetApplication() string {
	return mmt.Application
}

func (mmt *MockMigrationTarget) GetAll() map[string]string {
	return mmt.Storage
}

func (mmt *MockMigrationTarget) Debug() {
	for key, value := range mmt.Storage {
		fmt.Printf("[%v]->[%v]\n", key, value)
	}
}

func (mmt *MockMigrationTarget) SetStatus(state *MigrationState) error {
	mmt.MigrationState = state
	return nil
}

func (mmt *MockMigrationTarget) GetStatus() (*MigrationState, error) {
	return mmt.MigrationState, nil
}

func (mmt *MockMigrationTarget) IsPerformingFullMigration() bool {
	return true
}

func (mmt *MockMigrationTarget) GetMigrationTip() (*Migration, error) {
	if mmt.MigrationState == nil {
		return nil, nil
	}
	return mmt.MigrationState.Migrations.Last()
}

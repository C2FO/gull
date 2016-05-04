package gull

import "fmt"

type MigrationTarget interface {
	Set(path string, value string) error
	Get(path string) (string, error)
	GetEnvironment() string
	GetAll() map[string]string
	Debug()
	GetStatus() (*MigrationState, error)
	SetStatus(state *MigrationState) error
}

type MockMigrationTarget struct {
	Storage        map[string]string
	Environment    string
	MigrationState *MigrationState
}

func NewMockMigrationTarget(environment string) *MockMigrationTarget {
	return &MockMigrationTarget{
		Storage:     map[string]string{},
		Environment: environment,
	}
}

func (mmt *MockMigrationTarget) Set(path string, value string) error {
	mmt.Storage[path] = value
	return nil
}

func (mmt *MockMigrationTarget) Get(path string) (string, error) {
	val, ok := mmt.Storage[path]
	if !ok {
		return "", fmt.Errorf("Unable to find path %v in the mock store", path)
	}
	return val, nil
}

func (mmt *MockMigrationTarget) GetEnvironment() string {
	return mmt.Environment
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

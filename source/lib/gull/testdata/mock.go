package testdata

import "fmt"

type MockMigrationTarget struct {
	Storage     map[string]string
	Environment string
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

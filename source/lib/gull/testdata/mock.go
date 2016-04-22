package testdata

import "fmt"

type MockMigrationTarget struct{}

func (mmt *MockMigrationTarget) Set(path string, value string) error {
	fmt.Printf("Setting leaf [%+v]->[%+v]\n", path, value)
	return nil
}

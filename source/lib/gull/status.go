package gull

import "fmt"

type Status struct {
	MigrationTarget MigrationTarget
}

func NewStatus(target MigrationTarget) *Status {
	return &Status{
		MigrationTarget: target,
	}
}

func (s *Status) Check() error {
	migrationState, err := s.MigrationTarget.GetStatus()
	if err != nil {
		fmt.Println("No status was found in the migration target host")
		return nil
	}
	last, err := migrationState.Migrations.Last()
	if err != nil {
		fmt.Println("No migrations were found for the provided environment")
		return nil
	} else {
		fmt.Printf("Current migration tip is [%v]\n", last.Id)
		fmt.Printf("This environment was migrated at [%v]\n", migrationState.Created)
		fmt.Printf("There are [%v] applied migrations\n", migrationState.Migrations.Count())
	}
	return nil
}

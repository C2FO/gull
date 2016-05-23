package gull

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
		s.MigrationTarget.GetLogger().Info("No status was found in the migration target host")
		return nil
	}
	last, err := migrationState.Migrations.Last()
	if err != nil {
		s.MigrationTarget.GetLogger().Info("No migrations were found for the provided environment")
		return nil
	} else {
		s.MigrationTarget.GetLogger().Info("Current migration tip is [%v]\n", last.Name)
		s.MigrationTarget.GetLogger().Info("This environment was migrated at [%v]\n", migrationState.Created)
		s.MigrationTarget.GetLogger().Info("There are [%v] applied migrations\n", migrationState.Migrations.Len())
	}
	return nil
}

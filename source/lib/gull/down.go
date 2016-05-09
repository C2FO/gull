package gull

type Down struct {
	MigrationTarget MigrationTarget
	Migrations      *Migrations
}

func NewDown(target MigrationTarget) *Down {
	return &Down{
		MigrationTarget: target,
	}
}

func (d *Down) Migrate() error {
	status, err := d.MigrationTarget.GetStatus()
	if err != nil {
		return err
	}
	d.Migrations = status.Migrations
	_, err = status.Migrations.Pop()
	if err != nil {
		return err
	}
	_ = d.MigrationTarget.DeleteEnvironment()
	return status.Migrations.Apply(d.MigrationTarget)
}

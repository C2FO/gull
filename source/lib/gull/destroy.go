package gull

type Destroy struct {
	MigrationTarget MigrationTarget
}

func NewDestroy(target MigrationTarget) *Destroy {
	return &Destroy{
		MigrationTarget: target,
	}
}

func (d *Destroy) Execute() error {
	return d.MigrationTarget.DeleteEnvironment()
}

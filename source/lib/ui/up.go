package ui

import "log"

func (o *Options) Up() {
	log.Printf("TODO: Implement Up subcommand")
	// An 'Up' will walk through all migrations and apply 'default' to the target environment.
	// All migrations will then be walked again, applying any migrations containing the target environment.
}

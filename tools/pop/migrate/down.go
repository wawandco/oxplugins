package migrate

import (
	"context"

	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/pop/v5"
	"github.com/spf13/pflag"
)

type MigrateDown struct {
	*Logger
	migrations packd.Walkable

	connectionName string
	steps          int
	flags          *pflag.FlagSet
}

func (mu MigrateDown) Name() string {
	return "down"
}

func (mu MigrateDown) HelpText() string {
	return "Runs migrations down passed steps, 1 by default"
}

func (mu MigrateDown) ParentName() string {
	return "migrate"
}

// Run will run migrations on the current folder, it will look for the
// migrations folder and attempt to run the migrations using internal
// pop tooling
func (mu *MigrateDown) Run(ctx context.Context, root string, args []string) error {
	pop.SetLogger(mu.Log)

	conn := pop.Connections[mu.connectionName]
	if conn == nil {
		return ErrCouldNotFindConnection
	}

	mig, err := pop.NewMigrationBox(mu.migrations, conn)
	if err != nil {
		return err
	}

	err = mig.Down(mu.steps)
	if err != nil {
		return err
	}

	return nil
}

func (mu *MigrateDown) ParseFlags(args []string) {
	mu.flags = pflag.NewFlagSet(mu.Name(), pflag.ContinueOnError)

	mu.flags.StringVarP(&mu.connectionName, "conn", "", "development", "the name of the connection to use")
	mu.flags.IntVarP(&mu.steps, "steps", "s", 1, "how many migrations to run")
	mu.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (mu *MigrateDown) Flags() *pflag.FlagSet {
	return mu.flags
}

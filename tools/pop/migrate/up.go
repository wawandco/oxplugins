package migrate

import (
	"context"

	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/pop/v5"
	"github.com/spf13/pflag"
)

type MigrateUp struct {
	migrations packd.Walkable

	connectionName string
	steps          int
	flags          *pflag.FlagSet
}

func (mu MigrateUp) Name() string {
	return "up"
}

func (mu MigrateUp) ParentName() string {
	return "migrate"
}

func (mu MigrateUp) HelpText() string {
	return "Runs migrations up passed steps, all by default"
}

// Run will run migrations on the current folder, it will look for the
// migrations folder and attempt to run the migrations using internal
// pop tooling
func (mu *MigrateUp) Run(ctx context.Context, root string, args []string) error {
	conn := pop.Connections[mu.connectionName]
	if conn == nil {
		return ErrCouldNotFindConnection
	}

	mig, err := pop.NewMigrationBox(mu.migrations, conn)
	if err != nil {
		return err
	}

	return mig.Up()
}

func (mu *MigrateUp) ParseFlags(args []string) {
	mu.flags = pflag.NewFlagSet(mu.Name(), pflag.ContinueOnError)
	mu.flags.StringVarP(&mu.connectionName, "conn", "", "development", "the name of the connection to use")
	mu.flags.IntVarP(&mu.steps, "steps", "s", 0, "how many migrations to run")
	mu.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (mu *MigrateUp) Flags() *pflag.FlagSet {
	return mu.flags
}

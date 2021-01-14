package migrate

import (
	"context"
	"errors"

	"github.com/gobuffalo/packd"
	"github.com/wawandco/oxpecker/plugins"
)

var (
	_ plugins.Command        = (*Command)(nil)
	_ plugins.PluginReceiver = (*Command)(nil)

	ErrCouldNotFindConnection = errors.New("could not find connection by name")
	ErrNotEnoughArgs          = errors.New("not enough args, please specify direction p.e ox pop migrate up")
	ErrInvalidInstruction     = errors.New("invalid instruction for migrate command")
)

type Migrator interface {
	plugins.Command

	Direction() string
	SetConn(string)
}

type Command struct {
	subcommands []plugins.Command
}

//HelpText resturns the help Text of build function
func (m Command) HelpText() string {
	return "Uses soda to run pop migrations"
}

func (m *Command) Name() string {
	return "migrate"
}

func (m *Command) ParentName() string {
	return "db"
}

func (m *Command) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		return ErrNotEnoughArgs
	}

	name := args[1]
	for _, mig := range m.subcommands {
		if mig.Name() != name {
			continue
		}

		return mig.Run(ctx, root, args)
	}

	return ErrInvalidInstruction
}

func (m *Command) Subcommands() []plugins.Command {
	return m.subcommands
}

func (m *Command) Receive(pls []plugins.Plugin) {
	for _, pl := range pls {
		c, ok := pl.(plugins.Command)
		if !ok || c.ParentName() != m.Name() {
			continue
		}

		m.subcommands = append(m.subcommands, c)
	}
}

func Plugins(migrations packd.Box) []plugins.Plugin {

	pl := []plugins.Plugin{
		&Command{},
		&MigrateUp{migrations: migrations},
		&MigrateDown{migrations: migrations},
	}

	return pl
}

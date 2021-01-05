package migrate

import (
	"context"
	"errors"

	"github.com/wawandco/oxpecker/plugins"
)

var (
	_ plugins.Command = (*Plugin)(nil)

	ErrCouldNotFindConnection = errors.New("could not find connection by name")
	ErrNotEnoughArgs          = errors.New("not enough args, please specify direction p.e ox pop migrate up")
	ErrInvalidInstruction     = errors.New("invalid instruction for migrate command")
)

type Plugin struct {
	subcommands []plugins.Command
}

//HelpText resturns the help Text of build function
func (m Plugin) HelpText() string {
	return "Runs migrations on the current folder"
}

func (m *Plugin) Name() string {
	return "migrate"
}

func (m *Plugin) ParentName() string {
	return "pop"
}

func (m *Plugin) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		return ErrNotEnoughArgs
	}

	name := args[1]
	for _, mig := range m.subcommands {
		if mig.Name() == name {
			return mig.Run(ctx, root, args)
		}

		// if cm := mig.(plugins.CommandNamer); cm.CommandName() == name {
		// 	return mig.Run(ctx, root, args)
		// }

	}

	return ErrInvalidInstruction
}

func (m *Plugin) Receive(pls []plugins.Plugin) {
	for _, plugin := range pls {
		pl, ok := plugin.(plugins.Command)
		if !ok {
			continue
		}

		if pl.ParentName() != m.Name() {
			continue
		}

		m.subcommands = append(m.subcommands, pl)
	}
}

func (m *Plugin) Subcommands() []plugins.Command {
	return m.subcommands
}

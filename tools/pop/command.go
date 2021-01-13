package pop

import (
	"context"
	"errors"
	"fmt"

	"github.com/wawandco/oxpecker/plugins"
	"github.com/wawandco/oxplugins/tools/pop/migrate"
)

// Ensuring pop.Plugin is a command
var (
	_ plugins.Command        = (*Command)(nil)
	_ plugins.Plugin         = (*Command)(nil)
	_ plugins.HelpTexter     = (*Command)(nil)
	_ plugins.PluginReceiver = (*Command)(nil)
	_ plugins.Subcommander   = (*Command)(nil)

	ErrCommandNotFound = errors.New("subcommand not found")
)

type Command struct {
	// subcommands we will invoke depending on parameters
	// these are filled when Receive is called.
	subcommands []plugins.Command
}

func (p *Command) Name() string {
	return "pop"
}

func (p *Command) ParentName() string {
	return ""
}

//HelpText resturns the help Text of build function
func (b Command) HelpText() string {
	return "provides commands for pop common tasks"
}

func (b *Command) Receive(plugins []plugins.Plugin) {
	for _, plugin := range plugins {
		if mig, ok := plugin.(*migrate.Command); ok {
			b.subcommands = append(b.subcommands, mig)
			continue
		}

		// Other subcommands
	}
}

func (b *Command) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		return ErrCommandNotFound
	}

	for _, cm := range b.subcommands {
		if cm.Name() != args[1] {
			continue
		}

		if fp, ok := cm.(plugins.FlagParser); ok {
			fp.ParseFlags(args[1:])
		}

		return cm.Run(ctx, root, args[1:])
	}

	fmt.Printf("did not find %v subcommand, try ox help pop.\n", args[1])
	return nil
}

func (b *Command) Subcommands() []plugins.Command {
	return b.subcommands
}

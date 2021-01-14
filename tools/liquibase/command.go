package liquibase

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/wawandco/oxpecker/plugins"
)

var _ plugins.Command = (*Command)(nil)
var _ plugins.HelpTexter = (*Command)(nil)

type Command struct {
	connectionName string
	connections    map[string]URLProvider
	flags          *pflag.FlagSet
}

func (lb Command) Name() string {
	return "migrate"
}

func (lb Command) ParentName() string {
	return "db"
}

func (lb Command) HelpText() string {
	return "runs Liquibase command to update database specified with --conn flag"
}

func (lb *Command) ParseFlags(args []string) {
	lb.flags = pflag.NewFlagSet(lb.Name(), pflag.ContinueOnError)
	lb.flags.StringVarP(&lb.connectionName, "conn", "", "development", "the name of the connection to use")
	lb.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (lb *Command) Flags() *pflag.FlagSet {
	return lb.flags
}

func (lb *Command) Run(ctx context.Context, root string, args []string) error {
	fmt.Println(args)
	return lb.update(lb.connectionName)
}

func (lb *Command) RunBeforeTest(ctx context.Context, root string, args []string) error {
	return lb.update("test")
}

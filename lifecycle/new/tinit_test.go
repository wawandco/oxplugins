package new_test

import (
	"context"

	"github.com/spf13/pflag"
	"github.com/wawandco/oxplugins/lifecycle/new"
)

// For testing purposes
var int = 0

var _ new.Initializer = (*Tinit)(nil)
var _ new.AfterInitializer = (*Tinit)(nil)

type Tinit struct {
	afterCalled bool
	called      bool
}

func (t Tinit) Name() string { return "tinit" }

func (t *Tinit) Initialize(ctx context.Context, root string, args []string) error {
	t.called = true
	return nil
}

func (t *Tinit) AfterInitialize(ctx context.Context, root string, args []string) error {
	t.afterCalled = true
	return nil
}

func (t *Tinit) ParseFlags([]string) {}
func (t *Tinit) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("tinit", pflag.ContinueOnError)
}

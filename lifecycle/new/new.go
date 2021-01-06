package new

import "github.com/wawandco/oxpecker/plugins"


package dev

import (
	"context"
	"fmt"
	"sync"

	"github.com/wawandco/oxpecker/plugins"
)

var _ plugins.Command = (*Command)(nil)
var _ plugins.FlagParser = (*Command)(nil)

// Command to generate New applications.
type Command struct {
	initializers []Initializer
	afterInitializers []AfterInitializer
}

func (d Command) Name() string {
	return "new"
}

func (d Command) ParentName() string {
	return ""
}

//HelpText resturns the help Text of build function
func (d Command) HelpText() string {
	return "Generates a new app with registered plugins"
}

// Run calls NPM or yarn to start webpack watching the assets
// Also starts refresh listening for the changes in Go files.
func (d *Command) Run(ctx context.Context, root string, args []string) error {
	for _, ini := range d.initializers {
		err := ini.Initialize(ctx, root)
		if err != nil {
			return err
		}
	}

	for _, aini := range d.afterInitializers {
		err := aini.AfterInitialize(ctx, root)
		if err != nil {
			return err
		}
	}

	return nil
}
// Receive and store initializers
func (d *Command) Receive(plugins []plugins.Plugin) {
	for _, tool := range plugins {
		i, ok := tool.(Initializer)
		if ok {
			d.developers = append(d.initializers, i)
		}

		ai, ok = tool.(AfterInitializer)
		if ok {
			d.developers = append(d.afterInitializers, ai)
		}
	}
}
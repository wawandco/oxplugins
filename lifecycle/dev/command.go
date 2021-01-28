package dev

import (
	"context"
	"fmt"
	"sync"

	"github.com/wawandco/oxplugins/plugins"
)

var _ plugins.Command = (*Command)(nil)

// Command is the dev command, it runs the dev plugins, each one on a different
// go routine. the detail to what happen on each of these plugins is up to
// each of the Developer plugins.
type Command struct {
	developers []Developer
	beforeDevs []BeforeDeveloper
}

func (d Command) Name() string {
	return "dev"
}

func (d Command) ParentName() string {
	return ""
}

//HelpText returns the help Text of build function
func (d Command) HelpText() string {
	return "calls NPM or yarn to start webpack watching the assetst"
}

// Run calls each of the beforedeveloper plugins and then
// executes Developer plugins in parallel.
func (d *Command) Run(ctx context.Context, root string, args []string) error {
	for _, bd := range d.beforeDevs {
		err := bd.BeforeDevelop(ctx, root)
		if err != nil {
			return err
		}
	}

	var wg sync.WaitGroup
	for _, tool := range d.developers {
		// Each of the tools runs in parallel
		wg.Add(1)
		go func(t Developer) {
			err := t.Develop(ctx, root)
			if err != nil {
				fmt.Println(err)
			}

			wg.Done()
		}(tool)
	}

	wg.Wait()
	return nil
}

// Receive Developer and BeforeDeveloper plugins and store these
// in the Command to be used when the command is invoked.
func (d *Command) Receive(plugins []plugins.Plugin) {
	for _, tool := range plugins {
		if ptool, ok := tool.(Developer); ok {
			d.developers = append(d.developers, ptool)
		}

		if bdev, ok := tool.(BeforeDeveloper); ok {
			d.beforeDevs = append(d.beforeDevs, bdev)
		}
	}
}

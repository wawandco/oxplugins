package oxplugins

import (
	"github.com/wawandco/oxpecker/plugins"
	"github.com/wawandco/oxplugins/lifecycle/build"
	"github.com/wawandco/oxplugins/lifecycle/cli"
	"github.com/wawandco/oxplugins/lifecycle/dev"
	"github.com/wawandco/oxplugins/lifecycle/fix"
	"github.com/wawandco/oxplugins/lifecycle/generate"
	"github.com/wawandco/oxplugins/lifecycle/new"
	"github.com/wawandco/oxplugins/lifecycle/test"
	"github.com/wawandco/oxplugins/tools/envy"
	"github.com/wawandco/oxplugins/tools/flect"
	"github.com/wawandco/oxplugins/tools/grift"
	"github.com/wawandco/oxplugins/tools/ox"
	"github.com/wawandco/oxplugins/tools/pop"
	"github.com/wawandco/oxplugins/tools/refresh"
	"github.com/wawandco/oxplugins/tools/standard"
	"github.com/wawandco/oxplugins/tools/webpack"
	"github.com/wawandco/oxplugins/tools/yarn"
)

// Base plugins for Wawandco applications lifecycle. While oxplugins
// has other plugins this list is the base that is used across most of
// the apps we do. Feel free to add the rest in your cmd/ox/main.go file.
var Base = []plugins.Plugin{
	// Tools plugins.
	&webpack.Plugin{},
	&refresh.Plugin{},
	&yarn.Plugin{},

	// Developer Lifecycle plugins
	&build.Command{},
	&dev.Command{},
	&test.Command{},
	&fix.Command{},
	&generate.Command{},
	&new.Command{},
	&grift.Command{},

	// Builders
	&cli.Builder{},
	&standard.Builder{},

	// Fixers
	&standard.Fixer{},

	// Generators
	&ox.Generator{},

	// Initializer
	&refresh.Initializer{},
	&flect.Initializer{},

	// Testers
	&standard.Tester{},
	&envy.Tester{},
	&pop.Tester{},
}

package oxplugins

import (
	"github.com/wawandco/oxplugins/lifecycle/build"
	"github.com/wawandco/oxplugins/lifecycle/dev"
	"github.com/wawandco/oxplugins/lifecycle/fix"
	"github.com/wawandco/oxplugins/lifecycle/generate"
	"github.com/wawandco/oxplugins/lifecycle/new"
	"github.com/wawandco/oxplugins/lifecycle/test"
	"github.com/wawandco/oxplugins/plugins"
	"github.com/wawandco/oxplugins/tools/buffalo/action"
	"github.com/wawandco/oxplugins/tools/buffalo/app"
	"github.com/wawandco/oxplugins/tools/buffalo/cmd"
	"github.com/wawandco/oxplugins/tools/buffalo/embedded"
	"github.com/wawandco/oxplugins/tools/buffalo/folders"
	"github.com/wawandco/oxplugins/tools/buffalo/model"
	"github.com/wawandco/oxplugins/tools/buffalo/template"
	"github.com/wawandco/oxplugins/tools/cli/help"
	"github.com/wawandco/oxplugins/tools/envy"
	"github.com/wawandco/oxplugins/tools/grift"
	"github.com/wawandco/oxplugins/tools/node"
	"github.com/wawandco/oxplugins/tools/ox"
	"github.com/wawandco/oxplugins/tools/refresh"
	"github.com/wawandco/oxplugins/tools/standard"
	"github.com/wawandco/oxplugins/tools/webpack"
	"github.com/wawandco/oxplugins/tools/yarn"
)

// Base plugins for Wawandco applications lifecycle. While oxplugins
// has other plugins this list is the base that is used across most of
// the apps we do. Feel free to add the rest in your cmd/ox/main.go file.
var Base = []plugins.Plugin{
	// Help
	&help.Command{},

	// Builders
	&build.Command{},
	&node.Builder{},
	&standard.Builder{},

	// Generators
	&generate.Command{},
	&ox.Generator{},
	&template.Generator{},
	&model.Generator{},
	&action.Generator{},

	// Initializers
	&new.Command{},
	&folders.Initializer{},
	&refresh.Initializer{},
	&model.Initializer{},
	&embedded.Initializer{},
	&template.Initializer{},
	&cmd.Initializer{},
	&app.Initializer{},

	// Testers
	&test.Command{},
	&standard.Tester{},
	&envy.Tester{},

	// Fixers
	&fix.Command{},
	&standard.Fixer{},

	// Developer Lifecycle plugins
	&dev.Command{},
	&envy.Developer{},

	// Tools plugins.
	&webpack.Plugin{},
	&refresh.Plugin{},
	&yarn.Plugin{},

	// &flect.Initializer{},
	// &docker.Initializer{},

	// Other
	&grift.Command{},
}
